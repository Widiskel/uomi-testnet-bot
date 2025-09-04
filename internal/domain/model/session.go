package model

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type Session struct {
	Account       string
	AccIdx        int
	Address       string
	PublicKey     common.Address
	PrivateKey    *ecdsa.PrivateKey
	WalletBalance WalletBalance
}
