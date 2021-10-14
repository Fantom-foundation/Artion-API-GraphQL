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

// FantomAuctionMetaData contains all meta data concerning the FantomAuction contract.
var FantomAuctionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unitPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"}],\"name\":\"AuctionResulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FantomAuctionContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PauseToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionEndTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionStartTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"UpdateBidWithdrawalLockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"UpdateMinBidIncrement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resulted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidWithdrawalLockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"minBidReserve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_resulted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minBid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getHighestBidder\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBids\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"reclaimERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"resultAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleIsPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"updateAuctionEndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"updateAuctionReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"}],\"name\":\"updateAuctionStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"updateBidWithdrawalLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"updateMinBidIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomAuctionMetaData.ABI instead.
var FantomAuctionABI = FantomAuctionMetaData.ABI

// FantomAuction is an auto generated Go binding around an Ethereum contract.
type FantomAuction struct {
	FantomAuctionCaller     // Read-only binding to the contract
	FantomAuctionTransactor // Write-only binding to the contract
	FantomAuctionFilterer   // Log filterer for contract events
}

// FantomAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomAuctionSession struct {
	Contract     *FantomAuction    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FantomAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomAuctionCallerSession struct {
	Contract *FantomAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FantomAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomAuctionTransactorSession struct {
	Contract     *FantomAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FantomAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomAuctionRaw struct {
	Contract *FantomAuction // Generic contract binding to access the raw methods on
}

// FantomAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomAuctionCallerRaw struct {
	Contract *FantomAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// FantomAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomAuctionTransactorRaw struct {
	Contract *FantomAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomAuction creates a new instance of FantomAuction, bound to a specific deployed contract.
func NewFantomAuction(address common.Address, backend bind.ContractBackend) (*FantomAuction, error) {
	contract, err := bindFantomAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomAuction{FantomAuctionCaller: FantomAuctionCaller{contract: contract}, FantomAuctionTransactor: FantomAuctionTransactor{contract: contract}, FantomAuctionFilterer: FantomAuctionFilterer{contract: contract}}, nil
}

// NewFantomAuctionCaller creates a new read-only instance of FantomAuction, bound to a specific deployed contract.
func NewFantomAuctionCaller(address common.Address, caller bind.ContractCaller) (*FantomAuctionCaller, error) {
	contract, err := bindFantomAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionCaller{contract: contract}, nil
}

// NewFantomAuctionTransactor creates a new write-only instance of FantomAuction, bound to a specific deployed contract.
func NewFantomAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomAuctionTransactor, error) {
	contract, err := bindFantomAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionTransactor{contract: contract}, nil
}

// NewFantomAuctionFilterer creates a new log filterer instance of FantomAuction, bound to a specific deployed contract.
func NewFantomAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomAuctionFilterer, error) {
	contract, err := bindFantomAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionFilterer{contract: contract}, nil
}

// bindFantomAuction binds a generic wrapper to an already deployed contract.
func bindFantomAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAuction *FantomAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAuction.Contract.FantomAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAuction *FantomAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuction.Contract.FantomAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAuction *FantomAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAuction.Contract.FantomAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAuction *FantomAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAuction *FantomAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAuction *FantomAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAuction.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuction *FantomAuctionCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuction *FantomAuctionSession) AddressRegistry() (common.Address, error) {
	return _FantomAuction.Contract.AddressRegistry(&_FantomAuction.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomAuction *FantomAuctionCallerSession) AddressRegistry() (common.Address, error) {
	return _FantomAuction.Contract.AddressRegistry(&_FantomAuction.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuction *FantomAuctionCaller) Auctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "auctions", arg0, arg1)

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
func (_FantomAuction *FantomAuctionSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuction.Contract.Auctions(&_FantomAuction.CallOpts, arg0, arg1)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_FantomAuction *FantomAuctionCallerSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _FantomAuction.Contract.Auctions(&_FantomAuction.CallOpts, arg0, arg1)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuction *FantomAuctionCaller) BidWithdrawalLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "bidWithdrawalLockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuction *FantomAuctionSession) BidWithdrawalLockTime() (*big.Int, error) {
	return _FantomAuction.Contract.BidWithdrawalLockTime(&_FantomAuction.CallOpts)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_FantomAuction *FantomAuctionCallerSession) BidWithdrawalLockTime() (*big.Int, error) {
	return _FantomAuction.Contract.BidWithdrawalLockTime(&_FantomAuction.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_FantomAuction *FantomAuctionCaller) GetAuction(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "getAuction", _nftAddress, _tokenId)

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
func (_FantomAuction *FantomAuctionSession) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	return _FantomAuction.Contract.GetAuction(&_FantomAuction.CallOpts, _nftAddress, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_FantomAuction *FantomAuctionCallerSession) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	return _FantomAuction.Contract.GetAuction(&_FantomAuction.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_FantomAuction *FantomAuctionCaller) GetHighestBidder(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "getHighestBidder", _nftAddress, _tokenId)

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
func (_FantomAuction *FantomAuctionSession) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuction.Contract.GetHighestBidder(&_FantomAuction.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_FantomAuction *FantomAuctionCallerSession) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuction.Contract.GetHighestBidder(&_FantomAuction.CallOpts, _nftAddress, _tokenId)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_FantomAuction *FantomAuctionCaller) HighestBids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "highestBids", arg0, arg1)

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
func (_FantomAuction *FantomAuctionSession) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuction.Contract.HighestBids(&_FantomAuction.CallOpts, arg0, arg1)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_FantomAuction *FantomAuctionCallerSession) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _FantomAuction.Contract.HighestBids(&_FantomAuction.CallOpts, arg0, arg1)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuction *FantomAuctionCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuction *FantomAuctionSession) IsPaused() (bool, error) {
	return _FantomAuction.Contract.IsPaused(&_FantomAuction.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_FantomAuction *FantomAuctionCallerSession) IsPaused() (bool, error) {
	return _FantomAuction.Contract.IsPaused(&_FantomAuction.CallOpts)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuction *FantomAuctionCaller) MinBidIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "minBidIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuction *FantomAuctionSession) MinBidIncrement() (*big.Int, error) {
	return _FantomAuction.Contract.MinBidIncrement(&_FantomAuction.CallOpts)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_FantomAuction *FantomAuctionCallerSession) MinBidIncrement() (*big.Int, error) {
	return _FantomAuction.Contract.MinBidIncrement(&_FantomAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuction *FantomAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuction *FantomAuctionSession) Owner() (common.Address, error) {
	return _FantomAuction.Contract.Owner(&_FantomAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAuction *FantomAuctionCallerSession) Owner() (common.Address, error) {
	return _FantomAuction.Contract.Owner(&_FantomAuction.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuction *FantomAuctionCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuction *FantomAuctionSession) PlatformFee() (*big.Int, error) {
	return _FantomAuction.Contract.PlatformFee(&_FantomAuction.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_FantomAuction *FantomAuctionCallerSession) PlatformFee() (*big.Int, error) {
	return _FantomAuction.Contract.PlatformFee(&_FantomAuction.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuction *FantomAuctionCaller) PlatformFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAuction.contract.Call(opts, &out, "platformFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuction *FantomAuctionSession) PlatformFeeRecipient() (common.Address, error) {
	return _FantomAuction.Contract.PlatformFeeRecipient(&_FantomAuction.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_FantomAuction *FantomAuctionCallerSession) PlatformFeeRecipient() (common.Address, error) {
	return _FantomAuction.Contract.PlatformFeeRecipient(&_FantomAuction.CallOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "cancelAuction", _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionSession) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.CancelAuction(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionTransactorSession) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.CancelAuction(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_FantomAuction *FantomAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "createAuction", _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_FantomAuction *FantomAuctionSession) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.CreateAuction(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_FantomAuction *FantomAuctionTransactorSession) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.CreateAuction(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuction *FantomAuctionTransactor) Initialize(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "initialize", _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuction *FantomAuctionSession) Initialize(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.Initialize(&_FantomAuction.TransactOpts, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _platformFeeRecipient) returns()
func (_FantomAuction *FantomAuctionTransactorSession) Initialize(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.Initialize(&_FantomAuction.TransactOpts, _platformFeeRecipient)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuction *FantomAuctionTransactor) PlaceBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "placeBid", _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuction *FantomAuctionSession) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.PlaceBid(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) returns()
func (_FantomAuction *FantomAuctionTransactorSession) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.PlaceBid(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid0 is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId) payable returns()
func (_FantomAuction *FantomAuctionTransactor) PlaceBid0(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "placeBid0", _nftAddress, _tokenId)
}

// PlaceBid0 is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId) payable returns()
func (_FantomAuction *FantomAuctionSession) PlaceBid0(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.PlaceBid0(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// PlaceBid0 is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId) payable returns()
func (_FantomAuction *FantomAuctionTransactorSession) PlaceBid0(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.PlaceBid0(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuction *FantomAuctionTransactor) ReclaimERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "reclaimERC20", _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuction *FantomAuctionSession) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.ReclaimERC20(&_FantomAuction.TransactOpts, _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_FantomAuction *FantomAuctionTransactorSession) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.ReclaimERC20(&_FantomAuction.TransactOpts, _tokenContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuction *FantomAuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuction *FantomAuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAuction.Contract.RenounceOwnership(&_FantomAuction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAuction *FantomAuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAuction.Contract.RenounceOwnership(&_FantomAuction.TransactOpts)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionTransactor) ResultAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "resultAuction", _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionSession) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.ResultAuction(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionTransactorSession) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.ResultAuction(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuction *FantomAuctionTransactor) ToggleIsPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "toggleIsPaused")
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuction *FantomAuctionSession) ToggleIsPaused() (*types.Transaction, error) {
	return _FantomAuction.Contract.ToggleIsPaused(&_FantomAuction.TransactOpts)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_FantomAuction *FantomAuctionTransactorSession) ToggleIsPaused() (*types.Transaction, error) {
	return _FantomAuction.Contract.ToggleIsPaused(&_FantomAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuction *FantomAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuction *FantomAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.TransferOwnership(&_FantomAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAuction *FantomAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.TransferOwnership(&_FantomAuction.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuction *FantomAuctionSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAddressRegistry(&_FantomAuction.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAddressRegistry(&_FantomAuction.TransactOpts, _registry)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdateAuctionEndTime(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updateAuctionEndTime", _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_FantomAuction *FantomAuctionSession) UpdateAuctionEndTime(_nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAuctionEndTime(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdateAuctionEndTime(_nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAuctionEndTime(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdateAuctionReservePrice(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updateAuctionReservePrice", _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuction *FantomAuctionSession) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAuctionReservePrice(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAuctionReservePrice(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdateAuctionStartTime(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updateAuctionStartTime", _nftAddress, _tokenId, _startTime)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_FantomAuction *FantomAuctionSession) UpdateAuctionStartTime(_nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAuctionStartTime(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _startTime)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdateAuctionStartTime(_nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateAuctionStartTime(&_FantomAuction.TransactOpts, _nftAddress, _tokenId, _startTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdateBidWithdrawalLockTime(opts *bind.TransactOpts, _bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updateBidWithdrawalLockTime", _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuction *FantomAuctionSession) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateBidWithdrawalLockTime(&_FantomAuction.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateBidWithdrawalLockTime(&_FantomAuction.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdateMinBidIncrement(opts *bind.TransactOpts, _minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updateMinBidIncrement", _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuction *FantomAuctionSession) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateMinBidIncrement(&_FantomAuction.TransactOpts, _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdateMinBidIncrement(&_FantomAuction.TransactOpts, _minBidIncrement)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuction *FantomAuctionSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdatePlatformFee(&_FantomAuction.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdatePlatformFee(&_FantomAuction.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuction *FantomAuctionTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuction *FantomAuctionSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdatePlatformFeeRecipient(&_FantomAuction.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomAuction *FantomAuctionTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomAuction.Contract.UpdatePlatformFeeRecipient(&_FantomAuction.TransactOpts, _platformFeeRecipient)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionTransactor) WithdrawBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.contract.Transact(opts, "withdrawBid", _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionSession) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.WithdrawBid(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_FantomAuction *FantomAuctionTransactorSession) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomAuction.Contract.WithdrawBid(&_FantomAuction.TransactOpts, _nftAddress, _tokenId)
}

// FantomAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the FantomAuction contract.
type FantomAuctionAuctionCancelledIterator struct {
	Event *FantomAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *FantomAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionAuctionCancelled)
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
		it.Event = new(FantomAuctionAuctionCancelled)
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
func (it *FantomAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionAuctionCancelled represents a AuctionCancelled event raised by the FantomAuction contract.
type FantomAuctionAuctionCancelled struct {
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_FantomAuction *FantomAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionAuctionCancelledIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionAuctionCancelledIterator{contract: _FantomAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_FantomAuction *FantomAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *FantomAuctionAuctionCancelled, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionAuctionCancelled)
				if err := _FantomAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseAuctionCancelled(log types.Log) (*FantomAuctionAuctionCancelled, error) {
	event := new(FantomAuctionAuctionCancelled)
	if err := _FantomAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the FantomAuction contract.
type FantomAuctionAuctionCreatedIterator struct {
	Event *FantomAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *FantomAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionAuctionCreated)
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
		it.Event = new(FantomAuctionAuctionCreated)
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
func (it *FantomAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionAuctionCreated represents a AuctionCreated event raised by the FantomAuction contract.
type FantomAuctionAuctionCreated struct {
	NftAddress common.Address
	TokenId    *big.Int
	PayToken   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_FantomAuction *FantomAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionAuctionCreatedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionAuctionCreatedIterator{contract: _FantomAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_FantomAuction *FantomAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *FantomAuctionAuctionCreated, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionAuctionCreated)
				if err := _FantomAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseAuctionCreated(log types.Log) (*FantomAuctionAuctionCreated, error) {
	event := new(FantomAuctionAuctionCreated)
	if err := _FantomAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionAuctionResultedIterator is returned from FilterAuctionResulted and is used to iterate over the raw logs and unpacked data for AuctionResulted events raised by the FantomAuction contract.
type FantomAuctionAuctionResultedIterator struct {
	Event *FantomAuctionAuctionResulted // Event containing the contract specifics and raw log

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
func (it *FantomAuctionAuctionResultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionAuctionResulted)
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
		it.Event = new(FantomAuctionAuctionResulted)
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
func (it *FantomAuctionAuctionResultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionAuctionResultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionAuctionResulted represents a AuctionResulted event raised by the FantomAuction contract.
type FantomAuctionAuctionResulted struct {
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
func (_FantomAuction *FantomAuctionFilterer) FilterAuctionResulted(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (*FantomAuctionAuctionResultedIterator, error) {

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

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionAuctionResultedIterator{contract: _FantomAuction.contract, event: "AuctionResulted", logs: logs, sub: sub}, nil
}

// WatchAuctionResulted is a free log subscription operation binding the contract event 0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b.
//
// Solidity: event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid)
func (_FantomAuction *FantomAuctionFilterer) WatchAuctionResulted(opts *bind.WatchOpts, sink chan<- *FantomAuctionAuctionResulted, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionAuctionResulted)
				if err := _FantomAuction.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseAuctionResulted(log types.Log) (*FantomAuctionAuctionResulted, error) {
	event := new(FantomAuctionAuctionResulted)
	if err := _FantomAuction.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the FantomAuction contract.
type FantomAuctionBidPlacedIterator struct {
	Event *FantomAuctionBidPlaced // Event containing the contract specifics and raw log

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
func (it *FantomAuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionBidPlaced)
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
		it.Event = new(FantomAuctionBidPlaced)
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
func (it *FantomAuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionBidPlaced represents a BidPlaced event raised by the FantomAuction contract.
type FantomAuctionBidPlaced struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuction *FantomAuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionBidPlacedIterator, error) {

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

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionBidPlacedIterator{contract: _FantomAuction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuction *FantomAuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *FantomAuctionBidPlaced, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionBidPlaced)
				if err := _FantomAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseBidPlaced(log types.Log) (*FantomAuctionBidPlaced, error) {
	event := new(FantomAuctionBidPlaced)
	if err := _FantomAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionBidRefundedIterator is returned from FilterBidRefunded and is used to iterate over the raw logs and unpacked data for BidRefunded events raised by the FantomAuction contract.
type FantomAuctionBidRefundedIterator struct {
	Event *FantomAuctionBidRefunded // Event containing the contract specifics and raw log

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
func (it *FantomAuctionBidRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionBidRefunded)
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
		it.Event = new(FantomAuctionBidRefunded)
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
func (it *FantomAuctionBidRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionBidRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionBidRefunded represents a BidRefunded event raised by the FantomAuction contract.
type FantomAuctionBidRefunded struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidRefunded is a free log retrieval operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuction *FantomAuctionFilterer) FilterBidRefunded(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionBidRefundedIterator, error) {

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

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionBidRefundedIterator{contract: _FantomAuction.contract, event: "BidRefunded", logs: logs, sub: sub}, nil
}

// WatchBidRefunded is a free log subscription operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuction *FantomAuctionFilterer) WatchBidRefunded(opts *bind.WatchOpts, sink chan<- *FantomAuctionBidRefunded, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionBidRefunded)
				if err := _FantomAuction.contract.UnpackLog(event, "BidRefunded", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseBidRefunded(log types.Log) (*FantomAuctionBidRefunded, error) {
	event := new(FantomAuctionBidRefunded)
	if err := _FantomAuction.contract.UnpackLog(event, "BidRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionBidWithdrawnIterator is returned from FilterBidWithdrawn and is used to iterate over the raw logs and unpacked data for BidWithdrawn events raised by the FantomAuction contract.
type FantomAuctionBidWithdrawnIterator struct {
	Event *FantomAuctionBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *FantomAuctionBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionBidWithdrawn)
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
		it.Event = new(FantomAuctionBidWithdrawn)
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
func (it *FantomAuctionBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionBidWithdrawn represents a BidWithdrawn event raised by the FantomAuction contract.
type FantomAuctionBidWithdrawn struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidWithdrawn is a free log retrieval operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuction *FantomAuctionFilterer) FilterBidWithdrawn(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*FantomAuctionBidWithdrawnIterator, error) {

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

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionBidWithdrawnIterator{contract: _FantomAuction.contract, event: "BidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBidWithdrawn is a free log subscription operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_FantomAuction *FantomAuctionFilterer) WatchBidWithdrawn(opts *bind.WatchOpts, sink chan<- *FantomAuctionBidWithdrawn, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionBidWithdrawn)
				if err := _FantomAuction.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseBidWithdrawn(log types.Log) (*FantomAuctionBidWithdrawn, error) {
	event := new(FantomAuctionBidWithdrawn)
	if err := _FantomAuction.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionFantomAuctionContractDeployedIterator is returned from FilterFantomAuctionContractDeployed and is used to iterate over the raw logs and unpacked data for FantomAuctionContractDeployed events raised by the FantomAuction contract.
type FantomAuctionFantomAuctionContractDeployedIterator struct {
	Event *FantomAuctionFantomAuctionContractDeployed // Event containing the contract specifics and raw log

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
func (it *FantomAuctionFantomAuctionContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionFantomAuctionContractDeployed)
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
		it.Event = new(FantomAuctionFantomAuctionContractDeployed)
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
func (it *FantomAuctionFantomAuctionContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionFantomAuctionContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionFantomAuctionContractDeployed represents a FantomAuctionContractDeployed event raised by the FantomAuction contract.
type FantomAuctionFantomAuctionContractDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFantomAuctionContractDeployed is a free log retrieval operation binding the contract event 0x237706cf61496186f74c42dca1d920d0abb19e2ea389fa58fe9906377aa75a2f.
//
// Solidity: event FantomAuctionContractDeployed()
func (_FantomAuction *FantomAuctionFilterer) FilterFantomAuctionContractDeployed(opts *bind.FilterOpts) (*FantomAuctionFantomAuctionContractDeployedIterator, error) {

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "FantomAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionFantomAuctionContractDeployedIterator{contract: _FantomAuction.contract, event: "FantomAuctionContractDeployed", logs: logs, sub: sub}, nil
}

// WatchFantomAuctionContractDeployed is a free log subscription operation binding the contract event 0x237706cf61496186f74c42dca1d920d0abb19e2ea389fa58fe9906377aa75a2f.
//
// Solidity: event FantomAuctionContractDeployed()
func (_FantomAuction *FantomAuctionFilterer) WatchFantomAuctionContractDeployed(opts *bind.WatchOpts, sink chan<- *FantomAuctionFantomAuctionContractDeployed) (event.Subscription, error) {

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "FantomAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionFantomAuctionContractDeployed)
				if err := _FantomAuction.contract.UnpackLog(event, "FantomAuctionContractDeployed", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseFantomAuctionContractDeployed(log types.Log) (*FantomAuctionFantomAuctionContractDeployed, error) {
	event := new(FantomAuctionFantomAuctionContractDeployed)
	if err := _FantomAuction.contract.UnpackLog(event, "FantomAuctionContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomAuction contract.
type FantomAuctionOwnershipTransferredIterator struct {
	Event *FantomAuctionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomAuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionOwnershipTransferred)
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
		it.Event = new(FantomAuctionOwnershipTransferred)
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
func (it *FantomAuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionOwnershipTransferred represents a OwnershipTransferred event raised by the FantomAuction contract.
type FantomAuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAuction *FantomAuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomAuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionOwnershipTransferredIterator{contract: _FantomAuction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAuction *FantomAuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomAuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionOwnershipTransferred)
				if err := _FantomAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseOwnershipTransferred(log types.Log) (*FantomAuctionOwnershipTransferred, error) {
	event := new(FantomAuctionOwnershipTransferred)
	if err := _FantomAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionPauseToggledIterator is returned from FilterPauseToggled and is used to iterate over the raw logs and unpacked data for PauseToggled events raised by the FantomAuction contract.
type FantomAuctionPauseToggledIterator struct {
	Event *FantomAuctionPauseToggled // Event containing the contract specifics and raw log

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
func (it *FantomAuctionPauseToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionPauseToggled)
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
		it.Event = new(FantomAuctionPauseToggled)
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
func (it *FantomAuctionPauseToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionPauseToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionPauseToggled represents a PauseToggled event raised by the FantomAuction contract.
type FantomAuctionPauseToggled struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauseToggled is a free log retrieval operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_FantomAuction *FantomAuctionFilterer) FilterPauseToggled(opts *bind.FilterOpts) (*FantomAuctionPauseToggledIterator, error) {

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionPauseToggledIterator{contract: _FantomAuction.contract, event: "PauseToggled", logs: logs, sub: sub}, nil
}

// WatchPauseToggled is a free log subscription operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_FantomAuction *FantomAuctionFilterer) WatchPauseToggled(opts *bind.WatchOpts, sink chan<- *FantomAuctionPauseToggled) (event.Subscription, error) {

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionPauseToggled)
				if err := _FantomAuction.contract.UnpackLog(event, "PauseToggled", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParsePauseToggled(log types.Log) (*FantomAuctionPauseToggled, error) {
	event := new(FantomAuctionPauseToggled)
	if err := _FantomAuction.contract.UnpackLog(event, "PauseToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionUpdateAuctionEndTimeIterator is returned from FilterUpdateAuctionEndTime and is used to iterate over the raw logs and unpacked data for UpdateAuctionEndTime events raised by the FantomAuction contract.
type FantomAuctionUpdateAuctionEndTimeIterator struct {
	Event *FantomAuctionUpdateAuctionEndTime // Event containing the contract specifics and raw log

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
func (it *FantomAuctionUpdateAuctionEndTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionUpdateAuctionEndTime)
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
		it.Event = new(FantomAuctionUpdateAuctionEndTime)
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
func (it *FantomAuctionUpdateAuctionEndTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionUpdateAuctionEndTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionUpdateAuctionEndTime represents a UpdateAuctionEndTime event raised by the FantomAuction contract.
type FantomAuctionUpdateAuctionEndTime struct {
	NftAddress common.Address
	TokenId    *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionEndTime is a free log retrieval operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_FantomAuction *FantomAuctionFilterer) FilterUpdateAuctionEndTime(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionUpdateAuctionEndTimeIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "UpdateAuctionEndTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionUpdateAuctionEndTimeIterator{contract: _FantomAuction.contract, event: "UpdateAuctionEndTime", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionEndTime is a free log subscription operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_FantomAuction *FantomAuctionFilterer) WatchUpdateAuctionEndTime(opts *bind.WatchOpts, sink chan<- *FantomAuctionUpdateAuctionEndTime, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "UpdateAuctionEndTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionUpdateAuctionEndTime)
				if err := _FantomAuction.contract.UnpackLog(event, "UpdateAuctionEndTime", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseUpdateAuctionEndTime(log types.Log) (*FantomAuctionUpdateAuctionEndTime, error) {
	event := new(FantomAuctionUpdateAuctionEndTime)
	if err := _FantomAuction.contract.UnpackLog(event, "UpdateAuctionEndTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionUpdateAuctionReservePriceIterator is returned from FilterUpdateAuctionReservePrice and is used to iterate over the raw logs and unpacked data for UpdateAuctionReservePrice events raised by the FantomAuction contract.
type FantomAuctionUpdateAuctionReservePriceIterator struct {
	Event *FantomAuctionUpdateAuctionReservePrice // Event containing the contract specifics and raw log

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
func (it *FantomAuctionUpdateAuctionReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionUpdateAuctionReservePrice)
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
		it.Event = new(FantomAuctionUpdateAuctionReservePrice)
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
func (it *FantomAuctionUpdateAuctionReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionUpdateAuctionReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionUpdateAuctionReservePrice represents a UpdateAuctionReservePrice event raised by the FantomAuction contract.
type FantomAuctionUpdateAuctionReservePrice struct {
	NftAddress   common.Address
	TokenId      *big.Int
	PayToken     common.Address
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionReservePrice is a free log retrieval operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_FantomAuction *FantomAuctionFilterer) FilterUpdateAuctionReservePrice(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionUpdateAuctionReservePriceIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionUpdateAuctionReservePriceIterator{contract: _FantomAuction.contract, event: "UpdateAuctionReservePrice", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionReservePrice is a free log subscription operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_FantomAuction *FantomAuctionFilterer) WatchUpdateAuctionReservePrice(opts *bind.WatchOpts, sink chan<- *FantomAuctionUpdateAuctionReservePrice, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionUpdateAuctionReservePrice)
				if err := _FantomAuction.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseUpdateAuctionReservePrice(log types.Log) (*FantomAuctionUpdateAuctionReservePrice, error) {
	event := new(FantomAuctionUpdateAuctionReservePrice)
	if err := _FantomAuction.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionUpdateAuctionStartTimeIterator is returned from FilterUpdateAuctionStartTime and is used to iterate over the raw logs and unpacked data for UpdateAuctionStartTime events raised by the FantomAuction contract.
type FantomAuctionUpdateAuctionStartTimeIterator struct {
	Event *FantomAuctionUpdateAuctionStartTime // Event containing the contract specifics and raw log

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
func (it *FantomAuctionUpdateAuctionStartTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionUpdateAuctionStartTime)
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
		it.Event = new(FantomAuctionUpdateAuctionStartTime)
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
func (it *FantomAuctionUpdateAuctionStartTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionUpdateAuctionStartTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionUpdateAuctionStartTime represents a UpdateAuctionStartTime event raised by the FantomAuction contract.
type FantomAuctionUpdateAuctionStartTime struct {
	NftAddress common.Address
	TokenId    *big.Int
	StartTime  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionStartTime is a free log retrieval operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_FantomAuction *FantomAuctionFilterer) FilterUpdateAuctionStartTime(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*FantomAuctionUpdateAuctionStartTimeIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "UpdateAuctionStartTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FantomAuctionUpdateAuctionStartTimeIterator{contract: _FantomAuction.contract, event: "UpdateAuctionStartTime", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionStartTime is a free log subscription operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_FantomAuction *FantomAuctionFilterer) WatchUpdateAuctionStartTime(opts *bind.WatchOpts, sink chan<- *FantomAuctionUpdateAuctionStartTime, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "UpdateAuctionStartTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionUpdateAuctionStartTime)
				if err := _FantomAuction.contract.UnpackLog(event, "UpdateAuctionStartTime", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseUpdateAuctionStartTime(log types.Log) (*FantomAuctionUpdateAuctionStartTime, error) {
	event := new(FantomAuctionUpdateAuctionStartTime)
	if err := _FantomAuction.contract.UnpackLog(event, "UpdateAuctionStartTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionUpdateBidWithdrawalLockTimeIterator is returned from FilterUpdateBidWithdrawalLockTime and is used to iterate over the raw logs and unpacked data for UpdateBidWithdrawalLockTime events raised by the FantomAuction contract.
type FantomAuctionUpdateBidWithdrawalLockTimeIterator struct {
	Event *FantomAuctionUpdateBidWithdrawalLockTime // Event containing the contract specifics and raw log

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
func (it *FantomAuctionUpdateBidWithdrawalLockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionUpdateBidWithdrawalLockTime)
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
		it.Event = new(FantomAuctionUpdateBidWithdrawalLockTime)
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
func (it *FantomAuctionUpdateBidWithdrawalLockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionUpdateBidWithdrawalLockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionUpdateBidWithdrawalLockTime represents a UpdateBidWithdrawalLockTime event raised by the FantomAuction contract.
type FantomAuctionUpdateBidWithdrawalLockTime struct {
	BidWithdrawalLockTime *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateBidWithdrawalLockTime is a free log retrieval operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_FantomAuction *FantomAuctionFilterer) FilterUpdateBidWithdrawalLockTime(opts *bind.FilterOpts) (*FantomAuctionUpdateBidWithdrawalLockTimeIterator, error) {

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionUpdateBidWithdrawalLockTimeIterator{contract: _FantomAuction.contract, event: "UpdateBidWithdrawalLockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateBidWithdrawalLockTime is a free log subscription operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_FantomAuction *FantomAuctionFilterer) WatchUpdateBidWithdrawalLockTime(opts *bind.WatchOpts, sink chan<- *FantomAuctionUpdateBidWithdrawalLockTime) (event.Subscription, error) {

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionUpdateBidWithdrawalLockTime)
				if err := _FantomAuction.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseUpdateBidWithdrawalLockTime(log types.Log) (*FantomAuctionUpdateBidWithdrawalLockTime, error) {
	event := new(FantomAuctionUpdateBidWithdrawalLockTime)
	if err := _FantomAuction.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionUpdateMinBidIncrementIterator is returned from FilterUpdateMinBidIncrement and is used to iterate over the raw logs and unpacked data for UpdateMinBidIncrement events raised by the FantomAuction contract.
type FantomAuctionUpdateMinBidIncrementIterator struct {
	Event *FantomAuctionUpdateMinBidIncrement // Event containing the contract specifics and raw log

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
func (it *FantomAuctionUpdateMinBidIncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionUpdateMinBidIncrement)
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
		it.Event = new(FantomAuctionUpdateMinBidIncrement)
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
func (it *FantomAuctionUpdateMinBidIncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionUpdateMinBidIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionUpdateMinBidIncrement represents a UpdateMinBidIncrement event raised by the FantomAuction contract.
type FantomAuctionUpdateMinBidIncrement struct {
	MinBidIncrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinBidIncrement is a free log retrieval operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_FantomAuction *FantomAuctionFilterer) FilterUpdateMinBidIncrement(opts *bind.FilterOpts) (*FantomAuctionUpdateMinBidIncrementIterator, error) {

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionUpdateMinBidIncrementIterator{contract: _FantomAuction.contract, event: "UpdateMinBidIncrement", logs: logs, sub: sub}, nil
}

// WatchUpdateMinBidIncrement is a free log subscription operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_FantomAuction *FantomAuctionFilterer) WatchUpdateMinBidIncrement(opts *bind.WatchOpts, sink chan<- *FantomAuctionUpdateMinBidIncrement) (event.Subscription, error) {

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionUpdateMinBidIncrement)
				if err := _FantomAuction.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseUpdateMinBidIncrement(log types.Log) (*FantomAuctionUpdateMinBidIncrement, error) {
	event := new(FantomAuctionUpdateMinBidIncrement)
	if err := _FantomAuction.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomAuction contract.
type FantomAuctionUpdatePlatformFeeIterator struct {
	Event *FantomAuctionUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomAuctionUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionUpdatePlatformFee)
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
		it.Event = new(FantomAuctionUpdatePlatformFee)
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
func (it *FantomAuctionUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionUpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomAuction contract.
type FantomAuctionUpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomAuction *FantomAuctionFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomAuctionUpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionUpdatePlatformFeeIterator{contract: _FantomAuction.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_FantomAuction *FantomAuctionFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomAuctionUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionUpdatePlatformFee)
				if err := _FantomAuction.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseUpdatePlatformFee(log types.Log) (*FantomAuctionUpdatePlatformFee, error) {
	event := new(FantomAuctionUpdatePlatformFee)
	if err := _FantomAuction.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomAuctionUpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the FantomAuction contract.
type FantomAuctionUpdatePlatformFeeRecipientIterator struct {
	Event *FantomAuctionUpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomAuctionUpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAuctionUpdatePlatformFeeRecipient)
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
		it.Event = new(FantomAuctionUpdatePlatformFeeRecipient)
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
func (it *FantomAuctionUpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAuctionUpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAuctionUpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the FantomAuction contract.
type FantomAuctionUpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomAuction *FantomAuctionFilterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*FantomAuctionUpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _FantomAuction.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomAuctionUpdatePlatformFeeRecipientIterator{contract: _FantomAuction.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomAuction *FantomAuctionFilterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomAuctionUpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomAuction.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAuctionUpdatePlatformFeeRecipient)
				if err := _FantomAuction.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_FantomAuction *FantomAuctionFilterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*FantomAuctionUpdatePlatformFeeRecipient, error) {
	event := new(FantomAuctionUpdatePlatformFeeRecipient)
	if err := _FantomAuction.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
