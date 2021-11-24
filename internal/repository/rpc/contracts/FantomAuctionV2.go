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

// FantomAuctionV2MetaData contains all meta data concerning the FantomAuctionV2 contract.
var FantomAuctionV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unitPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"}],\"name\":\"AuctionResulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FantomAuctionContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PauseToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"UpdateBidWithdrawalLockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"UpdateMinBidIncrement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resulted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidWithdrawalLockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"minBidReserve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_resulted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minBid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getHighestBidder\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBids\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAuctionLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"reclaimERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"resultAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"resultFailedAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleIsPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"updateAuctionReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"updateBidWithdrawalLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"updateMinBidIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomAuctionV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomAuctionV2MetaData.ABI instead.
var FantomAuctionV2ABI = FantomAuctionV2MetaData.ABI

// FantomAuctionV2 is an auto generated Go binding around an Ethereum contract.
type FantomAuctionV2 struct {
	FantomAuctionV2Caller     // Read-only binding to the contract
	FantomAuctionV2Transactor // Write-only binding to the contract
	FantomAuctionV2Filterer   // Log filterer for contract events
}

// FantomAuctionV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FantomAuctionV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomAuctionV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomAuctionV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomAuctionV2Session struct {
	Contract     *FantomAuctionV2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FantomAuctionV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomAuctionV2CallerSession struct {
	Contract *FantomAuctionV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FantomAuctionV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomAuctionV2TransactorSession struct {
	Contract     *FantomAuctionV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FantomAuctionV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FantomAuctionV2Raw struct {
	Contract *FantomAuctionV2 // Generic contract binding to access the raw methods on
}

// FantomAuctionV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomAuctionV2CallerRaw struct {
	Contract *FantomAuctionV2Caller // Generic read-only contract binding to access the raw methods on
}

// FantomAuctionV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomAuctionV2TransactorRaw struct {
	Contract *FantomAuctionV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomAuctionV2 creates a new instance of FantomAuctionV2, bound to a specific deployed contract.
func NewFantomAuctionV2(address common.Address, backend bind.ContractBackend) (*FantomAuctionV2, error) {
	contract, err := bindFantomAuctionV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2{FantomAuctionV2Caller: FantomAuctionV2Caller{contract: contract}, FantomAuctionV2Transactor: FantomAuctionV2Transactor{contract: contract}, FantomAuctionV2Filterer: FantomAuctionV2Filterer{contract: contract}}, nil
}

// NewFantomAuctionV2Caller creates a new read-only instance of FantomAuctionV2, bound to a specific deployed contract.
func NewFantomAuctionV2Caller(address common.Address, caller bind.ContractCaller) (*FantomAuctionV2Caller, error) {
	contract, err := bindFantomAuctionV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2Caller{contract: contract}, nil
}

// NewFantomAuctionV2Transactor creates a new write-only instance of FantomAuctionV2, bound to a specific deployed contract.
func NewFantomAuctionV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FantomAuctionV2Transactor, error) {
	contract, err := bindFantomAuctionV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2Transactor{contract: contract}, nil
}

// NewFantomAuctionV2Filterer creates a new log filterer instance of FantomAuctionV2, bound to a specific deployed contract.
func NewFantomAuctionV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FantomAuctionV2Filterer, error) {
	contract, err := bindFantomAuctionV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2Filterer{contract: contract}, nil
}

// bindFantomAuctionV2 binds a generic wrapper to an already deployed contract.
func bindFantomAuctionV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomAuctionV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAuctionV2 *FantomAuctionV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAuctionV2.Contract.FantomAuctionV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAuctionV2 *FantomAuctionV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.FantomAuctionV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAuctionV2 *FantomAuctionV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.FantomAuctionV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAuctionV2 *FantomAuctionV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAuctionV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAuctionV2 *FantomAuctionV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAuctionV2 *FantomAuctionV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2Caller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2Session) AddressRegistry() (common.Address, error) {
	return _FantomAuctionV2.Contract.AddressRegistry(&_FantomAuctionV2.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) AddressRegistry() (common.Address, error) {
	return _FantomAuctionV2.Contract.AddressRegistry(&_FantomAuctionV2.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuctionV2 *FantomAuctionV2Caller) Auctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "auctions", arg0, arg1)

	outstruct := new(struct {
		Owner        common.Address
		PayToken     common.Address
		MinBid       *big.Int
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
	outstruct.MinBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReservePrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Resulted = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuctionV2 *FantomAuctionV2Session) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuctionV2.Contract.Auctions(&_FantomAuctionV2.CallOpts, arg0, arg1)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuctionV2.Contract.Auctions(&_FantomAuctionV2.CallOpts, arg0, arg1)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Caller) BidWithdrawalLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "bidWithdrawalLockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Session) BidWithdrawalLockTime() (*big.Int, error) {
	return _FantomAuctionV2.Contract.BidWithdrawalLockTime(&_FantomAuctionV2.CallOpts)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) BidWithdrawalLockTime() (*big.Int, error) {
	return _FantomAuctionV2.Contract.BidWithdrawalLockTime(&_FantomAuctionV2.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_FantomAuctionV2 *FantomAuctionV2Caller) GetAuction(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "getAuction", _nftAddress, _tokenId)

	outstruct := new(struct {
		Owner        common.Address
		PayToken     common.Address
		ReservePrice *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Resulted     bool
		MinBid       *big.Int
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
	outstruct.MinBid = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_FantomAuctionV2 *FantomAuctionV2Session) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	return _FantomAuctionV2.Contract.GetAuction(&_FantomAuctionV2.CallOpts, _nftAddress, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	return _FantomAuctionV2.Contract.GetAuction(&_FantomAuctionV2.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_FantomAuctionV2 *FantomAuctionV2Caller) GetHighestBidder(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "getHighestBidder", _nftAddress, _tokenId)

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
func (_FantomAuctionV2 *FantomAuctionV2Session) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV2.Contract.GetHighestBidder(&_FantomAuctionV2.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV2.Contract.GetHighestBidder(&_FantomAuctionV2.CallOpts, _nftAddress, _tokenId)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_FantomAuctionV2 *FantomAuctionV2Caller) HighestBids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "highestBids", arg0, arg1)

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
func (_FantomAuctionV2 *FantomAuctionV2Session) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV2.Contract.HighestBids(&_FantomAuctionV2.CallOpts, arg0, arg1)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuctionV2.Contract.HighestBids(&_FantomAuctionV2.CallOpts, arg0, arg1)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuctionV2 *FantomAuctionV2Caller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuctionV2 *FantomAuctionV2Session) IsPaused() (bool, error) {
	return _FantomAuctionV2.Contract.IsPaused(&_FantomAuctionV2.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) IsPaused() (bool, error) {
	return _FantomAuctionV2.Contract.IsPaused(&_FantomAuctionV2.CallOpts)
}

// MaxAuctionLength is a free data retrieval call binding the contract method 0x0e519ef9.
//
// Solidity: function maxAuctionLength() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Caller) MaxAuctionLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "maxAuctionLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAuctionLength is a free data retrieval call binding the contract method 0x0e519ef9.
//
// Solidity: function maxAuctionLength() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Session) MaxAuctionLength() (*big.Int, error) {
	return _FantomAuctionV2.Contract.MaxAuctionLength(&_FantomAuctionV2.CallOpts)
}

// MaxAuctionLength is a free data retrieval call binding the contract method 0x0e519ef9.
//
// Solidity: function maxAuctionLength() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) MaxAuctionLength() (*big.Int, error) {
	return _FantomAuctionV2.Contract.MaxAuctionLength(&_FantomAuctionV2.CallOpts)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Caller) MinBidIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "minBidIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Session) MinBidIncrement() (*big.Int, error) {
	return _FantomAuctionV2.Contract.MinBidIncrement(&_FantomAuctionV2.CallOpts)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) MinBidIncrement() (*big.Int, error) {
	return _FantomAuctionV2.Contract.MinBidIncrement(&_FantomAuctionV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2Session) Owner() (common.Address, error) {
	return _FantomAuctionV2.Contract.Owner(&_FantomAuctionV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) Owner() (common.Address, error) {
	return _FantomAuctionV2.Contract.Owner(&_FantomAuctionV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Caller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2Session) PlatformFee() (*big.Int, error) {
	return _FantomAuctionV2.Contract.PlatformFee(&_FantomAuctionV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) PlatformFee() (*big.Int, error) {
	return _FantomAuctionV2.Contract.PlatformFee(&_FantomAuctionV2.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2Caller) PlatformFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuctionV2.contract.Call(opts, &out, "platformFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2Session) PlatformFeeRecipient() (common.Address, error) {
	return _FantomAuctionV2.Contract.PlatformFeeRecipient(&_FantomAuctionV2.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuctionV2 *FantomAuctionV2CallerSession) PlatformFeeRecipient() (common.Address, error) {
	return _FantomAuctionV2.Contract.PlatformFeeRecipient(&_FantomAuctionV2.CallOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) CancelAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "cancelAuction", _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.CancelAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.CancelAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) CreateAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "createAuction", _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.CreateAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.CreateAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) Initialize(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "initialize", _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) Initialize(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.Initialize(&_FantomAuctionV2.TransactOpts, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) Initialize(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.Initialize(&_FantomAuctionV2.TransactOpts, _platformFeeRecipient)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_FantomAuctionV2 *FantomAuctionV2Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_FantomAuctionV2 *FantomAuctionV2Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.OnERC721Received(&_FantomAuctionV2.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.OnERC721Received(&_FantomAuctionV2.TransactOpts, arg0, arg1, arg2, arg3)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) PlaceBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "placeBid", _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.PlaceBid(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.PlaceBid(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) ReclaimERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "reclaimERC20", _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ReclaimERC20(&_FantomAuctionV2.TransactOpts, _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ReclaimERC20(&_FantomAuctionV2.TransactOpts, _tokenContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.RenounceOwnership(&_FantomAuctionV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.RenounceOwnership(&_FantomAuctionV2.TransactOpts)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) ResultAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "resultAuction", _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ResultAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ResultAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// ResultFailedAuction is a paid mutator transaction binding the contract method 0xb6a9711f.
//
// Solidity: function resultFailedAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) ResultFailedAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "resultFailedAuction", _nftAddress, _tokenId)
}

// ResultFailedAuction is a paid mutator transaction binding the contract method 0xb6a9711f.
//
// Solidity: function resultFailedAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) ResultFailedAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ResultFailedAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// ResultFailedAuction is a paid mutator transaction binding the contract method 0xb6a9711f.
//
// Solidity: function resultFailedAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) ResultFailedAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ResultFailedAuction(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) ToggleIsPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "toggleIsPaused")
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) ToggleIsPaused() (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ToggleIsPaused(&_FantomAuctionV2.TransactOpts)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) ToggleIsPaused() (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.ToggleIsPaused(&_FantomAuctionV2.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.TransferOwnership(&_FantomAuctionV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.TransferOwnership(&_FantomAuctionV2.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateAddressRegistry(&_FantomAuctionV2.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateAddressRegistry(&_FantomAuctionV2.TransactOpts, _registry)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) UpdateAuctionReservePrice(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "updateAuctionReservePrice", _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateAuctionReservePrice(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateAuctionReservePrice(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) UpdateBidWithdrawalLockTime(opts *bind.TransactOpts, _bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "updateBidWithdrawalLockTime", _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateBidWithdrawalLockTime(&_FantomAuctionV2.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateBidWithdrawalLockTime(&_FantomAuctionV2.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) UpdateMinBidIncrement(opts *bind.TransactOpts, _minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "updateMinBidIncrement", _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateMinBidIncrement(&_FantomAuctionV2.TransactOpts, _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdateMinBidIncrement(&_FantomAuctionV2.TransactOpts, _minBidIncrement)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdatePlatformFee(&_FantomAuctionV2.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdatePlatformFee(&_FantomAuctionV2.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdatePlatformFeeRecipient(&_FantomAuctionV2.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.UpdatePlatformFeeRecipient(&_FantomAuctionV2.TransactOpts, _platformFeeRecipient)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Transactor) WithdrawBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.contract.Transact(opts, "withdrawBid", _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2Session) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.WithdrawBid(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuctionV2 *FantomAuctionV2TransactorSession) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuctionV2.Contract.WithdrawBid(&_FantomAuctionV2.TransactOpts, _nftAddress, _tokenId)
}

// FantomAuctionV2AuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the FantomAuctionV2 contract.
type FantomAuctionV2AuctionCancelledIterator struct {
	Event *FantomAuctionV2AuctionCancelled // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2AuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2AuctionCancelled)
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
		it.Event = new(FantomAuctionV2AuctionCancelled)
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
func (it *FantomAuctionV2AuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2AuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2AuctionCancelled represents a AuctionCancelled event raised by the FantomAuctionV2 contract.
type FantomAuctionV2AuctionCancelled struct {
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterAuctionCancelled(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV2AuctionCancelledIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2AuctionCancelledIterator{contract: _FantomAuctionV2.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2AuctionCancelled, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2AuctionCancelled)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseAuctionCancelled(log types.Log) (*FantomAuctionV2AuctionCancelled, error) {
	event := new(FantomAuctionV2AuctionCancelled)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2AuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the FantomAuctionV2 contract.
type FantomAuctionV2AuctionCreatedIterator struct {
	Event *FantomAuctionV2AuctionCreated // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2AuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2AuctionCreated)
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
		it.Event = new(FantomAuctionV2AuctionCreated)
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
func (it *FantomAuctionV2AuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2AuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2AuctionCreated represents a AuctionCreated event raised by the FantomAuctionV2 contract.
type FantomAuctionV2AuctionCreated struct {
	NftAddress common.Address
	TokenId    *big.Int
	PayToken   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterAuctionCreated(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV2AuctionCreatedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2AuctionCreatedIterator{contract: _FantomAuctionV2.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2AuctionCreated, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2AuctionCreated)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseAuctionCreated(log types.Log) (*FantomAuctionV2AuctionCreated, error) {
	event := new(FantomAuctionV2AuctionCreated)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2AuctionResultedIterator is returned from FilterAuctionResulted and is used to iterate over the raw logs and unpacked data for AuctionResulted events raised by the FantomAuctionV2 contract.
type FantomAuctionV2AuctionResultedIterator struct {
	Event *FantomAuctionV2AuctionResulted // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2AuctionResultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2AuctionResulted)
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
		it.Event = new(FantomAuctionV2AuctionResulted)
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
func (it *FantomAuctionV2AuctionResultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2AuctionResultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2AuctionResulted represents a AuctionResulted event raised by the FantomAuctionV2 contract.
type FantomAuctionV2AuctionResulted struct {
	OldOwner   common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Winner     common.Address
	PayToken   common.Address
	UnitPrice  *big.Int
	WinningBid *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionResulted is a free log retrieval operation binding the contract event 0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b.
//
// Solidity: event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterAuctionResulted(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (*FantomAuctionV2AuctionResultedIterator, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2AuctionResultedIterator{contract: _FantomAuctionV2.contract, event: "AuctionResulted", logs: logs, sub: sub}, nil
}

// WatchAuctionResulted is a free log subscription operation binding the contract event 0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b.
//
// Solidity: event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchAuctionResulted(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2AuctionResulted, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2AuctionResulted)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
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

// ParseAuctionResulted is a log parse operation binding the contract event 0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b.
//
// Solidity: event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseAuctionResulted(log types.Log) (*FantomAuctionV2AuctionResulted, error) {
	event := new(FantomAuctionV2AuctionResulted)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2BidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the FantomAuctionV2 contract.
type FantomAuctionV2BidPlacedIterator struct {
	Event *FantomAuctionV2BidPlaced // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2BidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2BidPlaced)
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
		it.Event = new(FantomAuctionV2BidPlaced)
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
func (it *FantomAuctionV2BidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2BidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2BidPlaced represents a BidPlaced event raised by the FantomAuctionV2 contract.
type FantomAuctionV2BidPlaced struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterBidPlaced(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionV2BidPlacedIterator, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2BidPlacedIterator{contract: _FantomAuctionV2.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2BidPlaced, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2BidPlaced)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseBidPlaced(log types.Log) (*FantomAuctionV2BidPlaced, error) {
	event := new(FantomAuctionV2BidPlaced)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2BidRefundedIterator is returned from FilterBidRefunded and is used to iterate over the raw logs and unpacked data for BidRefunded events raised by the FantomAuctionV2 contract.
type FantomAuctionV2BidRefundedIterator struct {
	Event *FantomAuctionV2BidRefunded // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2BidRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2BidRefunded)
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
		it.Event = new(FantomAuctionV2BidRefunded)
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
func (it *FantomAuctionV2BidRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2BidRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2BidRefunded represents a BidRefunded event raised by the FantomAuctionV2 contract.
type FantomAuctionV2BidRefunded struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidRefunded is a free log retrieval operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterBidRefunded(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionV2BidRefundedIterator, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2BidRefundedIterator{contract: _FantomAuctionV2.contract, event: "BidRefunded", logs: logs, sub: sub}, nil
}

// WatchBidRefunded is a free log subscription operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchBidRefunded(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2BidRefunded, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2BidRefunded)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "BidRefunded", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseBidRefunded(log types.Log) (*FantomAuctionV2BidRefunded, error) {
	event := new(FantomAuctionV2BidRefunded)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "BidRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2BidWithdrawnIterator is returned from FilterBidWithdrawn and is used to iterate over the raw logs and unpacked data for BidWithdrawn events raised by the FantomAuctionV2 contract.
type FantomAuctionV2BidWithdrawnIterator struct {
	Event *FantomAuctionV2BidWithdrawn // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2BidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2BidWithdrawn)
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
		it.Event = new(FantomAuctionV2BidWithdrawn)
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
func (it *FantomAuctionV2BidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2BidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2BidWithdrawn represents a BidWithdrawn event raised by the FantomAuctionV2 contract.
type FantomAuctionV2BidWithdrawn struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidWithdrawn is a free log retrieval operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterBidWithdrawn(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionV2BidWithdrawnIterator, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2BidWithdrawnIterator{contract: _FantomAuctionV2.contract, event: "BidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBidWithdrawn is a free log subscription operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchBidWithdrawn(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2BidWithdrawn, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2BidWithdrawn)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseBidWithdrawn(log types.Log) (*FantomAuctionV2BidWithdrawn, error) {
	event := new(FantomAuctionV2BidWithdrawn)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2FantomAuctionContractDeployedIterator is returned from FilterFantomAuctionContractDeployed and is used to iterate over the raw logs and unpacked data for FantomAuctionContractDeployed events raised by the FantomAuctionV2 contract.
type FantomAuctionV2FantomAuctionContractDeployedIterator struct {
	Event *FantomAuctionV2FantomAuctionContractDeployed // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2FantomAuctionContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2FantomAuctionContractDeployed)
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
		it.Event = new(FantomAuctionV2FantomAuctionContractDeployed)
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
func (it *FantomAuctionV2FantomAuctionContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2FantomAuctionContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2FantomAuctionContractDeployed represents a FantomAuctionContractDeployed event raised by the FantomAuctionV2 contract.
type FantomAuctionV2FantomAuctionContractDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFantomAuctionContractDeployed is a free log retrieval operation binding the contract event 0x237706cf61496186f74c42dca1d920d0abb19e2ea389fa58fe9906377aa75a2f.
//
// Solidity: event FantomAuctionContractDeployed()
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterFantomAuctionContractDeployed(opts *bind.FilterOpts) (*FantomAuctionV2FantomAuctionContractDeployedIterator, error) {

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "FantomAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2FantomAuctionContractDeployedIterator{contract: _FantomAuctionV2.contract, event: "FantomAuctionContractDeployed", logs: logs, sub: sub}, nil
}

// WatchFantomAuctionContractDeployed is a free log subscription operation binding the contract event 0x237706cf61496186f74c42dca1d920d0abb19e2ea389fa58fe9906377aa75a2f.
//
// Solidity: event FantomAuctionContractDeployed()
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchFantomAuctionContractDeployed(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2FantomAuctionContractDeployed) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "FantomAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2FantomAuctionContractDeployed)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "FantomAuctionContractDeployed", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseFantomAuctionContractDeployed(log types.Log) (*FantomAuctionV2FantomAuctionContractDeployed, error) {
	event := new(FantomAuctionV2FantomAuctionContractDeployed)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "FantomAuctionContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomAuctionV2 contract.
type FantomAuctionV2OwnershipTransferredIterator struct {
	Event *FantomAuctionV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2OwnershipTransferred)
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
		it.Event = new(FantomAuctionV2OwnershipTransferred)
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
func (it *FantomAuctionV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2OwnershipTransferred represents a OwnershipTransferred event raised by the FantomAuctionV2 contract.
type FantomAuctionV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomAuctionV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2OwnershipTransferredIterator{contract: _FantomAuctionV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2OwnershipTransferred)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseOwnershipTransferred(log types.Log) (*FantomAuctionV2OwnershipTransferred, error) {
	event := new(FantomAuctionV2OwnershipTransferred)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2PauseToggledIterator is returned from FilterPauseToggled and is used to iterate over the raw logs and unpacked data for PauseToggled events raised by the FantomAuctionV2 contract.
type FantomAuctionV2PauseToggledIterator struct {
	Event *FantomAuctionV2PauseToggled // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2PauseToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2PauseToggled)
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
		it.Event = new(FantomAuctionV2PauseToggled)
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
func (it *FantomAuctionV2PauseToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2PauseToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2PauseToggled represents a PauseToggled event raised by the FantomAuctionV2 contract.
type FantomAuctionV2PauseToggled struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauseToggled is a free log retrieval operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterPauseToggled(opts *bind.FilterOpts) (*FantomAuctionV2PauseToggledIterator, error) {

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2PauseToggledIterator{contract: _FantomAuctionV2.contract, event: "PauseToggled", logs: logs, sub: sub}, nil
}

// WatchPauseToggled is a free log subscription operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchPauseToggled(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2PauseToggled) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2PauseToggled)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "PauseToggled", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParsePauseToggled(log types.Log) (*FantomAuctionV2PauseToggled, error) {
	event := new(FantomAuctionV2PauseToggled)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "PauseToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2UpdateAuctionReservePriceIterator is returned from FilterUpdateAuctionReservePrice and is used to iterate over the raw logs and unpacked data for UpdateAuctionReservePrice events raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdateAuctionReservePriceIterator struct {
	Event *FantomAuctionV2UpdateAuctionReservePrice // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2UpdateAuctionReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2UpdateAuctionReservePrice)
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
		it.Event = new(FantomAuctionV2UpdateAuctionReservePrice)
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
func (it *FantomAuctionV2UpdateAuctionReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2UpdateAuctionReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2UpdateAuctionReservePrice represents a UpdateAuctionReservePrice event raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdateAuctionReservePrice struct {
	NftAddress   common.Address
	TokenId      *big.Int
	PayToken     common.Address
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionReservePrice is a free log retrieval operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterUpdateAuctionReservePrice(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionV2UpdateAuctionReservePriceIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2UpdateAuctionReservePriceIterator{contract: _FantomAuctionV2.contract, event: "UpdateAuctionReservePrice", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionReservePrice is a free log subscription operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchUpdateAuctionReservePrice(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2UpdateAuctionReservePrice, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2UpdateAuctionReservePrice)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseUpdateAuctionReservePrice(log types.Log) (*FantomAuctionV2UpdateAuctionReservePrice, error) {
	event := new(FantomAuctionV2UpdateAuctionReservePrice)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2UpdateBidWithdrawalLockTimeIterator is returned from FilterUpdateBidWithdrawalLockTime and is used to iterate over the raw logs and unpacked data for UpdateBidWithdrawalLockTime events raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdateBidWithdrawalLockTimeIterator struct {
	Event *FantomAuctionV2UpdateBidWithdrawalLockTime // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2UpdateBidWithdrawalLockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2UpdateBidWithdrawalLockTime)
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
		it.Event = new(FantomAuctionV2UpdateBidWithdrawalLockTime)
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
func (it *FantomAuctionV2UpdateBidWithdrawalLockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2UpdateBidWithdrawalLockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2UpdateBidWithdrawalLockTime represents a UpdateBidWithdrawalLockTime event raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdateBidWithdrawalLockTime struct {
	BidWithdrawalLockTime *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateBidWithdrawalLockTime is a free log retrieval operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterUpdateBidWithdrawalLockTime(opts *bind.FilterOpts) (*FantomAuctionV2UpdateBidWithdrawalLockTimeIterator, error) {

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2UpdateBidWithdrawalLockTimeIterator{contract: _FantomAuctionV2.contract, event: "UpdateBidWithdrawalLockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateBidWithdrawalLockTime is a free log subscription operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchUpdateBidWithdrawalLockTime(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2UpdateBidWithdrawalLockTime) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2UpdateBidWithdrawalLockTime)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseUpdateBidWithdrawalLockTime(log types.Log) (*FantomAuctionV2UpdateBidWithdrawalLockTime, error) {
	event := new(FantomAuctionV2UpdateBidWithdrawalLockTime)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2UpdateMinBidIncrementIterator is returned from FilterUpdateMinBidIncrement and is used to iterate over the raw logs and unpacked data for UpdateMinBidIncrement events raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdateMinBidIncrementIterator struct {
	Event *FantomAuctionV2UpdateMinBidIncrement // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2UpdateMinBidIncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2UpdateMinBidIncrement)
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
		it.Event = new(FantomAuctionV2UpdateMinBidIncrement)
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
func (it *FantomAuctionV2UpdateMinBidIncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2UpdateMinBidIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2UpdateMinBidIncrement represents a UpdateMinBidIncrement event raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdateMinBidIncrement struct {
	MinBidIncrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinBidIncrement is a free log retrieval operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterUpdateMinBidIncrement(opts *bind.FilterOpts) (*FantomAuctionV2UpdateMinBidIncrementIterator, error) {

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2UpdateMinBidIncrementIterator{contract: _FantomAuctionV2.contract, event: "UpdateMinBidIncrement", logs: logs, sub: sub}, nil
}

// WatchUpdateMinBidIncrement is a free log subscription operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchUpdateMinBidIncrement(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2UpdateMinBidIncrement) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2UpdateMinBidIncrement)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseUpdateMinBidIncrement(log types.Log) (*FantomAuctionV2UpdateMinBidIncrement, error) {
	event := new(FantomAuctionV2UpdateMinBidIncrement)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2UpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdatePlatformFeeIterator struct {
	Event *FantomAuctionV2UpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2UpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2UpdatePlatformFee)
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
		it.Event = new(FantomAuctionV2UpdatePlatformFee)
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
func (it *FantomAuctionV2UpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2UpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2UpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomAuctionV2UpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2UpdatePlatformFeeIterator{contract: _FantomAuctionV2.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2UpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2UpdatePlatformFee)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseUpdatePlatformFee(log types.Log) (*FantomAuctionV2UpdatePlatformFee, error) {
	event := new(FantomAuctionV2UpdatePlatformFee)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionV2UpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdatePlatformFeeRecipientIterator struct {
	Event *FantomAuctionV2UpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomAuctionV2UpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionV2UpdatePlatformFeeRecipient)
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
		it.Event = new(FantomAuctionV2UpdatePlatformFeeRecipient)
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
func (it *FantomAuctionV2UpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionV2UpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionV2UpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the FantomAuctionV2 contract.
type FantomAuctionV2UpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*FantomAuctionV2UpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _FantomAuctionV2.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionV2UpdatePlatformFeeRecipientIterator{contract: _FantomAuctionV2.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomAuctionV2 *FantomAuctionV2Filterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomAuctionV2UpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomAuctionV2.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionV2UpdatePlatformFeeRecipient)
				if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_FantomAuctionV2 *FantomAuctionV2Filterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*FantomAuctionV2UpdatePlatformFeeRecipient, error) {
	event := new(FantomAuctionV2UpdatePlatformFeeRecipient)
	if err := _FantomAuctionV2.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
