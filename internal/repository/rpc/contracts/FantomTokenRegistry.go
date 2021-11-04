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

// FantomTokenRegistryMetaData contains all meta data concerning the FantomTokenRegistry contract.
var FantomTokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FantomTokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomTokenRegistryMetaData.ABI instead.
var FantomTokenRegistryABI = FantomTokenRegistryMetaData.ABI

// FantomTokenRegistry is an auto generated Go binding around an Ethereum contract.
type FantomTokenRegistry struct {
	FantomTokenRegistryCaller     // Read-only binding to the contract
	FantomTokenRegistryTransactor // Write-only binding to the contract
	FantomTokenRegistryFilterer   // Log filterer for contract events
}

// FantomTokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomTokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomTokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomTokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomTokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomTokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomTokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomTokenRegistrySession struct {
	Contract     *FantomTokenRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FantomTokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomTokenRegistryCallerSession struct {
	Contract *FantomTokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FantomTokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomTokenRegistryTransactorSession struct {
	Contract     *FantomTokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FantomTokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomTokenRegistryRaw struct {
	Contract *FantomTokenRegistry // Generic contract binding to access the raw methods on
}

// FantomTokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomTokenRegistryCallerRaw struct {
	Contract *FantomTokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FantomTokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomTokenRegistryTransactorRaw struct {
	Contract *FantomTokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomTokenRegistry creates a new instance of FantomTokenRegistry, bound to a specific deployed contract.
func NewFantomTokenRegistry(address common.Address, backend bind.ContractBackend) (*FantomTokenRegistry, error) {
	contract, err := bindFantomTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomTokenRegistry{FantomTokenRegistryCaller: FantomTokenRegistryCaller{contract: contract}, FantomTokenRegistryTransactor: FantomTokenRegistryTransactor{contract: contract}, FantomTokenRegistryFilterer: FantomTokenRegistryFilterer{contract: contract}}, nil
}

// NewFantomTokenRegistryCaller creates a new read-only instance of FantomTokenRegistry, bound to a specific deployed contract.
func NewFantomTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*FantomTokenRegistryCaller, error) {
	contract, err := bindFantomTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomTokenRegistryCaller{contract: contract}, nil
}

// NewFantomTokenRegistryTransactor creates a new write-only instance of FantomTokenRegistry, bound to a specific deployed contract.
func NewFantomTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomTokenRegistryTransactor, error) {
	contract, err := bindFantomTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomTokenRegistryTransactor{contract: contract}, nil
}

// NewFantomTokenRegistryFilterer creates a new log filterer instance of FantomTokenRegistry, bound to a specific deployed contract.
func NewFantomTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomTokenRegistryFilterer, error) {
	contract, err := bindFantomTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomTokenRegistryFilterer{contract: contract}, nil
}

// bindFantomTokenRegistry binds a generic wrapper to an already deployed contract.
func bindFantomTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomTokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomTokenRegistry *FantomTokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomTokenRegistry.Contract.FantomTokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomTokenRegistry *FantomTokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.FantomTokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomTokenRegistry *FantomTokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.FantomTokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomTokenRegistry *FantomTokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomTokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomTokenRegistry *FantomTokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomTokenRegistry *FantomTokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// Enabled is a free data retrieval call binding the contract method 0x98afdfe3.
//
// Solidity: function enabled(address ) view returns(bool)
func (_FantomTokenRegistry *FantomTokenRegistryCaller) Enabled(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FantomTokenRegistry.contract.Call(opts, &out, "enabled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enabled is a free data retrieval call binding the contract method 0x98afdfe3.
//
// Solidity: function enabled(address ) view returns(bool)
func (_FantomTokenRegistry *FantomTokenRegistrySession) Enabled(arg0 common.Address) (bool, error) {
	return _FantomTokenRegistry.Contract.Enabled(&_FantomTokenRegistry.CallOpts, arg0)
}

// Enabled is a free data retrieval call binding the contract method 0x98afdfe3.
//
// Solidity: function enabled(address ) view returns(bool)
func (_FantomTokenRegistry *FantomTokenRegistryCallerSession) Enabled(arg0 common.Address) (bool, error) {
	return _FantomTokenRegistry.Contract.Enabled(&_FantomTokenRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomTokenRegistry *FantomTokenRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomTokenRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomTokenRegistry *FantomTokenRegistrySession) Owner() (common.Address, error) {
	return _FantomTokenRegistry.Contract.Owner(&_FantomTokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomTokenRegistry *FantomTokenRegistryCallerSession) Owner() (common.Address, error) {
	return _FantomTokenRegistry.Contract.Owner(&_FantomTokenRegistry.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address token) returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactor) Add(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.contract.Transact(opts, "add", token)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address token) returns()
func (_FantomTokenRegistry *FantomTokenRegistrySession) Add(token common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.Add(&_FantomTokenRegistry.TransactOpts, token)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address token) returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactorSession) Add(token common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.Add(&_FantomTokenRegistry.TransactOpts, token)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address token) returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactor) Remove(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.contract.Transact(opts, "remove", token)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address token) returns()
func (_FantomTokenRegistry *FantomTokenRegistrySession) Remove(token common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.Remove(&_FantomTokenRegistry.TransactOpts, token)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address token) returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactorSession) Remove(token common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.Remove(&_FantomTokenRegistry.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomTokenRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomTokenRegistry *FantomTokenRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.RenounceOwnership(&_FantomTokenRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.RenounceOwnership(&_FantomTokenRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomTokenRegistry *FantomTokenRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.TransferOwnership(&_FantomTokenRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomTokenRegistry *FantomTokenRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomTokenRegistry.Contract.TransferOwnership(&_FantomTokenRegistry.TransactOpts, newOwner)
}

// FantomTokenRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomTokenRegistry contract.
type FantomTokenRegistryOwnershipTransferredIterator struct {
	Event *FantomTokenRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomTokenRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomTokenRegistryOwnershipTransferred)
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
		it.Event = new(FantomTokenRegistryOwnershipTransferred)
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
func (it *FantomTokenRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomTokenRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomTokenRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the FantomTokenRegistry contract.
type FantomTokenRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomTokenRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomTokenRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomTokenRegistryOwnershipTransferredIterator{contract: _FantomTokenRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomTokenRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomTokenRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomTokenRegistryOwnershipTransferred)
				if err := _FantomTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*FantomTokenRegistryOwnershipTransferred, error) {
	event := new(FantomTokenRegistryOwnershipTransferred)
	if err := _FantomTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomTokenRegistryTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the FantomTokenRegistry contract.
type FantomTokenRegistryTokenAddedIterator struct {
	Event *FantomTokenRegistryTokenAdded // Event containing the contract specifics and raw log

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
func (it *FantomTokenRegistryTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomTokenRegistryTokenAdded)
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
		it.Event = new(FantomTokenRegistryTokenAdded)
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
func (it *FantomTokenRegistryTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomTokenRegistryTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomTokenRegistryTokenAdded represents a TokenAdded event raised by the FantomTokenRegistry contract.
type FantomTokenRegistryTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address token)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*FantomTokenRegistryTokenAddedIterator, error) {

	logs, sub, err := _FantomTokenRegistry.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &FantomTokenRegistryTokenAddedIterator{contract: _FantomTokenRegistry.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address token)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *FantomTokenRegistryTokenAdded) (event.Subscription, error) {

	logs, sub, err := _FantomTokenRegistry.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomTokenRegistryTokenAdded)
				if err := _FantomTokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address token)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) ParseTokenAdded(log types.Log) (*FantomTokenRegistryTokenAdded, error) {
	event := new(FantomTokenRegistryTokenAdded)
	if err := _FantomTokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomTokenRegistryTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the FantomTokenRegistry contract.
type FantomTokenRegistryTokenRemovedIterator struct {
	Event *FantomTokenRegistryTokenRemoved // Event containing the contract specifics and raw log

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
func (it *FantomTokenRegistryTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomTokenRegistryTokenRemoved)
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
		it.Event = new(FantomTokenRegistryTokenRemoved)
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
func (it *FantomTokenRegistryTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomTokenRegistryTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomTokenRegistryTokenRemoved represents a TokenRemoved event raised by the FantomTokenRegistry contract.
type FantomTokenRegistryTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address token)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*FantomTokenRegistryTokenRemovedIterator, error) {

	logs, sub, err := _FantomTokenRegistry.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &FantomTokenRegistryTokenRemovedIterator{contract: _FantomTokenRegistry.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address token)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *FantomTokenRegistryTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _FantomTokenRegistry.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomTokenRegistryTokenRemoved)
				if err := _FantomTokenRegistry.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address token)
func (_FantomTokenRegistry *FantomTokenRegistryFilterer) ParseTokenRemoved(log types.Log) (*FantomTokenRegistryTokenRemoved, error) {
	event := new(FantomTokenRegistryTokenRemoved)
	if err := _FantomTokenRegistry.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
