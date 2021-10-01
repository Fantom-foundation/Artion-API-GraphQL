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

// FantomNFTTradableMetaData contains all meta data concerning the FantomNFTTradable contract.
var FantomNFTTradableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeReceipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"UpdateFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeReceipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomNFTTradableABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomNFTTradableMetaData.ABI instead.
var FantomNFTTradableABI = FantomNFTTradableMetaData.ABI

// FantomNFTTradable is an auto generated Go binding around an Ethereum contract.
type FantomNFTTradable struct {
	FantomNFTTradableCaller     // Read-only binding to the contract
	FantomNFTTradableTransactor // Write-only binding to the contract
	FantomNFTTradableFilterer   // Log filterer for contract events
}

// FantomNFTTradableCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomNFTTradableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomNFTTradableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomNFTTradableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomNFTTradableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomNFTTradableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomNFTTradableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomNFTTradableSession struct {
	Contract     *FantomNFTTradable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FantomNFTTradableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomNFTTradableCallerSession struct {
	Contract *FantomNFTTradableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FantomNFTTradableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomNFTTradableTransactorSession struct {
	Contract     *FantomNFTTradableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FantomNFTTradableRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomNFTTradableRaw struct {
	Contract *FantomNFTTradable // Generic contract binding to access the raw methods on
}

// FantomNFTTradableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomNFTTradableCallerRaw struct {
	Contract *FantomNFTTradableCaller // Generic read-only contract binding to access the raw methods on
}

// FantomNFTTradableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomNFTTradableTransactorRaw struct {
	Contract *FantomNFTTradableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomNFTTradable creates a new instance of FantomNFTTradable, bound to a specific deployed contract.
func NewFantomNFTTradable(address common.Address, backend bind.ContractBackend) (*FantomNFTTradable, error) {
	contract, err := bindFantomNFTTradable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradable{FantomNFTTradableCaller: FantomNFTTradableCaller{contract: contract}, FantomNFTTradableTransactor: FantomNFTTradableTransactor{contract: contract}, FantomNFTTradableFilterer: FantomNFTTradableFilterer{contract: contract}}, nil
}

// NewFantomNFTTradableCaller creates a new read-only instance of FantomNFTTradable, bound to a specific deployed contract.
func NewFantomNFTTradableCaller(address common.Address, caller bind.ContractCaller) (*FantomNFTTradableCaller, error) {
	contract, err := bindFantomNFTTradable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableCaller{contract: contract}, nil
}

// NewFantomNFTTradableTransactor creates a new write-only instance of FantomNFTTradable, bound to a specific deployed contract.
func NewFantomNFTTradableTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomNFTTradableTransactor, error) {
	contract, err := bindFantomNFTTradable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableTransactor{contract: contract}, nil
}

// NewFantomNFTTradableFilterer creates a new log filterer instance of FantomNFTTradable, bound to a specific deployed contract.
func NewFantomNFTTradableFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomNFTTradableFilterer, error) {
	contract, err := bindFantomNFTTradable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableFilterer{contract: contract}, nil
}

// bindFantomNFTTradable binds a generic wrapper to an already deployed contract.
func bindFantomNFTTradable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomNFTTradableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomNFTTradable *FantomNFTTradableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomNFTTradable.Contract.FantomNFTTradableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomNFTTradable *FantomNFTTradableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.FantomNFTTradableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomNFTTradable *FantomNFTTradableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.FantomNFTTradableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomNFTTradable *FantomNFTTradableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomNFTTradable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomNFTTradable *FantomNFTTradableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomNFTTradable *FantomNFTTradableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FantomNFTTradable.Contract.BalanceOf(&_FantomNFTTradable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FantomNFTTradable.Contract.BalanceOf(&_FantomNFTTradable.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableSession) BaseURI() (string, error) {
	return _FantomNFTTradable.Contract.BaseURI(&_FantomNFTTradable.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) BaseURI() (string, error) {
	return _FantomNFTTradable.Contract.BaseURI(&_FantomNFTTradable.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCaller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomNFTTradable *FantomNFTTradableSession) FeeReceipient() (common.Address, error) {
	return _FantomNFTTradable.Contract.FeeReceipient(&_FantomNFTTradable.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) FeeReceipient() (common.Address, error) {
	return _FantomNFTTradable.Contract.FeeReceipient(&_FantomNFTTradable.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FantomNFTTradable *FantomNFTTradableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FantomNFTTradable.Contract.GetApproved(&_FantomNFTTradable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FantomNFTTradable.Contract.GetApproved(&_FantomNFTTradable.CallOpts, tokenId)
}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableCaller) IsApproved(opts *bind.CallOpts, _tokenId *big.Int, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "isApproved", _tokenId, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableSession) IsApproved(_tokenId *big.Int, _operator common.Address) (bool, error) {
	return _FantomNFTTradable.Contract.IsApproved(&_FantomNFTTradable.CallOpts, _tokenId, _operator)
}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) IsApproved(_tokenId *big.Int, _operator common.Address) (bool, error) {
	return _FantomNFTTradable.Contract.IsApproved(&_FantomNFTTradable.CallOpts, _tokenId, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FantomNFTTradable.Contract.IsApprovedForAll(&_FantomNFTTradable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FantomNFTTradable.Contract.IsApprovedForAll(&_FantomNFTTradable.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableSession) Name() (string, error) {
	return _FantomNFTTradable.Contract.Name(&_FantomNFTTradable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) Name() (string, error) {
	return _FantomNFTTradable.Contract.Name(&_FantomNFTTradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomNFTTradable *FantomNFTTradableSession) Owner() (common.Address, error) {
	return _FantomNFTTradable.Contract.Owner(&_FantomNFTTradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) Owner() (common.Address, error) {
	return _FantomNFTTradable.Contract.Owner(&_FantomNFTTradable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FantomNFTTradable *FantomNFTTradableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FantomNFTTradable.Contract.OwnerOf(&_FantomNFTTradable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FantomNFTTradable.Contract.OwnerOf(&_FantomNFTTradable.CallOpts, tokenId)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableSession) PlatformFee() (*big.Int, error) {
	return _FantomNFTTradable.Contract.PlatformFee(&_FantomNFTTradable.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) PlatformFee() (*big.Int, error) {
	return _FantomNFTTradable.Contract.PlatformFee(&_FantomNFTTradable.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FantomNFTTradable.Contract.SupportsInterface(&_FantomNFTTradable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FantomNFTTradable.Contract.SupportsInterface(&_FantomNFTTradable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableSession) Symbol() (string, error) {
	return _FantomNFTTradable.Contract.Symbol(&_FantomNFTTradable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) Symbol() (string, error) {
	return _FantomNFTTradable.Contract.Symbol(&_FantomNFTTradable.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _FantomNFTTradable.Contract.TokenByIndex(&_FantomNFTTradable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _FantomNFTTradable.Contract.TokenByIndex(&_FantomNFTTradable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FantomNFTTradable.Contract.TokenOfOwnerByIndex(&_FantomNFTTradable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FantomNFTTradable.Contract.TokenOfOwnerByIndex(&_FantomNFTTradable.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FantomNFTTradable *FantomNFTTradableSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FantomNFTTradable.Contract.TokenURI(&_FantomNFTTradable.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FantomNFTTradable.Contract.TokenURI(&_FantomNFTTradable.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomNFTTradable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableSession) TotalSupply() (*big.Int, error) {
	return _FantomNFTTradable.Contract.TotalSupply(&_FantomNFTTradable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FantomNFTTradable *FantomNFTTradableCallerSession) TotalSupply() (*big.Int, error) {
	return _FantomNFTTradable.Contract.TotalSupply(&_FantomNFTTradable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.Approve(&_FantomNFTTradable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.Approve(&_FantomNFTTradable.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "burn", _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.Burn(&_FantomNFTTradable.TransactOpts, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.Burn(&_FantomNFTTradable.TransactOpts, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _tokenUri) payable returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenUri string) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "mint", _to, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _tokenUri) payable returns()
func (_FantomNFTTradable *FantomNFTTradableSession) Mint(_to common.Address, _tokenUri string) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.Mint(&_FantomNFTTradable.TransactOpts, _to, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _tokenUri) payable returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) Mint(_to common.Address, _tokenUri string) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.Mint(&_FantomNFTTradable.TransactOpts, _to, _tokenUri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomNFTTradable *FantomNFTTradableSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.RenounceOwnership(&_FantomNFTTradable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.RenounceOwnership(&_FantomNFTTradable.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.SafeTransferFrom(&_FantomNFTTradable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.SafeTransferFrom(&_FantomNFTTradable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.SafeTransferFrom0(&_FantomNFTTradable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.SafeTransferFrom0(&_FantomNFTTradable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.SetApprovalForAll(&_FantomNFTTradable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.SetApprovalForAll(&_FantomNFTTradable.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.TransferFrom(&_FantomNFTTradable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.TransferFrom(&_FantomNFTTradable.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.TransferOwnership(&_FantomNFTTradable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.TransferOwnership(&_FantomNFTTradable.TransactOpts, newOwner)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeReceipient) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeReceipient common.Address) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "updateFeeRecipient", _feeReceipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeReceipient) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) UpdateFeeRecipient(_feeReceipient common.Address) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.UpdateFeeRecipient(&_FantomNFTTradable.TransactOpts, _feeReceipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeReceipient) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) UpdateFeeRecipient(_feeReceipient common.Address) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.UpdateFeeRecipient(&_FantomNFTTradable.TransactOpts, _feeReceipient)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomNFTTradable *FantomNFTTradableSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.UpdatePlatformFee(&_FantomNFTTradable.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomNFTTradable *FantomNFTTradableTransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomNFTTradable.Contract.UpdatePlatformFee(&_FantomNFTTradable.TransactOpts, _platformFee)
}

// FantomNFTTradableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FantomNFTTradable contract.
type FantomNFTTradableApprovalIterator struct {
	Event *FantomNFTTradableApproval // Event containing the contract specifics and raw log

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
func (it *FantomNFTTradableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTTradableApproval)
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
		it.Event = new(FantomNFTTradableApproval)
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
func (it *FantomNFTTradableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTTradableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTTradableApproval represents a Approval event raised by the FantomNFTTradable contract.
type FantomNFTTradableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FantomNFTTradable *FantomNFTTradableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FantomNFTTradableApprovalIterator, error) {

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

	logs, sub, err := _FantomNFTTradable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableApprovalIterator{contract: _FantomNFTTradable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FantomNFTTradable *FantomNFTTradableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FantomNFTTradableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FantomNFTTradable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTTradableApproval)
				if err := _FantomNFTTradable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_FantomNFTTradable *FantomNFTTradableFilterer) ParseApproval(log types.Log) (*FantomNFTTradableApproval, error) {
	event := new(FantomNFTTradableApproval)
	if err := _FantomNFTTradable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomNFTTradableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FantomNFTTradable contract.
type FantomNFTTradableApprovalForAllIterator struct {
	Event *FantomNFTTradableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FantomNFTTradableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTTradableApprovalForAll)
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
		it.Event = new(FantomNFTTradableApprovalForAll)
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
func (it *FantomNFTTradableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTTradableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTTradableApprovalForAll represents a ApprovalForAll event raised by the FantomNFTTradable contract.
type FantomNFTTradableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FantomNFTTradable *FantomNFTTradableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FantomNFTTradableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FantomNFTTradable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableApprovalForAllIterator{contract: _FantomNFTTradable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FantomNFTTradable *FantomNFTTradableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FantomNFTTradableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FantomNFTTradable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTTradableApprovalForAll)
				if err := _FantomNFTTradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_FantomNFTTradable *FantomNFTTradableFilterer) ParseApprovalForAll(log types.Log) (*FantomNFTTradableApprovalForAll, error) {
	event := new(FantomNFTTradableApprovalForAll)
	if err := _FantomNFTTradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomNFTTradableMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the FantomNFTTradable contract.
type FantomNFTTradableMintedIterator struct {
	Event *FantomNFTTradableMinted // Event containing the contract specifics and raw log

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
func (it *FantomNFTTradableMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTTradableMinted)
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
		it.Event = new(FantomNFTTradableMinted)
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
func (it *FantomNFTTradableMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTTradableMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTTradableMinted represents a Minted event raised by the FantomNFTTradable contract.
type FantomNFTTradableMinted struct {
	TokenId     *big.Int
	Beneficiary common.Address
	TokenUri    string
	Minter      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_FantomNFTTradable *FantomNFTTradableFilterer) FilterMinted(opts *bind.FilterOpts) (*FantomNFTTradableMintedIterator, error) {

	logs, sub, err := _FantomNFTTradable.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableMintedIterator{contract: _FantomNFTTradable.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_FantomNFTTradable *FantomNFTTradableFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *FantomNFTTradableMinted) (event.Subscription, error) {

	logs, sub, err := _FantomNFTTradable.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTTradableMinted)
				if err := _FantomNFTTradable.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_FantomNFTTradable *FantomNFTTradableFilterer) ParseMinted(log types.Log) (*FantomNFTTradableMinted, error) {
	event := new(FantomNFTTradableMinted)
	if err := _FantomNFTTradable.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomNFTTradableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomNFTTradable contract.
type FantomNFTTradableOwnershipTransferredIterator struct {
	Event *FantomNFTTradableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomNFTTradableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTTradableOwnershipTransferred)
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
		it.Event = new(FantomNFTTradableOwnershipTransferred)
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
func (it *FantomNFTTradableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTTradableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTTradableOwnershipTransferred represents a OwnershipTransferred event raised by the FantomNFTTradable contract.
type FantomNFTTradableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomNFTTradable *FantomNFTTradableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomNFTTradableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomNFTTradable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableOwnershipTransferredIterator{contract: _FantomNFTTradable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomNFTTradable *FantomNFTTradableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomNFTTradableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomNFTTradable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTTradableOwnershipTransferred)
				if err := _FantomNFTTradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomNFTTradable *FantomNFTTradableFilterer) ParseOwnershipTransferred(log types.Log) (*FantomNFTTradableOwnershipTransferred, error) {
	event := new(FantomNFTTradableOwnershipTransferred)
	if err := _FantomNFTTradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomNFTTradableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FantomNFTTradable contract.
type FantomNFTTradableTransferIterator struct {
	Event *FantomNFTTradableTransfer // Event containing the contract specifics and raw log

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
func (it *FantomNFTTradableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTTradableTransfer)
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
		it.Event = new(FantomNFTTradableTransfer)
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
func (it *FantomNFTTradableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTTradableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTTradableTransfer represents a Transfer event raised by the FantomNFTTradable contract.
type FantomNFTTradableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FantomNFTTradable *FantomNFTTradableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FantomNFTTradableTransferIterator, error) {

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

	logs, sub, err := _FantomNFTTradable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableTransferIterator{contract: _FantomNFTTradable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FantomNFTTradable *FantomNFTTradableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FantomNFTTradableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FantomNFTTradable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTTradableTransfer)
				if err := _FantomNFTTradable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FantomNFTTradable *FantomNFTTradableFilterer) ParseTransfer(log types.Log) (*FantomNFTTradableTransfer, error) {
	event := new(FantomNFTTradableTransfer)
	if err := _FantomNFTTradable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomNFTTradableUpdateFeeRecipientIterator is returned from FilterUpdateFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdateFeeRecipient events raised by the FantomNFTTradable contract.
type FantomNFTTradableUpdateFeeRecipientIterator struct {
	Event *FantomNFTTradableUpdateFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomNFTTradableUpdateFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTTradableUpdateFeeRecipient)
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
		it.Event = new(FantomNFTTradableUpdateFeeRecipient)
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
func (it *FantomNFTTradableUpdateFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTTradableUpdateFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTTradableUpdateFeeRecipient represents a UpdateFeeRecipient event raised by the FantomNFTTradable contract.
type FantomNFTTradableUpdateFeeRecipient struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeRecipient is a free log retrieval operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_FantomNFTTradable *FantomNFTTradableFilterer) FilterUpdateFeeRecipient(opts *bind.FilterOpts) (*FantomNFTTradableUpdateFeeRecipientIterator, error) {

	logs, sub, err := _FantomNFTTradable.contract.FilterLogs(opts, "UpdateFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableUpdateFeeRecipientIterator{contract: _FantomNFTTradable.contract, event: "UpdateFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeRecipient is a free log subscription operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_FantomNFTTradable *FantomNFTTradableFilterer) WatchUpdateFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomNFTTradableUpdateFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomNFTTradable.contract.WatchLogs(opts, "UpdateFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTTradableUpdateFeeRecipient)
				if err := _FantomNFTTradable.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
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

// ParseUpdateFeeRecipient is a log parse operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_FantomNFTTradable *FantomNFTTradableFilterer) ParseUpdateFeeRecipient(log types.Log) (*FantomNFTTradableUpdateFeeRecipient, error) {
	event := new(FantomNFTTradableUpdateFeeRecipient)
	if err := _FantomNFTTradable.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomNFTTradableUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomNFTTradable contract.
type FantomNFTTradableUpdatePlatformFeeIterator struct {
	Event *FantomNFTTradableUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomNFTTradableUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomNFTTradableUpdatePlatformFee)
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
		it.Event = new(FantomNFTTradableUpdatePlatformFee)
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
func (it *FantomNFTTradableUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomNFTTradableUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomNFTTradableUpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomNFTTradable contract.
type FantomNFTTradableUpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomNFTTradable *FantomNFTTradableFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomNFTTradableUpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomNFTTradable.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomNFTTradableUpdatePlatformFeeIterator{contract: _FantomNFTTradable.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomNFTTradable *FantomNFTTradableFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomNFTTradableUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomNFTTradable.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomNFTTradableUpdatePlatformFee)
				if err := _FantomNFTTradable.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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

// ParseUpdatePlatformFee is a log parse operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomNFTTradable *FantomNFTTradableFilterer) ParseUpdatePlatformFee(log types.Log) (*FantomNFTTradableUpdatePlatformFee, error) {
	event := new(FantomNFTTradableUpdatePlatformFee)
	if err := _FantomNFTTradable.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
