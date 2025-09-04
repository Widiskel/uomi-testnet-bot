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

// ITickLensPopulatedTick is an auto generated low-level Go binding around an user-defined struct.
type ITickLensPopulatedTick struct {
	Tick           *big.Int
	LiquidityNet   *big.Int
	LiquidityGross *big.Int
}

// UniswapV3TickLensMetaData contains all meta data concerning the UniswapV3TickLens contract.
var UniswapV3TickLensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"tickBitmapIndex\",\"type\":\"int16\"}],\"name\":\"getPopulatedTicksInWord\",\"outputs\":[{\"components\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"}],\"internalType\":\"structITickLens.PopulatedTick[]\",\"name\":\"populatedTicks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV3TickLensABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3TickLensMetaData.ABI instead.
var UniswapV3TickLensABI = UniswapV3TickLensMetaData.ABI

// UniswapV3TickLens is an auto generated Go binding around an Ethereum contract.
type UniswapV3TickLens struct {
	UniswapV3TickLensCaller     // Read-only binding to the contract
	UniswapV3TickLensTransactor // Write-only binding to the contract
	UniswapV3TickLensFilterer   // Log filterer for contract events
}

// UniswapV3TickLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3TickLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3TickLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3TickLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3TickLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3TickLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3TickLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3TickLensSession struct {
	Contract     *UniswapV3TickLens // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapV3TickLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3TickLensCallerSession struct {
	Contract *UniswapV3TickLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniswapV3TickLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3TickLensTransactorSession struct {
	Contract     *UniswapV3TickLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniswapV3TickLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3TickLensRaw struct {
	Contract *UniswapV3TickLens // Generic contract binding to access the raw methods on
}

// UniswapV3TickLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3TickLensCallerRaw struct {
	Contract *UniswapV3TickLensCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3TickLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3TickLensTransactorRaw struct {
	Contract *UniswapV3TickLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3TickLens creates a new instance of UniswapV3TickLens, bound to a specific deployed contract.
func NewUniswapV3TickLens(address common.Address, backend bind.ContractBackend) (*UniswapV3TickLens, error) {
	contract, err := bindUniswapV3TickLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3TickLens{UniswapV3TickLensCaller: UniswapV3TickLensCaller{contract: contract}, UniswapV3TickLensTransactor: UniswapV3TickLensTransactor{contract: contract}, UniswapV3TickLensFilterer: UniswapV3TickLensFilterer{contract: contract}}, nil
}

// NewUniswapV3TickLensCaller creates a new read-only instance of UniswapV3TickLens, bound to a specific deployed contract.
func NewUniswapV3TickLensCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3TickLensCaller, error) {
	contract, err := bindUniswapV3TickLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3TickLensCaller{contract: contract}, nil
}

// NewUniswapV3TickLensTransactor creates a new write-only instance of UniswapV3TickLens, bound to a specific deployed contract.
func NewUniswapV3TickLensTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3TickLensTransactor, error) {
	contract, err := bindUniswapV3TickLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3TickLensTransactor{contract: contract}, nil
}

// NewUniswapV3TickLensFilterer creates a new log filterer instance of UniswapV3TickLens, bound to a specific deployed contract.
func NewUniswapV3TickLensFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3TickLensFilterer, error) {
	contract, err := bindUniswapV3TickLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3TickLensFilterer{contract: contract}, nil
}

// bindUniswapV3TickLens binds a generic wrapper to an already deployed contract.
func bindUniswapV3TickLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3TickLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3TickLens *UniswapV3TickLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3TickLens.Contract.UniswapV3TickLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3TickLens *UniswapV3TickLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3TickLens.Contract.UniswapV3TickLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3TickLens *UniswapV3TickLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3TickLens.Contract.UniswapV3TickLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3TickLens *UniswapV3TickLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3TickLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3TickLens *UniswapV3TickLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3TickLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3TickLens *UniswapV3TickLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3TickLens.Contract.contract.Transact(opts, method, params...)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3TickLens *UniswapV3TickLensCaller) GetPopulatedTicksInWord(opts *bind.CallOpts, pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	var out []interface{}
	err := _UniswapV3TickLens.contract.Call(opts, &out, "getPopulatedTicksInWord", pool, tickBitmapIndex)

	if err != nil {
		return *new([]ITickLensPopulatedTick), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITickLensPopulatedTick)).(*[]ITickLensPopulatedTick)

	return out0, err

}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3TickLens *UniswapV3TickLensSession) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	return _UniswapV3TickLens.Contract.GetPopulatedTicksInWord(&_UniswapV3TickLens.CallOpts, pool, tickBitmapIndex)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3TickLens *UniswapV3TickLensCallerSession) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	return _UniswapV3TickLens.Contract.GetPopulatedTicksInWord(&_UniswapV3TickLens.CallOpts, pool, tickBitmapIndex)
}
