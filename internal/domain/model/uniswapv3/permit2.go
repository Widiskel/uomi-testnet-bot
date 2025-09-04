package uniswapV3Model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PermitDetails struct {
	Token      common.Address `abi:"token"`
	Amount     *big.Int       `abi:"amount"`
	Expiration *big.Int       `abi:"expiration"`
	Nonce      *big.Int       `abi:"nonce"`
}

type PermitSingle struct {
	Details     PermitDetails  `abi:"details"`
	Spender     common.Address `abi:"spender"`
	SigDeadline *big.Int       `abi:"sigDeadline"`
}
