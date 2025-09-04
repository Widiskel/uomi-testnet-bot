package chain

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/service"
	"github.com/widiskel/uomi-testnet-bot/internal/config"
	"github.com/widiskel/uomi-testnet-bot/internal/domain/model"
	"github.com/widiskel/uomi-testnet-bot/internal/platform/logger"
	"github.com/widiskel/uomi-testnet-bot/pkg/utils"
)

type EthersClient struct {
	client  *ethclient.Client
	network config.Network
	session *model.Session
	log     *logger.ClassLogger
	cfg     config.Config

	TransactionService *service.TransactionService
}

func New(session *model.Session, cfg config.Config, network config.Network) (*EthersClient, error) {
	scope := "[New EtherClient] Error :"
	ec := &EthersClient{network: network, session: session}
	ec.cfg = cfg
	ec.log = logger.NewLogger(ec, session)
	ec.log.Log(fmt.Sprintf("Initializing Ethers Client on %s...", network.Name))

	client, err := ethclient.Dial(network.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("%s failed to connect RPC (%s): %w", scope, network.Name, err)
	}
	txService := service.NewTransactionService(client, session, &network)

	ec.TransactionService = txService
	ec.client = client
	return ec, nil
}

func (e *EthersClient) Close() {
	if e.client != nil {
		e.client.Close()
	}
}

func (e *EthersClient) ConnectWallet() error {
	scope := "[ConnectWallet] Error :"
	data := strings.TrimSpace(e.session.Account)
	if data == "" {
		e.session.Address = ""
		return fmt.Errorf("%s invalid account input (seed or private key)", scope)
	}

	e.log.Log(fmt.Sprintf("Connecting to Account : %d", e.session.AccIdx+1))

	typ := utils.DetermineType(data)
	var addr common.Address
	var privateKey *ecdsa.PrivateKey

	switch typ {
	case "Secret Phrase":
		a, pk, err := utils.AddressFromMnemonic(data, "")
		if err != nil {
			e.session.Address = ""
			return fmt.Errorf("%s failed to read from seed phrase: %w", scope, err)
		}
		addr = a
		privateKey = pk
	case "Private Key":
		pk, err := utils.PrivateKeyFromHex(data)
		if err != nil {
			e.session.Address = ""
			return fmt.Errorf("%s invalid private key: %w", scope, err)
		}
		addr = crypto.PubkeyToAddress(pk.PublicKey)
		privateKey = pk
	default:
		e.session.Address = ""
		return fmt.Errorf("%s invalid account: Secret Phrase or Private Key required", scope)
	}

	e.session.Address = addr.Hex()
	e.session.PublicKey = addr
	e.session.PrivateKey = privateKey
	e.log.Log(fmt.Sprintf("Wallet connected %s", e.session.Address))
	return nil
}

func (e *EthersClient) GetWalletBalance(update ...bool) error {
	return e.TransactionService.GetWalletBalance(update...)
}

func (e *EthersClient) TransferNative(to ...common.Address) error {
	dest := e.session.PublicKey
	if len(to) > 0 {
		dest = to[0]
	}
	amount, amountInWei, err := utils.GenerateRandomAmount(e.cfg.TxAmountMin, e.cfg.TxAmountMax, e.network.Decimals)
	if err != nil {
		return fmt.Errorf("failed to generate random amount: %w", err)
	}
	e.log.Log(fmt.Sprintf("Preparing to Transfer %s %s \nto %s", amount, e.network.Symbol, dest.Hex()), 2000)
	e.TransactionService.TransferNative(e.session.PublicKey, amountInWei)
	e.log.Log(fmt.Sprintf("Successfully Transfer %s %s \nto %s", amount, e.network.Symbol, dest.Hex()), 2000)
	return nil
}

func (e *EthersClient) DepositToken(wrappedToken contracts.Token, amount string, amountInWei *big.Int) error {

	e.log.Log(fmt.Sprintf("Preparing to deposit %s %s", amount, e.network.Symbol), 2000)
	e.TransactionService.Deposit(wrappedToken, amountInWei)
	e.log.Log(fmt.Sprintf("Successfully wrapped %s %s", amount, e.network.Symbol), 2000)
	return nil
}

func (e *EthersClient) WithdrawToken(wrappedToken contracts.Token) error {
	amountInWei, err := e.TransactionService.GetTokenBalance(wrappedToken)
	if err != nil {
		return fmt.Errorf("failed to get %s balance: %w", wrappedToken.Symbol, err)
	}

	if amountInWei.Cmp(big.NewInt(0)) <= 0 {
		e.log.Log(fmt.Sprintf("No %s to withdraw.", wrappedToken.Symbol), 3000)
		return nil
	}
	amount := utils.FormatUnits(amountInWei, wrappedToken.Decimal)

	e.log.Log(fmt.Sprintf("Preparing to withdraw %s %s", amount, wrappedToken.Symbol), 2000)
	e.TransactionService.Withdraw(wrappedToken, amountInWei)
	e.log.Log(fmt.Sprintf("Successfully unwrapped %s %s", amount, e.network.Symbol), 2000)
	return nil
}

func (e *EthersClient) SwapExactInput(swapParam model.SwapModel) error {
	e.log.Log(fmt.Sprintf("Preparing to Swap Exact Input \n%s %s to ? %s", swapParam.Amount, swapParam.FromToken.Symbol, swapParam.ToToken.Symbol), 2000)
	_, err := e.TransactionService.SwapExactInputSingle(&swapParam)
	if err != nil {
		return fmt.Errorf("failed to swap: %w", err)
	}

	return nil
}

func (e *EthersClient) SwapUniversal(swapParam model.SwapModel) error {
	e.log.Log(fmt.Sprintf("Preparing to Swap Universal \n%s %s to ? %s", swapParam.Amount, swapParam.FromToken.Symbol, swapParam.ToToken.Symbol), 2000)
	_, err := e.TransactionService.SwapUniversal(&swapParam)
	if err != nil {
		return fmt.Errorf("failed to swap: %w", err)
	}

	return nil
}

func (e *EthersClient) ClearBalance(swapParam model.SwapModel) error {
	fromTokenBalance, _ := e.TransactionService.GetTokenBalance(swapParam.FromToken)
	swapParam.Amount = utils.FormatUnits(fromTokenBalance, swapParam.FromToken.Decimal)
	swapParam.AmountInWei = fromTokenBalance
	e.log.Log(fmt.Sprintf("Swap Back %s %s to ? %s", utils.FormatUnits(fromTokenBalance, swapParam.FromToken.Decimal), swapParam.FromToken.Symbol, swapParam.ToToken.Symbol), 2000)
	_, err := e.TransactionService.SwapUniversal(&swapParam)
	if err != nil {
		return fmt.Errorf("failed to swap: %w", err)
	}

	return nil
}
