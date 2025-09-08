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

// IAutoBetAutoBetInfo is an auto generated low-level Go binding around an user-defined struct.
type IAutoBetAutoBetInfo struct {
	User        common.Address
	GameId      *big.Int
	WinLimit    *big.Int
	LossLimit   *big.Int
	Status      *big.Int
	TotalAmount *big.Int
	AddWhenWin  *big.Int
	AddWhenLoss *big.Int
	BetAmount   *big.Int
	Bonus       *big.Int
	WinAmount   *big.Int
	LossAmount  *big.Int
}

// AutoBetMetaData contains all meta data concerning the AutoBet contract.
var AutoBetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"autoBetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lossThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addWhenWin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"add_WhenLoss\",\"type\":\"uint256\"}],\"name\":\"AutoBetEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"autoBetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AutoBetReturnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addAutoBetOrderRealBetAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"addAutoOrderId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gameAddress_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"autoBetSwitch_\",\"type\":\"bool\"}],\"name\":\"admin_set_env\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoBetId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"autoBetOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lossLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"add_when_win\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"add_when_loss\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lossAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoBetId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"autoBetReturn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoBetSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"autoOrderIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"changeAutoBetOrderAmountAndStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"changeAutoBetOrderStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lossLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"add_when_win\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"add_when_loss\",\"type\":\"uint256\"}],\"name\":\"createAutoBetOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractIGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"}],\"name\":\"getAutoBetOrderAddAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"}],\"name\":\"getAutoBetOrderStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewAutoId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"}],\"name\":\"get_auto_bet_order_info\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lossLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"add_when_win\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"add_when_loss\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lossAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIAutoBet.AutoBetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"autoId\",\"type\":\"uint256\"}],\"name\":\"get_auto_order_ids\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AutoBetABI is the input ABI used to generate the binding from.
// Deprecated: Use AutoBetMetaData.ABI instead.
var AutoBetABI = AutoBetMetaData.ABI

// AutoBet is an auto generated Go binding around an Ethereum contract.
type AutoBet struct {
	AutoBetCaller     // Read-only binding to the contract
	AutoBetTransactor // Write-only binding to the contract
	AutoBetFilterer   // Log filterer for contract events
}

// AutoBetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AutoBetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutoBetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AutoBetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutoBetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AutoBetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutoBetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AutoBetSession struct {
	Contract     *AutoBet          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AutoBetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AutoBetCallerSession struct {
	Contract *AutoBetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AutoBetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AutoBetTransactorSession struct {
	Contract     *AutoBetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AutoBetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AutoBetRaw struct {
	Contract *AutoBet // Generic contract binding to access the raw methods on
}

// AutoBetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AutoBetCallerRaw struct {
	Contract *AutoBetCaller // Generic read-only contract binding to access the raw methods on
}

// AutoBetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AutoBetTransactorRaw struct {
	Contract *AutoBetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAutoBet creates a new instance of AutoBet, bound to a specific deployed contract.
func NewAutoBet(address common.Address, backend bind.ContractBackend) (*AutoBet, error) {
	contract, err := bindAutoBet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AutoBet{AutoBetCaller: AutoBetCaller{contract: contract}, AutoBetTransactor: AutoBetTransactor{contract: contract}, AutoBetFilterer: AutoBetFilterer{contract: contract}}, nil
}

// NewAutoBetCaller creates a new read-only instance of AutoBet, bound to a specific deployed contract.
func NewAutoBetCaller(address common.Address, caller bind.ContractCaller) (*AutoBetCaller, error) {
	contract, err := bindAutoBet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutoBetCaller{contract: contract}, nil
}

// NewAutoBetTransactor creates a new write-only instance of AutoBet, bound to a specific deployed contract.
func NewAutoBetTransactor(address common.Address, transactor bind.ContractTransactor) (*AutoBetTransactor, error) {
	contract, err := bindAutoBet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutoBetTransactor{contract: contract}, nil
}

// NewAutoBetFilterer creates a new log filterer instance of AutoBet, bound to a specific deployed contract.
func NewAutoBetFilterer(address common.Address, filterer bind.ContractFilterer) (*AutoBetFilterer, error) {
	contract, err := bindAutoBet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutoBetFilterer{contract: contract}, nil
}

// bindAutoBet binds a generic wrapper to an already deployed contract.
func bindAutoBet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AutoBetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutoBet *AutoBetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutoBet.Contract.AutoBetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutoBet *AutoBetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutoBet.Contract.AutoBetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutoBet *AutoBetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutoBet.Contract.AutoBetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutoBet *AutoBetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutoBet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutoBet *AutoBetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutoBet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutoBet *AutoBetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutoBet.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetSession) ADMINROLE() ([32]byte, error) {
	return _AutoBet.Contract.ADMINROLE(&_AutoBet.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetCallerSession) ADMINROLE() ([32]byte, error) {
	return _AutoBet.Contract.ADMINROLE(&_AutoBet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AutoBet.Contract.DEFAULTADMINROLE(&_AutoBet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AutoBet.Contract.DEFAULTADMINROLE(&_AutoBet.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetCaller) SERVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "SERVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetSession) SERVERROLE() ([32]byte, error) {
	return _AutoBet.Contract.SERVERROLE(&_AutoBet.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_AutoBet *AutoBetCallerSession) SERVERROLE() ([32]byte, error) {
	return _AutoBet.Contract.SERVERROLE(&_AutoBet.CallOpts)
}

// AutoBetId is a free data retrieval call binding the contract method 0x481efe78.
//
// Solidity: function autoBetId() view returns(uint256)
func (_AutoBet *AutoBetCaller) AutoBetId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "autoBetId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoBetId is a free data retrieval call binding the contract method 0x481efe78.
//
// Solidity: function autoBetId() view returns(uint256)
func (_AutoBet *AutoBetSession) AutoBetId() (*big.Int, error) {
	return _AutoBet.Contract.AutoBetId(&_AutoBet.CallOpts)
}

// AutoBetId is a free data retrieval call binding the contract method 0x481efe78.
//
// Solidity: function autoBetId() view returns(uint256)
func (_AutoBet *AutoBetCallerSession) AutoBetId() (*big.Int, error) {
	return _AutoBet.Contract.AutoBetId(&_AutoBet.CallOpts)
}

// AutoBetOrder is a free data retrieval call binding the contract method 0xf9a7b86e.
//
// Solidity: function autoBetOrder(uint256 ) view returns(address user, uint256 gameId, uint256 winLimit, uint256 lossLimit, uint256 status, uint256 totalAmount, uint256 add_when_win, uint256 add_when_loss, uint256 betAmount, uint256 bonus, uint256 winAmount, uint256 lossAmount)
func (_AutoBet *AutoBetCaller) AutoBetOrder(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User        common.Address
	GameId      *big.Int
	WinLimit    *big.Int
	LossLimit   *big.Int
	Status      *big.Int
	TotalAmount *big.Int
	AddWhenWin  *big.Int
	AddWhenLoss *big.Int
	BetAmount   *big.Int
	Bonus       *big.Int
	WinAmount   *big.Int
	LossAmount  *big.Int
}, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "autoBetOrder", arg0)

	outstruct := new(struct {
		User        common.Address
		GameId      *big.Int
		WinLimit    *big.Int
		LossLimit   *big.Int
		Status      *big.Int
		TotalAmount *big.Int
		AddWhenWin  *big.Int
		AddWhenLoss *big.Int
		BetAmount   *big.Int
		Bonus       *big.Int
		WinAmount   *big.Int
		LossAmount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GameId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WinLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LossLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AddWhenWin = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AddWhenLoss = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.BetAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Bonus = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.WinAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.LossAmount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AutoBetOrder is a free data retrieval call binding the contract method 0xf9a7b86e.
//
// Solidity: function autoBetOrder(uint256 ) view returns(address user, uint256 gameId, uint256 winLimit, uint256 lossLimit, uint256 status, uint256 totalAmount, uint256 add_when_win, uint256 add_when_loss, uint256 betAmount, uint256 bonus, uint256 winAmount, uint256 lossAmount)
func (_AutoBet *AutoBetSession) AutoBetOrder(arg0 *big.Int) (struct {
	User        common.Address
	GameId      *big.Int
	WinLimit    *big.Int
	LossLimit   *big.Int
	Status      *big.Int
	TotalAmount *big.Int
	AddWhenWin  *big.Int
	AddWhenLoss *big.Int
	BetAmount   *big.Int
	Bonus       *big.Int
	WinAmount   *big.Int
	LossAmount  *big.Int
}, error) {
	return _AutoBet.Contract.AutoBetOrder(&_AutoBet.CallOpts, arg0)
}

// AutoBetOrder is a free data retrieval call binding the contract method 0xf9a7b86e.
//
// Solidity: function autoBetOrder(uint256 ) view returns(address user, uint256 gameId, uint256 winLimit, uint256 lossLimit, uint256 status, uint256 totalAmount, uint256 add_when_win, uint256 add_when_loss, uint256 betAmount, uint256 bonus, uint256 winAmount, uint256 lossAmount)
func (_AutoBet *AutoBetCallerSession) AutoBetOrder(arg0 *big.Int) (struct {
	User        common.Address
	GameId      *big.Int
	WinLimit    *big.Int
	LossLimit   *big.Int
	Status      *big.Int
	TotalAmount *big.Int
	AddWhenWin  *big.Int
	AddWhenLoss *big.Int
	BetAmount   *big.Int
	Bonus       *big.Int
	WinAmount   *big.Int
	LossAmount  *big.Int
}, error) {
	return _AutoBet.Contract.AutoBetOrder(&_AutoBet.CallOpts, arg0)
}

// AutoBetSwitch is a free data retrieval call binding the contract method 0x250f9fb8.
//
// Solidity: function autoBetSwitch() view returns(bool)
func (_AutoBet *AutoBetCaller) AutoBetSwitch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "autoBetSwitch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutoBetSwitch is a free data retrieval call binding the contract method 0x250f9fb8.
//
// Solidity: function autoBetSwitch() view returns(bool)
func (_AutoBet *AutoBetSession) AutoBetSwitch() (bool, error) {
	return _AutoBet.Contract.AutoBetSwitch(&_AutoBet.CallOpts)
}

// AutoBetSwitch is a free data retrieval call binding the contract method 0x250f9fb8.
//
// Solidity: function autoBetSwitch() view returns(bool)
func (_AutoBet *AutoBetCallerSession) AutoBetSwitch() (bool, error) {
	return _AutoBet.Contract.AutoBetSwitch(&_AutoBet.CallOpts)
}

// AutoOrderIds is a free data retrieval call binding the contract method 0x5351367a.
//
// Solidity: function autoOrderIds(uint256 , uint256 ) view returns(uint256)
func (_AutoBet *AutoBetCaller) AutoOrderIds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "autoOrderIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoOrderIds is a free data retrieval call binding the contract method 0x5351367a.
//
// Solidity: function autoOrderIds(uint256 , uint256 ) view returns(uint256)
func (_AutoBet *AutoBetSession) AutoOrderIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AutoBet.Contract.AutoOrderIds(&_AutoBet.CallOpts, arg0, arg1)
}

// AutoOrderIds is a free data retrieval call binding the contract method 0x5351367a.
//
// Solidity: function autoOrderIds(uint256 , uint256 ) view returns(uint256)
func (_AutoBet *AutoBetCallerSession) AutoOrderIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AutoBet.Contract.AutoOrderIds(&_AutoBet.CallOpts, arg0, arg1)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_AutoBet *AutoBetCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_AutoBet *AutoBetSession) Game() (common.Address, error) {
	return _AutoBet.Contract.Game(&_AutoBet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_AutoBet *AutoBetCallerSession) Game() (common.Address, error) {
	return _AutoBet.Contract.Game(&_AutoBet.CallOpts)
}

// GetAutoBetOrderAddAmount is a free data retrieval call binding the contract method 0x2f37d991.
//
// Solidity: function getAutoBetOrderAddAmount(uint256 autoId) view returns(uint256, uint256)
func (_AutoBet *AutoBetCaller) GetAutoBetOrderAddAmount(opts *bind.CallOpts, autoId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "getAutoBetOrderAddAmount", autoId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAutoBetOrderAddAmount is a free data retrieval call binding the contract method 0x2f37d991.
//
// Solidity: function getAutoBetOrderAddAmount(uint256 autoId) view returns(uint256, uint256)
func (_AutoBet *AutoBetSession) GetAutoBetOrderAddAmount(autoId *big.Int) (*big.Int, *big.Int, error) {
	return _AutoBet.Contract.GetAutoBetOrderAddAmount(&_AutoBet.CallOpts, autoId)
}

// GetAutoBetOrderAddAmount is a free data retrieval call binding the contract method 0x2f37d991.
//
// Solidity: function getAutoBetOrderAddAmount(uint256 autoId) view returns(uint256, uint256)
func (_AutoBet *AutoBetCallerSession) GetAutoBetOrderAddAmount(autoId *big.Int) (*big.Int, *big.Int, error) {
	return _AutoBet.Contract.GetAutoBetOrderAddAmount(&_AutoBet.CallOpts, autoId)
}

// GetAutoBetOrderStatus is a free data retrieval call binding the contract method 0x24866e0b.
//
// Solidity: function getAutoBetOrderStatus(uint256 autoId) view returns(uint256)
func (_AutoBet *AutoBetCaller) GetAutoBetOrderStatus(opts *bind.CallOpts, autoId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "getAutoBetOrderStatus", autoId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAutoBetOrderStatus is a free data retrieval call binding the contract method 0x24866e0b.
//
// Solidity: function getAutoBetOrderStatus(uint256 autoId) view returns(uint256)
func (_AutoBet *AutoBetSession) GetAutoBetOrderStatus(autoId *big.Int) (*big.Int, error) {
	return _AutoBet.Contract.GetAutoBetOrderStatus(&_AutoBet.CallOpts, autoId)
}

// GetAutoBetOrderStatus is a free data retrieval call binding the contract method 0x24866e0b.
//
// Solidity: function getAutoBetOrderStatus(uint256 autoId) view returns(uint256)
func (_AutoBet *AutoBetCallerSession) GetAutoBetOrderStatus(autoId *big.Int) (*big.Int, error) {
	return _AutoBet.Contract.GetAutoBetOrderStatus(&_AutoBet.CallOpts, autoId)
}

// GetNewAutoId is a free data retrieval call binding the contract method 0x90634fa5.
//
// Solidity: function getNewAutoId() view returns(uint256)
func (_AutoBet *AutoBetCaller) GetNewAutoId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "getNewAutoId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewAutoId is a free data retrieval call binding the contract method 0x90634fa5.
//
// Solidity: function getNewAutoId() view returns(uint256)
func (_AutoBet *AutoBetSession) GetNewAutoId() (*big.Int, error) {
	return _AutoBet.Contract.GetNewAutoId(&_AutoBet.CallOpts)
}

// GetNewAutoId is a free data retrieval call binding the contract method 0x90634fa5.
//
// Solidity: function getNewAutoId() view returns(uint256)
func (_AutoBet *AutoBetCallerSession) GetNewAutoId() (*big.Int, error) {
	return _AutoBet.Contract.GetNewAutoId(&_AutoBet.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AutoBet *AutoBetCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AutoBet *AutoBetSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AutoBet.Contract.GetRoleAdmin(&_AutoBet.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AutoBet *AutoBetCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AutoBet.Contract.GetRoleAdmin(&_AutoBet.CallOpts, role)
}

// GetAutoBetOrderInfo is a free data retrieval call binding the contract method 0xfb66998c.
//
// Solidity: function get_auto_bet_order_info(uint256 autoId) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_AutoBet *AutoBetCaller) GetAutoBetOrderInfo(opts *bind.CallOpts, autoId *big.Int) (IAutoBetAutoBetInfo, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "get_auto_bet_order_info", autoId)

	if err != nil {
		return *new(IAutoBetAutoBetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IAutoBetAutoBetInfo)).(*IAutoBetAutoBetInfo)

	return out0, err

}

// GetAutoBetOrderInfo is a free data retrieval call binding the contract method 0xfb66998c.
//
// Solidity: function get_auto_bet_order_info(uint256 autoId) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_AutoBet *AutoBetSession) GetAutoBetOrderInfo(autoId *big.Int) (IAutoBetAutoBetInfo, error) {
	return _AutoBet.Contract.GetAutoBetOrderInfo(&_AutoBet.CallOpts, autoId)
}

// GetAutoBetOrderInfo is a free data retrieval call binding the contract method 0xfb66998c.
//
// Solidity: function get_auto_bet_order_info(uint256 autoId) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_AutoBet *AutoBetCallerSession) GetAutoBetOrderInfo(autoId *big.Int) (IAutoBetAutoBetInfo, error) {
	return _AutoBet.Contract.GetAutoBetOrderInfo(&_AutoBet.CallOpts, autoId)
}

// GetAutoOrderIds is a free data retrieval call binding the contract method 0x2974c07e.
//
// Solidity: function get_auto_order_ids(uint256 autoId) view returns(uint256[])
func (_AutoBet *AutoBetCaller) GetAutoOrderIds(opts *bind.CallOpts, autoId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "get_auto_order_ids", autoId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAutoOrderIds is a free data retrieval call binding the contract method 0x2974c07e.
//
// Solidity: function get_auto_order_ids(uint256 autoId) view returns(uint256[])
func (_AutoBet *AutoBetSession) GetAutoOrderIds(autoId *big.Int) ([]*big.Int, error) {
	return _AutoBet.Contract.GetAutoOrderIds(&_AutoBet.CallOpts, autoId)
}

// GetAutoOrderIds is a free data retrieval call binding the contract method 0x2974c07e.
//
// Solidity: function get_auto_order_ids(uint256 autoId) view returns(uint256[])
func (_AutoBet *AutoBetCallerSession) GetAutoOrderIds(autoId *big.Int) ([]*big.Int, error) {
	return _AutoBet.Contract.GetAutoOrderIds(&_AutoBet.CallOpts, autoId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AutoBet *AutoBetCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AutoBet *AutoBetSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AutoBet.Contract.HasRole(&_AutoBet.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AutoBet *AutoBetCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AutoBet.Contract.HasRole(&_AutoBet.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AutoBet *AutoBetCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AutoBet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AutoBet *AutoBetSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AutoBet.Contract.SupportsInterface(&_AutoBet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AutoBet *AutoBetCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AutoBet.Contract.SupportsInterface(&_AutoBet.CallOpts, interfaceId)
}

// AddAutoBetOrderRealBetAmount is a paid mutator transaction binding the contract method 0xa7a1548f.
//
// Solidity: function addAutoBetOrderRealBetAmount(uint256 autoId, uint256 amount) returns()
func (_AutoBet *AutoBetTransactor) AddAutoBetOrderRealBetAmount(opts *bind.TransactOpts, autoId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "addAutoBetOrderRealBetAmount", autoId, amount)
}

// AddAutoBetOrderRealBetAmount is a paid mutator transaction binding the contract method 0xa7a1548f.
//
// Solidity: function addAutoBetOrderRealBetAmount(uint256 autoId, uint256 amount) returns()
func (_AutoBet *AutoBetSession) AddAutoBetOrderRealBetAmount(autoId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.AddAutoBetOrderRealBetAmount(&_AutoBet.TransactOpts, autoId, amount)
}

// AddAutoBetOrderRealBetAmount is a paid mutator transaction binding the contract method 0xa7a1548f.
//
// Solidity: function addAutoBetOrderRealBetAmount(uint256 autoId, uint256 amount) returns()
func (_AutoBet *AutoBetTransactorSession) AddAutoBetOrderRealBetAmount(autoId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.AddAutoBetOrderRealBetAmount(&_AutoBet.TransactOpts, autoId, amount)
}

// AddAutoOrderId is a paid mutator transaction binding the contract method 0x03dcff45.
//
// Solidity: function addAutoOrderId(uint256 autoId, uint256 orderId) returns()
func (_AutoBet *AutoBetTransactor) AddAutoOrderId(opts *bind.TransactOpts, autoId *big.Int, orderId *big.Int) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "addAutoOrderId", autoId, orderId)
}

// AddAutoOrderId is a paid mutator transaction binding the contract method 0x03dcff45.
//
// Solidity: function addAutoOrderId(uint256 autoId, uint256 orderId) returns()
func (_AutoBet *AutoBetSession) AddAutoOrderId(autoId *big.Int, orderId *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.AddAutoOrderId(&_AutoBet.TransactOpts, autoId, orderId)
}

// AddAutoOrderId is a paid mutator transaction binding the contract method 0x03dcff45.
//
// Solidity: function addAutoOrderId(uint256 autoId, uint256 orderId) returns()
func (_AutoBet *AutoBetTransactorSession) AddAutoOrderId(autoId *big.Int, orderId *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.AddAutoOrderId(&_AutoBet.TransactOpts, autoId, orderId)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x163b9e65.
//
// Solidity: function admin_set_env(address gameAddress_, bool autoBetSwitch_) returns()
func (_AutoBet *AutoBetTransactor) AdminSetEnv(opts *bind.TransactOpts, gameAddress_ common.Address, autoBetSwitch_ bool) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "admin_set_env", gameAddress_, autoBetSwitch_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x163b9e65.
//
// Solidity: function admin_set_env(address gameAddress_, bool autoBetSwitch_) returns()
func (_AutoBet *AutoBetSession) AdminSetEnv(gameAddress_ common.Address, autoBetSwitch_ bool) (*types.Transaction, error) {
	return _AutoBet.Contract.AdminSetEnv(&_AutoBet.TransactOpts, gameAddress_, autoBetSwitch_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x163b9e65.
//
// Solidity: function admin_set_env(address gameAddress_, bool autoBetSwitch_) returns()
func (_AutoBet *AutoBetTransactorSession) AdminSetEnv(gameAddress_ common.Address, autoBetSwitch_ bool) (*types.Transaction, error) {
	return _AutoBet.Contract.AdminSetEnv(&_AutoBet.TransactOpts, gameAddress_, autoBetSwitch_)
}

// AutoBetReturn is a paid mutator transaction binding the contract method 0xb037b051.
//
// Solidity: function autoBetReturn(uint256 autoBetId_, uint256 gameId_, uint256 tokenId_) returns()
func (_AutoBet *AutoBetTransactor) AutoBetReturn(opts *bind.TransactOpts, autoBetId_ *big.Int, gameId_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "autoBetReturn", autoBetId_, gameId_, tokenId_)
}

// AutoBetReturn is a paid mutator transaction binding the contract method 0xb037b051.
//
// Solidity: function autoBetReturn(uint256 autoBetId_, uint256 gameId_, uint256 tokenId_) returns()
func (_AutoBet *AutoBetSession) AutoBetReturn(autoBetId_ *big.Int, gameId_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.AutoBetReturn(&_AutoBet.TransactOpts, autoBetId_, gameId_, tokenId_)
}

// AutoBetReturn is a paid mutator transaction binding the contract method 0xb037b051.
//
// Solidity: function autoBetReturn(uint256 autoBetId_, uint256 gameId_, uint256 tokenId_) returns()
func (_AutoBet *AutoBetTransactorSession) AutoBetReturn(autoBetId_ *big.Int, gameId_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.AutoBetReturn(&_AutoBet.TransactOpts, autoBetId_, gameId_, tokenId_)
}

// ChangeAutoBetOrderAmountAndStatus is a paid mutator transaction binding the contract method 0xa8e00b60.
//
// Solidity: function changeAutoBetOrderAmountAndStatus(uint256 autoId, bool isWin, uint256 tokenId, uint256 amount) returns()
func (_AutoBet *AutoBetTransactor) ChangeAutoBetOrderAmountAndStatus(opts *bind.TransactOpts, autoId *big.Int, isWin bool, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "changeAutoBetOrderAmountAndStatus", autoId, isWin, tokenId, amount)
}

// ChangeAutoBetOrderAmountAndStatus is a paid mutator transaction binding the contract method 0xa8e00b60.
//
// Solidity: function changeAutoBetOrderAmountAndStatus(uint256 autoId, bool isWin, uint256 tokenId, uint256 amount) returns()
func (_AutoBet *AutoBetSession) ChangeAutoBetOrderAmountAndStatus(autoId *big.Int, isWin bool, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.ChangeAutoBetOrderAmountAndStatus(&_AutoBet.TransactOpts, autoId, isWin, tokenId, amount)
}

// ChangeAutoBetOrderAmountAndStatus is a paid mutator transaction binding the contract method 0xa8e00b60.
//
// Solidity: function changeAutoBetOrderAmountAndStatus(uint256 autoId, bool isWin, uint256 tokenId, uint256 amount) returns()
func (_AutoBet *AutoBetTransactorSession) ChangeAutoBetOrderAmountAndStatus(autoId *big.Int, isWin bool, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.ChangeAutoBetOrderAmountAndStatus(&_AutoBet.TransactOpts, autoId, isWin, tokenId, amount)
}

// ChangeAutoBetOrderStatus is a paid mutator transaction binding the contract method 0x127cbbb6.
//
// Solidity: function changeAutoBetOrderStatus(uint256 autoId, uint256 status) returns()
func (_AutoBet *AutoBetTransactor) ChangeAutoBetOrderStatus(opts *bind.TransactOpts, autoId *big.Int, status *big.Int) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "changeAutoBetOrderStatus", autoId, status)
}

// ChangeAutoBetOrderStatus is a paid mutator transaction binding the contract method 0x127cbbb6.
//
// Solidity: function changeAutoBetOrderStatus(uint256 autoId, uint256 status) returns()
func (_AutoBet *AutoBetSession) ChangeAutoBetOrderStatus(autoId *big.Int, status *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.ChangeAutoBetOrderStatus(&_AutoBet.TransactOpts, autoId, status)
}

// ChangeAutoBetOrderStatus is a paid mutator transaction binding the contract method 0x127cbbb6.
//
// Solidity: function changeAutoBetOrderStatus(uint256 autoId, uint256 status) returns()
func (_AutoBet *AutoBetTransactorSession) ChangeAutoBetOrderStatus(autoId *big.Int, status *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.ChangeAutoBetOrderStatus(&_AutoBet.TransactOpts, autoId, status)
}

// CreateAutoBetOrder is a paid mutator transaction binding the contract method 0xf5a8a463.
//
// Solidity: function createAutoBetOrder(uint256 autoId, address user, uint256 gameId, uint256 winLimit, uint256 lossLimit, uint256 totalAmount, uint256 add_when_win, uint256 add_when_loss) returns()
func (_AutoBet *AutoBetTransactor) CreateAutoBetOrder(opts *bind.TransactOpts, autoId *big.Int, user common.Address, gameId *big.Int, winLimit *big.Int, lossLimit *big.Int, totalAmount *big.Int, add_when_win *big.Int, add_when_loss *big.Int) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "createAutoBetOrder", autoId, user, gameId, winLimit, lossLimit, totalAmount, add_when_win, add_when_loss)
}

// CreateAutoBetOrder is a paid mutator transaction binding the contract method 0xf5a8a463.
//
// Solidity: function createAutoBetOrder(uint256 autoId, address user, uint256 gameId, uint256 winLimit, uint256 lossLimit, uint256 totalAmount, uint256 add_when_win, uint256 add_when_loss) returns()
func (_AutoBet *AutoBetSession) CreateAutoBetOrder(autoId *big.Int, user common.Address, gameId *big.Int, winLimit *big.Int, lossLimit *big.Int, totalAmount *big.Int, add_when_win *big.Int, add_when_loss *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.CreateAutoBetOrder(&_AutoBet.TransactOpts, autoId, user, gameId, winLimit, lossLimit, totalAmount, add_when_win, add_when_loss)
}

// CreateAutoBetOrder is a paid mutator transaction binding the contract method 0xf5a8a463.
//
// Solidity: function createAutoBetOrder(uint256 autoId, address user, uint256 gameId, uint256 winLimit, uint256 lossLimit, uint256 totalAmount, uint256 add_when_win, uint256 add_when_loss) returns()
func (_AutoBet *AutoBetTransactorSession) CreateAutoBetOrder(autoId *big.Int, user common.Address, gameId *big.Int, winLimit *big.Int, lossLimit *big.Int, totalAmount *big.Int, add_when_win *big.Int, add_when_loss *big.Int) (*types.Transaction, error) {
	return _AutoBet.Contract.CreateAutoBetOrder(&_AutoBet.TransactOpts, autoId, user, gameId, winLimit, lossLimit, totalAmount, add_when_win, add_when_loss)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AutoBet *AutoBetTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AutoBet *AutoBetSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AutoBet.Contract.GrantRole(&_AutoBet.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AutoBet *AutoBetTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AutoBet.Contract.GrantRole(&_AutoBet.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AutoBet *AutoBetTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AutoBet *AutoBetSession) Initialize() (*types.Transaction, error) {
	return _AutoBet.Contract.Initialize(&_AutoBet.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AutoBet *AutoBetTransactorSession) Initialize() (*types.Transaction, error) {
	return _AutoBet.Contract.Initialize(&_AutoBet.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AutoBet *AutoBetTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AutoBet *AutoBetSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AutoBet.Contract.RenounceRole(&_AutoBet.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AutoBet *AutoBetTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AutoBet.Contract.RenounceRole(&_AutoBet.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AutoBet *AutoBetTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AutoBet.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AutoBet *AutoBetSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AutoBet.Contract.RevokeRole(&_AutoBet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AutoBet *AutoBetTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AutoBet.Contract.RevokeRole(&_AutoBet.TransactOpts, role, account)
}

// AutoBetAutoBetEventIterator is returned from FilterAutoBetEvent and is used to iterate over the raw logs and unpacked data for AutoBetEvent events raised by the AutoBet contract.
type AutoBetAutoBetEventIterator struct {
	Event *AutoBetAutoBetEvent // Event containing the contract specifics and raw log

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
func (it *AutoBetAutoBetEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutoBetAutoBetEvent)
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
		it.Event = new(AutoBetAutoBetEvent)
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
func (it *AutoBetAutoBetEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutoBetAutoBetEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutoBetAutoBetEvent represents a AutoBetEvent event raised by the AutoBet contract.
type AutoBetAutoBetEvent struct {
	AutoBetId     *big.Int
	GameId        *big.Int
	User          common.Address
	TotalAmount   *big.Int
	WinThreshold  *big.Int
	LossThreshold *big.Int
	AddWhenWin    *big.Int
	AddWhenLoss   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAutoBetEvent is a free log retrieval operation binding the contract event 0x4bbbda278af66e4f8dfe798606c185261a5597729b1bd9c77449fbdae16f3458.
//
// Solidity: event AutoBetEvent(uint256 indexed autoBetId, uint256 indexed gameId, address indexed user, uint256 totalAmount, uint256 winThreshold, uint256 lossThreshold, uint256 addWhenWin, uint256 add_WhenLoss)
func (_AutoBet *AutoBetFilterer) FilterAutoBetEvent(opts *bind.FilterOpts, autoBetId []*big.Int, gameId []*big.Int, user []common.Address) (*AutoBetAutoBetEventIterator, error) {

	var autoBetIdRule []interface{}
	for _, autoBetIdItem := range autoBetId {
		autoBetIdRule = append(autoBetIdRule, autoBetIdItem)
	}
	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AutoBet.contract.FilterLogs(opts, "AutoBetEvent", autoBetIdRule, gameIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AutoBetAutoBetEventIterator{contract: _AutoBet.contract, event: "AutoBetEvent", logs: logs, sub: sub}, nil
}

// WatchAutoBetEvent is a free log subscription operation binding the contract event 0x4bbbda278af66e4f8dfe798606c185261a5597729b1bd9c77449fbdae16f3458.
//
// Solidity: event AutoBetEvent(uint256 indexed autoBetId, uint256 indexed gameId, address indexed user, uint256 totalAmount, uint256 winThreshold, uint256 lossThreshold, uint256 addWhenWin, uint256 add_WhenLoss)
func (_AutoBet *AutoBetFilterer) WatchAutoBetEvent(opts *bind.WatchOpts, sink chan<- *AutoBetAutoBetEvent, autoBetId []*big.Int, gameId []*big.Int, user []common.Address) (event.Subscription, error) {

	var autoBetIdRule []interface{}
	for _, autoBetIdItem := range autoBetId {
		autoBetIdRule = append(autoBetIdRule, autoBetIdItem)
	}
	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AutoBet.contract.WatchLogs(opts, "AutoBetEvent", autoBetIdRule, gameIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutoBetAutoBetEvent)
				if err := _AutoBet.contract.UnpackLog(event, "AutoBetEvent", log); err != nil {
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

// ParseAutoBetEvent is a log parse operation binding the contract event 0x4bbbda278af66e4f8dfe798606c185261a5597729b1bd9c77449fbdae16f3458.
//
// Solidity: event AutoBetEvent(uint256 indexed autoBetId, uint256 indexed gameId, address indexed user, uint256 totalAmount, uint256 winThreshold, uint256 lossThreshold, uint256 addWhenWin, uint256 add_WhenLoss)
func (_AutoBet *AutoBetFilterer) ParseAutoBetEvent(log types.Log) (*AutoBetAutoBetEvent, error) {
	event := new(AutoBetAutoBetEvent)
	if err := _AutoBet.contract.UnpackLog(event, "AutoBetEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutoBetAutoBetReturnEventIterator is returned from FilterAutoBetReturnEvent and is used to iterate over the raw logs and unpacked data for AutoBetReturnEvent events raised by the AutoBet contract.
type AutoBetAutoBetReturnEventIterator struct {
	Event *AutoBetAutoBetReturnEvent // Event containing the contract specifics and raw log

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
func (it *AutoBetAutoBetReturnEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutoBetAutoBetReturnEvent)
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
		it.Event = new(AutoBetAutoBetReturnEvent)
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
func (it *AutoBetAutoBetReturnEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutoBetAutoBetReturnEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutoBetAutoBetReturnEvent represents a AutoBetReturnEvent event raised by the AutoBet contract.
type AutoBetAutoBetReturnEvent struct {
	AutoBetId *big.Int
	GameId    *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAutoBetReturnEvent is a free log retrieval operation binding the contract event 0xfe07506783c55fc766108fa91e8977f33a80aac492ac1e28bf4e17d184102050.
//
// Solidity: event AutoBetReturnEvent(uint256 indexed autoBetId, uint256 indexed gameId, uint256 amount)
func (_AutoBet *AutoBetFilterer) FilterAutoBetReturnEvent(opts *bind.FilterOpts, autoBetId []*big.Int, gameId []*big.Int) (*AutoBetAutoBetReturnEventIterator, error) {

	var autoBetIdRule []interface{}
	for _, autoBetIdItem := range autoBetId {
		autoBetIdRule = append(autoBetIdRule, autoBetIdItem)
	}
	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}

	logs, sub, err := _AutoBet.contract.FilterLogs(opts, "AutoBetReturnEvent", autoBetIdRule, gameIdRule)
	if err != nil {
		return nil, err
	}
	return &AutoBetAutoBetReturnEventIterator{contract: _AutoBet.contract, event: "AutoBetReturnEvent", logs: logs, sub: sub}, nil
}

// WatchAutoBetReturnEvent is a free log subscription operation binding the contract event 0xfe07506783c55fc766108fa91e8977f33a80aac492ac1e28bf4e17d184102050.
//
// Solidity: event AutoBetReturnEvent(uint256 indexed autoBetId, uint256 indexed gameId, uint256 amount)
func (_AutoBet *AutoBetFilterer) WatchAutoBetReturnEvent(opts *bind.WatchOpts, sink chan<- *AutoBetAutoBetReturnEvent, autoBetId []*big.Int, gameId []*big.Int) (event.Subscription, error) {

	var autoBetIdRule []interface{}
	for _, autoBetIdItem := range autoBetId {
		autoBetIdRule = append(autoBetIdRule, autoBetIdItem)
	}
	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}

	logs, sub, err := _AutoBet.contract.WatchLogs(opts, "AutoBetReturnEvent", autoBetIdRule, gameIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutoBetAutoBetReturnEvent)
				if err := _AutoBet.contract.UnpackLog(event, "AutoBetReturnEvent", log); err != nil {
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

// ParseAutoBetReturnEvent is a log parse operation binding the contract event 0xfe07506783c55fc766108fa91e8977f33a80aac492ac1e28bf4e17d184102050.
//
// Solidity: event AutoBetReturnEvent(uint256 indexed autoBetId, uint256 indexed gameId, uint256 amount)
func (_AutoBet *AutoBetFilterer) ParseAutoBetReturnEvent(log types.Log) (*AutoBetAutoBetReturnEvent, error) {
	event := new(AutoBetAutoBetReturnEvent)
	if err := _AutoBet.contract.UnpackLog(event, "AutoBetReturnEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutoBetInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AutoBet contract.
type AutoBetInitializedIterator struct {
	Event *AutoBetInitialized // Event containing the contract specifics and raw log

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
func (it *AutoBetInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutoBetInitialized)
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
		it.Event = new(AutoBetInitialized)
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
func (it *AutoBetInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutoBetInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutoBetInitialized represents a Initialized event raised by the AutoBet contract.
type AutoBetInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AutoBet *AutoBetFilterer) FilterInitialized(opts *bind.FilterOpts) (*AutoBetInitializedIterator, error) {

	logs, sub, err := _AutoBet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AutoBetInitializedIterator{contract: _AutoBet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AutoBet *AutoBetFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AutoBetInitialized) (event.Subscription, error) {

	logs, sub, err := _AutoBet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutoBetInitialized)
				if err := _AutoBet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AutoBet *AutoBetFilterer) ParseInitialized(log types.Log) (*AutoBetInitialized, error) {
	event := new(AutoBetInitialized)
	if err := _AutoBet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutoBetRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AutoBet contract.
type AutoBetRoleAdminChangedIterator struct {
	Event *AutoBetRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AutoBetRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutoBetRoleAdminChanged)
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
		it.Event = new(AutoBetRoleAdminChanged)
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
func (it *AutoBetRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutoBetRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutoBetRoleAdminChanged represents a RoleAdminChanged event raised by the AutoBet contract.
type AutoBetRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AutoBet *AutoBetFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AutoBetRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AutoBet.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AutoBetRoleAdminChangedIterator{contract: _AutoBet.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AutoBet *AutoBetFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AutoBetRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AutoBet.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutoBetRoleAdminChanged)
				if err := _AutoBet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AutoBet *AutoBetFilterer) ParseRoleAdminChanged(log types.Log) (*AutoBetRoleAdminChanged, error) {
	event := new(AutoBetRoleAdminChanged)
	if err := _AutoBet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutoBetRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AutoBet contract.
type AutoBetRoleGrantedIterator struct {
	Event *AutoBetRoleGranted // Event containing the contract specifics and raw log

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
func (it *AutoBetRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutoBetRoleGranted)
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
		it.Event = new(AutoBetRoleGranted)
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
func (it *AutoBetRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutoBetRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutoBetRoleGranted represents a RoleGranted event raised by the AutoBet contract.
type AutoBetRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AutoBet *AutoBetFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AutoBetRoleGrantedIterator, error) {

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

	logs, sub, err := _AutoBet.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AutoBetRoleGrantedIterator{contract: _AutoBet.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AutoBet *AutoBetFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AutoBetRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AutoBet.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutoBetRoleGranted)
				if err := _AutoBet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AutoBet *AutoBetFilterer) ParseRoleGranted(log types.Log) (*AutoBetRoleGranted, error) {
	event := new(AutoBetRoleGranted)
	if err := _AutoBet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutoBetRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AutoBet contract.
type AutoBetRoleRevokedIterator struct {
	Event *AutoBetRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AutoBetRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutoBetRoleRevoked)
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
		it.Event = new(AutoBetRoleRevoked)
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
func (it *AutoBetRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutoBetRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutoBetRoleRevoked represents a RoleRevoked event raised by the AutoBet contract.
type AutoBetRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AutoBet *AutoBetFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AutoBetRoleRevokedIterator, error) {

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

	logs, sub, err := _AutoBet.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AutoBetRoleRevokedIterator{contract: _AutoBet.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AutoBet *AutoBetFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AutoBetRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AutoBet.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutoBetRoleRevoked)
				if err := _AutoBet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AutoBet *AutoBetFilterer) ParseRoleRevoked(log types.Log) (*AutoBetRoleRevoked, error) {
	event := new(AutoBetRoleRevoked)
	if err := _AutoBet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
