package config

import (
	"fmt"

	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
)

type Network struct {
	Name     string
	ChainID  int
	RPCURL   string
	Explorer string
	Symbol   string
	Decimals int
}

var Networks = map[contracts.NetworkID]Network{
	contracts.UOMI_TESTNET: {
		Name:     "Uomi Testnet",
		ChainID:  4386,
		RPCURL:   "https://finney.uomi.ai",
		Explorer: "https://explorer.uomi.ai/",
		Symbol:   "UOMI",
		Decimals: 18,
	},
}

func GetNetwork(networkID contracts.NetworkID) Network {
	network, ok := Networks[networkID]
	if !ok {
		panic(fmt.Sprintf("networkID %v not found in Networks map", networkID))
	}
	return network
}

func GetNetworkID(networkToFind Network) (contracts.NetworkID, bool) {
	for id, network := range Networks {
		if network == networkToFind {
			return id, true
		}
	}
	return 0, false
}
