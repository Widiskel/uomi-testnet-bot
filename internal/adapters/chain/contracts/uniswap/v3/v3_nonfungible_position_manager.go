// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// INonfungiblePositionManagerCollectParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerCollectParams struct {
	TokenId    *big.Int
	Recipient  common.Address
	Amount0Max *big.Int
	Amount1Max *big.Int
}

// INonfungiblePositionManagerDecreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerDecreaseLiquidityParams struct {
	TokenId    *big.Int
	Liquidity  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Deadline   *big.Int
}

// INonfungiblePositionManagerIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerIncreaseLiquidityParams struct {
	TokenId        *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Deadline       *big.Int
}

// INonfungiblePositionManagerMintParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerMintParams struct {
	Token0         common.Address
	Token1         common.Address
	Fee            *big.Int
	TickLower      *big.Int
	TickUpper      *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Recipient      common.Address
	Deadline       *big.Int
}

// UniswapV3NonFungiblePositionManagerMetaData contains all meta data concerning the UniswapV3NonFungiblePositionManager contract.
var UniswapV3NonFungiblePositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenDescriptor_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"DecreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"IncreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Max\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Max\",\"type\":\"uint128\"}],\"internalType\":\"structINonfungiblePositionManager.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.DecreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"decreaseLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapV3NonFungiblePositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3NonFungiblePositionManagerMetaData.ABI instead.
var UniswapV3NonFungiblePositionManagerABI = UniswapV3NonFungiblePositionManagerMetaData.ABI

// UniswapV3NonFungiblePositionManager is an auto generated Go binding around an Ethereum contract.
type UniswapV3NonFungiblePositionManager struct {
	UniswapV3NonFungiblePositionManagerCaller     // Read-only binding to the contract
	UniswapV3NonFungiblePositionManagerTransactor // Write-only binding to the contract
	UniswapV3NonFungiblePositionManagerFilterer   // Log filterer for contract events
}

// UniswapV3NonFungiblePositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3NonFungiblePositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3NonFungiblePositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3NonFungiblePositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3NonFungiblePositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3NonFungiblePositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3NonFungiblePositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3NonFungiblePositionManagerSession struct {
	Contract     *UniswapV3NonFungiblePositionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// UniswapV3NonFungiblePositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3NonFungiblePositionManagerCallerSession struct {
	Contract *UniswapV3NonFungiblePositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// UniswapV3NonFungiblePositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3NonFungiblePositionManagerTransactorSession struct {
	Contract     *UniswapV3NonFungiblePositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// UniswapV3NonFungiblePositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3NonFungiblePositionManagerRaw struct {
	Contract *UniswapV3NonFungiblePositionManager // Generic contract binding to access the raw methods on
}

// UniswapV3NonFungiblePositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3NonFungiblePositionManagerCallerRaw struct {
	Contract *UniswapV3NonFungiblePositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3NonFungiblePositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3NonFungiblePositionManagerTransactorRaw struct {
	Contract *UniswapV3NonFungiblePositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3NonFungiblePositionManager creates a new instance of UniswapV3NonFungiblePositionManager, bound to a specific deployed contract.
func NewUniswapV3NonFungiblePositionManager(address common.Address, backend bind.ContractBackend) (*UniswapV3NonFungiblePositionManager, error) {
	contract, err := bindUniswapV3NonFungiblePositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManager{UniswapV3NonFungiblePositionManagerCaller: UniswapV3NonFungiblePositionManagerCaller{contract: contract}, UniswapV3NonFungiblePositionManagerTransactor: UniswapV3NonFungiblePositionManagerTransactor{contract: contract}, UniswapV3NonFungiblePositionManagerFilterer: UniswapV3NonFungiblePositionManagerFilterer{contract: contract}}, nil
}

// NewUniswapV3NonFungiblePositionManagerCaller creates a new read-only instance of UniswapV3NonFungiblePositionManager, bound to a specific deployed contract.
func NewUniswapV3NonFungiblePositionManagerCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3NonFungiblePositionManagerCaller, error) {
	contract, err := bindUniswapV3NonFungiblePositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerCaller{contract: contract}, nil
}

// NewUniswapV3NonFungiblePositionManagerTransactor creates a new write-only instance of UniswapV3NonFungiblePositionManager, bound to a specific deployed contract.
func NewUniswapV3NonFungiblePositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3NonFungiblePositionManagerTransactor, error) {
	contract, err := bindUniswapV3NonFungiblePositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerTransactor{contract: contract}, nil
}

// NewUniswapV3NonFungiblePositionManagerFilterer creates a new log filterer instance of UniswapV3NonFungiblePositionManager, bound to a specific deployed contract.
func NewUniswapV3NonFungiblePositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3NonFungiblePositionManagerFilterer, error) {
	contract, err := bindUniswapV3NonFungiblePositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerFilterer{contract: contract}, nil
}

// bindUniswapV3NonFungiblePositionManager binds a generic wrapper to an already deployed contract.
func bindUniswapV3NonFungiblePositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3NonFungiblePositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3NonFungiblePositionManager.Contract.UniswapV3NonFungiblePositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.UniswapV3NonFungiblePositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.UniswapV3NonFungiblePositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3NonFungiblePositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.DOMAINSEPARATOR(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.DOMAINSEPARATOR(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.PERMITTYPEHASH(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.PERMITTYPEHASH(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) WETH9() (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.WETH9(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) WETH9() (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.WETH9(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.BalanceOf(&_UniswapV3NonFungiblePositionManager.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.BalanceOf(&_UniswapV3NonFungiblePositionManager.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() pure returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() pure returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) BaseURI() (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.BaseURI(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() pure returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) BaseURI() (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.BaseURI(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Factory() (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Factory(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) Factory() (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Factory(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.GetApproved(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.GetApproved(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.IsApprovedForAll(&_UniswapV3NonFungiblePositionManager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.IsApprovedForAll(&_UniswapV3NonFungiblePositionManager.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Name() (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Name(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) Name() (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Name(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.OwnerOf(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.OwnerOf(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) Positions(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "positions", tokenId)

	outstruct := new(struct {
		Nonce                    *big.Int
		Operator                 common.Address
		Token0                   common.Address
		Token1                   common.Address
		Fee                      *big.Int
		TickLower                *big.Int
		TickUpper                *big.Int
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TickLower = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TickUpper = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Positions(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Positions(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SupportsInterface(&_UniswapV3NonFungiblePositionManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SupportsInterface(&_UniswapV3NonFungiblePositionManager.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Symbol() (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Symbol(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) Symbol() (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Symbol(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TokenByIndex(&_UniswapV3NonFungiblePositionManager.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TokenByIndex(&_UniswapV3NonFungiblePositionManager.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TokenOfOwnerByIndex(&_UniswapV3NonFungiblePositionManager.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TokenOfOwnerByIndex(&_UniswapV3NonFungiblePositionManager.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TokenURI(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TokenURI(&_UniswapV3NonFungiblePositionManager.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3NonFungiblePositionManager.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) TotalSupply() (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TotalSupply(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerCallerSession) TotalSupply() (*big.Int, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TotalSupply(&_UniswapV3NonFungiblePositionManager.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Approve(&_UniswapV3NonFungiblePositionManager.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Approve(&_UniswapV3NonFungiblePositionManager.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Burn(&_UniswapV3NonFungiblePositionManager.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Burn(&_UniswapV3NonFungiblePositionManager.TransactOpts, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) Collect(opts *bind.TransactOpts, params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "collect", params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Collect(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Collect(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.CreateAndInitializePoolIfNecessary(&_UniswapV3NonFungiblePositionManager.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.CreateAndInitializePoolIfNecessary(&_UniswapV3NonFungiblePositionManager.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) DecreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "decreaseLiquidity", params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.DecreaseLiquidity(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.DecreaseLiquidity(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) IncreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.IncreaseLiquidity(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.IncreaseLiquidity(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) Mint(opts *bind.TransactOpts, params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Mint(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Mint(&_UniswapV3NonFungiblePositionManager.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Multicall(&_UniswapV3NonFungiblePositionManager.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Multicall(&_UniswapV3NonFungiblePositionManager.TransactOpts, data)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "permit", spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Permit(&_UniswapV3NonFungiblePositionManager.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Permit(&_UniswapV3NonFungiblePositionManager.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) RefundETH() (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.RefundETH(&_UniswapV3NonFungiblePositionManager.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) RefundETH() (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.RefundETH(&_UniswapV3NonFungiblePositionManager.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SafeTransferFrom(&_UniswapV3NonFungiblePositionManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SafeTransferFrom(&_UniswapV3NonFungiblePositionManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SafeTransferFrom0(&_UniswapV3NonFungiblePositionManager.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SafeTransferFrom0(&_UniswapV3NonFungiblePositionManager.TransactOpts, from, to, tokenId, _data)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermit(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermit(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermitAllowed(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermitAllowed(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermitAllowedIfNecessary(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermitAllowedIfNecessary(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermitIfNecessary(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SelfPermitIfNecessary(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, value, deadline, v, r, s)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SetApprovalForAll(&_UniswapV3NonFungiblePositionManager.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SetApprovalForAll(&_UniswapV3NonFungiblePositionManager.TransactOpts, operator, approved)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SweepToken(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.SweepToken(&_UniswapV3NonFungiblePositionManager.TransactOpts, token, amountMinimum, recipient)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TransferFrom(&_UniswapV3NonFungiblePositionManager.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.TransferFrom(&_UniswapV3NonFungiblePositionManager.TransactOpts, from, to, tokenId)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.UniswapV3MintCallback(&_UniswapV3NonFungiblePositionManager.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.UniswapV3MintCallback(&_UniswapV3NonFungiblePositionManager.TransactOpts, amount0Owed, amount1Owed, data)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.UnwrapWETH9(&_UniswapV3NonFungiblePositionManager.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.UnwrapWETH9(&_UniswapV3NonFungiblePositionManager.TransactOpts, amountMinimum, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerSession) Receive() (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Receive(&_UniswapV3NonFungiblePositionManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV3NonFungiblePositionManager.Contract.Receive(&_UniswapV3NonFungiblePositionManager.TransactOpts)
}

// UniswapV3NonFungiblePositionManagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerApprovalIterator struct {
	Event *UniswapV3NonFungiblePositionManagerApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapV3NonFungiblePositionManagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3NonFungiblePositionManagerApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapV3NonFungiblePositionManagerApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapV3NonFungiblePositionManagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3NonFungiblePositionManagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3NonFungiblePositionManagerApproval represents a Approval event raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*UniswapV3NonFungiblePositionManagerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerApprovalIterator{contract: _UniswapV3NonFungiblePositionManager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UniswapV3NonFungiblePositionManagerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3NonFungiblePositionManagerApproval)
				if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) ParseApproval(log types.Log) (*UniswapV3NonFungiblePositionManagerApproval, error) {
	event := new(UniswapV3NonFungiblePositionManagerApproval)
	if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3NonFungiblePositionManagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerApprovalForAllIterator struct {
	Event *UniswapV3NonFungiblePositionManagerApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapV3NonFungiblePositionManagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3NonFungiblePositionManagerApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapV3NonFungiblePositionManagerApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapV3NonFungiblePositionManagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3NonFungiblePositionManagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3NonFungiblePositionManagerApprovalForAll represents a ApprovalForAll event raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*UniswapV3NonFungiblePositionManagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerApprovalForAllIterator{contract: _UniswapV3NonFungiblePositionManager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *UniswapV3NonFungiblePositionManagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3NonFungiblePositionManagerApprovalForAll)
				if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) ParseApprovalForAll(log types.Log) (*UniswapV3NonFungiblePositionManagerApprovalForAll, error) {
	event := new(UniswapV3NonFungiblePositionManagerApprovalForAll)
	if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3NonFungiblePositionManagerCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerCollectIterator struct {
	Event *UniswapV3NonFungiblePositionManagerCollect // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapV3NonFungiblePositionManagerCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3NonFungiblePositionManagerCollect)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapV3NonFungiblePositionManagerCollect)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapV3NonFungiblePositionManagerCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3NonFungiblePositionManagerCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3NonFungiblePositionManagerCollect represents a Collect event raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerCollect struct {
	TokenId   *big.Int
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) FilterCollect(opts *bind.FilterOpts, tokenId []*big.Int) (*UniswapV3NonFungiblePositionManagerCollectIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.FilterLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerCollectIterator{contract: _UniswapV3NonFungiblePositionManager.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *UniswapV3NonFungiblePositionManagerCollect, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.WatchLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3NonFungiblePositionManagerCollect)
				if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollect is a log parse operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) ParseCollect(log types.Log) (*UniswapV3NonFungiblePositionManagerCollect, error) {
	event := new(UniswapV3NonFungiblePositionManagerCollect)
	if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3NonFungiblePositionManagerDecreaseLiquidityIterator is returned from FilterDecreaseLiquidity and is used to iterate over the raw logs and unpacked data for DecreaseLiquidity events raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerDecreaseLiquidityIterator struct {
	Event *UniswapV3NonFungiblePositionManagerDecreaseLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapV3NonFungiblePositionManagerDecreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3NonFungiblePositionManagerDecreaseLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapV3NonFungiblePositionManagerDecreaseLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapV3NonFungiblePositionManagerDecreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3NonFungiblePositionManagerDecreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3NonFungiblePositionManagerDecreaseLiquidity represents a DecreaseLiquidity event raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerDecreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecreaseLiquidity is a free log retrieval operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) FilterDecreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*UniswapV3NonFungiblePositionManagerDecreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.FilterLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerDecreaseLiquidityIterator{contract: _UniswapV3NonFungiblePositionManager.contract, event: "DecreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchDecreaseLiquidity is a free log subscription operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) WatchDecreaseLiquidity(opts *bind.WatchOpts, sink chan<- *UniswapV3NonFungiblePositionManagerDecreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.WatchLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3NonFungiblePositionManagerDecreaseLiquidity)
				if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseLiquidity is a log parse operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) ParseDecreaseLiquidity(log types.Log) (*UniswapV3NonFungiblePositionManagerDecreaseLiquidity, error) {
	event := new(UniswapV3NonFungiblePositionManagerDecreaseLiquidity)
	if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3NonFungiblePositionManagerIncreaseLiquidityIterator is returned from FilterIncreaseLiquidity and is used to iterate over the raw logs and unpacked data for IncreaseLiquidity events raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerIncreaseLiquidityIterator struct {
	Event *UniswapV3NonFungiblePositionManagerIncreaseLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapV3NonFungiblePositionManagerIncreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3NonFungiblePositionManagerIncreaseLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapV3NonFungiblePositionManagerIncreaseLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapV3NonFungiblePositionManagerIncreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3NonFungiblePositionManagerIncreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3NonFungiblePositionManagerIncreaseLiquidity represents a IncreaseLiquidity event raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerIncreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreaseLiquidity is a free log retrieval operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) FilterIncreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*UniswapV3NonFungiblePositionManagerIncreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.FilterLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerIncreaseLiquidityIterator{contract: _UniswapV3NonFungiblePositionManager.contract, event: "IncreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchIncreaseLiquidity is a free log subscription operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) WatchIncreaseLiquidity(opts *bind.WatchOpts, sink chan<- *UniswapV3NonFungiblePositionManagerIncreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.WatchLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3NonFungiblePositionManagerIncreaseLiquidity)
				if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseLiquidity is a log parse operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) ParseIncreaseLiquidity(log types.Log) (*UniswapV3NonFungiblePositionManagerIncreaseLiquidity, error) {
	event := new(UniswapV3NonFungiblePositionManagerIncreaseLiquidity)
	if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3NonFungiblePositionManagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerTransferIterator struct {
	Event *UniswapV3NonFungiblePositionManagerTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapV3NonFungiblePositionManagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3NonFungiblePositionManagerTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapV3NonFungiblePositionManagerTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapV3NonFungiblePositionManagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3NonFungiblePositionManagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3NonFungiblePositionManagerTransfer represents a Transfer event raised by the UniswapV3NonFungiblePositionManager contract.
type UniswapV3NonFungiblePositionManagerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*UniswapV3NonFungiblePositionManagerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3NonFungiblePositionManagerTransferIterator{contract: _UniswapV3NonFungiblePositionManager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UniswapV3NonFungiblePositionManagerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UniswapV3NonFungiblePositionManager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3NonFungiblePositionManagerTransfer)
				if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_UniswapV3NonFungiblePositionManager *UniswapV3NonFungiblePositionManagerFilterer) ParseTransfer(log types.Log) (*UniswapV3NonFungiblePositionManagerTransfer, error) {
	event := new(UniswapV3NonFungiblePositionManagerTransfer)
	if err := _UniswapV3NonFungiblePositionManager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
