// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// RandomNumberOracleMetaData contains all meta data concerning the RandomNumberOracle contract.
var RandomNumberOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"ConsumerAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"ConsumerDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"ProviderAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"ProviderDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"}],\"name\":\"RandomNumberRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"RequestCanceled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"}],\"name\":\"allowConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"allowProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"}],\"name\":\"denyConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"denyProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_rnd\",\"type\":\"uint256\"}],\"name\":\"fulfillRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getConsumer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_seed\",\"type\":\"bytes32\"}],\"name\":\"requestRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RandomNumberOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomNumberOracleMetaData.ABI instead.
var RandomNumberOracleABI = RandomNumberOracleMetaData.ABI

// RandomNumberOracle is an auto generated Go binding around an Ethereum contract.
type RandomNumberOracle struct {
	RandomNumberOracleCaller     // Read-only binding to the contract
	RandomNumberOracleTransactor // Write-only binding to the contract
	RandomNumberOracleFilterer   // Log filterer for contract events
}

// RandomNumberOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomNumberOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomNumberOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomNumberOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomNumberOracleSession struct {
	Contract     *RandomNumberOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RandomNumberOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomNumberOracleCallerSession struct {
	Contract *RandomNumberOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RandomNumberOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomNumberOracleTransactorSession struct {
	Contract     *RandomNumberOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RandomNumberOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomNumberOracleRaw struct {
	Contract *RandomNumberOracle // Generic contract binding to access the raw methods on
}

// RandomNumberOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomNumberOracleCallerRaw struct {
	Contract *RandomNumberOracleCaller // Generic read-only contract binding to access the raw methods on
}

// RandomNumberOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomNumberOracleTransactorRaw struct {
	Contract *RandomNumberOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomNumberOracle creates a new instance of RandomNumberOracle, bound to a specific deployed contract.
func NewRandomNumberOracle(address common.Address, backend bind.ContractBackend) (*RandomNumberOracle, error) {
	contract, err := bindRandomNumberOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracle{RandomNumberOracleCaller: RandomNumberOracleCaller{contract: contract}, RandomNumberOracleTransactor: RandomNumberOracleTransactor{contract: contract}, RandomNumberOracleFilterer: RandomNumberOracleFilterer{contract: contract}}, nil
}

// NewRandomNumberOracleCaller creates a new read-only instance of RandomNumberOracle, bound to a specific deployed contract.
func NewRandomNumberOracleCaller(address common.Address, caller bind.ContractCaller) (*RandomNumberOracleCaller, error) {
	contract, err := bindRandomNumberOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleCaller{contract: contract}, nil
}

// NewRandomNumberOracleTransactor creates a new write-only instance of RandomNumberOracle, bound to a specific deployed contract.
func NewRandomNumberOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomNumberOracleTransactor, error) {
	contract, err := bindRandomNumberOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleTransactor{contract: contract}, nil
}

// NewRandomNumberOracleFilterer creates a new log filterer instance of RandomNumberOracle, bound to a specific deployed contract.
func NewRandomNumberOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomNumberOracleFilterer, error) {
	contract, err := bindRandomNumberOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleFilterer{contract: contract}, nil
}

// bindRandomNumberOracle binds a generic wrapper to an already deployed contract.
func bindRandomNumberOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomNumberOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomNumberOracle *RandomNumberOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomNumberOracle.Contract.RandomNumberOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomNumberOracle *RandomNumberOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.RandomNumberOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomNumberOracle *RandomNumberOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.RandomNumberOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomNumberOracle *RandomNumberOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomNumberOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomNumberOracle *RandomNumberOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomNumberOracle *RandomNumberOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.contract.Transact(opts, method, params...)
}

// GetConsumer is a free data retrieval call binding the contract method 0x993084e8.
//
// Solidity: function getConsumer(address ) view returns(bool)
func (_RandomNumberOracle *RandomNumberOracleCaller) GetConsumer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RandomNumberOracle.contract.Call(opts, &out, "getConsumer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetConsumer is a free data retrieval call binding the contract method 0x993084e8.
//
// Solidity: function getConsumer(address ) view returns(bool)
func (_RandomNumberOracle *RandomNumberOracleSession) GetConsumer(arg0 common.Address) (bool, error) {
	return _RandomNumberOracle.Contract.GetConsumer(&_RandomNumberOracle.CallOpts, arg0)
}

// GetConsumer is a free data retrieval call binding the contract method 0x993084e8.
//
// Solidity: function getConsumer(address ) view returns(bool)
func (_RandomNumberOracle *RandomNumberOracleCallerSession) GetConsumer(arg0 common.Address) (bool, error) {
	return _RandomNumberOracle.Contract.GetConsumer(&_RandomNumberOracle.CallOpts, arg0)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address ) view returns(uint256)
func (_RandomNumberOracle *RandomNumberOracleCaller) GetNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RandomNumberOracle.contract.Call(opts, &out, "getNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address ) view returns(uint256)
func (_RandomNumberOracle *RandomNumberOracleSession) GetNonce(arg0 common.Address) (*big.Int, error) {
	return _RandomNumberOracle.Contract.GetNonce(&_RandomNumberOracle.CallOpts, arg0)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address ) view returns(uint256)
func (_RandomNumberOracle *RandomNumberOracleCallerSession) GetNonce(arg0 common.Address) (*big.Int, error) {
	return _RandomNumberOracle.Contract.GetNonce(&_RandomNumberOracle.CallOpts, arg0)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_RandomNumberOracle *RandomNumberOracleCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomNumberOracle.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_RandomNumberOracle *RandomNumberOracleSession) GetOwner() (common.Address, error) {
	return _RandomNumberOracle.Contract.GetOwner(&_RandomNumberOracle.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_RandomNumberOracle *RandomNumberOracleCallerSession) GetOwner() (common.Address, error) {
	return _RandomNumberOracle.Contract.GetOwner(&_RandomNumberOracle.CallOpts)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address ) view returns(bool)
func (_RandomNumberOracle *RandomNumberOracleCaller) GetProvider(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RandomNumberOracle.contract.Call(opts, &out, "getProvider", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address ) view returns(bool)
func (_RandomNumberOracle *RandomNumberOracleSession) GetProvider(arg0 common.Address) (bool, error) {
	return _RandomNumberOracle.Contract.GetProvider(&_RandomNumberOracle.CallOpts, arg0)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address ) view returns(bool)
func (_RandomNumberOracle *RandomNumberOracleCallerSession) GetProvider(arg0 common.Address) (bool, error) {
	return _RandomNumberOracle.Contract.GetProvider(&_RandomNumberOracle.CallOpts, arg0)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 ) view returns(bytes32 seed, address consumer)
func (_RandomNumberOracle *RandomNumberOracleCaller) GetRequest(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Seed     [32]byte
	Consumer common.Address
}, error) {
	var out []interface{}
	err := _RandomNumberOracle.contract.Call(opts, &out, "getRequest", arg0)

	outstruct := new(struct {
		Seed     [32]byte
		Consumer common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seed = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Consumer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 ) view returns(bytes32 seed, address consumer)
func (_RandomNumberOracle *RandomNumberOracleSession) GetRequest(arg0 [32]byte) (struct {
	Seed     [32]byte
	Consumer common.Address
}, error) {
	return _RandomNumberOracle.Contract.GetRequest(&_RandomNumberOracle.CallOpts, arg0)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 ) view returns(bytes32 seed, address consumer)
func (_RandomNumberOracle *RandomNumberOracleCallerSession) GetRequest(arg0 [32]byte) (struct {
	Seed     [32]byte
	Consumer common.Address
}, error) {
	return _RandomNumberOracle.Contract.GetRequest(&_RandomNumberOracle.CallOpts, arg0)
}

// AllowConsumer is a paid mutator transaction binding the contract method 0x538f8f5e.
//
// Solidity: function allowConsumer(address _consumer) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) AllowConsumer(opts *bind.TransactOpts, _consumer common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "allowConsumer", _consumer)
}

// AllowConsumer is a paid mutator transaction binding the contract method 0x538f8f5e.
//
// Solidity: function allowConsumer(address _consumer) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) AllowConsumer(_consumer common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.AllowConsumer(&_RandomNumberOracle.TransactOpts, _consumer)
}

// AllowConsumer is a paid mutator transaction binding the contract method 0x538f8f5e.
//
// Solidity: function allowConsumer(address _consumer) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) AllowConsumer(_consumer common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.AllowConsumer(&_RandomNumberOracle.TransactOpts, _consumer)
}

// AllowProvider is a paid mutator transaction binding the contract method 0x34750126.
//
// Solidity: function allowProvider(address _provider) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) AllowProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "allowProvider", _provider)
}

// AllowProvider is a paid mutator transaction binding the contract method 0x34750126.
//
// Solidity: function allowProvider(address _provider) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) AllowProvider(_provider common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.AllowProvider(&_RandomNumberOracle.TransactOpts, _provider)
}

// AllowProvider is a paid mutator transaction binding the contract method 0x34750126.
//
// Solidity: function allowProvider(address _provider) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) AllowProvider(_provider common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.AllowProvider(&_RandomNumberOracle.TransactOpts, _provider)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 _requestID) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) CancelRequest(opts *bind.TransactOpts, _requestID [32]byte) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "cancelRequest", _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 _requestID) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) CancelRequest(_requestID [32]byte) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.CancelRequest(&_RandomNumberOracle.TransactOpts, _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 _requestID) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) CancelRequest(_requestID [32]byte) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.CancelRequest(&_RandomNumberOracle.TransactOpts, _requestID)
}

// DenyConsumer is a paid mutator transaction binding the contract method 0x6693a679.
//
// Solidity: function denyConsumer(address _consumer) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) DenyConsumer(opts *bind.TransactOpts, _consumer common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "denyConsumer", _consumer)
}

// DenyConsumer is a paid mutator transaction binding the contract method 0x6693a679.
//
// Solidity: function denyConsumer(address _consumer) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) DenyConsumer(_consumer common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.DenyConsumer(&_RandomNumberOracle.TransactOpts, _consumer)
}

// DenyConsumer is a paid mutator transaction binding the contract method 0x6693a679.
//
// Solidity: function denyConsumer(address _consumer) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) DenyConsumer(_consumer common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.DenyConsumer(&_RandomNumberOracle.TransactOpts, _consumer)
}

// DenyProvider is a paid mutator transaction binding the contract method 0x89d8de66.
//
// Solidity: function denyProvider(address _provider) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) DenyProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "denyProvider", _provider)
}

// DenyProvider is a paid mutator transaction binding the contract method 0x89d8de66.
//
// Solidity: function denyProvider(address _provider) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) DenyProvider(_provider common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.DenyProvider(&_RandomNumberOracle.TransactOpts, _provider)
}

// DenyProvider is a paid mutator transaction binding the contract method 0x89d8de66.
//
// Solidity: function denyProvider(address _provider) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) DenyProvider(_provider common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.DenyProvider(&_RandomNumberOracle.TransactOpts, _provider)
}

// FulfillRandomNumber is a paid mutator transaction binding the contract method 0x07dfe505.
//
// Solidity: function fulfillRandomNumber(bytes32 _requestID, uint256 _rnd) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) FulfillRandomNumber(opts *bind.TransactOpts, _requestID [32]byte, _rnd *big.Int) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "fulfillRandomNumber", _requestID, _rnd)
}

// FulfillRandomNumber is a paid mutator transaction binding the contract method 0x07dfe505.
//
// Solidity: function fulfillRandomNumber(bytes32 _requestID, uint256 _rnd) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) FulfillRandomNumber(_requestID [32]byte, _rnd *big.Int) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.FulfillRandomNumber(&_RandomNumberOracle.TransactOpts, _requestID, _rnd)
}

// FulfillRandomNumber is a paid mutator transaction binding the contract method 0x07dfe505.
//
// Solidity: function fulfillRandomNumber(bytes32 _requestID, uint256 _rnd) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) FulfillRandomNumber(_requestID [32]byte, _rnd *big.Int) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.FulfillRandomNumber(&_RandomNumberOracle.TransactOpts, _requestID, _rnd)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0xbd313d78.
//
// Solidity: function requestRandomNumber(bytes32 _seed) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) RequestRandomNumber(opts *bind.TransactOpts, _seed [32]byte) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "requestRandomNumber", _seed)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0xbd313d78.
//
// Solidity: function requestRandomNumber(bytes32 _seed) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) RequestRandomNumber(_seed [32]byte) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.RequestRandomNumber(&_RandomNumberOracle.TransactOpts, _seed)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0xbd313d78.
//
// Solidity: function requestRandomNumber(bytes32 _seed) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) RequestRandomNumber(_seed [32]byte) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.RequestRandomNumber(&_RandomNumberOracle.TransactOpts, _seed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RandomNumberOracle *RandomNumberOracleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.TransferOwnership(&_RandomNumberOracle.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RandomNumberOracle *RandomNumberOracleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RandomNumberOracle.Contract.TransferOwnership(&_RandomNumberOracle.TransactOpts, _newOwner)
}

// RandomNumberOracleConsumerAllowedIterator is returned from FilterConsumerAllowed and is used to iterate over the raw logs and unpacked data for ConsumerAllowed events raised by the RandomNumberOracle contract.
type RandomNumberOracleConsumerAllowedIterator struct {
	Event *RandomNumberOracleConsumerAllowed // Event containing the contract specifics and raw log

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
func (it *RandomNumberOracleConsumerAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberOracleConsumerAllowed)
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
		it.Event = new(RandomNumberOracleConsumerAllowed)
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
func (it *RandomNumberOracleConsumerAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberOracleConsumerAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberOracleConsumerAllowed represents a ConsumerAllowed event raised by the RandomNumberOracle contract.
type RandomNumberOracleConsumerAllowed struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsumerAllowed is a free log retrieval operation binding the contract event 0xc08fb44c85dce533411712fa7b72ea9148193f2b8160340de71713dc199a0770.
//
// Solidity: event ConsumerAllowed(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) FilterConsumerAllowed(opts *bind.FilterOpts) (*RandomNumberOracleConsumerAllowedIterator, error) {

	logs, sub, err := _RandomNumberOracle.contract.FilterLogs(opts, "ConsumerAllowed")
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleConsumerAllowedIterator{contract: _RandomNumberOracle.contract, event: "ConsumerAllowed", logs: logs, sub: sub}, nil
}

// WatchConsumerAllowed is a free log subscription operation binding the contract event 0xc08fb44c85dce533411712fa7b72ea9148193f2b8160340de71713dc199a0770.
//
// Solidity: event ConsumerAllowed(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) WatchConsumerAllowed(opts *bind.WatchOpts, sink chan<- *RandomNumberOracleConsumerAllowed) (event.Subscription, error) {

	logs, sub, err := _RandomNumberOracle.contract.WatchLogs(opts, "ConsumerAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberOracleConsumerAllowed)
				if err := _RandomNumberOracle.contract.UnpackLog(event, "ConsumerAllowed", log); err != nil {
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

// ParseConsumerAllowed is a log parse operation binding the contract event 0xc08fb44c85dce533411712fa7b72ea9148193f2b8160340de71713dc199a0770.
//
// Solidity: event ConsumerAllowed(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) ParseConsumerAllowed(log types.Log) (*RandomNumberOracleConsumerAllowed, error) {
	event := new(RandomNumberOracleConsumerAllowed)
	if err := _RandomNumberOracle.contract.UnpackLog(event, "ConsumerAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomNumberOracleConsumerDeniedIterator is returned from FilterConsumerDenied and is used to iterate over the raw logs and unpacked data for ConsumerDenied events raised by the RandomNumberOracle contract.
type RandomNumberOracleConsumerDeniedIterator struct {
	Event *RandomNumberOracleConsumerDenied // Event containing the contract specifics and raw log

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
func (it *RandomNumberOracleConsumerDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberOracleConsumerDenied)
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
		it.Event = new(RandomNumberOracleConsumerDenied)
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
func (it *RandomNumberOracleConsumerDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberOracleConsumerDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberOracleConsumerDenied represents a ConsumerDenied event raised by the RandomNumberOracle contract.
type RandomNumberOracleConsumerDenied struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsumerDenied is a free log retrieval operation binding the contract event 0xda10ec645b6dc1a68bcbe277a885d3882328aef5e916b06b07a45dd51a374ad8.
//
// Solidity: event ConsumerDenied(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) FilterConsumerDenied(opts *bind.FilterOpts) (*RandomNumberOracleConsumerDeniedIterator, error) {

	logs, sub, err := _RandomNumberOracle.contract.FilterLogs(opts, "ConsumerDenied")
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleConsumerDeniedIterator{contract: _RandomNumberOracle.contract, event: "ConsumerDenied", logs: logs, sub: sub}, nil
}

// WatchConsumerDenied is a free log subscription operation binding the contract event 0xda10ec645b6dc1a68bcbe277a885d3882328aef5e916b06b07a45dd51a374ad8.
//
// Solidity: event ConsumerDenied(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) WatchConsumerDenied(opts *bind.WatchOpts, sink chan<- *RandomNumberOracleConsumerDenied) (event.Subscription, error) {

	logs, sub, err := _RandomNumberOracle.contract.WatchLogs(opts, "ConsumerDenied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberOracleConsumerDenied)
				if err := _RandomNumberOracle.contract.UnpackLog(event, "ConsumerDenied", log); err != nil {
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

// ParseConsumerDenied is a log parse operation binding the contract event 0xda10ec645b6dc1a68bcbe277a885d3882328aef5e916b06b07a45dd51a374ad8.
//
// Solidity: event ConsumerDenied(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) ParseConsumerDenied(log types.Log) (*RandomNumberOracleConsumerDenied, error) {
	event := new(RandomNumberOracleConsumerDenied)
	if err := _RandomNumberOracle.contract.UnpackLog(event, "ConsumerDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomNumberOracleOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the RandomNumberOracle contract.
type RandomNumberOracleOwnerChangedIterator struct {
	Event *RandomNumberOracleOwnerChanged // Event containing the contract specifics and raw log

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
func (it *RandomNumberOracleOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberOracleOwnerChanged)
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
		it.Event = new(RandomNumberOracleOwnerChanged)
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
func (it *RandomNumberOracleOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberOracleOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberOracleOwnerChanged represents a OwnerChanged event raised by the RandomNumberOracle contract.
type RandomNumberOracleOwnerChanged struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address newOwner)
func (_RandomNumberOracle *RandomNumberOracleFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*RandomNumberOracleOwnerChangedIterator, error) {

	logs, sub, err := _RandomNumberOracle.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleOwnerChangedIterator{contract: _RandomNumberOracle.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address newOwner)
func (_RandomNumberOracle *RandomNumberOracleFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *RandomNumberOracleOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _RandomNumberOracle.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberOracleOwnerChanged)
				if err := _RandomNumberOracle.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address newOwner)
func (_RandomNumberOracle *RandomNumberOracleFilterer) ParseOwnerChanged(log types.Log) (*RandomNumberOracleOwnerChanged, error) {
	event := new(RandomNumberOracleOwnerChanged)
	if err := _RandomNumberOracle.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomNumberOracleProviderAllowedIterator is returned from FilterProviderAllowed and is used to iterate over the raw logs and unpacked data for ProviderAllowed events raised by the RandomNumberOracle contract.
type RandomNumberOracleProviderAllowedIterator struct {
	Event *RandomNumberOracleProviderAllowed // Event containing the contract specifics and raw log

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
func (it *RandomNumberOracleProviderAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberOracleProviderAllowed)
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
		it.Event = new(RandomNumberOracleProviderAllowed)
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
func (it *RandomNumberOracleProviderAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberOracleProviderAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberOracleProviderAllowed represents a ProviderAllowed event raised by the RandomNumberOracle contract.
type RandomNumberOracleProviderAllowed struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderAllowed is a free log retrieval operation binding the contract event 0x538bcc00b3cd482832cd1f1c9322147fde519bd7cca43c202c7f0375b3845ac0.
//
// Solidity: event ProviderAllowed(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) FilterProviderAllowed(opts *bind.FilterOpts) (*RandomNumberOracleProviderAllowedIterator, error) {

	logs, sub, err := _RandomNumberOracle.contract.FilterLogs(opts, "ProviderAllowed")
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleProviderAllowedIterator{contract: _RandomNumberOracle.contract, event: "ProviderAllowed", logs: logs, sub: sub}, nil
}

// WatchProviderAllowed is a free log subscription operation binding the contract event 0x538bcc00b3cd482832cd1f1c9322147fde519bd7cca43c202c7f0375b3845ac0.
//
// Solidity: event ProviderAllowed(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) WatchProviderAllowed(opts *bind.WatchOpts, sink chan<- *RandomNumberOracleProviderAllowed) (event.Subscription, error) {

	logs, sub, err := _RandomNumberOracle.contract.WatchLogs(opts, "ProviderAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberOracleProviderAllowed)
				if err := _RandomNumberOracle.contract.UnpackLog(event, "ProviderAllowed", log); err != nil {
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

// ParseProviderAllowed is a log parse operation binding the contract event 0x538bcc00b3cd482832cd1f1c9322147fde519bd7cca43c202c7f0375b3845ac0.
//
// Solidity: event ProviderAllowed(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) ParseProviderAllowed(log types.Log) (*RandomNumberOracleProviderAllowed, error) {
	event := new(RandomNumberOracleProviderAllowed)
	if err := _RandomNumberOracle.contract.UnpackLog(event, "ProviderAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomNumberOracleProviderDeniedIterator is returned from FilterProviderDenied and is used to iterate over the raw logs and unpacked data for ProviderDenied events raised by the RandomNumberOracle contract.
type RandomNumberOracleProviderDeniedIterator struct {
	Event *RandomNumberOracleProviderDenied // Event containing the contract specifics and raw log

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
func (it *RandomNumberOracleProviderDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberOracleProviderDenied)
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
		it.Event = new(RandomNumberOracleProviderDenied)
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
func (it *RandomNumberOracleProviderDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberOracleProviderDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberOracleProviderDenied represents a ProviderDenied event raised by the RandomNumberOracle contract.
type RandomNumberOracleProviderDenied struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderDenied is a free log retrieval operation binding the contract event 0x47de4abe86fecbf25e1c3c5468d05e3db8f3cbb1fdea31c5e278fdb004e38d71.
//
// Solidity: event ProviderDenied(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) FilterProviderDenied(opts *bind.FilterOpts) (*RandomNumberOracleProviderDeniedIterator, error) {

	logs, sub, err := _RandomNumberOracle.contract.FilterLogs(opts, "ProviderDenied")
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleProviderDeniedIterator{contract: _RandomNumberOracle.contract, event: "ProviderDenied", logs: logs, sub: sub}, nil
}

// WatchProviderDenied is a free log subscription operation binding the contract event 0x47de4abe86fecbf25e1c3c5468d05e3db8f3cbb1fdea31c5e278fdb004e38d71.
//
// Solidity: event ProviderDenied(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) WatchProviderDenied(opts *bind.WatchOpts, sink chan<- *RandomNumberOracleProviderDenied) (event.Subscription, error) {

	logs, sub, err := _RandomNumberOracle.contract.WatchLogs(opts, "ProviderDenied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberOracleProviderDenied)
				if err := _RandomNumberOracle.contract.UnpackLog(event, "ProviderDenied", log); err != nil {
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

// ParseProviderDenied is a log parse operation binding the contract event 0x47de4abe86fecbf25e1c3c5468d05e3db8f3cbb1fdea31c5e278fdb004e38d71.
//
// Solidity: event ProviderDenied(address provider)
func (_RandomNumberOracle *RandomNumberOracleFilterer) ParseProviderDenied(log types.Log) (*RandomNumberOracleProviderDenied, error) {
	event := new(RandomNumberOracleProviderDenied)
	if err := _RandomNumberOracle.contract.UnpackLog(event, "ProviderDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomNumberOracleRandomNumberRequestedIterator is returned from FilterRandomNumberRequested and is used to iterate over the raw logs and unpacked data for RandomNumberRequested events raised by the RandomNumberOracle contract.
type RandomNumberOracleRandomNumberRequestedIterator struct {
	Event *RandomNumberOracleRandomNumberRequested // Event containing the contract specifics and raw log

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
func (it *RandomNumberOracleRandomNumberRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberOracleRandomNumberRequested)
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
		it.Event = new(RandomNumberOracleRandomNumberRequested)
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
func (it *RandomNumberOracleRandomNumberRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberOracleRandomNumberRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberOracleRandomNumberRequested represents a RandomNumberRequested event raised by the RandomNumberOracle contract.
type RandomNumberOracleRandomNumberRequested struct {
	RequestID [32]byte
	Seed      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomNumberRequested is a free log retrieval operation binding the contract event 0xac2e43d9741627d0f2e7a61dba4f97dfa56414d39e787163b0e6dbde34e3a6b2.
//
// Solidity: event RandomNumberRequested(bytes32 requestID, bytes32 seed)
func (_RandomNumberOracle *RandomNumberOracleFilterer) FilterRandomNumberRequested(opts *bind.FilterOpts) (*RandomNumberOracleRandomNumberRequestedIterator, error) {

	logs, sub, err := _RandomNumberOracle.contract.FilterLogs(opts, "RandomNumberRequested")
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleRandomNumberRequestedIterator{contract: _RandomNumberOracle.contract, event: "RandomNumberRequested", logs: logs, sub: sub}, nil
}

// WatchRandomNumberRequested is a free log subscription operation binding the contract event 0xac2e43d9741627d0f2e7a61dba4f97dfa56414d39e787163b0e6dbde34e3a6b2.
//
// Solidity: event RandomNumberRequested(bytes32 requestID, bytes32 seed)
func (_RandomNumberOracle *RandomNumberOracleFilterer) WatchRandomNumberRequested(opts *bind.WatchOpts, sink chan<- *RandomNumberOracleRandomNumberRequested) (event.Subscription, error) {

	logs, sub, err := _RandomNumberOracle.contract.WatchLogs(opts, "RandomNumberRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberOracleRandomNumberRequested)
				if err := _RandomNumberOracle.contract.UnpackLog(event, "RandomNumberRequested", log); err != nil {
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

// ParseRandomNumberRequested is a log parse operation binding the contract event 0xac2e43d9741627d0f2e7a61dba4f97dfa56414d39e787163b0e6dbde34e3a6b2.
//
// Solidity: event RandomNumberRequested(bytes32 requestID, bytes32 seed)
func (_RandomNumberOracle *RandomNumberOracleFilterer) ParseRandomNumberRequested(log types.Log) (*RandomNumberOracleRandomNumberRequested, error) {
	event := new(RandomNumberOracleRandomNumberRequested)
	if err := _RandomNumberOracle.contract.UnpackLog(event, "RandomNumberRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomNumberOracleRequestCanceledIterator is returned from FilterRequestCanceled and is used to iterate over the raw logs and unpacked data for RequestCanceled events raised by the RandomNumberOracle contract.
type RandomNumberOracleRequestCanceledIterator struct {
	Event *RandomNumberOracleRequestCanceled // Event containing the contract specifics and raw log

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
func (it *RandomNumberOracleRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberOracleRequestCanceled)
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
		it.Event = new(RandomNumberOracleRequestCanceled)
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
func (it *RandomNumberOracleRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberOracleRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberOracleRequestCanceled represents a RequestCanceled event raised by the RandomNumberOracle contract.
type RandomNumberOracleRequestCanceled struct {
	RequestID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestCanceled is a free log retrieval operation binding the contract event 0x7efba54574fc3f1379813d6e09d06b5e9913bd264f1468a34879d577d2110f54.
//
// Solidity: event RequestCanceled(bytes32 requestID)
func (_RandomNumberOracle *RandomNumberOracleFilterer) FilterRequestCanceled(opts *bind.FilterOpts) (*RandomNumberOracleRequestCanceledIterator, error) {

	logs, sub, err := _RandomNumberOracle.contract.FilterLogs(opts, "RequestCanceled")
	if err != nil {
		return nil, err
	}
	return &RandomNumberOracleRequestCanceledIterator{contract: _RandomNumberOracle.contract, event: "RequestCanceled", logs: logs, sub: sub}, nil
}

// WatchRequestCanceled is a free log subscription operation binding the contract event 0x7efba54574fc3f1379813d6e09d06b5e9913bd264f1468a34879d577d2110f54.
//
// Solidity: event RequestCanceled(bytes32 requestID)
func (_RandomNumberOracle *RandomNumberOracleFilterer) WatchRequestCanceled(opts *bind.WatchOpts, sink chan<- *RandomNumberOracleRequestCanceled) (event.Subscription, error) {

	logs, sub, err := _RandomNumberOracle.contract.WatchLogs(opts, "RequestCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberOracleRequestCanceled)
				if err := _RandomNumberOracle.contract.UnpackLog(event, "RequestCanceled", log); err != nil {
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

// ParseRequestCanceled is a log parse operation binding the contract event 0x7efba54574fc3f1379813d6e09d06b5e9913bd264f1468a34879d577d2110f54.
//
// Solidity: event RequestCanceled(bytes32 requestID)
func (_RandomNumberOracle *RandomNumberOracleFilterer) ParseRequestCanceled(log types.Log) (*RandomNumberOracleRequestCanceled, error) {
	event := new(RandomNumberOracleRequestCanceled)
	if err := _RandomNumberOracle.contract.UnpackLog(event, "RequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
