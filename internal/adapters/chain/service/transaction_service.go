package service

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
	"github.com/widiskel/uomi-testnet-bot/internal/config"
	"github.com/widiskel/uomi-testnet-bot/internal/domain/model"
	"github.com/widiskel/uomi-testnet-bot/internal/platform/logger"
	"github.com/widiskel/uomi-testnet-bot/pkg/utils"
)

type TransactionService struct {
	Client  *ethclient.Client
	Session *model.Session
	Network *config.Network
	Log     *logger.ClassLogger
}

func NewTransactionService(client *ethclient.Client, session *model.Session, network *config.Network) *TransactionService {
	ts := &TransactionService{
		Client:  client,
		Session: session,
		Network: network,
	}
	ts.Log = logger.NewLogger(ts, session)
	return ts
}

func (e *TransactionService) GetWalletBalance(update ...bool) error {
	isUpdate := false
	if len(update) > 0 {
		isUpdate = update[0]
	}

	if !isUpdate {
		e.Log.Log("Fetching Wallet balances...")
	}
	e.Session.WalletBalance.Balances = nil

	nativeBal, err := e.GetNativeBalance()
	if err != nil {
		e.Log.Log(fmt.Sprintf("Could not fetch native balance: %v", err))
	} else {
		tokenBalance := model.TokenBalance{
			Token:      contracts.Token{Symbol: e.Network.Symbol, Native: true, Decimal: e.Network.Decimals},
			Balance:    *nativeBal,
			BalanceStr: utils.FormatUnits(nativeBal, e.Network.Decimals),
		}
		e.Session.WalletBalance.Balances = append(e.Session.WalletBalance.Balances, tokenBalance)
	}

	networkID, ok := config.GetNetworkID(*e.Network)
	if !ok {
		panic("Network ID not found")
	}

	wrappedToken := contracts.GetToken(networkID, contracts.WRAPPED)
	wrappedTokenBalance, err := e.GetTokenBalance(wrappedToken)
	if err != nil {
		e.Log.Log(fmt.Sprintf("Could not fetch Wrapped Token %s balance: %v", wrappedToken.Symbol, err), 3000)
	} else {
		tokenBalance := model.TokenBalance{
			Token:      wrappedToken,
			Balance:    *wrappedTokenBalance,
			BalanceStr: utils.FormatUnits(wrappedTokenBalance, wrappedToken.Decimal),
		}
		e.Session.WalletBalance.Balances = append(e.Session.WalletBalance.Balances, tokenBalance)
	}

	usdcToken := contracts.GetToken(networkID, contracts.USDC)
	usdcBal, err := e.GetTokenBalance(usdcToken)
	if err != nil {
		e.Log.Log(fmt.Sprintf("Could not fetch USDC balance: %v", err))
	} else {
		tokenBalance := model.TokenBalance{
			Token:      usdcToken,
			Balance:    *usdcBal,
			BalanceStr: utils.FormatUnits(usdcBal, usdcToken.Decimal),
		}
		e.Session.WalletBalance.Balances = append(e.Session.WalletBalance.Balances, tokenBalance)
	}

	syntToken := contracts.GetToken(networkID, contracts.SYNT)
	syntBal, err := e.GetTokenBalance(syntToken)
	if err != nil {
		e.Log.Log(fmt.Sprintf("Could not fetch SYNT balance: %v", err))
	} else {
		tokenBalance := model.TokenBalance{
			Token:      syntToken,
			Balance:    *syntBal,
			BalanceStr: utils.FormatUnits(syntBal, syntToken.Decimal),
		}
		e.Session.WalletBalance.Balances = append(e.Session.WalletBalance.Balances, tokenBalance)
	}

	return nil
}

func (s *TransactionService) TransferNative(to common.Address, amountInWei *big.Int) (*types.Receipt, error) {
	tx, err := s.BuildTransaction(to, amountInWei, nil)
	if err != nil {
		return nil, err
	}
	signedTx, err := s.ExecuteTransaction(tx)
	if err != nil {
		return nil, err
	}
	receipt, err := s.GetTransactorResult(signedTx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *TransactionService) Deposit(wrappedToken contracts.Token, amount *big.Int) (*types.Receipt, error) {
	data, err := s.prepareDepositData()
	if err != nil {
		return nil, err
	}
	tx, err := s.BuildTransaction(wrappedToken.Address, amount, data)
	if err != nil {
		return nil, err
	}
	signedTx, err := s.ExecuteTransaction(tx)
	if err != nil {
		return nil, err
	}
	receipt, err := s.GetTransactorResult(signedTx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *TransactionService) Withdraw(wrappedToken contracts.Token, amount *big.Int) (*types.Receipt, error) {
	data, err := s.prepareWithdrawData(amount)
	if err != nil {
		return nil, err
	}
	tx, err := s.BuildTransaction(wrappedToken.Address, big.NewInt(0), data)
	if err != nil {
		return nil, err
	}
	signedTx, err := s.ExecuteTransaction(tx)
	if err != nil {
		return nil, err
	}
	receipt, err := s.GetTransactorResult(signedTx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *TransactionService) SwapExactInputSingle(swapParam *model.SwapModel) (receipt *types.Receipt, err error) {
	data, amountOut, err := s.prepareSwapData(swapParam)
	if err != nil {
		return nil, err
	}

	tx, err := s.BuildTransaction(swapParam.RouterAddress, big.NewInt(0), data)
	if err != nil {
		return nil, err
	}

	signedTx, err := s.ExecuteTransaction(tx)
	if err != nil {
		return nil, err
	}
	receipt, err = s.GetTransactorResult(signedTx)
	if err != nil {
		return nil, err
	}
	s.Log.Log(fmt.Sprintf("Successfully Swap Exact Input \n%s %s to %s %s", utils.FormatUnits(swapParam.AmountInWei, swapParam.FromToken.Decimal), swapParam.FromToken.Symbol, *amountOut, swapParam.ToToken.Symbol), 3000)
	return receipt, err
}

func (s *TransactionService) SwapUniversal(swapParam *model.SwapModel) (receipt *types.Receipt, err error) {
	data, amountOut, err := s.prepareUniversalSwapData(swapParam)
	if err != nil {
		return nil, err
	}

	txVal := big.NewInt(0)
	if swapParam.FromToken.Native {
		txVal = swapParam.AmountInWei
	}

	tx, err := s.BuildTransaction(swapParam.UniversalRouterAddress, txVal, data)
	if err != nil {
		return nil, err
	}

	signedTx, err := s.ExecuteTransaction(tx)
	if err != nil {
		return nil, err
	}
	receipt, err = s.GetTransactorResult(signedTx)
	if err != nil {
		return nil, err
	}
	s.Log.Log(fmt.Sprintf("Successfully Swap Universal \n%s %s to %s %s", utils.FormatUnits(swapParam.AmountInWei, swapParam.FromToken.Decimal), swapParam.FromToken.Symbol, *amountOut, swapParam.ToToken.Symbol), 3000)
	return receipt, err
}
