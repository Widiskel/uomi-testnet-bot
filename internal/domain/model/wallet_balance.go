package model

import (
	"math/big"

	"github.com/widiskel/uomi-testnet-bot/internal/adapters/chain/contracts"
)

type TokenBalance struct {
	Token      contracts.Token
	Balance    big.Int
	BalanceStr string
}

type WalletBalance struct {
	Balances []TokenBalance
}
