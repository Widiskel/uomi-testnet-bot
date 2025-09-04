package dodo

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/http"
	"github.com/widiskel/uomi-testnet-bot/internal/config"
	"github.com/widiskel/uomi-testnet-bot/internal/domain/model"
	dodoModel "github.com/widiskel/uomi-testnet-bot/internal/domain/model/dodo"
	"github.com/widiskel/uomi-testnet-bot/internal/platform/logger"
	"github.com/widiskel/uomi-testnet-bot/pkg/utils"
)

type DodoClient struct {
	Ec      *chain.EthersClient
	ac      *http.APIClient
	session *model.Session
	network config.Network
	log     *logger.ClassLogger
	apiKey  string
}

func New(ec *chain.EthersClient, ac *http.APIClient, network config.Network, session *model.Session) (*DodoClient, error) {
	dc := &DodoClient{Ec: ec, ac: ac, session: session, network: network, apiKey: "a37546505892e1a952"}
	dc.log = logger.NewLogger(dc, session)
	dc.log.Log(fmt.Sprintf("Initializing Dodo Client on %s...", dc.network.Name), 0)
	return dc, nil
}

func (dc *DodoClient) getRoute(from contracts.Token, to contracts.Token, amount *big.Int) (*dodoModel.DodoRouteData, error) {
	params := dodoModel.DodoRouteParams{
		ChainID:          strconv.Itoa(dc.network.ChainID),
		DeadLine:         strconv.FormatInt(time.Now().Add(10*time.Minute).Unix(), 10),
		APIKey:           dc.apiKey,
		Slippage:         "0.5",
		Source:           "dodoV2AndMixWasm",
		ToTokenAddress:   to.Address.Hex(),
		FromTokenAddress: from.Address.Hex(),
		UserAddr:         dc.session.PublicKey.Hex(),
		EstimateGas:      "false",
		FromAmount:       amount.String(),
	}
	dc.log.LogObject("Route Param", params)

	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	queryString := v.Encode()
	fullURL := "https://api.dodoex.io/route-service/v2/widget/getdodoroute?" + queryString

	responseData, err := dc.ac.Fetch(fullURL, nil)
	if err != nil {
		if httpErr, ok := err.(*http.HTTPError); ok {
			errorBody := string(httpErr.Body)
			if strings.Contains(strings.ToLower(errorBody), "no route") {
				return nil, fmt.Errorf("DODO API returned final error: no route found")
			} else {
				return dc.getRoute(from, to, amount)
			}
		}
		return nil, fmt.Errorf("failed to get DODO route %v", err)
	}

	jsonBytes, err := json.Marshal(responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DODO route response: %w", err)
	}

	var dodoResponse dodoModel.DodoAPIResponse
	if err := json.Unmarshal(jsonBytes, &dodoResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal DODO route response: %w", err)
	}

	if dodoResponse.Status != 200 {
		var errorData string
		if err := json.Unmarshal(dodoResponse.Data, &errorData); err == nil {
			return nil, fmt.Errorf("DODO API error: %s", errorData)
		}
		return nil, fmt.Errorf("DODO API returned non-200 status: %d", dodoResponse.Status)
	}

	var routeData dodoModel.DodoRouteData
	if err := json.Unmarshal(dodoResponse.Data, &routeData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal successful route data: %w", err)
	}

	return &routeData, nil
}

func (dc *DodoClient) Swap(swapParam *model.SwapModel) error {
	if err := dc.Ec.TransactionService.CheckAndFullFillBalance(swapParam); err != nil {
		return err
	}

	dc.Ec.TransactionService.ApproveTokenSpend(swapParam.FromToken, swapParam.RouterAddress, swapParam.AmountInWei)
	dc.log.Log(fmt.Sprintf("Preparing to Swap Exact Input \n%s %s to ? %s", swapParam.Amount, swapParam.FromToken.Symbol, swapParam.ToToken.Symbol), 2000)

	routeData, err := dc.getRoute(swapParam.FromToken, swapParam.ToToken, swapParam.AmountInWei)
	if err != nil {
		return fmt.Errorf("failed to get route data: %w", err)
	}

	txValue := big.NewInt(0)
	if swapParam.FromToken.Native {
		txValue = swapParam.AmountInWei
	}

	tx, err := dc.Ec.TransactionService.BuildTransaction(swapParam.RouterAddress, txValue, common.FromHex(routeData.Data))
	if err != nil {
		return err
	}
	dc.log.Log(fmt.Sprintf("Swapping \n%s %s to %v %s", swapParam.Amount, swapParam.FromToken.Symbol, routeData.ResAmount, swapParam.ToToken.Symbol), 2000)

	signedTx, err := dc.Ec.TransactionService.ExecuteTransaction(tx)
	if err != nil {
		return err
	}
	_, err = dc.Ec.TransactionService.GetTransactorResult(signedTx)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DodoClient) ClearBalance(swapParam model.SwapModel) error {
	fromTokenBalance, _ := dc.Ec.TransactionService.GetTokenBalance(swapParam.FromToken)
	dc.log.Log(fmt.Sprintf("Swap Back %s %s to ? %s", utils.FormatUnits(fromTokenBalance, swapParam.FromToken.Decimal), swapParam.FromToken.Symbol, swapParam.ToToken.Symbol), 2000)

	routeData, err := dc.getRoute(swapParam.FromToken, swapParam.ToToken, fromTokenBalance)
	if err != nil {
		return fmt.Errorf("failed to get route data: %w", err)
	}

	dc.log.Log(fmt.Sprintf("Swapping Back \n%s %s to %v %s", utils.FormatUnits(fromTokenBalance, swapParam.FromToken.Decimal), swapParam.FromToken.Symbol, routeData.ResAmount, swapParam.ToToken.Symbol), 2000)

	txValue := big.NewInt(0)
	if swapParam.FromToken.Native {
		txValue = swapParam.AmountInWei
	}
	tx, err := dc.Ec.TransactionService.BuildTransaction(swapParam.RouterAddress, txValue, common.FromHex(routeData.Data))
	if err != nil {
		return err
	}

	signedTx, err := dc.Ec.TransactionService.ExecuteTransaction(tx)
	if err != nil {
		return err
	}
	_, err = dc.Ec.TransactionService.GetTransactorResult(signedTx)
	if err != nil {
		return err
	}
	return nil
}
