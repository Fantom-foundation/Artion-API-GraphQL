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

// FantomArtTradableMetaData contains all meta data concerning the FantomArtTradable contract.
var FantomArtTradableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeReceipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"_exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokenID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOperator\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FantomArtTradableABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomArtTradableMetaData.ABI instead.
var FantomArtTradableABI = FantomArtTradableMetaData.ABI

// FantomArtTradable is an auto generated Go binding around an Ethereum contract.
type FantomArtTradable struct {
	FantomArtTradableCaller     // Read-only binding to the contract
	FantomArtTradableTransactor // Write-only binding to the contract
	FantomArtTradableFilterer   // Log filterer for contract events
}

// FantomArtTradableCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomArtTradableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomArtTradableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomArtTradableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomArtTradableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomArtTradableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomArtTradableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomArtTradableSession struct {
	Contract     *FantomArtTradable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FantomArtTradableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomArtTradableCallerSession struct {
	Contract *FantomArtTradableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FantomArtTradableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomArtTradableTransactorSession struct {
	Contract     *FantomArtTradableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FantomArtTradableRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomArtTradableRaw struct {
	Contract *FantomArtTradable // Generic contract binding to access the raw methods on
}

// FantomArtTradableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomArtTradableCallerRaw struct {
	Contract *FantomArtTradableCaller // Generic read-only contract binding to access the raw methods on
}

// FantomArtTradableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomArtTradableTransactorRaw struct {
	Contract *FantomArtTradableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomArtTradable creates a new instance of FantomArtTradable, bound to a specific deployed contract.
func NewFantomArtTradable(address common.Address, backend bind.ContractBackend) (*FantomArtTradable, error) {
	contract, err := bindFantomArtTradable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradable{FantomArtTradableCaller: FantomArtTradableCaller{contract: contract}, FantomArtTradableTransactor: FantomArtTradableTransactor{contract: contract}, FantomArtTradableFilterer: FantomArtTradableFilterer{contract: contract}}, nil
}

// NewFantomArtTradableCaller creates a new read-only instance of FantomArtTradable, bound to a specific deployed contract.
func NewFantomArtTradableCaller(address common.Address, caller bind.ContractCaller) (*FantomArtTradableCaller, error) {
	contract, err := bindFantomArtTradable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableCaller{contract: contract}, nil
}

// NewFantomArtTradableTransactor creates a new write-only instance of FantomArtTradable, bound to a specific deployed contract.
func NewFantomArtTradableTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomArtTradableTransactor, error) {
	contract, err := bindFantomArtTradable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableTransactor{contract: contract}, nil
}

// NewFantomArtTradableFilterer creates a new log filterer instance of FantomArtTradable, bound to a specific deployed contract.
func NewFantomArtTradableFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomArtTradableFilterer, error) {
	contract, err := bindFantomArtTradable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableFilterer{contract: contract}, nil
}

// bindFantomArtTradable binds a generic wrapper to an already deployed contract.
func bindFantomArtTradable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomArtTradableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomArtTradable *FantomArtTradableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomArtTradable.Contract.FantomArtTradableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomArtTradable *FantomArtTradableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.FantomArtTradableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomArtTradable *FantomArtTradableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.FantomArtTradableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomArtTradable *FantomArtTradableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomArtTradable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomArtTradable *FantomArtTradableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomArtTradable *FantomArtTradableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0xf8e76cc0.
//
// Solidity: function _exists(uint256 _id) view returns(bool)
func (_FantomArtTradable *FantomArtTradableCaller) Exists(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "_exists", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf8e76cc0.
//
// Solidity: function _exists(uint256 _id) view returns(bool)
func (_FantomArtTradable *FantomArtTradableSession) Exists(_id *big.Int) (bool, error) {
	return _FantomArtTradable.Contract.Exists(&_FantomArtTradable.CallOpts, _id)
}

// Exists is a free data retrieval call binding the contract method 0xf8e76cc0.
//
// Solidity: function _exists(uint256 _id) view returns(bool)
func (_FantomArtTradable *FantomArtTradableCallerSession) Exists(_id *big.Int) (bool, error) {
	return _FantomArtTradable.Contract.Exists(&_FantomArtTradable.CallOpts, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "balanceOf", _owner, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _FantomArtTradable.Contract.BalanceOf(&_FantomArtTradable.CallOpts, _owner, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCallerSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _FantomArtTradable.Contract.BalanceOf(&_FantomArtTradable.CallOpts, _owner, _id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_FantomArtTradable *FantomArtTradableCaller) BalanceOfBatch(opts *bind.CallOpts, _owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "balanceOfBatch", _owners, _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_FantomArtTradable *FantomArtTradableSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _FantomArtTradable.Contract.BalanceOfBatch(&_FantomArtTradable.CallOpts, _owners, _ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_FantomArtTradable *FantomArtTradableCallerSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _FantomArtTradable.Contract.BalanceOfBatch(&_FantomArtTradable.CallOpts, _owners, _ids)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_FantomArtTradable *FantomArtTradableCaller) Creators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "creators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_FantomArtTradable *FantomArtTradableSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _FantomArtTradable.Contract.Creators(&_FantomArtTradable.CallOpts, arg0)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_FantomArtTradable *FantomArtTradableCallerSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _FantomArtTradable.Contract.Creators(&_FantomArtTradable.CallOpts, arg0)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomArtTradable *FantomArtTradableCaller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomArtTradable *FantomArtTradableSession) FeeReceipient() (common.Address, error) {
	return _FantomArtTradable.Contract.FeeReceipient(&_FantomArtTradable.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomArtTradable *FantomArtTradableCallerSession) FeeReceipient() (common.Address, error) {
	return _FantomArtTradable.Contract.FeeReceipient(&_FantomArtTradable.CallOpts)
}

// GetCurrentTokenID is a free data retrieval call binding the contract method 0xc4a6d0d2.
//
// Solidity: function getCurrentTokenID() view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCaller) GetCurrentTokenID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "getCurrentTokenID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTokenID is a free data retrieval call binding the contract method 0xc4a6d0d2.
//
// Solidity: function getCurrentTokenID() view returns(uint256)
func (_FantomArtTradable *FantomArtTradableSession) GetCurrentTokenID() (*big.Int, error) {
	return _FantomArtTradable.Contract.GetCurrentTokenID(&_FantomArtTradable.CallOpts)
}

// GetCurrentTokenID is a free data retrieval call binding the contract method 0xc4a6d0d2.
//
// Solidity: function getCurrentTokenID() view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCallerSession) GetCurrentTokenID() (*big.Int, error) {
	return _FantomArtTradable.Contract.GetCurrentTokenID(&_FantomArtTradable.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_FantomArtTradable *FantomArtTradableCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_FantomArtTradable *FantomArtTradableSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _FantomArtTradable.Contract.IsApprovedForAll(&_FantomArtTradable.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_FantomArtTradable *FantomArtTradableCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _FantomArtTradable.Contract.IsApprovedForAll(&_FantomArtTradable.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FantomArtTradable *FantomArtTradableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FantomArtTradable *FantomArtTradableSession) Name() (string, error) {
	return _FantomArtTradable.Contract.Name(&_FantomArtTradable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FantomArtTradable *FantomArtTradableCallerSession) Name() (string, error) {
	return _FantomArtTradable.Contract.Name(&_FantomArtTradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomArtTradable *FantomArtTradableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomArtTradable *FantomArtTradableSession) Owner() (common.Address, error) {
	return _FantomArtTradable.Contract.Owner(&_FantomArtTradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomArtTradable *FantomArtTradableCallerSession) Owner() (common.Address, error) {
	return _FantomArtTradable.Contract.Owner(&_FantomArtTradable.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomArtTradable *FantomArtTradableSession) PlatformFee() (*big.Int, error) {
	return _FantomArtTradable.Contract.PlatformFee(&_FantomArtTradable.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCallerSession) PlatformFee() (*big.Int, error) {
	return _FantomArtTradable.Contract.PlatformFee(&_FantomArtTradable.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_FantomArtTradable *FantomArtTradableCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_FantomArtTradable *FantomArtTradableSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _FantomArtTradable.Contract.SupportsInterface(&_FantomArtTradable.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_FantomArtTradable *FantomArtTradableCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _FantomArtTradable.Contract.SupportsInterface(&_FantomArtTradable.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FantomArtTradable *FantomArtTradableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FantomArtTradable *FantomArtTradableSession) Symbol() (string, error) {
	return _FantomArtTradable.Contract.Symbol(&_FantomArtTradable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FantomArtTradable *FantomArtTradableCallerSession) Symbol() (string, error) {
	return _FantomArtTradable.Contract.Symbol(&_FantomArtTradable.CallOpts)
}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCaller) TokenSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "tokenSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableSession) TokenSupply(arg0 *big.Int) (*big.Int, error) {
	return _FantomArtTradable.Contract.TokenSupply(&_FantomArtTradable.CallOpts, arg0)
}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCallerSession) TokenSupply(arg0 *big.Int) (*big.Int, error) {
	return _FantomArtTradable.Contract.TokenSupply(&_FantomArtTradable.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCaller) TotalSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "totalSupply", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _FantomArtTradable.Contract.TotalSupply(&_FantomArtTradable.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_FantomArtTradable *FantomArtTradableCallerSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _FantomArtTradable.Contract.TotalSupply(&_FantomArtTradable.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_FantomArtTradable *FantomArtTradableCaller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _FantomArtTradable.contract.Call(opts, &out, "uri", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_FantomArtTradable *FantomArtTradableSession) Uri(_id *big.Int) (string, error) {
	return _FantomArtTradable.Contract.Uri(&_FantomArtTradable.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_FantomArtTradable *FantomArtTradableCallerSession) Uri(_id *big.Int) (string, error) {
	return _FantomArtTradable.Contract.Uri(&_FantomArtTradable.CallOpts, _id)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address _to, uint256 _supply, string _uri) payable returns()
func (_FantomArtTradable *FantomArtTradableTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _supply *big.Int, _uri string) (*types.Transaction, error) {
	return _FantomArtTradable.contract.Transact(opts, "mint", _to, _supply, _uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address _to, uint256 _supply, string _uri) payable returns()
func (_FantomArtTradable *FantomArtTradableSession) Mint(_to common.Address, _supply *big.Int, _uri string) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.Mint(&_FantomArtTradable.TransactOpts, _to, _supply, _uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address _to, uint256 _supply, string _uri) payable returns()
func (_FantomArtTradable *FantomArtTradableTransactorSession) Mint(_to common.Address, _supply *big.Int, _uri string) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.Mint(&_FantomArtTradable.TransactOpts, _to, _supply, _uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomArtTradable *FantomArtTradableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomArtTradable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomArtTradable *FantomArtTradableSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomArtTradable.Contract.RenounceOwnership(&_FantomArtTradable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomArtTradable *FantomArtTradableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomArtTradable.Contract.RenounceOwnership(&_FantomArtTradable.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_FantomArtTradable *FantomArtTradableTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomArtTradable.contract.Transact(opts, "safeBatchTransferFrom", _from, _to, _ids, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_FantomArtTradable *FantomArtTradableSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.SafeBatchTransferFrom(&_FantomArtTradable.TransactOpts, _from, _to, _ids, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_FantomArtTradable *FantomArtTradableTransactorSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.SafeBatchTransferFrom(&_FantomArtTradable.TransactOpts, _from, _to, _ids, _amounts, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_FantomArtTradable *FantomArtTradableTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomArtTradable.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_FantomArtTradable *FantomArtTradableSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.SafeTransferFrom(&_FantomArtTradable.TransactOpts, _from, _to, _id, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_FantomArtTradable *FantomArtTradableTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.SafeTransferFrom(&_FantomArtTradable.TransactOpts, _from, _to, _id, _amount, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_FantomArtTradable *FantomArtTradableTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _FantomArtTradable.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_FantomArtTradable *FantomArtTradableSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.SetApprovalForAll(&_FantomArtTradable.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_FantomArtTradable *FantomArtTradableTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.SetApprovalForAll(&_FantomArtTradable.TransactOpts, _operator, _approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomArtTradable *FantomArtTradableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomArtTradable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomArtTradable *FantomArtTradableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.TransferOwnership(&_FantomArtTradable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomArtTradable *FantomArtTradableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomArtTradable.Contract.TransferOwnership(&_FantomArtTradable.TransactOpts, newOwner)
}

// FantomArtTradableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FantomArtTradable contract.
type FantomArtTradableApprovalForAllIterator struct {
	Event *FantomArtTradableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FantomArtTradableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomArtTradableApprovalForAll)
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
		it.Event = new(FantomArtTradableApprovalForAll)
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
func (it *FantomArtTradableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomArtTradableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomArtTradableApprovalForAll represents a ApprovalForAll event raised by the FantomArtTradable contract.
type FantomArtTradableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_FantomArtTradable *FantomArtTradableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*FantomArtTradableApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _FantomArtTradable.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableApprovalForAllIterator{contract: _FantomArtTradable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_FantomArtTradable *FantomArtTradableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FantomArtTradableApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _FantomArtTradable.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomArtTradableApprovalForAll)
				if err := _FantomArtTradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_FantomArtTradable *FantomArtTradableFilterer) ParseApprovalForAll(log types.Log) (*FantomArtTradableApprovalForAll, error) {
	event := new(FantomArtTradableApprovalForAll)
	if err := _FantomArtTradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomArtTradableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomArtTradable contract.
type FantomArtTradableOwnershipTransferredIterator struct {
	Event *FantomArtTradableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomArtTradableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomArtTradableOwnershipTransferred)
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
		it.Event = new(FantomArtTradableOwnershipTransferred)
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
func (it *FantomArtTradableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomArtTradableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomArtTradableOwnershipTransferred represents a OwnershipTransferred event raised by the FantomArtTradable contract.
type FantomArtTradableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomArtTradable *FantomArtTradableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomArtTradableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomArtTradable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableOwnershipTransferredIterator{contract: _FantomArtTradable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomArtTradable *FantomArtTradableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomArtTradableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomArtTradable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomArtTradableOwnershipTransferred)
				if err := _FantomArtTradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomArtTradable *FantomArtTradableFilterer) ParseOwnershipTransferred(log types.Log) (*FantomArtTradableOwnershipTransferred, error) {
	event := new(FantomArtTradableOwnershipTransferred)
	if err := _FantomArtTradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomArtTradableTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the FantomArtTradable contract.
type FantomArtTradableTransferBatchIterator struct {
	Event *FantomArtTradableTransferBatch // Event containing the contract specifics and raw log

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
func (it *FantomArtTradableTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomArtTradableTransferBatch)
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
		it.Event = new(FantomArtTradableTransferBatch)
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
func (it *FantomArtTradableTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomArtTradableTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomArtTradableTransferBatch represents a TransferBatch event raised by the FantomArtTradable contract.
type FantomArtTradableTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_FantomArtTradable *FantomArtTradableFilterer) FilterTransferBatch(opts *bind.FilterOpts, _operator []common.Address, _from []common.Address, _to []common.Address) (*FantomArtTradableTransferBatchIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FantomArtTradable.contract.FilterLogs(opts, "TransferBatch", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableTransferBatchIterator{contract: _FantomArtTradable.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_FantomArtTradable *FantomArtTradableFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *FantomArtTradableTransferBatch, _operator []common.Address, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FantomArtTradable.contract.WatchLogs(opts, "TransferBatch", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomArtTradableTransferBatch)
				if err := _FantomArtTradable.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_FantomArtTradable *FantomArtTradableFilterer) ParseTransferBatch(log types.Log) (*FantomArtTradableTransferBatch, error) {
	event := new(FantomArtTradableTransferBatch)
	if err := _FantomArtTradable.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomArtTradableTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the FantomArtTradable contract.
type FantomArtTradableTransferSingleIterator struct {
	Event *FantomArtTradableTransferSingle // Event containing the contract specifics and raw log

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
func (it *FantomArtTradableTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomArtTradableTransferSingle)
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
		it.Event = new(FantomArtTradableTransferSingle)
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
func (it *FantomArtTradableTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomArtTradableTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomArtTradableTransferSingle represents a TransferSingle event raised by the FantomArtTradable contract.
type FantomArtTradableTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_FantomArtTradable *FantomArtTradableFilterer) FilterTransferSingle(opts *bind.FilterOpts, _operator []common.Address, _from []common.Address, _to []common.Address) (*FantomArtTradableTransferSingleIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FantomArtTradable.contract.FilterLogs(opts, "TransferSingle", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableTransferSingleIterator{contract: _FantomArtTradable.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_FantomArtTradable *FantomArtTradableFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *FantomArtTradableTransferSingle, _operator []common.Address, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FantomArtTradable.contract.WatchLogs(opts, "TransferSingle", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomArtTradableTransferSingle)
				if err := _FantomArtTradable.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_FantomArtTradable *FantomArtTradableFilterer) ParseTransferSingle(log types.Log) (*FantomArtTradableTransferSingle, error) {
	event := new(FantomArtTradableTransferSingle)
	if err := _FantomArtTradable.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomArtTradableURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the FantomArtTradable contract.
type FantomArtTradableURIIterator struct {
	Event *FantomArtTradableURI // Event containing the contract specifics and raw log

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
func (it *FantomArtTradableURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomArtTradableURI)
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
		it.Event = new(FantomArtTradableURI)
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
func (it *FantomArtTradableURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomArtTradableURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomArtTradableURI represents a URI event raised by the FantomArtTradable contract.
type FantomArtTradableURI struct {
	Uri string
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string _uri, uint256 indexed _id)
func (_FantomArtTradable *FantomArtTradableFilterer) FilterURI(opts *bind.FilterOpts, _id []*big.Int) (*FantomArtTradableURIIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _FantomArtTradable.contract.FilterLogs(opts, "URI", _idRule)
	if err != nil {
		return nil, err
	}
	return &FantomArtTradableURIIterator{contract: _FantomArtTradable.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string _uri, uint256 indexed _id)
func (_FantomArtTradable *FantomArtTradableFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *FantomArtTradableURI, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _FantomArtTradable.contract.WatchLogs(opts, "URI", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomArtTradableURI)
				if err := _FantomArtTradable.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string _uri, uint256 indexed _id)
func (_FantomArtTradable *FantomArtTradableFilterer) ParseURI(log types.Log) (*FantomArtTradableURI, error) {
	event := new(FantomArtTradableURI)
	if err := _FantomArtTradable.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
