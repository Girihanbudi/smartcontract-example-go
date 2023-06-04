// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package filesystem

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

// FileSystemFile is an auto generated low-level Go binding around an user-defined struct.
type FileSystemFile struct {
	Name string
	File string
}

// FilesystemMetaData contains all meta data concerning the Filesystem contract.
var FilesystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"addFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFile\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"}],\"internalType\":\"structFileSystem.File\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506105b2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063248bfc3b146100465780632fa816921461005b57806330ae913714610079575b600080fd5b610059610054366004610347565b610081565b005b6100636100cd565b60405161007091906103f1565b60405180910390f35b610059610222565b60408051808201825283815260208082018490523360009081529081905291909120815181906100b190826104bc565b50602082015160018201906100c690826104bc565b5050505050565b60408051808201909152606080825260208201523360009081526020819052604090819020815180830190925280548290829061010990610433565b80601f016020809104026020016040519081016040528092919081815260200182805461013590610433565b80156101825780601f1061015757610100808354040283529160200191610182565b820191906000526020600020905b81548152906001019060200180831161016557829003601f168201915b5050505050815260200160018201805461019b90610433565b80601f01602080910402602001604051908101604052809291908181526020018280546101c790610433565b80156102145780601f106101e957610100808354040283529160200191610214565b820191906000526020600020905b8154815290600101906020018083116101f757829003601f168201915b505050505081525050905090565b3360009081526020819052604081209061023c828261024e565b61024a60018301600061024e565b5050565b50805461025a90610433565b6000825580601f1061026a575050565b601f016020900490600052602060002090810190610288919061028b565b50565b5b808211156102a0576000815560010161028c565b5090565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126102cb57600080fd5b813567ffffffffffffffff808211156102e6576102e66102a4565b604051601f8301601f19908116603f0116810190828211818310171561030e5761030e6102a4565b8160405283815286602085880101111561032757600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561035a57600080fd5b823567ffffffffffffffff8082111561037257600080fd5b61037e868387016102ba565b9350602085013591508082111561039457600080fd5b506103a1858286016102ba565b9150509250929050565b6000815180845260005b818110156103d1576020818501810151868301820152016103b5565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600082516040602084015261040d60608401826103ab565b90506020840151601f1984830301604085015261042a82826103ab565b95945050505050565b600181811c9082168061044757607f821691505b60208210810361046757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156104b757600081815260208120601f850160051c810160208610156104945750805b601f850160051c820191505b818110156104b3578281556001016104a0565b5050505b505050565b815167ffffffffffffffff8111156104d6576104d66102a4565b6104ea816104e48454610433565b8461046d565b602080601f83116001811461051f57600084156105075750858301515b600019600386901b1c1916600185901b1785556104b3565b600085815260208120601f198616915b8281101561054e5788860151825594840194600190910190840161052f565b508582101561056c5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220a3f951f0c2838adc9bd1d92eae103233730922b8e21d34b63248baa957d41d3a64736f6c63430008130033",
}

// FilesystemABI is the input ABI used to generate the binding from.
// Deprecated: Use FilesystemMetaData.ABI instead.
var FilesystemABI = FilesystemMetaData.ABI

// FilesystemBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FilesystemMetaData.Bin instead.
var FilesystemBin = FilesystemMetaData.Bin

// DeployFilesystem deploys a new Ethereum contract, binding an instance of Filesystem to it.
func DeployFilesystem(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Filesystem, error) {
	parsed, err := FilesystemMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FilesystemBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Filesystem{FilesystemCaller: FilesystemCaller{contract: contract}, FilesystemTransactor: FilesystemTransactor{contract: contract}, FilesystemFilterer: FilesystemFilterer{contract: contract}}, nil
}

// Filesystem is an auto generated Go binding around an Ethereum contract.
type Filesystem struct {
	FilesystemCaller     // Read-only binding to the contract
	FilesystemTransactor // Write-only binding to the contract
	FilesystemFilterer   // Log filterer for contract events
}

// FilesystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type FilesystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilesystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FilesystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilesystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FilesystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilesystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FilesystemSession struct {
	Contract     *Filesystem       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FilesystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FilesystemCallerSession struct {
	Contract *FilesystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FilesystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FilesystemTransactorSession struct {
	Contract     *FilesystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FilesystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type FilesystemRaw struct {
	Contract *Filesystem // Generic contract binding to access the raw methods on
}

// FilesystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FilesystemCallerRaw struct {
	Contract *FilesystemCaller // Generic read-only contract binding to access the raw methods on
}

// FilesystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FilesystemTransactorRaw struct {
	Contract *FilesystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFilesystem creates a new instance of Filesystem, bound to a specific deployed contract.
func NewFilesystem(address common.Address, backend bind.ContractBackend) (*Filesystem, error) {
	contract, err := bindFilesystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Filesystem{FilesystemCaller: FilesystemCaller{contract: contract}, FilesystemTransactor: FilesystemTransactor{contract: contract}, FilesystemFilterer: FilesystemFilterer{contract: contract}}, nil
}

// NewFilesystemCaller creates a new read-only instance of Filesystem, bound to a specific deployed contract.
func NewFilesystemCaller(address common.Address, caller bind.ContractCaller) (*FilesystemCaller, error) {
	contract, err := bindFilesystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FilesystemCaller{contract: contract}, nil
}

// NewFilesystemTransactor creates a new write-only instance of Filesystem, bound to a specific deployed contract.
func NewFilesystemTransactor(address common.Address, transactor bind.ContractTransactor) (*FilesystemTransactor, error) {
	contract, err := bindFilesystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FilesystemTransactor{contract: contract}, nil
}

// NewFilesystemFilterer creates a new log filterer instance of Filesystem, bound to a specific deployed contract.
func NewFilesystemFilterer(address common.Address, filterer bind.ContractFilterer) (*FilesystemFilterer, error) {
	contract, err := bindFilesystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FilesystemFilterer{contract: contract}, nil
}

// bindFilesystem binds a generic wrapper to an already deployed contract.
func bindFilesystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FilesystemMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Filesystem *FilesystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Filesystem.Contract.FilesystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Filesystem *FilesystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Filesystem.Contract.FilesystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Filesystem *FilesystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Filesystem.Contract.FilesystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Filesystem *FilesystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Filesystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Filesystem *FilesystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Filesystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Filesystem *FilesystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Filesystem.Contract.contract.Transact(opts, method, params...)
}

// GetFile is a free data retrieval call binding the contract method 0x2fa81692.
//
// Solidity: function getFile() view returns((string,string))
func (_Filesystem *FilesystemCaller) GetFile(opts *bind.CallOpts) (FileSystemFile, error) {
	var out []interface{}
	err := _Filesystem.contract.Call(opts, &out, "getFile")

	if err != nil {
		return *new(FileSystemFile), err
	}

	out0 := *abi.ConvertType(out[0], new(FileSystemFile)).(*FileSystemFile)

	return out0, err

}

// GetFile is a free data retrieval call binding the contract method 0x2fa81692.
//
// Solidity: function getFile() view returns((string,string))
func (_Filesystem *FilesystemSession) GetFile() (FileSystemFile, error) {
	return _Filesystem.Contract.GetFile(&_Filesystem.CallOpts)
}

// GetFile is a free data retrieval call binding the contract method 0x2fa81692.
//
// Solidity: function getFile() view returns((string,string))
func (_Filesystem *FilesystemCallerSession) GetFile() (FileSystemFile, error) {
	return _Filesystem.Contract.GetFile(&_Filesystem.CallOpts)
}

// AddFile is a paid mutator transaction binding the contract method 0x248bfc3b.
//
// Solidity: function addFile(string _name, string _file) returns()
func (_Filesystem *FilesystemTransactor) AddFile(opts *bind.TransactOpts, _name string, _file string) (*types.Transaction, error) {
	return _Filesystem.contract.Transact(opts, "addFile", _name, _file)
}

// AddFile is a paid mutator transaction binding the contract method 0x248bfc3b.
//
// Solidity: function addFile(string _name, string _file) returns()
func (_Filesystem *FilesystemSession) AddFile(_name string, _file string) (*types.Transaction, error) {
	return _Filesystem.Contract.AddFile(&_Filesystem.TransactOpts, _name, _file)
}

// AddFile is a paid mutator transaction binding the contract method 0x248bfc3b.
//
// Solidity: function addFile(string _name, string _file) returns()
func (_Filesystem *FilesystemTransactorSession) AddFile(_name string, _file string) (*types.Transaction, error) {
	return _Filesystem.Contract.AddFile(&_Filesystem.TransactOpts, _name, _file)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x30ae9137.
//
// Solidity: function deleteFile() returns()
func (_Filesystem *FilesystemTransactor) DeleteFile(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Filesystem.contract.Transact(opts, "deleteFile")
}

// DeleteFile is a paid mutator transaction binding the contract method 0x30ae9137.
//
// Solidity: function deleteFile() returns()
func (_Filesystem *FilesystemSession) DeleteFile() (*types.Transaction, error) {
	return _Filesystem.Contract.DeleteFile(&_Filesystem.TransactOpts)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x30ae9137.
//
// Solidity: function deleteFile() returns()
func (_Filesystem *FilesystemTransactorSession) DeleteFile() (*types.Transaction, error) {
	return _Filesystem.Contract.DeleteFile(&_Filesystem.TransactOpts)
}
