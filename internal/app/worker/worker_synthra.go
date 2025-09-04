package worker

import (
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
	"github.com/widiskel/uomi-testnet-bot/internal/config"
	"github.com/widiskel/uomi-testnet-bot/internal/domain/model"
	"github.com/widiskel/uomi-testnet-bot/pkg/utils"
)

func executeSynthraOperation(ec *chain.EthersClient, cfg config.Config) error {
	amount, amountInWei, _ := utils.GenerateRandomAmount(cfg.TxAmountMin, cfg.TxAmountMax, config.GetNetwork(contracts.UOMI_TESTNET).Decimals)

	//START OF UOMI USDC
	if err := ec.SwapUniversal(model.SwapModel{
		FromToken:              contracts.GetToken(contracts.UOMI_TESTNET, contracts.NATIVE),
		ToToken:                contracts.GetToken(contracts.UOMI_TESTNET, contracts.USDC),
		Amount:                 amount,
		AmountInWei:            amountInWei,
		Dapp:                   contracts.SYTHRA,
		NetworkId:              contracts.UOMI_TESTNET,
		QuoterAddress:          contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_QUOTERV2),
		UniversalRouterAddress: contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_UNIVERSAL_ROUTER),
	}); err != nil {
		ec.Close()
		return err
	}

	if err := ec.ClearBalance(model.SwapModel{
		FromToken:              contracts.GetToken(contracts.UOMI_TESTNET, contracts.USDC),
		ToToken:                contracts.GetToken(contracts.UOMI_TESTNET, contracts.NATIVE),
		Dapp:                   contracts.SYTHRA,
		NetworkId:              contracts.UOMI_TESTNET,
		QuoterAddress:          contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_QUOTERV2),
		UniversalRouterAddress: contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_UNIVERSAL_ROUTER),
		Permit2Address:         contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_PERMIT2),
	}); err != nil {
		ec.Close()
		return err
	}
	// END OF UOMI TO USDC

	//START OF UOMI SYNT
	if err := ec.SwapUniversal(model.SwapModel{
		FromToken:              contracts.GetToken(contracts.UOMI_TESTNET, contracts.NATIVE),
		ToToken:                contracts.GetToken(contracts.UOMI_TESTNET, contracts.SYNT),
		Amount:                 amount,
		AmountInWei:            amountInWei,
		Dapp:                   contracts.SYTHRA,
		NetworkId:              contracts.UOMI_TESTNET,
		QuoterAddress:          contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_QUOTERV2),
		UniversalRouterAddress: contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_UNIVERSAL_ROUTER),
	}); err != nil {
		ec.Close()
		return err
	}

	if err := ec.ClearBalance(model.SwapModel{
		FromToken:              contracts.GetToken(contracts.UOMI_TESTNET, contracts.SYNT),
		ToToken:                contracts.GetToken(contracts.UOMI_TESTNET, contracts.NATIVE),
		Dapp:                   contracts.SYTHRA,
		NetworkId:              contracts.UOMI_TESTNET,
		QuoterAddress:          contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_QUOTERV2),
		UniversalRouterAddress: contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_UNIVERSAL_ROUTER),
		Permit2Address:         contracts.GetContractAddress(contracts.UOMI_TESTNET, contracts.SYTHRA, contracts.UNISWAPV3_PERMIT2),
	}); err != nil {
		ec.Close()
		return err
	}
	// END OF UOMI TO SYNT

	return nil
}
