package contracts

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Address common.Address
	Symbol  string
	Decimal int
	Native  bool
	WETH    bool
}

var NetworkTokens = map[NetworkID]map[TokenID]Token{
	UOMI_TESTNET: {
		NATIVE: {
			Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Decimal: 18,
			Symbol:  "UOMI",
			Native:  true,
			WETH:    false,
		},
		WRAPPED: {
			Address: common.HexToAddress("0x5FCa78E132dF589c1c799F906dC867124a2567b2"),
			Decimal: 18,
			Symbol:  "WUOMI",
			Native:  false,
			WETH:    true,
		},
		USDC: {
			Address: common.HexToAddress("0xAA9C4829415BCe70c434b7349b628017C59EC2b1"),
			Decimal: 18,
			Symbol:  "USDC",
			Native:  false,
			WETH:    false,
		},
		SYNT: {
			Address: common.HexToAddress("0x2922B2Ca5EB6b02fc5E1EBE57Fc1972eBB99F7e0"),
			Decimal: 18,
			Symbol:  "SYNT",
			Native:  false,
			WETH:    false,
		},
	},
}

func GetToken(networkID NetworkID, tokenID TokenID) Token {
	network, ok := NetworkTokens[networkID]
	if !ok {
		panic(fmt.Sprintf("networkID %v not found in NetworkTokens map", networkID))
	}
	token, ok := network[tokenID]
	if !ok {
		panic(fmt.Sprintf("tokenID %v not found for networkID %v", tokenID, networkID))
	}
	return token
}

func GetTokenID(networkID NetworkID, tokenToFind Token) (TokenID, bool) {
	network, ok := NetworkTokens[networkID]
	if !ok {
		return 0, false
	}

	for id, token := range network {
		if token == tokenToFind {
			return id, true
		}
	}

	return 0, false
}
