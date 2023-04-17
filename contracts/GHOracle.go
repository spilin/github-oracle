// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prUrl\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ghUsername\",\"type\":\"string\"}],\"name\":\"Query\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"response\",\"type\":\"bool\"}],\"name\":\"Response\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ghUsername\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"response\",\"type\":\"bool\"}],\"name\":\"respond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805534801561001457600080fd5b50610531806100246000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806381ade3071461003b578063a9c577c514610057575b600080fd5b61005560048036038101906100509190610292565b610073565b005b610071600480360381019061006c9190610378565b6100d1565b005b600080549050600160008082825461008b91906103e7565b92505081905550807fef53eec7639d9fe59f45e1b3e3706008c1ef33456637a921b771ec4d5ac7b60484846040516100c492919061049a565b60405180910390a2505050565b806001600084815260200190815260200160002060006101000a81548160ff021916908315150217905550817fccb55462f4afaf05410376bbf3a404c04081149cfc3d0f2a46ce1c99fdc0cc9c8260405161012c91906104e0565b60405180910390a25050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61019f82610156565b810181811067ffffffffffffffff821117156101be576101bd610167565b5b80604052505050565b60006101d1610138565b90506101dd8282610196565b919050565b600067ffffffffffffffff8211156101fd576101fc610167565b5b61020682610156565b9050602081019050919050565b82818337600083830152505050565b6000610235610230846101e2565b6101c7565b90508281526020810184848401111561025157610250610151565b5b61025c848285610213565b509392505050565b600082601f8301126102795761027861014c565b5b8135610289848260208601610222565b91505092915050565b600080604083850312156102a9576102a8610142565b5b600083013567ffffffffffffffff8111156102c7576102c6610147565b5b6102d385828601610264565b925050602083013567ffffffffffffffff8111156102f4576102f3610147565b5b61030085828601610264565b9150509250929050565b6000819050919050565b61031d8161030a565b811461032857600080fd5b50565b60008135905061033a81610314565b92915050565b60008115159050919050565b61035581610340565b811461036057600080fd5b50565b6000813590506103728161034c565b92915050565b6000806040838503121561038f5761038e610142565b5b600061039d8582860161032b565b92505060206103ae85828601610363565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006103f28261030a565b91506103fd8361030a565b9250828201905080821115610415576104146103b8565b5b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561045557808201518184015260208101905061043a565b60008484015250505050565b600061046c8261041b565b6104768185610426565b9350610486818560208601610437565b61048f81610156565b840191505092915050565b600060408201905081810360008301526104b48185610461565b905081810360208301526104c88184610461565b90509392505050565b6104da81610340565b82525050565b60006020820190506104f560008301846104d1565b9291505056fea2646970667358221220dc8aaebdb06b1fb05373c0d17501bd31c8084df30decdfdea0746bbb28790db864736f6c63430008130033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Query is a paid mutator transaction binding the contract method 0x81ade307.
//
// Solidity: function query(string prUrl, string ghUsername) returns()
func (_Oracle *OracleTransactor) Query(opts *bind.TransactOpts, prUrl string, ghUsername string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "query", prUrl, ghUsername)
}

// Query is a paid mutator transaction binding the contract method 0x81ade307.
//
// Solidity: function query(string prUrl, string ghUsername) returns()
func (_Oracle *OracleSession) Query(prUrl string, ghUsername string) (*types.Transaction, error) {
	return _Oracle.Contract.Query(&_Oracle.TransactOpts, prUrl, ghUsername)
}

// Query is a paid mutator transaction binding the contract method 0x81ade307.
//
// Solidity: function query(string prUrl, string ghUsername) returns()
func (_Oracle *OracleTransactorSession) Query(prUrl string, ghUsername string) (*types.Transaction, error) {
	return _Oracle.Contract.Query(&_Oracle.TransactOpts, prUrl, ghUsername)
}

// Respond is a paid mutator transaction binding the contract method 0xa9c577c5.
//
// Solidity: function respond(uint256 id, bool response) returns()
func (_Oracle *OracleTransactor) Respond(opts *bind.TransactOpts, id *big.Int, response bool) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "respond", id, response)
}

// Respond is a paid mutator transaction binding the contract method 0xa9c577c5.
//
// Solidity: function respond(uint256 id, bool response) returns()
func (_Oracle *OracleSession) Respond(id *big.Int, response bool) (*types.Transaction, error) {
	return _Oracle.Contract.Respond(&_Oracle.TransactOpts, id, response)
}

// Respond is a paid mutator transaction binding the contract method 0xa9c577c5.
//
// Solidity: function respond(uint256 id, bool response) returns()
func (_Oracle *OracleTransactorSession) Respond(id *big.Int, response bool) (*types.Transaction, error) {
	return _Oracle.Contract.Respond(&_Oracle.TransactOpts, id, response)
}

// OracleQueryIterator is returned from FilterQuery and is used to iterate over the raw logs and unpacked data for Query events raised by the Oracle contract.
type OracleQueryIterator struct {
	Event *OracleQuery // Event containing the contract specifics and raw log

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
func (it *OracleQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleQuery)
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
		it.Event = new(OracleQuery)
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
func (it *OracleQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleQuery represents a Query event raised by the Oracle contract.
type OracleQuery struct {
	Id         *big.Int
	PrUrl      string
	GhUsername string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuery is a free log retrieval operation binding the contract event 0xef53eec7639d9fe59f45e1b3e3706008c1ef33456637a921b771ec4d5ac7b604.
//
// Solidity: event Query(uint256 indexed id, string prUrl, string ghUsername)
func (_Oracle *OracleFilterer) FilterQuery(opts *bind.FilterOpts, id []*big.Int) (*OracleQueryIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Query", idRule)
	if err != nil {
		return nil, err
	}
	return &OracleQueryIterator{contract: _Oracle.contract, event: "Query", logs: logs, sub: sub}, nil
}

// WatchQuery is a free log subscription operation binding the contract event 0xef53eec7639d9fe59f45e1b3e3706008c1ef33456637a921b771ec4d5ac7b604.
//
// Solidity: event Query(uint256 indexed id, string prUrl, string ghUsername)
func (_Oracle *OracleFilterer) WatchQuery(opts *bind.WatchOpts, sink chan<- *OracleQuery, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Query", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleQuery)
				if err := _Oracle.contract.UnpackLog(event, "Query", log); err != nil {
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

// ParseQuery is a log parse operation binding the contract event 0xef53eec7639d9fe59f45e1b3e3706008c1ef33456637a921b771ec4d5ac7b604.
//
// Solidity: event Query(uint256 indexed id, string prUrl, string ghUsername)
func (_Oracle *OracleFilterer) ParseQuery(log types.Log) (*OracleQuery, error) {
	event := new(OracleQuery)
	if err := _Oracle.contract.UnpackLog(event, "Query", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleResponseIterator is returned from FilterResponse and is used to iterate over the raw logs and unpacked data for Response events raised by the Oracle contract.
type OracleResponseIterator struct {
	Event *OracleResponse // Event containing the contract specifics and raw log

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
func (it *OracleResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleResponse)
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
		it.Event = new(OracleResponse)
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
func (it *OracleResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleResponse represents a Response event raised by the Oracle contract.
type OracleResponse struct {
	Id       *big.Int
	Response bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResponse is a free log retrieval operation binding the contract event 0xccb55462f4afaf05410376bbf3a404c04081149cfc3d0f2a46ce1c99fdc0cc9c.
//
// Solidity: event Response(uint256 indexed id, bool response)
func (_Oracle *OracleFilterer) FilterResponse(opts *bind.FilterOpts, id []*big.Int) (*OracleResponseIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Response", idRule)
	if err != nil {
		return nil, err
	}
	return &OracleResponseIterator{contract: _Oracle.contract, event: "Response", logs: logs, sub: sub}, nil
}

// WatchResponse is a free log subscription operation binding the contract event 0xccb55462f4afaf05410376bbf3a404c04081149cfc3d0f2a46ce1c99fdc0cc9c.
//
// Solidity: event Response(uint256 indexed id, bool response)
func (_Oracle *OracleFilterer) WatchResponse(opts *bind.WatchOpts, sink chan<- *OracleResponse, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Response", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleResponse)
				if err := _Oracle.contract.UnpackLog(event, "Response", log); err != nil {
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

// ParseResponse is a log parse operation binding the contract event 0xccb55462f4afaf05410376bbf3a404c04081149cfc3d0f2a46ce1c99fdc0cc9c.
//
// Solidity: event Response(uint256 indexed id, bool response)
func (_Oracle *OracleFilterer) ParseResponse(log types.Log) (*OracleResponse, error) {
	event := new(OracleResponse)
	if err := _Oracle.contract.UnpackLog(event, "Response", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
