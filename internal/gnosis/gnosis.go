// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gnosis

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

// BaseGuardMetaData contains all meta data concerning the BaseGuard contract.
var BaseGuardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"checkAfterExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"checkTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"93271368": "checkAfterExecution(bytes32,bool)",
		"75f0bb52": "checkTransaction(address,uint256,bytes,uint8,uint256,uint256,uint256,address,address,bytes,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// BaseGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseGuardMetaData.ABI instead.
var BaseGuardABI = BaseGuardMetaData.ABI

// Deprecated: Use BaseGuardMetaData.Sigs instead.
// BaseGuardFuncSigs maps the 4-byte function signature to its string representation.
var BaseGuardFuncSigs = BaseGuardMetaData.Sigs

// BaseGuard is an auto generated Go binding around an Ethereum contract.
type BaseGuard struct {
	BaseGuardCaller     // Read-only binding to the contract
	BaseGuardTransactor // Write-only binding to the contract
	BaseGuardFilterer   // Log filterer for contract events
}

// BaseGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseGuardSession struct {
	Contract     *BaseGuard        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseGuardCallerSession struct {
	Contract *BaseGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BaseGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseGuardTransactorSession struct {
	Contract     *BaseGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BaseGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseGuardRaw struct {
	Contract *BaseGuard // Generic contract binding to access the raw methods on
}

// BaseGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseGuardCallerRaw struct {
	Contract *BaseGuardCaller // Generic read-only contract binding to access the raw methods on
}

// BaseGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseGuardTransactorRaw struct {
	Contract *BaseGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseGuard creates a new instance of BaseGuard, bound to a specific deployed contract.
func NewBaseGuard(address common.Address, backend bind.ContractBackend) (*BaseGuard, error) {
	contract, err := bindBaseGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseGuard{BaseGuardCaller: BaseGuardCaller{contract: contract}, BaseGuardTransactor: BaseGuardTransactor{contract: contract}, BaseGuardFilterer: BaseGuardFilterer{contract: contract}}, nil
}

// NewBaseGuardCaller creates a new read-only instance of BaseGuard, bound to a specific deployed contract.
func NewBaseGuardCaller(address common.Address, caller bind.ContractCaller) (*BaseGuardCaller, error) {
	contract, err := bindBaseGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseGuardCaller{contract: contract}, nil
}

// NewBaseGuardTransactor creates a new write-only instance of BaseGuard, bound to a specific deployed contract.
func NewBaseGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseGuardTransactor, error) {
	contract, err := bindBaseGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseGuardTransactor{contract: contract}, nil
}

// NewBaseGuardFilterer creates a new log filterer instance of BaseGuard, bound to a specific deployed contract.
func NewBaseGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseGuardFilterer, error) {
	contract, err := bindBaseGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseGuardFilterer{contract: contract}, nil
}

// bindBaseGuard binds a generic wrapper to an already deployed contract.
func bindBaseGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseGuard *BaseGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseGuard.Contract.BaseGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseGuard *BaseGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseGuard.Contract.BaseGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseGuard *BaseGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseGuard.Contract.BaseGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseGuard *BaseGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseGuard *BaseGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseGuard *BaseGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseGuard.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseGuard *BaseGuardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BaseGuard.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseGuard *BaseGuardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseGuard.Contract.SupportsInterface(&_BaseGuard.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseGuard *BaseGuardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseGuard.Contract.SupportsInterface(&_BaseGuard.CallOpts, interfaceId)
}

// CheckAfterExecution is a paid mutator transaction binding the contract method 0x93271368.
//
// Solidity: function checkAfterExecution(bytes32 txHash, bool success) returns()
func (_BaseGuard *BaseGuardTransactor) CheckAfterExecution(opts *bind.TransactOpts, txHash [32]byte, success bool) (*types.Transaction, error) {
	return _BaseGuard.contract.Transact(opts, "checkAfterExecution", txHash, success)
}

// CheckAfterExecution is a paid mutator transaction binding the contract method 0x93271368.
//
// Solidity: function checkAfterExecution(bytes32 txHash, bool success) returns()
func (_BaseGuard *BaseGuardSession) CheckAfterExecution(txHash [32]byte, success bool) (*types.Transaction, error) {
	return _BaseGuard.Contract.CheckAfterExecution(&_BaseGuard.TransactOpts, txHash, success)
}

// CheckAfterExecution is a paid mutator transaction binding the contract method 0x93271368.
//
// Solidity: function checkAfterExecution(bytes32 txHash, bool success) returns()
func (_BaseGuard *BaseGuardTransactorSession) CheckAfterExecution(txHash [32]byte, success bool) (*types.Transaction, error) {
	return _BaseGuard.Contract.CheckAfterExecution(&_BaseGuard.TransactOpts, txHash, success)
}

// CheckTransaction is a paid mutator transaction binding the contract method 0x75f0bb52.
//
// Solidity: function checkTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, address msgSender) returns()
func (_BaseGuard *BaseGuardTransactor) CheckTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte, msgSender common.Address) (*types.Transaction, error) {
	return _BaseGuard.contract.Transact(opts, "checkTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures, msgSender)
}

// CheckTransaction is a paid mutator transaction binding the contract method 0x75f0bb52.
//
// Solidity: function checkTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, address msgSender) returns()
func (_BaseGuard *BaseGuardSession) CheckTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte, msgSender common.Address) (*types.Transaction, error) {
	return _BaseGuard.Contract.CheckTransaction(&_BaseGuard.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures, msgSender)
}

// CheckTransaction is a paid mutator transaction binding the contract method 0x75f0bb52.
//
// Solidity: function checkTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, address msgSender) returns()
func (_BaseGuard *BaseGuardTransactorSession) CheckTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte, msgSender common.Address) (*types.Transaction, error) {
	return _BaseGuard.Contract.CheckTransaction(&_BaseGuard.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures, msgSender)
}

// EnumMetaData contains all meta data concerning the Enum contract.
var EnumMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220684bee2cac9c90890d100e47e005e2cd5b0564758567c50e959308cd2ab74a7764736f6c634300080d0033",
}

// EnumABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumMetaData.ABI instead.
var EnumABI = EnumMetaData.ABI

// EnumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumMetaData.Bin instead.
var EnumBin = EnumMetaData.Bin

// DeployEnum deploys a new Ethereum contract, binding an instance of Enum to it.
func DeployEnum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Enum, error) {
	parsed, err := EnumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Enum{EnumCaller: EnumCaller{contract: contract}, EnumTransactor: EnumTransactor{contract: contract}, EnumFilterer: EnumFilterer{contract: contract}}, nil
}

// Enum is an auto generated Go binding around an Ethereum contract.
type Enum struct {
	EnumCaller     // Read-only binding to the contract
	EnumTransactor // Write-only binding to the contract
	EnumFilterer   // Log filterer for contract events
}

// EnumCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumSession struct {
	Contract     *Enum             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumCallerSession struct {
	Contract *EnumCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EnumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumTransactorSession struct {
	Contract     *EnumTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumRaw struct {
	Contract *Enum // Generic contract binding to access the raw methods on
}

// EnumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumCallerRaw struct {
	Contract *EnumCaller // Generic read-only contract binding to access the raw methods on
}

// EnumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumTransactorRaw struct {
	Contract *EnumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnum creates a new instance of Enum, bound to a specific deployed contract.
func NewEnum(address common.Address, backend bind.ContractBackend) (*Enum, error) {
	contract, err := bindEnum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Enum{EnumCaller: EnumCaller{contract: contract}, EnumTransactor: EnumTransactor{contract: contract}, EnumFilterer: EnumFilterer{contract: contract}}, nil
}

// NewEnumCaller creates a new read-only instance of Enum, bound to a specific deployed contract.
func NewEnumCaller(address common.Address, caller bind.ContractCaller) (*EnumCaller, error) {
	contract, err := bindEnum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumCaller{contract: contract}, nil
}

// NewEnumTransactor creates a new write-only instance of Enum, bound to a specific deployed contract.
func NewEnumTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumTransactor, error) {
	contract, err := bindEnum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumTransactor{contract: contract}, nil
}

// NewEnumFilterer creates a new log filterer instance of Enum, bound to a specific deployed contract.
func NewEnumFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumFilterer, error) {
	contract, err := bindEnum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumFilterer{contract: contract}, nil
}

// bindEnum binds a generic wrapper to an already deployed contract.
func bindEnum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enum *EnumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Enum.Contract.EnumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enum *EnumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enum.Contract.EnumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enum *EnumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enum.Contract.EnumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enum *EnumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Enum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enum *EnumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enum *EnumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enum.Contract.contract.Transact(opts, method, params...)
}

// EtherPaymentFallbackMetaData contains all meta data concerning the EtherPaymentFallback contract.
var EtherPaymentFallbackMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50607a8061001e6000396000f3fe608060405236603f5760405134815233907f3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d9060200160405180910390a2005b600080fdfea26469706673582212205c7489e942a4c928ea50c38de33e89a4cc0a5aea80bcd77adbb67c120832926364736f6c634300080d0033",
}

// EtherPaymentFallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherPaymentFallbackMetaData.ABI instead.
var EtherPaymentFallbackABI = EtherPaymentFallbackMetaData.ABI

// EtherPaymentFallbackBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtherPaymentFallbackMetaData.Bin instead.
var EtherPaymentFallbackBin = EtherPaymentFallbackMetaData.Bin

// DeployEtherPaymentFallback deploys a new Ethereum contract, binding an instance of EtherPaymentFallback to it.
func DeployEtherPaymentFallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EtherPaymentFallback, error) {
	parsed, err := EtherPaymentFallbackMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherPaymentFallbackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherPaymentFallback{EtherPaymentFallbackCaller: EtherPaymentFallbackCaller{contract: contract}, EtherPaymentFallbackTransactor: EtherPaymentFallbackTransactor{contract: contract}, EtherPaymentFallbackFilterer: EtherPaymentFallbackFilterer{contract: contract}}, nil
}

// EtherPaymentFallback is an auto generated Go binding around an Ethereum contract.
type EtherPaymentFallback struct {
	EtherPaymentFallbackCaller     // Read-only binding to the contract
	EtherPaymentFallbackTransactor // Write-only binding to the contract
	EtherPaymentFallbackFilterer   // Log filterer for contract events
}

// EtherPaymentFallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherPaymentFallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherPaymentFallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherPaymentFallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherPaymentFallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherPaymentFallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherPaymentFallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherPaymentFallbackSession struct {
	Contract     *EtherPaymentFallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtherPaymentFallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherPaymentFallbackCallerSession struct {
	Contract *EtherPaymentFallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// EtherPaymentFallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherPaymentFallbackTransactorSession struct {
	Contract     *EtherPaymentFallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// EtherPaymentFallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherPaymentFallbackRaw struct {
	Contract *EtherPaymentFallback // Generic contract binding to access the raw methods on
}

// EtherPaymentFallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherPaymentFallbackCallerRaw struct {
	Contract *EtherPaymentFallbackCaller // Generic read-only contract binding to access the raw methods on
}

// EtherPaymentFallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherPaymentFallbackTransactorRaw struct {
	Contract *EtherPaymentFallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherPaymentFallback creates a new instance of EtherPaymentFallback, bound to a specific deployed contract.
func NewEtherPaymentFallback(address common.Address, backend bind.ContractBackend) (*EtherPaymentFallback, error) {
	contract, err := bindEtherPaymentFallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherPaymentFallback{EtherPaymentFallbackCaller: EtherPaymentFallbackCaller{contract: contract}, EtherPaymentFallbackTransactor: EtherPaymentFallbackTransactor{contract: contract}, EtherPaymentFallbackFilterer: EtherPaymentFallbackFilterer{contract: contract}}, nil
}

// NewEtherPaymentFallbackCaller creates a new read-only instance of EtherPaymentFallback, bound to a specific deployed contract.
func NewEtherPaymentFallbackCaller(address common.Address, caller bind.ContractCaller) (*EtherPaymentFallbackCaller, error) {
	contract, err := bindEtherPaymentFallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherPaymentFallbackCaller{contract: contract}, nil
}

// NewEtherPaymentFallbackTransactor creates a new write-only instance of EtherPaymentFallback, bound to a specific deployed contract.
func NewEtherPaymentFallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherPaymentFallbackTransactor, error) {
	contract, err := bindEtherPaymentFallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherPaymentFallbackTransactor{contract: contract}, nil
}

// NewEtherPaymentFallbackFilterer creates a new log filterer instance of EtherPaymentFallback, bound to a specific deployed contract.
func NewEtherPaymentFallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherPaymentFallbackFilterer, error) {
	contract, err := bindEtherPaymentFallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherPaymentFallbackFilterer{contract: contract}, nil
}

// bindEtherPaymentFallback binds a generic wrapper to an already deployed contract.
func bindEtherPaymentFallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherPaymentFallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherPaymentFallback *EtherPaymentFallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherPaymentFallback.Contract.EtherPaymentFallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherPaymentFallback *EtherPaymentFallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherPaymentFallback.Contract.EtherPaymentFallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherPaymentFallback *EtherPaymentFallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherPaymentFallback.Contract.EtherPaymentFallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherPaymentFallback *EtherPaymentFallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherPaymentFallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherPaymentFallback *EtherPaymentFallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherPaymentFallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherPaymentFallback *EtherPaymentFallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherPaymentFallback.Contract.contract.Transact(opts, method, params...)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherPaymentFallback *EtherPaymentFallbackTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherPaymentFallback.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherPaymentFallback *EtherPaymentFallbackSession) Receive() (*types.Transaction, error) {
	return _EtherPaymentFallback.Contract.Receive(&_EtherPaymentFallback.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherPaymentFallback *EtherPaymentFallbackTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherPaymentFallback.Contract.Receive(&_EtherPaymentFallback.TransactOpts)
}

// EtherPaymentFallbackSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the EtherPaymentFallback contract.
type EtherPaymentFallbackSafeReceivedIterator struct {
	Event *EtherPaymentFallbackSafeReceived // Event containing the contract specifics and raw log

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
func (it *EtherPaymentFallbackSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherPaymentFallbackSafeReceived)
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
		it.Event = new(EtherPaymentFallbackSafeReceived)
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
func (it *EtherPaymentFallbackSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherPaymentFallbackSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherPaymentFallbackSafeReceived represents a SafeReceived event raised by the EtherPaymentFallback contract.
type EtherPaymentFallbackSafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_EtherPaymentFallback *EtherPaymentFallbackFilterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*EtherPaymentFallbackSafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EtherPaymentFallback.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &EtherPaymentFallbackSafeReceivedIterator{contract: _EtherPaymentFallback.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_EtherPaymentFallback *EtherPaymentFallbackFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *EtherPaymentFallbackSafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EtherPaymentFallback.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherPaymentFallbackSafeReceived)
				if err := _EtherPaymentFallback.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_EtherPaymentFallback *EtherPaymentFallbackFilterer) ParseSafeReceived(log types.Log) (*EtherPaymentFallbackSafeReceived, error) {
	event := new(EtherPaymentFallbackSafeReceived)
	if err := _EtherPaymentFallback.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorMetaData contains all meta data concerning the Executor contract.
var ExecutorMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212206e5fb0215ddd85defefa8bb2203792bbb92078f782ad383c44bd8c25ebe51b9d64736f6c634300080d0033",
}

// ExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorMetaData.ABI instead.
var ExecutorABI = ExecutorMetaData.ABI

// ExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutorMetaData.Bin instead.
var ExecutorBin = ExecutorMetaData.Bin

// DeployExecutor deploys a new Ethereum contract, binding an instance of Executor to it.
func DeployExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// Executor is an auto generated Go binding around an Ethereum contract.
type Executor struct {
	ExecutorCaller     // Read-only binding to the contract
	ExecutorTransactor // Write-only binding to the contract
	ExecutorFilterer   // Log filterer for contract events
}

// ExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorSession struct {
	Contract     *Executor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorCallerSession struct {
	Contract *ExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorTransactorSession struct {
	Contract     *ExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorRaw struct {
	Contract *Executor // Generic contract binding to access the raw methods on
}

// ExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorCallerRaw struct {
	Contract *ExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorTransactorRaw struct {
	Contract *ExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutor creates a new instance of Executor, bound to a specific deployed contract.
func NewExecutor(address common.Address, backend bind.ContractBackend) (*Executor, error) {
	contract, err := bindExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// NewExecutorCaller creates a new read-only instance of Executor, bound to a specific deployed contract.
func NewExecutorCaller(address common.Address, caller bind.ContractCaller) (*ExecutorCaller, error) {
	contract, err := bindExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorCaller{contract: contract}, nil
}

// NewExecutorTransactor creates a new write-only instance of Executor, bound to a specific deployed contract.
func NewExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorTransactor, error) {
	contract, err := bindExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransactor{contract: contract}, nil
}

// NewExecutorFilterer creates a new log filterer instance of Executor, bound to a specific deployed contract.
func NewExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorFilterer, error) {
	contract, err := bindExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorFilterer{contract: contract}, nil
}

// bindExecutor binds a generic wrapper to an already deployed contract.
func bindExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.ExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transact(opts, method, params...)
}

// FallbackManagerMetaData contains all meta data concerning the FallbackManager contract.
var FallbackManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f08a0323": "setFallbackHandler(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506101ab806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f08a032314610084575b7f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d580548061005557005b36600080373360601b365260008060143601600080855af190503d6000803e8061007e573d6000fd5b503d6000f35b610097610092366004610145565b610099565b005b6100a1610108565b6100c9817f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d555565b6040516001600160a01b03821681527f5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b09060200160405180910390a150565b3330146101435760405162461bcd60e51b8152602060048201526005602482015264475330333160d81b604482015260640160405180910390fd5b565b60006020828403121561015757600080fd5b81356001600160a01b038116811461016e57600080fd5b939250505056fea264697066735822122097d0f0a8fbc23dd45dd16ca0f12c1e1ad263e50abe88e103bcdf8978912e8cbb64736f6c634300080d0033",
}

// FallbackManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FallbackManagerMetaData.ABI instead.
var FallbackManagerABI = FallbackManagerMetaData.ABI

// Deprecated: Use FallbackManagerMetaData.Sigs instead.
// FallbackManagerFuncSigs maps the 4-byte function signature to its string representation.
var FallbackManagerFuncSigs = FallbackManagerMetaData.Sigs

// FallbackManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FallbackManagerMetaData.Bin instead.
var FallbackManagerBin = FallbackManagerMetaData.Bin

// DeployFallbackManager deploys a new Ethereum contract, binding an instance of FallbackManager to it.
func DeployFallbackManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FallbackManager, error) {
	parsed, err := FallbackManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FallbackManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FallbackManager{FallbackManagerCaller: FallbackManagerCaller{contract: contract}, FallbackManagerTransactor: FallbackManagerTransactor{contract: contract}, FallbackManagerFilterer: FallbackManagerFilterer{contract: contract}}, nil
}

// FallbackManager is an auto generated Go binding around an Ethereum contract.
type FallbackManager struct {
	FallbackManagerCaller     // Read-only binding to the contract
	FallbackManagerTransactor // Write-only binding to the contract
	FallbackManagerFilterer   // Log filterer for contract events
}

// FallbackManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FallbackManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FallbackManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FallbackManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FallbackManagerSession struct {
	Contract     *FallbackManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FallbackManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FallbackManagerCallerSession struct {
	Contract *FallbackManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FallbackManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FallbackManagerTransactorSession struct {
	Contract     *FallbackManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FallbackManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FallbackManagerRaw struct {
	Contract *FallbackManager // Generic contract binding to access the raw methods on
}

// FallbackManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FallbackManagerCallerRaw struct {
	Contract *FallbackManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FallbackManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FallbackManagerTransactorRaw struct {
	Contract *FallbackManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFallbackManager creates a new instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManager(address common.Address, backend bind.ContractBackend) (*FallbackManager, error) {
	contract, err := bindFallbackManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FallbackManager{FallbackManagerCaller: FallbackManagerCaller{contract: contract}, FallbackManagerTransactor: FallbackManagerTransactor{contract: contract}, FallbackManagerFilterer: FallbackManagerFilterer{contract: contract}}, nil
}

// NewFallbackManagerCaller creates a new read-only instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManagerCaller(address common.Address, caller bind.ContractCaller) (*FallbackManagerCaller, error) {
	contract, err := bindFallbackManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackManagerCaller{contract: contract}, nil
}

// NewFallbackManagerTransactor creates a new write-only instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FallbackManagerTransactor, error) {
	contract, err := bindFallbackManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackManagerTransactor{contract: contract}, nil
}

// NewFallbackManagerFilterer creates a new log filterer instance of FallbackManager, bound to a specific deployed contract.
func NewFallbackManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FallbackManagerFilterer, error) {
	contract, err := bindFallbackManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FallbackManagerFilterer{contract: contract}, nil
}

// bindFallbackManager binds a generic wrapper to an already deployed contract.
func bindFallbackManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FallbackManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FallbackManager *FallbackManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FallbackManager.Contract.FallbackManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FallbackManager *FallbackManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FallbackManager.Contract.FallbackManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FallbackManager *FallbackManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FallbackManager.Contract.FallbackManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FallbackManager *FallbackManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FallbackManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FallbackManager *FallbackManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FallbackManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FallbackManager *FallbackManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FallbackManager.Contract.contract.Transact(opts, method, params...)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_FallbackManager *FallbackManagerTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _FallbackManager.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_FallbackManager *FallbackManagerSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _FallbackManager.Contract.SetFallbackHandler(&_FallbackManager.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_FallbackManager *FallbackManagerTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _FallbackManager.Contract.SetFallbackHandler(&_FallbackManager.TransactOpts, handler)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FallbackManager *FallbackManagerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FallbackManager.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FallbackManager *FallbackManagerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FallbackManager.Contract.Fallback(&_FallbackManager.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FallbackManager *FallbackManagerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FallbackManager.Contract.Fallback(&_FallbackManager.TransactOpts, calldata)
}

// FallbackManagerChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the FallbackManager contract.
type FallbackManagerChangedFallbackHandlerIterator struct {
	Event *FallbackManagerChangedFallbackHandler // Event containing the contract specifics and raw log

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
func (it *FallbackManagerChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FallbackManagerChangedFallbackHandler)
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
		it.Event = new(FallbackManagerChangedFallbackHandler)
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
func (it *FallbackManagerChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FallbackManagerChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FallbackManagerChangedFallbackHandler represents a ChangedFallbackHandler event raised by the FallbackManager contract.
type FallbackManagerChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_FallbackManager *FallbackManagerFilterer) FilterChangedFallbackHandler(opts *bind.FilterOpts) (*FallbackManagerChangedFallbackHandlerIterator, error) {

	logs, sub, err := _FallbackManager.contract.FilterLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return &FallbackManagerChangedFallbackHandlerIterator{contract: _FallbackManager.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_FallbackManager *FallbackManagerFilterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *FallbackManagerChangedFallbackHandler) (event.Subscription, error) {

	logs, sub, err := _FallbackManager.contract.WatchLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FallbackManagerChangedFallbackHandler)
				if err := _FallbackManager.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
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

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_FallbackManager *FallbackManagerFilterer) ParseChangedFallbackHandler(log types.Log) (*FallbackManagerChangedFallbackHandler, error) {
	event := new(FallbackManagerChangedFallbackHandler)
	if err := _FallbackManager.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeMetaData contains all meta data concerning the GnosisSafe contract.
var GnosisSafeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"approvedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApproveHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"SafeSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"encodeTransactionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"requiredTxGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"ffa1ad74": "VERSION()",
		"0d582f13": "addOwnerWithThreshold(address,uint256)",
		"d4d9bdcd": "approveHash(bytes32)",
		"7d832974": "approvedHashes(address,bytes32)",
		"694e80c3": "changeThreshold(uint256)",
		"12fb68e0": "checkNSignatures(bytes32,bytes,bytes,uint256)",
		"934f3a11": "checkSignatures(bytes32,bytes,bytes)",
		"e009cfde": "disableModule(address,address)",
		"f698da25": "domainSeparator()",
		"610b5925": "enableModule(address)",
		"e86637db": "encodeTransactionData(address,uint256,bytes,uint8,uint256,uint256,uint256,address,address,uint256)",
		"6a761202": "execTransaction(address,uint256,bytes,uint8,uint256,uint256,uint256,address,address,bytes)",
		"468721a7": "execTransactionFromModule(address,uint256,bytes,uint8)",
		"5229073f": "execTransactionFromModuleReturnData(address,uint256,bytes,uint8)",
		"3408e470": "getChainId()",
		"cc2f8452": "getModulesPaginated(address,uint256)",
		"a0e67e2b": "getOwners()",
		"5624b25b": "getStorageAt(uint256,uint256)",
		"e75235b8": "getThreshold()",
		"d8d11f78": "getTransactionHash(address,uint256,bytes,uint8,uint256,uint256,uint256,address,address,uint256)",
		"2d9ad53d": "isModuleEnabled(address)",
		"2f54bf6e": "isOwner(address)",
		"affed0e0": "nonce()",
		"f8dc5dd9": "removeOwner(address,address,uint256)",
		"c4ca3a9c": "requiredTxGas(address,uint256,bytes,uint8)",
		"f08a0323": "setFallbackHandler(address)",
		"e19a9dd9": "setGuard(address)",
		"b63e800d": "setup(address[],uint256,address,bytes,address,address,uint256,address)",
		"5ae6bd37": "signedMessages(bytes32)",
		"b4faba09": "simulateAndRevert(address,bytes)",
		"e318b52b": "swapOwner(address,address,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506001600455613101806100256000396000f3fe6080604052600436106101dc5760003560e01c8063affed0e011610102578063e19a9dd911610095578063f08a032311610064578063f08a032314610620578063f698da2514610640578063f8dc5dd914610655578063ffa1ad741461067557610218565b8063e19a9dd9146105ab578063e318b52b146105cb578063e75235b8146105eb578063e86637db1461060057610218565b8063cc2f8452116100d1578063cc2f84521461051d578063d4d9bdcd1461054b578063d8d11f781461056b578063e009cfde1461058b57610218565b8063affed0e0146104a7578063b4faba09146104bd578063b63e800d146104dd578063c4ca3a9c146104fd57610218565b80635624b25b1161017a5780636a761202116101495780636a7612021461041a5780637d8329741461042d578063934f3a1114610465578063a0e67e2b1461048557610218565b80635624b25b146103805780635ae6bd37146103ad578063610b5925146103da578063694e80c3146103fa57610218565b80632f54bf6e116101b65780632f54bf6e146102f55780633408e47014610315578063468721a7146103325780635229073f1461035257610218565b80630d582f131461027e57806312fb68e0146102a05780632d9ad53d146102c057610218565b366102185760405134815233907f3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d9060200160405180910390a2005b34801561022457600080fd5b507f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d580548061024f57005b36600080373360601b365260008060143601600080855af190503d6000803e80610278573d6000fd5b503d6000f35b34801561028a57600080fd5b5061029e610299366004612554565b6106a6565b005b3480156102ac57600080fd5b5061029e6102bb366004612622565b610806565b3480156102cc57600080fd5b506102e06102db366004612696565b610c71565b60405190151581526020015b60405180910390f35b34801561030157600080fd5b506102e0610310366004612696565b610cac565b34801561032157600080fd5b50465b6040519081526020016102ec565b34801561033e57600080fd5b506102e061034d3660046126c2565b610ce4565b34801561035e57600080fd5b5061037261036d3660046126c2565b610dbb565b6040516102ec929190612778565b34801561038c57600080fd5b506103a061039b366004612793565b610df1565b6040516102ec91906127b5565b3480156103b957600080fd5b506103246103c83660046127c8565b60076020526000908152604090205481565b3480156103e657600080fd5b5061029e6103f5366004612696565b610e76565b34801561040657600080fd5b5061029e6104153660046127c8565b610fb8565b6102e0610428366004612829565b611050565b34801561043957600080fd5b50610324610448366004612554565b600860209081526000928352604080842090915290825290205481565b34801561047157600080fd5b5061029e610480366004612901565b611399565b34801561049157600080fd5b5061049a6113e3565b6040516102ec91906129b1565b3480156104b357600080fd5b5061032460055481565b3480156104c957600080fd5b5061029e6104d83660046129c4565b6114d3565b3480156104e957600080fd5b5061029e6104f8366004612a13565b6114f6565b34801561050957600080fd5b50610324610518366004612b07565b611617565b34801561052957600080fd5b5061053d610538366004612554565b6116b1565b6040516102ec929190612b77565b34801561055757600080fd5b5061029e6105663660046127c8565b6117aa565b34801561057757600080fd5b50610324610586366004612ba1565b61183f565b34801561059757600080fd5b5061029e6105a6366004612c61565b61186c565b3480156105b757600080fd5b5061029e6105c6366004612696565b61199b565b3480156105d757600080fd5b5061029e6105e6366004612c9a565b611ab2565b3480156105f757600080fd5b50600454610324565b34801561060c57600080fd5b506103a061061b366004612ba1565b611ca1565b34801561062c57600080fd5b5061029e61063b366004612696565b611d7a565b34801561064c57600080fd5b50610324611de3565b34801561066157600080fd5b5061029e610670366004612ce5565b611e3a565b34801561068157600080fd5b506103a0604051806040016040528060058152602001640312e332e360dc1b81525081565b6106ae611fad565b6001600160a01b038216158015906106d057506001600160a01b038216600114155b80156106e557506001600160a01b0382163014155b61070a5760405162461bcd60e51b815260040161070190612d26565b60405180910390fd5b6001600160a01b0382811660009081526002602052604090205416156107425760405162461bcd60e51b815260040161070190612d45565b60026020527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e080546001600160a01b038481166000818152604081208054939094166001600160a01b0319938416179093556001835283549091161790915560038054916107af83612d7a565b90915550506040516001600160a01b03831681527f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea269060200160405180910390a180600454146108025761080281610fb8565b5050565b610811816041611fe6565b825110156108495760405162461bcd60e51b8152602060048201526005602482015264047533032360dc1b6044820152606401610701565b6000808060008060005b86811015610c65576041818102890160208101516040820151919092015160ff16955090935091506000849003610a24579193508391610894876041611fe6565b8210156108cb5760405162461bcd60e51b8152602060048201526005602482015264475330323160d81b6044820152606401610701565b87516108d8836020612022565b111561090e5760405162461bcd60e51b815260206004820152600560248201526423a998191960d91b6044820152606401610701565b60208289018101518951909161093190839061092b908790612022565b90612022565b11156109675760405162461bcd60e51b8152602060048201526005602482015264475330323360d81b6044820152606401610701565b6040516320c13b0b60e01b8082528a8501602001916001600160a01b038916906320c13b0b9061099d908f908690600401612d93565b602060405180830381865afa1580156109ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109de9190612db8565b6001600160e01b03191614610a1d5760405162461bcd60e51b815260206004820152600560248201526411d4cc0c8d60da1b6044820152606401610701565b5050610bcb565b8360ff16600103610aa6579193508391336001600160a01b0384161480610a6d57506001600160a01b03851660009081526008602090815260408083208d845290915290205415155b610aa15760405162461bcd60e51b8152602060048201526005602482015264475330323560d81b6044820152606401610701565b610bcb565b601e8460ff161115610b6b576040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018b9052600190605c0160405160208183030381529060405280519060200120600486610b0b9190612de2565b6040805160008152602081018083529390935260ff90911690820152606081018590526080810184905260a0016020604051602081039080840390855afa158015610b5a573d6000803e3d6000fd5b505050602060405103519450610bcb565b6040805160008152602081018083528c905260ff861691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610bbe573d6000803e3d6000fd5b5050506020604051035194505b856001600160a01b0316856001600160a01b0316118015610c0557506001600160a01b038581166000908152600260205260409020541615155b8015610c1b57506001600160a01b038516600114155b610c4f5760405162461bcd60e51b815260206004820152600560248201526423a998191b60d91b6044820152606401610701565b8495508080610c5d90612d7a565b915050610853565b50505050505050505050565b600060016001600160a01b03831614801590610ca657506001600160a01b038281166000908152600160205260409020541615155b92915050565b60006001600160a01b038216600114801590610ca65750506001600160a01b0390811660009081526002602052604090205416151590565b600033600114801590610d0e5750336000908152600160205260409020546001600160a01b031615155b610d425760405162461bcd60e51b815260206004820152600560248201526411d4cc4c0d60da1b6044820152606401610701565b610d4f858585855a61203e565b90508015610d875760405133907f6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb890600090a2610db3565b60405133907facd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd37590600090a25b949350505050565b60006060610dcb86868686610ce4565b915060405160203d0181016040523d81523d6000602083013e8091505094509492505050565b60606000610e00836020612e05565b6001600160401b03811115610e1757610e17612580565b6040519080825280601f01601f191660200182016040528015610e41576020820181803683370190505b50905060005b83811015610e6e578481015460208083028401015280610e6681612d7a565b915050610e47565b509392505050565b610e7e611fad565b6001600160a01b03811615801590610ea057506001600160a01b038116600114155b610ed45760405162461bcd60e51b8152602060048201526005602482015264475331303160d81b6044820152606401610701565b6001600160a01b038181166000908152600160205260409020541615610f245760405162461bcd60e51b815260206004820152600560248201526423a998981960d91b6044820152606401610701565b600160208181527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f80546001600160a01b03858116600081815260408082208054949095166001600160a01b031994851617909455959095528254168417909155519182527fecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f844091015b60405180910390a150565b610fc0611fad565b600354811115610fe25760405162461bcd60e51b815260040161070190612e24565b600181101561101b5760405162461bcd60e51b815260206004820152600560248201526423a999181960d91b6044820152606401610701565b60048190556040518181527f610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c9390602001610fad565b600080600061106a8e8e8e8e8e8e8e8e8e8e600554611ca1565b60058054919250600061107c83612d7a565b9091555050805160208201209150611095828286611399565b5060006110c07f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c85490565b90506001600160a01b0381161561114657806001600160a01b03166375f0bb528f8f8f8f8f8f8f8f8f8f8f336040518d63ffffffff1660e01b81526004016111139c9b9a99989796959493929190612e7b565b600060405180830381600087803b15801561112d57600080fd5b505af1158015611141573d6000803e3d6000fd5b505050505b6111726111558a6109c4612f40565b603f6111628c6040612e05565b61116c9190612f58565b90612085565b61117e906101f4612f40565b5a10156111b55760405162461bcd60e51b8152602060048201526005602482015264047533031360dc1b6044820152606401610701565b60005a90506112268f8f8f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e8c600014611213578e61203e565b6109c45a6112219190612f7a565b61203e565b93506112335a829061209c565b9050838061124057508915155b8061124a57508715155b61127e5760405162461bcd60e51b8152602060048201526005602482015264475330313360d81b6044820152606401610701565b6000881561129657611293828b8b8b8b6120b7565b90505b84156112da5760408051858152602081018390527f442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e910160405180910390a1611314565b60408051858152602081018390527f23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23910160405180910390a15b50506001600160a01b0381161561138857604051631264e26d60e31b81526004810183905283151560248201526001600160a01b03821690639327136890604401600060405180830381600087803b15801561136f57600080fd5b505af1158015611383573d6000803e3d6000fd5b505050505b50509b9a5050505050505050505050565b600454806113d15760405162461bcd60e51b8152602060048201526005602482015264475330303160d81b6044820152606401610701565b6113dd84848484610806565b50505050565b606060006003546001600160401b0381111561140157611401612580565b60405190808252806020026020018201604052801561142a578160200160208202803683370190505b506001600090815260026020527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e054919250906001600160a01b03165b6001600160a01b0381166001146114cb578083838151811061148b5761148b612f91565b6001600160a01b039283166020918202929092018101919091529181166000908152600290925260409091205416816114c381612d7a565b925050611467565b509092915050565b600080825160208401855af480600052503d6020523d600060403e60403d016000fd5b6115348a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508c92506121bd915050565b6001600160a01b0384161561156b5761156b847f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d555565b6115ab8787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506123a392505050565b81156115c2576115c0826000600186856120b7565b505b336001600160a01b03167f141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a88b8b8b8b89604051611603959493929190612fa7565b60405180910390a250505050505050505050565b6000805a9050611660878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525089925050505a61203e565b61166957600080fd5b60005a6116769083612f7a565b90508060405160200161168b91815260200190565b60408051601f198184030181529082905262461bcd60e51b8252610701916004016127b5565b60606000826001600160401b038111156116cd576116cd612580565b6040519080825280602002602001820160405280156116f6578160200160208202803683370190505b506001600160a01b0380861660009081526001602052604081205492945091165b6001600160a01b0381161580159061173957506001600160a01b038116600114155b801561174457508482105b1561179c578084838151811061175c5761175c612f91565b6001600160a01b0392831660209182029290920181019190915291811660009081526001909252604090912054168161179481612d7a565b925050611717565b908352919491935090915050565b336000908152600260205260409020546001600160a01b03166117f75760405162461bcd60e51b8152602060048201526005602482015264047533033360dc1b6044820152606401610701565b336000818152600860209081526040808320858452909152808220600190555183917ff2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c91a350565b60006118548c8c8c8c8c8c8c8c8c8c8c611ca1565b8051906020012090509b9a5050505050505050505050565b611874611fad565b6001600160a01b0381161580159061189657506001600160a01b038116600114155b6118ca5760405162461bcd60e51b8152602060048201526005602482015264475331303160d81b6044820152606401610701565b6001600160a01b0382811660009081526001602052604090205481169082161461191e5760405162461bcd60e51b8152602060048201526005602482015264475331303360d81b6044820152606401610701565b6001600160a01b038181166000818152600160209081526040808320805488871685528285208054919097166001600160a01b03199182161790965592849052825490941690915591519081527faab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace405427691015b60405180910390a15050565b6119a3611fad565b6001600160a01b03811615611a55576040516301ffc9a760e01b815263736bd41d60e11b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa1580156119fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a219190613013565b611a555760405162461bcd60e51b8152602060048201526005602482015264047533330360dc1b6044820152606401610701565b7f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c88181556040516001600160a01b03831681527f1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa29060200161198f565b611aba611fad565b6001600160a01b03811615801590611adc57506001600160a01b038116600114155b8015611af157506001600160a01b0381163014155b611b0d5760405162461bcd60e51b815260040161070190612d26565b6001600160a01b038181166000908152600260205260409020541615611b455760405162461bcd60e51b815260040161070190612d45565b6001600160a01b03821615801590611b6757506001600160a01b038216600114155b611b835760405162461bcd60e51b815260040161070190612d26565b6001600160a01b03838116600090815260026020526040902054811690831614611bd75760405162461bcd60e51b8152602060048201526005602482015264475332303560d81b6044820152606401610701565b6001600160a01b038281166000818152600260209081526040808320805487871680865283862080549289166001600160a01b0319938416179055968a1685528285208054821690971790965592849052825490941690915591519081527ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf910160405180910390a16040516001600160a01b03821681527f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea269060200160405180910390a1505050565b606060007fbb8310d486368db6bd6f849402fdd73ad53d316b5a4b2644ad6efe0f941286d860001b8d8d8d8d604051611cdb929190613035565b604051908190038120611d01949392918e908e908e908e908e908e908e90602001613045565b60408051601f1981840301815291905280516020909101209050601960f81b600160f81b611d2d611de3565b6040516001600160f81b031993841660208201529290911660218301526022820152604281018290526062016040516020818303038152906040529150509b9a5050505050505050505050565b611d82611fad565b611daa817f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d555565b6040516001600160a01b03821681527f5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b090602001610fad565b60007f47e79534a245952e8b16893a336b85a3d9ea9fa8c573f3d803afb92a794692184660408051602081019390935282015230606082015260800160405160208183030381529060405280519060200120905090565b611e42611fad565b806001600354611e529190612f7a565b1015611e705760405162461bcd60e51b815260040161070190612e24565b6001600160a01b03821615801590611e9257506001600160a01b038216600114155b611eae5760405162461bcd60e51b815260040161070190612d26565b6001600160a01b03838116600090815260026020526040902054811690831614611f025760405162461bcd60e51b8152602060048201526005602482015264475332303560d81b6044820152606401610701565b6001600160a01b03828116600081815260026020526040808220805488861684529183208054929095166001600160a01b03199283161790945591815282549091169091556003805491611f55836130b4565b90915550506040516001600160a01b03831681527ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf9060200160405180910390a18060045414611fa857611fa881610fb8565b505050565b333014611fe45760405162461bcd60e51b8152602060048201526005602482015264475330333160d81b6044820152606401610701565b565b600082600003611ff857506000610ca6565b60006120048385612e05565b9050826120118583612f58565b1461201b57600080fd5b9392505050565b60008061202f8385612f40565b90508381101561201b57600080fd5b6000600183600181111561205457612054612e43565b0361206c576000808551602087018986f4905061207c565b600080855160208701888a87f190505b95945050505050565b600081831015612095578161201b565b5090919050565b6000828211156120ab57600080fd5b6000610db38385612f7a565b6000806001600160a01b038316156120cf57826120d1565b325b90506001600160a01b038416612164576121033a86106120f1573a6120f3565b855b6120fd8989612022565b90611fe6565b6040519092506001600160a01b0382169083156108fc029084906000818181858888f1935050505061215f5760405162461bcd60e51b8152602060048201526005602482015264475330313160d81b6044820152606401610701565b6121b3565b612172856120fd8989612022565b915061217f84828461249d565b6121b35760405162461bcd60e51b815260206004820152600560248201526423a998189960d91b6044820152606401610701565b5095945050505050565b600454156121f55760405162461bcd60e51b8152602060048201526005602482015264047533230360dc1b6044820152606401610701565b81518111156122165760405162461bcd60e51b815260040161070190612e24565b600181101561224f5760405162461bcd60e51b815260206004820152600560248201526423a999181960d91b6044820152606401610701565b600160005b835181101561237057600084828151811061227157612271612f91565b6020026020010151905060006001600160a01b0316816001600160a01b0316141580156122a857506001600160a01b038116600114155b80156122bd57506001600160a01b0381163014155b80156122db5750806001600160a01b0316836001600160a01b031614155b6122f75760405162461bcd60e51b815260040161070190612d26565b6001600160a01b03818116600090815260026020526040902054161561232f5760405162461bcd60e51b815260040161070190612d45565b6001600160a01b03928316600090815260026020526040902080546001600160a01b031916938216939093179092558061236881612d7a565b915050612254565b506001600160a01b0316600090815260026020526040902080546001600160a01b03191660011790559051600355600455565b600160008190526020527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f546001600160a01b03161561240d5760405162461bcd60e51b8152602060048201526005602482015264047533130360dc1b6044820152606401610701565b6001600081905260208190527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f80546001600160a01b03191690911790556001600160a01b03821615610802576124698260008360015a61203e565b6108025760405162461bcd60e51b8152602060048201526005602482015264047533030360dc1b6044820152606401610701565b604080516001600160a01b03841660248201526044808201849052825180830390910181526064909101909152602080820180516001600160e01b031663a9059cbb60e01b1781528251600093929184919082896127105a03f13d801561250f57602081146125175760009350612522565b819350612522565b600051158215171593505b5050509392505050565b6001600160a01b038116811461254157600080fd5b50565b803561254f8161252c565b919050565b6000806040838503121561256757600080fd5b82356125728161252c565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126125a757600080fd5b81356001600160401b03808211156125c1576125c1612580565b604051601f8301601f19908116603f011681019082821181831017156125e9576125e9612580565b8160405283815286602085880101111561260257600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561263857600080fd5b8435935060208501356001600160401b038082111561265657600080fd5b61266288838901612596565b9450604087013591508082111561267857600080fd5b5061268587828801612596565b949793965093946060013593505050565b6000602082840312156126a857600080fd5b813561201b8161252c565b80356002811061254f57600080fd5b600080600080608085870312156126d857600080fd5b84356126e38161252c565b93506020850135925060408501356001600160401b0381111561270557600080fd5b61271187828801612596565b925050612720606086016126b3565b905092959194509250565b6000815180845260005b8181101561275157602081850181015186830182015201612735565b81811115612763576000602083870101525b50601f01601f19169290920160200192915050565b8215158152604060208201526000610db3604083018461272b565b600080604083850312156127a657600080fd5b50508035926020909101359150565b60208152600061201b602083018461272b565b6000602082840312156127da57600080fd5b5035919050565b60008083601f8401126127f357600080fd5b5081356001600160401b0381111561280a57600080fd5b60208301915083602082850101111561282257600080fd5b9250929050565b60008060008060008060008060008060006101408c8e03121561284b57600080fd5b6128548c612544565b9a5060208c013599506001600160401b038060408e0135111561287657600080fd5b6128868e60408f01358f016127e1565b909a50985061289760608e016126b3565b975060808d0135965060a08d0135955060c08d013594506128ba60e08e01612544565b93506128c96101008e01612544565b9250806101208e013511156128dd57600080fd5b506128ef8d6101208e01358e01612596565b90509295989b509295989b9093969950565b60008060006060848603121561291657600080fd5b8335925060208401356001600160401b038082111561293457600080fd5b61294087838801612596565b9350604086013591508082111561295657600080fd5b5061296386828701612596565b9150509250925092565b600081518084526020808501945080840160005b838110156129a65781516001600160a01b031687529582019590820190600101612981565b509495945050505050565b60208152600061201b602083018461296d565b600080604083850312156129d757600080fd5b82356129e28161252c565b915060208301356001600160401b038111156129fd57600080fd5b612a0985828601612596565b9150509250929050565b6000806000806000806000806000806101008b8d031215612a3357600080fd5b8a356001600160401b0380821115612a4a57600080fd5b818d0191508d601f830112612a5e57600080fd5b813581811115612a6d57600080fd5b8e60208260051b8501011115612a8257600080fd5b60208381019d50909b508d01359950612a9d60408e01612544565b985060608d0135915080821115612ab357600080fd5b50612ac08d828e016127e1565b9097509550612ad3905060808c01612544565b9350612ae160a08c01612544565b925060c08b01359150612af660e08c01612544565b90509295989b9194979a5092959850565b600080600080600060808688031215612b1f57600080fd5b8535612b2a8161252c565b94506020860135935060408601356001600160401b03811115612b4c57600080fd5b612b58888289016127e1565b9094509250612b6b9050606087016126b3565b90509295509295909350565b604081526000612b8a604083018561296d565b905060018060a01b03831660208301529392505050565b60008060008060008060008060008060006101408c8e031215612bc357600080fd5b8b35612bce8161252c565b9a5060208c0135995060408c01356001600160401b03811115612bf057600080fd5b612bfc8e828f016127e1565b909a509850612c0f905060608d016126b3565b965060808c0135955060a08c0135945060c08c0135935060e08c0135612c348161252c565b92506101008c0135612c458161252c565b809250506101208c013590509295989b509295989b9093969950565b60008060408385031215612c7457600080fd5b8235612c7f8161252c565b91506020830135612c8f8161252c565b809150509250929050565b600080600060608486031215612caf57600080fd5b8335612cba8161252c565b92506020840135612cca8161252c565b91506040840135612cda8161252c565b809150509250925092565b600080600060608486031215612cfa57600080fd5b8335612d058161252c565b92506020840135612d158161252c565b929592945050506040919091013590565b602080825260059082015264475332303360d81b604082015260600190565b60208082526005908201526411d4cc8c0d60da1b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600060018201612d8c57612d8c612d64565b5060010190565b604081526000612da6604083018561272b565b828103602084015261207c818561272b565b600060208284031215612dca57600080fd5b81516001600160e01b03198116811461201b57600080fd5b600060ff821660ff841680821015612dfc57612dfc612d64565b90039392505050565b6000816000190483118215151615612e1f57612e1f612d64565b500290565b602080825260059082015264475332303160d81b604082015260600190565b634e487b7160e01b600052602160045260246000fd5b60028110612e7757634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b038d168152602081018c90526101606040820181905281018a905260006101808b8d828501376000838d01820152601f8c01601f19168301612ec7606085018d612e59565b8a60808501528960a08501528860c0850152612eee60e08501896001600160a01b03169052565b6001600160a01b0387166101008501528184820301610120850152612f158282018761272b565b92505050612f2f6101408301846001600160a01b03169052565b9d9c50505050505050505050505050565b60008219821115612f5357612f53612d64565b500190565b600082612f7557634e487b7160e01b600052601260045260246000fd5b500490565b600082821015612f8c57612f8c612d64565b500390565b634e487b7160e01b600052603260045260246000fd5b6080808252810185905260008660a08301825b88811015612fea578235612fcd8161252c565b6001600160a01b0316825260209283019290910190600101612fba565b50602084019690965250506001600160a01b039283166040820152911660609091015292915050565b60006020828403121561302557600080fd5b8151801515811461201b57600080fd5b8183823760009101908152919050565b8b81526001600160a01b038b81166020830152604082018b9052606082018a9052610160820190613079608084018b612e59565b60a083019890985260c082019690965260e0810194909452918516610100840152909316610120820152610140019190915295945050505050565b6000816130c3576130c3612d64565b50600019019056fea2646970667358221220c6d941cdf665bda02888a500669529b03010645e453d4d9bfae8c8beb2b355c964736f6c634300080d0033",
}

// GnosisSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use GnosisSafeMetaData.ABI instead.
var GnosisSafeABI = GnosisSafeMetaData.ABI

// Deprecated: Use GnosisSafeMetaData.Sigs instead.
// GnosisSafeFuncSigs maps the 4-byte function signature to its string representation.
var GnosisSafeFuncSigs = GnosisSafeMetaData.Sigs

// GnosisSafeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GnosisSafeMetaData.Bin instead.
var GnosisSafeBin = GnosisSafeMetaData.Bin

// DeployGnosisSafe deploys a new Ethereum contract, binding an instance of GnosisSafe to it.
func DeployGnosisSafe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GnosisSafe, error) {
	parsed, err := GnosisSafeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GnosisSafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GnosisSafe{GnosisSafeCaller: GnosisSafeCaller{contract: contract}, GnosisSafeTransactor: GnosisSafeTransactor{contract: contract}, GnosisSafeFilterer: GnosisSafeFilterer{contract: contract}}, nil
}

// GnosisSafe is an auto generated Go binding around an Ethereum contract.
type GnosisSafe struct {
	GnosisSafeCaller     // Read-only binding to the contract
	GnosisSafeTransactor // Write-only binding to the contract
	GnosisSafeFilterer   // Log filterer for contract events
}

// GnosisSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeSession struct {
	Contract     *GnosisSafe       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnosisSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeCallerSession struct {
	Contract *GnosisSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GnosisSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeTransactorSession struct {
	Contract     *GnosisSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GnosisSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeRaw struct {
	Contract *GnosisSafe // Generic contract binding to access the raw methods on
}

// GnosisSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeCallerRaw struct {
	Contract *GnosisSafeCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeTransactorRaw struct {
	Contract *GnosisSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafe creates a new instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafe(address common.Address, backend bind.ContractBackend) (*GnosisSafe, error) {
	contract, err := bindGnosisSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafe{GnosisSafeCaller: GnosisSafeCaller{contract: contract}, GnosisSafeTransactor: GnosisSafeTransactor{contract: contract}, GnosisSafeFilterer: GnosisSafeFilterer{contract: contract}}, nil
}

// NewGnosisSafeCaller creates a new read-only instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafeCaller(address common.Address, caller bind.ContractCaller) (*GnosisSafeCaller, error) {
	contract, err := bindGnosisSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeCaller{contract: contract}, nil
}

// NewGnosisSafeTransactor creates a new write-only instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeTransactor, error) {
	contract, err := bindGnosisSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeTransactor{contract: contract}, nil
}

// NewGnosisSafeFilterer creates a new log filterer instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeFilterer, error) {
	contract, err := bindGnosisSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeFilterer{contract: contract}, nil
}

// bindGnosisSafe binds a generic wrapper to an already deployed contract.
func bindGnosisSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisSafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafe *GnosisSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafe.Contract.GnosisSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafe *GnosisSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafe.Contract.GnosisSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafe *GnosisSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafe.Contract.GnosisSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafe *GnosisSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafe *GnosisSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafe *GnosisSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafe.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafe *GnosisSafeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafe *GnosisSafeSession) VERSION() (string, error) {
	return _GnosisSafe.Contract.VERSION(&_GnosisSafe.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafe *GnosisSafeCallerSession) VERSION() (string, error) {
	return _GnosisSafe.Contract.VERSION(&_GnosisSafe.CallOpts)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) ApprovedHashes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "approvedHashes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.ApprovedHashes(&_GnosisSafe.CallOpts, arg0, arg1)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.ApprovedHashes(&_GnosisSafe.CallOpts, arg0, arg1)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafe *GnosisSafeCaller) CheckNSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "checkNSignatures", dataHash, data, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafe *GnosisSafeSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _GnosisSafe.Contract.CheckNSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafe *GnosisSafeCallerSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _GnosisSafe.Contract.CheckNSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafe *GnosisSafeCaller) CheckSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte) error {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "checkSignatures", dataHash, data, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafe *GnosisSafeSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _GnosisSafe.Contract.CheckSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafe *GnosisSafeCallerSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _GnosisSafe.Contract.CheckSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafe *GnosisSafeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafe *GnosisSafeSession) DomainSeparator() ([32]byte, error) {
	return _GnosisSafe.Contract.DomainSeparator(&_GnosisSafe.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafe *GnosisSafeCallerSession) DomainSeparator() ([32]byte, error) {
	return _GnosisSafe.Contract.DomainSeparator(&_GnosisSafe.CallOpts)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafe *GnosisSafeCaller) EncodeTransactionData(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "encodeTransactionData", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafe *GnosisSafeSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.EncodeTransactionData(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafe *GnosisSafeCallerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.EncodeTransactionData(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) GetChainId() (*big.Int, error) {
	return _GnosisSafe.Contract.GetChainId(&_GnosisSafe.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) GetChainId() (*big.Int, error) {
	return _GnosisSafe.Contract.GetChainId(&_GnosisSafe.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafe *GnosisSafeCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafe *GnosisSafeSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _GnosisSafe.Contract.GetModulesPaginated(&_GnosisSafe.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafe *GnosisSafeCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _GnosisSafe.Contract.GetModulesPaginated(&_GnosisSafe.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafe *GnosisSafeCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafe *GnosisSafeSession) GetOwners() ([]common.Address, error) {
	return _GnosisSafe.Contract.GetOwners(&_GnosisSafe.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafe *GnosisSafeCallerSession) GetOwners() ([]common.Address, error) {
	return _GnosisSafe.Contract.GetOwners(&_GnosisSafe.CallOpts)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafe *GnosisSafeCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafe *GnosisSafeSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.GetStorageAt(&_GnosisSafe.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafe *GnosisSafeCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.GetStorageAt(&_GnosisSafe.CallOpts, offset, length)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) GetThreshold() (*big.Int, error) {
	return _GnosisSafe.Contract.GetThreshold(&_GnosisSafe.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) GetThreshold() (*big.Int, error) {
	return _GnosisSafe.Contract.GetThreshold(&_GnosisSafe.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafe *GnosisSafeCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafe *GnosisSafeSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafe.Contract.GetTransactionHash(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafe *GnosisSafeCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafe.Contract.GetTransactionHash(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafe *GnosisSafeCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafe *GnosisSafeSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsModuleEnabled(&_GnosisSafe.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafe *GnosisSafeCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsModuleEnabled(&_GnosisSafe.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafe *GnosisSafeCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafe *GnosisSafeSession) IsOwner(owner common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsOwner(&_GnosisSafe.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafe *GnosisSafeCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsOwner(&_GnosisSafe.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) Nonce() (*big.Int, error) {
	return _GnosisSafe.Contract.Nonce(&_GnosisSafe.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) Nonce() (*big.Int, error) {
	return _GnosisSafe.Contract.Nonce(&_GnosisSafe.CallOpts)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) SignedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "signedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.SignedMessages(&_GnosisSafe.CallOpts, arg0)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.SignedMessages(&_GnosisSafe.CallOpts, arg0)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.AddOwnerWithThreshold(&_GnosisSafe.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.AddOwnerWithThreshold(&_GnosisSafe.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafe *GnosisSafeTransactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafe *GnosisSafeSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ApproveHash(&_GnosisSafe.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ApproveHash(&_GnosisSafe.TransactOpts, hashToApprove)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ChangeThreshold(&_GnosisSafe.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ChangeThreshold(&_GnosisSafe.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafe *GnosisSafeTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafe *GnosisSafeSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.DisableModule(&_GnosisSafe.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.DisableModule(&_GnosisSafe.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafe *GnosisSafeTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafe *GnosisSafeSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.EnableModule(&_GnosisSafe.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.EnableModule(&_GnosisSafe.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafe *GnosisSafeTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafe *GnosisSafeSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransaction(&_GnosisSafe.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafe *GnosisSafeTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransaction(&_GnosisSafe.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafe *GnosisSafeTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafe *GnosisSafeSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModule(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafe *GnosisSafeTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModule(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafe *GnosisSafeTransactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafe *GnosisSafeSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModuleReturnData(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafe *GnosisSafeTransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModuleReturnData(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RemoveOwner(&_GnosisSafe.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RemoveOwner(&_GnosisSafe.TransactOpts, prevOwner, owner, _threshold)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_GnosisSafe *GnosisSafeTransactor) RequiredTxGas(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "requiredTxGas", to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_GnosisSafe *GnosisSafeSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RequiredTxGas(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_GnosisSafe *GnosisSafeTransactorSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RequiredTxGas(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafe *GnosisSafeTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafe *GnosisSafeSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetFallbackHandler(&_GnosisSafe.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetFallbackHandler(&_GnosisSafe.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafe *GnosisSafeTransactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafe *GnosisSafeSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetGuard(&_GnosisSafe.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetGuard(&_GnosisSafe.TransactOpts, guard)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafe *GnosisSafeTransactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "setup", _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafe *GnosisSafeSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Setup(&_GnosisSafe.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Setup(&_GnosisSafe.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafe *GnosisSafeTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafe *GnosisSafeSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SimulateAndRevert(&_GnosisSafe.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SimulateAndRevert(&_GnosisSafe.TransactOpts, targetContract, calldataPayload)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafe *GnosisSafeTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafe *GnosisSafeSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SwapOwner(&_GnosisSafe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SwapOwner(&_GnosisSafe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafe *GnosisSafeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafe *GnosisSafeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Fallback(&_GnosisSafe.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafe *GnosisSafeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Fallback(&_GnosisSafe.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafe *GnosisSafeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafe.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafe *GnosisSafeSession) Receive() (*types.Transaction, error) {
	return _GnosisSafe.Contract.Receive(&_GnosisSafe.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafe *GnosisSafeTransactorSession) Receive() (*types.Transaction, error) {
	return _GnosisSafe.Contract.Receive(&_GnosisSafe.TransactOpts)
}

// GnosisSafeAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the GnosisSafe contract.
type GnosisSafeAddedOwnerIterator struct {
	Event *GnosisSafeAddedOwner // Event containing the contract specifics and raw log

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
func (it *GnosisSafeAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeAddedOwner)
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
		it.Event = new(GnosisSafeAddedOwner)
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
func (it *GnosisSafeAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeAddedOwner represents a AddedOwner event raised by the GnosisSafe contract.
type GnosisSafeAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) FilterAddedOwner(opts *bind.FilterOpts) (*GnosisSafeAddedOwnerIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeAddedOwnerIterator{contract: _GnosisSafe.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *GnosisSafeAddedOwner) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeAddedOwner)
				if err := _GnosisSafe.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) ParseAddedOwner(log types.Log) (*GnosisSafeAddedOwner, error) {
	event := new(GnosisSafeAddedOwner)
	if err := _GnosisSafe.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeApproveHashIterator is returned from FilterApproveHash and is used to iterate over the raw logs and unpacked data for ApproveHash events raised by the GnosisSafe contract.
type GnosisSafeApproveHashIterator struct {
	Event *GnosisSafeApproveHash // Event containing the contract specifics and raw log

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
func (it *GnosisSafeApproveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeApproveHash)
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
		it.Event = new(GnosisSafeApproveHash)
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
func (it *GnosisSafeApproveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeApproveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeApproveHash represents a ApproveHash event raised by the GnosisSafe contract.
type GnosisSafeApproveHash struct {
	ApprovedHash [32]byte
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproveHash is a free log retrieval operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafe *GnosisSafeFilterer) FilterApproveHash(opts *bind.FilterOpts, approvedHash [][32]byte, owner []common.Address) (*GnosisSafeApproveHashIterator, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeApproveHashIterator{contract: _GnosisSafe.contract, event: "ApproveHash", logs: logs, sub: sub}, nil
}

// WatchApproveHash is a free log subscription operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafe *GnosisSafeFilterer) WatchApproveHash(opts *bind.WatchOpts, sink chan<- *GnosisSafeApproveHash, approvedHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeApproveHash)
				if err := _GnosisSafe.contract.UnpackLog(event, "ApproveHash", log); err != nil {
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

// ParseApproveHash is a log parse operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafe *GnosisSafeFilterer) ParseApproveHash(log types.Log) (*GnosisSafeApproveHash, error) {
	event := new(GnosisSafeApproveHash)
	if err := _GnosisSafe.contract.UnpackLog(event, "ApproveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the GnosisSafe contract.
type GnosisSafeChangedFallbackHandlerIterator struct {
	Event *GnosisSafeChangedFallbackHandler // Event containing the contract specifics and raw log

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
func (it *GnosisSafeChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeChangedFallbackHandler)
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
		it.Event = new(GnosisSafeChangedFallbackHandler)
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
func (it *GnosisSafeChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeChangedFallbackHandler represents a ChangedFallbackHandler event raised by the GnosisSafe contract.
type GnosisSafeChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_GnosisSafe *GnosisSafeFilterer) FilterChangedFallbackHandler(opts *bind.FilterOpts) (*GnosisSafeChangedFallbackHandlerIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeChangedFallbackHandlerIterator{contract: _GnosisSafe.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_GnosisSafe *GnosisSafeFilterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *GnosisSafeChangedFallbackHandler) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeChangedFallbackHandler)
				if err := _GnosisSafe.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
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

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_GnosisSafe *GnosisSafeFilterer) ParseChangedFallbackHandler(log types.Log) (*GnosisSafeChangedFallbackHandler, error) {
	event := new(GnosisSafeChangedFallbackHandler)
	if err := _GnosisSafe.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the GnosisSafe contract.
type GnosisSafeChangedGuardIterator struct {
	Event *GnosisSafeChangedGuard // Event containing the contract specifics and raw log

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
func (it *GnosisSafeChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeChangedGuard)
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
		it.Event = new(GnosisSafeChangedGuard)
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
func (it *GnosisSafeChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeChangedGuard represents a ChangedGuard event raised by the GnosisSafe contract.
type GnosisSafeChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GnosisSafe *GnosisSafeFilterer) FilterChangedGuard(opts *bind.FilterOpts) (*GnosisSafeChangedGuardIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeChangedGuardIterator{contract: _GnosisSafe.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GnosisSafe *GnosisSafeFilterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *GnosisSafeChangedGuard) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeChangedGuard)
				if err := _GnosisSafe.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
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

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GnosisSafe *GnosisSafeFilterer) ParseChangedGuard(log types.Log) (*GnosisSafeChangedGuard, error) {
	event := new(GnosisSafeChangedGuard)
	if err := _GnosisSafe.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the GnosisSafe contract.
type GnosisSafeChangedThresholdIterator struct {
	Event *GnosisSafeChangedThreshold // Event containing the contract specifics and raw log

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
func (it *GnosisSafeChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeChangedThreshold)
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
		it.Event = new(GnosisSafeChangedThreshold)
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
func (it *GnosisSafeChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeChangedThreshold represents a ChangedThreshold event raised by the GnosisSafe contract.
type GnosisSafeChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafe *GnosisSafeFilterer) FilterChangedThreshold(opts *bind.FilterOpts) (*GnosisSafeChangedThresholdIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeChangedThresholdIterator{contract: _GnosisSafe.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafe *GnosisSafeFilterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *GnosisSafeChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeChangedThreshold)
				if err := _GnosisSafe.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
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

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafe *GnosisSafeFilterer) ParseChangedThreshold(log types.Log) (*GnosisSafeChangedThreshold, error) {
	event := new(GnosisSafeChangedThreshold)
	if err := _GnosisSafe.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the GnosisSafe contract.
type GnosisSafeDisabledModuleIterator struct {
	Event *GnosisSafeDisabledModule // Event containing the contract specifics and raw log

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
func (it *GnosisSafeDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeDisabledModule)
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
		it.Event = new(GnosisSafeDisabledModule)
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
func (it *GnosisSafeDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeDisabledModule represents a DisabledModule event raised by the GnosisSafe contract.
type GnosisSafeDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) FilterDisabledModule(opts *bind.FilterOpts) (*GnosisSafeDisabledModuleIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeDisabledModuleIterator{contract: _GnosisSafe.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *GnosisSafeDisabledModule) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeDisabledModule)
				if err := _GnosisSafe.contract.UnpackLog(event, "DisabledModule", log); err != nil {
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

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) ParseDisabledModule(log types.Log) (*GnosisSafeDisabledModule, error) {
	event := new(GnosisSafeDisabledModule)
	if err := _GnosisSafe.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the GnosisSafe contract.
type GnosisSafeEnabledModuleIterator struct {
	Event *GnosisSafeEnabledModule // Event containing the contract specifics and raw log

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
func (it *GnosisSafeEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeEnabledModule)
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
		it.Event = new(GnosisSafeEnabledModule)
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
func (it *GnosisSafeEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeEnabledModule represents a EnabledModule event raised by the GnosisSafe contract.
type GnosisSafeEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) FilterEnabledModule(opts *bind.FilterOpts) (*GnosisSafeEnabledModuleIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeEnabledModuleIterator{contract: _GnosisSafe.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *GnosisSafeEnabledModule) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeEnabledModule)
				if err := _GnosisSafe.contract.UnpackLog(event, "EnabledModule", log); err != nil {
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

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) ParseEnabledModule(log types.Log) (*GnosisSafeEnabledModule, error) {
	event := new(GnosisSafeEnabledModule)
	if err := _GnosisSafe.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the GnosisSafe contract.
type GnosisSafeExecutionFailureIterator struct {
	Event *GnosisSafeExecutionFailure // Event containing the contract specifics and raw log

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
func (it *GnosisSafeExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionFailure)
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
		it.Event = new(GnosisSafeExecutionFailure)
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
func (it *GnosisSafeExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionFailure represents a ExecutionFailure event raised by the GnosisSafe contract.
type GnosisSafeExecutionFailure struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionFailure(opts *bind.FilterOpts) (*GnosisSafeExecutionFailureIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionFailureIterator{contract: _GnosisSafe.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionFailure) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionFailure)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionFailure(log types.Log) (*GnosisSafeExecutionFailure, error) {
	event := new(GnosisSafeExecutionFailure)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleFailureIterator struct {
	Event *GnosisSafeExecutionFromModuleFailure // Event containing the contract specifics and raw log

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
func (it *GnosisSafeExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionFromModuleFailure)
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
		it.Event = new(GnosisSafeExecutionFromModuleFailure)
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
func (it *GnosisSafeExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionFromModuleFailureIterator{contract: _GnosisSafe.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionFromModuleFailure)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
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

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionFromModuleFailure(log types.Log) (*GnosisSafeExecutionFromModuleFailure, error) {
	event := new(GnosisSafeExecutionFromModuleFailure)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleSuccessIterator struct {
	Event *GnosisSafeExecutionFromModuleSuccess // Event containing the contract specifics and raw log

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
func (it *GnosisSafeExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionFromModuleSuccess)
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
		it.Event = new(GnosisSafeExecutionFromModuleSuccess)
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
func (it *GnosisSafeExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionFromModuleSuccessIterator{contract: _GnosisSafe.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionFromModuleSuccess)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
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

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionFromModuleSuccess(log types.Log) (*GnosisSafeExecutionFromModuleSuccess, error) {
	event := new(GnosisSafeExecutionFromModuleSuccess)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the GnosisSafe contract.
type GnosisSafeExecutionSuccessIterator struct {
	Event *GnosisSafeExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *GnosisSafeExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionSuccess)
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
		it.Event = new(GnosisSafeExecutionSuccess)
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
func (it *GnosisSafeExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionSuccess represents a ExecutionSuccess event raised by the GnosisSafe contract.
type GnosisSafeExecutionSuccess struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionSuccess(opts *bind.FilterOpts) (*GnosisSafeExecutionSuccessIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionSuccessIterator{contract: _GnosisSafe.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionSuccess) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionSuccess)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionSuccess(log types.Log) (*GnosisSafeExecutionSuccess, error) {
	event := new(GnosisSafeExecutionSuccess)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the GnosisSafe contract.
type GnosisSafeRemovedOwnerIterator struct {
	Event *GnosisSafeRemovedOwner // Event containing the contract specifics and raw log

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
func (it *GnosisSafeRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeRemovedOwner)
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
		it.Event = new(GnosisSafeRemovedOwner)
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
func (it *GnosisSafeRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeRemovedOwner represents a RemovedOwner event raised by the GnosisSafe contract.
type GnosisSafeRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) FilterRemovedOwner(opts *bind.FilterOpts) (*GnosisSafeRemovedOwnerIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeRemovedOwnerIterator{contract: _GnosisSafe.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *GnosisSafeRemovedOwner) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeRemovedOwner)
				if err := _GnosisSafe.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) ParseRemovedOwner(log types.Log) (*GnosisSafeRemovedOwner, error) {
	event := new(GnosisSafeRemovedOwner)
	if err := _GnosisSafe.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the GnosisSafe contract.
type GnosisSafeSafeReceivedIterator struct {
	Event *GnosisSafeSafeReceived // Event containing the contract specifics and raw log

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
func (it *GnosisSafeSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeSafeReceived)
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
		it.Event = new(GnosisSafeSafeReceived)
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
func (it *GnosisSafeSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeSafeReceived represents a SafeReceived event raised by the GnosisSafe contract.
type GnosisSafeSafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafe *GnosisSafeFilterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*GnosisSafeSafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeSafeReceivedIterator{contract: _GnosisSafe.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafe *GnosisSafeFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *GnosisSafeSafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeSafeReceived)
				if err := _GnosisSafe.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafe *GnosisSafeFilterer) ParseSafeReceived(log types.Log) (*GnosisSafeSafeReceived, error) {
	event := new(GnosisSafeSafeReceived)
	if err := _GnosisSafe.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeSafeSetupIterator is returned from FilterSafeSetup and is used to iterate over the raw logs and unpacked data for SafeSetup events raised by the GnosisSafe contract.
type GnosisSafeSafeSetupIterator struct {
	Event *GnosisSafeSafeSetup // Event containing the contract specifics and raw log

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
func (it *GnosisSafeSafeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeSafeSetup)
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
		it.Event = new(GnosisSafeSafeSetup)
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
func (it *GnosisSafeSafeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeSafeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeSafeSetup represents a SafeSetup event raised by the GnosisSafe contract.
type GnosisSafeSafeSetup struct {
	Initiator       common.Address
	Owners          []common.Address
	Threshold       *big.Int
	Initializer     common.Address
	FallbackHandler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSafeSetup is a free log retrieval operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafe *GnosisSafeFilterer) FilterSafeSetup(opts *bind.FilterOpts, initiator []common.Address) (*GnosisSafeSafeSetupIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeSafeSetupIterator{contract: _GnosisSafe.contract, event: "SafeSetup", logs: logs, sub: sub}, nil
}

// WatchSafeSetup is a free log subscription operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafe *GnosisSafeFilterer) WatchSafeSetup(opts *bind.WatchOpts, sink chan<- *GnosisSafeSafeSetup, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeSafeSetup)
				if err := _GnosisSafe.contract.UnpackLog(event, "SafeSetup", log); err != nil {
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

// ParseSafeSetup is a log parse operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafe *GnosisSafeFilterer) ParseSafeSetup(log types.Log) (*GnosisSafeSafeSetup, error) {
	event := new(GnosisSafeSafeSetup)
	if err := _GnosisSafe.contract.UnpackLog(event, "SafeSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeSignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the GnosisSafe contract.
type GnosisSafeSignMsgIterator struct {
	Event *GnosisSafeSignMsg // Event containing the contract specifics and raw log

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
func (it *GnosisSafeSignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeSignMsg)
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
		it.Event = new(GnosisSafeSignMsg)
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
func (it *GnosisSafeSignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeSignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeSignMsg represents a SignMsg event raised by the GnosisSafe contract.
type GnosisSafeSignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafe *GnosisSafeFilterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*GnosisSafeSignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeSignMsgIterator{contract: _GnosisSafe.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafe *GnosisSafeFilterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *GnosisSafeSignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeSignMsg)
				if err := _GnosisSafe.contract.UnpackLog(event, "SignMsg", log); err != nil {
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

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafe *GnosisSafeFilterer) ParseSignMsg(log types.Log) (*GnosisSafeSignMsg, error) {
	event := new(GnosisSafeSignMsg)
	if err := _GnosisSafe.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeMathMetaData contains all meta data concerning the GnosisSafeMath contract.
var GnosisSafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201c65d753eb16585c1ae57d7cc778efe0243851140783053b6ae2fb24096063a064736f6c634300080d0033",
}

// GnosisSafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use GnosisSafeMathMetaData.ABI instead.
var GnosisSafeMathABI = GnosisSafeMathMetaData.ABI

// GnosisSafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GnosisSafeMathMetaData.Bin instead.
var GnosisSafeMathBin = GnosisSafeMathMetaData.Bin

// DeployGnosisSafeMath deploys a new Ethereum contract, binding an instance of GnosisSafeMath to it.
func DeployGnosisSafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GnosisSafeMath, error) {
	parsed, err := GnosisSafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GnosisSafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GnosisSafeMath{GnosisSafeMathCaller: GnosisSafeMathCaller{contract: contract}, GnosisSafeMathTransactor: GnosisSafeMathTransactor{contract: contract}, GnosisSafeMathFilterer: GnosisSafeMathFilterer{contract: contract}}, nil
}

// GnosisSafeMath is an auto generated Go binding around an Ethereum contract.
type GnosisSafeMath struct {
	GnosisSafeMathCaller     // Read-only binding to the contract
	GnosisSafeMathTransactor // Write-only binding to the contract
	GnosisSafeMathFilterer   // Log filterer for contract events
}

// GnosisSafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeMathSession struct {
	Contract     *GnosisSafeMath   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnosisSafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeMathCallerSession struct {
	Contract *GnosisSafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GnosisSafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeMathTransactorSession struct {
	Contract     *GnosisSafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GnosisSafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeMathRaw struct {
	Contract *GnosisSafeMath // Generic contract binding to access the raw methods on
}

// GnosisSafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeMathCallerRaw struct {
	Contract *GnosisSafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeMathTransactorRaw struct {
	Contract *GnosisSafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafeMath creates a new instance of GnosisSafeMath, bound to a specific deployed contract.
func NewGnosisSafeMath(address common.Address, backend bind.ContractBackend) (*GnosisSafeMath, error) {
	contract, err := bindGnosisSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMath{GnosisSafeMathCaller: GnosisSafeMathCaller{contract: contract}, GnosisSafeMathTransactor: GnosisSafeMathTransactor{contract: contract}, GnosisSafeMathFilterer: GnosisSafeMathFilterer{contract: contract}}, nil
}

// NewGnosisSafeMathCaller creates a new read-only instance of GnosisSafeMath, bound to a specific deployed contract.
func NewGnosisSafeMathCaller(address common.Address, caller bind.ContractCaller) (*GnosisSafeMathCaller, error) {
	contract, err := bindGnosisSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMathCaller{contract: contract}, nil
}

// NewGnosisSafeMathTransactor creates a new write-only instance of GnosisSafeMath, bound to a specific deployed contract.
func NewGnosisSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeMathTransactor, error) {
	contract, err := bindGnosisSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMathTransactor{contract: contract}, nil
}

// NewGnosisSafeMathFilterer creates a new log filterer instance of GnosisSafeMath, bound to a specific deployed contract.
func NewGnosisSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeMathFilterer, error) {
	contract, err := bindGnosisSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMathFilterer{contract: contract}, nil
}

// bindGnosisSafeMath binds a generic wrapper to an already deployed contract.
func bindGnosisSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisSafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeMath *GnosisSafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeMath.Contract.GnosisSafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeMath *GnosisSafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeMath.Contract.GnosisSafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeMath *GnosisSafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeMath.Contract.GnosisSafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeMath *GnosisSafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeMath *GnosisSafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeMath *GnosisSafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeMath.Contract.contract.Transact(opts, method, params...)
}

// GuardMetaData contains all meta data concerning the Guard contract.
var GuardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"checkAfterExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"checkTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"93271368": "checkAfterExecution(bytes32,bool)",
		"75f0bb52": "checkTransaction(address,uint256,bytes,uint8,uint256,uint256,uint256,address,address,bytes,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// GuardABI is the input ABI used to generate the binding from.
// Deprecated: Use GuardMetaData.ABI instead.
var GuardABI = GuardMetaData.ABI

// Deprecated: Use GuardMetaData.Sigs instead.
// GuardFuncSigs maps the 4-byte function signature to its string representation.
var GuardFuncSigs = GuardMetaData.Sigs

// Guard is an auto generated Go binding around an Ethereum contract.
type Guard struct {
	GuardCaller     // Read-only binding to the contract
	GuardTransactor // Write-only binding to the contract
	GuardFilterer   // Log filterer for contract events
}

// GuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuardSession struct {
	Contract     *Guard            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuardCallerSession struct {
	Contract *GuardCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuardTransactorSession struct {
	Contract     *GuardTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuardRaw struct {
	Contract *Guard // Generic contract binding to access the raw methods on
}

// GuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuardCallerRaw struct {
	Contract *GuardCaller // Generic read-only contract binding to access the raw methods on
}

// GuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuardTransactorRaw struct {
	Contract *GuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuard creates a new instance of Guard, bound to a specific deployed contract.
func NewGuard(address common.Address, backend bind.ContractBackend) (*Guard, error) {
	contract, err := bindGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Guard{GuardCaller: GuardCaller{contract: contract}, GuardTransactor: GuardTransactor{contract: contract}, GuardFilterer: GuardFilterer{contract: contract}}, nil
}

// NewGuardCaller creates a new read-only instance of Guard, bound to a specific deployed contract.
func NewGuardCaller(address common.Address, caller bind.ContractCaller) (*GuardCaller, error) {
	contract, err := bindGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuardCaller{contract: contract}, nil
}

// NewGuardTransactor creates a new write-only instance of Guard, bound to a specific deployed contract.
func NewGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*GuardTransactor, error) {
	contract, err := bindGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuardTransactor{contract: contract}, nil
}

// NewGuardFilterer creates a new log filterer instance of Guard, bound to a specific deployed contract.
func NewGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*GuardFilterer, error) {
	contract, err := bindGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuardFilterer{contract: contract}, nil
}

// bindGuard binds a generic wrapper to an already deployed contract.
func bindGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Guard *GuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Guard.Contract.GuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Guard *GuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Guard.Contract.GuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Guard *GuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Guard.Contract.GuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Guard *GuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Guard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Guard *GuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Guard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Guard *GuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Guard.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Guard *GuardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Guard.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Guard *GuardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Guard.Contract.SupportsInterface(&_Guard.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Guard *GuardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Guard.Contract.SupportsInterface(&_Guard.CallOpts, interfaceId)
}

// CheckAfterExecution is a paid mutator transaction binding the contract method 0x93271368.
//
// Solidity: function checkAfterExecution(bytes32 txHash, bool success) returns()
func (_Guard *GuardTransactor) CheckAfterExecution(opts *bind.TransactOpts, txHash [32]byte, success bool) (*types.Transaction, error) {
	return _Guard.contract.Transact(opts, "checkAfterExecution", txHash, success)
}

// CheckAfterExecution is a paid mutator transaction binding the contract method 0x93271368.
//
// Solidity: function checkAfterExecution(bytes32 txHash, bool success) returns()
func (_Guard *GuardSession) CheckAfterExecution(txHash [32]byte, success bool) (*types.Transaction, error) {
	return _Guard.Contract.CheckAfterExecution(&_Guard.TransactOpts, txHash, success)
}

// CheckAfterExecution is a paid mutator transaction binding the contract method 0x93271368.
//
// Solidity: function checkAfterExecution(bytes32 txHash, bool success) returns()
func (_Guard *GuardTransactorSession) CheckAfterExecution(txHash [32]byte, success bool) (*types.Transaction, error) {
	return _Guard.Contract.CheckAfterExecution(&_Guard.TransactOpts, txHash, success)
}

// CheckTransaction is a paid mutator transaction binding the contract method 0x75f0bb52.
//
// Solidity: function checkTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, address msgSender) returns()
func (_Guard *GuardTransactor) CheckTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte, msgSender common.Address) (*types.Transaction, error) {
	return _Guard.contract.Transact(opts, "checkTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures, msgSender)
}

// CheckTransaction is a paid mutator transaction binding the contract method 0x75f0bb52.
//
// Solidity: function checkTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, address msgSender) returns()
func (_Guard *GuardSession) CheckTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte, msgSender common.Address) (*types.Transaction, error) {
	return _Guard.Contract.CheckTransaction(&_Guard.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures, msgSender)
}

// CheckTransaction is a paid mutator transaction binding the contract method 0x75f0bb52.
//
// Solidity: function checkTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, address msgSender) returns()
func (_Guard *GuardTransactorSession) CheckTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte, msgSender common.Address) (*types.Transaction, error) {
	return _Guard.Contract.CheckTransaction(&_Guard.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures, msgSender)
}

// GuardManagerMetaData contains all meta data concerning the GuardManager contract.
var GuardManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e19a9dd9": "setGuard(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610229806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e19a9dd914610030575b600080fd5b61004361003e3660046101a1565b610045565b005b61004d610168565b6001600160a01b03811615610104576040516301ffc9a760e01b815263736bd41d60e11b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa1580156100a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100cb91906101d1565b6101045760405162461bcd60e51b8152602060048201526005602482015264047533330360dc1b60448201526064015b60405180910390fd5b7f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c88181556040516001600160a01b03831681527f1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa29060200160405180910390a15050565b33301461019f5760405162461bcd60e51b8152602060048201526005602482015264475330333160d81b60448201526064016100fb565b565b6000602082840312156101b357600080fd5b81356001600160a01b03811681146101ca57600080fd5b9392505050565b6000602082840312156101e357600080fd5b815180151581146101ca57600080fdfea2646970667358221220553bcfbead9321885f8805df47be18a0eb0eebdf7b55c3996fc51dc13a554c7364736f6c634300080d0033",
}

// GuardManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use GuardManagerMetaData.ABI instead.
var GuardManagerABI = GuardManagerMetaData.ABI

// Deprecated: Use GuardManagerMetaData.Sigs instead.
// GuardManagerFuncSigs maps the 4-byte function signature to its string representation.
var GuardManagerFuncSigs = GuardManagerMetaData.Sigs

// GuardManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GuardManagerMetaData.Bin instead.
var GuardManagerBin = GuardManagerMetaData.Bin

// DeployGuardManager deploys a new Ethereum contract, binding an instance of GuardManager to it.
func DeployGuardManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GuardManager, error) {
	parsed, err := GuardManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GuardManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GuardManager{GuardManagerCaller: GuardManagerCaller{contract: contract}, GuardManagerTransactor: GuardManagerTransactor{contract: contract}, GuardManagerFilterer: GuardManagerFilterer{contract: contract}}, nil
}

// GuardManager is an auto generated Go binding around an Ethereum contract.
type GuardManager struct {
	GuardManagerCaller     // Read-only binding to the contract
	GuardManagerTransactor // Write-only binding to the contract
	GuardManagerFilterer   // Log filterer for contract events
}

// GuardManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuardManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuardManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuardManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuardManagerSession struct {
	Contract     *GuardManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GuardManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuardManagerCallerSession struct {
	Contract *GuardManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GuardManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuardManagerTransactorSession struct {
	Contract     *GuardManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GuardManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuardManagerRaw struct {
	Contract *GuardManager // Generic contract binding to access the raw methods on
}

// GuardManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuardManagerCallerRaw struct {
	Contract *GuardManagerCaller // Generic read-only contract binding to access the raw methods on
}

// GuardManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuardManagerTransactorRaw struct {
	Contract *GuardManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuardManager creates a new instance of GuardManager, bound to a specific deployed contract.
func NewGuardManager(address common.Address, backend bind.ContractBackend) (*GuardManager, error) {
	contract, err := bindGuardManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuardManager{GuardManagerCaller: GuardManagerCaller{contract: contract}, GuardManagerTransactor: GuardManagerTransactor{contract: contract}, GuardManagerFilterer: GuardManagerFilterer{contract: contract}}, nil
}

// NewGuardManagerCaller creates a new read-only instance of GuardManager, bound to a specific deployed contract.
func NewGuardManagerCaller(address common.Address, caller bind.ContractCaller) (*GuardManagerCaller, error) {
	contract, err := bindGuardManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuardManagerCaller{contract: contract}, nil
}

// NewGuardManagerTransactor creates a new write-only instance of GuardManager, bound to a specific deployed contract.
func NewGuardManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*GuardManagerTransactor, error) {
	contract, err := bindGuardManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuardManagerTransactor{contract: contract}, nil
}

// NewGuardManagerFilterer creates a new log filterer instance of GuardManager, bound to a specific deployed contract.
func NewGuardManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*GuardManagerFilterer, error) {
	contract, err := bindGuardManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuardManagerFilterer{contract: contract}, nil
}

// bindGuardManager binds a generic wrapper to an already deployed contract.
func bindGuardManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GuardManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardManager *GuardManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardManager.Contract.GuardManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardManager *GuardManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardManager.Contract.GuardManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardManager *GuardManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardManager.Contract.GuardManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardManager *GuardManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardManager *GuardManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardManager *GuardManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardManager.Contract.contract.Transact(opts, method, params...)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GuardManager *GuardManagerTransactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _GuardManager.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GuardManager *GuardManagerSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GuardManager.Contract.SetGuard(&_GuardManager.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GuardManager *GuardManagerTransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GuardManager.Contract.SetGuard(&_GuardManager.TransactOpts, guard)
}

// GuardManagerChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the GuardManager contract.
type GuardManagerChangedGuardIterator struct {
	Event *GuardManagerChangedGuard // Event containing the contract specifics and raw log

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
func (it *GuardManagerChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardManagerChangedGuard)
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
		it.Event = new(GuardManagerChangedGuard)
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
func (it *GuardManagerChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardManagerChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardManagerChangedGuard represents a ChangedGuard event raised by the GuardManager contract.
type GuardManagerChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GuardManager *GuardManagerFilterer) FilterChangedGuard(opts *bind.FilterOpts) (*GuardManagerChangedGuardIterator, error) {

	logs, sub, err := _GuardManager.contract.FilterLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return &GuardManagerChangedGuardIterator{contract: _GuardManager.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GuardManager *GuardManagerFilterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *GuardManagerChangedGuard) (event.Subscription, error) {

	logs, sub, err := _GuardManager.contract.WatchLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardManagerChangedGuard)
				if err := _GuardManager.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
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

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GuardManager *GuardManagerFilterer) ParseChangedGuard(log types.Log) (*GuardManagerChangedGuard, error) {
	event := new(GuardManagerChangedGuard)
	if err := _GuardManager.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// Deprecated: Use IERC165MetaData.Sigs instead.
// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = IERC165MetaData.Sigs

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// ISignatureValidatorMetaData contains all meta data concerning the ISignatureValidator contract.
var ISignatureValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"20c13b0b": "isValidSignature(bytes,bytes)",
	},
}

// ISignatureValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ISignatureValidatorMetaData.ABI instead.
var ISignatureValidatorABI = ISignatureValidatorMetaData.ABI

// Deprecated: Use ISignatureValidatorMetaData.Sigs instead.
// ISignatureValidatorFuncSigs maps the 4-byte function signature to its string representation.
var ISignatureValidatorFuncSigs = ISignatureValidatorMetaData.Sigs

// ISignatureValidator is an auto generated Go binding around an Ethereum contract.
type ISignatureValidator struct {
	ISignatureValidatorCaller     // Read-only binding to the contract
	ISignatureValidatorTransactor // Write-only binding to the contract
	ISignatureValidatorFilterer   // Log filterer for contract events
}

// ISignatureValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISignatureValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISignatureValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISignatureValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISignatureValidatorSession struct {
	Contract     *ISignatureValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISignatureValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISignatureValidatorCallerSession struct {
	Contract *ISignatureValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ISignatureValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISignatureValidatorTransactorSession struct {
	Contract     *ISignatureValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ISignatureValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISignatureValidatorRaw struct {
	Contract *ISignatureValidator // Generic contract binding to access the raw methods on
}

// ISignatureValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISignatureValidatorCallerRaw struct {
	Contract *ISignatureValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ISignatureValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISignatureValidatorTransactorRaw struct {
	Contract *ISignatureValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISignatureValidator creates a new instance of ISignatureValidator, bound to a specific deployed contract.
func NewISignatureValidator(address common.Address, backend bind.ContractBackend) (*ISignatureValidator, error) {
	contract, err := bindISignatureValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidator{ISignatureValidatorCaller: ISignatureValidatorCaller{contract: contract}, ISignatureValidatorTransactor: ISignatureValidatorTransactor{contract: contract}, ISignatureValidatorFilterer: ISignatureValidatorFilterer{contract: contract}}, nil
}

// NewISignatureValidatorCaller creates a new read-only instance of ISignatureValidator, bound to a specific deployed contract.
func NewISignatureValidatorCaller(address common.Address, caller bind.ContractCaller) (*ISignatureValidatorCaller, error) {
	contract, err := bindISignatureValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidatorCaller{contract: contract}, nil
}

// NewISignatureValidatorTransactor creates a new write-only instance of ISignatureValidator, bound to a specific deployed contract.
func NewISignatureValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ISignatureValidatorTransactor, error) {
	contract, err := bindISignatureValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidatorTransactor{contract: contract}, nil
}

// NewISignatureValidatorFilterer creates a new log filterer instance of ISignatureValidator, bound to a specific deployed contract.
func NewISignatureValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ISignatureValidatorFilterer, error) {
	contract, err := bindISignatureValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidatorFilterer{contract: contract}, nil
}

// bindISignatureValidator binds a generic wrapper to an already deployed contract.
func bindISignatureValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISignatureValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureValidator *ISignatureValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureValidator.Contract.ISignatureValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureValidator *ISignatureValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureValidator.Contract.ISignatureValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureValidator *ISignatureValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureValidator.Contract.ISignatureValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureValidator *ISignatureValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureValidator *ISignatureValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureValidator *ISignatureValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureValidator.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_ISignatureValidator *ISignatureValidatorCaller) IsValidSignature(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ISignatureValidator.contract.Call(opts, &out, "isValidSignature", _data, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_ISignatureValidator *ISignatureValidatorSession) IsValidSignature(_data []byte, _signature []byte) ([4]byte, error) {
	return _ISignatureValidator.Contract.IsValidSignature(&_ISignatureValidator.CallOpts, _data, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_ISignatureValidator *ISignatureValidatorCallerSession) IsValidSignature(_data []byte, _signature []byte) ([4]byte, error) {
	return _ISignatureValidator.Contract.IsValidSignature(&_ISignatureValidator.CallOpts, _data, _signature)
}

// ISignatureValidatorConstantsMetaData contains all meta data concerning the ISignatureValidatorConstants contract.
var ISignatureValidatorConstantsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212207b18b107bc3b571b1bb6ac0466b159323d4730b29efbc15ad2ebfcb32e3eda1164736f6c634300080d0033",
}

// ISignatureValidatorConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use ISignatureValidatorConstantsMetaData.ABI instead.
var ISignatureValidatorConstantsABI = ISignatureValidatorConstantsMetaData.ABI

// ISignatureValidatorConstantsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ISignatureValidatorConstantsMetaData.Bin instead.
var ISignatureValidatorConstantsBin = ISignatureValidatorConstantsMetaData.Bin

// DeployISignatureValidatorConstants deploys a new Ethereum contract, binding an instance of ISignatureValidatorConstants to it.
func DeployISignatureValidatorConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ISignatureValidatorConstants, error) {
	parsed, err := ISignatureValidatorConstantsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ISignatureValidatorConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ISignatureValidatorConstants{ISignatureValidatorConstantsCaller: ISignatureValidatorConstantsCaller{contract: contract}, ISignatureValidatorConstantsTransactor: ISignatureValidatorConstantsTransactor{contract: contract}, ISignatureValidatorConstantsFilterer: ISignatureValidatorConstantsFilterer{contract: contract}}, nil
}

// ISignatureValidatorConstants is an auto generated Go binding around an Ethereum contract.
type ISignatureValidatorConstants struct {
	ISignatureValidatorConstantsCaller     // Read-only binding to the contract
	ISignatureValidatorConstantsTransactor // Write-only binding to the contract
	ISignatureValidatorConstantsFilterer   // Log filterer for contract events
}

// ISignatureValidatorConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISignatureValidatorConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureValidatorConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISignatureValidatorConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureValidatorConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISignatureValidatorConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureValidatorConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISignatureValidatorConstantsSession struct {
	Contract     *ISignatureValidatorConstants // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ISignatureValidatorConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISignatureValidatorConstantsCallerSession struct {
	Contract *ISignatureValidatorConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ISignatureValidatorConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISignatureValidatorConstantsTransactorSession struct {
	Contract     *ISignatureValidatorConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ISignatureValidatorConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISignatureValidatorConstantsRaw struct {
	Contract *ISignatureValidatorConstants // Generic contract binding to access the raw methods on
}

// ISignatureValidatorConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISignatureValidatorConstantsCallerRaw struct {
	Contract *ISignatureValidatorConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// ISignatureValidatorConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISignatureValidatorConstantsTransactorRaw struct {
	Contract *ISignatureValidatorConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISignatureValidatorConstants creates a new instance of ISignatureValidatorConstants, bound to a specific deployed contract.
func NewISignatureValidatorConstants(address common.Address, backend bind.ContractBackend) (*ISignatureValidatorConstants, error) {
	contract, err := bindISignatureValidatorConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidatorConstants{ISignatureValidatorConstantsCaller: ISignatureValidatorConstantsCaller{contract: contract}, ISignatureValidatorConstantsTransactor: ISignatureValidatorConstantsTransactor{contract: contract}, ISignatureValidatorConstantsFilterer: ISignatureValidatorConstantsFilterer{contract: contract}}, nil
}

// NewISignatureValidatorConstantsCaller creates a new read-only instance of ISignatureValidatorConstants, bound to a specific deployed contract.
func NewISignatureValidatorConstantsCaller(address common.Address, caller bind.ContractCaller) (*ISignatureValidatorConstantsCaller, error) {
	contract, err := bindISignatureValidatorConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidatorConstantsCaller{contract: contract}, nil
}

// NewISignatureValidatorConstantsTransactor creates a new write-only instance of ISignatureValidatorConstants, bound to a specific deployed contract.
func NewISignatureValidatorConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISignatureValidatorConstantsTransactor, error) {
	contract, err := bindISignatureValidatorConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidatorConstantsTransactor{contract: contract}, nil
}

// NewISignatureValidatorConstantsFilterer creates a new log filterer instance of ISignatureValidatorConstants, bound to a specific deployed contract.
func NewISignatureValidatorConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISignatureValidatorConstantsFilterer, error) {
	contract, err := bindISignatureValidatorConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISignatureValidatorConstantsFilterer{contract: contract}, nil
}

// bindISignatureValidatorConstants binds a generic wrapper to an already deployed contract.
func bindISignatureValidatorConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISignatureValidatorConstantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureValidatorConstants *ISignatureValidatorConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureValidatorConstants.Contract.ISignatureValidatorConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureValidatorConstants *ISignatureValidatorConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureValidatorConstants.Contract.ISignatureValidatorConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureValidatorConstants *ISignatureValidatorConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureValidatorConstants.Contract.ISignatureValidatorConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureValidatorConstants *ISignatureValidatorConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureValidatorConstants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureValidatorConstants *ISignatureValidatorConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureValidatorConstants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureValidatorConstants *ISignatureValidatorConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureValidatorConstants.Contract.contract.Transact(opts, method, params...)
}

// ModuleManagerMetaData contains all meta data concerning the ModuleManager contract.
var ModuleManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e009cfde": "disableModule(address,address)",
		"610b5925": "enableModule(address)",
		"468721a7": "execTransactionFromModule(address,uint256,bytes,uint8)",
		"5229073f": "execTransactionFromModuleReturnData(address,uint256,bytes,uint8)",
		"cc2f8452": "getModulesPaginated(address,uint256)",
		"2d9ad53d": "isModuleEnabled(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061091f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632d9ad53d14610067578063468721a71461008f5780635229073f146100a2578063610b5925146100c3578063cc2f8452146100d8578063e009cfde146100f9575b600080fd5b61007a61007536600461065a565b61010c565b60405190151581526020015b60405180910390f35b61007a61009d3660046106a1565b610147565b6100b56100b03660046106a1565b610223565b60405161008692919061077d565b6100d66100d136600461065a565b610259565b005b6100eb6100e63660046107dc565b610399565b604051610086929190610806565b6100d6610107366004610863565b610492565b600060016001600160a01b0383161480159061014157506001600160a01b038281166000908152602081905260409020541615155b92915050565b6000336001148015906101715750336000908152602081905260409020546001600160a01b031615155b6101aa5760405162461bcd60e51b815260206004820152600560248201526411d4cc4c0d60da1b60448201526064015b60405180910390fd5b6101b7858585855a6105be565b905080156101ef5760405133907f6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb890600090a261021b565b60405133907facd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd37590600090a25b949350505050565b6000606061023386868686610147565b915060405160203d0181016040523d81523d6000602083013e8091505094509492505050565b610261610605565b6001600160a01b0381161580159061028357506001600160a01b038116600114155b6102b75760405162461bcd60e51b8152602060048201526005602482015264475331303160d81b60448201526064016101a1565b6001600160a01b0381811660009081526020819052604090205416156103075760405162461bcd60e51b815260206004820152600560248201526423a998981960d91b60448201526064016101a1565b600060208181527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b0385811680865260408087208054939094166001600160a01b03199384161790935560019095528254168417909155519182527fecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440910160405180910390a150565b606060008267ffffffffffffffff8111156103b6576103b661067c565b6040519080825280602002602001820160405280156103df578160200160208202803683370190505b506001600160a01b0380861660009081526020819052604081205492945091165b6001600160a01b0381161580159061042257506001600160a01b038116600114155b801561042d57508482105b15610484578084838151811061044557610445610896565b6001600160a01b03928316602091820292909201810191909152918116600090815291829052604090912054168161047c816108ac565b925050610400565b908352919491935090915050565b61049a610605565b6001600160a01b038116158015906104bc57506001600160a01b038116600114155b6104f05760405162461bcd60e51b8152602060048201526005602482015264475331303160d81b60448201526064016101a1565b6001600160a01b038281166000908152602081905260409020548116908216146105445760405162461bcd60e51b8152602060048201526005602482015264475331303360d81b60448201526064016101a1565b6001600160a01b03818116600081815260208181526040808320805488871685528285208054919097166001600160a01b03199182161790965592849052825490941690915591519081527faab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276910160405180910390a15050565b600060018360018111156105d4576105d46108d3565b036105ec576000808551602087018986f490506105fc565b600080855160208701888a87f190505b95945050505050565b33301461063c5760405162461bcd60e51b8152602060048201526005602482015264475330333160d81b60448201526064016101a1565b565b80356001600160a01b038116811461065557600080fd5b919050565b60006020828403121561066c57600080fd5b6106758261063e565b9392505050565b634e487b7160e01b600052604160045260246000fd5b80356002811061065557600080fd5b600080600080608085870312156106b757600080fd5b6106c08561063e565b935060208501359250604085013567ffffffffffffffff808211156106e457600080fd5b818701915087601f8301126106f857600080fd5b81358181111561070a5761070a61067c565b604051601f8201601f19908116603f011681019083821181831017156107325761073261067c565b816040528281528a602084870101111561074b57600080fd5b82602086016020830137600060208483010152809650505050505061077260608601610692565b905092959194509250565b821515815260006020604081840152835180604085015260005b818110156107b357858101830151858201606001528201610797565b818111156107c5576000606083870101525b50601f01601f191692909201606001949350505050565b600080604083850312156107ef57600080fd5b6107f88361063e565b946020939093013593505050565b604080825283519082018190526000906020906060840190828701845b828110156108485781516001600160a01b031684529284019290840190600101610823565b5050506001600160a01b039490941692019190915250919050565b6000806040838503121561087657600080fd5b61087f8361063e565b915061088d6020840161063e565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b6000600182016108cc57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220c6102addab68c59b9ff079ae4aa86a247c6ecf6011b01ca1d1565a12baa4922f64736f6c634300080d0033",
}

// ModuleManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ModuleManagerMetaData.ABI instead.
var ModuleManagerABI = ModuleManagerMetaData.ABI

// Deprecated: Use ModuleManagerMetaData.Sigs instead.
// ModuleManagerFuncSigs maps the 4-byte function signature to its string representation.
var ModuleManagerFuncSigs = ModuleManagerMetaData.Sigs

// ModuleManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModuleManagerMetaData.Bin instead.
var ModuleManagerBin = ModuleManagerMetaData.Bin

// DeployModuleManager deploys a new Ethereum contract, binding an instance of ModuleManager to it.
func DeployModuleManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModuleManager, error) {
	parsed, err := ModuleManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModuleManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModuleManager{ModuleManagerCaller: ModuleManagerCaller{contract: contract}, ModuleManagerTransactor: ModuleManagerTransactor{contract: contract}, ModuleManagerFilterer: ModuleManagerFilterer{contract: contract}}, nil
}

// ModuleManager is an auto generated Go binding around an Ethereum contract.
type ModuleManager struct {
	ModuleManagerCaller     // Read-only binding to the contract
	ModuleManagerTransactor // Write-only binding to the contract
	ModuleManagerFilterer   // Log filterer for contract events
}

// ModuleManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModuleManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModuleManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModuleManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModuleManagerSession struct {
	Contract     *ModuleManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModuleManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModuleManagerCallerSession struct {
	Contract *ModuleManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ModuleManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModuleManagerTransactorSession struct {
	Contract     *ModuleManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ModuleManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModuleManagerRaw struct {
	Contract *ModuleManager // Generic contract binding to access the raw methods on
}

// ModuleManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModuleManagerCallerRaw struct {
	Contract *ModuleManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ModuleManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModuleManagerTransactorRaw struct {
	Contract *ModuleManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModuleManager creates a new instance of ModuleManager, bound to a specific deployed contract.
func NewModuleManager(address common.Address, backend bind.ContractBackend) (*ModuleManager, error) {
	contract, err := bindModuleManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModuleManager{ModuleManagerCaller: ModuleManagerCaller{contract: contract}, ModuleManagerTransactor: ModuleManagerTransactor{contract: contract}, ModuleManagerFilterer: ModuleManagerFilterer{contract: contract}}, nil
}

// NewModuleManagerCaller creates a new read-only instance of ModuleManager, bound to a specific deployed contract.
func NewModuleManagerCaller(address common.Address, caller bind.ContractCaller) (*ModuleManagerCaller, error) {
	contract, err := bindModuleManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModuleManagerCaller{contract: contract}, nil
}

// NewModuleManagerTransactor creates a new write-only instance of ModuleManager, bound to a specific deployed contract.
func NewModuleManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ModuleManagerTransactor, error) {
	contract, err := bindModuleManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModuleManagerTransactor{contract: contract}, nil
}

// NewModuleManagerFilterer creates a new log filterer instance of ModuleManager, bound to a specific deployed contract.
func NewModuleManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ModuleManagerFilterer, error) {
	contract, err := bindModuleManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModuleManagerFilterer{contract: contract}, nil
}

// bindModuleManager binds a generic wrapper to an already deployed contract.
func bindModuleManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModuleManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModuleManager *ModuleManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModuleManager.Contract.ModuleManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModuleManager *ModuleManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModuleManager.Contract.ModuleManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModuleManager *ModuleManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModuleManager.Contract.ModuleManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModuleManager *ModuleManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModuleManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModuleManager *ModuleManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModuleManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModuleManager *ModuleManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModuleManager.Contract.contract.Transact(opts, method, params...)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_ModuleManager *ModuleManagerCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _ModuleManager.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_ModuleManager *ModuleManagerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _ModuleManager.Contract.GetModulesPaginated(&_ModuleManager.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_ModuleManager *ModuleManagerCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _ModuleManager.Contract.GetModulesPaginated(&_ModuleManager.CallOpts, start, pageSize)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_ModuleManager *ModuleManagerCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _ModuleManager.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_ModuleManager *ModuleManagerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _ModuleManager.Contract.IsModuleEnabled(&_ModuleManager.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_ModuleManager *ModuleManagerCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _ModuleManager.Contract.IsModuleEnabled(&_ModuleManager.CallOpts, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_ModuleManager *ModuleManagerTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _ModuleManager.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_ModuleManager *ModuleManagerSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _ModuleManager.Contract.DisableModule(&_ModuleManager.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_ModuleManager *ModuleManagerTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _ModuleManager.Contract.DisableModule(&_ModuleManager.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_ModuleManager *ModuleManagerTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _ModuleManager.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_ModuleManager *ModuleManagerSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _ModuleManager.Contract.EnableModule(&_ModuleManager.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_ModuleManager *ModuleManagerTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _ModuleManager.Contract.EnableModule(&_ModuleManager.TransactOpts, module)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_ModuleManager *ModuleManagerTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ModuleManager.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_ModuleManager *ModuleManagerSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ModuleManager.Contract.ExecTransactionFromModule(&_ModuleManager.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_ModuleManager *ModuleManagerTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ModuleManager.Contract.ExecTransactionFromModule(&_ModuleManager.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_ModuleManager *ModuleManagerTransactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ModuleManager.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_ModuleManager *ModuleManagerSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ModuleManager.Contract.ExecTransactionFromModuleReturnData(&_ModuleManager.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_ModuleManager *ModuleManagerTransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ModuleManager.Contract.ExecTransactionFromModuleReturnData(&_ModuleManager.TransactOpts, to, value, data, operation)
}

// ModuleManagerDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the ModuleManager contract.
type ModuleManagerDisabledModuleIterator struct {
	Event *ModuleManagerDisabledModule // Event containing the contract specifics and raw log

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
func (it *ModuleManagerDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModuleManagerDisabledModule)
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
		it.Event = new(ModuleManagerDisabledModule)
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
func (it *ModuleManagerDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModuleManagerDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModuleManagerDisabledModule represents a DisabledModule event raised by the ModuleManager contract.
type ModuleManagerDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_ModuleManager *ModuleManagerFilterer) FilterDisabledModule(opts *bind.FilterOpts) (*ModuleManagerDisabledModuleIterator, error) {

	logs, sub, err := _ModuleManager.contract.FilterLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return &ModuleManagerDisabledModuleIterator{contract: _ModuleManager.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_ModuleManager *ModuleManagerFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *ModuleManagerDisabledModule) (event.Subscription, error) {

	logs, sub, err := _ModuleManager.contract.WatchLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModuleManagerDisabledModule)
				if err := _ModuleManager.contract.UnpackLog(event, "DisabledModule", log); err != nil {
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

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_ModuleManager *ModuleManagerFilterer) ParseDisabledModule(log types.Log) (*ModuleManagerDisabledModule, error) {
	event := new(ModuleManagerDisabledModule)
	if err := _ModuleManager.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModuleManagerEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the ModuleManager contract.
type ModuleManagerEnabledModuleIterator struct {
	Event *ModuleManagerEnabledModule // Event containing the contract specifics and raw log

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
func (it *ModuleManagerEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModuleManagerEnabledModule)
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
		it.Event = new(ModuleManagerEnabledModule)
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
func (it *ModuleManagerEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModuleManagerEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModuleManagerEnabledModule represents a EnabledModule event raised by the ModuleManager contract.
type ModuleManagerEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_ModuleManager *ModuleManagerFilterer) FilterEnabledModule(opts *bind.FilterOpts) (*ModuleManagerEnabledModuleIterator, error) {

	logs, sub, err := _ModuleManager.contract.FilterLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return &ModuleManagerEnabledModuleIterator{contract: _ModuleManager.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_ModuleManager *ModuleManagerFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *ModuleManagerEnabledModule) (event.Subscription, error) {

	logs, sub, err := _ModuleManager.contract.WatchLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModuleManagerEnabledModule)
				if err := _ModuleManager.contract.UnpackLog(event, "EnabledModule", log); err != nil {
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

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_ModuleManager *ModuleManagerFilterer) ParseEnabledModule(log types.Log) (*ModuleManagerEnabledModule, error) {
	event := new(ModuleManagerEnabledModule)
	if err := _ModuleManager.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModuleManagerExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the ModuleManager contract.
type ModuleManagerExecutionFromModuleFailureIterator struct {
	Event *ModuleManagerExecutionFromModuleFailure // Event containing the contract specifics and raw log

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
func (it *ModuleManagerExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModuleManagerExecutionFromModuleFailure)
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
		it.Event = new(ModuleManagerExecutionFromModuleFailure)
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
func (it *ModuleManagerExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModuleManagerExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModuleManagerExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the ModuleManager contract.
type ModuleManagerExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_ModuleManager *ModuleManagerFilterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*ModuleManagerExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ModuleManager.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &ModuleManagerExecutionFromModuleFailureIterator{contract: _ModuleManager.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_ModuleManager *ModuleManagerFilterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *ModuleManagerExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ModuleManager.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModuleManagerExecutionFromModuleFailure)
				if err := _ModuleManager.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
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

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_ModuleManager *ModuleManagerFilterer) ParseExecutionFromModuleFailure(log types.Log) (*ModuleManagerExecutionFromModuleFailure, error) {
	event := new(ModuleManagerExecutionFromModuleFailure)
	if err := _ModuleManager.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModuleManagerExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the ModuleManager contract.
type ModuleManagerExecutionFromModuleSuccessIterator struct {
	Event *ModuleManagerExecutionFromModuleSuccess // Event containing the contract specifics and raw log

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
func (it *ModuleManagerExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModuleManagerExecutionFromModuleSuccess)
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
		it.Event = new(ModuleManagerExecutionFromModuleSuccess)
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
func (it *ModuleManagerExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModuleManagerExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModuleManagerExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the ModuleManager contract.
type ModuleManagerExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_ModuleManager *ModuleManagerFilterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*ModuleManagerExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ModuleManager.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &ModuleManagerExecutionFromModuleSuccessIterator{contract: _ModuleManager.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_ModuleManager *ModuleManagerFilterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *ModuleManagerExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ModuleManager.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModuleManagerExecutionFromModuleSuccess)
				if err := _ModuleManager.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
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

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_ModuleManager *ModuleManagerFilterer) ParseExecutionFromModuleSuccess(log types.Log) (*ModuleManagerExecutionFromModuleSuccess, error) {
	event := new(ModuleManagerExecutionFromModuleSuccess)
	if err := _ModuleManager.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerManagerMetaData contains all meta data concerning the OwnerManager contract.
var OwnerManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d582f13": "addOwnerWithThreshold(address,uint256)",
		"694e80c3": "changeThreshold(uint256)",
		"a0e67e2b": "getOwners()",
		"e75235b8": "getThreshold()",
		"2f54bf6e": "isOwner(address)",
		"f8dc5dd9": "removeOwner(address,address,uint256)",
		"e318b52b": "swapOwner(address,address,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610a6b806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063a0e67e2b1161005b578063a0e67e2b146100d2578063e318b52b146100e7578063e75235b8146100fa578063f8dc5dd91461010b57600080fd5b80630d582f13146100825780632f54bf6e14610097578063694e80c3146100bf575b600080fd5b61009561009036600461085c565b61011e565b005b6100aa6100a5366004610886565b610297565b60405190151581526020015b60405180910390f35b6100956100cd3660046108a8565b6102d2565b6100da610388565b6040516100b691906108c1565b6100956100f536600461090e565b610478565b6002546040519081526020016100b6565b610095610119366004610951565b61067d565b610126610807565b6001600160a01b0382161580159061014857506001600160a01b038216600114155b801561015d57506001600160a01b0382163014155b6101825760405162461bcd60e51b81526004016101799061098d565b60405180910390fd5b6001600160a01b0382811660009081526020819052604090205416156101d25760405162461bcd60e51b815260206004820152600560248201526411d4cc8c0d60da1b6044820152606401610179565b600060208190527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d80546001600160a01b03858116808552604085208054929093166001600160a01b03199283161790925560018085528354909116909117909155805491610240836109c2565b90915550506040516001600160a01b03831681527f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea269060200160405180910390a1806002541461029357610293816102d2565b5050565b60006001600160a01b0382166001148015906102cc57506001600160a01b038281166000908152602081905260409020541615155b92915050565b6102da610807565b6001548111156103145760405162461bcd60e51b8152602060048201526005602482015264475332303160d81b6044820152606401610179565b600181101561034d5760405162461bcd60e51b815260206004820152600560248201526423a999181960d91b6044820152606401610179565b60028190556040518181527f610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c939060200160405180910390a150565b6060600060015467ffffffffffffffff8111156103a7576103a76109db565b6040519080825280602002602001820160405280156103d0578160200160208202803683370190505b506001600090815260208190527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d54919250906001600160a01b03165b6001600160a01b0381166001146104705780838381518110610431576104316109f1565b6001600160a01b039283166020918202929092018101919091529181166000908152918290526040909120541681610468816109c2565b92505061040d565b509092915050565b610480610807565b6001600160a01b038116158015906104a257506001600160a01b038116600114155b80156104b757506001600160a01b0381163014155b6104d35760405162461bcd60e51b81526004016101799061098d565b6001600160a01b0381811660009081526020819052604090205416156105235760405162461bcd60e51b815260206004820152600560248201526411d4cc8c0d60da1b6044820152606401610179565b6001600160a01b0382161580159061054557506001600160a01b038216600114155b6105615760405162461bcd60e51b81526004016101799061098d565b6001600160a01b038381166000908152602081905260409020548116908316146105b55760405162461bcd60e51b8152602060048201526005602482015264475332303560d81b6044820152606401610179565b6001600160a01b03828116600081815260208181526040808320805487871680865283862080549289166001600160a01b0319938416179055968a1685528285208054821690971790965592849052825490941690915591519081527ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf910160405180910390a16040516001600160a01b03821681527f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea269060200160405180910390a1505050565b610685610807565b80600180546106949190610a07565b10156106ca5760405162461bcd60e51b8152602060048201526005602482015264475332303160d81b6044820152606401610179565b6001600160a01b038216158015906106ec57506001600160a01b038216600114155b6107085760405162461bcd60e51b81526004016101799061098d565b6001600160a01b0383811660009081526020819052604090205481169083161461075c5760405162461bcd60e51b8152602060048201526005602482015264475332303560d81b6044820152606401610179565b6001600160a01b03828116600081815260208190526040808220805488861684529183208054929095166001600160a01b031992831617909455918152825490911690915560018054916107af83610a1e565b90915550506040516001600160a01b03831681527ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf9060200160405180910390a1806002541461080257610802816102d2565b505050565b33301461083e5760405162461bcd60e51b8152602060048201526005602482015264475330333160d81b6044820152606401610179565b565b80356001600160a01b038116811461085757600080fd5b919050565b6000806040838503121561086f57600080fd5b61087883610840565b946020939093013593505050565b60006020828403121561089857600080fd5b6108a182610840565b9392505050565b6000602082840312156108ba57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156109025783516001600160a01b0316835292840192918401916001016108dd565b50909695505050505050565b60008060006060848603121561092357600080fd5b61092c84610840565b925061093a60208501610840565b915061094860408501610840565b90509250925092565b60008060006060848603121561096657600080fd5b61096f84610840565b925061097d60208501610840565b9150604084013590509250925092565b602080825260059082015264475332303360d81b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000600182016109d4576109d46109ac565b5060010190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600082821015610a1957610a196109ac565b500390565b600081610a2d57610a2d6109ac565b50600019019056fea264697066735822122064f785f2fdb69e859916fdacd85838283bf847e87983dc083bdcb996ade7791664736f6c634300080d0033",
}

// OwnerManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnerManagerMetaData.ABI instead.
var OwnerManagerABI = OwnerManagerMetaData.ABI

// Deprecated: Use OwnerManagerMetaData.Sigs instead.
// OwnerManagerFuncSigs maps the 4-byte function signature to its string representation.
var OwnerManagerFuncSigs = OwnerManagerMetaData.Sigs

// OwnerManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnerManagerMetaData.Bin instead.
var OwnerManagerBin = OwnerManagerMetaData.Bin

// DeployOwnerManager deploys a new Ethereum contract, binding an instance of OwnerManager to it.
func DeployOwnerManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OwnerManager, error) {
	parsed, err := OwnerManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnerManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnerManager{OwnerManagerCaller: OwnerManagerCaller{contract: contract}, OwnerManagerTransactor: OwnerManagerTransactor{contract: contract}, OwnerManagerFilterer: OwnerManagerFilterer{contract: contract}}, nil
}

// OwnerManager is an auto generated Go binding around an Ethereum contract.
type OwnerManager struct {
	OwnerManagerCaller     // Read-only binding to the contract
	OwnerManagerTransactor // Write-only binding to the contract
	OwnerManagerFilterer   // Log filterer for contract events
}

// OwnerManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerManagerSession struct {
	Contract     *OwnerManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerManagerCallerSession struct {
	Contract *OwnerManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OwnerManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerManagerTransactorSession struct {
	Contract     *OwnerManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OwnerManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerManagerRaw struct {
	Contract *OwnerManager // Generic contract binding to access the raw methods on
}

// OwnerManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerManagerCallerRaw struct {
	Contract *OwnerManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerManagerTransactorRaw struct {
	Contract *OwnerManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnerManager creates a new instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManager(address common.Address, backend bind.ContractBackend) (*OwnerManager, error) {
	contract, err := bindOwnerManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnerManager{OwnerManagerCaller: OwnerManagerCaller{contract: contract}, OwnerManagerTransactor: OwnerManagerTransactor{contract: contract}, OwnerManagerFilterer: OwnerManagerFilterer{contract: contract}}, nil
}

// NewOwnerManagerCaller creates a new read-only instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManagerCaller(address common.Address, caller bind.ContractCaller) (*OwnerManagerCaller, error) {
	contract, err := bindOwnerManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerCaller{contract: contract}, nil
}

// NewOwnerManagerTransactor creates a new write-only instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerManagerTransactor, error) {
	contract, err := bindOwnerManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerTransactor{contract: contract}, nil
}

// NewOwnerManagerFilterer creates a new log filterer instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerManagerFilterer, error) {
	contract, err := bindOwnerManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerFilterer{contract: contract}, nil
}

// bindOwnerManager binds a generic wrapper to an already deployed contract.
func bindOwnerManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerManager *OwnerManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerManager.Contract.OwnerManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerManager *OwnerManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerManager.Contract.OwnerManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerManager *OwnerManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerManager.Contract.OwnerManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerManager *OwnerManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerManager *OwnerManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerManager *OwnerManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerManager.Contract.contract.Transact(opts, method, params...)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_OwnerManager *OwnerManagerCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OwnerManager.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_OwnerManager *OwnerManagerSession) GetOwners() ([]common.Address, error) {
	return _OwnerManager.Contract.GetOwners(&_OwnerManager.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_OwnerManager *OwnerManagerCallerSession) GetOwners() ([]common.Address, error) {
	return _OwnerManager.Contract.GetOwners(&_OwnerManager.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_OwnerManager *OwnerManagerCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwnerManager.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_OwnerManager *OwnerManagerSession) GetThreshold() (*big.Int, error) {
	return _OwnerManager.Contract.GetThreshold(&_OwnerManager.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_OwnerManager *OwnerManagerCallerSession) GetThreshold() (*big.Int, error) {
	return _OwnerManager.Contract.GetThreshold(&_OwnerManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_OwnerManager *OwnerManagerCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _OwnerManager.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_OwnerManager *OwnerManagerSession) IsOwner(owner common.Address) (bool, error) {
	return _OwnerManager.Contract.IsOwner(&_OwnerManager.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_OwnerManager *OwnerManagerCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _OwnerManager.Contract.IsOwner(&_OwnerManager.CallOpts, owner)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.Contract.AddOwnerWithThreshold(&_OwnerManager.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.Contract.AddOwnerWithThreshold(&_OwnerManager.TransactOpts, owner, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.Contract.ChangeThreshold(&_OwnerManager.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.Contract.ChangeThreshold(&_OwnerManager.TransactOpts, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.Contract.RemoveOwner(&_OwnerManager.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_OwnerManager *OwnerManagerTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _OwnerManager.Contract.RemoveOwner(&_OwnerManager.TransactOpts, prevOwner, owner, _threshold)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_OwnerManager *OwnerManagerTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _OwnerManager.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_OwnerManager *OwnerManagerSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _OwnerManager.Contract.SwapOwner(&_OwnerManager.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_OwnerManager *OwnerManagerTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _OwnerManager.Contract.SwapOwner(&_OwnerManager.TransactOpts, prevOwner, oldOwner, newOwner)
}

// OwnerManagerAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the OwnerManager contract.
type OwnerManagerAddedOwnerIterator struct {
	Event *OwnerManagerAddedOwner // Event containing the contract specifics and raw log

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
func (it *OwnerManagerAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerManagerAddedOwner)
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
		it.Event = new(OwnerManagerAddedOwner)
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
func (it *OwnerManagerAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerManagerAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerManagerAddedOwner represents a AddedOwner event raised by the OwnerManager contract.
type OwnerManagerAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_OwnerManager *OwnerManagerFilterer) FilterAddedOwner(opts *bind.FilterOpts) (*OwnerManagerAddedOwnerIterator, error) {

	logs, sub, err := _OwnerManager.contract.FilterLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return &OwnerManagerAddedOwnerIterator{contract: _OwnerManager.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_OwnerManager *OwnerManagerFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *OwnerManagerAddedOwner) (event.Subscription, error) {

	logs, sub, err := _OwnerManager.contract.WatchLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerManagerAddedOwner)
				if err := _OwnerManager.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_OwnerManager *OwnerManagerFilterer) ParseAddedOwner(log types.Log) (*OwnerManagerAddedOwner, error) {
	event := new(OwnerManagerAddedOwner)
	if err := _OwnerManager.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerManagerChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the OwnerManager contract.
type OwnerManagerChangedThresholdIterator struct {
	Event *OwnerManagerChangedThreshold // Event containing the contract specifics and raw log

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
func (it *OwnerManagerChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerManagerChangedThreshold)
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
		it.Event = new(OwnerManagerChangedThreshold)
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
func (it *OwnerManagerChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerManagerChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerManagerChangedThreshold represents a ChangedThreshold event raised by the OwnerManager contract.
type OwnerManagerChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_OwnerManager *OwnerManagerFilterer) FilterChangedThreshold(opts *bind.FilterOpts) (*OwnerManagerChangedThresholdIterator, error) {

	logs, sub, err := _OwnerManager.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &OwnerManagerChangedThresholdIterator{contract: _OwnerManager.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_OwnerManager *OwnerManagerFilterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *OwnerManagerChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _OwnerManager.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerManagerChangedThreshold)
				if err := _OwnerManager.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
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

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_OwnerManager *OwnerManagerFilterer) ParseChangedThreshold(log types.Log) (*OwnerManagerChangedThreshold, error) {
	event := new(OwnerManagerChangedThreshold)
	if err := _OwnerManager.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerManagerRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the OwnerManager contract.
type OwnerManagerRemovedOwnerIterator struct {
	Event *OwnerManagerRemovedOwner // Event containing the contract specifics and raw log

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
func (it *OwnerManagerRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerManagerRemovedOwner)
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
		it.Event = new(OwnerManagerRemovedOwner)
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
func (it *OwnerManagerRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerManagerRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerManagerRemovedOwner represents a RemovedOwner event raised by the OwnerManager contract.
type OwnerManagerRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_OwnerManager *OwnerManagerFilterer) FilterRemovedOwner(opts *bind.FilterOpts) (*OwnerManagerRemovedOwnerIterator, error) {

	logs, sub, err := _OwnerManager.contract.FilterLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return &OwnerManagerRemovedOwnerIterator{contract: _OwnerManager.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_OwnerManager *OwnerManagerFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *OwnerManagerRemovedOwner) (event.Subscription, error) {

	logs, sub, err := _OwnerManager.contract.WatchLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerManagerRemovedOwner)
				if err := _OwnerManager.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_OwnerManager *OwnerManagerFilterer) ParseRemovedOwner(log types.Log) (*OwnerManagerRemovedOwner, error) {
	event := new(OwnerManagerRemovedOwner)
	if err := _OwnerManager.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SecuredTokenTransferMetaData contains all meta data concerning the SecuredTokenTransfer contract.
var SecuredTokenTransferMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220e89a90711822198798c233f0d9ec72d23bf6e87c9e3d9fc62b57d4d78735850c64736f6c634300080d0033",
}

// SecuredTokenTransferABI is the input ABI used to generate the binding from.
// Deprecated: Use SecuredTokenTransferMetaData.ABI instead.
var SecuredTokenTransferABI = SecuredTokenTransferMetaData.ABI

// SecuredTokenTransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SecuredTokenTransferMetaData.Bin instead.
var SecuredTokenTransferBin = SecuredTokenTransferMetaData.Bin

// DeploySecuredTokenTransfer deploys a new Ethereum contract, binding an instance of SecuredTokenTransfer to it.
func DeploySecuredTokenTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SecuredTokenTransfer, error) {
	parsed, err := SecuredTokenTransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SecuredTokenTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecuredTokenTransfer{SecuredTokenTransferCaller: SecuredTokenTransferCaller{contract: contract}, SecuredTokenTransferTransactor: SecuredTokenTransferTransactor{contract: contract}, SecuredTokenTransferFilterer: SecuredTokenTransferFilterer{contract: contract}}, nil
}

// SecuredTokenTransfer is an auto generated Go binding around an Ethereum contract.
type SecuredTokenTransfer struct {
	SecuredTokenTransferCaller     // Read-only binding to the contract
	SecuredTokenTransferTransactor // Write-only binding to the contract
	SecuredTokenTransferFilterer   // Log filterer for contract events
}

// SecuredTokenTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecuredTokenTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecuredTokenTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecuredTokenTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecuredTokenTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SecuredTokenTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecuredTokenTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecuredTokenTransferSession struct {
	Contract     *SecuredTokenTransfer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SecuredTokenTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecuredTokenTransferCallerSession struct {
	Contract *SecuredTokenTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SecuredTokenTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecuredTokenTransferTransactorSession struct {
	Contract     *SecuredTokenTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SecuredTokenTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecuredTokenTransferRaw struct {
	Contract *SecuredTokenTransfer // Generic contract binding to access the raw methods on
}

// SecuredTokenTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecuredTokenTransferCallerRaw struct {
	Contract *SecuredTokenTransferCaller // Generic read-only contract binding to access the raw methods on
}

// SecuredTokenTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecuredTokenTransferTransactorRaw struct {
	Contract *SecuredTokenTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecuredTokenTransfer creates a new instance of SecuredTokenTransfer, bound to a specific deployed contract.
func NewSecuredTokenTransfer(address common.Address, backend bind.ContractBackend) (*SecuredTokenTransfer, error) {
	contract, err := bindSecuredTokenTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecuredTokenTransfer{SecuredTokenTransferCaller: SecuredTokenTransferCaller{contract: contract}, SecuredTokenTransferTransactor: SecuredTokenTransferTransactor{contract: contract}, SecuredTokenTransferFilterer: SecuredTokenTransferFilterer{contract: contract}}, nil
}

// NewSecuredTokenTransferCaller creates a new read-only instance of SecuredTokenTransfer, bound to a specific deployed contract.
func NewSecuredTokenTransferCaller(address common.Address, caller bind.ContractCaller) (*SecuredTokenTransferCaller, error) {
	contract, err := bindSecuredTokenTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecuredTokenTransferCaller{contract: contract}, nil
}

// NewSecuredTokenTransferTransactor creates a new write-only instance of SecuredTokenTransfer, bound to a specific deployed contract.
func NewSecuredTokenTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*SecuredTokenTransferTransactor, error) {
	contract, err := bindSecuredTokenTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecuredTokenTransferTransactor{contract: contract}, nil
}

// NewSecuredTokenTransferFilterer creates a new log filterer instance of SecuredTokenTransfer, bound to a specific deployed contract.
func NewSecuredTokenTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*SecuredTokenTransferFilterer, error) {
	contract, err := bindSecuredTokenTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecuredTokenTransferFilterer{contract: contract}, nil
}

// bindSecuredTokenTransfer binds a generic wrapper to an already deployed contract.
func bindSecuredTokenTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SecuredTokenTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecuredTokenTransfer *SecuredTokenTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecuredTokenTransfer.Contract.SecuredTokenTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecuredTokenTransfer *SecuredTokenTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecuredTokenTransfer.Contract.SecuredTokenTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecuredTokenTransfer *SecuredTokenTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecuredTokenTransfer.Contract.SecuredTokenTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecuredTokenTransfer *SecuredTokenTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecuredTokenTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecuredTokenTransfer *SecuredTokenTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecuredTokenTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecuredTokenTransfer *SecuredTokenTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecuredTokenTransfer.Contract.contract.Transact(opts, method, params...)
}

// SelfAuthorizedMetaData contains all meta data concerning the SelfAuthorized contract.
var SelfAuthorizedMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122056b0797a519b0ef612c6219051e87a88476b872b624afed8ef81f4ce191285da64736f6c634300080d0033",
}

// SelfAuthorizedABI is the input ABI used to generate the binding from.
// Deprecated: Use SelfAuthorizedMetaData.ABI instead.
var SelfAuthorizedABI = SelfAuthorizedMetaData.ABI

// SelfAuthorizedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SelfAuthorizedMetaData.Bin instead.
var SelfAuthorizedBin = SelfAuthorizedMetaData.Bin

// DeploySelfAuthorized deploys a new Ethereum contract, binding an instance of SelfAuthorized to it.
func DeploySelfAuthorized(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SelfAuthorized, error) {
	parsed, err := SelfAuthorizedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SelfAuthorizedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SelfAuthorized{SelfAuthorizedCaller: SelfAuthorizedCaller{contract: contract}, SelfAuthorizedTransactor: SelfAuthorizedTransactor{contract: contract}, SelfAuthorizedFilterer: SelfAuthorizedFilterer{contract: contract}}, nil
}

// SelfAuthorized is an auto generated Go binding around an Ethereum contract.
type SelfAuthorized struct {
	SelfAuthorizedCaller     // Read-only binding to the contract
	SelfAuthorizedTransactor // Write-only binding to the contract
	SelfAuthorizedFilterer   // Log filterer for contract events
}

// SelfAuthorizedCaller is an auto generated read-only Go binding around an Ethereum contract.
type SelfAuthorizedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthorizedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SelfAuthorizedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthorizedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SelfAuthorizedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthorizedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SelfAuthorizedSession struct {
	Contract     *SelfAuthorized   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SelfAuthorizedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SelfAuthorizedCallerSession struct {
	Contract *SelfAuthorizedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SelfAuthorizedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SelfAuthorizedTransactorSession struct {
	Contract     *SelfAuthorizedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SelfAuthorizedRaw is an auto generated low-level Go binding around an Ethereum contract.
type SelfAuthorizedRaw struct {
	Contract *SelfAuthorized // Generic contract binding to access the raw methods on
}

// SelfAuthorizedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SelfAuthorizedCallerRaw struct {
	Contract *SelfAuthorizedCaller // Generic read-only contract binding to access the raw methods on
}

// SelfAuthorizedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SelfAuthorizedTransactorRaw struct {
	Contract *SelfAuthorizedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSelfAuthorized creates a new instance of SelfAuthorized, bound to a specific deployed contract.
func NewSelfAuthorized(address common.Address, backend bind.ContractBackend) (*SelfAuthorized, error) {
	contract, err := bindSelfAuthorized(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfAuthorized{SelfAuthorizedCaller: SelfAuthorizedCaller{contract: contract}, SelfAuthorizedTransactor: SelfAuthorizedTransactor{contract: contract}, SelfAuthorizedFilterer: SelfAuthorizedFilterer{contract: contract}}, nil
}

// NewSelfAuthorizedCaller creates a new read-only instance of SelfAuthorized, bound to a specific deployed contract.
func NewSelfAuthorizedCaller(address common.Address, caller bind.ContractCaller) (*SelfAuthorizedCaller, error) {
	contract, err := bindSelfAuthorized(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfAuthorizedCaller{contract: contract}, nil
}

// NewSelfAuthorizedTransactor creates a new write-only instance of SelfAuthorized, bound to a specific deployed contract.
func NewSelfAuthorizedTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfAuthorizedTransactor, error) {
	contract, err := bindSelfAuthorized(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfAuthorizedTransactor{contract: contract}, nil
}

// NewSelfAuthorizedFilterer creates a new log filterer instance of SelfAuthorized, bound to a specific deployed contract.
func NewSelfAuthorizedFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfAuthorizedFilterer, error) {
	contract, err := bindSelfAuthorized(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfAuthorizedFilterer{contract: contract}, nil
}

// bindSelfAuthorized binds a generic wrapper to an already deployed contract.
func bindSelfAuthorized(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SelfAuthorizedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfAuthorized *SelfAuthorizedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfAuthorized.Contract.SelfAuthorizedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfAuthorized *SelfAuthorizedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfAuthorized.Contract.SelfAuthorizedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfAuthorized *SelfAuthorizedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfAuthorized.Contract.SelfAuthorizedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfAuthorized *SelfAuthorizedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfAuthorized.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfAuthorized *SelfAuthorizedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfAuthorized.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfAuthorized *SelfAuthorizedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfAuthorized.Contract.contract.Transact(opts, method, params...)
}

// SignatureDecoderMetaData contains all meta data concerning the SignatureDecoder contract.
var SignatureDecoderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122005e1652404d9982f9c9bccbd494043f6722d09c1bf6301143fc698b5f64587ae64736f6c634300080d0033",
}

// SignatureDecoderABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureDecoderMetaData.ABI instead.
var SignatureDecoderABI = SignatureDecoderMetaData.ABI

// SignatureDecoderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignatureDecoderMetaData.Bin instead.
var SignatureDecoderBin = SignatureDecoderMetaData.Bin

// DeploySignatureDecoder deploys a new Ethereum contract, binding an instance of SignatureDecoder to it.
func DeploySignatureDecoder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignatureDecoder, error) {
	parsed, err := SignatureDecoderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignatureDecoderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignatureDecoder{SignatureDecoderCaller: SignatureDecoderCaller{contract: contract}, SignatureDecoderTransactor: SignatureDecoderTransactor{contract: contract}, SignatureDecoderFilterer: SignatureDecoderFilterer{contract: contract}}, nil
}

// SignatureDecoder is an auto generated Go binding around an Ethereum contract.
type SignatureDecoder struct {
	SignatureDecoderCaller     // Read-only binding to the contract
	SignatureDecoderTransactor // Write-only binding to the contract
	SignatureDecoderFilterer   // Log filterer for contract events
}

// SignatureDecoderCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureDecoderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureDecoderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureDecoderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureDecoderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureDecoderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureDecoderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureDecoderSession struct {
	Contract     *SignatureDecoder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignatureDecoderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureDecoderCallerSession struct {
	Contract *SignatureDecoderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SignatureDecoderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureDecoderTransactorSession struct {
	Contract     *SignatureDecoderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SignatureDecoderRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureDecoderRaw struct {
	Contract *SignatureDecoder // Generic contract binding to access the raw methods on
}

// SignatureDecoderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureDecoderCallerRaw struct {
	Contract *SignatureDecoderCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureDecoderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureDecoderTransactorRaw struct {
	Contract *SignatureDecoderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureDecoder creates a new instance of SignatureDecoder, bound to a specific deployed contract.
func NewSignatureDecoder(address common.Address, backend bind.ContractBackend) (*SignatureDecoder, error) {
	contract, err := bindSignatureDecoder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureDecoder{SignatureDecoderCaller: SignatureDecoderCaller{contract: contract}, SignatureDecoderTransactor: SignatureDecoderTransactor{contract: contract}, SignatureDecoderFilterer: SignatureDecoderFilterer{contract: contract}}, nil
}

// NewSignatureDecoderCaller creates a new read-only instance of SignatureDecoder, bound to a specific deployed contract.
func NewSignatureDecoderCaller(address common.Address, caller bind.ContractCaller) (*SignatureDecoderCaller, error) {
	contract, err := bindSignatureDecoder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureDecoderCaller{contract: contract}, nil
}

// NewSignatureDecoderTransactor creates a new write-only instance of SignatureDecoder, bound to a specific deployed contract.
func NewSignatureDecoderTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureDecoderTransactor, error) {
	contract, err := bindSignatureDecoder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureDecoderTransactor{contract: contract}, nil
}

// NewSignatureDecoderFilterer creates a new log filterer instance of SignatureDecoder, bound to a specific deployed contract.
func NewSignatureDecoderFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureDecoderFilterer, error) {
	contract, err := bindSignatureDecoder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureDecoderFilterer{contract: contract}, nil
}

// bindSignatureDecoder binds a generic wrapper to an already deployed contract.
func bindSignatureDecoder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureDecoderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureDecoder *SignatureDecoderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureDecoder.Contract.SignatureDecoderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureDecoder *SignatureDecoderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureDecoder.Contract.SignatureDecoderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureDecoder *SignatureDecoderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureDecoder.Contract.SignatureDecoderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureDecoder *SignatureDecoderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureDecoder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureDecoder *SignatureDecoderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureDecoder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureDecoder *SignatureDecoderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureDecoder.Contract.contract.Transact(opts, method, params...)
}

// SingletonMetaData contains all meta data concerning the Singleton contract.
var SingletonMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220a3951dee9496fc0120e66c0ee7da9c607d9b6b344dc51757c36615dde0b76c5464736f6c634300080d0033",
}

// SingletonABI is the input ABI used to generate the binding from.
// Deprecated: Use SingletonMetaData.ABI instead.
var SingletonABI = SingletonMetaData.ABI

// SingletonBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SingletonMetaData.Bin instead.
var SingletonBin = SingletonMetaData.Bin

// DeploySingleton deploys a new Ethereum contract, binding an instance of Singleton to it.
func DeploySingleton(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Singleton, error) {
	parsed, err := SingletonMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SingletonBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Singleton{SingletonCaller: SingletonCaller{contract: contract}, SingletonTransactor: SingletonTransactor{contract: contract}, SingletonFilterer: SingletonFilterer{contract: contract}}, nil
}

// Singleton is an auto generated Go binding around an Ethereum contract.
type Singleton struct {
	SingletonCaller     // Read-only binding to the contract
	SingletonTransactor // Write-only binding to the contract
	SingletonFilterer   // Log filterer for contract events
}

// SingletonCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingletonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingletonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingletonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingletonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SingletonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingletonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingletonSession struct {
	Contract     *Singleton        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SingletonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingletonCallerSession struct {
	Contract *SingletonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SingletonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingletonTransactorSession struct {
	Contract     *SingletonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SingletonRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingletonRaw struct {
	Contract *Singleton // Generic contract binding to access the raw methods on
}

// SingletonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingletonCallerRaw struct {
	Contract *SingletonCaller // Generic read-only contract binding to access the raw methods on
}

// SingletonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingletonTransactorRaw struct {
	Contract *SingletonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingleton creates a new instance of Singleton, bound to a specific deployed contract.
func NewSingleton(address common.Address, backend bind.ContractBackend) (*Singleton, error) {
	contract, err := bindSingleton(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Singleton{SingletonCaller: SingletonCaller{contract: contract}, SingletonTransactor: SingletonTransactor{contract: contract}, SingletonFilterer: SingletonFilterer{contract: contract}}, nil
}

// NewSingletonCaller creates a new read-only instance of Singleton, bound to a specific deployed contract.
func NewSingletonCaller(address common.Address, caller bind.ContractCaller) (*SingletonCaller, error) {
	contract, err := bindSingleton(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingletonCaller{contract: contract}, nil
}

// NewSingletonTransactor creates a new write-only instance of Singleton, bound to a specific deployed contract.
func NewSingletonTransactor(address common.Address, transactor bind.ContractTransactor) (*SingletonTransactor, error) {
	contract, err := bindSingleton(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingletonTransactor{contract: contract}, nil
}

// NewSingletonFilterer creates a new log filterer instance of Singleton, bound to a specific deployed contract.
func NewSingletonFilterer(address common.Address, filterer bind.ContractFilterer) (*SingletonFilterer, error) {
	contract, err := bindSingleton(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingletonFilterer{contract: contract}, nil
}

// bindSingleton binds a generic wrapper to an already deployed contract.
func bindSingleton(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SingletonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Singleton *SingletonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Singleton.Contract.SingletonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Singleton *SingletonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Singleton.Contract.SingletonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Singleton *SingletonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Singleton.Contract.SingletonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Singleton *SingletonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Singleton.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Singleton *SingletonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Singleton.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Singleton *SingletonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Singleton.Contract.contract.Transact(opts, method, params...)
}

// StorageAccessibleMetaData contains all meta data concerning the StorageAccessible contract.
var StorageAccessibleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5624b25b": "getStorageAt(uint256,uint256)",
		"b4faba09": "simulateAndRevert(address,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610303806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635624b25b1461003b578063b4faba0914610064575b600080fd5b61004e610049366004610122565b610079565b60405161005b9190610144565b60405180910390f35b6100776100723660046101af565b6100ff565b005b60606000610088836020610295565b67ffffffffffffffff8111156100a0576100a0610199565b6040519080825280601f01601f1916602001820160405280156100ca576020820181803683370190505b50905060005b838110156100f75784810154602080830284010152806100ef816102b4565b9150506100d0565b509392505050565b600080825160208401855af480600052503d6020523d600060403e60403d016000fd5b6000806040838503121561013557600080fd5b50508035926020909101359150565b600060208083528351808285015260005b8181101561017157858101830151858201604001528201610155565b81811115610183576000604083870101525b50601f01601f1916929092016040019392505050565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156101c257600080fd5b82356001600160a01b03811681146101d957600080fd5b9150602083013567ffffffffffffffff808211156101f657600080fd5b818501915085601f83011261020a57600080fd5b81358181111561021c5761021c610199565b604051601f8201601f19908116603f0116810190838211818310171561024457610244610199565b8160405282815288602084870101111561025d57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b634e487b7160e01b600052601160045260246000fd5b60008160001904831182151516156102af576102af61027f565b500290565b6000600182016102c6576102c661027f565b506001019056fea26469706673582212204b82da09855dc9cd636a09bbeb5517a91a1e774b6ab2eefb4282e7ab54a61f9d64736f6c634300080d0033",
}

// StorageAccessibleABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageAccessibleMetaData.ABI instead.
var StorageAccessibleABI = StorageAccessibleMetaData.ABI

// Deprecated: Use StorageAccessibleMetaData.Sigs instead.
// StorageAccessibleFuncSigs maps the 4-byte function signature to its string representation.
var StorageAccessibleFuncSigs = StorageAccessibleMetaData.Sigs

// StorageAccessibleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageAccessibleMetaData.Bin instead.
var StorageAccessibleBin = StorageAccessibleMetaData.Bin

// DeployStorageAccessible deploys a new Ethereum contract, binding an instance of StorageAccessible to it.
func DeployStorageAccessible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageAccessible, error) {
	parsed, err := StorageAccessibleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageAccessibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageAccessible{StorageAccessibleCaller: StorageAccessibleCaller{contract: contract}, StorageAccessibleTransactor: StorageAccessibleTransactor{contract: contract}, StorageAccessibleFilterer: StorageAccessibleFilterer{contract: contract}}, nil
}

// StorageAccessible is an auto generated Go binding around an Ethereum contract.
type StorageAccessible struct {
	StorageAccessibleCaller     // Read-only binding to the contract
	StorageAccessibleTransactor // Write-only binding to the contract
	StorageAccessibleFilterer   // Log filterer for contract events
}

// StorageAccessibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageAccessibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageAccessibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageAccessibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageAccessibleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageAccessibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageAccessibleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageAccessibleSession struct {
	Contract     *StorageAccessible // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageAccessibleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageAccessibleCallerSession struct {
	Contract *StorageAccessibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StorageAccessibleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageAccessibleTransactorSession struct {
	Contract     *StorageAccessibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StorageAccessibleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageAccessibleRaw struct {
	Contract *StorageAccessible // Generic contract binding to access the raw methods on
}

// StorageAccessibleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageAccessibleCallerRaw struct {
	Contract *StorageAccessibleCaller // Generic read-only contract binding to access the raw methods on
}

// StorageAccessibleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageAccessibleTransactorRaw struct {
	Contract *StorageAccessibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageAccessible creates a new instance of StorageAccessible, bound to a specific deployed contract.
func NewStorageAccessible(address common.Address, backend bind.ContractBackend) (*StorageAccessible, error) {
	contract, err := bindStorageAccessible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageAccessible{StorageAccessibleCaller: StorageAccessibleCaller{contract: contract}, StorageAccessibleTransactor: StorageAccessibleTransactor{contract: contract}, StorageAccessibleFilterer: StorageAccessibleFilterer{contract: contract}}, nil
}

// NewStorageAccessibleCaller creates a new read-only instance of StorageAccessible, bound to a specific deployed contract.
func NewStorageAccessibleCaller(address common.Address, caller bind.ContractCaller) (*StorageAccessibleCaller, error) {
	contract, err := bindStorageAccessible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageAccessibleCaller{contract: contract}, nil
}

// NewStorageAccessibleTransactor creates a new write-only instance of StorageAccessible, bound to a specific deployed contract.
func NewStorageAccessibleTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageAccessibleTransactor, error) {
	contract, err := bindStorageAccessible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageAccessibleTransactor{contract: contract}, nil
}

// NewStorageAccessibleFilterer creates a new log filterer instance of StorageAccessible, bound to a specific deployed contract.
func NewStorageAccessibleFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageAccessibleFilterer, error) {
	contract, err := bindStorageAccessible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageAccessibleFilterer{contract: contract}, nil
}

// bindStorageAccessible binds a generic wrapper to an already deployed contract.
func bindStorageAccessible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageAccessibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageAccessible *StorageAccessibleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageAccessible.Contract.StorageAccessibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageAccessible *StorageAccessibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageAccessible.Contract.StorageAccessibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageAccessible *StorageAccessibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageAccessible.Contract.StorageAccessibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageAccessible *StorageAccessibleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageAccessible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageAccessible *StorageAccessibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageAccessible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageAccessible *StorageAccessibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageAccessible.Contract.contract.Transact(opts, method, params...)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_StorageAccessible *StorageAccessibleCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StorageAccessible.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_StorageAccessible *StorageAccessibleSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _StorageAccessible.Contract.GetStorageAt(&_StorageAccessible.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_StorageAccessible *StorageAccessibleCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _StorageAccessible.Contract.GetStorageAt(&_StorageAccessible.CallOpts, offset, length)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_StorageAccessible *StorageAccessibleTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _StorageAccessible.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_StorageAccessible *StorageAccessibleSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _StorageAccessible.Contract.SimulateAndRevert(&_StorageAccessible.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_StorageAccessible *StorageAccessibleTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _StorageAccessible.Contract.SimulateAndRevert(&_StorageAccessible.TransactOpts, targetContract, calldataPayload)
}
