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

// FantomMarketplaceMetaData contains all meta data concerning the FantomMarketplace contract.
var FantomMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ItemCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"name\":\"ItemListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unitPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ItemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OfferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"platformFee\",\"type\":\"uint16\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collectionRoyalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_platformFee\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTime\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"registerCollectionRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"}],\"name\":\"registerRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_platformFee\",\"type\":\"uint16\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"validateItemSold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomMarketplaceMetaData.ABI instead.
var FantomMarketplaceABI = FantomMarketplaceMetaData.ABI

// FantomMarketplace is an auto generated Go binding around an Ethereum contract.
type FantomMarketplace struct {
	FantomMarketplaceCaller     // Read-only binding to the contract
	FantomMarketplaceTransactor // Write-only binding to the contract
	FantomMarketplaceFilterer   // Log filterer for contract events
}

// FantomMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomMarketplaceSession struct {
	Contract     *FantomMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FantomMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomMarketplaceCallerSession struct {
	Contract *FantomMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FantomMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomMarketplaceTransactorSession struct {
	Contract     *FantomMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FantomMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomMarketplaceRaw struct {
	Contract *FantomMarketplace // Generic contract binding to access the raw methods on
}

// FantomMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomMarketplaceCallerRaw struct {
	Contract *FantomMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// FantomMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomMarketplaceTransactorRaw struct {
	Contract *FantomMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomMarketplace creates a new instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplace(address common.Address, backend bind.ContractBackend) (*FantomMarketplace, error) {
	contract, err := bindFantomMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplace{FantomMarketplaceCaller: FantomMarketplaceCaller{contract: contract}, FantomMarketplaceTransactor: FantomMarketplaceTransactor{contract: contract}, FantomMarketplaceFilterer: FantomMarketplaceFilterer{contract: contract}}, nil
}

// NewFantomMarketplaceCaller creates a new read-only instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*FantomMarketplaceCaller, error) {
	contract, err := bindFantomMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceCaller{contract: contract}, nil
}

// NewFantomMarketplaceTransactor creates a new write-only instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomMarketplaceTransactor, error) {
	contract, err := bindFantomMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceTransactor{contract: contract}, nil
}

// NewFantomMarketplaceFilterer creates a new log filterer instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomMarketplaceFilterer, error) {
	contract, err := bindFantomMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceFilterer{contract: contract}, nil
}

// bindFantomMarketplace binds a generic wrapper to an already deployed contract.
func bindFantomMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomMarketplace *FantomMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomMarketplace.Contract.FantomMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomMarketplace *FantomMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.FantomMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomMarketplace *FantomMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.FantomMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomMarketplace *FantomMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomMarketplace *FantomMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomMarketplace *FantomMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) AddressRegistry() (common.Address, error) {
	return _FantomMarketplace.Contract.AddressRegistry(&_FantomMarketplace.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) AddressRegistry() (common.Address, error) {
	return _FantomMarketplace.Contract.AddressRegistry(&_FantomMarketplace.CallOpts)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplace *FantomMarketplaceCaller) CollectionRoyalties(opts *bind.CallOpts, arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "collectionRoyalties", arg0)

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
func (_FantomMarketplace *FantomMarketplaceSession) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _FantomMarketplace.Contract.CollectionRoyalties(&_FantomMarketplace.CallOpts, arg0)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplace *FantomMarketplaceCallerSession) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _FantomMarketplace.Contract.CollectionRoyalties(&_FantomMarketplace.CallOpts, arg0)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) FeeReceipient() (common.Address, error) {
	return _FantomMarketplace.Contract.FeeReceipient(&_FantomMarketplace.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) FeeReceipient() (common.Address, error) {
	return _FantomMarketplace.Contract.FeeReceipient(&_FantomMarketplace.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplace *FantomMarketplaceCaller) GetPrice(opts *bind.CallOpts, _payToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "getPrice", _payToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplace *FantomMarketplaceSession) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _FantomMarketplace.Contract.GetPrice(&_FantomMarketplace.CallOpts, _payToken)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplace *FantomMarketplaceCallerSession) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _FantomMarketplace.Contract.GetPrice(&_FantomMarketplace.CallOpts, _payToken)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "listings", arg0, arg1, arg2)

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
func (_FantomMarketplace *FantomMarketplaceSession) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomMarketplace.Contract.Listings(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomMarketplace.Contract.Listings(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) Minters(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "minters", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FantomMarketplace.Contract.Minters(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FantomMarketplace.Contract.Minters(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceCaller) Offers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "offers", arg0, arg1, arg2)

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
func (_FantomMarketplace *FantomMarketplaceSession) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _FantomMarketplace.Contract.Offers(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _FantomMarketplace.Contract.Offers(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) Owner() (common.Address, error) {
	return _FantomMarketplace.Contract.Owner(&_FantomMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Owner() (common.Address, error) {
	return _FantomMarketplace.Contract.Owner(&_FantomMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCaller) PlatformFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceSession) PlatformFee() (uint16, error) {
	return _FantomMarketplace.Contract.PlatformFee(&_FantomMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCallerSession) PlatformFee() (uint16, error) {
	return _FantomMarketplace.Contract.PlatformFee(&_FantomMarketplace.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCaller) Royalties(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint16, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "royalties", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceSession) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _FantomMarketplace.Contract.Royalties(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _FantomMarketplace.Contract.Royalties(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) AcceptOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "acceptOffer", _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplace *FantomMarketplaceSession) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.AcceptOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.AcceptOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _creator)
}

// BuyItem is a paid mutator transaction binding the contract method 0x259ca365.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _owner) payable returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) BuyItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "buyItem", _nftAddress, _tokenId, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x259ca365.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _owner) payable returns()
func (_FantomMarketplace *FantomMarketplaceSession) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x259ca365.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _owner) payable returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _owner)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) BuyItem0(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "buyItem0", _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplace *FantomMarketplaceSession) BuyItem0(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem0(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) BuyItem0(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem0(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) CancelListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "cancelListing", _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceSession) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) CancelOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "cancelOffer", _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceSession) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) CreateOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "createOffer", _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplace *FantomMarketplaceSession) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CreateOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CreateOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xf39e69a4.
//
// Solidity: function initialize(address _feeRecipient, uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) Initialize(opts *bind.TransactOpts, _feeRecipient common.Address, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "initialize", _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf39e69a4.
//
// Solidity: function initialize(address _feeRecipient, uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceSession) Initialize(_feeRecipient common.Address, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.Initialize(&_FantomMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf39e69a4.
//
// Solidity: function initialize(address _feeRecipient, uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) Initialize(_feeRecipient common.Address, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.Initialize(&_FantomMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) ListItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "listItem", _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplace *FantomMarketplaceSession) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ListItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ListItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) RegisterCollectionRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "registerCollectionRoyalty", _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceSession) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterCollectionRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterCollectionRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) RegisterRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "registerRoyalty", _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplace *FantomMarketplaceSession) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplace *FantomMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RenounceOwnership(&_FantomMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RenounceOwnership(&_FantomMarketplace.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplace *FantomMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.TransferOwnership(&_FantomMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.TransferOwnership(&_FantomMarketplace.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateAddressRegistry(&_FantomMarketplace.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateAddressRegistry(&_FantomMarketplace.TransactOpts, _registry)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdateListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updateListing", _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFee(&_FantomMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFee(&_FantomMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFeeRecipient(&_FantomMarketplace.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFeeRecipient(&_FantomMarketplace.TransactOpts, _platformFeeRecipient)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) ValidateItemSold(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "validateItemSold", _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplace *FantomMarketplaceSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ValidateItemSold(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ValidateItemSold(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// FantomMarketplaceItemCanceledIterator is returned from FilterItemCanceled and is used to iterate over the raw logs and unpacked data for ItemCanceled events raised by the FantomMarketplace contract.
type FantomMarketplaceItemCanceledIterator struct {
	Event *FantomMarketplaceItemCanceled // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemCanceled)
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
		it.Event = new(FantomMarketplaceItemCanceled)
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
func (it *FantomMarketplaceItemCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemCanceled represents a ItemCanceled event raised by the FantomMarketplace contract.
type FantomMarketplaceItemCanceled struct {
	Owner   common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterItemCanceled is a free log retrieval operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemCanceled(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceItemCanceledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemCanceledIterator{contract: _FantomMarketplace.contract, event: "ItemCanceled", logs: logs, sub: sub}, nil
}

// WatchItemCanceled is a free log subscription operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemCanceled(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemCanceled, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemCanceled)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemCanceled(log types.Log) (*FantomMarketplaceItemCanceled, error) {
	event := new(FantomMarketplaceItemCanceled)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceItemListedIterator is returned from FilterItemListed and is used to iterate over the raw logs and unpacked data for ItemListed events raised by the FantomMarketplace contract.
type FantomMarketplaceItemListedIterator struct {
	Event *FantomMarketplaceItemListed // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemListed)
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
		it.Event = new(FantomMarketplaceItemListed)
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
func (it *FantomMarketplaceItemListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemListed represents a ItemListed event raised by the FantomMarketplace contract.
type FantomMarketplaceItemListed struct {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemListed(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceItemListedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemListedIterator{contract: _FantomMarketplace.contract, event: "ItemListed", logs: logs, sub: sub}, nil
}

// WatchItemListed is a free log subscription operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemListed(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemListed, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemListed)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemListed(log types.Log) (*FantomMarketplaceItemListed, error) {
	event := new(FantomMarketplaceItemListed)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the FantomMarketplace contract.
type FantomMarketplaceItemSoldIterator struct {
	Event *FantomMarketplaceItemSold // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemSold)
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
		it.Event = new(FantomMarketplaceItemSold)
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
func (it *FantomMarketplaceItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemSold represents a ItemSold event raised by the FantomMarketplace contract.
type FantomMarketplaceItemSold struct {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemSold(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, nft []common.Address) (*FantomMarketplaceItemSoldIterator, error) {

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

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemSoldIterator{contract: _FantomMarketplace.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemSold, seller []common.Address, buyer []common.Address, nft []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemSold)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemSold(log types.Log) (*FantomMarketplaceItemSold, error) {
	event := new(FantomMarketplaceItemSold)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceItemUpdatedIterator is returned from FilterItemUpdated and is used to iterate over the raw logs and unpacked data for ItemUpdated events raised by the FantomMarketplace contract.
type FantomMarketplaceItemUpdatedIterator struct {
	Event *FantomMarketplaceItemUpdated // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemUpdated)
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
		it.Event = new(FantomMarketplaceItemUpdated)
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
func (it *FantomMarketplaceItemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemUpdated represents a ItemUpdated event raised by the FantomMarketplace contract.
type FantomMarketplaceItemUpdated struct {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemUpdated(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceItemUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemUpdatedIterator{contract: _FantomMarketplace.contract, event: "ItemUpdated", logs: logs, sub: sub}, nil
}

// WatchItemUpdated is a free log subscription operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemUpdated(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemUpdated, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemUpdated)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemUpdated(log types.Log) (*FantomMarketplaceItemUpdated, error) {
	event := new(FantomMarketplaceItemUpdated)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCanceledIterator struct {
	Event *FantomMarketplaceOfferCanceled // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceOfferCanceled)
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
		it.Event = new(FantomMarketplaceOfferCanceled)
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
func (it *FantomMarketplaceOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceOfferCanceled represents a OfferCanceled event raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCanceled struct {
	Creator common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterOfferCanceled(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*FantomMarketplaceOfferCanceledIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceOfferCanceledIterator{contract: _FantomMarketplace.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceOfferCanceled, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceOfferCanceled)
				if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseOfferCanceled(log types.Log) (*FantomMarketplaceOfferCanceled, error) {
	event := new(FantomMarketplaceOfferCanceled)
	if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCreatedIterator struct {
	Event *FantomMarketplaceOfferCreated // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceOfferCreated)
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
		it.Event = new(FantomMarketplaceOfferCreated)
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
func (it *FantomMarketplaceOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceOfferCreated represents a OfferCreated event raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCreated struct {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterOfferCreated(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*FantomMarketplaceOfferCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceOfferCreatedIterator{contract: _FantomMarketplace.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceOfferCreated, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceOfferCreated)
				if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseOfferCreated(log types.Log) (*FantomMarketplaceOfferCreated, error) {
	event := new(FantomMarketplaceOfferCreated)
	if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomMarketplace contract.
type FantomMarketplaceOwnershipTransferredIterator struct {
	Event *FantomMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceOwnershipTransferred)
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
		it.Event = new(FantomMarketplaceOwnershipTransferred)
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
func (it *FantomMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the FantomMarketplace contract.
type FantomMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceOwnershipTransferredIterator{contract: _FantomMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceOwnershipTransferred)
				if err := _FantomMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*FantomMarketplaceOwnershipTransferred, error) {
	event := new(FantomMarketplaceOwnershipTransferred)
	if err := _FantomMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFeeIterator struct {
	Event *FantomMarketplaceUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceUpdatePlatformFee)
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
		it.Event = new(FantomMarketplaceUpdatePlatformFee)
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
func (it *FantomMarketplaceUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceUpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFee struct {
	PlatformFee uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomMarketplaceUpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceUpdatePlatformFeeIterator{contract: _FantomMarketplace.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceUpdatePlatformFee)
				if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseUpdatePlatformFee(log types.Log) (*FantomMarketplaceUpdatePlatformFee, error) {
	event := new(FantomMarketplaceUpdatePlatformFee)
	if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceUpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFeeRecipientIterator struct {
	Event *FantomMarketplaceUpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceUpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceUpdatePlatformFeeRecipient)
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
		it.Event = new(FantomMarketplaceUpdatePlatformFeeRecipient)
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
func (it *FantomMarketplaceUpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceUpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceUpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*FantomMarketplaceUpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceUpdatePlatformFeeRecipientIterator{contract: _FantomMarketplace.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceUpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceUpdatePlatformFeeRecipient)
				if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*FantomMarketplaceUpdatePlatformFeeRecipient, error) {
	event := new(FantomMarketplaceUpdatePlatformFeeRecipient)
	if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
