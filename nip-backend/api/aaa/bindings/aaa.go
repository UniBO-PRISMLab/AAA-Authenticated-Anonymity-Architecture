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

// AAAMetaData contains all meta data concerning the AAA contract.
var AAAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nodes\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"words\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redundancyFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"symK\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sid\",\"type\":\"bytes32\"}],\"name\":\"PIDEncryptionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encSID\",\"type\":\"bytes\"}],\"name\":\"PhraseComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hashedWord\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toNode\",\"type\":\"address\"}],\"name\":\"RedundantWordRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"wordHash\",\"type\":\"bytes32\"}],\"name\":\"RedundantWordSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"userPK\",\"type\":\"bytes\"}],\"name\":\"SIDEncryptionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"SeedPhraseProtocolInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"userPK\",\"type\":\"bytes\"}],\"name\":\"WordRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"wordHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"WordSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REDUNDANCY_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WORDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"getPhrase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"started\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"pk\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"encWords\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRedundantWords\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"words\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"nodePKs\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pkHash\",\"type\":\"bytes32\"}],\"name\":\"getSACRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"getSID\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sid\",\"type\":\"bytes32\"}],\"name\":\"getSIDRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"getSelectedNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"}],\"name\":\"getWords\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"words\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redundantNodesByPID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"removeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sac\",\"type\":\"bytes\"}],\"name\":\"sacExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pk\",\"type\":\"bytes\"}],\"name\":\"seedPhraseGenerationProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"selectedNodesByPID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encPID\",\"type\":\"bytes\"}],\"name\":\"submitEncryptedPID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encSID\",\"type\":\"bytes\"}],\"name\":\"submitEncryptedSID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encryptedWord\",\"type\":\"bytes\"}],\"name\":\"submitEncryptedWord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encryptedWord\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"wordIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodePK\",\"type\":\"bytes\"}],\"name\":\"submitRedundantWord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sac\",\"type\":\"bytes\"}],\"name\":\"submitSAC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sac\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"pkHash\",\"type\":\"bytes32\"}],\"name\":\"submitSACRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AAAABI is the input ABI used to generate the binding from.
// Deprecated: Use AAAMetaData.ABI instead.
var AAAABI = AAAMetaData.ABI

// AAA is an auto generated Go binding around an Ethereum contract.
type AAA struct {
	AAACaller     // Read-only binding to the contract
	AAATransactor // Write-only binding to the contract
	AAAFilterer   // Log filterer for contract events
}

// AAACaller is an auto generated read-only Go binding around an Ethereum contract.
type AAACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAATransactor is an auto generated write-only Go binding around an Ethereum contract.
type AAATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AAAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AAASession struct {
	Contract     *AAA              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AAACallerSession struct {
	Contract *AAACaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AAATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AAATransactorSession struct {
	Contract     *AAATransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAARaw is an auto generated low-level Go binding around an Ethereum contract.
type AAARaw struct {
	Contract *AAA // Generic contract binding to access the raw methods on
}

// AAACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AAACallerRaw struct {
	Contract *AAACaller // Generic read-only contract binding to access the raw methods on
}

// AAATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AAATransactorRaw struct {
	Contract *AAATransactor // Generic write-only contract binding to access the raw methods on
}

// NewAAA creates a new instance of AAA, bound to a specific deployed contract.
func NewAAA(address common.Address, backend bind.ContractBackend) (*AAA, error) {
	contract, err := bindAAA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AAA{AAACaller: AAACaller{contract: contract}, AAATransactor: AAATransactor{contract: contract}, AAAFilterer: AAAFilterer{contract: contract}}, nil
}

// NewAAACaller creates a new read-only instance of AAA, bound to a specific deployed contract.
func NewAAACaller(address common.Address, caller bind.ContractCaller) (*AAACaller, error) {
	contract, err := bindAAA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AAACaller{contract: contract}, nil
}

// NewAAATransactor creates a new write-only instance of AAA, bound to a specific deployed contract.
func NewAAATransactor(address common.Address, transactor bind.ContractTransactor) (*AAATransactor, error) {
	contract, err := bindAAA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AAATransactor{contract: contract}, nil
}

// NewAAAFilterer creates a new log filterer instance of AAA, bound to a specific deployed contract.
func NewAAAFilterer(address common.Address, filterer bind.ContractFilterer) (*AAAFilterer, error) {
	contract, err := bindAAA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AAAFilterer{contract: contract}, nil
}

// bindAAA binds a generic wrapper to an already deployed contract.
func bindAAA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AAAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAA *AAARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAA.Contract.AAACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAA *AAARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAA.Contract.AAATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAA *AAARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAA.Contract.AAATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAA *AAACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAA *AAATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAA *AAATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAA.Contract.contract.Transact(opts, method, params...)
}

// REDUNDANCYFACTOR is a free data retrieval call binding the contract method 0x288a1319.
//
// Solidity: function REDUNDANCY_FACTOR() view returns(uint256)
func (_AAA *AAACaller) REDUNDANCYFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "REDUNDANCY_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDUNDANCYFACTOR is a free data retrieval call binding the contract method 0x288a1319.
//
// Solidity: function REDUNDANCY_FACTOR() view returns(uint256)
func (_AAA *AAASession) REDUNDANCYFACTOR() (*big.Int, error) {
	return _AAA.Contract.REDUNDANCYFACTOR(&_AAA.CallOpts)
}

// REDUNDANCYFACTOR is a free data retrieval call binding the contract method 0x288a1319.
//
// Solidity: function REDUNDANCY_FACTOR() view returns(uint256)
func (_AAA *AAACallerSession) REDUNDANCYFACTOR() (*big.Int, error) {
	return _AAA.Contract.REDUNDANCYFACTOR(&_AAA.CallOpts)
}

// WORDS is a free data retrieval call binding the contract method 0x4ecb3a31.
//
// Solidity: function WORDS() view returns(uint256)
func (_AAA *AAACaller) WORDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "WORDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WORDS is a free data retrieval call binding the contract method 0x4ecb3a31.
//
// Solidity: function WORDS() view returns(uint256)
func (_AAA *AAASession) WORDS() (*big.Int, error) {
	return _AAA.Contract.WORDS(&_AAA.CallOpts)
}

// WORDS is a free data retrieval call binding the contract method 0x4ecb3a31.
//
// Solidity: function WORDS() view returns(uint256)
func (_AAA *AAACallerSession) WORDS() (*big.Int, error) {
	return _AAA.Contract.WORDS(&_AAA.CallOpts)
}

// GetPhrase is a free data retrieval call binding the contract method 0xcef2ccc9.
//
// Solidity: function getPhrase(bytes32 pid) view returns(bool started, bytes pk, bytes[] encWords)
func (_AAA *AAACaller) GetPhrase(opts *bind.CallOpts, pid [32]byte) (struct {
	Started  bool
	Pk       []byte
	EncWords [][]byte
}, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "getPhrase", pid)

	outstruct := new(struct {
		Started  bool
		Pk       []byte
		EncWords [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Started = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Pk = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.EncWords = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// GetPhrase is a free data retrieval call binding the contract method 0xcef2ccc9.
//
// Solidity: function getPhrase(bytes32 pid) view returns(bool started, bytes pk, bytes[] encWords)
func (_AAA *AAASession) GetPhrase(pid [32]byte) (struct {
	Started  bool
	Pk       []byte
	EncWords [][]byte
}, error) {
	return _AAA.Contract.GetPhrase(&_AAA.CallOpts, pid)
}

// GetPhrase is a free data retrieval call binding the contract method 0xcef2ccc9.
//
// Solidity: function getPhrase(bytes32 pid) view returns(bool started, bytes pk, bytes[] encWords)
func (_AAA *AAACallerSession) GetPhrase(pid [32]byte) (struct {
	Started  bool
	Pk       []byte
	EncWords [][]byte
}, error) {
	return _AAA.Contract.GetPhrase(&_AAA.CallOpts, pid)
}

// GetRedundantWords is a free data retrieval call binding the contract method 0x934ed944.
//
// Solidity: function getRedundantWords(bytes32 pid, uint256 index) view returns(bytes[] words, bytes[] nodePKs)
func (_AAA *AAACaller) GetRedundantWords(opts *bind.CallOpts, pid [32]byte, index *big.Int) (struct {
	Words   [][]byte
	NodePKs [][]byte
}, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "getRedundantWords", pid, index)

	outstruct := new(struct {
		Words   [][]byte
		NodePKs [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Words = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.NodePKs = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// GetRedundantWords is a free data retrieval call binding the contract method 0x934ed944.
//
// Solidity: function getRedundantWords(bytes32 pid, uint256 index) view returns(bytes[] words, bytes[] nodePKs)
func (_AAA *AAASession) GetRedundantWords(pid [32]byte, index *big.Int) (struct {
	Words   [][]byte
	NodePKs [][]byte
}, error) {
	return _AAA.Contract.GetRedundantWords(&_AAA.CallOpts, pid, index)
}

// GetRedundantWords is a free data retrieval call binding the contract method 0x934ed944.
//
// Solidity: function getRedundantWords(bytes32 pid, uint256 index) view returns(bytes[] words, bytes[] nodePKs)
func (_AAA *AAACallerSession) GetRedundantWords(pid [32]byte, index *big.Int) (struct {
	Words   [][]byte
	NodePKs [][]byte
}, error) {
	return _AAA.Contract.GetRedundantWords(&_AAA.CallOpts, pid, index)
}

// GetSACRecord is a free data retrieval call binding the contract method 0xd1a83232.
//
// Solidity: function getSACRecord(bytes32 pkHash) view returns(bytes)
func (_AAA *AAACaller) GetSACRecord(opts *bind.CallOpts, pkHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "getSACRecord", pkHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSACRecord is a free data retrieval call binding the contract method 0xd1a83232.
//
// Solidity: function getSACRecord(bytes32 pkHash) view returns(bytes)
func (_AAA *AAASession) GetSACRecord(pkHash [32]byte) ([]byte, error) {
	return _AAA.Contract.GetSACRecord(&_AAA.CallOpts, pkHash)
}

// GetSACRecord is a free data retrieval call binding the contract method 0xd1a83232.
//
// Solidity: function getSACRecord(bytes32 pkHash) view returns(bytes)
func (_AAA *AAACallerSession) GetSACRecord(pkHash [32]byte) ([]byte, error) {
	return _AAA.Contract.GetSACRecord(&_AAA.CallOpts, pkHash)
}

// GetSID is a free data retrieval call binding the contract method 0xb7ed02d3.
//
// Solidity: function getSID(bytes32 pid) view returns(bytes)
func (_AAA *AAACaller) GetSID(opts *bind.CallOpts, pid [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "getSID", pid)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSID is a free data retrieval call binding the contract method 0xb7ed02d3.
//
// Solidity: function getSID(bytes32 pid) view returns(bytes)
func (_AAA *AAASession) GetSID(pid [32]byte) ([]byte, error) {
	return _AAA.Contract.GetSID(&_AAA.CallOpts, pid)
}

// GetSID is a free data retrieval call binding the contract method 0xb7ed02d3.
//
// Solidity: function getSID(bytes32 pid) view returns(bytes)
func (_AAA *AAACallerSession) GetSID(pid [32]byte) ([]byte, error) {
	return _AAA.Contract.GetSID(&_AAA.CallOpts, pid)
}

// GetSIDRecord is a free data retrieval call binding the contract method 0x239db46f.
//
// Solidity: function getSIDRecord(bytes32 sid) view returns(bytes, bytes)
func (_AAA *AAACaller) GetSIDRecord(opts *bind.CallOpts, sid [32]byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "getSIDRecord", sid)

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
func (_AAA *AAASession) GetSIDRecord(sid [32]byte) ([]byte, []byte, error) {
	return _AAA.Contract.GetSIDRecord(&_AAA.CallOpts, sid)
}

// GetSIDRecord is a free data retrieval call binding the contract method 0x239db46f.
//
// Solidity: function getSIDRecord(bytes32 sid) view returns(bytes, bytes)
func (_AAA *AAACallerSession) GetSIDRecord(sid [32]byte) ([]byte, []byte, error) {
	return _AAA.Contract.GetSIDRecord(&_AAA.CallOpts, sid)
}

// GetSelectedNodes is a free data retrieval call binding the contract method 0xc7b4d5a3.
//
// Solidity: function getSelectedNodes(bytes32 pid) view returns(address[])
func (_AAA *AAACaller) GetSelectedNodes(opts *bind.CallOpts, pid [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "getSelectedNodes", pid)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSelectedNodes is a free data retrieval call binding the contract method 0xc7b4d5a3.
//
// Solidity: function getSelectedNodes(bytes32 pid) view returns(address[])
func (_AAA *AAASession) GetSelectedNodes(pid [32]byte) ([]common.Address, error) {
	return _AAA.Contract.GetSelectedNodes(&_AAA.CallOpts, pid)
}

// GetSelectedNodes is a free data retrieval call binding the contract method 0xc7b4d5a3.
//
// Solidity: function getSelectedNodes(bytes32 pid) view returns(address[])
func (_AAA *AAACallerSession) GetSelectedNodes(pid [32]byte) ([]common.Address, error) {
	return _AAA.Contract.GetSelectedNodes(&_AAA.CallOpts, pid)
}

// GetWords is a free data retrieval call binding the contract method 0x72f43639.
//
// Solidity: function getWords(bytes32 pid) view returns(bytes[] words)
func (_AAA *AAACaller) GetWords(opts *bind.CallOpts, pid [32]byte) ([][]byte, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "getWords", pid)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetWords is a free data retrieval call binding the contract method 0x72f43639.
//
// Solidity: function getWords(bytes32 pid) view returns(bytes[] words)
func (_AAA *AAASession) GetWords(pid [32]byte) ([][]byte, error) {
	return _AAA.Contract.GetWords(&_AAA.CallOpts, pid)
}

// GetWords is a free data retrieval call binding the contract method 0x72f43639.
//
// Solidity: function getWords(bytes32 pid) view returns(bytes[] words)
func (_AAA *AAACallerSession) GetWords(pid [32]byte) ([][]byte, error) {
	return _AAA.Contract.GetWords(&_AAA.CallOpts, pid)
}

// IsNode is a free data retrieval call binding the contract method 0x01750152.
//
// Solidity: function isNode(address ) view returns(bool)
func (_AAA *AAACaller) IsNode(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "isNode", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNode is a free data retrieval call binding the contract method 0x01750152.
//
// Solidity: function isNode(address ) view returns(bool)
func (_AAA *AAASession) IsNode(arg0 common.Address) (bool, error) {
	return _AAA.Contract.IsNode(&_AAA.CallOpts, arg0)
}

// IsNode is a free data retrieval call binding the contract method 0x01750152.
//
// Solidity: function isNode(address ) view returns(bool)
func (_AAA *AAACallerSession) IsNode(arg0 common.Address) (bool, error) {
	return _AAA.Contract.IsNode(&_AAA.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_AAA *AAACaller) NodeList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "nodeList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_AAA *AAASession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _AAA.Contract.NodeList(&_AAA.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_AAA *AAACallerSession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _AAA.Contract.NodeList(&_AAA.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAA *AAACaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAA *AAASession) Owner() (common.Address, error) {
	return _AAA.Contract.Owner(&_AAA.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAA *AAACallerSession) Owner() (common.Address, error) {
	return _AAA.Contract.Owner(&_AAA.CallOpts)
}

// RedundantNodesByPID is a free data retrieval call binding the contract method 0x5e1616d2.
//
// Solidity: function redundantNodesByPID(bytes32 , uint256 , uint256 ) view returns(address)
func (_AAA *AAACaller) RedundantNodesByPID(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "redundantNodesByPID", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedundantNodesByPID is a free data retrieval call binding the contract method 0x5e1616d2.
//
// Solidity: function redundantNodesByPID(bytes32 , uint256 , uint256 ) view returns(address)
func (_AAA *AAASession) RedundantNodesByPID(arg0 [32]byte, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _AAA.Contract.RedundantNodesByPID(&_AAA.CallOpts, arg0, arg1, arg2)
}

// RedundantNodesByPID is a free data retrieval call binding the contract method 0x5e1616d2.
//
// Solidity: function redundantNodesByPID(bytes32 , uint256 , uint256 ) view returns(address)
func (_AAA *AAACallerSession) RedundantNodesByPID(arg0 [32]byte, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _AAA.Contract.RedundantNodesByPID(&_AAA.CallOpts, arg0, arg1, arg2)
}

// SacExists is a free data retrieval call binding the contract method 0x78d14d8a.
//
// Solidity: function sacExists(bytes sac) view returns(bool)
func (_AAA *AAACaller) SacExists(opts *bind.CallOpts, sac []byte) (bool, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "sacExists", sac)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SacExists is a free data retrieval call binding the contract method 0x78d14d8a.
//
// Solidity: function sacExists(bytes sac) view returns(bool)
func (_AAA *AAASession) SacExists(sac []byte) (bool, error) {
	return _AAA.Contract.SacExists(&_AAA.CallOpts, sac)
}

// SacExists is a free data retrieval call binding the contract method 0x78d14d8a.
//
// Solidity: function sacExists(bytes sac) view returns(bool)
func (_AAA *AAACallerSession) SacExists(sac []byte) (bool, error) {
	return _AAA.Contract.SacExists(&_AAA.CallOpts, sac)
}

// SelectedNodesByPID is a free data retrieval call binding the contract method 0x6a978a99.
//
// Solidity: function selectedNodesByPID(bytes32 , uint256 ) view returns(address)
func (_AAA *AAACaller) SelectedNodesByPID(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AAA.contract.Call(opts, &out, "selectedNodesByPID", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SelectedNodesByPID is a free data retrieval call binding the contract method 0x6a978a99.
//
// Solidity: function selectedNodesByPID(bytes32 , uint256 ) view returns(address)
func (_AAA *AAASession) SelectedNodesByPID(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AAA.Contract.SelectedNodesByPID(&_AAA.CallOpts, arg0, arg1)
}

// SelectedNodesByPID is a free data retrieval call binding the contract method 0x6a978a99.
//
// Solidity: function selectedNodesByPID(bytes32 , uint256 ) view returns(address)
func (_AAA *AAACallerSession) SelectedNodesByPID(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AAA.Contract.SelectedNodesByPID(&_AAA.CallOpts, arg0, arg1)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address node) returns()
func (_AAA *AAATransactor) AddNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "addNode", node)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address node) returns()
func (_AAA *AAASession) AddNode(node common.Address) (*types.Transaction, error) {
	return _AAA.Contract.AddNode(&_AAA.TransactOpts, node)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address node) returns()
func (_AAA *AAATransactorSession) AddNode(node common.Address) (*types.Transaction, error) {
	return _AAA.Contract.AddNode(&_AAA.TransactOpts, node)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address node) returns()
func (_AAA *AAATransactor) RemoveNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "removeNode", node)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address node) returns()
func (_AAA *AAASession) RemoveNode(node common.Address) (*types.Transaction, error) {
	return _AAA.Contract.RemoveNode(&_AAA.TransactOpts, node)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address node) returns()
func (_AAA *AAATransactorSession) RemoveNode(node common.Address) (*types.Transaction, error) {
	return _AAA.Contract.RemoveNode(&_AAA.TransactOpts, node)
}

// SeedPhraseGenerationProtocol is a paid mutator transaction binding the contract method 0x3b33892a.
//
// Solidity: function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) returns()
func (_AAA *AAATransactor) SeedPhraseGenerationProtocol(opts *bind.TransactOpts, pid [32]byte, pk []byte) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "seedPhraseGenerationProtocol", pid, pk)
}

// SeedPhraseGenerationProtocol is a paid mutator transaction binding the contract method 0x3b33892a.
//
// Solidity: function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) returns()
func (_AAA *AAASession) SeedPhraseGenerationProtocol(pid [32]byte, pk []byte) (*types.Transaction, error) {
	return _AAA.Contract.SeedPhraseGenerationProtocol(&_AAA.TransactOpts, pid, pk)
}

// SeedPhraseGenerationProtocol is a paid mutator transaction binding the contract method 0x3b33892a.
//
// Solidity: function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) returns()
func (_AAA *AAATransactorSession) SeedPhraseGenerationProtocol(pid [32]byte, pk []byte) (*types.Transaction, error) {
	return _AAA.Contract.SeedPhraseGenerationProtocol(&_AAA.TransactOpts, pid, pk)
}

// SubmitEncryptedPID is a paid mutator transaction binding the contract method 0x00e12e3a.
//
// Solidity: function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) returns()
func (_AAA *AAATransactor) SubmitEncryptedPID(opts *bind.TransactOpts, pid [32]byte, sid [32]byte, encPID []byte) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "submitEncryptedPID", pid, sid, encPID)
}

// SubmitEncryptedPID is a paid mutator transaction binding the contract method 0x00e12e3a.
//
// Solidity: function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) returns()
func (_AAA *AAASession) SubmitEncryptedPID(pid [32]byte, sid [32]byte, encPID []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitEncryptedPID(&_AAA.TransactOpts, pid, sid, encPID)
}

// SubmitEncryptedPID is a paid mutator transaction binding the contract method 0x00e12e3a.
//
// Solidity: function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) returns()
func (_AAA *AAATransactorSession) SubmitEncryptedPID(pid [32]byte, sid [32]byte, encPID []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitEncryptedPID(&_AAA.TransactOpts, pid, sid, encPID)
}

// SubmitEncryptedSID is a paid mutator transaction binding the contract method 0x4542552c.
//
// Solidity: function submitEncryptedSID(bytes32 pid, bytes encSID) returns()
func (_AAA *AAATransactor) SubmitEncryptedSID(opts *bind.TransactOpts, pid [32]byte, encSID []byte) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "submitEncryptedSID", pid, encSID)
}

// SubmitEncryptedSID is a paid mutator transaction binding the contract method 0x4542552c.
//
// Solidity: function submitEncryptedSID(bytes32 pid, bytes encSID) returns()
func (_AAA *AAASession) SubmitEncryptedSID(pid [32]byte, encSID []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitEncryptedSID(&_AAA.TransactOpts, pid, encSID)
}

// SubmitEncryptedSID is a paid mutator transaction binding the contract method 0x4542552c.
//
// Solidity: function submitEncryptedSID(bytes32 pid, bytes encSID) returns()
func (_AAA *AAATransactorSession) SubmitEncryptedSID(pid [32]byte, encSID []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitEncryptedSID(&_AAA.TransactOpts, pid, encSID)
}

// SubmitEncryptedWord is a paid mutator transaction binding the contract method 0xb14e798b.
//
// Solidity: function submitEncryptedWord(bytes32 pid, bytes encryptedWord) returns()
func (_AAA *AAATransactor) SubmitEncryptedWord(opts *bind.TransactOpts, pid [32]byte, encryptedWord []byte) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "submitEncryptedWord", pid, encryptedWord)
}

// SubmitEncryptedWord is a paid mutator transaction binding the contract method 0xb14e798b.
//
// Solidity: function submitEncryptedWord(bytes32 pid, bytes encryptedWord) returns()
func (_AAA *AAASession) SubmitEncryptedWord(pid [32]byte, encryptedWord []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitEncryptedWord(&_AAA.TransactOpts, pid, encryptedWord)
}

// SubmitEncryptedWord is a paid mutator transaction binding the contract method 0xb14e798b.
//
// Solidity: function submitEncryptedWord(bytes32 pid, bytes encryptedWord) returns()
func (_AAA *AAATransactorSession) SubmitEncryptedWord(pid [32]byte, encryptedWord []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitEncryptedWord(&_AAA.TransactOpts, pid, encryptedWord)
}

// SubmitRedundantWord is a paid mutator transaction binding the contract method 0xcef0ceba.
//
// Solidity: function submitRedundantWord(bytes32 pid, bytes encryptedWord, uint256 wordIndex, bytes nodePK) returns()
func (_AAA *AAATransactor) SubmitRedundantWord(opts *bind.TransactOpts, pid [32]byte, encryptedWord []byte, wordIndex *big.Int, nodePK []byte) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "submitRedundantWord", pid, encryptedWord, wordIndex, nodePK)
}

// SubmitRedundantWord is a paid mutator transaction binding the contract method 0xcef0ceba.
//
// Solidity: function submitRedundantWord(bytes32 pid, bytes encryptedWord, uint256 wordIndex, bytes nodePK) returns()
func (_AAA *AAASession) SubmitRedundantWord(pid [32]byte, encryptedWord []byte, wordIndex *big.Int, nodePK []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitRedundantWord(&_AAA.TransactOpts, pid, encryptedWord, wordIndex, nodePK)
}

// SubmitRedundantWord is a paid mutator transaction binding the contract method 0xcef0ceba.
//
// Solidity: function submitRedundantWord(bytes32 pid, bytes encryptedWord, uint256 wordIndex, bytes nodePK) returns()
func (_AAA *AAATransactorSession) SubmitRedundantWord(pid [32]byte, encryptedWord []byte, wordIndex *big.Int, nodePK []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitRedundantWord(&_AAA.TransactOpts, pid, encryptedWord, wordIndex, nodePK)
}

// SubmitSAC is a paid mutator transaction binding the contract method 0x7491b589.
//
// Solidity: function submitSAC(bytes sac) returns()
func (_AAA *AAATransactor) SubmitSAC(opts *bind.TransactOpts, sac []byte) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "submitSAC", sac)
}

// SubmitSAC is a paid mutator transaction binding the contract method 0x7491b589.
//
// Solidity: function submitSAC(bytes sac) returns()
func (_AAA *AAASession) SubmitSAC(sac []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitSAC(&_AAA.TransactOpts, sac)
}

// SubmitSAC is a paid mutator transaction binding the contract method 0x7491b589.
//
// Solidity: function submitSAC(bytes sac) returns()
func (_AAA *AAATransactorSession) SubmitSAC(sac []byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitSAC(&_AAA.TransactOpts, sac)
}

// SubmitSACRecord is a paid mutator transaction binding the contract method 0x9516a01c.
//
// Solidity: function submitSACRecord(bytes sac, bytes32 pkHash) returns()
func (_AAA *AAATransactor) SubmitSACRecord(opts *bind.TransactOpts, sac []byte, pkHash [32]byte) (*types.Transaction, error) {
	return _AAA.contract.Transact(opts, "submitSACRecord", sac, pkHash)
}

// SubmitSACRecord is a paid mutator transaction binding the contract method 0x9516a01c.
//
// Solidity: function submitSACRecord(bytes sac, bytes32 pkHash) returns()
func (_AAA *AAASession) SubmitSACRecord(sac []byte, pkHash [32]byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitSACRecord(&_AAA.TransactOpts, sac, pkHash)
}

// SubmitSACRecord is a paid mutator transaction binding the contract method 0x9516a01c.
//
// Solidity: function submitSACRecord(bytes sac, bytes32 pkHash) returns()
func (_AAA *AAATransactorSession) SubmitSACRecord(sac []byte, pkHash [32]byte) (*types.Transaction, error) {
	return _AAA.Contract.SubmitSACRecord(&_AAA.TransactOpts, sac, pkHash)
}

// AAANodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the AAA contract.
type AAANodeAddedIterator struct {
	Event *AAANodeAdded // Event containing the contract specifics and raw log

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
func (it *AAANodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAANodeAdded)
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
		it.Event = new(AAANodeAdded)
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
func (it *AAANodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAANodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAANodeAdded represents a NodeAdded event raised by the AAA contract.
type AAANodeAdded struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address node)
func (_AAA *AAAFilterer) FilterNodeAdded(opts *bind.FilterOpts) (*AAANodeAddedIterator, error) {

	logs, sub, err := _AAA.contract.FilterLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return &AAANodeAddedIterator{contract: _AAA.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address node)
func (_AAA *AAAFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *AAANodeAdded) (event.Subscription, error) {

	logs, sub, err := _AAA.contract.WatchLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAANodeAdded)
				if err := _AAA.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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
func (_AAA *AAAFilterer) ParseNodeAdded(log types.Log) (*AAANodeAdded, error) {
	event := new(AAANodeAdded)
	if err := _AAA.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAANodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the AAA contract.
type AAANodeRemovedIterator struct {
	Event *AAANodeRemoved // Event containing the contract specifics and raw log

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
func (it *AAANodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAANodeRemoved)
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
		it.Event = new(AAANodeRemoved)
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
func (it *AAANodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAANodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAANodeRemoved represents a NodeRemoved event raised by the AAA contract.
type AAANodeRemoved struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address node)
func (_AAA *AAAFilterer) FilterNodeRemoved(opts *bind.FilterOpts) (*AAANodeRemovedIterator, error) {

	logs, sub, err := _AAA.contract.FilterLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return &AAANodeRemovedIterator{contract: _AAA.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address node)
func (_AAA *AAAFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *AAANodeRemoved) (event.Subscription, error) {

	logs, sub, err := _AAA.contract.WatchLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAANodeRemoved)
				if err := _AAA.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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
func (_AAA *AAAFilterer) ParseNodeRemoved(log types.Log) (*AAANodeRemoved, error) {
	event := new(AAANodeRemoved)
	if err := _AAA.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAPIDEncryptionRequestedIterator is returned from FilterPIDEncryptionRequested and is used to iterate over the raw logs and unpacked data for PIDEncryptionRequested events raised by the AAA contract.
type AAAPIDEncryptionRequestedIterator struct {
	Event *AAAPIDEncryptionRequested // Event containing the contract specifics and raw log

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
func (it *AAAPIDEncryptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAPIDEncryptionRequested)
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
		it.Event = new(AAAPIDEncryptionRequested)
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
func (it *AAAPIDEncryptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAPIDEncryptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAPIDEncryptionRequested represents a PIDEncryptionRequested event raised by the AAA contract.
type AAAPIDEncryptionRequested struct {
	Pid  [32]byte
	Node common.Address
	SymK [32]byte
	Sid  [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPIDEncryptionRequested is a free log retrieval operation binding the contract event 0x01c8b3ea5e773e12ac7fb27cd0c382713fd140d4db5916ee0c1bce601bdc4d50.
//
// Solidity: event PIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes32 symK, bytes32 sid)
func (_AAA *AAAFilterer) FilterPIDEncryptionRequested(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAAPIDEncryptionRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.FilterLogs(opts, "PIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAPIDEncryptionRequestedIterator{contract: _AAA.contract, event: "PIDEncryptionRequested", logs: logs, sub: sub}, nil
}

// WatchPIDEncryptionRequested is a free log subscription operation binding the contract event 0x01c8b3ea5e773e12ac7fb27cd0c382713fd140d4db5916ee0c1bce601bdc4d50.
//
// Solidity: event PIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes32 symK, bytes32 sid)
func (_AAA *AAAFilterer) WatchPIDEncryptionRequested(opts *bind.WatchOpts, sink chan<- *AAAPIDEncryptionRequested, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.WatchLogs(opts, "PIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAPIDEncryptionRequested)
				if err := _AAA.contract.UnpackLog(event, "PIDEncryptionRequested", log); err != nil {
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
func (_AAA *AAAFilterer) ParsePIDEncryptionRequested(log types.Log) (*AAAPIDEncryptionRequested, error) {
	event := new(AAAPIDEncryptionRequested)
	if err := _AAA.contract.UnpackLog(event, "PIDEncryptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAPhraseCompleteIterator is returned from FilterPhraseComplete and is used to iterate over the raw logs and unpacked data for PhraseComplete events raised by the AAA contract.
type AAAPhraseCompleteIterator struct {
	Event *AAAPhraseComplete // Event containing the contract specifics and raw log

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
func (it *AAAPhraseCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAPhraseComplete)
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
		it.Event = new(AAAPhraseComplete)
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
func (it *AAAPhraseCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAPhraseCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAPhraseComplete represents a PhraseComplete event raised by the AAA contract.
type AAAPhraseComplete struct {
	Pid    [32]byte
	EncSID []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPhraseComplete is a free log retrieval operation binding the contract event 0xa24f64aabeee785513b21a32ea1d294c10dd37c0b1ba653b03e9825bcfa1769f.
//
// Solidity: event PhraseComplete(bytes32 indexed pid, bytes encSID)
func (_AAA *AAAFilterer) FilterPhraseComplete(opts *bind.FilterOpts, pid [][32]byte) (*AAAPhraseCompleteIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAA.contract.FilterLogs(opts, "PhraseComplete", pidRule)
	if err != nil {
		return nil, err
	}
	return &AAAPhraseCompleteIterator{contract: _AAA.contract, event: "PhraseComplete", logs: logs, sub: sub}, nil
}

// WatchPhraseComplete is a free log subscription operation binding the contract event 0xa24f64aabeee785513b21a32ea1d294c10dd37c0b1ba653b03e9825bcfa1769f.
//
// Solidity: event PhraseComplete(bytes32 indexed pid, bytes encSID)
func (_AAA *AAAFilterer) WatchPhraseComplete(opts *bind.WatchOpts, sink chan<- *AAAPhraseComplete, pid [][32]byte) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAA.contract.WatchLogs(opts, "PhraseComplete", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAPhraseComplete)
				if err := _AAA.contract.UnpackLog(event, "PhraseComplete", log); err != nil {
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
func (_AAA *AAAFilterer) ParsePhraseComplete(log types.Log) (*AAAPhraseComplete, error) {
	event := new(AAAPhraseComplete)
	if err := _AAA.contract.UnpackLog(event, "PhraseComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAARedundantWordRequestedIterator is returned from FilterRedundantWordRequested and is used to iterate over the raw logs and unpacked data for RedundantWordRequested events raised by the AAA contract.
type AAARedundantWordRequestedIterator struct {
	Event *AAARedundantWordRequested // Event containing the contract specifics and raw log

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
func (it *AAARedundantWordRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAARedundantWordRequested)
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
		it.Event = new(AAARedundantWordRequested)
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
func (it *AAARedundantWordRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAARedundantWordRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAARedundantWordRequested represents a RedundantWordRequested event raised by the AAA contract.
type AAARedundantWordRequested struct {
	Pid        [32]byte
	Index      *big.Int
	HashedWord [32]byte
	ToNode     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedundantWordRequested is a free log retrieval operation binding the contract event 0x0da8c63ed8c37525265fb127b3b89d28f1fb59d2f08fd85df8dbe9fd667fe89e.
//
// Solidity: event RedundantWordRequested(bytes32 indexed pid, uint256 indexed index, bytes32 hashedWord, address indexed toNode)
func (_AAA *AAAFilterer) FilterRedundantWordRequested(opts *bind.FilterOpts, pid [][32]byte, index []*big.Int, toNode []common.Address) (*AAARedundantWordRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var toNodeRule []interface{}
	for _, toNodeItem := range toNode {
		toNodeRule = append(toNodeRule, toNodeItem)
	}

	logs, sub, err := _AAA.contract.FilterLogs(opts, "RedundantWordRequested", pidRule, indexRule, toNodeRule)
	if err != nil {
		return nil, err
	}
	return &AAARedundantWordRequestedIterator{contract: _AAA.contract, event: "RedundantWordRequested", logs: logs, sub: sub}, nil
}

// WatchRedundantWordRequested is a free log subscription operation binding the contract event 0x0da8c63ed8c37525265fb127b3b89d28f1fb59d2f08fd85df8dbe9fd667fe89e.
//
// Solidity: event RedundantWordRequested(bytes32 indexed pid, uint256 indexed index, bytes32 hashedWord, address indexed toNode)
func (_AAA *AAAFilterer) WatchRedundantWordRequested(opts *bind.WatchOpts, sink chan<- *AAARedundantWordRequested, pid [][32]byte, index []*big.Int, toNode []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var toNodeRule []interface{}
	for _, toNodeItem := range toNode {
		toNodeRule = append(toNodeRule, toNodeItem)
	}

	logs, sub, err := _AAA.contract.WatchLogs(opts, "RedundantWordRequested", pidRule, indexRule, toNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAARedundantWordRequested)
				if err := _AAA.contract.UnpackLog(event, "RedundantWordRequested", log); err != nil {
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

// ParseRedundantWordRequested is a log parse operation binding the contract event 0x0da8c63ed8c37525265fb127b3b89d28f1fb59d2f08fd85df8dbe9fd667fe89e.
//
// Solidity: event RedundantWordRequested(bytes32 indexed pid, uint256 indexed index, bytes32 hashedWord, address indexed toNode)
func (_AAA *AAAFilterer) ParseRedundantWordRequested(log types.Log) (*AAARedundantWordRequested, error) {
	event := new(AAARedundantWordRequested)
	if err := _AAA.contract.UnpackLog(event, "RedundantWordRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAARedundantWordSubmittedIterator is returned from FilterRedundantWordSubmitted and is used to iterate over the raw logs and unpacked data for RedundantWordSubmitted events raised by the AAA contract.
type AAARedundantWordSubmittedIterator struct {
	Event *AAARedundantWordSubmitted // Event containing the contract specifics and raw log

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
func (it *AAARedundantWordSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAARedundantWordSubmitted)
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
		it.Event = new(AAARedundantWordSubmitted)
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
func (it *AAARedundantWordSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAARedundantWordSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAARedundantWordSubmitted represents a RedundantWordSubmitted event raised by the AAA contract.
type AAARedundantWordSubmitted struct {
	Pid      [32]byte
	Index    *big.Int
	Node     common.Address
	WordHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedundantWordSubmitted is a free log retrieval operation binding the contract event 0x2e48f225c7128740b00ba988b438fb1cdd63d84a479c8d821ed3e39e20e37b09.
//
// Solidity: event RedundantWordSubmitted(bytes32 indexed pid, uint256 indexed index, address indexed node, bytes32 wordHash)
func (_AAA *AAAFilterer) FilterRedundantWordSubmitted(opts *bind.FilterOpts, pid [][32]byte, index []*big.Int, node []common.Address) (*AAARedundantWordSubmittedIterator, error) {

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

	logs, sub, err := _AAA.contract.FilterLogs(opts, "RedundantWordSubmitted", pidRule, indexRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAARedundantWordSubmittedIterator{contract: _AAA.contract, event: "RedundantWordSubmitted", logs: logs, sub: sub}, nil
}

// WatchRedundantWordSubmitted is a free log subscription operation binding the contract event 0x2e48f225c7128740b00ba988b438fb1cdd63d84a479c8d821ed3e39e20e37b09.
//
// Solidity: event RedundantWordSubmitted(bytes32 indexed pid, uint256 indexed index, address indexed node, bytes32 wordHash)
func (_AAA *AAAFilterer) WatchRedundantWordSubmitted(opts *bind.WatchOpts, sink chan<- *AAARedundantWordSubmitted, pid [][32]byte, index []*big.Int, node []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AAA.contract.WatchLogs(opts, "RedundantWordSubmitted", pidRule, indexRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAARedundantWordSubmitted)
				if err := _AAA.contract.UnpackLog(event, "RedundantWordSubmitted", log); err != nil {
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
func (_AAA *AAAFilterer) ParseRedundantWordSubmitted(log types.Log) (*AAARedundantWordSubmitted, error) {
	event := new(AAARedundantWordSubmitted)
	if err := _AAA.contract.UnpackLog(event, "RedundantWordSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAASIDEncryptionRequestedIterator is returned from FilterSIDEncryptionRequested and is used to iterate over the raw logs and unpacked data for SIDEncryptionRequested events raised by the AAA contract.
type AAASIDEncryptionRequestedIterator struct {
	Event *AAASIDEncryptionRequested // Event containing the contract specifics and raw log

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
func (it *AAASIDEncryptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAASIDEncryptionRequested)
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
		it.Event = new(AAASIDEncryptionRequested)
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
func (it *AAASIDEncryptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAASIDEncryptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAASIDEncryptionRequested represents a SIDEncryptionRequested event raised by the AAA contract.
type AAASIDEncryptionRequested struct {
	Pid    [32]byte
	Node   common.Address
	Sid    []byte
	UserPK []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSIDEncryptionRequested is a free log retrieval operation binding the contract event 0xe0ea39226dca1f5e473a57d8db19ad9f9578535bda4579ebdc7659681cba31f5.
//
// Solidity: event SIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes sid, bytes userPK)
func (_AAA *AAAFilterer) FilterSIDEncryptionRequested(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAASIDEncryptionRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.FilterLogs(opts, "SIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAASIDEncryptionRequestedIterator{contract: _AAA.contract, event: "SIDEncryptionRequested", logs: logs, sub: sub}, nil
}

// WatchSIDEncryptionRequested is a free log subscription operation binding the contract event 0xe0ea39226dca1f5e473a57d8db19ad9f9578535bda4579ebdc7659681cba31f5.
//
// Solidity: event SIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes sid, bytes userPK)
func (_AAA *AAAFilterer) WatchSIDEncryptionRequested(opts *bind.WatchOpts, sink chan<- *AAASIDEncryptionRequested, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.WatchLogs(opts, "SIDEncryptionRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAASIDEncryptionRequested)
				if err := _AAA.contract.UnpackLog(event, "SIDEncryptionRequested", log); err != nil {
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
func (_AAA *AAAFilterer) ParseSIDEncryptionRequested(log types.Log) (*AAASIDEncryptionRequested, error) {
	event := new(AAASIDEncryptionRequested)
	if err := _AAA.contract.UnpackLog(event, "SIDEncryptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAASeedPhraseProtocolInitiatedIterator is returned from FilterSeedPhraseProtocolInitiated and is used to iterate over the raw logs and unpacked data for SeedPhraseProtocolInitiated events raised by the AAA contract.
type AAASeedPhraseProtocolInitiatedIterator struct {
	Event *AAASeedPhraseProtocolInitiated // Event containing the contract specifics and raw log

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
func (it *AAASeedPhraseProtocolInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAASeedPhraseProtocolInitiated)
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
		it.Event = new(AAASeedPhraseProtocolInitiated)
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
func (it *AAASeedPhraseProtocolInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAASeedPhraseProtocolInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAASeedPhraseProtocolInitiated represents a SeedPhraseProtocolInitiated event raised by the AAA contract.
type AAASeedPhraseProtocolInitiated struct {
	Pid [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSeedPhraseProtocolInitiated is a free log retrieval operation binding the contract event 0x9668f7eb50ac3a206e62fc4a82eaec134ba6cc78cc89f656c295469e89e8a7ee.
//
// Solidity: event SeedPhraseProtocolInitiated(bytes32 indexed pid)
func (_AAA *AAAFilterer) FilterSeedPhraseProtocolInitiated(opts *bind.FilterOpts, pid [][32]byte) (*AAASeedPhraseProtocolInitiatedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAA.contract.FilterLogs(opts, "SeedPhraseProtocolInitiated", pidRule)
	if err != nil {
		return nil, err
	}
	return &AAASeedPhraseProtocolInitiatedIterator{contract: _AAA.contract, event: "SeedPhraseProtocolInitiated", logs: logs, sub: sub}, nil
}

// WatchSeedPhraseProtocolInitiated is a free log subscription operation binding the contract event 0x9668f7eb50ac3a206e62fc4a82eaec134ba6cc78cc89f656c295469e89e8a7ee.
//
// Solidity: event SeedPhraseProtocolInitiated(bytes32 indexed pid)
func (_AAA *AAAFilterer) WatchSeedPhraseProtocolInitiated(opts *bind.WatchOpts, sink chan<- *AAASeedPhraseProtocolInitiated, pid [][32]byte) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AAA.contract.WatchLogs(opts, "SeedPhraseProtocolInitiated", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAASeedPhraseProtocolInitiated)
				if err := _AAA.contract.UnpackLog(event, "SeedPhraseProtocolInitiated", log); err != nil {
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
func (_AAA *AAAFilterer) ParseSeedPhraseProtocolInitiated(log types.Log) (*AAASeedPhraseProtocolInitiated, error) {
	event := new(AAASeedPhraseProtocolInitiated)
	if err := _AAA.contract.UnpackLog(event, "SeedPhraseProtocolInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAWordRequestedIterator is returned from FilterWordRequested and is used to iterate over the raw logs and unpacked data for WordRequested events raised by the AAA contract.
type AAAWordRequestedIterator struct {
	Event *AAAWordRequested // Event containing the contract specifics and raw log

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
func (it *AAAWordRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAWordRequested)
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
		it.Event = new(AAAWordRequested)
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
func (it *AAAWordRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAWordRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAWordRequested represents a WordRequested event raised by the AAA contract.
type AAAWordRequested struct {
	Pid    [32]byte
	Node   common.Address
	UserPK []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWordRequested is a free log retrieval operation binding the contract event 0x3354c6cdc2d692d5c1229f7955ac412bd1990d76e481b696603e2ec30f4b2b48.
//
// Solidity: event WordRequested(bytes32 indexed pid, address indexed node, bytes userPK)
func (_AAA *AAAFilterer) FilterWordRequested(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAAWordRequestedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.FilterLogs(opts, "WordRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAWordRequestedIterator{contract: _AAA.contract, event: "WordRequested", logs: logs, sub: sub}, nil
}

// WatchWordRequested is a free log subscription operation binding the contract event 0x3354c6cdc2d692d5c1229f7955ac412bd1990d76e481b696603e2ec30f4b2b48.
//
// Solidity: event WordRequested(bytes32 indexed pid, address indexed node, bytes userPK)
func (_AAA *AAAFilterer) WatchWordRequested(opts *bind.WatchOpts, sink chan<- *AAAWordRequested, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.WatchLogs(opts, "WordRequested", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAWordRequested)
				if err := _AAA.contract.UnpackLog(event, "WordRequested", log); err != nil {
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
func (_AAA *AAAFilterer) ParseWordRequested(log types.Log) (*AAAWordRequested, error) {
	event := new(AAAWordRequested)
	if err := _AAA.contract.UnpackLog(event, "WordRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAAWordSubmittedIterator is returned from FilterWordSubmitted and is used to iterate over the raw logs and unpacked data for WordSubmitted events raised by the AAA contract.
type AAAWordSubmittedIterator struct {
	Event *AAAWordSubmitted // Event containing the contract specifics and raw log

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
func (it *AAAWordSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAAWordSubmitted)
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
		it.Event = new(AAAWordSubmitted)
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
func (it *AAAWordSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAAWordSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAAWordSubmitted represents a WordSubmitted event raised by the AAA contract.
type AAAWordSubmitted struct {
	Pid      [32]byte
	Node     common.Address
	WordHash [32]byte
	Index    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWordSubmitted is a free log retrieval operation binding the contract event 0x2fe138d30802f64a9d3b26d90b0dae52cfb6d93ea2a79a4e0ed823c0922b1c48.
//
// Solidity: event WordSubmitted(bytes32 indexed pid, address indexed node, bytes32 wordHash, uint256 index)
func (_AAA *AAAFilterer) FilterWordSubmitted(opts *bind.FilterOpts, pid [][32]byte, node []common.Address) (*AAAWordSubmittedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.FilterLogs(opts, "WordSubmitted", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &AAAWordSubmittedIterator{contract: _AAA.contract, event: "WordSubmitted", logs: logs, sub: sub}, nil
}

// WatchWordSubmitted is a free log subscription operation binding the contract event 0x2fe138d30802f64a9d3b26d90b0dae52cfb6d93ea2a79a4e0ed823c0922b1c48.
//
// Solidity: event WordSubmitted(bytes32 indexed pid, address indexed node, bytes32 wordHash, uint256 index)
func (_AAA *AAAFilterer) WatchWordSubmitted(opts *bind.WatchOpts, sink chan<- *AAAWordSubmitted, pid [][32]byte, node []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AAA.contract.WatchLogs(opts, "WordSubmitted", pidRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAAWordSubmitted)
				if err := _AAA.contract.UnpackLog(event, "WordSubmitted", log); err != nil {
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

// ParseWordSubmitted is a log parse operation binding the contract event 0x2fe138d30802f64a9d3b26d90b0dae52cfb6d93ea2a79a4e0ed823c0922b1c48.
//
// Solidity: event WordSubmitted(bytes32 indexed pid, address indexed node, bytes32 wordHash, uint256 index)
func (_AAA *AAAFilterer) ParseWordSubmitted(log types.Log) (*AAAWordSubmitted, error) {
	event := new(AAAWordSubmitted)
	if err := _AAA.contract.UnpackLog(event, "WordSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
