package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
)

type FulfillBalanceFunc func() error

type SwapModel struct {
	FromToken                contracts.Token
	ToToken                  contracts.Token
	Amount                   string
	AmountInWei              *big.Int
	Dapp                     contracts.DAppID
	NetworkId                contracts.NetworkID
	QuoterAddress            common.Address
	RouterAddress            common.Address
	UniversalRouterAddress   common.Address
	Permit2Address           common.Address
	CustomRouterAbi          *string
	CustomUniversalRouterAbi *string
	FulfillBalance           FulfillBalanceFunc
}
