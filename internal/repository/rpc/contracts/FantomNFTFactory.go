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

// FantomNFTFactoryMetaData contains all meta data concerning the FantomNFTFactory contract.
var FantomNFTFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createNFTContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"}],\"name\":\"updateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"}],\"name\":\"updateBundleMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"}],\"name\":\"updateMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintFee\",\"type\":\"uint256\"}],\"name\":\"updateMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomNFTFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomNFTFactoryMetaData.ABI instead.
var FantomNFTFactoryABI = FantomNFTFactoryMetaData.ABI

// FantomNFTFactory is an auto generated Go binding around an Ethereum contract.
type FantomNFTFactory struct {
	FantomNFTFactoryCaller     // Read-only binding to the contract
	FantomNFTFactoryTransactor // Write-only binding to the contract
	FantomNFTFactoryFilterer   // Log filterer for contract events
}

// FantomNFTFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomNFTFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomNFTFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomNFTFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomNFTFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomNFTFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomNFTFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomNFTFactorySession struct {
	Contract     *FantomNFTFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FantomNFTFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomNFTFactoryCallerSession struct {
	Contract *FantomNFTFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FantomNFTFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomNFTFactoryTransactorSession struct {
	Contract     *FantomNFTFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FantomNFTFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomNFTFactoryRaw struct {
	Contract *FantomNFTFactory // Generic contract binding to access the raw methods on
}

// FantomNFTFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomNFTFactoryCallerRaw struct {
	Contract *FantomNFTFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FantomNFTFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomNFTFactoryTransactorRaw struct {
	Contract *FantomNFTFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomNFTFactory creates a new instance of FantomNFTFactory, bound to a specific deployed contract.
func NewFantomNFTFactory(address common.Address, backend bind.ContractBackend) (*FantomNFTFactory, error) {
	contract, err := bindFantomNFTFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomNFTFactory{FantomNFTFactoryCaller: FantomNFTFactoryCaller{contract: contract}, FantomNFTFactoryTransactor: FantomNFTFactoryTransactor{contract: contract}, FantomNFTFactoryFilterer: FantomNFTFactoryFilterer{contract: contract}}, nil
}

// NewFantomNFTFactoryCaller creates a new read-only instance of FantomNFTFactory, bound to a specific deployed contract.
func NewFantomNFTFactoryCaller(address common.Address, caller bind.ContractCaller) (*FantomNFTFactoryCaller, error) {
	contract, err := bindFantomNFTFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomNFTFactoryCaller{contract: contract}, nil
}

// NewFantomNFTFactoryTransactor creates a new write-only instance of FantomNFTFactory, bound to a specific deployed contract.
func NewFantomNFTFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomNFTFactoryTransactor, error) {
	contract, err := bindFantomNFTFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomNFTFactoryTransactor{contract: contract}, nil
}

// NewFantomNFTFactoryFilterer creates a new log filterer instance of FantomNFTFactory, bound to a specific deployed contract.
func NewFantomNFTFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomNFTFactoryFilterer, error) {
	contract, err := bindFantomNFTFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomNFTFactoryFilterer{contract: contract}, nil
}

// bindFantomNFTFactory binds a generic wrapper to an already deployed contract.
func bindFantomNFTFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomNFTFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomNFTFactory *FantomNFTFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomNFTFactory.Contract.FantomNFTFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomNFTFactory *FantomNFTFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.FantomNFTFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomNFTFactory *FantomNFTFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.FantomNFTFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomNFTFactory *FantomNFTFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomNFTFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomNFTFactory *FantomNFTFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomNFTFactory *FantomNFTFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "auction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_FantomNFTFactory *FantomNFTFactorySession) Auction() (common.Address, error) {
	return _FantomNFTFactory.Contract.Auction(&_FantomNFTFactory.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) Auction() (common.Address, error) {
	return _FantomNFTFactory.Contract.Auction(&_FantomNFTFactory.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCaller) BundleMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "bundleMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomNFTFactory *FantomNFTFactorySession) BundleMarketplace() (common.Address, error) {
	return _FantomNFTFactory.Contract.BundleMarketplace(&_FantomNFTFactory.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) BundleMarketplace() (common.Address, error) {
	return _FantomNFTFactory.Contract.BundleMarketplace(&_FantomNFTFactory.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_FantomNFTFactory *FantomNFTFactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_FantomNFTFactory *FantomNFTFactorySession) Exists(arg0 common.Address) (bool, error) {
	return _FantomNFTFactory.Contract.Exists(&_FantomNFTFactory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _FantomNFTFactory.Contract.Exists(&_FantomNFTFactory.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FantomNFTFactory *FantomNFTFactorySession) FeeRecipient() (common.Address, error) {
	return _FantomNFTFactory.Contract.FeeRecipient(&_FantomNFTFactory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) FeeRecipient() (common.Address, error) {
	return _FantomNFTFactory.Contract.FeeRecipient(&_FantomNFTFactory.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomNFTFactory *FantomNFTFactorySession) Marketplace() (common.Address, error) {
	return _FantomNFTFactory.Contract.Marketplace(&_FantomNFTFactory.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) Marketplace() (common.Address, error) {
	return _FantomNFTFactory.Contract.Marketplace(&_FantomNFTFactory.CallOpts)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_FantomNFTFactory *FantomNFTFactoryCaller) MintFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "mintFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_FantomNFTFactory *FantomNFTFactorySession) MintFee() (*big.Int, error) {
	return _FantomNFTFactory.Contract.MintFee(&_FantomNFTFactory.CallOpts)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) MintFee() (*big.Int, error) {
	return _FantomNFTFactory.Contract.MintFee(&_FantomNFTFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomNFTFactory *FantomNFTFactorySession) Owner() (common.Address, error) {
	return _FantomNFTFactory.Contract.Owner(&_FantomNFTFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) Owner() (common.Address, error) {
	return _FantomNFTFactory.Contract.Owner(&_FantomNFTFactory.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomNFTFactory *FantomNFTFactoryCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomNFTFactory.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomNFTFactory *FantomNFTFactorySession) PlatformFee() (*big.Int, error) {
	return _FantomNFTFactory.Contract.PlatformFee(&_FantomNFTFactory.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomNFTFactory *FantomNFTFactoryCallerSession) PlatformFee() (*big.Int, error) {
	return _FantomNFTFactory.Contract.PlatformFee(&_FantomNFTFactory.CallOpts)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x14d76ffc.
//
// Solidity: function createNFTContract(string _name, string _symbol) payable returns(address)
func (_FantomNFTFactory *FantomNFTFactoryTransactor) CreateNFTContract(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "createNFTContract", _name, _symbol)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x14d76ffc.
//
// Solidity: function createNFTContract(string _name, string _symbol) payable returns(address)
func (_FantomNFTFactory *FantomNFTFactorySession) CreateNFTContract(_name string, _symbol string) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.CreateNFTContract(&_FantomNFTFactory.TransactOpts, _name, _symbol)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x14d76ffc.
//
// Solidity: function createNFTContract(string _name, string _symbol) payable returns(address)
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) CreateNFTContract(_name string, _symbol string) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.CreateNFTContract(&_FantomNFTFactory.TransactOpts, _name, _symbol)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomNFTFactory *FantomNFTFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.RenounceOwnership(&_FantomNFTFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.RenounceOwnership(&_FantomNFTFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomNFTFactory *FantomNFTFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.TransferOwnership(&_FantomNFTFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.TransferOwnership(&_FantomNFTFactory.TransactOpts, newOwner)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) UpdateAuction(opts *bind.TransactOpts, _auction common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "updateAuction", _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_FantomNFTFactory *FantomNFTFactorySession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateAuction(&_FantomNFTFactory.TransactOpts, _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateAuction(&_FantomNFTFactory.TransactOpts, _auction)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) UpdateBundleMarketplace(opts *bind.TransactOpts, _bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "updateBundleMarketplace", _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomNFTFactory *FantomNFTFactorySession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateBundleMarketplace(&_FantomNFTFactory.TransactOpts, _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateBundleMarketplace(&_FantomNFTFactory.TransactOpts, _bundleMarketplace)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "updateFeeRecipient", _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_FantomNFTFactory *FantomNFTFactorySession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateFeeRecipient(&_FantomNFTFactory.TransactOpts, _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateFeeRecipient(&_FantomNFTFactory.TransactOpts, _feeRecipient)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) UpdateMarketplace(opts *bind.TransactOpts, _marketplace common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "updateMarketplace", _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomNFTFactory *FantomNFTFactorySession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateMarketplace(&_FantomNFTFactory.TransactOpts, _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateMarketplace(&_FantomNFTFactory.TransactOpts, _marketplace)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x84017e52.
//
// Solidity: function updateMintFee(uint256 _mintFee) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) UpdateMintFee(opts *bind.TransactOpts, _mintFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "updateMintFee", _mintFee)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x84017e52.
//
// Solidity: function updateMintFee(uint256 _mintFee) returns()
func (_FantomNFTFactory *FantomNFTFactorySession) UpdateMintFee(_mintFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateMintFee(&_FantomNFTFactory.TransactOpts, _mintFee)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x84017e52.
//
// Solidity: function updateMintFee(uint256 _mintFee) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) UpdateMintFee(_mintFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdateMintFee(&_FantomNFTFactory.TransactOpts, _mintFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTFactory.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomNFTFactory *FantomNFTFactorySession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdatePlatformFee(&_FantomNFTFactory.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomNFTFactory *FantomNFTFactoryTransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTFactory.Contract.UpdatePlatformFee(&_FantomNFTFactory.TransactOpts, _platformFee)
}

// FantomNFTFactoryContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the FantomNFTFactory contract.
type FantomNFTFactoryContractCreatedIterator struct {
	Event *FantomNFTFactoryContractCreated // Event containing the contract specifics and raw log

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
func (it *FantomNFTFactoryContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTFactoryContractCreated)
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
		it.Event = new(FantomNFTFactoryContractCreated)
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
func (it *FantomNFTFactoryContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTFactoryContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTFactoryContractCreated represents a ContractCreated event raised by the FantomNFTFactory contract.
type FantomNFTFactoryContractCreated struct {
	Creator common.Address
	Nft     common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b.
//
// Solidity: event ContractCreated(address creator, address nft)
func (_FantomNFTFactory *FantomNFTFactoryFilterer) FilterContractCreated(opts *bind.FilterOpts) (*FantomNFTFactoryContractCreatedIterator, error) {

	logs, sub, err := _FantomNFTFactory.contract.FilterLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return &FantomNFTFactoryContractCreatedIterator{contract: _FantomNFTFactory.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b.
//
// Solidity: event ContractCreated(address creator, address nft)
func (_FantomNFTFactory *FantomNFTFactoryFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *FantomNFTFactoryContractCreated) (event.Subscription, error) {

	logs, sub, err := _FantomNFTFactory.contract.WatchLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTFactoryContractCreated)
				if err := _FantomNFTFactory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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
func (_FantomNFTFactory *FantomNFTFactoryFilterer) ParseContractCreated(log types.Log) (*FantomNFTFactoryContractCreated, error) {
	event := new(FantomNFTFactoryContractCreated)
	if err := _FantomNFTFactory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomNFTFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomNFTFactory contract.
type FantomNFTFactoryOwnershipTransferredIterator struct {
	Event *FantomNFTFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomNFTFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTFactoryOwnershipTransferred)
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
		it.Event = new(FantomNFTFactoryOwnershipTransferred)
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
func (it *FantomNFTFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the FantomNFTFactory contract.
type FantomNFTFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomNFTFactory *FantomNFTFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomNFTFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomNFTFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomNFTFactoryOwnershipTransferredIterator{contract: _FantomNFTFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomNFTFactory *FantomNFTFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomNFTFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomNFTFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTFactoryOwnershipTransferred)
				if err := _FantomNFTFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomNFTFactory *FantomNFTFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*FantomNFTFactoryOwnershipTransferred, error) {
	event := new(FantomNFTFactoryOwnershipTransferred)
	if err := _FantomNFTFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
