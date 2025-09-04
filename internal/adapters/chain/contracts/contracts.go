package contracts

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

var NetworkDAppContracts = map[NetworkID]map[DAppID]map[ContractID]common.Address{
	UOMI_TESTNET: {
		SYTHRA: {
			UNISWAPV3_UNIVERSAL_ROUTER:           common.HexToAddress("0x197EEAd5Fe3DB82c4Cd55C5752Bc87AEdE11f230"),
			UNISWAPV3_FACTORY:                    common.HexToAddress("0x7CE5b44F2d05babd29caE68557F52ab051265F01"),
			UNISWAPV3_MULTICALL:                  common.HexToAddress("0xaDD90b7787B22106e10E4530dfc9d58D4c508791"),
			UNISWAPV3_QUOTER:                     common.HexToAddress("0x0000000000000000000000000000000000000000"),
			UNISWAPV3_QUOTERV2:                   common.HexToAddress("0xCcB2B2F8395e4462d28703469F84c95293845332"),
			UNISWAPV3_V3MIGRATOR:                 common.HexToAddress("0xde4d72aB8f4E5B2b3eA80FBe7FcFFE7687e929e2"),
			UNISWAPV3_NONFUNGIBLEPOSITIONMANAGER: common.HexToAddress("0x906515Dc7c32ab887C8B8Dce6463ac3a7816Af38"),
			UNISWAPV3_TICKLENS:                   common.HexToAddress("0xD36cA9255dea7837cE1D5B816B3b8d89c3D41152"),
			UNISWAPV3_SWAPROUTER:                 common.HexToAddress("0x2046bAA610FFCF4FBfaCE6bB5c3178f51773db82"),
			UNISWAPV3_PERMIT2:                    common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3"),
		},
	},
}

func GetContractAddress(networkID NetworkID, dappID DAppID, contractID ContractID) common.Address {
	network, ok := NetworkDAppContracts[networkID]
	if !ok {
		panic(fmt.Sprintf("networkID %v not found in NetworkDAppContracts map", networkID))
	}

	dapp, ok := network[dappID]
	if !ok {
		panic(fmt.Sprintf("dappID %v not found for networkID %v", dappID, networkID))
	}

	addr, ok := dapp[contractID]
	if !ok {
		panic(fmt.Sprintf("contractID %v not found for dappID %v", contractID, dappID))
	}

	return addr
}

func GetInfoByContractAddress(addressToFind common.Address) (NetworkID, DAppID, bool) {
	for networkID, dapps := range NetworkDAppContracts {
		for dappID, contracts := range dapps {
			for _, contractAddress := range contracts {
				if contractAddress == addressToFind {
					return networkID, dappID, true
				}
			}
		}
	}
	return 0, 0, false
}
