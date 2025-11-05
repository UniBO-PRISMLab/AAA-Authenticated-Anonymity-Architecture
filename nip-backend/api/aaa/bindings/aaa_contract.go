// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// AAAContractMetaData contains all meta data concerning the AAAContract contract.
var AAAContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nodes\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"words\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redundancyFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"symK\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sid\",\"type\":\"bytes32\"}],\"name\":\"PIDEncryptionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encSID\",\"type\":\"bytes\"}],\"name\":\"PhraseComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromNode\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toNode\",\"type\":\"address\"}],\"name\":\"RedundancyRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"wordHash\",\"type\":\"bytes32\"}],\"name\":\"RedundantWordSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"userPK\",\"type\":\"bytes\"}],\"name\":\"SIDEncryptionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"SeedPhraseProtocolInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"userPK\",\"type\":\"bytes\"}],\"name\":\"WordRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"wordHash\",\"type\":\"bytes32\"}],\"name\":\"WordSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REDUNDANCY_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WORDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pk\",\"type\":\"bytes\"}],\"name\":\"getSACRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"getSID\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sid\",\"type\":\"bytes32\"}],\"name\":\"getSIDRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"getSelectedNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"getWords\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"removeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pk\",\"type\":\"bytes\"}],\"name\":\"seedPhraseGenerationProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"selectedNodesByPID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encPID\",\"type\":\"bytes\"}],\"name\":\"submitEncryptedPID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encSID\",\"type\":\"bytes\"}],\"name\":\"submitEncryptedSID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encryptedWord\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"nodePK\",\"type\":\"bytes\"}],\"name\":\"submitEncryptedWord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sac\",\"type\":\"uint256\"}],\"name\":\"submitSAC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sac\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pk\",\"type\":\"bytes\"}],\"name\":\"submitSACRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AAAContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AAAContractMetaData.ABI instead.
var AAAContractABI = AAAContractMetaData.ABI

// AAAContract is an auto generated Go binding around an Ethereum contract.
type AAAContract struct {
	AAAContractCaller     // Read-only binding to the contract
	AAAContractTransactor // Write-only binding to the contract
	AAAContractFilterer   // Log filterer for contract events
}

// AAAContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AAAContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAAContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AAAContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAAContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AAAContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAAContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AAAContractSession struct {
	Contract     *AAAContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAAContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AAAContractCallerSession struct {
	Contract *AAAContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AAAContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AAAContractTransactorSession struct {
	Contract     *AAAContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AAAContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AAAContractRaw struct {
	Contract *AAAContract // Generic contract binding to access the raw methods on
}

// AAAContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AAAContractCallerRaw struct {
	Contract *AAAContractCaller // Generic read-only contract binding to access the raw methods on
}

// AAAContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AAAContractTransactorRaw struct {
	Contract *AAAContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAAAContract creates a new instance of AAAContract, bound to a specific deployed contract.
func NewAAAContract(address common.Address, backend bind.ContractBackend) (*AAAContract, error) {
	contract, err := bindAAAContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AAAContract{AAAContractCaller: AAAContractCaller{contract: contract}, AAAContractTransactor: AAAContractTransactor{contract: contract}, AAAContractFilterer: AAAContractFilterer{contract: contract}}, nil
}

// NewAAAContractCaller creates a new read-only instance of AAAContract, bound to a specific deployed contract.
func NewAAAContractCaller(address common.Address, caller bind.ContractCaller) (*AAAContractCaller, error) {
	contract, err := bindAAAContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AAAContractCaller{contract: contract}, nil
}

// NewAAAContractTransactor creates a new write-only instance of AAAContract, bound to a specific deployed contract.
func NewAAAContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AAAContractTransactor, error) {
	contract, err := bindAAAContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AAAContractTransactor{contract: contract}, nil
}

// NewAAAContractFilterer creates a new log filterer instance of AAAContract, bound to a specific deployed contract.
func NewAAAContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AAAContractFilterer, error) {
	contract, err := bindAAAContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AAAContractFilterer{contract: contract}, nil
}

// bindAAAContract binds a generic wrapper to an already deployed contract.
func bindAAAContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AAAContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAAContract *AAAContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAAContract.Contract.AAAContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAAContract *AAAContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAAContract.Contract.AAAContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAAContract *AAAContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAAContract.Contract.AAAContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAAContract *AAAContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAAContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAAContract *AAAContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAAContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAAContract *AAAContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAAContract.Contract.contract.Transact(opts, method, params...)
}

// REDUNDANCYFACTOR is a free data retrieval call binding the contract method 0x288a1319.
//
// Solidity: function REDUNDANCY_FACTOR() view returns(uint256)
func (_AAAContract *AAAContractCaller) REDUNDANCYFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "REDUNDANCY_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDUNDANCYFACTOR is a free data retrieval call binding the contract method 0x288a1319.
//
// Solidity: function REDUNDANCY_FACTOR() view returns(uint256)
func (_AAAContract *AAAContractSession) REDUNDANCYFACTOR() (*big.Int, error) {
	return _AAAContract.Contract.REDUNDANCYFACTOR(&_AAAContract.CallOpts)
}

// REDUNDANCYFACTOR is a free data retrieval call binding the contract method 0x288a1319.
//
// Solidity: function REDUNDANCY_FACTOR() view returns(uint256)
func (_AAAContract *AAAContractCallerSession) REDUNDANCYFACTOR() (*big.Int, error) {
	return _AAAContract.Contract.REDUNDANCYFACTOR(&_AAAContract.CallOpts)
}

// WORDS is a free data retrieval call binding the contract method 0x4ecb3a31.
//
// Solidity: function WORDS() view returns(uint256)
func (_AAAContract *AAAContractCaller) WORDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "WORDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WORDS is a free data retrieval call binding the contract method 0x4ecb3a31.
//
// Solidity: function WORDS() view returns(uint256)
func (_AAAContract *AAAContractSession) WORDS() (*big.Int, error) {
	return _AAAContract.Contract.WORDS(&_AAAContract.CallOpts)
}

// WORDS is a free data retrieval call binding the contract method 0x4ecb3a31.
//
// Solidity: function WORDS() view returns(uint256)
func (_AAAContract *AAAContractCallerSession) WORDS() (*big.Int, error) {
	return _AAAContract.Contract.WORDS(&_AAAContract.CallOpts)
}

// GetSACRecord is a free data retrieval call binding the contract method 0xf0307989.
//
// Solidity: function getSACRecord(bytes pk) view returns(uint256)
func (_AAAContract *AAAContractCaller) GetSACRecord(opts *bind.CallOpts, pk []byte) (*big.Int, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "getSACRecord", pk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSACRecord is a free data retrieval call binding the contract method 0xf0307989.
//
// Solidity: function getSACRecord(bytes pk) view returns(uint256)
func (_AAAContract *AAAContractSession) GetSACRecord(pk []byte) (*big.Int, error) {
	return _AAAContract.Contract.GetSACRecord(&_AAAContract.CallOpts, pk)
}

// GetSACRecord is a free data retrieval call binding the contract method 0xf0307989.
//
// Solidity: function getSACRecord(bytes pk) view returns(uint256)
func (_AAAContract *AAAContractCallerSession) GetSACRecord(pk []byte) (*big.Int, error) {
	return _AAAContract.Contract.GetSACRecord(&_AAAContract.CallOpts, pk)
}

// GetSID is a free data retrieval call binding the contract method 0xb7ed02d3.
//
// Solidity: function getSID(bytes32 pid) view returns(bytes)
func (_AAAContract *AAAContractCaller) GetSID(opts *bind.CallOpts, pid [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "getSID", pid)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSID is a free data retrieval call binding the contract method 0xb7ed02d3.
//
// Solidity: function getSID(bytes32 pid) view returns(bytes)
func (_AAAContract *AAAContractSession) GetSID(pid [32]byte) ([]byte, error) {
	return _AAAContract.Contract.GetSID(&_AAAContract.CallOpts, pid)
}

// GetSID is a free data retrieval call binding the contract method 0xb7ed02d3.
//
// Solidity: function getSID(bytes32 pid) view returns(bytes)
func (_AAAContract *AAAContractCallerSession) GetSID(pid [32]byte) ([]byte, error) {
	return _AAAContract.Contract.GetSID(&_AAAContract.CallOpts, pid)
}

// GetSIDRecord is a free data retrieval call binding the contract method 0x239db46f.
//
// Solidity: function getSIDRecord(bytes32 sid) view returns(bytes, bytes)
func (_AAAContract *AAAContractCaller) GetSIDRecord(opts *bind.CallOpts, sid [32]byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "getSIDRecord", sid)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetSIDRecord is a free data retrieval call binding the contract method 0x239db46f.
//
// Solidity: function getSIDRecord(bytes32 sid) view returns(bytes, bytes)
func (_AAAContract *AAAContractSession) GetSIDRecord(sid [32]byte) ([]byte, []byte, error) {
	return _AAAContract.Contract.GetSIDRecord(&_AAAContract.CallOpts, sid)
}

// GetSIDRecord is a free data retrieval call binding the contract method 0x239db46f.
//
// Solidity: function getSIDRecord(bytes32 sid) view returns(bytes, bytes)
func (_AAAContract *AAAContractCallerSession) GetSIDRecord(sid [32]byte) ([]byte, []byte, error) {
	return _AAAContract.Contract.GetSIDRecord(&_AAAContract.CallOpts, sid)
}

// GetSelectedNodes is a free data retrieval call binding the contract method 0xc7b4d5a3.
//
// Solidity: function getSelectedNodes(bytes32 pid) view returns(address[])
func (_AAAContract *AAAContractCaller) GetSelectedNodes(opts *bind.CallOpts, pid [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "getSelectedNodes", pid)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSelectedNodes is a free data retrieval call binding the contract method 0xc7b4d5a3.
//
// Solidity: function getSelectedNodes(bytes32 pid) view returns(address[])
func (_AAAContract *AAAContractSession) GetSelectedNodes(pid [32]byte) ([]common.Address, error) {
	return _AAAContract.Contract.GetSelectedNodes(&_AAAContract.CallOpts, pid)
}

// GetSelectedNodes is a free data retrieval call binding the contract method 0xc7b4d5a3.
//
// Solidity: function getSelectedNodes(bytes32 pid) view returns(address[])
func (_AAAContract *AAAContractCallerSession) GetSelectedNodes(pid [32]byte) ([]common.Address, error) {
	return _AAAContract.Contract.GetSelectedNodes(&_AAAContract.CallOpts, pid)
}

// GetWords is a free data retrieval call binding the contract method 0x72f43639.
//
// Solidity: function getWords(bytes32 pid) view returns(bytes[])
func (_AAAContract *AAAContractCaller) GetWords(opts *bind.CallOpts, pid [32]byte) ([][]byte, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "getWords", pid)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetWords is a free data retrieval call binding the contract method 0x72f43639.
//
// Solidity: function getWords(bytes32 pid) view returns(bytes[])
func (_AAAContract *AAAContractSession) GetWords(pid [32]byte) ([][]byte, error) {
	return _AAAContract.Contract.GetWords(&_AAAContract.CallOpts, pid)
}

// GetWords is a free data retrieval call binding the contract method 0x72f43639.
//
// Solidity: function getWords(bytes32 pid) view returns(bytes[])
func (_AAAContract *AAAContractCallerSession) GetWords(pid [32]byte) ([][]byte, error) {
	return _AAAContract.Contract.GetWords(&_AAAContract.CallOpts, pid)
}

// IsNode is a free data retrieval call binding the contract method 0x01750152.
//
// Solidity: function isNode(address ) view returns(bool)
func (_AAAContract *AAAContractCaller) IsNode(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "isNode", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNode is a free data retrieval call binding the contract method 0x01750152.
//
// Solidity: function isNode(address ) view returns(bool)
func (_AAAContract *AAAContractSession) IsNode(arg0 common.Address) (bool, error) {
	return _AAAContract.Contract.IsNode(&_AAAContract.CallOpts, arg0)
}

// IsNode is a free data retrieval call binding the contract method 0x01750152.
//
// Solidity: function isNode(address ) view returns(bool)
func (_AAAContract *AAAContractCallerSession) IsNode(arg0 common.Address) (bool, error) {
	return _AAAContract.Contract.IsNode(&_AAAContract.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_AAAContract *AAAContractCaller) NodeList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "nodeList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_AAAContract *AAAContractSession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _AAAContract.Contract.NodeList(&_AAAContract.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_AAAContract *AAAContractCallerSession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _AAAContract.Contract.NodeList(&_AAAContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAAContract *AAAContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAAContract *AAAContractSession) Owner() (common.Address, error) {
	return _AAAContract.Contract.Owner(&_AAAContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAAContract *AAAContractCallerSession) Owner() (common.Address, error) {
	return _AAAContract.Contract.Owner(&_AAAContract.CallOpts)
}

// SelectedNodesByPID is a free data retrieval call binding the contract method 0x6a978a99.
//
// Solidity: function selectedNodesByPID(bytes32 , uint256 ) view returns(address)
func (_AAAContract *AAAContractCaller) SelectedNodesByPID(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AAAContract.contract.Call(opts, &out, "selectedNodesByPID", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SelectedNodesByPID is a free data retrieval call binding the contract method 0x6a978a99.
//
// Solidity: function selectedNodesByPID(bytes32 , uint256 ) view returns(address)
func (_AAAContract *AAAContractSession) SelectedNodesByPID(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AAAContract.Contract.SelectedNodesByPID(&_AAAContract.CallOpts, arg0, arg1)
}

// SelectedNodesByPID is a free data retrieval call binding the contract method 0x6a978a99.
//
// Solidity: function selectedNodesByPID(bytes32 , uint256 ) view returns(address)
func (_AAAContract *AAAContractCallerSession) SelectedNodesByPID(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AAAContract.Contract.SelectedNodesByPID(&_AAAContract.CallOpts, arg0, arg1)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address node) returns()
func (_AAAContract *AAAContractTransactor) AddNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "addNode", node)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address node) returns()
func (_AAAContract *AAAContractSession) AddNode(node common.Address) (*types.Transaction, error) {
	return _AAAContract.Contract.AddNode(&_AAAContract.TransactOpts, node)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address node) returns()
func (_AAAContract *AAAContractTransactorSession) AddNode(node common.Address) (*types.Transaction, error) {
	return _AAAContract.Contract.AddNode(&_AAAContract.TransactOpts, node)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address node) returns()
func (_AAAContract *AAAContractTransactor) RemoveNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "removeNode", node)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address node) returns()
func (_AAAContract *AAAContractSession) RemoveNode(node common.Address) (*types.Transaction, error) {
	return _AAAContract.Contract.RemoveNode(&_AAAContract.TransactOpts, node)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address node) returns()
func (_AAAContract *AAAContractTransactorSession) RemoveNode(node common.Address) (*types.Transaction, error) {
	return _AAAContract.Contract.RemoveNode(&_AAAContract.TransactOpts, node)
}

// SeedPhraseGenerationProtocol is a paid mutator transaction binding the contract method 0x3b33892a.
//
// Solidity: function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) returns()
func (_AAAContract *AAAContractTransactor) SeedPhraseGenerationProtocol(opts *bind.TransactOpts, pid [32]byte, pk []byte) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "seedPhraseGenerationProtocol", pid, pk)
}

// SeedPhraseGenerationProtocol is a paid mutator transaction binding the contract method 0x3b33892a.
//
// Solidity: function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) returns()
func (_AAAContract *AAAContractSession) SeedPhraseGenerationProtocol(pid [32]byte, pk []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SeedPhraseGenerationProtocol(&_AAAContract.TransactOpts, pid, pk)
}

// SeedPhraseGenerationProtocol is a paid mutator transaction binding the contract method 0x3b33892a.
//
// Solidity: function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) returns()
func (_AAAContract *AAAContractTransactorSession) SeedPhraseGenerationProtocol(pid [32]byte, pk []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SeedPhraseGenerationProtocol(&_AAAContract.TransactOpts, pid, pk)
}

// SubmitEncryptedPID is a paid mutator transaction binding the contract method 0x00e12e3a.
//
// Solidity: function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) returns()
func (_AAAContract *AAAContractTransactor) SubmitEncryptedPID(opts *bind.TransactOpts, pid [32]byte, sid [32]byte, encPID []byte) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "submitEncryptedPID", pid, sid, encPID)
}

// SubmitEncryptedPID is a paid mutator transaction binding the contract method 0x00e12e3a.
//
// Solidity: function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) returns()
func (_AAAContract *AAAContractSession) SubmitEncryptedPID(pid [32]byte, sid [32]byte, encPID []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitEncryptedPID(&_AAAContract.TransactOpts, pid, sid, encPID)
}

// SubmitEncryptedPID is a paid mutator transaction binding the contract method 0x00e12e3a.
//
// Solidity: function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) returns()
func (_AAAContract *AAAContractTransactorSession) SubmitEncryptedPID(pid [32]byte, sid [32]byte, encPID []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitEncryptedPID(&_AAAContract.TransactOpts, pid, sid, encPID)
}

// SubmitEncryptedSID is a paid mutator transaction binding the contract method 0x4542552c.
//
// Solidity: function submitEncryptedSID(bytes32 pid, bytes encSID) returns()
func (_AAAContract *AAAContractTransactor) SubmitEncryptedSID(opts *bind.TransactOpts, pid [32]byte, encSID []byte) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "submitEncryptedSID", pid, encSID)
}

// SubmitEncryptedSID is a paid mutator transaction binding the contract method 0x4542552c.
//
// Solidity: function submitEncryptedSID(bytes32 pid, bytes encSID) returns()
func (_AAAContract *AAAContractSession) SubmitEncryptedSID(pid [32]byte, encSID []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitEncryptedSID(&_AAAContract.TransactOpts, pid, encSID)
}

// SubmitEncryptedSID is a paid mutator transaction binding the contract method 0x4542552c.
//
// Solidity: function submitEncryptedSID(bytes32 pid, bytes encSID) returns()
func (_AAAContract *AAAContractTransactorSession) SubmitEncryptedSID(pid [32]byte, encSID []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitEncryptedSID(&_AAAContract.TransactOpts, pid, encSID)
}

// SubmitEncryptedWord is a paid mutator transaction binding the contract method 0x0470c9e9.
//
// Solidity: function submitEncryptedWord(bytes32 pid, bytes encryptedWord, bytes nodePK) returns()
func (_AAAContract *AAAContractTransactor) SubmitEncryptedWord(opts *bind.TransactOpts, pid [32]byte, encryptedWord []byte, nodePK []byte) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "submitEncryptedWord", pid, encryptedWord, nodePK)
}

// SubmitEncryptedWord is a paid mutator transaction binding the contract method 0x0470c9e9.
//
// Solidity: function submitEncryptedWord(bytes32 pid, bytes encryptedWord, bytes nodePK) returns()
func (_AAAContract *AAAContractSession) SubmitEncryptedWord(pid [32]byte, encryptedWord []byte, nodePK []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitEncryptedWord(&_AAAContract.TransactOpts, pid, encryptedWord, nodePK)
}

// SubmitEncryptedWord is a paid mutator transaction binding the contract method 0x0470c9e9.
//
// Solidity: function submitEncryptedWord(bytes32 pid, bytes encryptedWord, bytes nodePK) returns()
func (_AAAContract *AAAContractTransactorSession) SubmitEncryptedWord(pid [32]byte, encryptedWord []byte, nodePK []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitEncryptedWord(&_AAAContract.TransactOpts, pid, encryptedWord, nodePK)
}

// SubmitSAC is a paid mutator transaction binding the contract method 0xff20d844.
//
// Solidity: function submitSAC(uint256 sac) returns()
func (_AAAContract *AAAContractTransactor) SubmitSAC(opts *bind.TransactOpts, sac *big.Int) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "submitSAC", sac)
}

// SubmitSAC is a paid mutator transaction binding the contract method 0xff20d844.
//
// Solidity: function submitSAC(uint256 sac) returns()
func (_AAAContract *AAAContractSession) SubmitSAC(sac *big.Int) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitSAC(&_AAAContract.TransactOpts, sac)
}

// SubmitSAC is a paid mutator transaction binding the contract method 0xff20d844.
//
// Solidity: function submitSAC(uint256 sac) returns()
func (_AAAContract *AAAContractTransactorSession) SubmitSAC(sac *big.Int) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitSAC(&_AAAContract.TransactOpts, sac)
}

// SubmitSACRecord is a paid mutator transaction binding the contract method 0x7c5d77c0.
//
// Solidity: function submitSACRecord(uint256 sac, bytes pk) returns()
func (_AAAContract *AAAContractTransactor) SubmitSACRecord(opts *bind.TransactOpts, sac *big.Int, pk []byte) (*types.Transaction, error) {
	return _AAAContract.contract.Transact(opts, "submitSACRecord", sac, pk)
}

// SubmitSACRecord is a paid mutator transaction binding the contract method 0x7c5d77c0.
//
// Solidity: function submitSACRecord(uint256 sac, bytes pk) returns()
func (_AAAContract *AAAContractSession) SubmitSACRecord(sac *big.Int, pk []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitSACRecord(&_AAAContract.TransactOpts, sac, pk)
}

// SubmitSACRecord is a paid mutator transaction binding the contract method 0x7c5d77c0.
//
// Solidity: function submitSACRecord(uint256 sac, bytes pk) returns()
func (_AAAContract *AAAContractTransactorSession) SubmitSACRecord(sac *big.Int, pk []byte) (*types.Transaction, error) {
	return _AAAContract.Contract.SubmitSACRecord(&_AAAContract.TransactOpts, sac, pk)
}

// AAAContractNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the AAAContract contract.
type AAAContractNodeAddedIterator struct {
	Event *AAAContractNodeAdded // Event containing the contract specifics and raw log

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
func (it *AAAContractNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractNodeAdded)
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
		it.Event = new(AAAContractNodeAdded)
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
func (it *AAAContractNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractNodeAdded represents a NodeAdded event raised by the AAAContract contract.
type AAAContractNodeAdded struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address node)
func (_AAAContract *AAAContractFilterer) FilterNodeAdded(opts *bind.FilterOpts) (*AAAContractNodeAddedIterator, error) {

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return &AAAContractNodeAddedIterator{contract: _AAAContract.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address node)
func (_AAAContract *AAAContractFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *AAAContractNodeAdded) (event.Subscription, error) {

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractNodeAdded)
				if err := _AAAContract.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address node)
func (_AAAContract *AAAContractFilterer) ParseNodeAdded(log types.Log) (*AAAContractNodeAdded, error) {
	event := new(AAAContractNodeAdded)
	if err := _AAAContract.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractNodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the AAAContract contract.
type AAAContractNodeRemovedIterator struct {
	Event *AAAContractNodeRemoved // Event containing the contract specifics and raw log

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
func (it *AAAContractNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractNodeRemoved)
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
		it.Event = new(AAAContractNodeRemoved)
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
func (it *AAAContractNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractNodeRemoved represents a NodeRemoved event raised by the AAAContract contract.
type AAAContractNodeRemoved struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address node)
func (_AAAContract *AAAContractFilterer) FilterNodeRemoved(opts *bind.FilterOpts) (*AAAContractNodeRemovedIterator, error) {

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return &AAAContractNodeRemovedIterator{contract: _AAAContract.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address node)
func (_AAAContract *AAAContractFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *AAAContractNodeRemoved) (event.Subscription, error) {

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractNodeRemoved)
				if err := _AAAContract.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

// ParseNodeRemoved is a log parse operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address node)
func (_AAAContract *AAAContractFilterer) ParseNodeRemoved(log types.Log) (*AAAContractNodeRemoved, error) {
	event := new(AAAContractNodeRemoved)
	if err := _AAAContract.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractPIDEncryptionRequestedIterator is returned from FilterPIDEncryptionRequested and is used to iterate over the raw logs and unpacked data for PIDEncryptionRequested events raised by the AAAContract contract.
type AAAContractPIDEncryptionRequestedIterator struct {
	Event *AAAContractPIDEncryptionRequested // Event containing the contract specifics and raw log

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
func (it *AAAContractPIDEncryptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractPIDEncryptionRequested)
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
		it.Event = new(AAAContractPIDEncryptionRequested)
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
func (it *AAAContractPIDEncryptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractPIDEncryptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractPIDEncryptionRequested represents a PIDEncryptionRequested event raised by the AAAContract contract.
type AAAContractPIDEncryptionRequested struct {
	Pid  [32]byte
	Node common.Address
	SymK [32]byte
	Sid  [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPIDEncryptionRequested is a free log retrieval operation binding the contract event 0x01c8b3ea5e773e12ac7fb27cd0c382713fd140d4db5916ee0c1bce601bdc4d50.
//
// Solidity: event PIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes32 symK, bytes32 sid)
func (_AAAContract *AAAContractFilterer) FilterPIDEncryptionRequested(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAAContractPIDEncryptionRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "PIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractPIDEncryptionRequestedIterator{contract: _AAAContract.contract, event: "PIDEncryptionRequested", logs: logs, sub: sub}, nil
}

// WatchPIDEncryptionRequested is a free log subscription operation binding the contract event 0x01c8b3ea5e773e12ac7fb27cd0c382713fd140d4db5916ee0c1bce601bdc4d50.
//
// Solidity: event PIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes32 symK, bytes32 sid)
func (_AAAContract *AAAContractFilterer) WatchPIDEncryptionRequested(opts *bind.WatchOpts, sink chan<- *AAAContractPIDEncryptionRequested, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "PIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractPIDEncryptionRequested)
				if err := _AAAContract.contract.UnpackLog(event, "PIDEncryptionRequested", log); err != nil {
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

// ParsePIDEncryptionRequested is a log parse operation binding the contract event 0x01c8b3ea5e773e12ac7fb27cd0c382713fd140d4db5916ee0c1bce601bdc4d50.
//
// Solidity: event PIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes32 symK, bytes32 sid)
func (_AAAContract *AAAContractFilterer) ParsePIDEncryptionRequested(log types.Log) (*AAAContractPIDEncryptionRequested, error) {
	event := new(AAAContractPIDEncryptionRequested)
	if err := _AAAContract.contract.UnpackLog(event, "PIDEncryptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractPhraseCompleteIterator is returned from FilterPhraseComplete and is used to iterate over the raw logs and unpacked data for PhraseComplete events raised by the AAAContract contract.
type AAAContractPhraseCompleteIterator struct {
	Event *AAAContractPhraseComplete // Event containing the contract specifics and raw log

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
func (it *AAAContractPhraseCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractPhraseComplete)
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
		it.Event = new(AAAContractPhraseComplete)
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
func (it *AAAContractPhraseCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractPhraseCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractPhraseComplete represents a PhraseComplete event raised by the AAAContract contract.
type AAAContractPhraseComplete struct {
	Pid    [32]byte
	EncSID []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPhraseComplete is a free log retrieval operation binding the contract event 0xa24f64aabeee785513b21a32ea1d294c10dd37c0b1ba653b03e9825bcfa1769f.
//
// Solidity: event PhraseComplete(bytes32 indexed pid, bytes encSID)
func (_AAAContract *AAAContractFilterer) FilterPhraseComplete(opts *bind.FilterOpts, pid [][32]byte) (*AAAContractPhraseCompleteIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "PhraseComplete", pidRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractPhraseCompleteIterator{contract: _AAAContract.contract, event: "PhraseComplete", logs: logs, sub: sub}, nil
}

// WatchPhraseComplete is a free log subscription operation binding the contract event 0xa24f64aabeee785513b21a32ea1d294c10dd37c0b1ba653b03e9825bcfa1769f.
//
// Solidity: event PhraseComplete(bytes32 indexed pid, bytes encSID)
func (_AAAContract *AAAContractFilterer) WatchPhraseComplete(opts *bind.WatchOpts, sink chan<- *AAAContractPhraseComplete, pid [][32]byte) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "PhraseComplete", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractPhraseComplete)
				if err := _AAAContract.contract.UnpackLog(event, "PhraseComplete", log); err != nil {
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

// ParsePhraseComplete is a log parse operation binding the contract event 0xa24f64aabeee785513b21a32ea1d294c10dd37c0b1ba653b03e9825bcfa1769f.
//
// Solidity: event PhraseComplete(bytes32 indexed pid, bytes encSID)
func (_AAAContract *AAAContractFilterer) ParsePhraseComplete(log types.Log) (*AAAContractPhraseComplete, error) {
	event := new(AAAContractPhraseComplete)
	if err := _AAAContract.contract.UnpackLog(event, "PhraseComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractRedundancyRequestedIterator is returned from FilterRedundancyRequested and is used to iterate over the raw logs and unpacked data for RedundancyRequested events raised by the AAAContract contract.
type AAAContractRedundancyRequestedIterator struct {
	Event *AAAContractRedundancyRequested // Event containing the contract specifics and raw log

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
func (it *AAAContractRedundancyRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractRedundancyRequested)
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
		it.Event = new(AAAContractRedundancyRequested)
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
func (it *AAAContractRedundancyRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractRedundancyRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractRedundancyRequested represents a RedundancyRequested event raised by the AAAContract contract.
type AAAContractRedundancyRequested struct {
	Pid      [32]byte
	Index    *big.Int
	FromNode common.Address
	ToNode   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedundancyRequested is a free log retrieval operation binding the contract event 0x69cf9406242d538f49d684d5ac357be065c1d26c7054702774fa1179a7b4d726.
//
// Solidity: event RedundancyRequested(bytes32 indexed pid, uint256 indexed index, address indexed fromNode, address toNode)
func (_AAAContract *AAAContractFilterer) FilterRedundancyRequested(opts *bind.FilterOpts, pid [][32]byte, index []*big.Int, fromNode []common.Address) (*AAAContractRedundancyRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var fromNodeRule []interface{}
	for _, fromNodeItem := range fromNode {
		fromNodeRule = append(fromNodeRule, fromNodeItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "RedundancyRequested", pidRule, indexRule, fromNodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractRedundancyRequestedIterator{contract: _AAAContract.contract, event: "RedundancyRequested", logs: logs, sub: sub}, nil
}

// WatchRedundancyRequested is a free log subscription operation binding the contract event 0x69cf9406242d538f49d684d5ac357be065c1d26c7054702774fa1179a7b4d726.
//
// Solidity: event RedundancyRequested(bytes32 indexed pid, uint256 indexed index, address indexed fromNode, address toNode)
func (_AAAContract *AAAContractFilterer) WatchRedundancyRequested(opts *bind.WatchOpts, sink chan<- *AAAContractRedundancyRequested, pid [][32]byte, index []*big.Int, fromNode []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var fromNodeRule []interface{}
	for _, fromNodeItem := range fromNode {
		fromNodeRule = append(fromNodeRule, fromNodeItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "RedundancyRequested", pidRule, indexRule, fromNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractRedundancyRequested)
				if err := _AAAContract.contract.UnpackLog(event, "RedundancyRequested", log); err != nil {
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

// ParseRedundancyRequested is a log parse operation binding the contract event 0x69cf9406242d538f49d684d5ac357be065c1d26c7054702774fa1179a7b4d726.
//
// Solidity: event RedundancyRequested(bytes32 indexed pid, uint256 indexed index, address indexed fromNode, address toNode)
func (_AAAContract *AAAContractFilterer) ParseRedundancyRequested(log types.Log) (*AAAContractRedundancyRequested, error) {
	event := new(AAAContractRedundancyRequested)
	if err := _AAAContract.contract.UnpackLog(event, "RedundancyRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractRedundantWordSubmittedIterator is returned from FilterRedundantWordSubmitted and is used to iterate over the raw logs and unpacked data for RedundantWordSubmitted events raised by the AAAContract contract.
type AAAContractRedundantWordSubmittedIterator struct {
	Event *AAAContractRedundantWordSubmitted // Event containing the contract specifics and raw log

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
func (it *AAAContractRedundantWordSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractRedundantWordSubmitted)
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
		it.Event = new(AAAContractRedundantWordSubmitted)
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
func (it *AAAContractRedundantWordSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractRedundantWordSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractRedundantWordSubmitted represents a RedundantWordSubmitted event raised by the AAAContract contract.
type AAAContractRedundantWordSubmitted struct {
	Pid      [32]byte
	Index    *big.Int
	Node     common.Address
	WordHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedundantWordSubmitted is a free log retrieval operation binding the contract event 0x2e48f225c7128740b00ba988b438fb1cdd63d84a479c8d821ed3e39e20e37b09.
//
// Solidity: event RedundantWordSubmitted(bytes32 indexed pid, uint256 indexed index, address indexed node, bytes32 wordHash)
func (_AAAContract *AAAContractFilterer) FilterRedundantWordSubmitted(opts *bind.FilterOpts, pid [][32]byte, index []*big.Int, node []common.Address) (*AAAContractRedundantWordSubmittedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "RedundantWordSubmitted", pidRule, indexRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractRedundantWordSubmittedIterator{contract: _AAAContract.contract, event: "RedundantWordSubmitted", logs: logs, sub: sub}, nil
}

// WatchRedundantWordSubmitted is a free log subscription operation binding the contract event 0x2e48f225c7128740b00ba988b438fb1cdd63d84a479c8d821ed3e39e20e37b09.
//
// Solidity: event RedundantWordSubmitted(bytes32 indexed pid, uint256 indexed index, address indexed node, bytes32 wordHash)
func (_AAAContract *AAAContractFilterer) WatchRedundantWordSubmitted(opts *bind.WatchOpts, sink chan<- *AAAContractRedundantWordSubmitted, pid [][32]byte, index []*big.Int, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "RedundantWordSubmitted", pidRule, indexRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractRedundantWordSubmitted)
				if err := _AAAContract.contract.UnpackLog(event, "RedundantWordSubmitted", log); err != nil {
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

// ParseRedundantWordSubmitted is a log parse operation binding the contract event 0x2e48f225c7128740b00ba988b438fb1cdd63d84a479c8d821ed3e39e20e37b09.
//
// Solidity: event RedundantWordSubmitted(bytes32 indexed pid, uint256 indexed index, address indexed node, bytes32 wordHash)
func (_AAAContract *AAAContractFilterer) ParseRedundantWordSubmitted(log types.Log) (*AAAContractRedundantWordSubmitted, error) {
	event := new(AAAContractRedundantWordSubmitted)
	if err := _AAAContract.contract.UnpackLog(event, "RedundantWordSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractSIDEncryptionRequestedIterator is returned from FilterSIDEncryptionRequested and is used to iterate over the raw logs and unpacked data for SIDEncryptionRequested events raised by the AAAContract contract.
type AAAContractSIDEncryptionRequestedIterator struct {
	Event *AAAContractSIDEncryptionRequested // Event containing the contract specifics and raw log

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
func (it *AAAContractSIDEncryptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractSIDEncryptionRequested)
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
		it.Event = new(AAAContractSIDEncryptionRequested)
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
func (it *AAAContractSIDEncryptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractSIDEncryptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractSIDEncryptionRequested represents a SIDEncryptionRequested event raised by the AAAContract contract.
type AAAContractSIDEncryptionRequested struct {
	Pid    [32]byte
	Node   common.Address
	Sid    []byte
	UserPK []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSIDEncryptionRequested is a free log retrieval operation binding the contract event 0xe0ea39226dca1f5e473a57d8db19ad9f9578535bda4579ebdc7659681cba31f5.
//
// Solidity: event SIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes sid, bytes userPK)
func (_AAAContract *AAAContractFilterer) FilterSIDEncryptionRequested(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAAContractSIDEncryptionRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "SIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractSIDEncryptionRequestedIterator{contract: _AAAContract.contract, event: "SIDEncryptionRequested", logs: logs, sub: sub}, nil
}

// WatchSIDEncryptionRequested is a free log subscription operation binding the contract event 0xe0ea39226dca1f5e473a57d8db19ad9f9578535bda4579ebdc7659681cba31f5.
//
// Solidity: event SIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes sid, bytes userPK)
func (_AAAContract *AAAContractFilterer) WatchSIDEncryptionRequested(opts *bind.WatchOpts, sink chan<- *AAAContractSIDEncryptionRequested, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "SIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractSIDEncryptionRequested)
				if err := _AAAContract.contract.UnpackLog(event, "SIDEncryptionRequested", log); err != nil {
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

// ParseSIDEncryptionRequested is a log parse operation binding the contract event 0xe0ea39226dca1f5e473a57d8db19ad9f9578535bda4579ebdc7659681cba31f5.
//
// Solidity: event SIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes sid, bytes userPK)
func (_AAAContract *AAAContractFilterer) ParseSIDEncryptionRequested(log types.Log) (*AAAContractSIDEncryptionRequested, error) {
	event := new(AAAContractSIDEncryptionRequested)
	if err := _AAAContract.contract.UnpackLog(event, "SIDEncryptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractSeedPhraseProtocolInitiatedIterator is returned from FilterSeedPhraseProtocolInitiated and is used to iterate over the raw logs and unpacked data for SeedPhraseProtocolInitiated events raised by the AAAContract contract.
type AAAContractSeedPhraseProtocolInitiatedIterator struct {
	Event *AAAContractSeedPhraseProtocolInitiated // Event containing the contract specifics and raw log

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
func (it *AAAContractSeedPhraseProtocolInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractSeedPhraseProtocolInitiated)
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
		it.Event = new(AAAContractSeedPhraseProtocolInitiated)
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
func (it *AAAContractSeedPhraseProtocolInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractSeedPhraseProtocolInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractSeedPhraseProtocolInitiated represents a SeedPhraseProtocolInitiated event raised by the AAAContract contract.
type AAAContractSeedPhraseProtocolInitiated struct {
	Pid [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSeedPhraseProtocolInitiated is a free log retrieval operation binding the contract event 0x9668f7eb50ac3a206e62fc4a82eaec134ba6cc78cc89f656c295469e89e8a7ee.
//
// Solidity: event SeedPhraseProtocolInitiated(bytes32 indexed pid)
func (_AAAContract *AAAContractFilterer) FilterSeedPhraseProtocolInitiated(opts *bind.FilterOpts, pid [][32]byte) (*AAAContractSeedPhraseProtocolInitiatedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "SeedPhraseProtocolInitiated", pidRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractSeedPhraseProtocolInitiatedIterator{contract: _AAAContract.contract, event: "SeedPhraseProtocolInitiated", logs: logs, sub: sub}, nil
}

// WatchSeedPhraseProtocolInitiated is a free log subscription operation binding the contract event 0x9668f7eb50ac3a206e62fc4a82eaec134ba6cc78cc89f656c295469e89e8a7ee.
//
// Solidity: event SeedPhraseProtocolInitiated(bytes32 indexed pid)
func (_AAAContract *AAAContractFilterer) WatchSeedPhraseProtocolInitiated(opts *bind.WatchOpts, sink chan<- *AAAContractSeedPhraseProtocolInitiated, pid [][32]byte) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "SeedPhraseProtocolInitiated", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractSeedPhraseProtocolInitiated)
				if err := _AAAContract.contract.UnpackLog(event, "SeedPhraseProtocolInitiated", log); err != nil {
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

// ParseSeedPhraseProtocolInitiated is a log parse operation binding the contract event 0x9668f7eb50ac3a206e62fc4a82eaec134ba6cc78cc89f656c295469e89e8a7ee.
//
// Solidity: event SeedPhraseProtocolInitiated(bytes32 indexed pid)
func (_AAAContract *AAAContractFilterer) ParseSeedPhraseProtocolInitiated(log types.Log) (*AAAContractSeedPhraseProtocolInitiated, error) {
	event := new(AAAContractSeedPhraseProtocolInitiated)
	if err := _AAAContract.contract.UnpackLog(event, "SeedPhraseProtocolInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractWordRequestedIterator is returned from FilterWordRequested and is used to iterate over the raw logs and unpacked data for WordRequested events raised by the AAAContract contract.
type AAAContractWordRequestedIterator struct {
	Event *AAAContractWordRequested // Event containing the contract specifics and raw log

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
func (it *AAAContractWordRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractWordRequested)
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
		it.Event = new(AAAContractWordRequested)
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
func (it *AAAContractWordRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractWordRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractWordRequested represents a WordRequested event raised by the AAAContract contract.
type AAAContractWordRequested struct {
	Pid    [32]byte
	Node   common.Address
	UserPK []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWordRequested is a free log retrieval operation binding the contract event 0x3354c6cdc2d692d5c1229f7955ac412bd1990d76e481b696603e2ec30f4b2b48.
//
// Solidity: event WordRequested(bytes32 indexed pid, address indexed node, bytes userPK)
func (_AAAContract *AAAContractFilterer) FilterWordRequested(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAAContractWordRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "WordRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractWordRequestedIterator{contract: _AAAContract.contract, event: "WordRequested", logs: logs, sub: sub}, nil
}

// WatchWordRequested is a free log subscription operation binding the contract event 0x3354c6cdc2d692d5c1229f7955ac412bd1990d76e481b696603e2ec30f4b2b48.
//
// Solidity: event WordRequested(bytes32 indexed pid, address indexed node, bytes userPK)
func (_AAAContract *AAAContractFilterer) WatchWordRequested(opts *bind.WatchOpts, sink chan<- *AAAContractWordRequested, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "WordRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractWordRequested)
				if err := _AAAContract.contract.UnpackLog(event, "WordRequested", log); err != nil {
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

// ParseWordRequested is a log parse operation binding the contract event 0x3354c6cdc2d692d5c1229f7955ac412bd1990d76e481b696603e2ec30f4b2b48.
//
// Solidity: event WordRequested(bytes32 indexed pid, address indexed node, bytes userPK)
func (_AAAContract *AAAContractFilterer) ParseWordRequested(log types.Log) (*AAAContractWordRequested, error) {
	event := new(AAAContractWordRequested)
	if err := _AAAContract.contract.UnpackLog(event, "WordRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAContractWordSubmittedIterator is returned from FilterWordSubmitted and is used to iterate over the raw logs and unpacked data for WordSubmitted events raised by the AAAContract contract.
type AAAContractWordSubmittedIterator struct {
	Event *AAAContractWordSubmitted // Event containing the contract specifics and raw log

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
func (it *AAAContractWordSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAContractWordSubmitted)
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
		it.Event = new(AAAContractWordSubmitted)
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
func (it *AAAContractWordSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAContractWordSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAContractWordSubmitted represents a WordSubmitted event raised by the AAAContract contract.
type AAAContractWordSubmitted struct {
	Pid      [32]byte
	Node     common.Address
	WordHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWordSubmitted is a free log retrieval operation binding the contract event 0xd5f0fe48f0da0cd8acd3d0c9c62dff193e60ca8d2bc0344c6f1d03059115c251.
//
// Solidity: event WordSubmitted(bytes32 indexed pid, address indexed node, bytes32 wordHash)
func (_AAAContract *AAAContractFilterer) FilterWordSubmitted(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAAContractWordSubmittedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.FilterLogs(opts, "WordSubmitted", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAContractWordSubmittedIterator{contract: _AAAContract.contract, event: "WordSubmitted", logs: logs, sub: sub}, nil
}

// WatchWordSubmitted is a free log subscription operation binding the contract event 0xd5f0fe48f0da0cd8acd3d0c9c62dff193e60ca8d2bc0344c6f1d03059115c251.
//
// Solidity: event WordSubmitted(bytes32 indexed pid, address indexed node, bytes32 wordHash)
func (_AAAContract *AAAContractFilterer) WatchWordSubmitted(opts *bind.WatchOpts, sink chan<- *AAAContractWordSubmitted, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAAContract.contract.WatchLogs(opts, "WordSubmitted", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAContractWordSubmitted)
				if err := _AAAContract.contract.UnpackLog(event, "WordSubmitted", log); err != nil {
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

// ParseWordSubmitted is a log parse operation binding the contract event 0xd5f0fe48f0da0cd8acd3d0c9c62dff193e60ca8d2bc0344c6f1d03059115c251.
//
// Solidity: event WordSubmitted(bytes32 indexed pid, address indexed node, bytes32 wordHash)
func (_AAAContract *AAAContractFilterer) ParseWordSubmitted(log types.Log) (*AAAContractWordSubmitted, error) {
	event := new(AAAContractWordSubmitted)
	if err := _AAAContract.contract.UnpackLog(event, "WordSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
