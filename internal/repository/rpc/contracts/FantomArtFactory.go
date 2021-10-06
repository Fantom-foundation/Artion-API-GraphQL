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

// FantomArtFactoryMetaData contains all meta data concerning the FantomArtFactory contract.
var FantomArtFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bundleMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createNFTContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"}],\"name\":\"updateBundleMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"}],\"name\":\"updateMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintFee\",\"type\":\"uint256\"}],\"name\":\"updateMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomArtFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomArtFactoryMetaData.ABI instead.
var FantomArtFactoryABI = FantomArtFactoryMetaData.ABI

// FantomArtFactory is an auto generated Go binding around an Ethereum contract.
type FantomArtFactory struct {
	FantomArtFactoryCaller     // Read-only binding to the contract
	FantomArtFactoryTransactor // Write-only binding to the contract
	FantomArtFactoryFilterer   // Log filterer for contract events
}

// FantomArtFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomArtFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomArtFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomArtFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomArtFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomArtFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomArtFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomArtFactorySession struct {
	Contract     *FantomArtFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FantomArtFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomArtFactoryCallerSession struct {
	Contract *FantomArtFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FantomArtFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomArtFactoryTransactorSession struct {
	Contract     *FantomArtFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FantomArtFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomArtFactoryRaw struct {
	Contract *FantomArtFactory // Generic contract binding to access the raw methods on
}

// FantomArtFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomArtFactoryCallerRaw struct {
	Contract *FantomArtFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FantomArtFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomArtFactoryTransactorRaw struct {
	Contract *FantomArtFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomArtFactory creates a new instance of FantomArtFactory, bound to a specific deployed contract.
func NewFantomArtFactory(address common.Address, backend bind.ContractBackend) (*FantomArtFactory, error) {
	contract, err := bindFantomArtFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomArtFactory{FantomArtFactoryCaller: FantomArtFactoryCaller{contract: contract}, FantomArtFactoryTransactor: FantomArtFactoryTransactor{contract: contract}, FantomArtFactoryFilterer: FantomArtFactoryFilterer{contract: contract}}, nil
}

// NewFantomArtFactoryCaller creates a new read-only instance of FantomArtFactory, bound to a specific deployed contract.
func NewFantomArtFactoryCaller(address common.Address, caller bind.ContractCaller) (*FantomArtFactoryCaller, error) {
	contract, err := bindFantomArtFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomArtFactoryCaller{contract: contract}, nil
}

// NewFantomArtFactoryTransactor creates a new write-only instance of FantomArtFactory, bound to a specific deployed contract.
func NewFantomArtFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomArtFactoryTransactor, error) {
	contract, err := bindFantomArtFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomArtFactoryTransactor{contract: contract}, nil
}

// NewFantomArtFactoryFilterer creates a new log filterer instance of FantomArtFactory, bound to a specific deployed contract.
func NewFantomArtFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomArtFactoryFilterer, error) {
	contract, err := bindFantomArtFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomArtFactoryFilterer{contract: contract}, nil
}

// bindFantomArtFactory binds a generic wrapper to an already deployed contract.
func bindFantomArtFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomArtFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomArtFactory *FantomArtFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomArtFactory.Contract.FantomArtFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomArtFactory *FantomArtFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.FantomArtFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomArtFactory *FantomArtFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.FantomArtFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomArtFactory *FantomArtFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomArtFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomArtFactory *FantomArtFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomArtFactory *FantomArtFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.contract.Transact(opts, method, params...)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCaller) BundleMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomArtFactory.contract.Call(opts, &out, "bundleMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomArtFactory *FantomArtFactorySession) BundleMarketplace() (common.Address, error) {
	return _FantomArtFactory.Contract.BundleMarketplace(&_FantomArtFactory.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCallerSession) BundleMarketplace() (common.Address, error) {
	return _FantomArtFactory.Contract.BundleMarketplace(&_FantomArtFactory.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_FantomArtFactory *FantomArtFactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FantomArtFactory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_FantomArtFactory *FantomArtFactorySession) Exists(arg0 common.Address) (bool, error) {
	return _FantomArtFactory.Contract.Exists(&_FantomArtFactory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_FantomArtFactory *FantomArtFactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _FantomArtFactory.Contract.Exists(&_FantomArtFactory.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomArtFactory.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FantomArtFactory *FantomArtFactorySession) FeeRecipient() (common.Address, error) {
	return _FantomArtFactory.Contract.FeeRecipient(&_FantomArtFactory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCallerSession) FeeRecipient() (common.Address, error) {
	return _FantomArtFactory.Contract.FeeRecipient(&_FantomArtFactory.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomArtFactory.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomArtFactory *FantomArtFactorySession) Marketplace() (common.Address, error) {
	return _FantomArtFactory.Contract.Marketplace(&_FantomArtFactory.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCallerSession) Marketplace() (common.Address, error) {
	return _FantomArtFactory.Contract.Marketplace(&_FantomArtFactory.CallOpts)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_FantomArtFactory *FantomArtFactoryCaller) MintFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomArtFactory.contract.Call(opts, &out, "mintFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_FantomArtFactory *FantomArtFactorySession) MintFee() (*big.Int, error) {
	return _FantomArtFactory.Contract.MintFee(&_FantomArtFactory.CallOpts)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_FantomArtFactory *FantomArtFactoryCallerSession) MintFee() (*big.Int, error) {
	return _FantomArtFactory.Contract.MintFee(&_FantomArtFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomArtFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomArtFactory *FantomArtFactorySession) Owner() (common.Address, error) {
	return _FantomArtFactory.Contract.Owner(&_FantomArtFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomArtFactory *FantomArtFactoryCallerSession) Owner() (common.Address, error) {
	return _FantomArtFactory.Contract.Owner(&_FantomArtFactory.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomArtFactory *FantomArtFactoryCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomArtFactory.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomArtFactory *FantomArtFactorySession) PlatformFee() (*big.Int, error) {
	return _FantomArtFactory.Contract.PlatformFee(&_FantomArtFactory.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomArtFactory *FantomArtFactoryCallerSession) PlatformFee() (*big.Int, error) {
	return _FantomArtFactory.Contract.PlatformFee(&_FantomArtFactory.CallOpts)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x14d76ffc.
//
// Solidity: function createNFTContract(string _name, string _symbol) payable returns(address)
func (_FantomArtFactory *FantomArtFactoryTransactor) CreateNFTContract(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "createNFTContract", _name, _symbol)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x14d76ffc.
//
// Solidity: function createNFTContract(string _name, string _symbol) payable returns(address)
func (_FantomArtFactory *FantomArtFactorySession) CreateNFTContract(_name string, _symbol string) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.CreateNFTContract(&_FantomArtFactory.TransactOpts, _name, _symbol)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x14d76ffc.
//
// Solidity: function createNFTContract(string _name, string _symbol) payable returns(address)
func (_FantomArtFactory *FantomArtFactoryTransactorSession) CreateNFTContract(_name string, _symbol string) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.CreateNFTContract(&_FantomArtFactory.TransactOpts, _name, _symbol)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomArtFactory *FantomArtFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomArtFactory *FantomArtFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomArtFactory.Contract.RenounceOwnership(&_FantomArtFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomArtFactory *FantomArtFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomArtFactory.Contract.RenounceOwnership(&_FantomArtFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomArtFactory *FantomArtFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomArtFactory *FantomArtFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.TransferOwnership(&_FantomArtFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomArtFactory *FantomArtFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.TransferOwnership(&_FantomArtFactory.TransactOpts, newOwner)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomArtFactory *FantomArtFactoryTransactor) UpdateBundleMarketplace(opts *bind.TransactOpts, _bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "updateBundleMarketplace", _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomArtFactory *FantomArtFactorySession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateBundleMarketplace(&_FantomArtFactory.TransactOpts, _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomArtFactory *FantomArtFactoryTransactorSession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateBundleMarketplace(&_FantomArtFactory.TransactOpts, _bundleMarketplace)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_FantomArtFactory *FantomArtFactoryTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "updateFeeRecipient", _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_FantomArtFactory *FantomArtFactorySession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateFeeRecipient(&_FantomArtFactory.TransactOpts, _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_FantomArtFactory *FantomArtFactoryTransactorSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateFeeRecipient(&_FantomArtFactory.TransactOpts, _feeRecipient)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomArtFactory *FantomArtFactoryTransactor) UpdateMarketplace(opts *bind.TransactOpts, _marketplace common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "updateMarketplace", _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomArtFactory *FantomArtFactorySession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateMarketplace(&_FantomArtFactory.TransactOpts, _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomArtFactory *FantomArtFactoryTransactorSession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateMarketplace(&_FantomArtFactory.TransactOpts, _marketplace)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x84017e52.
//
// Solidity: function updateMintFee(uint256 _mintFee) returns()
func (_FantomArtFactory *FantomArtFactoryTransactor) UpdateMintFee(opts *bind.TransactOpts, _mintFee *big.Int) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "updateMintFee", _mintFee)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x84017e52.
//
// Solidity: function updateMintFee(uint256 _mintFee) returns()
func (_FantomArtFactory *FantomArtFactorySession) UpdateMintFee(_mintFee *big.Int) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateMintFee(&_FantomArtFactory.TransactOpts, _mintFee)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x84017e52.
//
// Solidity: function updateMintFee(uint256 _mintFee) returns()
func (_FantomArtFactory *FantomArtFactoryTransactorSession) UpdateMintFee(_mintFee *big.Int) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdateMintFee(&_FantomArtFactory.TransactOpts, _mintFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomArtFactory *FantomArtFactoryTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomArtFactory.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomArtFactory *FantomArtFactorySession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdatePlatformFee(&_FantomArtFactory.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomArtFactory *FantomArtFactoryTransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomArtFactory.Contract.UpdatePlatformFee(&_FantomArtFactory.TransactOpts, _platformFee)
}

// FantomArtFactoryContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the FantomArtFactory contract.
type FantomArtFactoryContractCreatedIterator struct {
	Event *FantomArtFactoryContractCreated // Event containing the contract specifics and raw log

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
func (it *FantomArtFactoryContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomArtFactoryContractCreated)
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
		it.Event = new(FantomArtFactoryContractCreated)
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
func (it *FantomArtFactoryContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomArtFactoryContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomArtFactoryContractCreated represents a ContractCreated event raised by the FantomArtFactory contract.
type FantomArtFactoryContractCreated struct {
	Creator common.Address
	Nft     common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b.
//
// Solidity: event ContractCreated(address creator, address nft)
func (_FantomArtFactory *FantomArtFactoryFilterer) FilterContractCreated(opts *bind.FilterOpts) (*FantomArtFactoryContractCreatedIterator, error) {

	logs, sub, err := _FantomArtFactory.contract.FilterLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return &FantomArtFactoryContractCreatedIterator{contract: _FantomArtFactory.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b.
//
// Solidity: event ContractCreated(address creator, address nft)
func (_FantomArtFactory *FantomArtFactoryFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *FantomArtFactoryContractCreated) (event.Subscription, error) {

	logs, sub, err := _FantomArtFactory.contract.WatchLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomArtFactoryContractCreated)
				if err := _FantomArtFactory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ParseContractCreated is a log parse operation binding the contract event 0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b.
//
// Solidity: event ContractCreated(address creator, address nft)
func (_FantomArtFactory *FantomArtFactoryFilterer) ParseContractCreated(log types.Log) (*FantomArtFactoryContractCreated, error) {
	event := new(FantomArtFactoryContractCreated)
	if err := _FantomArtFactory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomArtFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomArtFactory contract.
type FantomArtFactoryOwnershipTransferredIterator struct {
	Event *FantomArtFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomArtFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomArtFactoryOwnershipTransferred)
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
		it.Event = new(FantomArtFactoryOwnershipTransferred)
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
func (it *FantomArtFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomArtFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomArtFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the FantomArtFactory contract.
type FantomArtFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomArtFactory *FantomArtFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomArtFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomArtFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomArtFactoryOwnershipTransferredIterator{contract: _FantomArtFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomArtFactory *FantomArtFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomArtFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomArtFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomArtFactoryOwnershipTransferred)
				if err := _FantomArtFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomArtFactory *FantomArtFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*FantomArtFactoryOwnershipTransferred, error) {
	event := new(FantomArtFactoryOwnershipTransferred)
	if err := _FantomArtFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
