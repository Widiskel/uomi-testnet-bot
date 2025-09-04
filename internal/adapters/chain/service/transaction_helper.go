package service

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	apitypes "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts/token"
	uniswapv3 "github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts/uniswap/v3"

	"github.com/widiskel/uomi-testnet-bot/internal/domain/model"
	uniswapV3Model "github.com/widiskel/uomi-testnet-bot/internal/domain/model/uniswapv3"
	"github.com/widiskel/uomi-testnet-bot/pkg/utils"
)

func (s *TransactionService) GetNativeBalance() (*big.Int, error) {
	address := common.HexToAddress(s.Session.Address)
	return s.Client.BalanceAt(context.Background(), address, nil)
}

func (s *TransactionService) GetTokenBalance(contract contracts.Token) (*big.Int, error) {
	tokenContract, _ := token.NewERC20Caller(contract.Address, s.Client)
	return tokenContract.BalanceOf(&bind.CallOpts{Context: context.Background()}, s.Session.PublicKey)
}

func (s *TransactionService) approve(tokenToApprove contracts.Token, spender common.Address) (*types.Receipt, error) {
	abi, err := token.ERC20MetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("could not get erc20 abi: %v", err))
	}

	calldata, err := abi.Pack("approve", spender, utils.MaxUint256())
	if err != nil {
		panic(fmt.Sprintf("failed to pack arguments for approve: %v", err))
	}

	s.Log.Log(fmt.Sprintf("Approving %s for %s...", tokenToApprove.Symbol, spender.Hex()), 0)
	tx, err := s.BuildTransaction(tokenToApprove.Address, big.NewInt(0), calldata)
	if err != nil {
		return nil, err
	}

	signedTx, err := s.ExecuteTransaction(tx)
	if err != nil {
		return nil, err
	}

	return s.GetTransactorResult(signedTx)
}

func (s *TransactionService) ApproveTokenSpend(tokenToApprove contracts.Token, spender common.Address, amount *big.Int) error {
	s.Log.Log(fmt.Sprintf("Checking allowance for %s...", tokenToApprove.Symbol), 500)
	tokenContract, err := token.NewERC20Caller(tokenToApprove.Address, s.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate token contract: %w", err)
	}

	allowance, err := tokenContract.Allowance(&bind.CallOpts{}, s.Session.PublicKey, spender)
	if err != nil {
		return fmt.Errorf("failed to get allowance: %w", err)
	}

	if allowance.Cmp(amount) < 0 {
		s.Log.Log(fmt.Sprintf("Allowance too low (%s). Approving %s to spend...", utils.FormatUnits(allowance, tokenToApprove.Decimal), tokenToApprove.Symbol), 0)

		receipt, err := s.approve(tokenToApprove, spender)
		if err != nil {
			return fmt.Errorf("failed to approve token: %w", err)
		}

		s.Log.Log(fmt.Sprintf("Approval successful! Tx: %s", receipt.TxHash.Hex()), 500)

	} else {
		s.Log.Log(fmt.Sprintf("Sufficient allowance found for %s.", tokenToApprove.Symbol), 0)
	}

	return nil
}

func (s *TransactionService) CheckAndFullFillBalance(swapParam *model.SwapModel) error {
	s.Log.Log(fmt.Sprintf("Checking balance for %s...", swapParam.FromToken.Symbol), 0)
	currentBalance, err := s.GetTokenBalance(swapParam.FromToken)
	if err != nil {
		return fmt.Errorf("failed to get FromToken balance: %w", err)
	}

	if currentBalance.Cmp(swapParam.AmountInWei) < 0 {
		s.Log.Log(fmt.Sprintf("Balance insufficient for %s. Executing fulfillBalance callback...", swapParam.FromToken.Symbol), 0)
		if err := swapParam.FulfillBalance(); err != nil {
			return fmt.Errorf("fulfillBalance callback failed: %w", err)
		}
		s.Log.Log("Balance now fulfilled.", 1000)
	}
	return nil
}

func (s *TransactionService) prepareDepositData() ([]byte, error) {
	abi, _ := token.WETH9MetaData.GetAbi()
	return abi.Pack("deposit")
}

func (s *TransactionService) prepareWithdrawData(amount *big.Int) ([]byte, error) {
	abi, _ := token.WETH9MetaData.GetAbi()
	return abi.Pack("withdraw", amount)
}

func (s *TransactionService) getPoolInfo(data *model.SwapModel) (poolAddress common.Address, poolFee *big.Int, err error) {
	s.Log.Log("Getting Swap Pool Information ...", 100)
	quoterContract, _ := uniswapv3.NewUniswapV3QuoterV2Caller(data.QuoterAddress, s.Client)
	factoryAddress, _ := quoterContract.Factory(&bind.CallOpts{Context: context.Background()})
	factory, _ := uniswapv3.NewUniswapV3FactoryCaller(factoryAddress, s.Client)
	feeTiers := []*big.Int{
		big.NewInt(500),
		big.NewInt(3000),
		big.NewInt(10000),
	}
	fromToken := data.FromToken.Address
	toToken := data.ToToken.Address
	if data.FromToken.Native {
		fromToken = contracts.GetToken(data.NetworkId, contracts.WRAPPED).Address
	}
	if data.ToToken.Native {
		toToken = contracts.GetToken(data.NetworkId, contracts.WRAPPED).Address
	}

	for _, fee := range feeTiers {
		path := []common.Address{fromToken, toToken}
		s.Log.JustLog(fmt.Sprintf("Get Pool for %v fee %v", path, fee))
		addr, err := factory.GetPool(&bind.CallOpts{}, fromToken, toToken, fee)
		if err != nil {
			return addr, big.NewInt(3000), nil
		}

		if addr != (common.Address{}) {
			s.Log.JustLog(fmt.Sprintf("Pool found with fee: %d", fee.Int64()))
			return addr, fee, nil
		}
	}

	return common.Address{}, nil, fmt.Errorf("no liquidity pool found for the token pair")
}

func (s *TransactionService) calculateAmountOut(quoterAddress common.Address, quoteParam uniswapv3.IQuoterV2QuoteExactInputSingleParams) (*big.Int, error) {
	abi, err := uniswapv3.UniswapV3QuoterV2MetaData.GetAbi()
	amountOut := big.NewInt(0)
	if err != nil {
		panic(fmt.Sprintf("could not get quoter v2 abi: %v", err))
	}

	calldata, err := abi.Pack("quoteExactInputSingle",
		quoteParam)
	if err != nil {
		panic(fmt.Sprintf("failed to pack arguments for quoteExactInputSingle: %v", err))
	}
	resultBytes, err := s.Call(quoterAddress, quoteParam.AmountIn, calldata)

	if err != nil {
		return amountOut, fmt.Errorf("failed to perform static call to quoter: %w", err)
	}

	if len(resultBytes) == 0 {
		return amountOut, fmt.Errorf("quoter reverted, pool for the given pair/fee likely does not exist %v", resultBytes)
	}

	results, err := abi.Unpack("quoteExactInputSingle", resultBytes)
	if err != nil {
		return amountOut, fmt.Errorf("failed to unpack results from quoter v1: %v", err)
	}

	amountOut, ok := results[0].(*big.Int)
	if !ok {
		panic("could not cast result[0] to *big.Int for amountOut")
	}

	s.Log.JustLog(fmt.Sprintf("Amount Out Calculated : %v", amountOut))

	return amountOut, nil

}

func (s *TransactionService) prepareSwapData(swapParam *model.SwapModel) ([]byte, *string, error) {
	s.Log.LogObject("Swap Param Args", swapParam)
	if err := s.CheckAndFullFillBalance(swapParam); err != nil {
		return nil, nil, err
	}

	s.ApproveTokenSpend(swapParam.FromToken, swapParam.RouterAddress, swapParam.AmountInWei)

	_, fee, _ := s.getPoolInfo(swapParam)

	quoteParam := uniswapv3.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           swapParam.FromToken.Address,
		TokenOut:          swapParam.ToToken.Address,
		Fee:               fee,
		AmountIn:          swapParam.AmountInWei,
		SqrtPriceLimitX96: big.NewInt(0),
	}

	amountOut, err := s.calculateAmountOut(swapParam.QuoterAddress, quoteParam)
	if err != nil {
		return nil, nil, err
	}

	s.Log.Log(fmt.Sprintf("Preparing to Swap Exact Input \n%s %s to %s %s", utils.FormatUnits(swapParam.AmountInWei, swapParam.FromToken.Decimal), swapParam.FromToken.Symbol, utils.FormatUnits(amountOut, swapParam.ToToken.Decimal), swapParam.ToToken.Symbol), 2000)
	deadline := big.NewInt(time.Now().Add(10 * time.Minute).Unix())

	amountOutMinimum := new(big.Int)
	amountOutMinimum.Mul(amountOut, big.NewInt(80))
	amountOutMinimum.Div(amountOutMinimum, big.NewInt(100))

	routerAbi, err := uniswapv3.UniswapV3SwapRouterMetaData.GetAbi()
	if err != nil {
		panic("could not get default router abi")
	}

	if swapParam.CustomRouterAbi != nil {
		if swapParam.CustomRouterAbi != nil {
			parsedCustomABI, err := abi.JSON(strings.NewReader(*swapParam.CustomRouterAbi))
			if err != nil {
				panic("could not parse custom router abi")
			}
			routerAbi = &parsedCustomABI
		}
	}

	swapParams := uniswapv3.ISwapRouterExactInputSingleParams{
		TokenIn:           swapParam.FromToken.Address,
		TokenOut:          swapParam.ToToken.Address,
		Fee:               fee,
		Recipient:         s.Session.PublicKey,
		Deadline:          deadline,
		AmountIn:          swapParam.AmountInWei,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: quoteParam.SqrtPriceLimitX96,
	}
	s.Log.LogObject("Swap Param", swapParams)
	callData, err := routerAbi.Pack("exactInputSingle", swapParams)
	amountOutString := utils.FormatUnits(amountOut, swapParam.ToToken.Decimal)

	return callData, &amountOutString, err
}

func (s *TransactionService) prepareUniversalSwapData(swapParam *model.SwapModel) ([]byte, *string, error) {
	const PayableFlag byte = 0x80

	s.Log.LogObject("Swap Param Args", swapParam)

	from := swapParam.FromToken
	to := swapParam.ToToken
	wrapped := contracts.GetToken(swapParam.NetworkId, contracts.WRAPPED)

	var commands []byte
	var inputs [][]byte

	tokenInForQuote := from.Address
	if from.Native {
		tokenInForQuote = wrapped.Address
	}
	tokenOutForQuote := to.Address
	if to.Native {
		tokenOutForQuote = wrapped.Address
	}

	_, fee, err := s.getPoolInfo(swapParam)
	if err != nil {
		return nil, nil, fmt.Errorf("get pool info: %w", err)
	}
	quoteParam := uniswapv3.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           tokenInForQuote,
		TokenOut:          tokenOutForQuote,
		Fee:               fee,
		AmountIn:          swapParam.AmountInWei,
		SqrtPriceLimitX96: big.NewInt(0),
	}
	s.Log.LogObject("Queto Param", quoteParam)
	amountOut, err := s.calculateAmountOut(swapParam.QuoterAddress, quoteParam)
	if err != nil {
		return nil, nil, fmt.Errorf("quote error: %w", err)
	}
	amountOutString := utils.FormatUnits(amountOut, swapParam.ToToken.Decimal)
	s.Log.Log(fmt.Sprintf("Preparing to Swap Universal \n%s %s to %s %s",
		utils.FormatUnits(swapParam.AmountInWei, from.Decimal), from.Symbol,
		utils.FormatUnits(amountOut, to.Decimal), to.Symbol), 2000)
	amountOutMinimum := big.NewInt(0)
	if amountOut.Cmp(big.NewInt(0)) > 0 {
		amountOutMinimum.Mul(amountOut, big.NewInt(80))
		amountOutMinimum.Div(amountOutMinimum, big.NewInt(100))
		s.Log.JustLog(fmt.Sprintf("Amount Out After slippage : %s %s",
			utils.FormatUnits(amountOutMinimum, to.Decimal), to.Symbol))
	}

	deadline := big.NewInt(time.Now().Add(10 * time.Minute).Unix())

	if from.Native {
		s.Log.Log("Native token detected, preparing WRAP_ETH command...", 1000)
		commands = append(commands, byte(contracts.WRAP_ETH)|PayableFlag)
		wrapABI, err := abi.JSON(strings.NewReader(contracts.UniversalCommandABIInputs[contracts.WRAP_ETH]))
		if err != nil {
			return nil, nil, fmt.Errorf("parse WRAP_ETH ABI: %w", err)
		}
		wrapMethod := wrapABI.Methods["wrap"]
		wrapInputs, err := wrapMethod.Inputs.Pack(
			swapParam.UniversalRouterAddress,
			swapParam.AmountInWei,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("encode wrap inputs: %w", err)
		}
		inputs = append(inputs, wrapInputs)
	} else {
		if err := s.CheckAndFullFillBalance(swapParam); err != nil {
			return nil, nil, err
		}
		if err := s.ApproveTokenSpend(from, swapParam.Permit2Address, swapParam.AmountInWei); err != nil {
			return nil, nil, err
		}
		needPermit, permitInput, err := s.preparePermit2PermitIfNeeded(swapParam, swapParam.AmountInWei)
		if err != nil {
			return nil, nil, err
		}
		if needPermit {
			commands = append(commands, byte(contracts.PERMIT2_PERMIT))
			inputs = append(inputs, permitInput)
		}
	}

	tokenInForPath := from.Address
	if from.Native {
		tokenInForPath = wrapped.Address
	}
	tokenOutForPath := to.Address
	if to.Native {
		tokenOutForPath = wrapped.Address
	}
	path := utils.EncodePath(tokenInForPath, tokenOutForPath, fee)
	if l := len(path); l < 43 || (l-20)%23 != 0 {
		return nil, nil, fmt.Errorf("invalid v3 path length %d", l)
	}

	commands = append(commands, byte(contracts.V3_SWAP_EXACT_IN))
	swapABI, err := abi.JSON(strings.NewReader(contracts.UniversalCommandABIInputs[contracts.V3_SWAP_EXACT_IN]))
	if err != nil {
		return nil, nil, fmt.Errorf("parse V3_SWAP_EXACT_IN ABI: %w", err)
	}
	swapMethod := swapABI.Methods["swap"]
	recipientForSwap := s.Session.PublicKey
	if to.Native {
		recipientForSwap = swapParam.UniversalRouterAddress
	}

	payerIsSender := !from.Native
	amountInForSwap := new(big.Int).Set(swapParam.AmountInWei)
	swapData, err := swapMethod.Inputs.Pack(
		recipientForSwap,
		amountInForSwap,
		amountOutMinimum,
		path,
		payerIsSender,
	)

	if err != nil {
		return nil, nil, fmt.Errorf("encode v3 swap inputs: %w", err)
	}
	inputs = append(inputs, swapData)

	if to.Native {
		commands = append(commands, byte(contracts.UNWRAP_WETH))
		unwrapABI, err := abi.JSON(strings.NewReader(contracts.UniversalCommandABIInputs[contracts.UNWRAP_WETH]))
		if err != nil {
			return nil, nil, fmt.Errorf("parse UNWRAP_WETH ABI: %w", err)
		}
		unwrapMethod := unwrapABI.Methods["unwrap"]
		unwrapInputs, err := unwrapMethod.Inputs.Pack(
			s.Session.PublicKey,
			amountOutMinimum,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("encode unwrap inputs: %w", err)
		}
		inputs = append(inputs, unwrapInputs)
	}

	execABI, err := abi.JSON(strings.NewReader(`[
		{"inputs":[
			{"internalType":"bytes","name":"commands","type":"bytes"},
			{"internalType":"bytes[]","name":"inputs","type":"bytes[]"},
			{"internalType":"uint256","name":"deadline","type":"uint256"}
		],"name":"execute","outputs":[],"stateMutability":"payable","type":"function"}
	]`))
	if err != nil {
		return nil, nil, fmt.Errorf("parse execute ABI: %w", err)
	}

	s.Log.JustLog(fmt.Sprintf("Universal Router Commands: %x", commands))
	s.Log.LogObject("Universal Router Inputs", inputs)
	s.Log.LogObject("Universal Router Payer is Sender", payerIsSender)

	callData, err := execABI.Pack("execute", commands, inputs, deadline)
	if err != nil {
		return nil, nil, fmt.Errorf("pack execute: %w", err)
	}

	return callData, &amountOutString, nil
}

func (s *TransactionService) BuildTransaction(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	nonce, err := s.Client.PendingNonceAt(context.Background(), s.Session.PublicKey)
	if err != nil {
		return nil, err
	}
	gasPrice, err := s.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: s.Session.PublicKey, To: &to, Value: value, Data: data, GasPrice: gasPrice}
	s.Log.LogObject("TX Msg", msg)

	gasLimit, err := s.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		revertReason := s.getRevertReason(msg)
		if revertReason != "" {
			if revertReason == "InsufficientETH" {
				bal, _ := s.GetNativeBalance()
				s.Log.JustLog(fmt.Sprintf("User ETH Balance : %v", bal))
			}
			return nil, fmt.Errorf("transaction would revert: %s", revertReason)
		}
		return nil, fmt.Errorf("\nfailed to estimate gas: %w", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &to,
		Value:    value,
		Data:     data,
	})
	s.Log.LogObject("Tx Data", tx)
	return tx, nil
}

type rpcError interface {
	ErrorData() interface{}
}

func (s *TransactionService) getRevertReason(msg ethereum.CallMsg) string {
	_, err := s.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		if rpcErr, ok := err.(rpcError); ok {
			if revertData, ok := rpcErr.ErrorData().(string); ok {

				revertBytes := common.FromHex(revertData)
				unpacked, errUnpack := abi.UnpackRevert(revertBytes)
				if errUnpack == nil {
					return unpacked
				}

				if len(revertBytes) >= 4 {
					var errorID [4]byte
					copy(errorID[:], revertBytes[:4])

					routerABI, _ := uniswapv3.UniswapV3UniversalRouterMetaData.GetAbi()
					if contractErr, errParse := routerABI.ErrorByID(errorID); errParse == nil {
						return contractErr.Name
					}
				}
			}
		}
	}
	return "unknown revert reason"
}

func (s *TransactionService) ExecuteTransaction(tx *types.Transaction) (*types.Transaction, error) {
	s.Log.Log("Executing transaction...", 0)
	chainID := big.NewInt(int64(s.Network.ChainID))
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), s.Session.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}
	err = s.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}
	return signedTx, nil
}

func (s *TransactionService) GetTransactorResult(signedTx *types.Transaction) (*types.Receipt, error) {
	s.Log.Log(fmt.Sprintf("Tx sent, waiting for confirmation... \nHash: %s", signedTx.Hash().Hex()), 0)
	receipt, err := bind.WaitMined(context.Background(), s.Client, signedTx)
	if err != nil {
		s.Log.Log("Transaction failed to confirm.", 5000)
		return nil, fmt.Errorf("failed while waiting for tx to be mined: %w", err)
	}
	explorerURL := fmt.Sprintf("%stx/%s", s.Network.Explorer, receipt.TxHash.Hex())
	s.GetWalletBalance(true)
	s.Log.Log(fmt.Sprintf("Transaction confirmed! View on explorer: \n%s", explorerURL), 5000)
	return receipt, nil
}

func (s *TransactionService) Call(to common.Address, value *big.Int, data []byte) ([]byte, error) {
	msg := ethereum.CallMsg{From: s.Session.PublicKey, To: &to, Data: data}
	resultBytes, err := s.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		revertReason := s.getRevertReason(msg)
		if revertReason != "" {
			return nil, fmt.Errorf("transaction would revert: %s", revertReason)
		}
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	return resultBytes, nil
}

func (s *TransactionService) preparePermit2PermitIfNeeded(swapParam *model.SwapModel, amountIn *big.Int) (bool, []byte, error) {
	s.Log.Log("Checking Permit2 allowance...", 500)

	owner := s.Session.PublicKey
	tokenIn := swapParam.FromToken.Address
	router := swapParam.UniversalRouterAddress
	permit2 := swapParam.Permit2Address

	p2, err := uniswapv3.NewPermit2(permit2, s.Client)
	if err != nil {
		return false, nil, fmt.Errorf("new permit2: %w", err)
	}
	allow, err := p2.Allowance(&bind.CallOpts{}, owner, tokenIn, router)
	if err != nil {
		return false, nil, fmt.Errorf("permit2.allowance: %w", err)
	}

	now := big.NewInt(time.Now().Unix())
	hasEnough := allow.Amount != nil && allow.Amount.Cmp(amountIn) >= 0
	notExpired := allow.Expiration != nil && allow.Expiration.Cmp(now) > 0

	s.Log.LogObject("Permit2 allowance", map[string]string{
		"owner":   owner.Hex(),
		"token":   tokenIn.Hex(),
		"spender": router.Hex(),
		"amount": func() string {
			if allow.Amount != nil {
				return allow.Amount.String()
			}
			return "0"
		}(),
		"expiration": func() string {
			if allow.Expiration != nil {
				return allow.Expiration.String()
			}
			return "0"
		}(),
		"nonce": func() string {
			if allow.Nonce != nil {
				return allow.Nonce.String()
			}
			return "0"
		}(),
		"needAmount": amountIn.String(),
		"now":        now.String(),
	})

	if hasEnough && notExpired {
		s.Log.Log("Permit2 allowance is sufficient and not expired, skipping permit", 500)
		return false, nil, nil
	}

	s.Log.Log("Preparing Permit2 signature...", 500)

	maxUint160 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 160), big.NewInt(1))
	exp := time.Now().Add(24 * time.Hour).Unix()
	if exp < 0 {
		exp = 0
	}
	const u48Max = (1 << 48) - 1
	if uint64(exp) > u48Max {
		exp = u48Max
	}
	expiration := big.NewInt(exp)
	sigDeadline := big.NewInt(time.Now().Add(30 * time.Minute).Unix())

	s.Log.LogObject("Permit2 params", map[string]string{
		"maxUint160":  maxUint160.String(),
		"expiration":  expiration.String(),
		"sigDeadline": sigDeadline.String(),
	})

	chainID := gethmath.HexOrDecimal256(*big.NewInt(int64(s.Network.ChainID)))

	td := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"PermitDetails": {
				{Name: "token", Type: "address"},
				{Name: "amount", Type: "uint160"},
				{Name: "expiration", Type: "uint48"},
				{Name: "nonce", Type: "uint48"},
			},
			"PermitSingle": {
				{Name: "details", Type: "PermitDetails"},
				{Name: "spender", Type: "address"},
				{Name: "sigDeadline", Type: "uint256"},
			},
		},
		PrimaryType: "PermitSingle",
		Domain: apitypes.TypedDataDomain{
			Name:              "Permit2",
			ChainId:           &chainID,
			VerifyingContract: permit2.Hex(),
		},
		Message: apitypes.TypedDataMessage{
			"details": apitypes.TypedDataMessage{
				"token":      tokenIn.Hex(),
				"amount":     maxUint160.String(),
				"expiration": expiration.String(),
				"nonce": func() string {
					if allow.Nonce != nil {
						return allow.Nonce.String()
					}
					return "0"
				}(),
			},
			"spender":     router.Hex(),
			"sigDeadline": sigDeadline.String(),
		},
	}

	s.Log.LogObject("EIP712 Domain", td.Domain)
	s.Log.LogObject("EIP712 Message", td.Message)

	ds, err := td.HashStruct("EIP712Domain", td.Domain.Map())
	if err != nil {
		return false, nil, err
	}
	mh, err := td.HashStruct(td.PrimaryType, td.Message)
	if err != nil {
		return false, nil, err
	}
	digest := crypto.Keccak256([]byte{0x19, 0x01}, ds, mh)

	s.Log.JustLog(fmt.Sprintf("EIP712 domainSeparator: 0x%x", ds))
	s.Log.JustLog(fmt.Sprintf("EIP712 messageHash: 0x%x", mh))
	s.Log.JustLog(fmt.Sprintf("EIP712 digest: 0x%x", digest))

	sig, err := crypto.Sign(digest, s.Session.PrivateKey)
	if err != nil {
		return false, nil, err
	}
	if sig[64] < 27 {
		sig[64] += 27
	}
	s.Log.Log(fmt.Sprintf("Signature generated len=%d v=%d", len(sig), sig[64]), 500)

	pABI, err := abi.JSON(strings.NewReader(contracts.UniversalCommandABIInputs[contracts.PERMIT2_PERMIT]))
	if err != nil {
		return false, nil, fmt.Errorf("parse PERMIT2_PERMIT ABI: %w", err)
	}
	method := pABI.Methods["permit"]
	s.Log.Log("PERMIT2_PERMIT ABI parsed", 300)

	nonce := big.NewInt(0)
	if allow.Nonce != nil {
		nonce = allow.Nonce
	}

	perm := uniswapV3Model.PermitSingle{
		Details: uniswapV3Model.PermitDetails{
			Token:      tokenIn,
			Amount:     maxUint160,
			Expiration: expiration,
			Nonce:      nonce,
		},
		Spender:     router,
		SigDeadline: sigDeadline,
	}

	s.Log.LogObject("PermitSingle payload", perm)

	inputBytes, err := method.Inputs.Pack(perm, sig)
	if err != nil {
		return false, nil, fmt.Errorf("encode PERMIT2_PERMIT inputs: %w", err)
	}

	s.Log.Log(fmt.Sprintf("PERMIT2_PERMIT calldata encoded (%d bytes)", len(inputBytes)), 800)
	return true, inputBytes, nil
}
