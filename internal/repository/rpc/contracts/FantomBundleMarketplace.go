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

// FantomBundleMarketplaceMetaData contains all meta data concerning the FantomBundleMarketplace contract.
var FantomBundleMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"}],\"name\":\"ItemCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"name\":\"ItemListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unitPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nft\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"quantity\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ItemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"}],\"name\":\"OfferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"}],\"name\":\"getListing\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_nftAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTime\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"validateItemSold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomBundleMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomBundleMarketplaceMetaData.ABI instead.
var FantomBundleMarketplaceABI = FantomBundleMarketplaceMetaData.ABI

// FantomBundleMarketplace is an auto generated Go binding around an Ethereum contract.
type FantomBundleMarketplace struct {
	FantomBundleMarketplaceCaller     // Read-only binding to the contract
	FantomBundleMarketplaceTransactor // Write-only binding to the contract
	FantomBundleMarketplaceFilterer   // Log filterer for contract events
}

// FantomBundleMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomBundleMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomBundleMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomBundleMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomBundleMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomBundleMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomBundleMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomBundleMarketplaceSession struct {
	Contract     *FantomBundleMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FantomBundleMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomBundleMarketplaceCallerSession struct {
	Contract *FantomBundleMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// FantomBundleMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomBundleMarketplaceTransactorSession struct {
	Contract     *FantomBundleMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// FantomBundleMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomBundleMarketplaceRaw struct {
	Contract *FantomBundleMarketplace // Generic contract binding to access the raw methods on
}

// FantomBundleMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomBundleMarketplaceCallerRaw struct {
	Contract *FantomBundleMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// FantomBundleMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomBundleMarketplaceTransactorRaw struct {
	Contract *FantomBundleMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomBundleMarketplace creates a new instance of FantomBundleMarketplace, bound to a specific deployed contract.
func NewFantomBundleMarketplace(address common.Address, backend bind.ContractBackend) (*FantomBundleMarketplace, error) {
	contract, err := bindFantomBundleMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplace{FantomBundleMarketplaceCaller: FantomBundleMarketplaceCaller{contract: contract}, FantomBundleMarketplaceTransactor: FantomBundleMarketplaceTransactor{contract: contract}, FantomBundleMarketplaceFilterer: FantomBundleMarketplaceFilterer{contract: contract}}, nil
}

// NewFantomBundleMarketplaceCaller creates a new read-only instance of FantomBundleMarketplace, bound to a specific deployed contract.
func NewFantomBundleMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*FantomBundleMarketplaceCaller, error) {
	contract, err := bindFantomBundleMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceCaller{contract: contract}, nil
}

// NewFantomBundleMarketplaceTransactor creates a new write-only instance of FantomBundleMarketplace, bound to a specific deployed contract.
func NewFantomBundleMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomBundleMarketplaceTransactor, error) {
	contract, err := bindFantomBundleMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceTransactor{contract: contract}, nil
}

// NewFantomBundleMarketplaceFilterer creates a new log filterer instance of FantomBundleMarketplace, bound to a specific deployed contract.
func NewFantomBundleMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomBundleMarketplaceFilterer, error) {
	contract, err := bindFantomBundleMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceFilterer{contract: contract}, nil
}

// bindFantomBundleMarketplace binds a generic wrapper to an already deployed contract.
func bindFantomBundleMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomBundleMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomBundleMarketplace *FantomBundleMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomBundleMarketplace.Contract.FantomBundleMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomBundleMarketplace *FantomBundleMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.FantomBundleMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomBundleMarketplace *FantomBundleMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.FantomBundleMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomBundleMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) AddressRegistry() (common.Address, error) {
	return _FantomBundleMarketplace.Contract.AddressRegistry(&_FantomBundleMarketplace.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) AddressRegistry() (common.Address, error) {
	return _FantomBundleMarketplace.Contract.AddressRegistry(&_FantomBundleMarketplace.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) FeeReceipient() (common.Address, error) {
	return _FantomBundleMarketplace.Contract.FeeReceipient(&_FantomBundleMarketplace.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) FeeReceipient() (common.Address, error) {
	return _FantomBundleMarketplace.Contract.FeeReceipient(&_FantomBundleMarketplace.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x49deb977.
//
// Solidity: function getListing(address _owner, string _bundleID) view returns(address[] nfts, uint256[] tokenIds, uint256[] quantities, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) GetListing(opts *bind.CallOpts, _owner common.Address, _bundleID string) (struct {
	Nfts         []common.Address
	TokenIds     []*big.Int
	Quantities   []*big.Int
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "getListing", _owner, _bundleID)

	outstruct := new(struct {
		Nfts         []common.Address
		TokenIds     []*big.Int
		Quantities   []*big.Int
		Price        *big.Int
		StartingTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nfts = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Quantities = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetListing is a free data retrieval call binding the contract method 0x49deb977.
//
// Solidity: function getListing(address _owner, string _bundleID) view returns(address[] nfts, uint256[] tokenIds, uint256[] quantities, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) GetListing(_owner common.Address, _bundleID string) (struct {
	Nfts         []common.Address
	TokenIds     []*big.Int
	Quantities   []*big.Int
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomBundleMarketplace.Contract.GetListing(&_FantomBundleMarketplace.CallOpts, _owner, _bundleID)
}

// GetListing is a free data retrieval call binding the contract method 0x49deb977.
//
// Solidity: function getListing(address _owner, string _bundleID) view returns(address[] nfts, uint256[] tokenIds, uint256[] quantities, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) GetListing(_owner common.Address, _bundleID string) (struct {
	Nfts         []common.Address
	TokenIds     []*big.Int
	Quantities   []*big.Int
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomBundleMarketplace.Contract.GetListing(&_FantomBundleMarketplace.CallOpts, _owner, _bundleID)
}

// Listings is a free data retrieval call binding the contract method 0xaa3a6b36.
//
// Solidity: function listings(address , bytes32 ) view returns(address payToken, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "listings", arg0, arg1)

	outstruct := new(struct {
		PayToken     common.Address
		Price        *big.Int
		StartingTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xaa3a6b36.
//
// Solidity: function listings(address , bytes32 ) view returns(address payToken, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) Listings(arg0 common.Address, arg1 [32]byte) (struct {
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomBundleMarketplace.Contract.Listings(&_FantomBundleMarketplace.CallOpts, arg0, arg1)
}

// Listings is a free data retrieval call binding the contract method 0xaa3a6b36.
//
// Solidity: function listings(address , bytes32 ) view returns(address payToken, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) Listings(arg0 common.Address, arg1 [32]byte) (struct {
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomBundleMarketplace.Contract.Listings(&_FantomBundleMarketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xa5851902.
//
// Solidity: function offers(bytes32 , address ) view returns(address payToken, uint256 price, uint256 deadline)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) Offers(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
}, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "offers", arg0, arg1)

	outstruct := new(struct {
		PayToken common.Address
		Price    *big.Int
		Deadline *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0xa5851902.
//
// Solidity: function offers(bytes32 , address ) view returns(address payToken, uint256 price, uint256 deadline)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) Offers(arg0 [32]byte, arg1 common.Address) (struct {
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
}, error) {
	return _FantomBundleMarketplace.Contract.Offers(&_FantomBundleMarketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xa5851902.
//
// Solidity: function offers(bytes32 , address ) view returns(address payToken, uint256 price, uint256 deadline)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) Offers(arg0 [32]byte, arg1 common.Address) (struct {
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
}, error) {
	return _FantomBundleMarketplace.Contract.Offers(&_FantomBundleMarketplace.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) Owner() (common.Address, error) {
	return _FantomBundleMarketplace.Contract.Owner(&_FantomBundleMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) Owner() (common.Address, error) {
	return _FantomBundleMarketplace.Contract.Owner(&_FantomBundleMarketplace.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xfb8ad6ff.
//
// Solidity: function owners(bytes32 ) view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) Owners(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xfb8ad6ff.
//
// Solidity: function owners(bytes32 ) view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) Owners(arg0 [32]byte) (common.Address, error) {
	return _FantomBundleMarketplace.Contract.Owners(&_FantomBundleMarketplace.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0xfb8ad6ff.
//
// Solidity: function owners(bytes32 ) view returns(address)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) Owners(arg0 [32]byte) (common.Address, error) {
	return _FantomBundleMarketplace.Contract.Owners(&_FantomBundleMarketplace.CallOpts, arg0)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomBundleMarketplace.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) PlatformFee() (*big.Int, error) {
	return _FantomBundleMarketplace.Contract.PlatformFee(&_FantomBundleMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomBundleMarketplace *FantomBundleMarketplaceCallerSession) PlatformFee() (*big.Int, error) {
	return _FantomBundleMarketplace.Contract.PlatformFee(&_FantomBundleMarketplace.CallOpts)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x84328e20.
//
// Solidity: function acceptOffer(string _bundleID, address _creator) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) AcceptOffer(opts *bind.TransactOpts, _bundleID string, _creator common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "acceptOffer", _bundleID, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x84328e20.
//
// Solidity: function acceptOffer(string _bundleID, address _creator) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) AcceptOffer(_bundleID string, _creator common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.AcceptOffer(&_FantomBundleMarketplace.TransactOpts, _bundleID, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x84328e20.
//
// Solidity: function acceptOffer(string _bundleID, address _creator) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) AcceptOffer(_bundleID string, _creator common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.AcceptOffer(&_FantomBundleMarketplace.TransactOpts, _bundleID, _creator)
}

// BuyItem is a paid mutator transaction binding the contract method 0x13f19926.
//
// Solidity: function buyItem(string _bundleID) payable returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) BuyItem(opts *bind.TransactOpts, _bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "buyItem", _bundleID)
}

// BuyItem is a paid mutator transaction binding the contract method 0x13f19926.
//
// Solidity: function buyItem(string _bundleID) payable returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) BuyItem(_bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.BuyItem(&_FantomBundleMarketplace.TransactOpts, _bundleID)
}

// BuyItem is a paid mutator transaction binding the contract method 0x13f19926.
//
// Solidity: function buyItem(string _bundleID) payable returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) BuyItem(_bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.BuyItem(&_FantomBundleMarketplace.TransactOpts, _bundleID)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x9d8820c4.
//
// Solidity: function buyItem(string _bundleID, address _payToken) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) BuyItem0(opts *bind.TransactOpts, _bundleID string, _payToken common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "buyItem0", _bundleID, _payToken)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x9d8820c4.
//
// Solidity: function buyItem(string _bundleID, address _payToken) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) BuyItem0(_bundleID string, _payToken common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.BuyItem0(&_FantomBundleMarketplace.TransactOpts, _bundleID, _payToken)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x9d8820c4.
//
// Solidity: function buyItem(string _bundleID, address _payToken) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) BuyItem0(_bundleID string, _payToken common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.BuyItem0(&_FantomBundleMarketplace.TransactOpts, _bundleID, _payToken)
}

// CancelListing is a paid mutator transaction binding the contract method 0x943f2261.
//
// Solidity: function cancelListing(string _bundleID) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) CancelListing(opts *bind.TransactOpts, _bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "cancelListing", _bundleID)
}

// CancelListing is a paid mutator transaction binding the contract method 0x943f2261.
//
// Solidity: function cancelListing(string _bundleID) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) CancelListing(_bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.CancelListing(&_FantomBundleMarketplace.TransactOpts, _bundleID)
}

// CancelListing is a paid mutator transaction binding the contract method 0x943f2261.
//
// Solidity: function cancelListing(string _bundleID) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) CancelListing(_bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.CancelListing(&_FantomBundleMarketplace.TransactOpts, _bundleID)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x2d63ce27.
//
// Solidity: function cancelOffer(string _bundleID) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) CancelOffer(opts *bind.TransactOpts, _bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "cancelOffer", _bundleID)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x2d63ce27.
//
// Solidity: function cancelOffer(string _bundleID) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) CancelOffer(_bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.CancelOffer(&_FantomBundleMarketplace.TransactOpts, _bundleID)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x2d63ce27.
//
// Solidity: function cancelOffer(string _bundleID) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) CancelOffer(_bundleID string) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.CancelOffer(&_FantomBundleMarketplace.TransactOpts, _bundleID)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xe831eebf.
//
// Solidity: function createOffer(string _bundleID, address _payToken, uint256 _price, uint256 _deadline) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) CreateOffer(opts *bind.TransactOpts, _bundleID string, _payToken common.Address, _price *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "createOffer", _bundleID, _payToken, _price, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xe831eebf.
//
// Solidity: function createOffer(string _bundleID, address _payToken, uint256 _price, uint256 _deadline) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) CreateOffer(_bundleID string, _payToken common.Address, _price *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.CreateOffer(&_FantomBundleMarketplace.TransactOpts, _bundleID, _payToken, _price, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xe831eebf.
//
// Solidity: function createOffer(string _bundleID, address _payToken, uint256 _price, uint256 _deadline) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) CreateOffer(_bundleID string, _payToken common.Address, _price *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.CreateOffer(&_FantomBundleMarketplace.TransactOpts, _bundleID, _payToken, _price, _deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) Initialize(opts *bind.TransactOpts, _feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "initialize", _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) Initialize(_feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.Initialize(&_FantomBundleMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) Initialize(_feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.Initialize(&_FantomBundleMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// ListItem is a paid mutator transaction binding the contract method 0x4994bfe5.
//
// Solidity: function listItem(string _bundleID, address[] _nftAddresses, uint256[] _tokenIds, uint256[] _quantities, address _payToken, uint256 _price, uint256 _startingTime) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) ListItem(opts *bind.TransactOpts, _bundleID string, _nftAddresses []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _payToken common.Address, _price *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "listItem", _bundleID, _nftAddresses, _tokenIds, _quantities, _payToken, _price, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x4994bfe5.
//
// Solidity: function listItem(string _bundleID, address[] _nftAddresses, uint256[] _tokenIds, uint256[] _quantities, address _payToken, uint256 _price, uint256 _startingTime) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) ListItem(_bundleID string, _nftAddresses []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _payToken common.Address, _price *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.ListItem(&_FantomBundleMarketplace.TransactOpts, _bundleID, _nftAddresses, _tokenIds, _quantities, _payToken, _price, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x4994bfe5.
//
// Solidity: function listItem(string _bundleID, address[] _nftAddresses, uint256[] _tokenIds, uint256[] _quantities, address _payToken, uint256 _price, uint256 _startingTime) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) ListItem(_bundleID string, _nftAddresses []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _payToken common.Address, _price *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.ListItem(&_FantomBundleMarketplace.TransactOpts, _bundleID, _nftAddresses, _tokenIds, _quantities, _payToken, _price, _startingTime)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.RenounceOwnership(&_FantomBundleMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.RenounceOwnership(&_FantomBundleMarketplace.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.TransferOwnership(&_FantomBundleMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.TransferOwnership(&_FantomBundleMarketplace.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdateAddressRegistry(&_FantomBundleMarketplace.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdateAddressRegistry(&_FantomBundleMarketplace.TransactOpts, _registry)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xab0962dd.
//
// Solidity: function updateListing(string _bundleID, address _payToken, uint256 _newPrice) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) UpdateListing(opts *bind.TransactOpts, _bundleID string, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "updateListing", _bundleID, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xab0962dd.
//
// Solidity: function updateListing(string _bundleID, address _payToken, uint256 _newPrice) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) UpdateListing(_bundleID string, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdateListing(&_FantomBundleMarketplace.TransactOpts, _bundleID, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xab0962dd.
//
// Solidity: function updateListing(string _bundleID, address _payToken, uint256 _newPrice) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) UpdateListing(_bundleID string, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdateListing(&_FantomBundleMarketplace.TransactOpts, _bundleID, _payToken, _newPrice)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdatePlatformFee(&_FantomBundleMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdatePlatformFee(&_FantomBundleMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdatePlatformFeeRecipient(&_FantomBundleMarketplace.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.UpdatePlatformFeeRecipient(&_FantomBundleMarketplace.TransactOpts, _platformFeeRecipient)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, uint256 _quantity) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactor) ValidateItemSold(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.contract.Transact(opts, "validateItemSold", _nftAddress, _tokenId, _quantity)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, uint256 _quantity) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.ValidateItemSold(&_FantomBundleMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, uint256 _quantity) returns()
func (_FantomBundleMarketplace *FantomBundleMarketplaceTransactorSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _FantomBundleMarketplace.Contract.ValidateItemSold(&_FantomBundleMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity)
}

// FantomBundleMarketplaceItemCanceledIterator is returned from FilterItemCanceled and is used to iterate over the raw logs and unpacked data for ItemCanceled events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemCanceledIterator struct {
	Event *FantomBundleMarketplaceItemCanceled // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceItemCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceItemCanceled)
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
		it.Event = new(FantomBundleMarketplaceItemCanceled)
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
func (it *FantomBundleMarketplaceItemCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceItemCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceItemCanceled represents a ItemCanceled event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemCanceled struct {
	Owner    common.Address
	BundleID string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemCanceled is a free log retrieval operation binding the contract event 0x71a0e856c3470e868cc72d9c038e190004e167ea6490f2063682cd37679a302a.
//
// Solidity: event ItemCanceled(address indexed owner, string bundleID)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterItemCanceled(opts *bind.FilterOpts, owner []common.Address) (*FantomBundleMarketplaceItemCanceledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "ItemCanceled", ownerRule)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceItemCanceledIterator{contract: _FantomBundleMarketplace.contract, event: "ItemCanceled", logs: logs, sub: sub}, nil
}

// WatchItemCanceled is a free log subscription operation binding the contract event 0x71a0e856c3470e868cc72d9c038e190004e167ea6490f2063682cd37679a302a.
//
// Solidity: event ItemCanceled(address indexed owner, string bundleID)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchItemCanceled(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceItemCanceled, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "ItemCanceled", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceItemCanceled)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
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

// ParseItemCanceled is a log parse operation binding the contract event 0x71a0e856c3470e868cc72d9c038e190004e167ea6490f2063682cd37679a302a.
//
// Solidity: event ItemCanceled(address indexed owner, string bundleID)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseItemCanceled(log types.Log) (*FantomBundleMarketplaceItemCanceled, error) {
	event := new(FantomBundleMarketplaceItemCanceled)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceItemListedIterator is returned from FilterItemListed and is used to iterate over the raw logs and unpacked data for ItemListed events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemListedIterator struct {
	Event *FantomBundleMarketplaceItemListed // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceItemListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceItemListed)
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
		it.Event = new(FantomBundleMarketplaceItemListed)
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
func (it *FantomBundleMarketplaceItemListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceItemListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceItemListed represents a ItemListed event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemListed struct {
	Owner        common.Address
	BundleID     string
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemListed is a free log retrieval operation binding the contract event 0x1dd1f9de6505bd1f4c42382f88f7e9475dedc3c10ff5bd77cf14490a5d74fab2.
//
// Solidity: event ItemListed(address indexed owner, string bundleID, address payToken, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterItemListed(opts *bind.FilterOpts, owner []common.Address) (*FantomBundleMarketplaceItemListedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "ItemListed", ownerRule)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceItemListedIterator{contract: _FantomBundleMarketplace.contract, event: "ItemListed", logs: logs, sub: sub}, nil
}

// WatchItemListed is a free log subscription operation binding the contract event 0x1dd1f9de6505bd1f4c42382f88f7e9475dedc3c10ff5bd77cf14490a5d74fab2.
//
// Solidity: event ItemListed(address indexed owner, string bundleID, address payToken, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchItemListed(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceItemListed, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "ItemListed", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceItemListed)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
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

// ParseItemListed is a log parse operation binding the contract event 0x1dd1f9de6505bd1f4c42382f88f7e9475dedc3c10ff5bd77cf14490a5d74fab2.
//
// Solidity: event ItemListed(address indexed owner, string bundleID, address payToken, uint256 price, uint256 startingTime)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseItemListed(log types.Log) (*FantomBundleMarketplaceItemListed, error) {
	event := new(FantomBundleMarketplaceItemListed)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemSoldIterator struct {
	Event *FantomBundleMarketplaceItemSold // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceItemSold)
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
		it.Event = new(FantomBundleMarketplaceItemSold)
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
func (it *FantomBundleMarketplaceItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceItemSold represents a ItemSold event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemSold struct {
	Seller    common.Address
	Buyer     common.Address
	BundleID  string
	PayToken  common.Address
	UnitPrice *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0x586a7960ce4f3631680ea77efd69fa758b24388a9a5183f452209aa9f1c74bfd.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, string bundleID, address payToken, int256 unitPrice, uint256 price)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterItemSold(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address) (*FantomBundleMarketplaceItemSoldIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "ItemSold", sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceItemSoldIterator{contract: _FantomBundleMarketplace.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x586a7960ce4f3631680ea77efd69fa758b24388a9a5183f452209aa9f1c74bfd.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, string bundleID, address payToken, int256 unitPrice, uint256 price)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceItemSold, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "ItemSold", sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceItemSold)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0x586a7960ce4f3631680ea77efd69fa758b24388a9a5183f452209aa9f1c74bfd.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, string bundleID, address payToken, int256 unitPrice, uint256 price)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseItemSold(log types.Log) (*FantomBundleMarketplaceItemSold, error) {
	event := new(FantomBundleMarketplaceItemSold)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceItemUpdatedIterator is returned from FilterItemUpdated and is used to iterate over the raw logs and unpacked data for ItemUpdated events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemUpdatedIterator struct {
	Event *FantomBundleMarketplaceItemUpdated // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceItemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceItemUpdated)
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
		it.Event = new(FantomBundleMarketplaceItemUpdated)
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
func (it *FantomBundleMarketplaceItemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceItemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceItemUpdated represents a ItemUpdated event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceItemUpdated struct {
	Owner    common.Address
	BundleID string
	Nft      []common.Address
	TokenId  []*big.Int
	Quantity []*big.Int
	PayToken common.Address
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemUpdated is a free log retrieval operation binding the contract event 0xbf004834f59fa5d831659d6ca458fe7e8675daefebc459ac53a3cab2b6816bf4.
//
// Solidity: event ItemUpdated(address indexed owner, string bundleID, address[] nft, uint256[] tokenId, uint256[] quantity, address payToken, uint256 newPrice)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterItemUpdated(opts *bind.FilterOpts, owner []common.Address) (*FantomBundleMarketplaceItemUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "ItemUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceItemUpdatedIterator{contract: _FantomBundleMarketplace.contract, event: "ItemUpdated", logs: logs, sub: sub}, nil
}

// WatchItemUpdated is a free log subscription operation binding the contract event 0xbf004834f59fa5d831659d6ca458fe7e8675daefebc459ac53a3cab2b6816bf4.
//
// Solidity: event ItemUpdated(address indexed owner, string bundleID, address[] nft, uint256[] tokenId, uint256[] quantity, address payToken, uint256 newPrice)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchItemUpdated(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceItemUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "ItemUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceItemUpdated)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
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

// ParseItemUpdated is a log parse operation binding the contract event 0xbf004834f59fa5d831659d6ca458fe7e8675daefebc459ac53a3cab2b6816bf4.
//
// Solidity: event ItemUpdated(address indexed owner, string bundleID, address[] nft, uint256[] tokenId, uint256[] quantity, address payToken, uint256 newPrice)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseItemUpdated(log types.Log) (*FantomBundleMarketplaceItemUpdated, error) {
	event := new(FantomBundleMarketplaceItemUpdated)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceOfferCanceledIterator struct {
	Event *FantomBundleMarketplaceOfferCanceled // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceOfferCanceled)
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
		it.Event = new(FantomBundleMarketplaceOfferCanceled)
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
func (it *FantomBundleMarketplaceOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceOfferCanceled represents a OfferCanceled event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceOfferCanceled struct {
	Creator  common.Address
	BundleID string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0x7d7a082d5e9005cb920fdc04023f730c7776927c8a728347783ac7a8d2b31e7c.
//
// Solidity: event OfferCanceled(address indexed creator, string bundleID)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterOfferCanceled(opts *bind.FilterOpts, creator []common.Address) (*FantomBundleMarketplaceOfferCanceledIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "OfferCanceled", creatorRule)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceOfferCanceledIterator{contract: _FantomBundleMarketplace.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0x7d7a082d5e9005cb920fdc04023f730c7776927c8a728347783ac7a8d2b31e7c.
//
// Solidity: event OfferCanceled(address indexed creator, string bundleID)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceOfferCanceled, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "OfferCanceled", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceOfferCanceled)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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

// ParseOfferCanceled is a log parse operation binding the contract event 0x7d7a082d5e9005cb920fdc04023f730c7776927c8a728347783ac7a8d2b31e7c.
//
// Solidity: event OfferCanceled(address indexed creator, string bundleID)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseOfferCanceled(log types.Log) (*FantomBundleMarketplaceOfferCanceled, error) {
	event := new(FantomBundleMarketplaceOfferCanceled)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceOfferCreatedIterator struct {
	Event *FantomBundleMarketplaceOfferCreated // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceOfferCreated)
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
		it.Event = new(FantomBundleMarketplaceOfferCreated)
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
func (it *FantomBundleMarketplaceOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceOfferCreated represents a OfferCreated event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceOfferCreated struct {
	Creator  common.Address
	BundleID string
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0x6c59d59bc0c5e094ffd184dd57ea02b8f8e2e2299fd34ffebdd4888864b547b9.
//
// Solidity: event OfferCreated(address indexed creator, string bundleID, address payToken, uint256 price, uint256 deadline)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterOfferCreated(opts *bind.FilterOpts, creator []common.Address) (*FantomBundleMarketplaceOfferCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "OfferCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceOfferCreatedIterator{contract: _FantomBundleMarketplace.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0x6c59d59bc0c5e094ffd184dd57ea02b8f8e2e2299fd34ffebdd4888864b547b9.
//
// Solidity: event OfferCreated(address indexed creator, string bundleID, address payToken, uint256 price, uint256 deadline)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceOfferCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "OfferCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceOfferCreated)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0x6c59d59bc0c5e094ffd184dd57ea02b8f8e2e2299fd34ffebdd4888864b547b9.
//
// Solidity: event OfferCreated(address indexed creator, string bundleID, address payToken, uint256 price, uint256 deadline)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseOfferCreated(log types.Log) (*FantomBundleMarketplaceOfferCreated, error) {
	event := new(FantomBundleMarketplaceOfferCreated)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceOwnershipTransferredIterator struct {
	Event *FantomBundleMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceOwnershipTransferred)
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
		it.Event = new(FantomBundleMarketplaceOwnershipTransferred)
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
func (it *FantomBundleMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomBundleMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceOwnershipTransferredIterator{contract: _FantomBundleMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceOwnershipTransferred)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*FantomBundleMarketplaceOwnershipTransferred, error) {
	event := new(FantomBundleMarketplaceOwnershipTransferred)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceUpdatePlatformFeeIterator struct {
	Event *FantomBundleMarketplaceUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceUpdatePlatformFee)
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
		it.Event = new(FantomBundleMarketplaceUpdatePlatformFee)
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
func (it *FantomBundleMarketplaceUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceUpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceUpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomBundleMarketplaceUpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceUpdatePlatformFeeIterator{contract: _FantomBundleMarketplace.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceUpdatePlatformFee)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseUpdatePlatformFee(log types.Log) (*FantomBundleMarketplaceUpdatePlatformFee, error) {
	event := new(FantomBundleMarketplaceUpdatePlatformFee)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomBundleMarketplaceUpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceUpdatePlatformFeeRecipientIterator struct {
	Event *FantomBundleMarketplaceUpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomBundleMarketplaceUpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomBundleMarketplaceUpdatePlatformFeeRecipient)
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
		it.Event = new(FantomBundleMarketplaceUpdatePlatformFeeRecipient)
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
func (it *FantomBundleMarketplaceUpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomBundleMarketplaceUpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomBundleMarketplaceUpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the FantomBundleMarketplace contract.
type FantomBundleMarketplaceUpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*FantomBundleMarketplaceUpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _FantomBundleMarketplace.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomBundleMarketplaceUpdatePlatformFeeRecipientIterator{contract: _FantomBundleMarketplace.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomBundleMarketplaceUpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomBundleMarketplace.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomBundleMarketplaceUpdatePlatformFeeRecipient)
				if err := _FantomBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_FantomBundleMarketplace *FantomBundleMarketplaceFilterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*FantomBundleMarketplaceUpdatePlatformFeeRecipient, error) {
	event := new(FantomBundleMarketplaceUpdatePlatformFeeRecipient)
	if err := _FantomBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
