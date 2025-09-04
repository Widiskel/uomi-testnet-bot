package worker

import (
	"fmt"
	"strings"

	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
	"github.com/widiskel/uomi-testnet-bot/internal/config"
	"github.com/widiskel/uomi-testnet-bot/internal/domain/model"
	"github.com/widiskel/uomi-testnet-bot/internal/platform/logger"
)

type Worker struct {
	workerSession model.Session
}

func handleError(worker *Worker, log *logger.ClassLogger, err error) (shouldStop bool) {
	errMsg := err.Error()
	fatalSubstrings := []string{
		"invalid account input",
		"failed to read from seed phrase",
		"invalid private key",
	}

	for _, sub := range fatalSubstrings {
		if strings.Contains(errMsg, sub) {
			log.Log(fmt.Sprintf("FATAL: %s. Worker for accounts %d will stop.", errMsg, worker.workerSession.AccIdx+1))
			return true
		}
	}

	log.Log(fmt.Sprintf("%s, Retrying after 10 seconds", errMsg), 10000)
	return false
}

func Run(account string, index int, cfg config.Config) {
	session := model.Session{Account: account, AccIdx: index, Address: "-"}
	worker := Worker{
		workerSession: session,
	}
	log := logger.NewNamed(fmt.Sprintf("Operation - Account %d", session.AccIdx+1), &session)

	// apiClient, err := http.NewAPIClient("", &session)
	// if err != nil {
	// 	log.Log(fmt.Sprintf("FATAL: Could not Initialize API Client %v", err), 0)
	// 	return
	// }

	ec, err := chain.New(&session, cfg, config.GetNetwork(contracts.UOMI_TESTNET))
	if err != nil {
		log.Log(fmt.Sprintf("FATAL: Could not establish initial connection: %v", err), 0)
		return
	}
	defer ec.Close()

	if err := ec.ConnectWallet(); err != nil {
		log.Log(fmt.Sprintf("FATAL: Invalid wallet credentials: %v", err), 0)
		return
	}

	for {
		if err := ec.GetWalletBalance(); err != nil {
			ec.Close()
			if handleError(&worker, log, err) {
				return
			}
			continue
		}

		// amount, amountInWei, _ := utils.GenerateRandomAmount(cfg.TxAmountMin, cfg.TxAmountMax, config.GetNetwork(contracts.UOMI_TESTNET).Decimals)
		// if err := ec.TransferNative(); err != nil {
		// 	ec.Close()
		// 	if handleError(&worker, log, err) {
		// 		return
		// 	}
		// 	continue
		// }

		// if err := ec.DepositToken(contracts.GetToken(contracts.UOMI_TESTNET, contracts.WRAPPED), amount, amountInWei); err != nil {
		// 	ec.Close()
		// 	if handleError(&worker, log, err) {
		// 		return
		// 	}
		// 	continue
		// }

		// if err := ec.WithdrawToken(contracts.GetToken(contracts.UOMI_TESTNET, contracts.WRAPPED)); err != nil {
		// 	ec.Close()
		// 	if handleError(&worker, log, err) {
		// 		return
		// 	}
		// 	continue
		// }

		if err := executeSynthraOperation(ec, cfg); err != nil {
			ec.Close()
			if handleError(&worker, log, err) {
				return
			}
			continue
		}

		log.Log("Account processing complete. Sleeping...", 1000000)
	}
}
