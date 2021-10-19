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

// FantomAuctionV1MetaData contains all meta data concerning the FantomAuctionV1 contract.
var FantomAuctionV1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unitPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"}],\"name\":\"AuctionResulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FantomAuctionContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PauseToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionEndTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionStartTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"UpdateBidWithdrawalLockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"UpdateMinBidIncrement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resulted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidWithdrawalLockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_resulted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getHighestBidder\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBids\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"reclaimERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"resultAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleIsPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"updateAuctionEndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"updateAuctionReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"}],\"name\":\"updateAuctionStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"updateBidWithdrawalLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"updateMinBidIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomAuctionV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomAuctionV1MetaData.ABI instead.
var FantomAuctionV1ABI = FantomAuctionV1MetaData.ABI

// FantomAuctionV1 is an auto generated Go binding around an Ethereum contract.
type FantomAuctionV1 struct {
	FantomAuctionV1Caller     // Read-only binding to the contract
	FantomAuctionV1Transactor // Write-only binding to the contract
	FantomAuctionV1Filterer   // Log filterer for contract events
}

// FantomAuctionV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type FantomAuctionV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomAuctionV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomAuctionV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomAuctionV1Session struct {
	Contract     *FantomAuctionV1  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FantomAuctionV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomAuctionV1CallerSession struct {
	Contract *FantomAuctionV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FantomAuctionV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomAuctionV1TransactorSession struct {
	Contract     *FantomAuctionV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FantomAuctionV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type FantomAuctionV1Raw struct {
	Contract *FantomAuctionV1 // Generic contract binding to access the raw methods on
}

// FantomAuctionV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomAuctionV1CallerRaw struct {
	Contract *FantomAuctionV1Caller // Generic read-only contract binding to access the raw methods on
}

// FantomAuctionV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomAuctionV1TransactorRaw struct {
	Contract *FantomAuctionV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomAuctionV1 creates a new instance of FantomAuctionV1, bound to a specific deployed contract.
func NewFantomAuctionV1(address common.Address, backend bind.ContractBackend) (*FantomAuctionV1, error) {
	contract, err := bindFantomAuctionV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1{FantomAuctionV1Caller: FantomAuctionV1Caller{contract: contract}, FantomAuctionV1Transactor: FantomAuctionV1Transactor{contract: contract}, FantomAuctionV1Filterer: FantomAuctionV1Filterer{contract: contract}}, nil
}

// NewFantomAuctionV1Caller creates a new read-only instance of FantomAuctionV1, bound to a specific deployed contract.
func NewFantomAuctionV1Caller(address common.Address, caller bind.ContractCaller) (*FantomAuctionV1Caller, error) {
	contract, err := bindFantomAuctionV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1Caller{contract: contract}, nil
}

// NewFantomAuctionV1Transactor creates a new write-only instance of FantomAuctionV1, bound to a specific deployed contract.
func NewFantomAuctionV1Transactor(address common.Address, transactor bind.ContractTransactor) (*FantomAuctionV1Transactor, error) {
	contract, err := bindFantomAuctionV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1Transactor{contract: contract}, nil
}

// NewFantomAuctionV1Filterer creates a new log filterer instance of FantomAuctionV1, bound to a specific deployed contract.
func NewFantomAuctionV1Filterer(address common.Address, filterer bind.ContractFilterer) (*FantomAuctionV1Filterer, error) {
	contract, err := bindFantomAuctionV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1Filterer{contract: contract}, nil
}

// bindFantomAuctionV1 binds a generic wrapper to an already deployed contract.
func bindFantomAuctionV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomAuctionV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAuctionV1 *FantomAuctionV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAuctionV1.Contract.FantomAuctionV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAuctionV1 *FantomAuctionV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.FantomAuctionV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAuctionV1 *FantomAuctionV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.FantomAuctionV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAuctionV1 *FantomAuctionV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAuctionV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAuctionV1 *FantomAuctionV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAuctionV1 *FantomAuctionV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1Caller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1Session) AddressRegistry() (common.Address, error) {
	return _FantomAuctionV1.Contract.AddressRegistry(&_FantomAuctionV1.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) AddressRegistry() (common.Address, error) {
	return _FantomAuctionV1.Contract.AddressRegistry(&_FantomAuctionV1.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuctionV1 *FantomAuctionV1Caller) Auctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "auctions", arg0, arg1)

	outstruct := new(struct {
		Owner        common.Address
		PayToken     common.Address
		ReservePrice *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Resulted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReservePrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Resulted = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuctionV1 *FantomAuctionV1Session) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuctionV1.Contract.Auctions(&_FantomAuctionV1.CallOpts, arg0, arg1)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuctionV1.Contract.Auctions(&_FantomAuctionV1.CallOpts, arg0, arg1)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1Caller) BidWithdrawalLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "bidWithdrawalLockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1Session) BidWithdrawalLockTime() (*big.Int, error) {
	return _FantomAuctionV1.Contract.BidWithdrawalLockTime(&_FantomAuctionV1.CallOpts)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) BidWithdrawalLockTime() (*big.Int, error) {
	return _FantomAuctionV1.Contract.BidWithdrawalLockTime(&_FantomAuctionV1.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted)
func (_FantomAuctionV1 *FantomAuctionV1Caller) GetAuction(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "getAuction", _nftAddress, _tokenId)

	outstruct := new(struct {
		Owner        common.Address
		PayToken     common.Address
		ReservePrice *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Resulted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReservePrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Resulted = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted)
func (_FantomAuctionV1 *FantomAuctionV1Session) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuctionV1.Contract.GetAuction(&_FantomAuctionV1.CallOpts, _nftAddress, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuctionV1.Contract.GetAuction(&_FantomAuctionV1.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_FantomAuctionV1 *FantomAuctionV1Caller) GetHighestBidder(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "getHighestBidder", _nftAddress, _tokenId)

	outstruct := new(struct {
		Bidder      common.Address
		Bid         *big.Int
		LastBidTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Bid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastBidTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_FantomAuctionV1 *FantomAuctionV1Session) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV1.Contract.GetHighestBidder(&_FantomAuctionV1.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV1.Contract.GetHighestBidder(&_FantomAuctionV1.CallOpts, _nftAddress, _tokenId)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_FantomAuctionV1 *FantomAuctionV1Caller) HighestBids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "highestBids", arg0, arg1)

	outstruct := new(struct {
		Bidder      common.Address
		Bid         *big.Int
		LastBidTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Bid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastBidTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_FantomAuctionV1 *FantomAuctionV1Session) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV1.Contract.HighestBids(&_FantomAuctionV1.CallOpts, arg0, arg1)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV1.Contract.HighestBids(&_FantomAuctionV1.CallOpts, arg0, arg1)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuctionV1 *FantomAuctionV1Caller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuctionV1 *FantomAuctionV1Session) IsPaused() (bool, error) {
	return _FantomAuctionV1.Contract.IsPaused(&_FantomAuctionV1.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) IsPaused() (bool, error) {
	return _FantomAuctionV1.Contract.IsPaused(&_FantomAuctionV1.CallOpts)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1Caller) MinBidIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "minBidIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1Session) MinBidIncrement() (*big.Int, error) {
	return _FantomAuctionV1.Contract.MinBidIncrement(&_FantomAuctionV1.CallOpts)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) MinBidIncrement() (*big.Int, error) {
	return _FantomAuctionV1.Contract.MinBidIncrement(&_FantomAuctionV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1Session) Owner() (common.Address, error) {
	return _FantomAuctionV1.Contract.Owner(&_FantomAuctionV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) Owner() (common.Address, error) {
	return _FantomAuctionV1.Contract.Owner(&_FantomAuctionV1.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1Caller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1Session) PlatformFee() (*big.Int, error) {
	return _FantomAuctionV1.Contract.PlatformFee(&_FantomAuctionV1.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) PlatformFee() (*big.Int, error) {
	return _FantomAuctionV1.Contract.PlatformFee(&_FantomAuctionV1.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1Caller) PlatformFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuctionV1.contract.Call(opts, &out, "platformFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1Session) PlatformFeeRecipient() (common.Address, error) {
	return _FantomAuctionV1.Contract.PlatformFeeRecipient(&_FantomAuctionV1.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuctionV1 *FantomAuctionV1CallerSession) PlatformFeeRecipient() (common.Address, error) {
	return _FantomAuctionV1.Contract.PlatformFeeRecipient(&_FantomAuctionV1.CallOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) CancelAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "cancelAuction", _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.CancelAuction(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.CancelAuction(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x14ec4106.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, uint256 _endTimestamp) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) CreateAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "createAuction", _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x14ec4106.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, uint256 _endTimestamp) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.CreateAuction(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x14ec4106.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, uint256 _endTimestamp) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.CreateAuction(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, _endTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) Initialize(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "initialize", _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) Initialize(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.Initialize(&_FantomAuctionV1.TransactOpts, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) Initialize(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.Initialize(&_FantomAuctionV1.TransactOpts, _platformFeeRecipient)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) PlaceBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "placeBid", _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.PlaceBid(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.PlaceBid(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid0 is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId) payable returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) PlaceBid0(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "placeBid0", _nftAddress, _tokenId)
}

// PlaceBid0 is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId) payable returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) PlaceBid0(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.PlaceBid0(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// PlaceBid0 is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId) payable returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) PlaceBid0(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.PlaceBid0(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) ReclaimERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "reclaimERC20", _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.ReclaimERC20(&_FantomAuctionV1.TransactOpts, _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.ReclaimERC20(&_FantomAuctionV1.TransactOpts, _tokenContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.RenounceOwnership(&_FantomAuctionV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.RenounceOwnership(&_FantomAuctionV1.TransactOpts)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) ResultAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "resultAuction", _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.ResultAuction(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.ResultAuction(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) ToggleIsPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "toggleIsPaused")
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) ToggleIsPaused() (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.ToggleIsPaused(&_FantomAuctionV1.TransactOpts)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) ToggleIsPaused() (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.ToggleIsPaused(&_FantomAuctionV1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.TransferOwnership(&_FantomAuctionV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.TransferOwnership(&_FantomAuctionV1.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAddressRegistry(&_FantomAuctionV1.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAddressRegistry(&_FantomAuctionV1.TransactOpts, _registry)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdateAuctionEndTime(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updateAuctionEndTime", _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdateAuctionEndTime(_nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAuctionEndTime(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdateAuctionEndTime(_nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAuctionEndTime(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdateAuctionReservePrice(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updateAuctionReservePrice", _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAuctionReservePrice(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAuctionReservePrice(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdateAuctionStartTime(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updateAuctionStartTime", _nftAddress, _tokenId, _startTime)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdateAuctionStartTime(_nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAuctionStartTime(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _startTime)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdateAuctionStartTime(_nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateAuctionStartTime(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId, _startTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdateBidWithdrawalLockTime(opts *bind.TransactOpts, _bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updateBidWithdrawalLockTime", _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateBidWithdrawalLockTime(&_FantomAuctionV1.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateBidWithdrawalLockTime(&_FantomAuctionV1.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdateMinBidIncrement(opts *bind.TransactOpts, _minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updateMinBidIncrement", _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateMinBidIncrement(&_FantomAuctionV1.TransactOpts, _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdateMinBidIncrement(&_FantomAuctionV1.TransactOpts, _minBidIncrement)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdatePlatformFee(&_FantomAuctionV1.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdatePlatformFee(&_FantomAuctionV1.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdatePlatformFeeRecipient(&_FantomAuctionV1.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.UpdatePlatformFeeRecipient(&_FantomAuctionV1.TransactOpts, _platformFeeRecipient)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1Transactor) WithdrawBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.contract.Transact(opts, "withdrawBid", _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1Session) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.WithdrawBid(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV1 *FantomAuctionV1TransactorSession) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV1.Contract.WithdrawBid(&_FantomAuctionV1.TransactOpts, _nftAddress, _tokenId)
}

// FantomAuctionV1AuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the FantomAuctionV1 contract.
type FantomAuctionV1AuctionCancelledIterator struct {
	Event *FantomAuctionV1AuctionCancelled // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1AuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1AuctionCancelled)
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
		it.Event = new(FantomAuctionV1AuctionCancelled)
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
func (it *FantomAuctionV1AuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1AuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1AuctionCancelled represents a AuctionCancelled event raised by the FantomAuctionV1 contract.
type FantomAuctionV1AuctionCancelled struct {
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterAuctionCancelled(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV1AuctionCancelledIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1AuctionCancelledIterator{contract: _FantomAuctionV1.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1AuctionCancelled, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1AuctionCancelled)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseAuctionCancelled(log types.Log) (*FantomAuctionV1AuctionCancelled, error) {
	event := new(FantomAuctionV1AuctionCancelled)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1AuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the FantomAuctionV1 contract.
type FantomAuctionV1AuctionCreatedIterator struct {
	Event *FantomAuctionV1AuctionCreated // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1AuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1AuctionCreated)
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
		it.Event = new(FantomAuctionV1AuctionCreated)
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
func (it *FantomAuctionV1AuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1AuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1AuctionCreated represents a AuctionCreated event raised by the FantomAuctionV1 contract.
type FantomAuctionV1AuctionCreated struct {
	NftAddress common.Address
	TokenId    *big.Int
	PayToken   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterAuctionCreated(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV1AuctionCreatedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1AuctionCreatedIterator{contract: _FantomAuctionV1.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1AuctionCreated, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1AuctionCreated)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseAuctionCreated(log types.Log) (*FantomAuctionV1AuctionCreated, error) {
	event := new(FantomAuctionV1AuctionCreated)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1AuctionResultedIterator is returned from FilterAuctionResulted and is used to iterate over the raw logs and unpacked data for AuctionResulted events raised by the FantomAuctionV1 contract.
type FantomAuctionV1AuctionResultedIterator struct {
	Event *FantomAuctionV1AuctionResulted // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1AuctionResultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1AuctionResulted)
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
		it.Event = new(FantomAuctionV1AuctionResulted)
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
func (it *FantomAuctionV1AuctionResultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1AuctionResultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1AuctionResulted represents a AuctionResulted event raised by the FantomAuctionV1 contract.
type FantomAuctionV1AuctionResulted struct {
	NftAddress common.Address
	TokenId    *big.Int
	Winner     common.Address
	PayToken   common.Address
	UnitPrice  *big.Int
	WinningBid *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionResulted is a free log retrieval operation binding the contract event 0x0427247b7170429ed72451c5779575834ddc75e9788eb907e40025c62ed7a258.
//
// Solidity: event AuctionResulted(address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterAuctionResulted(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (*FantomAuctionV1AuctionResultedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1AuctionResultedIterator{contract: _FantomAuctionV1.contract, event: "AuctionResulted", logs: logs, sub: sub}, nil
}

// WatchAuctionResulted is a free log subscription operation binding the contract event 0x0427247b7170429ed72451c5779575834ddc75e9788eb907e40025c62ed7a258.
//
// Solidity: event AuctionResulted(address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchAuctionResulted(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1AuctionResulted, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1AuctionResulted)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
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

// ParseAuctionResulted is a log parse operation binding the contract event 0x0427247b7170429ed72451c5779575834ddc75e9788eb907e40025c62ed7a258.
//
// Solidity: event AuctionResulted(address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseAuctionResulted(log types.Log) (*FantomAuctionV1AuctionResulted, error) {
	event := new(FantomAuctionV1AuctionResulted)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1BidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the FantomAuctionV1 contract.
type FantomAuctionV1BidPlacedIterator struct {
	Event *FantomAuctionV1BidPlaced // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1BidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1BidPlaced)
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
		it.Event = new(FantomAuctionV1BidPlaced)
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
func (it *FantomAuctionV1BidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1BidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1BidPlaced represents a BidPlaced event raised by the FantomAuctionV1 contract.
type FantomAuctionV1BidPlaced struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterBidPlaced(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionV1BidPlacedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1BidPlacedIterator{contract: _FantomAuctionV1.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1BidPlaced, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1BidPlaced)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseBidPlaced(log types.Log) (*FantomAuctionV1BidPlaced, error) {
	event := new(FantomAuctionV1BidPlaced)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1BidRefundedIterator is returned from FilterBidRefunded and is used to iterate over the raw logs and unpacked data for BidRefunded events raised by the FantomAuctionV1 contract.
type FantomAuctionV1BidRefundedIterator struct {
	Event *FantomAuctionV1BidRefunded // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1BidRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1BidRefunded)
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
		it.Event = new(FantomAuctionV1BidRefunded)
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
func (it *FantomAuctionV1BidRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1BidRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1BidRefunded represents a BidRefunded event raised by the FantomAuctionV1 contract.
type FantomAuctionV1BidRefunded struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidRefunded is a free log retrieval operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterBidRefunded(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionV1BidRefundedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1BidRefundedIterator{contract: _FantomAuctionV1.contract, event: "BidRefunded", logs: logs, sub: sub}, nil
}

// WatchBidRefunded is a free log subscription operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchBidRefunded(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1BidRefunded, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1BidRefunded)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "BidRefunded", log); err != nil {
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

// ParseBidRefunded is a log parse operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseBidRefunded(log types.Log) (*FantomAuctionV1BidRefunded, error) {
	event := new(FantomAuctionV1BidRefunded)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "BidRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1BidWithdrawnIterator is returned from FilterBidWithdrawn and is used to iterate over the raw logs and unpacked data for BidWithdrawn events raised by the FantomAuctionV1 contract.
type FantomAuctionV1BidWithdrawnIterator struct {
	Event *FantomAuctionV1BidWithdrawn // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1BidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1BidWithdrawn)
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
		it.Event = new(FantomAuctionV1BidWithdrawn)
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
func (it *FantomAuctionV1BidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1BidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1BidWithdrawn represents a BidWithdrawn event raised by the FantomAuctionV1 contract.
type FantomAuctionV1BidWithdrawn struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidWithdrawn is a free log retrieval operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterBidWithdrawn(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionV1BidWithdrawnIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1BidWithdrawnIterator{contract: _FantomAuctionV1.contract, event: "BidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBidWithdrawn is a free log subscription operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchBidWithdrawn(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1BidWithdrawn, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1BidWithdrawn)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
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

// ParseBidWithdrawn is a log parse operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseBidWithdrawn(log types.Log) (*FantomAuctionV1BidWithdrawn, error) {
	event := new(FantomAuctionV1BidWithdrawn)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1FantomAuctionContractDeployedIterator is returned from FilterFantomAuctionContractDeployed and is used to iterate over the raw logs and unpacked data for FantomAuctionContractDeployed events raised by the FantomAuctionV1 contract.
type FantomAuctionV1FantomAuctionContractDeployedIterator struct {
	Event *FantomAuctionV1FantomAuctionContractDeployed // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1FantomAuctionContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1FantomAuctionContractDeployed)
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
		it.Event = new(FantomAuctionV1FantomAuctionContractDeployed)
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
func (it *FantomAuctionV1FantomAuctionContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1FantomAuctionContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1FantomAuctionContractDeployed represents a FantomAuctionContractDeployed event raised by the FantomAuctionV1 contract.
type FantomAuctionV1FantomAuctionContractDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFantomAuctionContractDeployed is a free log retrieval operation binding the contract event 0x237706cf61496186f74c42dca1d920d0abb19e2ea389fa58fe9906377aa75a2f.
//
// Solidity: event FantomAuctionContractDeployed()
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterFantomAuctionContractDeployed(opts *bind.FilterOpts) (*FantomAuctionV1FantomAuctionContractDeployedIterator, error) {

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "FantomAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1FantomAuctionContractDeployedIterator{contract: _FantomAuctionV1.contract, event: "FantomAuctionContractDeployed", logs: logs, sub: sub}, nil
}

// WatchFantomAuctionContractDeployed is a free log subscription operation binding the contract event 0x237706cf61496186f74c42dca1d920d0abb19e2ea389fa58fe9906377aa75a2f.
//
// Solidity: event FantomAuctionContractDeployed()
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchFantomAuctionContractDeployed(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1FantomAuctionContractDeployed) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "FantomAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1FantomAuctionContractDeployed)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "FantomAuctionContractDeployed", log); err != nil {
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

// ParseFantomAuctionContractDeployed is a log parse operation binding the contract event 0x237706cf61496186f74c42dca1d920d0abb19e2ea389fa58fe9906377aa75a2f.
//
// Solidity: event FantomAuctionContractDeployed()
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseFantomAuctionContractDeployed(log types.Log) (*FantomAuctionV1FantomAuctionContractDeployed, error) {
	event := new(FantomAuctionV1FantomAuctionContractDeployed)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "FantomAuctionContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomAuctionV1 contract.
type FantomAuctionV1OwnershipTransferredIterator struct {
	Event *FantomAuctionV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1OwnershipTransferred)
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
		it.Event = new(FantomAuctionV1OwnershipTransferred)
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
func (it *FantomAuctionV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1OwnershipTransferred represents a OwnershipTransferred event raised by the FantomAuctionV1 contract.
type FantomAuctionV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomAuctionV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1OwnershipTransferredIterator{contract: _FantomAuctionV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1OwnershipTransferred)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseOwnershipTransferred(log types.Log) (*FantomAuctionV1OwnershipTransferred, error) {
	event := new(FantomAuctionV1OwnershipTransferred)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1PauseToggledIterator is returned from FilterPauseToggled and is used to iterate over the raw logs and unpacked data for PauseToggled events raised by the FantomAuctionV1 contract.
type FantomAuctionV1PauseToggledIterator struct {
	Event *FantomAuctionV1PauseToggled // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1PauseToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1PauseToggled)
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
		it.Event = new(FantomAuctionV1PauseToggled)
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
func (it *FantomAuctionV1PauseToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1PauseToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1PauseToggled represents a PauseToggled event raised by the FantomAuctionV1 contract.
type FantomAuctionV1PauseToggled struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauseToggled is a free log retrieval operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterPauseToggled(opts *bind.FilterOpts) (*FantomAuctionV1PauseToggledIterator, error) {

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1PauseToggledIterator{contract: _FantomAuctionV1.contract, event: "PauseToggled", logs: logs, sub: sub}, nil
}

// WatchPauseToggled is a free log subscription operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchPauseToggled(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1PauseToggled) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1PauseToggled)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "PauseToggled", log); err != nil {
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

// ParsePauseToggled is a log parse operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParsePauseToggled(log types.Log) (*FantomAuctionV1PauseToggled, error) {
	event := new(FantomAuctionV1PauseToggled)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "PauseToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1UpdateAuctionEndTimeIterator is returned from FilterUpdateAuctionEndTime and is used to iterate over the raw logs and unpacked data for UpdateAuctionEndTime events raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateAuctionEndTimeIterator struct {
	Event *FantomAuctionV1UpdateAuctionEndTime // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1UpdateAuctionEndTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1UpdateAuctionEndTime)
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
		it.Event = new(FantomAuctionV1UpdateAuctionEndTime)
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
func (it *FantomAuctionV1UpdateAuctionEndTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1UpdateAuctionEndTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1UpdateAuctionEndTime represents a UpdateAuctionEndTime event raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateAuctionEndTime struct {
	NftAddress common.Address
	TokenId    *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionEndTime is a free log retrieval operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterUpdateAuctionEndTime(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV1UpdateAuctionEndTimeIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "UpdateAuctionEndTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1UpdateAuctionEndTimeIterator{contract: _FantomAuctionV1.contract, event: "UpdateAuctionEndTime", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionEndTime is a free log subscription operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchUpdateAuctionEndTime(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1UpdateAuctionEndTime, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "UpdateAuctionEndTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1UpdateAuctionEndTime)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateAuctionEndTime", log); err != nil {
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

// ParseUpdateAuctionEndTime is a log parse operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseUpdateAuctionEndTime(log types.Log) (*FantomAuctionV1UpdateAuctionEndTime, error) {
	event := new(FantomAuctionV1UpdateAuctionEndTime)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateAuctionEndTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1UpdateAuctionReservePriceIterator is returned from FilterUpdateAuctionReservePrice and is used to iterate over the raw logs and unpacked data for UpdateAuctionReservePrice events raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateAuctionReservePriceIterator struct {
	Event *FantomAuctionV1UpdateAuctionReservePrice // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1UpdateAuctionReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1UpdateAuctionReservePrice)
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
		it.Event = new(FantomAuctionV1UpdateAuctionReservePrice)
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
func (it *FantomAuctionV1UpdateAuctionReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1UpdateAuctionReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1UpdateAuctionReservePrice represents a UpdateAuctionReservePrice event raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateAuctionReservePrice struct {
	NftAddress   common.Address
	TokenId      *big.Int
	PayToken     common.Address
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionReservePrice is a free log retrieval operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterUpdateAuctionReservePrice(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV1UpdateAuctionReservePriceIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1UpdateAuctionReservePriceIterator{contract: _FantomAuctionV1.contract, event: "UpdateAuctionReservePrice", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionReservePrice is a free log subscription operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchUpdateAuctionReservePrice(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1UpdateAuctionReservePrice, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1UpdateAuctionReservePrice)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
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

// ParseUpdateAuctionReservePrice is a log parse operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseUpdateAuctionReservePrice(log types.Log) (*FantomAuctionV1UpdateAuctionReservePrice, error) {
	event := new(FantomAuctionV1UpdateAuctionReservePrice)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1UpdateAuctionStartTimeIterator is returned from FilterUpdateAuctionStartTime and is used to iterate over the raw logs and unpacked data for UpdateAuctionStartTime events raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateAuctionStartTimeIterator struct {
	Event *FantomAuctionV1UpdateAuctionStartTime // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1UpdateAuctionStartTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1UpdateAuctionStartTime)
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
		it.Event = new(FantomAuctionV1UpdateAuctionStartTime)
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
func (it *FantomAuctionV1UpdateAuctionStartTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1UpdateAuctionStartTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1UpdateAuctionStartTime represents a UpdateAuctionStartTime event raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateAuctionStartTime struct {
	NftAddress common.Address
	TokenId    *big.Int
	StartTime  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionStartTime is a free log retrieval operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterUpdateAuctionStartTime(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV1UpdateAuctionStartTimeIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "UpdateAuctionStartTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1UpdateAuctionStartTimeIterator{contract: _FantomAuctionV1.contract, event: "UpdateAuctionStartTime", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionStartTime is a free log subscription operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchUpdateAuctionStartTime(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1UpdateAuctionStartTime, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "UpdateAuctionStartTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1UpdateAuctionStartTime)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateAuctionStartTime", log); err != nil {
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

// ParseUpdateAuctionStartTime is a log parse operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseUpdateAuctionStartTime(log types.Log) (*FantomAuctionV1UpdateAuctionStartTime, error) {
	event := new(FantomAuctionV1UpdateAuctionStartTime)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateAuctionStartTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1UpdateBidWithdrawalLockTimeIterator is returned from FilterUpdateBidWithdrawalLockTime and is used to iterate over the raw logs and unpacked data for UpdateBidWithdrawalLockTime events raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateBidWithdrawalLockTimeIterator struct {
	Event *FantomAuctionV1UpdateBidWithdrawalLockTime // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1UpdateBidWithdrawalLockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1UpdateBidWithdrawalLockTime)
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
		it.Event = new(FantomAuctionV1UpdateBidWithdrawalLockTime)
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
func (it *FantomAuctionV1UpdateBidWithdrawalLockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1UpdateBidWithdrawalLockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1UpdateBidWithdrawalLockTime represents a UpdateBidWithdrawalLockTime event raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateBidWithdrawalLockTime struct {
	BidWithdrawalLockTime *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateBidWithdrawalLockTime is a free log retrieval operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterUpdateBidWithdrawalLockTime(opts *bind.FilterOpts) (*FantomAuctionV1UpdateBidWithdrawalLockTimeIterator, error) {

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1UpdateBidWithdrawalLockTimeIterator{contract: _FantomAuctionV1.contract, event: "UpdateBidWithdrawalLockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateBidWithdrawalLockTime is a free log subscription operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchUpdateBidWithdrawalLockTime(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1UpdateBidWithdrawalLockTime) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1UpdateBidWithdrawalLockTime)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
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

// ParseUpdateBidWithdrawalLockTime is a log parse operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseUpdateBidWithdrawalLockTime(log types.Log) (*FantomAuctionV1UpdateBidWithdrawalLockTime, error) {
	event := new(FantomAuctionV1UpdateBidWithdrawalLockTime)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1UpdateMinBidIncrementIterator is returned from FilterUpdateMinBidIncrement and is used to iterate over the raw logs and unpacked data for UpdateMinBidIncrement events raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateMinBidIncrementIterator struct {
	Event *FantomAuctionV1UpdateMinBidIncrement // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1UpdateMinBidIncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1UpdateMinBidIncrement)
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
		it.Event = new(FantomAuctionV1UpdateMinBidIncrement)
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
func (it *FantomAuctionV1UpdateMinBidIncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1UpdateMinBidIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1UpdateMinBidIncrement represents a UpdateMinBidIncrement event raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdateMinBidIncrement struct {
	MinBidIncrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinBidIncrement is a free log retrieval operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterUpdateMinBidIncrement(opts *bind.FilterOpts) (*FantomAuctionV1UpdateMinBidIncrementIterator, error) {

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1UpdateMinBidIncrementIterator{contract: _FantomAuctionV1.contract, event: "UpdateMinBidIncrement", logs: logs, sub: sub}, nil
}

// WatchUpdateMinBidIncrement is a free log subscription operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchUpdateMinBidIncrement(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1UpdateMinBidIncrement) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1UpdateMinBidIncrement)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
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

// ParseUpdateMinBidIncrement is a log parse operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseUpdateMinBidIncrement(log types.Log) (*FantomAuctionV1UpdateMinBidIncrement, error) {
	event := new(FantomAuctionV1UpdateMinBidIncrement)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1UpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdatePlatformFeeIterator struct {
	Event *FantomAuctionV1UpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1UpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1UpdatePlatformFee)
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
		it.Event = new(FantomAuctionV1UpdatePlatformFee)
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
func (it *FantomAuctionV1UpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1UpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1UpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomAuctionV1UpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1UpdatePlatformFeeIterator{contract: _FantomAuctionV1.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1UpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1UpdatePlatformFee)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseUpdatePlatformFee(log types.Log) (*FantomAuctionV1UpdatePlatformFee, error) {
	event := new(FantomAuctionV1UpdatePlatformFee)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV1UpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdatePlatformFeeRecipientIterator struct {
	Event *FantomAuctionV1UpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV1UpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV1UpdatePlatformFeeRecipient)
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
		it.Event = new(FantomAuctionV1UpdatePlatformFeeRecipient)
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
func (it *FantomAuctionV1UpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV1UpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV1UpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the FantomAuctionV1 contract.
type FantomAuctionV1UpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*FantomAuctionV1UpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _FantomAuctionV1.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV1UpdatePlatformFeeRecipientIterator{contract: _FantomAuctionV1.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomAuctionV1 *FantomAuctionV1Filterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomAuctionV1UpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV1.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV1UpdatePlatformFeeRecipient)
				if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_FantomAuctionV1 *FantomAuctionV1Filterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*FantomAuctionV1UpdatePlatformFeeRecipient, error) {
	event := new(FantomAuctionV1UpdatePlatformFeeRecipient)
	if err := _FantomAuctionV1.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
