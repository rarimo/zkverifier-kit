// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poseidonsmt

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

// SparseMerkleTreeNode is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeNode struct {
	NodeType   uint8
	ChildLeft  uint64
	ChildRight uint64
	NodeHash   [32]byte
	Key        [32]byte
	Value      [32]byte
}

// SparseMerkleTreeProof is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeProof struct {
	Root         [32]byte
	Siblings     [][32]byte
	Existence    bool
	Key          [32]byte
	Value        [32]byte
	AuxExistence bool
	AuxKey       [32]byte
	AuxValue     [32]byte
}

// PoseidonSMTMetaData contains all meta data concerning the PoseidonSMT contract.
var PoseidonSMTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ROOT_VALIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"treeHeight_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"registration_\",\"type\":\"address\"}],\"name\":\"__PoseidonSMT_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"element_\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getNodeByKey\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSparseMerkleTree.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"childLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"childRight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"auxKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxValue\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootLatest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registration\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newElement_\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoseidonSMTABI is the input ABI used to generate the binding from.
// Deprecated: Use PoseidonSMTMetaData.ABI instead.
var PoseidonSMTABI = PoseidonSMTMetaData.ABI

// PoseidonSMT is an auto generated Go binding around an Ethereum contract.
type PoseidonSMT struct {
	PoseidonSMTCaller     // Read-only binding to the contract
	PoseidonSMTTransactor // Write-only binding to the contract
	PoseidonSMTFilterer   // Log filterer for contract events
}

// PoseidonSMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoseidonSMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoseidonSMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoseidonSMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoseidonSMTSession struct {
	Contract     *PoseidonSMT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoseidonSMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoseidonSMTCallerSession struct {
	Contract *PoseidonSMTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoseidonSMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoseidonSMTTransactorSession struct {
	Contract     *PoseidonSMTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoseidonSMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoseidonSMTRaw struct {
	Contract *PoseidonSMT // Generic contract binding to access the raw methods on
}

// PoseidonSMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoseidonSMTCallerRaw struct {
	Contract *PoseidonSMTCaller // Generic read-only contract binding to access the raw methods on
}

// PoseidonSMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoseidonSMTTransactorRaw struct {
	Contract *PoseidonSMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoseidonSMT creates a new instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMT(address common.Address, backend bind.ContractBackend) (*PoseidonSMT, error) {
	contract, err := bindPoseidonSMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMT{PoseidonSMTCaller: PoseidonSMTCaller{contract: contract}, PoseidonSMTTransactor: PoseidonSMTTransactor{contract: contract}, PoseidonSMTFilterer: PoseidonSMTFilterer{contract: contract}}, nil
}

// NewPoseidonSMTCaller creates a new read-only instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMTCaller(address common.Address, caller bind.ContractCaller) (*PoseidonSMTCaller, error) {
	contract, err := bindPoseidonSMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTCaller{contract: contract}, nil
}

// NewPoseidonSMTTransactor creates a new write-only instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMTTransactor(address common.Address, transactor bind.ContractTransactor) (*PoseidonSMTTransactor, error) {
	contract, err := bindPoseidonSMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTTransactor{contract: contract}, nil
}

// NewPoseidonSMTFilterer creates a new log filterer instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMTFilterer(address common.Address, filterer bind.ContractFilterer) (*PoseidonSMTFilterer, error) {
	contract, err := bindPoseidonSMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTFilterer{contract: contract}, nil
}

// bindPoseidonSMT binds a generic wrapper to an already deployed contract.
func bindPoseidonSMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoseidonSMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoseidonSMT *PoseidonSMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoseidonSMT.Contract.PoseidonSMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoseidonSMT *PoseidonSMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoseidonSMT *PoseidonSMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoseidonSMT *PoseidonSMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoseidonSMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoseidonSMT *PoseidonSMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoseidonSMT *PoseidonSMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.contract.Transact(opts, method, params...)
}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCaller) ROOTVALIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "ROOT_VALIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTSession) ROOTVALIDITY() (*big.Int, error) {
	return _PoseidonSMT.Contract.ROOTVALIDITY(&_PoseidonSMT.CallOpts)
}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCallerSession) ROOTVALIDITY() (*big.Int, error) {
	return _PoseidonSMT.Contract.ROOTVALIDITY(&_PoseidonSMT.CallOpts)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCaller) GetNodeByKey(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeNode, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "getNodeByKey", key_)

	if err != nil {
		return *new(SparseMerkleTreeNode), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeNode)).(*SparseMerkleTreeNode)

	return out0, err

}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _PoseidonSMT.Contract.GetNodeByKey(&_PoseidonSMT.CallOpts, key_)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCallerSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _PoseidonSMT.Contract.GetNodeByKey(&_PoseidonSMT.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCaller) GetProof(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeProof, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "getProof", key_)

	if err != nil {
		return *new(SparseMerkleTreeProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeProof)).(*SparseMerkleTreeProof)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _PoseidonSMT.Contract.GetProof(&_PoseidonSMT.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCallerSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _PoseidonSMT.Contract.GetProof(&_PoseidonSMT.CallOpts, key_)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTSession) GetRoot() ([32]byte, error) {
	return _PoseidonSMT.Contract.GetRoot(&_PoseidonSMT.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTCallerSession) GetRoot() ([32]byte, error) {
	return _PoseidonSMT.Contract.GetRoot(&_PoseidonSMT.CallOpts)
}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCaller) IsRootLatest(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "isRootLatest", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTSession) IsRootLatest(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootLatest(&_PoseidonSMT.CallOpts, root_)
}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCallerSession) IsRootLatest(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootLatest(&_PoseidonSMT.CallOpts, root_)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCaller) IsRootValid(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "isRootValid", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTSession) IsRootValid(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootValid(&_PoseidonSMT.CallOpts, root_)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCallerSession) IsRootValid(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootValid(&_PoseidonSMT.CallOpts, root_)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_PoseidonSMT *PoseidonSMTCaller) Registration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "registration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_PoseidonSMT *PoseidonSMTSession) Registration() (common.Address, error) {
	return _PoseidonSMT.Contract.Registration(&_PoseidonSMT.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_PoseidonSMT *PoseidonSMTCallerSession) Registration() (common.Address, error) {
	return _PoseidonSMT.Contract.Registration(&_PoseidonSMT.CallOpts)
}

// PoseidonSMTInit is a paid mutator transaction binding the contract method 0x146b8412.
//
// Solidity: function __PoseidonSMT_init(uint256 treeHeight_, address registration_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) PoseidonSMTInit(opts *bind.TransactOpts, treeHeight_ *big.Int, registration_ common.Address) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "__PoseidonSMT_init", treeHeight_, registration_)
}

// PoseidonSMTInit is a paid mutator transaction binding the contract method 0x146b8412.
//
// Solidity: function __PoseidonSMT_init(uint256 treeHeight_, address registration_) returns()
func (_PoseidonSMT *PoseidonSMTSession) PoseidonSMTInit(treeHeight_ *big.Int, registration_ common.Address) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTInit(&_PoseidonSMT.TransactOpts, treeHeight_, registration_)
}

// PoseidonSMTInit is a paid mutator transaction binding the contract method 0x146b8412.
//
// Solidity: function __PoseidonSMT_init(uint256 treeHeight_, address registration_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) PoseidonSMTInit(treeHeight_ *big.Int, registration_ common.Address) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTInit(&_PoseidonSMT.TransactOpts, treeHeight_, registration_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) Add(opts *bind.TransactOpts, keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "add", keyOfElement_, element_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_PoseidonSMT *PoseidonSMTSession) Add(keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Add(&_PoseidonSMT.TransactOpts, keyOfElement_, element_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) Add(keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Add(&_PoseidonSMT.TransactOpts, keyOfElement_, element_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) Remove(opts *bind.TransactOpts, keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "remove", keyOfElement_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_PoseidonSMT *PoseidonSMTSession) Remove(keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Remove(&_PoseidonSMT.TransactOpts, keyOfElement_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) Remove(keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Remove(&_PoseidonSMT.TransactOpts, keyOfElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) Update(opts *bind.TransactOpts, keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "update", keyOfElement_, newElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_PoseidonSMT *PoseidonSMTSession) Update(keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Update(&_PoseidonSMT.TransactOpts, keyOfElement_, newElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) Update(keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Update(&_PoseidonSMT.TransactOpts, keyOfElement_, newElement_)
}

// PoseidonSMTInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoseidonSMT contract.
type PoseidonSMTInitializedIterator struct {
	Event *PoseidonSMTInitialized // Event containing the contract specifics and raw log

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
func (it *PoseidonSMTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoseidonSMTInitialized)
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
		it.Event = new(PoseidonSMTInitialized)
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
func (it *PoseidonSMTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoseidonSMTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoseidonSMTInitialized represents a Initialized event raised by the PoseidonSMT contract.
type PoseidonSMTInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoseidonSMT *PoseidonSMTFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoseidonSMTInitializedIterator, error) {

	logs, sub, err := _PoseidonSMT.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTInitializedIterator{contract: _PoseidonSMT.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoseidonSMT *PoseidonSMTFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoseidonSMTInitialized) (event.Subscription, error) {

	logs, sub, err := _PoseidonSMT.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoseidonSMTInitialized)
				if err := _PoseidonSMT.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoseidonSMT *PoseidonSMTFilterer) ParseInitialized(log types.Log) (*PoseidonSMTInitialized, error) {
	event := new(PoseidonSMTInitialized)
	if err := _PoseidonSMT.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
