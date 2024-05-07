// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registration

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

// RegistrationIdentityInfo is an auto generated low-level Go binding around an user-defined struct.
type RegistrationIdentityInfo struct {
	ActivePassport [32]byte
	IssueTimestamp uint64
}

// RegistrationPassport is an auto generated low-level Go binding around an user-defined struct.
type RegistrationPassport struct {
	DataType  [32]byte
	Signature []byte
	PublicKey []byte
}

// RegistrationPassportInfo is an auto generated low-level Go binding around an user-defined struct.
type RegistrationPassportInfo struct {
	ActiveIdentity         [32]byte
	IdentityReissueCounter uint64
}

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

// VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// RegistrationMetaData contains all meta data concerning the Registration contract.
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identityKey\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identityKey\",\"type\":\"bytes32\"}],\"name\":\"ReissuedIdentity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identityKey\",\"type\":\"bytes32\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ICAO_PREFIX\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVOKED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"treeHeight_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"icaoMasterTreeMerkleRoot_\",\"type\":\"bytes32\"}],\"name\":\"__Registration_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dispatcherType_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"dispatcher_\",\"type\":\"address\"}],\"name\":\"addDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"changeICAOMasterTreeRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getNodeByKey\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSparseMerkleTree.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"childLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"childRight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportKey_\",\"type\":\"bytes32\"}],\"name\":\"getPassportInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"activeIdentity\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"identityReissueCounter\",\"type\":\"uint64\"}],\"internalType\":\"structRegistration.PassportInfo\",\"name\":\"passportInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"activePassport\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"issueTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structRegistration.IdentityInfo\",\"name\":\"identityInfo_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"auxKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxValue\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"icaoMasterTreeMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"passportDispatchers\",\"outputs\":[{\"internalType\":\"contractIPassportDispatcher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"reissueIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dispatcherType_\",\"type\":\"bytes32\"}],\"name\":\"removeDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"internalType\":\"structRegistration.Passport\",\"name\":\"passport_\",\"type\":\"tuple\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrationMetaData.ABI instead.
var RegistrationABI = RegistrationMetaData.ABI

// Registration is an auto generated Go binding around an Ethereum contract.
type Registration struct {
	RegistrationCaller     // Read-only binding to the contract
	RegistrationTransactor // Write-only binding to the contract
	RegistrationFilterer   // Log filterer for contract events
}

// RegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrationSession struct {
	Contract     *Registration     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrationCallerSession struct {
	Contract *RegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrationTransactorSession struct {
	Contract     *RegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrationRaw struct {
	Contract *Registration // Generic contract binding to access the raw methods on
}

// RegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrationCallerRaw struct {
	Contract *RegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrationTransactorRaw struct {
	Contract *RegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistration creates a new instance of Registration, bound to a specific deployed contract.
func NewRegistration(address common.Address, backend bind.ContractBackend) (*Registration, error) {
	contract, err := bindRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registration{RegistrationCaller: RegistrationCaller{contract: contract}, RegistrationTransactor: RegistrationTransactor{contract: contract}, RegistrationFilterer: RegistrationFilterer{contract: contract}}, nil
}

// NewRegistrationCaller creates a new read-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationCaller(address common.Address, caller bind.ContractCaller) (*RegistrationCaller, error) {
	contract, err := bindRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationCaller{contract: contract}, nil
}

// NewRegistrationTransactor creates a new write-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationTransactor, error) {
	contract, err := bindRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationTransactor{contract: contract}, nil
}

// NewRegistrationFilterer creates a new log filterer instance of Registration, bound to a specific deployed contract.
func NewRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationFilterer, error) {
	contract, err := bindRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationFilterer{contract: contract}, nil
}

// bindRegistration binds a generic wrapper to an already deployed contract.
func bindRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.RegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transact(opts, method, params...)
}

// ICAOPREFIX is a free data retrieval call binding the contract method 0x1f649377.
//
// Solidity: function ICAO_PREFIX() view returns(string)
func (_Registration *RegistrationCaller) ICAOPREFIX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "ICAO_PREFIX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ICAOPREFIX is a free data retrieval call binding the contract method 0x1f649377.
//
// Solidity: function ICAO_PREFIX() view returns(string)
func (_Registration *RegistrationSession) ICAOPREFIX() (string, error) {
	return _Registration.Contract.ICAOPREFIX(&_Registration.CallOpts)
}

// ICAOPREFIX is a free data retrieval call binding the contract method 0x1f649377.
//
// Solidity: function ICAO_PREFIX() view returns(string)
func (_Registration *RegistrationCallerSession) ICAOPREFIX() (string, error) {
	return _Registration.Contract.ICAOPREFIX(&_Registration.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Registration *RegistrationCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Registration *RegistrationSession) P() (*big.Int, error) {
	return _Registration.Contract.P(&_Registration.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Registration *RegistrationCallerSession) P() (*big.Int, error) {
	return _Registration.Contract.P(&_Registration.CallOpts)
}

// REVOKED is a free data retrieval call binding the contract method 0xe6c292e8.
//
// Solidity: function REVOKED() view returns(bytes32)
func (_Registration *RegistrationCaller) REVOKED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "REVOKED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REVOKED is a free data retrieval call binding the contract method 0xe6c292e8.
//
// Solidity: function REVOKED() view returns(bytes32)
func (_Registration *RegistrationSession) REVOKED() ([32]byte, error) {
	return _Registration.Contract.REVOKED(&_Registration.CallOpts)
}

// REVOKED is a free data retrieval call binding the contract method 0xe6c292e8.
//
// Solidity: function REVOKED() view returns(bytes32)
func (_Registration *RegistrationCallerSession) REVOKED() ([32]byte, error) {
	return _Registration.Contract.REVOKED(&_Registration.CallOpts)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Registration *RegistrationCaller) GetNodeByKey(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeNode, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getNodeByKey", key_)

	if err != nil {
		return *new(SparseMerkleTreeNode), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeNode)).(*SparseMerkleTreeNode)

	return out0, err

}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Registration *RegistrationSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _Registration.Contract.GetNodeByKey(&_Registration.CallOpts, key_)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Registration *RegistrationCallerSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _Registration.Contract.GetNodeByKey(&_Registration.CallOpts, key_)
}

// GetPassportInfo is a free data retrieval call binding the contract method 0x69fab630.
//
// Solidity: function getPassportInfo(bytes32 passportKey_) view returns((bytes32,uint64) passportInfo_, (bytes32,uint64) identityInfo_)
func (_Registration *RegistrationCaller) GetPassportInfo(opts *bind.CallOpts, passportKey_ [32]byte) (struct {
	PassportInfo RegistrationPassportInfo
	IdentityInfo RegistrationIdentityInfo
}, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getPassportInfo", passportKey_)

	outstruct := new(struct {
		PassportInfo RegistrationPassportInfo
		IdentityInfo RegistrationIdentityInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PassportInfo = *abi.ConvertType(out[0], new(RegistrationPassportInfo)).(*RegistrationPassportInfo)
	outstruct.IdentityInfo = *abi.ConvertType(out[1], new(RegistrationIdentityInfo)).(*RegistrationIdentityInfo)

	return *outstruct, err

}

// GetPassportInfo is a free data retrieval call binding the contract method 0x69fab630.
//
// Solidity: function getPassportInfo(bytes32 passportKey_) view returns((bytes32,uint64) passportInfo_, (bytes32,uint64) identityInfo_)
func (_Registration *RegistrationSession) GetPassportInfo(passportKey_ [32]byte) (struct {
	PassportInfo RegistrationPassportInfo
	IdentityInfo RegistrationIdentityInfo
}, error) {
	return _Registration.Contract.GetPassportInfo(&_Registration.CallOpts, passportKey_)
}

// GetPassportInfo is a free data retrieval call binding the contract method 0x69fab630.
//
// Solidity: function getPassportInfo(bytes32 passportKey_) view returns((bytes32,uint64) passportInfo_, (bytes32,uint64) identityInfo_)
func (_Registration *RegistrationCallerSession) GetPassportInfo(passportKey_ [32]byte) (struct {
	PassportInfo RegistrationPassportInfo
	IdentityInfo RegistrationIdentityInfo
}, error) {
	return _Registration.Contract.GetPassportInfo(&_Registration.CallOpts, passportKey_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Registration *RegistrationCaller) GetProof(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeProof, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getProof", key_)

	if err != nil {
		return *new(SparseMerkleTreeProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeProof)).(*SparseMerkleTreeProof)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Registration *RegistrationSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _Registration.Contract.GetProof(&_Registration.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Registration *RegistrationCallerSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _Registration.Contract.GetProof(&_Registration.CallOpts, key_)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Registration *RegistrationCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Registration *RegistrationSession) GetRoot() ([32]byte, error) {
	return _Registration.Contract.GetRoot(&_Registration.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Registration *RegistrationCallerSession) GetRoot() ([32]byte, error) {
	return _Registration.Contract.GetRoot(&_Registration.CallOpts)
}

// IcaoMasterTreeMerkleRoot is a free data retrieval call binding the contract method 0x28093985.
//
// Solidity: function icaoMasterTreeMerkleRoot() view returns(bytes32)
func (_Registration *RegistrationCaller) IcaoMasterTreeMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "icaoMasterTreeMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IcaoMasterTreeMerkleRoot is a free data retrieval call binding the contract method 0x28093985.
//
// Solidity: function icaoMasterTreeMerkleRoot() view returns(bytes32)
func (_Registration *RegistrationSession) IcaoMasterTreeMerkleRoot() ([32]byte, error) {
	return _Registration.Contract.IcaoMasterTreeMerkleRoot(&_Registration.CallOpts)
}

// IcaoMasterTreeMerkleRoot is a free data retrieval call binding the contract method 0x28093985.
//
// Solidity: function icaoMasterTreeMerkleRoot() view returns(bytes32)
func (_Registration *RegistrationCallerSession) IcaoMasterTreeMerkleRoot() ([32]byte, error) {
	return _Registration.Contract.IcaoMasterTreeMerkleRoot(&_Registration.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registration *RegistrationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registration *RegistrationSession) Owner() (common.Address, error) {
	return _Registration.Contract.Owner(&_Registration.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registration *RegistrationCallerSession) Owner() (common.Address, error) {
	return _Registration.Contract.Owner(&_Registration.CallOpts)
}

// PassportDispatchers is a free data retrieval call binding the contract method 0xdd9492a7.
//
// Solidity: function passportDispatchers(bytes32 ) view returns(address)
func (_Registration *RegistrationCaller) PassportDispatchers(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "passportDispatchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PassportDispatchers is a free data retrieval call binding the contract method 0xdd9492a7.
//
// Solidity: function passportDispatchers(bytes32 ) view returns(address)
func (_Registration *RegistrationSession) PassportDispatchers(arg0 [32]byte) (common.Address, error) {
	return _Registration.Contract.PassportDispatchers(&_Registration.CallOpts, arg0)
}

// PassportDispatchers is a free data retrieval call binding the contract method 0xdd9492a7.
//
// Solidity: function passportDispatchers(bytes32 ) view returns(address)
func (_Registration *RegistrationCallerSession) PassportDispatchers(arg0 [32]byte) (common.Address, error) {
	return _Registration.Contract.PassportDispatchers(&_Registration.CallOpts, arg0)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Registration *RegistrationCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Registration *RegistrationSession) Signer() (common.Address, error) {
	return _Registration.Contract.Signer(&_Registration.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Registration *RegistrationCallerSession) Signer() (common.Address, error) {
	return _Registration.Contract.Signer(&_Registration.CallOpts)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x759bf7e1.
//
// Solidity: function __Registration_init(uint256 treeHeight_, address signer_, bytes32 icaoMasterTreeMerkleRoot_) returns()
func (_Registration *RegistrationTransactor) RegistrationInit(opts *bind.TransactOpts, treeHeight_ *big.Int, signer_ common.Address, icaoMasterTreeMerkleRoot_ [32]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "__Registration_init", treeHeight_, signer_, icaoMasterTreeMerkleRoot_)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x759bf7e1.
//
// Solidity: function __Registration_init(uint256 treeHeight_, address signer_, bytes32 icaoMasterTreeMerkleRoot_) returns()
func (_Registration *RegistrationSession) RegistrationInit(treeHeight_ *big.Int, signer_ common.Address, icaoMasterTreeMerkleRoot_ [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationInit(&_Registration.TransactOpts, treeHeight_, signer_, icaoMasterTreeMerkleRoot_)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x759bf7e1.
//
// Solidity: function __Registration_init(uint256 treeHeight_, address signer_, bytes32 icaoMasterTreeMerkleRoot_) returns()
func (_Registration *RegistrationTransactorSession) RegistrationInit(treeHeight_ *big.Int, signer_ common.Address, icaoMasterTreeMerkleRoot_ [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationInit(&_Registration.TransactOpts, treeHeight_, signer_, icaoMasterTreeMerkleRoot_)
}

// AddDispatcher is a paid mutator transaction binding the contract method 0x6fdc1f6c.
//
// Solidity: function addDispatcher(bytes32 dispatcherType_, address dispatcher_) returns()
func (_Registration *RegistrationTransactor) AddDispatcher(opts *bind.TransactOpts, dispatcherType_ [32]byte, dispatcher_ common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "addDispatcher", dispatcherType_, dispatcher_)
}

// AddDispatcher is a paid mutator transaction binding the contract method 0x6fdc1f6c.
//
// Solidity: function addDispatcher(bytes32 dispatcherType_, address dispatcher_) returns()
func (_Registration *RegistrationSession) AddDispatcher(dispatcherType_ [32]byte, dispatcher_ common.Address) (*types.Transaction, error) {
	return _Registration.Contract.AddDispatcher(&_Registration.TransactOpts, dispatcherType_, dispatcher_)
}

// AddDispatcher is a paid mutator transaction binding the contract method 0x6fdc1f6c.
//
// Solidity: function addDispatcher(bytes32 dispatcherType_, address dispatcher_) returns()
func (_Registration *RegistrationTransactorSession) AddDispatcher(dispatcherType_ [32]byte, dispatcher_ common.Address) (*types.Transaction, error) {
	return _Registration.Contract.AddDispatcher(&_Registration.TransactOpts, dispatcherType_, dispatcher_)
}

// ChangeICAOMasterTreeRoot is a paid mutator transaction binding the contract method 0x7ead04df.
//
// Solidity: function changeICAOMasterTreeRoot(bytes32 newRoot_, uint256 timestamp, bytes proof_) returns()
func (_Registration *RegistrationTransactor) ChangeICAOMasterTreeRoot(opts *bind.TransactOpts, newRoot_ [32]byte, timestamp *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "changeICAOMasterTreeRoot", newRoot_, timestamp, proof_)
}

// ChangeICAOMasterTreeRoot is a paid mutator transaction binding the contract method 0x7ead04df.
//
// Solidity: function changeICAOMasterTreeRoot(bytes32 newRoot_, uint256 timestamp, bytes proof_) returns()
func (_Registration *RegistrationSession) ChangeICAOMasterTreeRoot(newRoot_ [32]byte, timestamp *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _Registration.Contract.ChangeICAOMasterTreeRoot(&_Registration.TransactOpts, newRoot_, timestamp, proof_)
}

// ChangeICAOMasterTreeRoot is a paid mutator transaction binding the contract method 0x7ead04df.
//
// Solidity: function changeICAOMasterTreeRoot(bytes32 newRoot_, uint256 timestamp, bytes proof_) returns()
func (_Registration *RegistrationTransactorSession) ChangeICAOMasterTreeRoot(newRoot_ [32]byte, timestamp *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _Registration.Contract.ChangeICAOMasterTreeRoot(&_Registration.TransactOpts, newRoot_, timestamp, proof_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Registration *RegistrationTransactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Registration *RegistrationSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Registration.Contract.ChangeSigner(&_Registration.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Registration *RegistrationTransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Registration.Contract.ChangeSigner(&_Registration.TransactOpts, newSignerPubKey_, signature_)
}

// Register is a paid mutator transaction binding the contract method 0xd7656e78.
//
// Solidity: function register(uint256 identityKey_, uint256 dgCommit_, (bytes32,bytes,bytes) passport_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Registration *RegistrationTransactor) Register(opts *bind.TransactOpts, identityKey_ *big.Int, dgCommit_ *big.Int, passport_ RegistrationPassport, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "register", identityKey_, dgCommit_, passport_, zkPoints_)
}

// Register is a paid mutator transaction binding the contract method 0xd7656e78.
//
// Solidity: function register(uint256 identityKey_, uint256 dgCommit_, (bytes32,bytes,bytes) passport_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Registration *RegistrationSession) Register(identityKey_ *big.Int, dgCommit_ *big.Int, passport_ RegistrationPassport, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, identityKey_, dgCommit_, passport_, zkPoints_)
}

// Register is a paid mutator transaction binding the contract method 0xd7656e78.
//
// Solidity: function register(uint256 identityKey_, uint256 dgCommit_, (bytes32,bytes,bytes) passport_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Registration *RegistrationTransactorSession) Register(identityKey_ *big.Int, dgCommit_ *big.Int, passport_ RegistrationPassport, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, identityKey_, dgCommit_, passport_, zkPoints_)
}

// ReissueIdentity is a paid mutator transaction binding the contract method 0xf6ea8086.
//
// Solidity: function reissueIdentity(uint256 identityKey_, uint256 dgCommit_, (bytes32,bytes,bytes) passport_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Registration *RegistrationTransactor) ReissueIdentity(opts *bind.TransactOpts, identityKey_ *big.Int, dgCommit_ *big.Int, passport_ RegistrationPassport, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "reissueIdentity", identityKey_, dgCommit_, passport_, zkPoints_)
}

// ReissueIdentity is a paid mutator transaction binding the contract method 0xf6ea8086.
//
// Solidity: function reissueIdentity(uint256 identityKey_, uint256 dgCommit_, (bytes32,bytes,bytes) passport_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Registration *RegistrationSession) ReissueIdentity(identityKey_ *big.Int, dgCommit_ *big.Int, passport_ RegistrationPassport, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Registration.Contract.ReissueIdentity(&_Registration.TransactOpts, identityKey_, dgCommit_, passport_, zkPoints_)
}

// ReissueIdentity is a paid mutator transaction binding the contract method 0xf6ea8086.
//
// Solidity: function reissueIdentity(uint256 identityKey_, uint256 dgCommit_, (bytes32,bytes,bytes) passport_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Registration *RegistrationTransactorSession) ReissueIdentity(identityKey_ *big.Int, dgCommit_ *big.Int, passport_ RegistrationPassport, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Registration.Contract.ReissueIdentity(&_Registration.TransactOpts, identityKey_, dgCommit_, passport_, zkPoints_)
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x7e68c9ea.
//
// Solidity: function removeDispatcher(bytes32 dispatcherType_) returns()
func (_Registration *RegistrationTransactor) RemoveDispatcher(opts *bind.TransactOpts, dispatcherType_ [32]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "removeDispatcher", dispatcherType_)
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x7e68c9ea.
//
// Solidity: function removeDispatcher(bytes32 dispatcherType_) returns()
func (_Registration *RegistrationSession) RemoveDispatcher(dispatcherType_ [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.RemoveDispatcher(&_Registration.TransactOpts, dispatcherType_)
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x7e68c9ea.
//
// Solidity: function removeDispatcher(bytes32 dispatcherType_) returns()
func (_Registration *RegistrationTransactorSession) RemoveDispatcher(dispatcherType_ [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.RemoveDispatcher(&_Registration.TransactOpts, dispatcherType_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registration *RegistrationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registration *RegistrationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registration.Contract.RenounceOwnership(&_Registration.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registration *RegistrationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registration.Contract.RenounceOwnership(&_Registration.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xc8ca2267.
//
// Solidity: function revoke(uint256 identityKey_, (bytes32,bytes,bytes) passport_) returns()
func (_Registration *RegistrationTransactor) Revoke(opts *bind.TransactOpts, identityKey_ *big.Int, passport_ RegistrationPassport) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "revoke", identityKey_, passport_)
}

// Revoke is a paid mutator transaction binding the contract method 0xc8ca2267.
//
// Solidity: function revoke(uint256 identityKey_, (bytes32,bytes,bytes) passport_) returns()
func (_Registration *RegistrationSession) Revoke(identityKey_ *big.Int, passport_ RegistrationPassport) (*types.Transaction, error) {
	return _Registration.Contract.Revoke(&_Registration.TransactOpts, identityKey_, passport_)
}

// Revoke is a paid mutator transaction binding the contract method 0xc8ca2267.
//
// Solidity: function revoke(uint256 identityKey_, (bytes32,bytes,bytes) passport_) returns()
func (_Registration *RegistrationTransactorSession) Revoke(identityKey_ *big.Int, passport_ RegistrationPassport) (*types.Transaction, error) {
	return _Registration.Contract.Revoke(&_Registration.TransactOpts, identityKey_, passport_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registration *RegistrationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registration *RegistrationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registration.Contract.TransferOwnership(&_Registration.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registration *RegistrationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registration.Contract.TransferOwnership(&_Registration.TransactOpts, newOwner)
}

// RegistrationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registration contract.
type RegistrationInitializedIterator struct {
	Event *RegistrationInitialized // Event containing the contract specifics and raw log

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
func (it *RegistrationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationInitialized)
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
		it.Event = new(RegistrationInitialized)
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
func (it *RegistrationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationInitialized represents a Initialized event raised by the Registration contract.
type RegistrationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registration *RegistrationFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistrationInitializedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistrationInitializedIterator{contract: _Registration.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registration *RegistrationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistrationInitialized) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationInitialized)
				if err := _Registration.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseInitialized(log types.Log) (*RegistrationInitialized, error) {
	event := new(RegistrationInitialized)
	if err := _Registration.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registration contract.
type RegistrationOwnershipTransferredIterator struct {
	Event *RegistrationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationOwnershipTransferred)
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
		it.Event = new(RegistrationOwnershipTransferred)
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
func (it *RegistrationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationOwnershipTransferred represents a OwnershipTransferred event raised by the Registration contract.
type RegistrationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registration *RegistrationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registration.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationOwnershipTransferredIterator{contract: _Registration.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registration *RegistrationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registration.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationOwnershipTransferred)
				if err := _Registration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrationOwnershipTransferred, error) {
	event := new(RegistrationOwnershipTransferred)
	if err := _Registration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Registration contract.
type RegistrationRegisteredIterator struct {
	Event *RegistrationRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRegistered)
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
		it.Event = new(RegistrationRegistered)
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
func (it *RegistrationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRegistered represents a Registered event raised by the Registration contract.
type RegistrationRegistered struct {
	PassportKey [32]byte
	IdentityKey [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x8828766b1a3f99eda4bad24a0979129565597036bac9a94cdccd1d1e6ef17eb2.
//
// Solidity: event Registered(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) FilterRegistered(opts *bind.FilterOpts) (*RegistrationRegisteredIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &RegistrationRegisteredIterator{contract: _Registration.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x8828766b1a3f99eda4bad24a0979129565597036bac9a94cdccd1d1e6ef17eb2.
//
// Solidity: event Registered(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistrationRegistered) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRegistered)
				if err := _Registration.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x8828766b1a3f99eda4bad24a0979129565597036bac9a94cdccd1d1e6ef17eb2.
//
// Solidity: event Registered(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) ParseRegistered(log types.Log) (*RegistrationRegistered, error) {
	event := new(RegistrationRegistered)
	if err := _Registration.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationReissuedIdentityIterator is returned from FilterReissuedIdentity and is used to iterate over the raw logs and unpacked data for ReissuedIdentity events raised by the Registration contract.
type RegistrationReissuedIdentityIterator struct {
	Event *RegistrationReissuedIdentity // Event containing the contract specifics and raw log

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
func (it *RegistrationReissuedIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationReissuedIdentity)
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
		it.Event = new(RegistrationReissuedIdentity)
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
func (it *RegistrationReissuedIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationReissuedIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationReissuedIdentity represents a ReissuedIdentity event raised by the Registration contract.
type RegistrationReissuedIdentity struct {
	PassportKey [32]byte
	IdentityKey [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReissuedIdentity is a free log retrieval operation binding the contract event 0x144723860a72ec80ca97d869cc90e3cd4b29fa675265fc686f07af1b63646659.
//
// Solidity: event ReissuedIdentity(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) FilterReissuedIdentity(opts *bind.FilterOpts) (*RegistrationReissuedIdentityIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "ReissuedIdentity")
	if err != nil {
		return nil, err
	}
	return &RegistrationReissuedIdentityIterator{contract: _Registration.contract, event: "ReissuedIdentity", logs: logs, sub: sub}, nil
}

// WatchReissuedIdentity is a free log subscription operation binding the contract event 0x144723860a72ec80ca97d869cc90e3cd4b29fa675265fc686f07af1b63646659.
//
// Solidity: event ReissuedIdentity(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) WatchReissuedIdentity(opts *bind.WatchOpts, sink chan<- *RegistrationReissuedIdentity) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "ReissuedIdentity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationReissuedIdentity)
				if err := _Registration.contract.UnpackLog(event, "ReissuedIdentity", log); err != nil {
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

// ParseReissuedIdentity is a log parse operation binding the contract event 0x144723860a72ec80ca97d869cc90e3cd4b29fa675265fc686f07af1b63646659.
//
// Solidity: event ReissuedIdentity(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) ParseReissuedIdentity(log types.Log) (*RegistrationReissuedIdentity, error) {
	event := new(RegistrationReissuedIdentity)
	if err := _Registration.contract.UnpackLog(event, "ReissuedIdentity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the Registration contract.
type RegistrationRevokedIterator struct {
	Event *RegistrationRevoked // Event containing the contract specifics and raw log

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
func (it *RegistrationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRevoked)
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
		it.Event = new(RegistrationRevoked)
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
func (it *RegistrationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRevoked represents a Revoked event raised by the Registration contract.
type RegistrationRevoked struct {
	PassportKey [32]byte
	IdentityKey [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0xa75610f599cea10d1a6bb6332878af5fffa58f57b227cba0d0229fa90a63cc51.
//
// Solidity: event Revoked(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) FilterRevoked(opts *bind.FilterOpts) (*RegistrationRevokedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return &RegistrationRevokedIterator{contract: _Registration.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0xa75610f599cea10d1a6bb6332878af5fffa58f57b227cba0d0229fa90a63cc51.
//
// Solidity: event Revoked(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *RegistrationRevoked) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRevoked)
				if err := _Registration.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0xa75610f599cea10d1a6bb6332878af5fffa58f57b227cba0d0229fa90a63cc51.
//
// Solidity: event Revoked(bytes32 passportKey, bytes32 identityKey)
func (_Registration *RegistrationFilterer) ParseRevoked(log types.Log) (*RegistrationRevoked, error) {
	event := new(RegistrationRevoked)
	if err := _Registration.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
