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

// FantomRoyaltyRegistryMetaData contains all meta data concerning the FantomRoyaltyRegistry contract.
var FantomRoyaltyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyPercent\",\"type\":\"uint16\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyPercent\",\"type\":\"uint16\"}],\"name\":\"setRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomRoyaltyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomRoyaltyRegistryMetaData.ABI instead.
var FantomRoyaltyRegistryABI = FantomRoyaltyRegistryMetaData.ABI

// FantomRoyaltyRegistry is an auto generated Go binding around an Ethereum contract.
type FantomRoyaltyRegistry struct {
	FantomRoyaltyRegistryCaller     // Read-only binding to the contract
	FantomRoyaltyRegistryTransactor // Write-only binding to the contract
	FantomRoyaltyRegistryFilterer   // Log filterer for contract events
}

// FantomRoyaltyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomRoyaltyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomRoyaltyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomRoyaltyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomRoyaltyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomRoyaltyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomRoyaltyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomRoyaltyRegistrySession struct {
	Contract     *FantomRoyaltyRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FantomRoyaltyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomRoyaltyRegistryCallerSession struct {
	Contract *FantomRoyaltyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// FantomRoyaltyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomRoyaltyRegistryTransactorSession struct {
	Contract     *FantomRoyaltyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// FantomRoyaltyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomRoyaltyRegistryRaw struct {
	Contract *FantomRoyaltyRegistry // Generic contract binding to access the raw methods on
}

// FantomRoyaltyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomRoyaltyRegistryCallerRaw struct {
	Contract *FantomRoyaltyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FantomRoyaltyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomRoyaltyRegistryTransactorRaw struct {
	Contract *FantomRoyaltyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomRoyaltyRegistry creates a new instance of FantomRoyaltyRegistry, bound to a specific deployed contract.
func NewFantomRoyaltyRegistry(address common.Address, backend bind.ContractBackend) (*FantomRoyaltyRegistry, error) {
	contract, err := bindFantomRoyaltyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomRoyaltyRegistry{FantomRoyaltyRegistryCaller: FantomRoyaltyRegistryCaller{contract: contract}, FantomRoyaltyRegistryTransactor: FantomRoyaltyRegistryTransactor{contract: contract}, FantomRoyaltyRegistryFilterer: FantomRoyaltyRegistryFilterer{contract: contract}}, nil
}

// NewFantomRoyaltyRegistryCaller creates a new read-only instance of FantomRoyaltyRegistry, bound to a specific deployed contract.
func NewFantomRoyaltyRegistryCaller(address common.Address, caller bind.ContractCaller) (*FantomRoyaltyRegistryCaller, error) {
	contract, err := bindFantomRoyaltyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomRoyaltyRegistryCaller{contract: contract}, nil
}

// NewFantomRoyaltyRegistryTransactor creates a new write-only instance of FantomRoyaltyRegistry, bound to a specific deployed contract.
func NewFantomRoyaltyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomRoyaltyRegistryTransactor, error) {
	contract, err := bindFantomRoyaltyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomRoyaltyRegistryTransactor{contract: contract}, nil
}

// NewFantomRoyaltyRegistryFilterer creates a new log filterer instance of FantomRoyaltyRegistry, bound to a specific deployed contract.
func NewFantomRoyaltyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomRoyaltyRegistryFilterer, error) {
	contract, err := bindFantomRoyaltyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomRoyaltyRegistryFilterer{contract: contract}, nil
}

// bindFantomRoyaltyRegistry binds a generic wrapper to an already deployed contract.
func bindFantomRoyaltyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomRoyaltyRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomRoyaltyRegistry.Contract.FantomRoyaltyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.FantomRoyaltyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.FantomRoyaltyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomRoyaltyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.contract.Transact(opts, method, params...)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x538bd5ea.
//
// Solidity: function royaltyInfo(address _collection, uint256 _tokenId, uint256 _salePrice) view returns(address _receiver, uint256 _royaltyAmount)
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryCaller) RoyaltyInfo(opts *bind.CallOpts, _collection common.Address, _tokenId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _FantomRoyaltyRegistry.contract.Call(opts, &out, "royaltyInfo", _collection, _tokenId, _salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x538bd5ea.
//
// Solidity: function royaltyInfo(address _collection, uint256 _tokenId, uint256 _salePrice) view returns(address _receiver, uint256 _royaltyAmount)
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistrySession) RoyaltyInfo(_collection common.Address, _tokenId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _FantomRoyaltyRegistry.Contract.RoyaltyInfo(&_FantomRoyaltyRegistry.CallOpts, _collection, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x538bd5ea.
//
// Solidity: function royaltyInfo(address _collection, uint256 _tokenId, uint256 _salePrice) view returns(address _receiver, uint256 _royaltyAmount)
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryCallerSession) RoyaltyInfo(_collection common.Address, _tokenId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _FantomRoyaltyRegistry.Contract.RoyaltyInfo(&_FantomRoyaltyRegistry.CallOpts, _collection, _tokenId, _salePrice)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x1149d95e.
//
// Solidity: function setDefaultRoyalty(address _collection, address _receiver, uint16 _royaltyPercent) returns()
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryTransactor) SetDefaultRoyalty(opts *bind.TransactOpts, _collection common.Address, _receiver common.Address, _royaltyPercent uint16) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.contract.Transact(opts, "setDefaultRoyalty", _collection, _receiver, _royaltyPercent)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x1149d95e.
//
// Solidity: function setDefaultRoyalty(address _collection, address _receiver, uint16 _royaltyPercent) returns()
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistrySession) SetDefaultRoyalty(_collection common.Address, _receiver common.Address, _royaltyPercent uint16) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.SetDefaultRoyalty(&_FantomRoyaltyRegistry.TransactOpts, _collection, _receiver, _royaltyPercent)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x1149d95e.
//
// Solidity: function setDefaultRoyalty(address _collection, address _receiver, uint16 _royaltyPercent) returns()
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryTransactorSession) SetDefaultRoyalty(_collection common.Address, _receiver common.Address, _royaltyPercent uint16) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.SetDefaultRoyalty(&_FantomRoyaltyRegistry.TransactOpts, _collection, _receiver, _royaltyPercent)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x53059d78.
//
// Solidity: function setRoyalty(address _collection, uint256 _tokenId, address _receiver, uint16 _royaltyPercent) returns()
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryTransactor) SetRoyalty(opts *bind.TransactOpts, _collection common.Address, _tokenId *big.Int, _receiver common.Address, _royaltyPercent uint16) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.contract.Transact(opts, "setRoyalty", _collection, _tokenId, _receiver, _royaltyPercent)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x53059d78.
//
// Solidity: function setRoyalty(address _collection, uint256 _tokenId, address _receiver, uint16 _royaltyPercent) returns()
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistrySession) SetRoyalty(_collection common.Address, _tokenId *big.Int, _receiver common.Address, _royaltyPercent uint16) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.SetRoyalty(&_FantomRoyaltyRegistry.TransactOpts, _collection, _tokenId, _receiver, _royaltyPercent)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x53059d78.
//
// Solidity: function setRoyalty(address _collection, uint256 _tokenId, address _receiver, uint16 _royaltyPercent) returns()
func (_FantomRoyaltyRegistry *FantomRoyaltyRegistryTransactorSession) SetRoyalty(_collection common.Address, _tokenId *big.Int, _receiver common.Address, _royaltyPercent uint16) (*types.Transaction, error) {
	return _FantomRoyaltyRegistry.Contract.SetRoyalty(&_FantomRoyaltyRegistry.TransactOpts, _collection, _tokenId, _receiver, _royaltyPercent)
}
