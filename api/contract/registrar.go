// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RegistrarABI is the input ABI used to generate the binding from.
const RegistrarABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voter\",\"type\":\"bytes32\"}],\"name\":\"checkVoterRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voter\",\"type\":\"bytes32\"}],\"name\":\"registerVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistrarBin is the compiled bytecode used for deploying new contracts.
var RegistrarBin = "0x608060405234801561001057600080fd5b50610158806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806310103b721461003b578063793c5d211461007f575b600080fd5b6100676004803603602081101561005157600080fd5b81019080803590602001909291905050506100c3565b60405180821515815260200191505060405180910390f35b6100ab6004803603602081101561009557600080fd5b81019080803590602001909291905050506100ec565b60405180821515815260200191505060405180910390f35b600080600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000600160008084815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905091905056fea2646970667358221220c3f63ad665173cd93d5f688b2529f373a1214f3ee77041488c861c9c69d6b32564736f6c63430007030033"

// DeployRegistrar deploys a new Ethereum contract, binding an instance of Registrar to it.
func DeployRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registrar, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistrarBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registrar{RegistrarCaller: RegistrarCaller{contract: contract}, RegistrarTransactor: RegistrarTransactor{contract: contract}, RegistrarFilterer: RegistrarFilterer{contract: contract}}, nil
}

// Registrar is an auto generated Go binding around an Ethereum contract.
type Registrar struct {
	RegistrarCaller     // Read-only binding to the contract
	RegistrarTransactor // Write-only binding to the contract
	RegistrarFilterer   // Log filterer for contract events
}

// RegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarSession struct {
	Contract     *Registrar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarCallerSession struct {
	Contract *RegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarTransactorSession struct {
	Contract     *RegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarRaw struct {
	Contract *Registrar // Generic contract binding to access the raw methods on
}

// RegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarCallerRaw struct {
	Contract *RegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarTransactorRaw struct {
	Contract *RegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrar creates a new instance of Registrar, bound to a specific deployed contract.
func NewRegistrar(address common.Address, backend bind.ContractBackend) (*Registrar, error) {
	contract, err := bindRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrar{RegistrarCaller: RegistrarCaller{contract: contract}, RegistrarTransactor: RegistrarTransactor{contract: contract}, RegistrarFilterer: RegistrarFilterer{contract: contract}}, nil
}

// NewRegistrarCaller creates a new read-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarCaller(address common.Address, caller bind.ContractCaller) (*RegistrarCaller, error) {
	contract, err := bindRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarCaller{contract: contract}, nil
}

// NewRegistrarTransactor creates a new write-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarTransactor, error) {
	contract, err := bindRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarTransactor{contract: contract}, nil
}

// NewRegistrarFilterer creates a new log filterer instance of Registrar, bound to a specific deployed contract.
func NewRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrarFilterer, error) {
	contract, err := bindRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrarFilterer{contract: contract}, nil
}

// bindRegistrar binds a generic wrapper to an already deployed contract.
func bindRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.RegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transact(opts, method, params...)
}

// CheckVoterRegistered is a free data retrieval call binding the contract method 0x10103b72.
//
// Solidity: function checkVoterRegistered(bytes32 _voter) view returns(bool)
func (_Registrar *RegistrarCaller) CheckVoterRegistered(opts *bind.CallOpts, _voter [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registrar.contract.Call(opts, out, "checkVoterRegistered", _voter)
	return *ret0, err
}

// CheckVoterRegistered is a free data retrieval call binding the contract method 0x10103b72.
//
// Solidity: function checkVoterRegistered(bytes32 _voter) view returns(bool)
func (_Registrar *RegistrarSession) CheckVoterRegistered(_voter [32]byte) (bool, error) {
	return _Registrar.Contract.CheckVoterRegistered(&_Registrar.CallOpts, _voter)
}

// CheckVoterRegistered is a free data retrieval call binding the contract method 0x10103b72.
//
// Solidity: function checkVoterRegistered(bytes32 _voter) view returns(bool)
func (_Registrar *RegistrarCallerSession) CheckVoterRegistered(_voter [32]byte) (bool, error) {
	return _Registrar.Contract.CheckVoterRegistered(&_Registrar.CallOpts, _voter)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x793c5d21.
//
// Solidity: function registerVoter(bytes32 _voter) returns(bool)
func (_Registrar *RegistrarTransactor) RegisterVoter(opts *bind.TransactOpts, _voter [32]byte) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "registerVoter", _voter)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x793c5d21.
//
// Solidity: function registerVoter(bytes32 _voter) returns(bool)
func (_Registrar *RegistrarSession) RegisterVoter(_voter [32]byte) (*types.Transaction, error) {
	return _Registrar.Contract.RegisterVoter(&_Registrar.TransactOpts, _voter)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x793c5d21.
//
// Solidity: function registerVoter(bytes32 _voter) returns(bool)
func (_Registrar *RegistrarTransactorSession) RegisterVoter(_voter [32]byte) (*types.Transaction, error) {
	return _Registrar.Contract.RegisterVoter(&_Registrar.TransactOpts, _voter)
}
