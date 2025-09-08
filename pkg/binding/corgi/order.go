// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package corgi

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

// OrderMetaData contains all meta data concerning the Order contract.
var OrderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BetEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"DrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gameAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"autoBetAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gameCategoryAddress_\",\"type\":\"address\"}],\"name\":\"admin_set_env\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoBet\",\"outputs\":[{\"internalType\":\"contractIAutoBet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gameId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"autoId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"}],\"name\":\"checkOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"autoId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"seed_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rate_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hashExpired_\",\"type\":\"bool\"}],\"name\":\"draw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"contractIFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractIGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameCategory\",\"outputs\":[{\"internalType\":\"contractIGameCategory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"}],\"name\":\"getOrderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isOrderCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"seed\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ctime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OrderABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderMetaData.ABI instead.
var OrderABI = OrderMetaData.ABI

// Order is an auto generated Go binding around an Ethereum contract.
type Order struct {
	OrderCaller     // Read-only binding to the contract
	OrderTransactor // Write-only binding to the contract
	OrderFilterer   // Log filterer for contract events
}

// OrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderSession struct {
	Contract     *Order            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderCallerSession struct {
	Contract *OrderCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderTransactorSession struct {
	Contract     *OrderTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderRaw struct {
	Contract *Order // Generic contract binding to access the raw methods on
}

// OrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderCallerRaw struct {
	Contract *OrderCaller // Generic read-only contract binding to access the raw methods on
}

// OrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderTransactorRaw struct {
	Contract *OrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrder creates a new instance of Order, bound to a specific deployed contract.
func NewOrder(address common.Address, backend bind.ContractBackend) (*Order, error) {
	contract, err := bindOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Order{OrderCaller: OrderCaller{contract: contract}, OrderTransactor: OrderTransactor{contract: contract}, OrderFilterer: OrderFilterer{contract: contract}}, nil
}

// NewOrderCaller creates a new read-only instance of Order, bound to a specific deployed contract.
func NewOrderCaller(address common.Address, caller bind.ContractCaller) (*OrderCaller, error) {
	contract, err := bindOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderCaller{contract: contract}, nil
}

// NewOrderTransactor creates a new write-only instance of Order, bound to a specific deployed contract.
func NewOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderTransactor, error) {
	contract, err := bindOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderTransactor{contract: contract}, nil
}

// NewOrderFilterer creates a new log filterer instance of Order, bound to a specific deployed contract.
func NewOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderFilterer, error) {
	contract, err := bindOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderFilterer{contract: contract}, nil
}

// bindOrder binds a generic wrapper to an already deployed contract.
func bindOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Order.Contract.OrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Order.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Order *OrderCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Order *OrderSession) ADMINROLE() ([32]byte, error) {
	return _Order.Contract.ADMINROLE(&_Order.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Order *OrderCallerSession) ADMINROLE() ([32]byte, error) {
	return _Order.Contract.ADMINROLE(&_Order.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Order *OrderCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Order *OrderSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Order.Contract.DEFAULTADMINROLE(&_Order.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Order *OrderCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Order.Contract.DEFAULTADMINROLE(&_Order.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Order *OrderCaller) SERVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "SERVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Order *OrderSession) SERVERROLE() ([32]byte, error) {
	return _Order.Contract.SERVERROLE(&_Order.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Order *OrderCallerSession) SERVERROLE() ([32]byte, error) {
	return _Order.Contract.SERVERROLE(&_Order.CallOpts)
}

// AutoBet is a free data retrieval call binding the contract method 0x09392438.
//
// Solidity: function autoBet() view returns(address)
func (_Order *OrderCaller) AutoBet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "autoBet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AutoBet is a free data retrieval call binding the contract method 0x09392438.
//
// Solidity: function autoBet() view returns(address)
func (_Order *OrderSession) AutoBet() (common.Address, error) {
	return _Order.Contract.AutoBet(&_Order.CallOpts)
}

// AutoBet is a free data retrieval call binding the contract method 0x09392438.
//
// Solidity: function autoBet() view returns(address)
func (_Order *OrderCallerSession) AutoBet() (common.Address, error) {
	return _Order.Contract.AutoBet(&_Order.CallOpts)
}

// CheckOrder is a free data retrieval call binding the contract method 0x48efc672.
//
// Solidity: function checkOrder(uint256 orderId_) view returns()
func (_Order *OrderCaller) CheckOrder(opts *bind.CallOpts, orderId_ *big.Int) error {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "checkOrder", orderId_)

	if err != nil {
		return err
	}

	return err

}

// CheckOrder is a free data retrieval call binding the contract method 0x48efc672.
//
// Solidity: function checkOrder(uint256 orderId_) view returns()
func (_Order *OrderSession) CheckOrder(orderId_ *big.Int) error {
	return _Order.Contract.CheckOrder(&_Order.CallOpts, orderId_)
}

// CheckOrder is a free data retrieval call binding the contract method 0x48efc672.
//
// Solidity: function checkOrder(uint256 orderId_) view returns()
func (_Order *OrderCallerSession) CheckOrder(orderId_ *big.Int) error {
	return _Order.Contract.CheckOrder(&_Order.CallOpts, orderId_)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_Order *OrderCaller) Fee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_Order *OrderSession) Fee() (common.Address, error) {
	return _Order.Contract.Fee(&_Order.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_Order *OrderCallerSession) Fee() (common.Address, error) {
	return _Order.Contract.Fee(&_Order.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_Order *OrderCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_Order *OrderSession) Game() (common.Address, error) {
	return _Order.Contract.Game(&_Order.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_Order *OrderCallerSession) Game() (common.Address, error) {
	return _Order.Contract.Game(&_Order.CallOpts)
}

// GameCategory is a free data retrieval call binding the contract method 0x816928f0.
//
// Solidity: function gameCategory() view returns(address)
func (_Order *OrderCaller) GameCategory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "gameCategory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameCategory is a free data retrieval call binding the contract method 0x816928f0.
//
// Solidity: function gameCategory() view returns(address)
func (_Order *OrderSession) GameCategory() (common.Address, error) {
	return _Order.Contract.GameCategory(&_Order.CallOpts)
}

// GameCategory is a free data retrieval call binding the contract method 0x816928f0.
//
// Solidity: function gameCategory() view returns(address)
func (_Order *OrderCallerSession) GameCategory() (common.Address, error) {
	return _Order.Contract.GameCategory(&_Order.CallOpts)
}

// GetNewOrderId is a free data retrieval call binding the contract method 0x2d0e6001.
//
// Solidity: function getNewOrderId() view returns(uint256)
func (_Order *OrderCaller) GetNewOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "getNewOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewOrderId is a free data retrieval call binding the contract method 0x2d0e6001.
//
// Solidity: function getNewOrderId() view returns(uint256)
func (_Order *OrderSession) GetNewOrderId() (*big.Int, error) {
	return _Order.Contract.GetNewOrderId(&_Order.CallOpts)
}

// GetNewOrderId is a free data retrieval call binding the contract method 0x2d0e6001.
//
// Solidity: function getNewOrderId() view returns(uint256)
func (_Order *OrderCallerSession) GetNewOrderId() (*big.Int, error) {
	return _Order.Contract.GetNewOrderId(&_Order.CallOpts)
}

// GetOrderAmount is a free data retrieval call binding the contract method 0x67bb1fcf.
//
// Solidity: function getOrderAmount(uint256 orderId_) view returns(uint256)
func (_Order *OrderCaller) GetOrderAmount(opts *bind.CallOpts, orderId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "getOrderAmount", orderId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderAmount is a free data retrieval call binding the contract method 0x67bb1fcf.
//
// Solidity: function getOrderAmount(uint256 orderId_) view returns(uint256)
func (_Order *OrderSession) GetOrderAmount(orderId_ *big.Int) (*big.Int, error) {
	return _Order.Contract.GetOrderAmount(&_Order.CallOpts, orderId_)
}

// GetOrderAmount is a free data retrieval call binding the contract method 0x67bb1fcf.
//
// Solidity: function getOrderAmount(uint256 orderId_) view returns(uint256)
func (_Order *OrderCallerSession) GetOrderAmount(orderId_ *big.Int) (*big.Int, error) {
	return _Order.Contract.GetOrderAmount(&_Order.CallOpts, orderId_)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Order *OrderCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Order *OrderSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Order.Contract.GetRoleAdmin(&_Order.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Order *OrderCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Order.Contract.GetRoleAdmin(&_Order.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Order *OrderCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Order *OrderSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Order.Contract.HasRole(&_Order.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Order *OrderCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Order.Contract.HasRole(&_Order.CallOpts, role, account)
}

// IsOrderCanceled is a free data retrieval call binding the contract method 0x77632ec2.
//
// Solidity: function isOrderCanceled(uint256 ) view returns(bool)
func (_Order *OrderCaller) IsOrderCanceled(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "isOrderCanceled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderCanceled is a free data retrieval call binding the contract method 0x77632ec2.
//
// Solidity: function isOrderCanceled(uint256 ) view returns(bool)
func (_Order *OrderSession) IsOrderCanceled(arg0 *big.Int) (bool, error) {
	return _Order.Contract.IsOrderCanceled(&_Order.CallOpts, arg0)
}

// IsOrderCanceled is a free data retrieval call binding the contract method 0x77632ec2.
//
// Solidity: function isOrderCanceled(uint256 ) view returns(bool)
func (_Order *OrderCallerSession) IsOrderCanceled(arg0 *big.Int) (bool, error) {
	return _Order.Contract.IsOrderCanceled(&_Order.CallOpts, arg0)
}

// OrderId is a free data retrieval call binding the contract method 0x163de5e5.
//
// Solidity: function orderId() view returns(uint256)
func (_Order *OrderCaller) OrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "orderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderId is a free data retrieval call binding the contract method 0x163de5e5.
//
// Solidity: function orderId() view returns(uint256)
func (_Order *OrderSession) OrderId() (*big.Int, error) {
	return _Order.Contract.OrderId(&_Order.CallOpts)
}

// OrderId is a free data retrieval call binding the contract method 0x163de5e5.
//
// Solidity: function orderId() view returns(uint256)
func (_Order *OrderCallerSession) OrderId() (*big.Int, error) {
	return _Order.Contract.OrderId(&_Order.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 gameId, uint256 autoId, uint256 orderId, bytes32 hash, string seed, address user, uint256 amount, uint256 tokenId, uint256 bonus, uint256 status, uint256 ctime)
func (_Order *OrderCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	GameId  *big.Int
	AutoId  *big.Int
	OrderId *big.Int
	Hash    [32]byte
	Seed    string
	User    common.Address
	Amount  *big.Int
	TokenId *big.Int
	Bonus   *big.Int
	Status  *big.Int
	Ctime   *big.Int
}, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		GameId  *big.Int
		AutoId  *big.Int
		OrderId *big.Int
		Hash    [32]byte
		Seed    string
		User    common.Address
		Amount  *big.Int
		TokenId *big.Int
		Bonus   *big.Int
		Status  *big.Int
		Ctime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AutoId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OrderId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Hash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Seed = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.User = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Bonus = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Ctime = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 gameId, uint256 autoId, uint256 orderId, bytes32 hash, string seed, address user, uint256 amount, uint256 tokenId, uint256 bonus, uint256 status, uint256 ctime)
func (_Order *OrderSession) Orders(arg0 *big.Int) (struct {
	GameId  *big.Int
	AutoId  *big.Int
	OrderId *big.Int
	Hash    [32]byte
	Seed    string
	User    common.Address
	Amount  *big.Int
	TokenId *big.Int
	Bonus   *big.Int
	Status  *big.Int
	Ctime   *big.Int
}, error) {
	return _Order.Contract.Orders(&_Order.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 gameId, uint256 autoId, uint256 orderId, bytes32 hash, string seed, address user, uint256 amount, uint256 tokenId, uint256 bonus, uint256 status, uint256 ctime)
func (_Order *OrderCallerSession) Orders(arg0 *big.Int) (struct {
	GameId  *big.Int
	AutoId  *big.Int
	OrderId *big.Int
	Hash    [32]byte
	Seed    string
	User    common.Address
	Amount  *big.Int
	TokenId *big.Int
	Bonus   *big.Int
	Status  *big.Int
	Ctime   *big.Int
}, error) {
	return _Order.Contract.Orders(&_Order.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Order *OrderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Order *OrderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Order.Contract.SupportsInterface(&_Order.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Order *OrderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Order.Contract.SupportsInterface(&_Order.CallOpts, interfaceId)
}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(uint256)
func (_Order *OrderCaller) UserOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "userOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(uint256)
func (_Order *OrderSession) UserOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Order.Contract.UserOrders(&_Order.CallOpts, arg0, arg1)
}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(uint256)
func (_Order *OrderCallerSession) UserOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Order.Contract.UserOrders(&_Order.CallOpts, arg0, arg1)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x1f137039.
//
// Solidity: function admin_set_env(address feeAddress_, address gameAddress_, address autoBetAddress_, address gameCategoryAddress_) returns()
func (_Order *OrderTransactor) AdminSetEnv(opts *bind.TransactOpts, feeAddress_ common.Address, gameAddress_ common.Address, autoBetAddress_ common.Address, gameCategoryAddress_ common.Address) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "admin_set_env", feeAddress_, gameAddress_, autoBetAddress_, gameCategoryAddress_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x1f137039.
//
// Solidity: function admin_set_env(address feeAddress_, address gameAddress_, address autoBetAddress_, address gameCategoryAddress_) returns()
func (_Order *OrderSession) AdminSetEnv(feeAddress_ common.Address, gameAddress_ common.Address, autoBetAddress_ common.Address, gameCategoryAddress_ common.Address) (*types.Transaction, error) {
	return _Order.Contract.AdminSetEnv(&_Order.TransactOpts, feeAddress_, gameAddress_, autoBetAddress_, gameCategoryAddress_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x1f137039.
//
// Solidity: function admin_set_env(address feeAddress_, address gameAddress_, address autoBetAddress_, address gameCategoryAddress_) returns()
func (_Order *OrderTransactorSession) AdminSetEnv(feeAddress_ common.Address, gameAddress_ common.Address, autoBetAddress_ common.Address, gameCategoryAddress_ common.Address) (*types.Transaction, error) {
	return _Order.Contract.AdminSetEnv(&_Order.TransactOpts, feeAddress_, gameAddress_, autoBetAddress_, gameCategoryAddress_)
}

// Bet is a paid mutator transaction binding the contract method 0xcebcd88f.
//
// Solidity: function bet(address user_, uint256 gameId_, uint256 autoId_, uint256 orderId_, uint256 tokenId_, uint256 amount_, bytes32 hash_, bytes data_) returns()
func (_Order *OrderTransactor) Bet(opts *bind.TransactOpts, user_ common.Address, gameId_ *big.Int, autoId_ *big.Int, orderId_ *big.Int, tokenId_ *big.Int, amount_ *big.Int, hash_ [32]byte, data_ []byte) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "bet", user_, gameId_, autoId_, orderId_, tokenId_, amount_, hash_, data_)
}

// Bet is a paid mutator transaction binding the contract method 0xcebcd88f.
//
// Solidity: function bet(address user_, uint256 gameId_, uint256 autoId_, uint256 orderId_, uint256 tokenId_, uint256 amount_, bytes32 hash_, bytes data_) returns()
func (_Order *OrderSession) Bet(user_ common.Address, gameId_ *big.Int, autoId_ *big.Int, orderId_ *big.Int, tokenId_ *big.Int, amount_ *big.Int, hash_ [32]byte, data_ []byte) (*types.Transaction, error) {
	return _Order.Contract.Bet(&_Order.TransactOpts, user_, gameId_, autoId_, orderId_, tokenId_, amount_, hash_, data_)
}

// Bet is a paid mutator transaction binding the contract method 0xcebcd88f.
//
// Solidity: function bet(address user_, uint256 gameId_, uint256 autoId_, uint256 orderId_, uint256 tokenId_, uint256 amount_, bytes32 hash_, bytes data_) returns()
func (_Order *OrderTransactorSession) Bet(user_ common.Address, gameId_ *big.Int, autoId_ *big.Int, orderId_ *big.Int, tokenId_ *big.Int, amount_ *big.Int, hash_ [32]byte, data_ []byte) (*types.Transaction, error) {
	return _Order.Contract.Bet(&_Order.TransactOpts, user_, gameId_, autoId_, orderId_, tokenId_, amount_, hash_, data_)
}

// Draw is a paid mutator transaction binding the contract method 0xd9b849fa.
//
// Solidity: function draw(uint256 gameId_, uint256 autoId_, uint256 orderId_, string seed_, uint256 rate_, bool hashExpired_) returns()
func (_Order *OrderTransactor) Draw(opts *bind.TransactOpts, gameId_ *big.Int, autoId_ *big.Int, orderId_ *big.Int, seed_ string, rate_ *big.Int, hashExpired_ bool) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "draw", gameId_, autoId_, orderId_, seed_, rate_, hashExpired_)
}

// Draw is a paid mutator transaction binding the contract method 0xd9b849fa.
//
// Solidity: function draw(uint256 gameId_, uint256 autoId_, uint256 orderId_, string seed_, uint256 rate_, bool hashExpired_) returns()
func (_Order *OrderSession) Draw(gameId_ *big.Int, autoId_ *big.Int, orderId_ *big.Int, seed_ string, rate_ *big.Int, hashExpired_ bool) (*types.Transaction, error) {
	return _Order.Contract.Draw(&_Order.TransactOpts, gameId_, autoId_, orderId_, seed_, rate_, hashExpired_)
}

// Draw is a paid mutator transaction binding the contract method 0xd9b849fa.
//
// Solidity: function draw(uint256 gameId_, uint256 autoId_, uint256 orderId_, string seed_, uint256 rate_, bool hashExpired_) returns()
func (_Order *OrderTransactorSession) Draw(gameId_ *big.Int, autoId_ *big.Int, orderId_ *big.Int, seed_ string, rate_ *big.Int, hashExpired_ bool) (*types.Transaction, error) {
	return _Order.Contract.Draw(&_Order.TransactOpts, gameId_, autoId_, orderId_, seed_, rate_, hashExpired_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Order *OrderTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Order *OrderSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Order.Contract.GrantRole(&_Order.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Order *OrderTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Order.Contract.GrantRole(&_Order.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Order *OrderTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Order *OrderSession) Initialize() (*types.Transaction, error) {
	return _Order.Contract.Initialize(&_Order.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Order *OrderTransactorSession) Initialize() (*types.Transaction, error) {
	return _Order.Contract.Initialize(&_Order.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Order *OrderTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Order *OrderSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Order.Contract.RenounceRole(&_Order.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Order *OrderTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Order.Contract.RenounceRole(&_Order.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Order *OrderTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Order *OrderSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Order.Contract.RevokeRole(&_Order.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Order *OrderTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Order.Contract.RevokeRole(&_Order.TransactOpts, role, account)
}

// OrderBetEventIterator is returned from FilterBetEvent and is used to iterate over the raw logs and unpacked data for BetEvent events raised by the Order contract.
type OrderBetEventIterator struct {
	Event *OrderBetEvent // Event containing the contract specifics and raw log

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
func (it *OrderBetEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBetEvent)
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
		it.Event = new(OrderBetEvent)
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
func (it *OrderBetEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBetEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBetEvent represents a BetEvent event raised by the Order contract.
type OrderBetEvent struct {
	GameId  *big.Int
	AutoId  *big.Int
	OrderId *big.Int
	User    common.Address
	TokenId *big.Int
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBetEvent is a free log retrieval operation binding the contract event 0xda4a6dbd45fd2585cef08553979d11c0be03ef59aa3fb11fc72eb46a17d91c50.
//
// Solidity: event BetEvent(uint256 gameId, uint256 autoId, uint256 indexed orderId, address indexed user, uint256 tokenId, uint256 indexed amount, bytes data)
func (_Order *OrderFilterer) FilterBetEvent(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address, amount []*big.Int) (*OrderBetEventIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "BetEvent", orderIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBetEventIterator{contract: _Order.contract, event: "BetEvent", logs: logs, sub: sub}, nil
}

// WatchBetEvent is a free log subscription operation binding the contract event 0xda4a6dbd45fd2585cef08553979d11c0be03ef59aa3fb11fc72eb46a17d91c50.
//
// Solidity: event BetEvent(uint256 gameId, uint256 autoId, uint256 indexed orderId, address indexed user, uint256 tokenId, uint256 indexed amount, bytes data)
func (_Order *OrderFilterer) WatchBetEvent(opts *bind.WatchOpts, sink chan<- *OrderBetEvent, orderId []*big.Int, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "BetEvent", orderIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBetEvent)
				if err := _Order.contract.UnpackLog(event, "BetEvent", log); err != nil {
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

// ParseBetEvent is a log parse operation binding the contract event 0xda4a6dbd45fd2585cef08553979d11c0be03ef59aa3fb11fc72eb46a17d91c50.
//
// Solidity: event BetEvent(uint256 gameId, uint256 autoId, uint256 indexed orderId, address indexed user, uint256 tokenId, uint256 indexed amount, bytes data)
func (_Order *OrderFilterer) ParseBetEvent(log types.Log) (*OrderBetEvent, error) {
	event := new(OrderBetEvent)
	if err := _Order.contract.UnpackLog(event, "BetEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderDrawEventIterator is returned from FilterDrawEvent and is used to iterate over the raw logs and unpacked data for DrawEvent events raised by the Order contract.
type OrderDrawEventIterator struct {
	Event *OrderDrawEvent // Event containing the contract specifics and raw log

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
func (it *OrderDrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderDrawEvent)
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
		it.Event = new(OrderDrawEvent)
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
func (it *OrderDrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderDrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderDrawEvent represents a DrawEvent event raised by the Order contract.
type OrderDrawEvent struct {
	GameId  *big.Int
	OrderId *big.Int
	Status  *big.Int
	User    common.Address
	Bonus   *big.Int
	Refund  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDrawEvent is a free log retrieval operation binding the contract event 0x39a879a428a5382476f497ed2b6049a3a3a6a9588c2f5d10058391ee1207d972.
//
// Solidity: event DrawEvent(uint256 gameId, uint256 indexed orderId, uint256 indexed status, address indexed user, uint256 bonus, uint256 refund)
func (_Order *OrderFilterer) FilterDrawEvent(opts *bind.FilterOpts, orderId []*big.Int, status []*big.Int, user []common.Address) (*OrderDrawEventIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "DrawEvent", orderIdRule, statusRule, userRule)
	if err != nil {
		return nil, err
	}
	return &OrderDrawEventIterator{contract: _Order.contract, event: "DrawEvent", logs: logs, sub: sub}, nil
}

// WatchDrawEvent is a free log subscription operation binding the contract event 0x39a879a428a5382476f497ed2b6049a3a3a6a9588c2f5d10058391ee1207d972.
//
// Solidity: event DrawEvent(uint256 gameId, uint256 indexed orderId, uint256 indexed status, address indexed user, uint256 bonus, uint256 refund)
func (_Order *OrderFilterer) WatchDrawEvent(opts *bind.WatchOpts, sink chan<- *OrderDrawEvent, orderId []*big.Int, status []*big.Int, user []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "DrawEvent", orderIdRule, statusRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderDrawEvent)
				if err := _Order.contract.UnpackLog(event, "DrawEvent", log); err != nil {
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

// ParseDrawEvent is a log parse operation binding the contract event 0x39a879a428a5382476f497ed2b6049a3a3a6a9588c2f5d10058391ee1207d972.
//
// Solidity: event DrawEvent(uint256 gameId, uint256 indexed orderId, uint256 indexed status, address indexed user, uint256 bonus, uint256 refund)
func (_Order *OrderFilterer) ParseDrawEvent(log types.Log) (*OrderDrawEvent, error) {
	event := new(OrderDrawEvent)
	if err := _Order.contract.UnpackLog(event, "DrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Order contract.
type OrderInitializedIterator struct {
	Event *OrderInitialized // Event containing the contract specifics and raw log

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
func (it *OrderInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderInitialized)
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
		it.Event = new(OrderInitialized)
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
func (it *OrderInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderInitialized represents a Initialized event raised by the Order contract.
type OrderInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Order *OrderFilterer) FilterInitialized(opts *bind.FilterOpts) (*OrderInitializedIterator, error) {

	logs, sub, err := _Order.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OrderInitializedIterator{contract: _Order.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Order *OrderFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OrderInitialized) (event.Subscription, error) {

	logs, sub, err := _Order.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderInitialized)
				if err := _Order.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Order *OrderFilterer) ParseInitialized(log types.Log) (*OrderInitialized, error) {
	event := new(OrderInitialized)
	if err := _Order.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Order contract.
type OrderRoleAdminChangedIterator struct {
	Event *OrderRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OrderRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderRoleAdminChanged)
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
		it.Event = new(OrderRoleAdminChanged)
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
func (it *OrderRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderRoleAdminChanged represents a RoleAdminChanged event raised by the Order contract.
type OrderRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Order *OrderFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OrderRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OrderRoleAdminChangedIterator{contract: _Order.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Order *OrderFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OrderRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderRoleAdminChanged)
				if err := _Order.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Order *OrderFilterer) ParseRoleAdminChanged(log types.Log) (*OrderRoleAdminChanged, error) {
	event := new(OrderRoleAdminChanged)
	if err := _Order.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Order contract.
type OrderRoleGrantedIterator struct {
	Event *OrderRoleGranted // Event containing the contract specifics and raw log

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
func (it *OrderRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderRoleGranted)
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
		it.Event = new(OrderRoleGranted)
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
func (it *OrderRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderRoleGranted represents a RoleGranted event raised by the Order contract.
type OrderRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Order *OrderFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OrderRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OrderRoleGrantedIterator{contract: _Order.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Order *OrderFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OrderRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderRoleGranted)
				if err := _Order.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Order *OrderFilterer) ParseRoleGranted(log types.Log) (*OrderRoleGranted, error) {
	event := new(OrderRoleGranted)
	if err := _Order.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Order contract.
type OrderRoleRevokedIterator struct {
	Event *OrderRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OrderRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderRoleRevoked)
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
		it.Event = new(OrderRoleRevoked)
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
func (it *OrderRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderRoleRevoked represents a RoleRevoked event raised by the Order contract.
type OrderRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Order *OrderFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OrderRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OrderRoleRevokedIterator{contract: _Order.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Order *OrderFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OrderRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderRoleRevoked)
				if err := _Order.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Order *OrderFilterer) ParseRoleRevoked(log types.Log) (*OrderRoleRevoked, error) {
	event := new(OrderRoleRevoked)
	if err := _Order.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
