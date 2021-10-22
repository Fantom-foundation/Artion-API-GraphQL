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

// RandomTradeMetaData contains all meta data concerning the RandomTrade contract.
var RandomTradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rngOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PayTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PayTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newDecimals\",\"type\":\"uint8\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"purchaseID\",\"type\":\"bytes32\"}],\"name\":\"PurchaseCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"purchaseID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PurchaseCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"purchaseID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"PurchaseFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradeTokenID\",\"type\":\"bytes32\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"allowPayToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"cancelPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"clearPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_rnd\",\"type\":\"uint256\"}],\"name\":\"consumeRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"denyPayToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentNFT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"next\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"previous\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNFTAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNFTCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getPurchase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRNGOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradeEnds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradeStarts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnitPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnitPriceDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPayTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shift\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_priceDecimals\",\"type\":\"uint8\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setRNGOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RandomTradeABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomTradeMetaData.ABI instead.
var RandomTradeABI = RandomTradeMetaData.ABI

// RandomTrade is an auto generated Go binding around an Ethereum contract.
type RandomTrade struct {
	RandomTradeCaller     // Read-only binding to the contract
	RandomTradeTransactor // Write-only binding to the contract
	RandomTradeFilterer   // Log filterer for contract events
}

// RandomTradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomTradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomTradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomTradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomTradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomTradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomTradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomTradeSession struct {
	Contract     *RandomTrade      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomTradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomTradeCallerSession struct {
	Contract *RandomTradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RandomTradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomTradeTransactorSession struct {
	Contract     *RandomTradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RandomTradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomTradeRaw struct {
	Contract *RandomTrade // Generic contract binding to access the raw methods on
}

// RandomTradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomTradeCallerRaw struct {
	Contract *RandomTradeCaller // Generic read-only contract binding to access the raw methods on
}

// RandomTradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomTradeTransactorRaw struct {
	Contract *RandomTradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomTrade creates a new instance of RandomTrade, bound to a specific deployed contract.
func NewRandomTrade(address common.Address, backend bind.ContractBackend) (*RandomTrade, error) {
	contract, err := bindRandomTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomTrade{RandomTradeCaller: RandomTradeCaller{contract: contract}, RandomTradeTransactor: RandomTradeTransactor{contract: contract}, RandomTradeFilterer: RandomTradeFilterer{contract: contract}}, nil
}

// NewRandomTradeCaller creates a new read-only instance of RandomTrade, bound to a specific deployed contract.
func NewRandomTradeCaller(address common.Address, caller bind.ContractCaller) (*RandomTradeCaller, error) {
	contract, err := bindRandomTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomTradeCaller{contract: contract}, nil
}

// NewRandomTradeTransactor creates a new write-only instance of RandomTrade, bound to a specific deployed contract.
func NewRandomTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomTradeTransactor, error) {
	contract, err := bindRandomTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomTradeTransactor{contract: contract}, nil
}

// NewRandomTradeFilterer creates a new log filterer instance of RandomTrade, bound to a specific deployed contract.
func NewRandomTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomTradeFilterer, error) {
	contract, err := bindRandomTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomTradeFilterer{contract: contract}, nil
}

// bindRandomTrade binds a generic wrapper to an already deployed contract.
func bindRandomTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomTradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomTrade *RandomTradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomTrade.Contract.RandomTradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomTrade *RandomTradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomTrade.Contract.RandomTradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomTrade *RandomTradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomTrade.Contract.RandomTradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomTrade *RandomTradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomTrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomTrade *RandomTradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomTrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomTrade *RandomTradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomTrade.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentNFT is a free data retrieval call binding the contract method 0x777cc588.
//
// Solidity: function getCurrentNFT() view returns(bytes32)
func (_RandomTrade *RandomTradeCaller) GetCurrentNFT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getCurrentNFT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCurrentNFT is a free data retrieval call binding the contract method 0x777cc588.
//
// Solidity: function getCurrentNFT() view returns(bytes32)
func (_RandomTrade *RandomTradeSession) GetCurrentNFT() ([32]byte, error) {
	return _RandomTrade.Contract.GetCurrentNFT(&_RandomTrade.CallOpts)
}

// GetCurrentNFT is a free data retrieval call binding the contract method 0x777cc588.
//
// Solidity: function getCurrentNFT() view returns(bytes32)
func (_RandomTrade *RandomTradeCallerSession) GetCurrentNFT() ([32]byte, error) {
	return _RandomTrade.Contract.GetCurrentNFT(&_RandomTrade.CallOpts)
}

// GetNFT is a free data retrieval call binding the contract method 0x12223582.
//
// Solidity: function getNFT(bytes32 ) view returns(address nftContract, uint256 tokenID, bytes32 next, bytes32 previous)
func (_RandomTrade *RandomTradeCaller) GetNFT(opts *bind.CallOpts, arg0 [32]byte) (struct {
	NftContract common.Address
	TokenID     *big.Int
	Next        [32]byte
	Previous    [32]byte
}, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getNFT", arg0)

	outstruct := new(struct {
		NftContract common.Address
		TokenID     *big.Int
		Next        [32]byte
		Previous    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Next = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Previous = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetNFT is a free data retrieval call binding the contract method 0x12223582.
//
// Solidity: function getNFT(bytes32 ) view returns(address nftContract, uint256 tokenID, bytes32 next, bytes32 previous)
func (_RandomTrade *RandomTradeSession) GetNFT(arg0 [32]byte) (struct {
	NftContract common.Address
	TokenID     *big.Int
	Next        [32]byte
	Previous    [32]byte
}, error) {
	return _RandomTrade.Contract.GetNFT(&_RandomTrade.CallOpts, arg0)
}

// GetNFT is a free data retrieval call binding the contract method 0x12223582.
//
// Solidity: function getNFT(bytes32 ) view returns(address nftContract, uint256 tokenID, bytes32 next, bytes32 previous)
func (_RandomTrade *RandomTradeCallerSession) GetNFT(arg0 [32]byte) (struct {
	NftContract common.Address
	TokenID     *big.Int
	Next        [32]byte
	Previous    [32]byte
}, error) {
	return _RandomTrade.Contract.GetNFT(&_RandomTrade.CallOpts, arg0)
}

// GetNFTAvailable is a free data retrieval call binding the contract method 0x80e5fec0.
//
// Solidity: function getNFTAvailable() view returns(uint256)
func (_RandomTrade *RandomTradeCaller) GetNFTAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getNFTAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTAvailable is a free data retrieval call binding the contract method 0x80e5fec0.
//
// Solidity: function getNFTAvailable() view returns(uint256)
func (_RandomTrade *RandomTradeSession) GetNFTAvailable() (*big.Int, error) {
	return _RandomTrade.Contract.GetNFTAvailable(&_RandomTrade.CallOpts)
}

// GetNFTAvailable is a free data retrieval call binding the contract method 0x80e5fec0.
//
// Solidity: function getNFTAvailable() view returns(uint256)
func (_RandomTrade *RandomTradeCallerSession) GetNFTAvailable() (*big.Int, error) {
	return _RandomTrade.Contract.GetNFTAvailable(&_RandomTrade.CallOpts)
}

// GetNFTCount is a free data retrieval call binding the contract method 0x54279637.
//
// Solidity: function getNFTCount() view returns(uint256)
func (_RandomTrade *RandomTradeCaller) GetNFTCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getNFTCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTCount is a free data retrieval call binding the contract method 0x54279637.
//
// Solidity: function getNFTCount() view returns(uint256)
func (_RandomTrade *RandomTradeSession) GetNFTCount() (*big.Int, error) {
	return _RandomTrade.Contract.GetNFTCount(&_RandomTrade.CallOpts)
}

// GetNFTCount is a free data retrieval call binding the contract method 0x54279637.
//
// Solidity: function getNFTCount() view returns(uint256)
func (_RandomTrade *RandomTradeCallerSession) GetNFTCount() (*big.Int, error) {
	return _RandomTrade.Contract.GetNFTCount(&_RandomTrade.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_RandomTrade *RandomTradeCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_RandomTrade *RandomTradeSession) GetOwner() (common.Address, error) {
	return _RandomTrade.Contract.GetOwner(&_RandomTrade.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_RandomTrade *RandomTradeCallerSession) GetOwner() (common.Address, error) {
	return _RandomTrade.Contract.GetOwner(&_RandomTrade.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_RandomTrade *RandomTradeCaller) GetPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_RandomTrade *RandomTradeSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _RandomTrade.Contract.GetPrice(&_RandomTrade.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_RandomTrade *RandomTradeCallerSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _RandomTrade.Contract.GetPrice(&_RandomTrade.CallOpts, _token)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_RandomTrade *RandomTradeCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_RandomTrade *RandomTradeSession) GetPriceOracle() (common.Address, error) {
	return _RandomTrade.Contract.GetPriceOracle(&_RandomTrade.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_RandomTrade *RandomTradeCallerSession) GetPriceOracle() (common.Address, error) {
	return _RandomTrade.Contract.GetPriceOracle(&_RandomTrade.CallOpts)
}

// GetPurchase is a free data retrieval call binding the contract method 0x3082a402.
//
// Solidity: function getPurchase(bytes32 ) view returns(address buyer, address payToken, uint256 deposit)
func (_RandomTrade *RandomTradeCaller) GetPurchase(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Buyer    common.Address
	PayToken common.Address
	Deposit  *big.Int
}, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getPurchase", arg0)

	outstruct := new(struct {
		Buyer    common.Address
		PayToken common.Address
		Deposit  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPurchase is a free data retrieval call binding the contract method 0x3082a402.
//
// Solidity: function getPurchase(bytes32 ) view returns(address buyer, address payToken, uint256 deposit)
func (_RandomTrade *RandomTradeSession) GetPurchase(arg0 [32]byte) (struct {
	Buyer    common.Address
	PayToken common.Address
	Deposit  *big.Int
}, error) {
	return _RandomTrade.Contract.GetPurchase(&_RandomTrade.CallOpts, arg0)
}

// GetPurchase is a free data retrieval call binding the contract method 0x3082a402.
//
// Solidity: function getPurchase(bytes32 ) view returns(address buyer, address payToken, uint256 deposit)
func (_RandomTrade *RandomTradeCallerSession) GetPurchase(arg0 [32]byte) (struct {
	Buyer    common.Address
	PayToken common.Address
	Deposit  *big.Int
}, error) {
	return _RandomTrade.Contract.GetPurchase(&_RandomTrade.CallOpts, arg0)
}

// GetRNGOracle is a free data retrieval call binding the contract method 0x36521647.
//
// Solidity: function getRNGOracle() view returns(address)
func (_RandomTrade *RandomTradeCaller) GetRNGOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getRNGOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRNGOracle is a free data retrieval call binding the contract method 0x36521647.
//
// Solidity: function getRNGOracle() view returns(address)
func (_RandomTrade *RandomTradeSession) GetRNGOracle() (common.Address, error) {
	return _RandomTrade.Contract.GetRNGOracle(&_RandomTrade.CallOpts)
}

// GetRNGOracle is a free data retrieval call binding the contract method 0x36521647.
//
// Solidity: function getRNGOracle() view returns(address)
func (_RandomTrade *RandomTradeCallerSession) GetRNGOracle() (common.Address, error) {
	return _RandomTrade.Contract.GetRNGOracle(&_RandomTrade.CallOpts)
}

// GetTradeEnds is a free data retrieval call binding the contract method 0x325408ae.
//
// Solidity: function getTradeEnds() view returns(uint256)
func (_RandomTrade *RandomTradeCaller) GetTradeEnds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getTradeEnds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeEnds is a free data retrieval call binding the contract method 0x325408ae.
//
// Solidity: function getTradeEnds() view returns(uint256)
func (_RandomTrade *RandomTradeSession) GetTradeEnds() (*big.Int, error) {
	return _RandomTrade.Contract.GetTradeEnds(&_RandomTrade.CallOpts)
}

// GetTradeEnds is a free data retrieval call binding the contract method 0x325408ae.
//
// Solidity: function getTradeEnds() view returns(uint256)
func (_RandomTrade *RandomTradeCallerSession) GetTradeEnds() (*big.Int, error) {
	return _RandomTrade.Contract.GetTradeEnds(&_RandomTrade.CallOpts)
}

// GetTradeStarts is a free data retrieval call binding the contract method 0x433fcd91.
//
// Solidity: function getTradeStarts() view returns(uint256)
func (_RandomTrade *RandomTradeCaller) GetTradeStarts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getTradeStarts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeStarts is a free data retrieval call binding the contract method 0x433fcd91.
//
// Solidity: function getTradeStarts() view returns(uint256)
func (_RandomTrade *RandomTradeSession) GetTradeStarts() (*big.Int, error) {
	return _RandomTrade.Contract.GetTradeStarts(&_RandomTrade.CallOpts)
}

// GetTradeStarts is a free data retrieval call binding the contract method 0x433fcd91.
//
// Solidity: function getTradeStarts() view returns(uint256)
func (_RandomTrade *RandomTradeCallerSession) GetTradeStarts() (*big.Int, error) {
	return _RandomTrade.Contract.GetTradeStarts(&_RandomTrade.CallOpts)
}

// GetUnitPrice is a free data retrieval call binding the contract method 0x012a33aa.
//
// Solidity: function getUnitPrice() view returns(uint256)
func (_RandomTrade *RandomTradeCaller) GetUnitPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getUnitPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnitPrice is a free data retrieval call binding the contract method 0x012a33aa.
//
// Solidity: function getUnitPrice() view returns(uint256)
func (_RandomTrade *RandomTradeSession) GetUnitPrice() (*big.Int, error) {
	return _RandomTrade.Contract.GetUnitPrice(&_RandomTrade.CallOpts)
}

// GetUnitPrice is a free data retrieval call binding the contract method 0x012a33aa.
//
// Solidity: function getUnitPrice() view returns(uint256)
func (_RandomTrade *RandomTradeCallerSession) GetUnitPrice() (*big.Int, error) {
	return _RandomTrade.Contract.GetUnitPrice(&_RandomTrade.CallOpts)
}

// GetUnitPriceDecimals is a free data retrieval call binding the contract method 0x29bed128.
//
// Solidity: function getUnitPriceDecimals() view returns(uint8)
func (_RandomTrade *RandomTradeCaller) GetUnitPriceDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "getUnitPriceDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUnitPriceDecimals is a free data retrieval call binding the contract method 0x29bed128.
//
// Solidity: function getUnitPriceDecimals() view returns(uint8)
func (_RandomTrade *RandomTradeSession) GetUnitPriceDecimals() (uint8, error) {
	return _RandomTrade.Contract.GetUnitPriceDecimals(&_RandomTrade.CallOpts)
}

// GetUnitPriceDecimals is a free data retrieval call binding the contract method 0x29bed128.
//
// Solidity: function getUnitPriceDecimals() view returns(uint8)
func (_RandomTrade *RandomTradeCallerSession) GetUnitPriceDecimals() (uint8, error) {
	return _RandomTrade.Contract.GetUnitPriceDecimals(&_RandomTrade.CallOpts)
}

// IsPayTokenAllowed is a free data retrieval call binding the contract method 0x07bfa7a6.
//
// Solidity: function isPayTokenAllowed(address ) view returns(bool)
func (_RandomTrade *RandomTradeCaller) IsPayTokenAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RandomTrade.contract.Call(opts, &out, "isPayTokenAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPayTokenAllowed is a free data retrieval call binding the contract method 0x07bfa7a6.
//
// Solidity: function isPayTokenAllowed(address ) view returns(bool)
func (_RandomTrade *RandomTradeSession) IsPayTokenAllowed(arg0 common.Address) (bool, error) {
	return _RandomTrade.Contract.IsPayTokenAllowed(&_RandomTrade.CallOpts, arg0)
}

// IsPayTokenAllowed is a free data retrieval call binding the contract method 0x07bfa7a6.
//
// Solidity: function isPayTokenAllowed(address ) view returns(bool)
func (_RandomTrade *RandomTradeCallerSession) IsPayTokenAllowed(arg0 common.Address) (bool, error) {
	return _RandomTrade.Contract.IsPayTokenAllowed(&_RandomTrade.CallOpts, arg0)
}

// AllowPayToken is a paid mutator transaction binding the contract method 0x9aae1267.
//
// Solidity: function allowPayToken(address _token) returns()
func (_RandomTrade *RandomTradeTransactor) AllowPayToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "allowPayToken", _token)
}

// AllowPayToken is a paid mutator transaction binding the contract method 0x9aae1267.
//
// Solidity: function allowPayToken(address _token) returns()
func (_RandomTrade *RandomTradeSession) AllowPayToken(_token common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.AllowPayToken(&_RandomTrade.TransactOpts, _token)
}

// AllowPayToken is a paid mutator transaction binding the contract method 0x9aae1267.
//
// Solidity: function allowPayToken(address _token) returns()
func (_RandomTrade *RandomTradeTransactorSession) AllowPayToken(_token common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.AllowPayToken(&_RandomTrade.TransactOpts, _token)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0x951f38cf.
//
// Solidity: function cancelPurchase(bytes32 _id) returns()
func (_RandomTrade *RandomTradeTransactor) CancelPurchase(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "cancelPurchase", _id)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0x951f38cf.
//
// Solidity: function cancelPurchase(bytes32 _id) returns()
func (_RandomTrade *RandomTradeSession) CancelPurchase(_id [32]byte) (*types.Transaction, error) {
	return _RandomTrade.Contract.CancelPurchase(&_RandomTrade.TransactOpts, _id)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0x951f38cf.
//
// Solidity: function cancelPurchase(bytes32 _id) returns()
func (_RandomTrade *RandomTradeTransactorSession) CancelPurchase(_id [32]byte) (*types.Transaction, error) {
	return _RandomTrade.Contract.CancelPurchase(&_RandomTrade.TransactOpts, _id)
}

// ClearPurchase is a paid mutator transaction binding the contract method 0x59704d31.
//
// Solidity: function clearPurchase(bytes32 _id) returns()
func (_RandomTrade *RandomTradeTransactor) ClearPurchase(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "clearPurchase", _id)
}

// ClearPurchase is a paid mutator transaction binding the contract method 0x59704d31.
//
// Solidity: function clearPurchase(bytes32 _id) returns()
func (_RandomTrade *RandomTradeSession) ClearPurchase(_id [32]byte) (*types.Transaction, error) {
	return _RandomTrade.Contract.ClearPurchase(&_RandomTrade.TransactOpts, _id)
}

// ClearPurchase is a paid mutator transaction binding the contract method 0x59704d31.
//
// Solidity: function clearPurchase(bytes32 _id) returns()
func (_RandomTrade *RandomTradeTransactorSession) ClearPurchase(_id [32]byte) (*types.Transaction, error) {
	return _RandomTrade.Contract.ClearPurchase(&_RandomTrade.TransactOpts, _id)
}

// ConsumeRandomNumber is a paid mutator transaction binding the contract method 0xa3f03687.
//
// Solidity: function consumeRandomNumber(bytes32 _seed, uint256 _rnd) returns()
func (_RandomTrade *RandomTradeTransactor) ConsumeRandomNumber(opts *bind.TransactOpts, _seed [32]byte, _rnd *big.Int) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "consumeRandomNumber", _seed, _rnd)
}

// ConsumeRandomNumber is a paid mutator transaction binding the contract method 0xa3f03687.
//
// Solidity: function consumeRandomNumber(bytes32 _seed, uint256 _rnd) returns()
func (_RandomTrade *RandomTradeSession) ConsumeRandomNumber(_seed [32]byte, _rnd *big.Int) (*types.Transaction, error) {
	return _RandomTrade.Contract.ConsumeRandomNumber(&_RandomTrade.TransactOpts, _seed, _rnd)
}

// ConsumeRandomNumber is a paid mutator transaction binding the contract method 0xa3f03687.
//
// Solidity: function consumeRandomNumber(bytes32 _seed, uint256 _rnd) returns()
func (_RandomTrade *RandomTradeTransactorSession) ConsumeRandomNumber(_seed [32]byte, _rnd *big.Int) (*types.Transaction, error) {
	return _RandomTrade.Contract.ConsumeRandomNumber(&_RandomTrade.TransactOpts, _seed, _rnd)
}

// DenyPayToken is a paid mutator transaction binding the contract method 0x8c49f478.
//
// Solidity: function denyPayToken(address _token) returns()
func (_RandomTrade *RandomTradeTransactor) DenyPayToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "denyPayToken", _token)
}

// DenyPayToken is a paid mutator transaction binding the contract method 0x8c49f478.
//
// Solidity: function denyPayToken(address _token) returns()
func (_RandomTrade *RandomTradeSession) DenyPayToken(_token common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.DenyPayToken(&_RandomTrade.TransactOpts, _token)
}

// DenyPayToken is a paid mutator transaction binding the contract method 0x8c49f478.
//
// Solidity: function denyPayToken(address _token) returns()
func (_RandomTrade *RandomTradeTransactorSession) DenyPayToken(_token common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.DenyPayToken(&_RandomTrade.TransactOpts, _token)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes ) returns(bytes4)
func (_RandomTrade *RandomTradeTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "onERC721Received", operator, arg1, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes ) returns(bytes4)
func (_RandomTrade *RandomTradeSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _RandomTrade.Contract.OnERC721Received(&_RandomTrade.TransactOpts, operator, arg1, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes ) returns(bytes4)
func (_RandomTrade *RandomTradeTransactorSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _RandomTrade.Contract.OnERC721Received(&_RandomTrade.TransactOpts, operator, arg1, tokenId, arg3)
}

// Purchase is a paid mutator transaction binding the contract method 0x25b31a97.
//
// Solidity: function purchase(address _token) returns()
func (_RandomTrade *RandomTradeTransactor) Purchase(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "purchase", _token)
}

// Purchase is a paid mutator transaction binding the contract method 0x25b31a97.
//
// Solidity: function purchase(address _token) returns()
func (_RandomTrade *RandomTradeSession) Purchase(_token common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.Purchase(&_RandomTrade.TransactOpts, _token)
}

// Purchase is a paid mutator transaction binding the contract method 0x25b31a97.
//
// Solidity: function purchase(address _token) returns()
func (_RandomTrade *RandomTradeTransactorSession) Purchase(_token common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.Purchase(&_RandomTrade.TransactOpts, _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x36c5d724.
//
// Solidity: function removeToken(uint256 _shift) returns()
func (_RandomTrade *RandomTradeTransactor) RemoveToken(opts *bind.TransactOpts, _shift *big.Int) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "removeToken", _shift)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x36c5d724.
//
// Solidity: function removeToken(uint256 _shift) returns()
func (_RandomTrade *RandomTradeSession) RemoveToken(_shift *big.Int) (*types.Transaction, error) {
	return _RandomTrade.Contract.RemoveToken(&_RandomTrade.TransactOpts, _shift)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x36c5d724.
//
// Solidity: function removeToken(uint256 _shift) returns()
func (_RandomTrade *RandomTradeTransactorSession) RemoveToken(_shift *big.Int) (*types.Transaction, error) {
	return _RandomTrade.Contract.RemoveToken(&_RandomTrade.TransactOpts, _shift)
}

// SetPrice is a paid mutator transaction binding the contract method 0x96383c04.
//
// Solidity: function setPrice(uint256 _price, uint8 _priceDecimals) returns()
func (_RandomTrade *RandomTradeTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int, _priceDecimals uint8) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "setPrice", _price, _priceDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x96383c04.
//
// Solidity: function setPrice(uint256 _price, uint8 _priceDecimals) returns()
func (_RandomTrade *RandomTradeSession) SetPrice(_price *big.Int, _priceDecimals uint8) (*types.Transaction, error) {
	return _RandomTrade.Contract.SetPrice(&_RandomTrade.TransactOpts, _price, _priceDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x96383c04.
//
// Solidity: function setPrice(uint256 _price, uint8 _priceDecimals) returns()
func (_RandomTrade *RandomTradeTransactorSession) SetPrice(_price *big.Int, _priceDecimals uint8) (*types.Transaction, error) {
	return _RandomTrade.Contract.SetPrice(&_RandomTrade.TransactOpts, _price, _priceDecimals)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _oracle) returns()
func (_RandomTrade *RandomTradeTransactor) SetPriceOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "setPriceOracle", _oracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _oracle) returns()
func (_RandomTrade *RandomTradeSession) SetPriceOracle(_oracle common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.SetPriceOracle(&_RandomTrade.TransactOpts, _oracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _oracle) returns()
func (_RandomTrade *RandomTradeTransactorSession) SetPriceOracle(_oracle common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.SetPriceOracle(&_RandomTrade.TransactOpts, _oracle)
}

// SetRNGOracle is a paid mutator transaction binding the contract method 0xb0d60178.
//
// Solidity: function setRNGOracle(address _oracle) returns()
func (_RandomTrade *RandomTradeTransactor) SetRNGOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "setRNGOracle", _oracle)
}

// SetRNGOracle is a paid mutator transaction binding the contract method 0xb0d60178.
//
// Solidity: function setRNGOracle(address _oracle) returns()
func (_RandomTrade *RandomTradeSession) SetRNGOracle(_oracle common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.SetRNGOracle(&_RandomTrade.TransactOpts, _oracle)
}

// SetRNGOracle is a paid mutator transaction binding the contract method 0xb0d60178.
//
// Solidity: function setRNGOracle(address _oracle) returns()
func (_RandomTrade *RandomTradeTransactorSession) SetRNGOracle(_oracle common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.SetRNGOracle(&_RandomTrade.TransactOpts, _oracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RandomTrade *RandomTradeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RandomTrade.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RandomTrade *RandomTradeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.TransferOwnership(&_RandomTrade.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RandomTrade *RandomTradeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RandomTrade.Contract.TransferOwnership(&_RandomTrade.TransactOpts, _newOwner)
}

// RandomTradeOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the RandomTrade contract.
type RandomTradeOwnerChangedIterator struct {
	Event *RandomTradeOwnerChanged // Event containing the contract specifics and raw log

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
func (it *RandomTradeOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradeOwnerChanged)
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
		it.Event = new(RandomTradeOwnerChanged)
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
func (it *RandomTradeOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradeOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradeOwnerChanged represents a OwnerChanged event raised by the RandomTrade contract.
type RandomTradeOwnerChanged struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address newOwner)
func (_RandomTrade *RandomTradeFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*RandomTradeOwnerChangedIterator, error) {

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &RandomTradeOwnerChangedIterator{contract: _RandomTrade.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address newOwner)
func (_RandomTrade *RandomTradeFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *RandomTradeOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradeOwnerChanged)
				if err := _RandomTrade.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address newOwner)
func (_RandomTrade *RandomTradeFilterer) ParseOwnerChanged(log types.Log) (*RandomTradeOwnerChanged, error) {
	event := new(RandomTradeOwnerChanged)
	if err := _RandomTrade.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomTradePayTokenAddedIterator is returned from FilterPayTokenAdded and is used to iterate over the raw logs and unpacked data for PayTokenAdded events raised by the RandomTrade contract.
type RandomTradePayTokenAddedIterator struct {
	Event *RandomTradePayTokenAdded // Event containing the contract specifics and raw log

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
func (it *RandomTradePayTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradePayTokenAdded)
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
		it.Event = new(RandomTradePayTokenAdded)
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
func (it *RandomTradePayTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradePayTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradePayTokenAdded represents a PayTokenAdded event raised by the RandomTrade contract.
type RandomTradePayTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPayTokenAdded is a free log retrieval operation binding the contract event 0x7885671185cd1b30280a590458ac1fd9c40ec6fb6013288a46a00dd5b936ffee.
//
// Solidity: event PayTokenAdded(address token)
func (_RandomTrade *RandomTradeFilterer) FilterPayTokenAdded(opts *bind.FilterOpts) (*RandomTradePayTokenAddedIterator, error) {

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "PayTokenAdded")
	if err != nil {
		return nil, err
	}
	return &RandomTradePayTokenAddedIterator{contract: _RandomTrade.contract, event: "PayTokenAdded", logs: logs, sub: sub}, nil
}

// WatchPayTokenAdded is a free log subscription operation binding the contract event 0x7885671185cd1b30280a590458ac1fd9c40ec6fb6013288a46a00dd5b936ffee.
//
// Solidity: event PayTokenAdded(address token)
func (_RandomTrade *RandomTradeFilterer) WatchPayTokenAdded(opts *bind.WatchOpts, sink chan<- *RandomTradePayTokenAdded) (event.Subscription, error) {

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "PayTokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradePayTokenAdded)
				if err := _RandomTrade.contract.UnpackLog(event, "PayTokenAdded", log); err != nil {
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

// ParsePayTokenAdded is a log parse operation binding the contract event 0x7885671185cd1b30280a590458ac1fd9c40ec6fb6013288a46a00dd5b936ffee.
//
// Solidity: event PayTokenAdded(address token)
func (_RandomTrade *RandomTradeFilterer) ParsePayTokenAdded(log types.Log) (*RandomTradePayTokenAdded, error) {
	event := new(RandomTradePayTokenAdded)
	if err := _RandomTrade.contract.UnpackLog(event, "PayTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomTradePayTokenRemovedIterator is returned from FilterPayTokenRemoved and is used to iterate over the raw logs and unpacked data for PayTokenRemoved events raised by the RandomTrade contract.
type RandomTradePayTokenRemovedIterator struct {
	Event *RandomTradePayTokenRemoved // Event containing the contract specifics and raw log

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
func (it *RandomTradePayTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradePayTokenRemoved)
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
		it.Event = new(RandomTradePayTokenRemoved)
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
func (it *RandomTradePayTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradePayTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradePayTokenRemoved represents a PayTokenRemoved event raised by the RandomTrade contract.
type RandomTradePayTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPayTokenRemoved is a free log retrieval operation binding the contract event 0xf36e518b3a6b4afd64732dcf603fb05c0a7dc30d66044808e00bf257d3525e62.
//
// Solidity: event PayTokenRemoved(address token)
func (_RandomTrade *RandomTradeFilterer) FilterPayTokenRemoved(opts *bind.FilterOpts) (*RandomTradePayTokenRemovedIterator, error) {

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "PayTokenRemoved")
	if err != nil {
		return nil, err
	}
	return &RandomTradePayTokenRemovedIterator{contract: _RandomTrade.contract, event: "PayTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchPayTokenRemoved is a free log subscription operation binding the contract event 0xf36e518b3a6b4afd64732dcf603fb05c0a7dc30d66044808e00bf257d3525e62.
//
// Solidity: event PayTokenRemoved(address token)
func (_RandomTrade *RandomTradeFilterer) WatchPayTokenRemoved(opts *bind.WatchOpts, sink chan<- *RandomTradePayTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "PayTokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradePayTokenRemoved)
				if err := _RandomTrade.contract.UnpackLog(event, "PayTokenRemoved", log); err != nil {
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

// ParsePayTokenRemoved is a log parse operation binding the contract event 0xf36e518b3a6b4afd64732dcf603fb05c0a7dc30d66044808e00bf257d3525e62.
//
// Solidity: event PayTokenRemoved(address token)
func (_RandomTrade *RandomTradeFilterer) ParsePayTokenRemoved(log types.Log) (*RandomTradePayTokenRemoved, error) {
	event := new(RandomTradePayTokenRemoved)
	if err := _RandomTrade.contract.UnpackLog(event, "PayTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomTradePriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the RandomTrade contract.
type RandomTradePriceChangedIterator struct {
	Event *RandomTradePriceChanged // Event containing the contract specifics and raw log

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
func (it *RandomTradePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradePriceChanged)
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
		it.Event = new(RandomTradePriceChanged)
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
func (it *RandomTradePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradePriceChanged represents a PriceChanged event raised by the RandomTrade contract.
type RandomTradePriceChanged struct {
	NewPrice    *big.Int
	NewDecimals uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0xc0dde471040c20f58bf65c262285b2d1c3b6d8d99c67eb8b4c6d5c3545c5d61e.
//
// Solidity: event PriceChanged(uint256 newPrice, uint8 newDecimals)
func (_RandomTrade *RandomTradeFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*RandomTradePriceChangedIterator, error) {

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &RandomTradePriceChangedIterator{contract: _RandomTrade.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0xc0dde471040c20f58bf65c262285b2d1c3b6d8d99c67eb8b4c6d5c3545c5d61e.
//
// Solidity: event PriceChanged(uint256 newPrice, uint8 newDecimals)
func (_RandomTrade *RandomTradeFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *RandomTradePriceChanged) (event.Subscription, error) {

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradePriceChanged)
				if err := _RandomTrade.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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

// ParsePriceChanged is a log parse operation binding the contract event 0xc0dde471040c20f58bf65c262285b2d1c3b6d8d99c67eb8b4c6d5c3545c5d61e.
//
// Solidity: event PriceChanged(uint256 newPrice, uint8 newDecimals)
func (_RandomTrade *RandomTradeFilterer) ParsePriceChanged(log types.Log) (*RandomTradePriceChanged, error) {
	event := new(RandomTradePriceChanged)
	if err := _RandomTrade.contract.UnpackLog(event, "PriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomTradePurchaseCanceledIterator is returned from FilterPurchaseCanceled and is used to iterate over the raw logs and unpacked data for PurchaseCanceled events raised by the RandomTrade contract.
type RandomTradePurchaseCanceledIterator struct {
	Event *RandomTradePurchaseCanceled // Event containing the contract specifics and raw log

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
func (it *RandomTradePurchaseCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradePurchaseCanceled)
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
		it.Event = new(RandomTradePurchaseCanceled)
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
func (it *RandomTradePurchaseCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradePurchaseCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradePurchaseCanceled represents a PurchaseCanceled event raised by the RandomTrade contract.
type RandomTradePurchaseCanceled struct {
	Buyer      common.Address
	PurchaseID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPurchaseCanceled is a free log retrieval operation binding the contract event 0x5ac43bc19cb42a697cdf201fb2d91ba45d0962f91c9f2f2dfec6674297fd8410.
//
// Solidity: event PurchaseCanceled(address indexed buyer, bytes32 purchaseID)
func (_RandomTrade *RandomTradeFilterer) FilterPurchaseCanceled(opts *bind.FilterOpts, buyer []common.Address) (*RandomTradePurchaseCanceledIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "PurchaseCanceled", buyerRule)
	if err != nil {
		return nil, err
	}
	return &RandomTradePurchaseCanceledIterator{contract: _RandomTrade.contract, event: "PurchaseCanceled", logs: logs, sub: sub}, nil
}

// WatchPurchaseCanceled is a free log subscription operation binding the contract event 0x5ac43bc19cb42a697cdf201fb2d91ba45d0962f91c9f2f2dfec6674297fd8410.
//
// Solidity: event PurchaseCanceled(address indexed buyer, bytes32 purchaseID)
func (_RandomTrade *RandomTradeFilterer) WatchPurchaseCanceled(opts *bind.WatchOpts, sink chan<- *RandomTradePurchaseCanceled, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "PurchaseCanceled", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradePurchaseCanceled)
				if err := _RandomTrade.contract.UnpackLog(event, "PurchaseCanceled", log); err != nil {
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

// ParsePurchaseCanceled is a log parse operation binding the contract event 0x5ac43bc19cb42a697cdf201fb2d91ba45d0962f91c9f2f2dfec6674297fd8410.
//
// Solidity: event PurchaseCanceled(address indexed buyer, bytes32 purchaseID)
func (_RandomTrade *RandomTradeFilterer) ParsePurchaseCanceled(log types.Log) (*RandomTradePurchaseCanceled, error) {
	event := new(RandomTradePurchaseCanceled)
	if err := _RandomTrade.contract.UnpackLog(event, "PurchaseCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomTradePurchaseCreatedIterator is returned from FilterPurchaseCreated and is used to iterate over the raw logs and unpacked data for PurchaseCreated events raised by the RandomTrade contract.
type RandomTradePurchaseCreatedIterator struct {
	Event *RandomTradePurchaseCreated // Event containing the contract specifics and raw log

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
func (it *RandomTradePurchaseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradePurchaseCreated)
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
		it.Event = new(RandomTradePurchaseCreated)
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
func (it *RandomTradePurchaseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradePurchaseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradePurchaseCreated represents a PurchaseCreated event raised by the RandomTrade contract.
type RandomTradePurchaseCreated struct {
	Buyer      common.Address
	PurchaseID [32]byte
	PayToken   common.Address
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPurchaseCreated is a free log retrieval operation binding the contract event 0x35cd9b60e9c67cd3e02b9f6ad728a195879869b9bb7caf562bce5e1268c8239a.
//
// Solidity: event PurchaseCreated(address indexed buyer, bytes32 purchaseID, address payToken, uint256 price)
func (_RandomTrade *RandomTradeFilterer) FilterPurchaseCreated(opts *bind.FilterOpts, buyer []common.Address) (*RandomTradePurchaseCreatedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "PurchaseCreated", buyerRule)
	if err != nil {
		return nil, err
	}
	return &RandomTradePurchaseCreatedIterator{contract: _RandomTrade.contract, event: "PurchaseCreated", logs: logs, sub: sub}, nil
}

// WatchPurchaseCreated is a free log subscription operation binding the contract event 0x35cd9b60e9c67cd3e02b9f6ad728a195879869b9bb7caf562bce5e1268c8239a.
//
// Solidity: event PurchaseCreated(address indexed buyer, bytes32 purchaseID, address payToken, uint256 price)
func (_RandomTrade *RandomTradeFilterer) WatchPurchaseCreated(opts *bind.WatchOpts, sink chan<- *RandomTradePurchaseCreated, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "PurchaseCreated", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradePurchaseCreated)
				if err := _RandomTrade.contract.UnpackLog(event, "PurchaseCreated", log); err != nil {
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

// ParsePurchaseCreated is a log parse operation binding the contract event 0x35cd9b60e9c67cd3e02b9f6ad728a195879869b9bb7caf562bce5e1268c8239a.
//
// Solidity: event PurchaseCreated(address indexed buyer, bytes32 purchaseID, address payToken, uint256 price)
func (_RandomTrade *RandomTradeFilterer) ParsePurchaseCreated(log types.Log) (*RandomTradePurchaseCreated, error) {
	event := new(RandomTradePurchaseCreated)
	if err := _RandomTrade.contract.UnpackLog(event, "PurchaseCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomTradePurchaseFinishedIterator is returned from FilterPurchaseFinished and is used to iterate over the raw logs and unpacked data for PurchaseFinished events raised by the RandomTrade contract.
type RandomTradePurchaseFinishedIterator struct {
	Event *RandomTradePurchaseFinished // Event containing the contract specifics and raw log

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
func (it *RandomTradePurchaseFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradePurchaseFinished)
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
		it.Event = new(RandomTradePurchaseFinished)
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
func (it *RandomTradePurchaseFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradePurchaseFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradePurchaseFinished represents a PurchaseFinished event raised by the RandomTrade contract.
type RandomTradePurchaseFinished struct {
	Buyer      common.Address
	PurchaseID [32]byte
	Collection common.Address
	TokenID    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPurchaseFinished is a free log retrieval operation binding the contract event 0xc4992cca08aa5af71c2ae61e6a9e727cecf15ea1e5e258eae0fe5cac5e02a917.
//
// Solidity: event PurchaseFinished(address indexed buyer, bytes32 purchaseID, address collection, uint256 tokenID)
func (_RandomTrade *RandomTradeFilterer) FilterPurchaseFinished(opts *bind.FilterOpts, buyer []common.Address) (*RandomTradePurchaseFinishedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "PurchaseFinished", buyerRule)
	if err != nil {
		return nil, err
	}
	return &RandomTradePurchaseFinishedIterator{contract: _RandomTrade.contract, event: "PurchaseFinished", logs: logs, sub: sub}, nil
}

// WatchPurchaseFinished is a free log subscription operation binding the contract event 0xc4992cca08aa5af71c2ae61e6a9e727cecf15ea1e5e258eae0fe5cac5e02a917.
//
// Solidity: event PurchaseFinished(address indexed buyer, bytes32 purchaseID, address collection, uint256 tokenID)
func (_RandomTrade *RandomTradeFilterer) WatchPurchaseFinished(opts *bind.WatchOpts, sink chan<- *RandomTradePurchaseFinished, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "PurchaseFinished", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradePurchaseFinished)
				if err := _RandomTrade.contract.UnpackLog(event, "PurchaseFinished", log); err != nil {
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

// ParsePurchaseFinished is a log parse operation binding the contract event 0xc4992cca08aa5af71c2ae61e6a9e727cecf15ea1e5e258eae0fe5cac5e02a917.
//
// Solidity: event PurchaseFinished(address indexed buyer, bytes32 purchaseID, address collection, uint256 tokenID)
func (_RandomTrade *RandomTradeFilterer) ParsePurchaseFinished(log types.Log) (*RandomTradePurchaseFinished, error) {
	event := new(RandomTradePurchaseFinished)
	if err := _RandomTrade.contract.UnpackLog(event, "PurchaseFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomTradeTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the RandomTrade contract.
type RandomTradeTokenAddedIterator struct {
	Event *RandomTradeTokenAdded // Event containing the contract specifics and raw log

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
func (it *RandomTradeTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomTradeTokenAdded)
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
		it.Event = new(RandomTradeTokenAdded)
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
func (it *RandomTradeTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomTradeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomTradeTokenAdded represents a TokenAdded event raised by the RandomTrade contract.
type RandomTradeTokenAdded struct {
	Collection   common.Address
	TokenID      *big.Int
	TradeTokenID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x55da7a6a7d3b708ae471975643f183a0b160d485625acbbc628e25f2516304b7.
//
// Solidity: event TokenAdded(address collection, uint256 tokenID, bytes32 tradeTokenID)
func (_RandomTrade *RandomTradeFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*RandomTradeTokenAddedIterator, error) {

	logs, sub, err := _RandomTrade.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &RandomTradeTokenAddedIterator{contract: _RandomTrade.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x55da7a6a7d3b708ae471975643f183a0b160d485625acbbc628e25f2516304b7.
//
// Solidity: event TokenAdded(address collection, uint256 tokenID, bytes32 tradeTokenID)
func (_RandomTrade *RandomTradeFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *RandomTradeTokenAdded) (event.Subscription, error) {

	logs, sub, err := _RandomTrade.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomTradeTokenAdded)
				if err := _RandomTrade.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x55da7a6a7d3b708ae471975643f183a0b160d485625acbbc628e25f2516304b7.
//
// Solidity: event TokenAdded(address collection, uint256 tokenID, bytes32 tradeTokenID)
func (_RandomTrade *RandomTradeFilterer) ParseTokenAdded(log types.Log) (*RandomTradeTokenAdded, error) {
	event := new(RandomTradeTokenAdded)
	if err := _RandomTrade.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
