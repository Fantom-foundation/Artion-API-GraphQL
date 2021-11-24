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

// FantomMarketplaceV2MetaData contains all meta data concerning the FantomMarketplaceV2 contract.
var FantomMarketplaceV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ItemCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"name\":\"ItemListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unitPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ItemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OfferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"platformFee\",\"type\":\"uint16\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collectionRoyalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_platformFee\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTime\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"registerCollectionRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"}],\"name\":\"registerRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_platformFee\",\"type\":\"uint16\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"validateItemSold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomMarketplaceV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomMarketplaceV2MetaData.ABI instead.
var FantomMarketplaceV2ABI = FantomMarketplaceV2MetaData.ABI

// FantomMarketplaceV2 is an auto generated Go binding around an Ethereum contract.
type FantomMarketplaceV2 struct {
	FantomMarketplaceV2Caller     // Read-only binding to the contract
	FantomMarketplaceV2Transactor // Write-only binding to the contract
	FantomMarketplaceV2Filterer   // Log filterer for contract events
}

// FantomMarketplaceV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FantomMarketplaceV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomMarketplaceV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomMarketplaceV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomMarketplaceV2Session struct {
	Contract     *FantomMarketplaceV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FantomMarketplaceV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomMarketplaceV2CallerSession struct {
	Contract *FantomMarketplaceV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FantomMarketplaceV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomMarketplaceV2TransactorSession struct {
	Contract     *FantomMarketplaceV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FantomMarketplaceV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FantomMarketplaceV2Raw struct {
	Contract *FantomMarketplaceV2 // Generic contract binding to access the raw methods on
}

// FantomMarketplaceV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomMarketplaceV2CallerRaw struct {
	Contract *FantomMarketplaceV2Caller // Generic read-only contract binding to access the raw methods on
}

// FantomMarketplaceV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomMarketplaceV2TransactorRaw struct {
	Contract *FantomMarketplaceV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomMarketplaceV2 creates a new instance of FantomMarketplaceV2, bound to a specific deployed contract.
func NewFantomMarketplaceV2(address common.Address, backend bind.ContractBackend) (*FantomMarketplaceV2, error) {
	contract, err := bindFantomMarketplaceV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2{FantomMarketplaceV2Caller: FantomMarketplaceV2Caller{contract: contract}, FantomMarketplaceV2Transactor: FantomMarketplaceV2Transactor{contract: contract}, FantomMarketplaceV2Filterer: FantomMarketplaceV2Filterer{contract: contract}}, nil
}

// NewFantomMarketplaceV2Caller creates a new read-only instance of FantomMarketplaceV2, bound to a specific deployed contract.
func NewFantomMarketplaceV2Caller(address common.Address, caller bind.ContractCaller) (*FantomMarketplaceV2Caller, error) {
	contract, err := bindFantomMarketplaceV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2Caller{contract: contract}, nil
}

// NewFantomMarketplaceV2Transactor creates a new write-only instance of FantomMarketplaceV2, bound to a specific deployed contract.
func NewFantomMarketplaceV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FantomMarketplaceV2Transactor, error) {
	contract, err := bindFantomMarketplaceV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2Transactor{contract: contract}, nil
}

// NewFantomMarketplaceV2Filterer creates a new log filterer instance of FantomMarketplaceV2, bound to a specific deployed contract.
func NewFantomMarketplaceV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FantomMarketplaceV2Filterer, error) {
	contract, err := bindFantomMarketplaceV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2Filterer{contract: contract}, nil
}

// bindFantomMarketplaceV2 binds a generic wrapper to an already deployed contract.
func bindFantomMarketplaceV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomMarketplaceV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomMarketplaceV2 *FantomMarketplaceV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomMarketplaceV2.Contract.FantomMarketplaceV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomMarketplaceV2 *FantomMarketplaceV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.FantomMarketplaceV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomMarketplaceV2 *FantomMarketplaceV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.FantomMarketplaceV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomMarketplaceV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) AddressRegistry() (common.Address, error) {
	return _FantomMarketplaceV2.Contract.AddressRegistry(&_FantomMarketplaceV2.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) AddressRegistry() (common.Address, error) {
	return _FantomMarketplaceV2.Contract.AddressRegistry(&_FantomMarketplaceV2.CallOpts)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) CollectionRoyalties(opts *bind.CallOpts, arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "collectionRoyalties", arg0)

	outstruct := new(struct {
		Royalty      uint16
		Creator      common.Address
		FeeRecipient common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Royalty = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _FantomMarketplaceV2.Contract.CollectionRoyalties(&_FantomMarketplaceV2.CallOpts, arg0)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _FantomMarketplaceV2.Contract.CollectionRoyalties(&_FantomMarketplaceV2.CallOpts, arg0)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) FeeReceipient() (common.Address, error) {
	return _FantomMarketplaceV2.Contract.FeeReceipient(&_FantomMarketplaceV2.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) FeeReceipient() (common.Address, error) {
	return _FantomMarketplaceV2.Contract.FeeReceipient(&_FantomMarketplaceV2.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) GetPrice(opts *bind.CallOpts, _payToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "getPrice", _payToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _FantomMarketplaceV2.Contract.GetPrice(&_FantomMarketplaceV2.CallOpts, _payToken)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _FantomMarketplaceV2.Contract.GetPrice(&_FantomMarketplaceV2.CallOpts, _payToken)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) Listings(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "listings", arg0, arg1, arg2)

	outstruct := new(struct {
		Quantity     *big.Int
		PayToken     common.Address
		PricePerItem *big.Int
		StartingTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quantity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PricePerItem = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomMarketplaceV2.Contract.Listings(&_FantomMarketplaceV2.CallOpts, arg0, arg1, arg2)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomMarketplaceV2.Contract.Listings(&_FantomMarketplaceV2.CallOpts, arg0, arg1, arg2)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) Minters(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "minters", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FantomMarketplaceV2.Contract.Minters(&_FantomMarketplaceV2.CallOpts, arg0, arg1)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FantomMarketplaceV2.Contract.Minters(&_FantomMarketplaceV2.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) Offers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "offers", arg0, arg1, arg2)

	outstruct := new(struct {
		PayToken     common.Address
		Quantity     *big.Int
		PricePerItem *big.Int
		Deadline     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quantity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PricePerItem = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _FantomMarketplaceV2.Contract.Offers(&_FantomMarketplaceV2.CallOpts, arg0, arg1, arg2)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _FantomMarketplaceV2.Contract.Offers(&_FantomMarketplaceV2.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) Owner() (common.Address, error) {
	return _FantomMarketplaceV2.Contract.Owner(&_FantomMarketplaceV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) Owner() (common.Address, error) {
	return _FantomMarketplaceV2.Contract.Owner(&_FantomMarketplaceV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) PlatformFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) PlatformFee() (uint16, error) {
	return _FantomMarketplaceV2.Contract.PlatformFee(&_FantomMarketplaceV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) PlatformFee() (uint16, error) {
	return _FantomMarketplaceV2.Contract.PlatformFee(&_FantomMarketplaceV2.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Caller) Royalties(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint16, error) {
	var out []interface{}
	err := _FantomMarketplaceV2.contract.Call(opts, &out, "royalties", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _FantomMarketplaceV2.Contract.Royalties(&_FantomMarketplaceV2.CallOpts, arg0, arg1)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplaceV2 *FantomMarketplaceV2CallerSession) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _FantomMarketplaceV2.Contract.Royalties(&_FantomMarketplaceV2.CallOpts, arg0, arg1)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) AcceptOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "acceptOffer", _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.AcceptOffer(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.AcceptOffer(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _creator)
}

// BuyItem is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) BuyItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "buyItem", _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.BuyItem(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.BuyItem(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) CancelListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "cancelListing", _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.CancelListing(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.CancelListing(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) CancelOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "cancelOffer", _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.CancelOffer(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.CancelOffer(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) CreateOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "createOffer", _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.CreateOffer(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.CreateOffer(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xf39e69a4.
//
// Solidity: function initialize(address _feeRecipient, uint16 _platformFee) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) Initialize(opts *bind.TransactOpts, _feeRecipient common.Address, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "initialize", _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf39e69a4.
//
// Solidity: function initialize(address _feeRecipient, uint16 _platformFee) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) Initialize(_feeRecipient common.Address, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.Initialize(&_FantomMarketplaceV2.TransactOpts, _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf39e69a4.
//
// Solidity: function initialize(address _feeRecipient, uint16 _platformFee) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) Initialize(_feeRecipient common.Address, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.Initialize(&_FantomMarketplaceV2.TransactOpts, _feeRecipient, _platformFee)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) ListItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "listItem", _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.ListItem(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.ListItem(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) RegisterCollectionRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "registerCollectionRoyalty", _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.RegisterCollectionRoyalty(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.RegisterCollectionRoyalty(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) RegisterRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "registerRoyalty", _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.RegisterRoyalty(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.RegisterRoyalty(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.RenounceOwnership(&_FantomMarketplaceV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.RenounceOwnership(&_FantomMarketplaceV2.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.TransferOwnership(&_FantomMarketplaceV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.TransferOwnership(&_FantomMarketplaceV2.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdateAddressRegistry(&_FantomMarketplaceV2.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdateAddressRegistry(&_FantomMarketplaceV2.TransactOpts, _registry)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) UpdateListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "updateListing", _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdateListing(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdateListing(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdatePlatformFee(&_FantomMarketplaceV2.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdatePlatformFee(&_FantomMarketplaceV2.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdatePlatformFeeRecipient(&_FantomMarketplaceV2.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.UpdatePlatformFeeRecipient(&_FantomMarketplaceV2.TransactOpts, _platformFeeRecipient)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Transactor) ValidateItemSold(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.contract.Transact(opts, "validateItemSold", _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2Session) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.ValidateItemSold(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplaceV2 *FantomMarketplaceV2TransactorSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplaceV2.Contract.ValidateItemSold(&_FantomMarketplaceV2.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// FantomMarketplaceV2ItemCanceledIterator is returned from FilterItemCanceled and is used to iterate over the raw logs and unpacked data for ItemCanceled events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemCanceledIterator struct {
	Event *FantomMarketplaceV2ItemCanceled // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2ItemCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2ItemCanceled)
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
		it.Event = new(FantomMarketplaceV2ItemCanceled)
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
func (it *FantomMarketplaceV2ItemCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2ItemCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2ItemCanceled represents a ItemCanceled event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemCanceled struct {
	Owner   common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterItemCanceled is a free log retrieval operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterItemCanceled(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceV2ItemCanceledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2ItemCanceledIterator{contract: _FantomMarketplaceV2.contract, event: "ItemCanceled", logs: logs, sub: sub}, nil
}

// WatchItemCanceled is a free log subscription operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchItemCanceled(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2ItemCanceled, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2ItemCanceled)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
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

// ParseItemCanceled is a log parse operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseItemCanceled(log types.Log) (*FantomMarketplaceV2ItemCanceled, error) {
	event := new(FantomMarketplaceV2ItemCanceled)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2ItemListedIterator is returned from FilterItemListed and is used to iterate over the raw logs and unpacked data for ItemListed events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemListedIterator struct {
	Event *FantomMarketplaceV2ItemListed // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2ItemListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2ItemListed)
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
		it.Event = new(FantomMarketplaceV2ItemListed)
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
func (it *FantomMarketplaceV2ItemListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2ItemListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2ItemListed represents a ItemListed event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemListed struct {
	Owner        common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemListed is a free log retrieval operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterItemListed(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceV2ItemListedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2ItemListedIterator{contract: _FantomMarketplaceV2.contract, event: "ItemListed", logs: logs, sub: sub}, nil
}

// WatchItemListed is a free log subscription operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchItemListed(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2ItemListed, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2ItemListed)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemListed", log); err != nil {
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

// ParseItemListed is a log parse operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseItemListed(log types.Log) (*FantomMarketplaceV2ItemListed, error) {
	event := new(FantomMarketplaceV2ItemListed)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2ItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemSoldIterator struct {
	Event *FantomMarketplaceV2ItemSold // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2ItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2ItemSold)
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
		it.Event = new(FantomMarketplaceV2ItemSold)
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
func (it *FantomMarketplaceV2ItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2ItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2ItemSold represents a ItemSold event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemSold struct {
	Seller       common.Address
	Buyer        common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	UnitPrice    *big.Int
	PricePerItem *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterItemSold(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, nft []common.Address) (*FantomMarketplaceV2ItemSoldIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2ItemSoldIterator{contract: _FantomMarketplaceV2.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2ItemSold, seller []common.Address, buyer []common.Address, nft []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2ItemSold)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseItemSold(log types.Log) (*FantomMarketplaceV2ItemSold, error) {
	event := new(FantomMarketplaceV2ItemSold)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2ItemUpdatedIterator is returned from FilterItemUpdated and is used to iterate over the raw logs and unpacked data for ItemUpdated events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemUpdatedIterator struct {
	Event *FantomMarketplaceV2ItemUpdated // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2ItemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2ItemUpdated)
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
		it.Event = new(FantomMarketplaceV2ItemUpdated)
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
func (it *FantomMarketplaceV2ItemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2ItemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2ItemUpdated represents a ItemUpdated event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2ItemUpdated struct {
	Owner    common.Address
	Nft      common.Address
	TokenId  *big.Int
	PayToken common.Address
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemUpdated is a free log retrieval operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterItemUpdated(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceV2ItemUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2ItemUpdatedIterator{contract: _FantomMarketplaceV2.contract, event: "ItemUpdated", logs: logs, sub: sub}, nil
}

// WatchItemUpdated is a free log subscription operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchItemUpdated(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2ItemUpdated, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2ItemUpdated)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
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

// ParseItemUpdated is a log parse operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseItemUpdated(log types.Log) (*FantomMarketplaceV2ItemUpdated, error) {
	event := new(FantomMarketplaceV2ItemUpdated)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2OfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2OfferCanceledIterator struct {
	Event *FantomMarketplaceV2OfferCanceled // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2OfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2OfferCanceled)
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
		it.Event = new(FantomMarketplaceV2OfferCanceled)
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
func (it *FantomMarketplaceV2OfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2OfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2OfferCanceled represents a OfferCanceled event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2OfferCanceled struct {
	Creator common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterOfferCanceled(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*FantomMarketplaceV2OfferCanceledIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2OfferCanceledIterator{contract: _FantomMarketplaceV2.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2OfferCanceled, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2OfferCanceled)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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

// ParseOfferCanceled is a log parse operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseOfferCanceled(log types.Log) (*FantomMarketplaceV2OfferCanceled, error) {
	event := new(FantomMarketplaceV2OfferCanceled)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2OfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2OfferCreatedIterator struct {
	Event *FantomMarketplaceV2OfferCreated // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2OfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2OfferCreated)
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
		it.Event = new(FantomMarketplaceV2OfferCreated)
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
func (it *FantomMarketplaceV2OfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2OfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2OfferCreated represents a OfferCreated event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2OfferCreated struct {
	Creator      common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	Deadline     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterOfferCreated(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*FantomMarketplaceV2OfferCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2OfferCreatedIterator{contract: _FantomMarketplaceV2.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2OfferCreated, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2OfferCreated)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseOfferCreated(log types.Log) (*FantomMarketplaceV2OfferCreated, error) {
	event := new(FantomMarketplaceV2OfferCreated)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2OwnershipTransferredIterator struct {
	Event *FantomMarketplaceV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2OwnershipTransferred)
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
		it.Event = new(FantomMarketplaceV2OwnershipTransferred)
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
func (it *FantomMarketplaceV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2OwnershipTransferred represents a OwnershipTransferred event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomMarketplaceV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2OwnershipTransferredIterator{contract: _FantomMarketplaceV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2OwnershipTransferred)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseOwnershipTransferred(log types.Log) (*FantomMarketplaceV2OwnershipTransferred, error) {
	event := new(FantomMarketplaceV2OwnershipTransferred)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2UpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2UpdatePlatformFeeIterator struct {
	Event *FantomMarketplaceV2UpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2UpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2UpdatePlatformFee)
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
		it.Event = new(FantomMarketplaceV2UpdatePlatformFee)
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
func (it *FantomMarketplaceV2UpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2UpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2UpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2UpdatePlatformFee struct {
	PlatformFee uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomMarketplaceV2UpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2UpdatePlatformFeeIterator{contract: _FantomMarketplaceV2.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2UpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2UpdatePlatformFee)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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

// ParseUpdatePlatformFee is a log parse operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseUpdatePlatformFee(log types.Log) (*FantomMarketplaceV2UpdatePlatformFee, error) {
	event := new(FantomMarketplaceV2UpdatePlatformFee)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceV2UpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2UpdatePlatformFeeRecipientIterator struct {
	Event *FantomMarketplaceV2UpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceV2UpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceV2UpdatePlatformFeeRecipient)
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
		it.Event = new(FantomMarketplaceV2UpdatePlatformFeeRecipient)
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
func (it *FantomMarketplaceV2UpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceV2UpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceV2UpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the FantomMarketplaceV2 contract.
type FantomMarketplaceV2UpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*FantomMarketplaceV2UpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _FantomMarketplaceV2.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceV2UpdatePlatformFeeRecipientIterator{contract: _FantomMarketplaceV2.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceV2UpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomMarketplaceV2.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceV2UpdatePlatformFeeRecipient)
				if err := _FantomMarketplaceV2.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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

// ParseUpdatePlatformFeeRecipient is a log parse operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplaceV2 *FantomMarketplaceV2Filterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*FantomMarketplaceV2UpdatePlatformFeeRecipient, error) {
	event := new(FantomMarketplaceV2UpdatePlatformFeeRecipient)
	if err := _FantomMarketplaceV2.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
