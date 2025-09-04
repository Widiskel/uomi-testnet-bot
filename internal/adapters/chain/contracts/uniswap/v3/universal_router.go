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

// RouterParameters is an auto generated low-level Go binding around an user-defined struct.
type RouterParameters struct {
	Permit2              common.Address
	Weth9                common.Address
	V2Factory            common.Address
	V3Factory            common.Address
	PairInitCodeHash     [32]byte
	PoolInitCodeHash     [32]byte
	V4PoolManager        common.Address
	V3NFTPositionManager common.Address
	V4PositionManager    common.Address
}

// UniswapV3UniversalRouterMetaData contains all meta data concerning the UniswapV3UniversalRouter contract.
var UniswapV3UniversalRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v3Factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pairInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"v4PoolManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v3NFTPositionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v4PositionManager\",\"type\":\"address\"}],\"internalType\":\"structRouterParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Currency\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"DeltaNotNegative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Currency\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"DeltaNotPositive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FromAddressIsNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"}],\"name\":\"InvalidAction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBips\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandType\",\"type\":\"uint256\"}],\"name\":\"InvalidCommandType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEthSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NotAuthorizedForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPoolManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMintAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"action\",\"type\":\"uint256\"}],\"name\":\"UnsupportedAction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidAmountOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidSwap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmountOutReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"name\":\"V4TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmountInRequested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountRequested\",\"type\":\"uint256\"}],\"name\":\"V4TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3_POSITION_MANAGER\",\"outputs\":[{\"internalType\":\"contractINonfungiblePositionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"V4_POSITION_MANAGER\",\"outputs\":[{\"internalType\":\"contractIPositionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolManager\",\"outputs\":[{\"internalType\":\"contractIPoolManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unlockCallback\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapV3UniversalRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3UniversalRouterMetaData.ABI instead.
var UniswapV3UniversalRouterABI = UniswapV3UniversalRouterMetaData.ABI

// UniswapV3UniversalRouter is an auto generated Go binding around an Ethereum contract.
type UniswapV3UniversalRouter struct {
	UniswapV3UniversalRouterCaller     // Read-only binding to the contract
	UniswapV3UniversalRouterTransactor // Write-only binding to the contract
	UniswapV3UniversalRouterFilterer   // Log filterer for contract events
}

// UniswapV3UniversalRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3UniversalRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3UniversalRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3UniversalRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3UniversalRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3UniversalRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3UniversalRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3UniversalRouterSession struct {
	Contract     *UniswapV3UniversalRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapV3UniversalRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3UniversalRouterCallerSession struct {
	Contract *UniswapV3UniversalRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// UniswapV3UniversalRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3UniversalRouterTransactorSession struct {
	Contract     *UniswapV3UniversalRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// UniswapV3UniversalRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3UniversalRouterRaw struct {
	Contract *UniswapV3UniversalRouter // Generic contract binding to access the raw methods on
}

// UniswapV3UniversalRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3UniversalRouterCallerRaw struct {
	Contract *UniswapV3UniversalRouterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3UniversalRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3UniversalRouterTransactorRaw struct {
	Contract *UniswapV3UniversalRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3UniversalRouter creates a new instance of UniswapV3UniversalRouter, bound to a specific deployed contract.
func NewUniswapV3UniversalRouter(address common.Address, backend bind.ContractBackend) (*UniswapV3UniversalRouter, error) {
	contract, err := bindUniswapV3UniversalRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3UniversalRouter{UniswapV3UniversalRouterCaller: UniswapV3UniversalRouterCaller{contract: contract}, UniswapV3UniversalRouterTransactor: UniswapV3UniversalRouterTransactor{contract: contract}, UniswapV3UniversalRouterFilterer: UniswapV3UniversalRouterFilterer{contract: contract}}, nil
}

// NewUniswapV3UniversalRouterCaller creates a new read-only instance of UniswapV3UniversalRouter, bound to a specific deployed contract.
func NewUniswapV3UniversalRouterCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3UniversalRouterCaller, error) {
	contract, err := bindUniswapV3UniversalRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3UniversalRouterCaller{contract: contract}, nil
}

// NewUniswapV3UniversalRouterTransactor creates a new write-only instance of UniswapV3UniversalRouter, bound to a specific deployed contract.
func NewUniswapV3UniversalRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3UniversalRouterTransactor, error) {
	contract, err := bindUniswapV3UniversalRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3UniversalRouterTransactor{contract: contract}, nil
}

// NewUniswapV3UniversalRouterFilterer creates a new log filterer instance of UniswapV3UniversalRouter, bound to a specific deployed contract.
func NewUniswapV3UniversalRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3UniversalRouterFilterer, error) {
	contract, err := bindUniswapV3UniversalRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3UniversalRouterFilterer{contract: contract}, nil
}

// bindUniswapV3UniversalRouter binds a generic wrapper to an already deployed contract.
func bindUniswapV3UniversalRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3UniversalRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3UniversalRouter.Contract.UniswapV3UniversalRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.UniswapV3UniversalRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.UniswapV3UniversalRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3UniversalRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.contract.Transact(opts, method, params...)
}

// V3POSITIONMANAGER is a free data retrieval call binding the contract method 0x817122dc.
//
// Solidity: function V3_POSITION_MANAGER() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCaller) V3POSITIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3UniversalRouter.contract.Call(opts, &out, "V3_POSITION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V3POSITIONMANAGER is a free data retrieval call binding the contract method 0x817122dc.
//
// Solidity: function V3_POSITION_MANAGER() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) V3POSITIONMANAGER() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.V3POSITIONMANAGER(&_UniswapV3UniversalRouter.CallOpts)
}

// V3POSITIONMANAGER is a free data retrieval call binding the contract method 0x817122dc.
//
// Solidity: function V3_POSITION_MANAGER() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCallerSession) V3POSITIONMANAGER() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.V3POSITIONMANAGER(&_UniswapV3UniversalRouter.CallOpts)
}

// V4POSITIONMANAGER is a free data retrieval call binding the contract method 0xd0c9f6cb.
//
// Solidity: function V4_POSITION_MANAGER() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCaller) V4POSITIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3UniversalRouter.contract.Call(opts, &out, "V4_POSITION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V4POSITIONMANAGER is a free data retrieval call binding the contract method 0xd0c9f6cb.
//
// Solidity: function V4_POSITION_MANAGER() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) V4POSITIONMANAGER() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.V4POSITIONMANAGER(&_UniswapV3UniversalRouter.CallOpts)
}

// V4POSITIONMANAGER is a free data retrieval call binding the contract method 0xd0c9f6cb.
//
// Solidity: function V4_POSITION_MANAGER() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCallerSession) V4POSITIONMANAGER() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.V4POSITIONMANAGER(&_UniswapV3UniversalRouter.CallOpts)
}

// MsgSender is a free data retrieval call binding the contract method 0xd737d0c7.
//
// Solidity: function msgSender() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCaller) MsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3UniversalRouter.contract.Call(opts, &out, "msgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgSender is a free data retrieval call binding the contract method 0xd737d0c7.
//
// Solidity: function msgSender() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) MsgSender() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.MsgSender(&_UniswapV3UniversalRouter.CallOpts)
}

// MsgSender is a free data retrieval call binding the contract method 0xd737d0c7.
//
// Solidity: function msgSender() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCallerSession) MsgSender() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.MsgSender(&_UniswapV3UniversalRouter.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3UniversalRouter.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) PoolManager() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.PoolManager(&_UniswapV3UniversalRouter.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterCallerSession) PoolManager() (common.Address, error) {
	return _UniswapV3UniversalRouter.Contract.PoolManager(&_UniswapV3UniversalRouter.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactor) Execute(opts *bind.TransactOpts, commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.contract.Transact(opts, "execute", commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.Execute(&_UniswapV3UniversalRouter.TransactOpts, commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactorSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.Execute(&_UniswapV3UniversalRouter.TransactOpts, commands, inputs)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactor) Execute0(opts *bind.TransactOpts, commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.contract.Transact(opts, "execute0", commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.Execute0(&_UniswapV3UniversalRouter.TransactOpts, commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactorSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.Execute0(&_UniswapV3UniversalRouter.TransactOpts, commands, inputs, deadline)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.UniswapV3SwapCallback(&_UniswapV3UniversalRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.UniswapV3SwapCallback(&_UniswapV3UniversalRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactor) UnlockCallback(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.contract.Transact(opts, "unlockCallback", data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.UnlockCallback(&_UniswapV3UniversalRouter.TransactOpts, data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactorSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.UnlockCallback(&_UniswapV3UniversalRouter.TransactOpts, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterSession) Receive() (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.Receive(&_UniswapV3UniversalRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3UniversalRouter *UniswapV3UniversalRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV3UniversalRouter.Contract.Receive(&_UniswapV3UniversalRouter.TransactOpts)
}
