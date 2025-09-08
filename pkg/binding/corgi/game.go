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

// GameMetaData contains all meta data concerning the Game contract.
var GameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"BindParentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"betSwitch_\",\"type\":\"bool\"}],\"name\":\"admin_set_bet_switch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"orderAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"autoBetAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jackPotAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gameCategoryAddress_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"betSwitch_\",\"type\":\"bool\"}],\"name\":\"admin_set_env\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auto_\",\"outputs\":[{\"internalType\":\"contractIAutoBet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"game_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initial_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"auto_set\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashArr\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"auto_bet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"game_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"betSingleSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"betSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"}],\"name\":\"bind_parent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"changeUserBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"game_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"order_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"seed\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hash_expired\",\"type\":\"bool\"}],\"name\":\"draw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drawData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"game_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auto_id\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"seed\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rate\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"hash_expired\",\"type\":\"bool[]\"}],\"name\":\"draw_auto_bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"game_id\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"order_id\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"seed\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rate\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"hash_expired\",\"type\":\"bool[]\"}],\"name\":\"draw_batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"contractIFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameCategory\",\"outputs\":[{\"internalType\":\"contractIGameCategory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasBind\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jackPot\",\"outputs\":[{\"internalType\":\"contractIJackpot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"order\",\"outputs\":[{\"internalType\":\"contractIOrder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userParent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userSons\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GameABI is the input ABI used to generate the binding from.
// Deprecated: Use GameMetaData.ABI instead.
var GameABI = GameMetaData.ABI

// Game is an auto generated Go binding around an Ethereum contract.
type Game struct {
	GameCaller     // Read-only binding to the contract
	GameTransactor // Write-only binding to the contract
	GameFilterer   // Log filterer for contract events
}

// GameCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameSession struct {
	Contract     *Game             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameCallerSession struct {
	Contract *GameCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameTransactorSession struct {
	Contract     *GameTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameRaw struct {
	Contract *Game // Generic contract binding to access the raw methods on
}

// GameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameCallerRaw struct {
	Contract *GameCaller // Generic read-only contract binding to access the raw methods on
}

// GameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameTransactorRaw struct {
	Contract *GameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGame creates a new instance of Game, bound to a specific deployed contract.
func NewGame(address common.Address, backend bind.ContractBackend) (*Game, error) {
	contract, err := bindGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

// NewGameCaller creates a new read-only instance of Game, bound to a specific deployed contract.
func NewGameCaller(address common.Address, caller bind.ContractCaller) (*GameCaller, error) {
	contract, err := bindGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameCaller{contract: contract}, nil
}

// NewGameTransactor creates a new write-only instance of Game, bound to a specific deployed contract.
func NewGameTransactor(address common.Address, transactor bind.ContractTransactor) (*GameTransactor, error) {
	contract, err := bindGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameTransactor{contract: contract}, nil
}

// NewGameFilterer creates a new log filterer instance of Game, bound to a specific deployed contract.
func NewGameFilterer(address common.Address, filterer bind.ContractFilterer) (*GameFilterer, error) {
	contract, err := bindGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameFilterer{contract: contract}, nil
}

// bindGame binds a generic wrapper to an already deployed contract.
func bindGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.GameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Game *GameCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Game *GameSession) ADMINROLE() ([32]byte, error) {
	return _Game.Contract.ADMINROLE(&_Game.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Game *GameCallerSession) ADMINROLE() ([32]byte, error) {
	return _Game.Contract.ADMINROLE(&_Game.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Game *GameCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Game *GameSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Game.Contract.DEFAULTADMINROLE(&_Game.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Game *GameCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Game.Contract.DEFAULTADMINROLE(&_Game.CallOpts)
}

// MANAGEROLE is a free data retrieval call binding the contract method 0x60a4b76a.
//
// Solidity: function MANAGE_ROLE() view returns(bytes32)
func (_Game *GameCaller) MANAGEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "MANAGE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEROLE is a free data retrieval call binding the contract method 0x60a4b76a.
//
// Solidity: function MANAGE_ROLE() view returns(bytes32)
func (_Game *GameSession) MANAGEROLE() ([32]byte, error) {
	return _Game.Contract.MANAGEROLE(&_Game.CallOpts)
}

// MANAGEROLE is a free data retrieval call binding the contract method 0x60a4b76a.
//
// Solidity: function MANAGE_ROLE() view returns(bytes32)
func (_Game *GameCallerSession) MANAGEROLE() ([32]byte, error) {
	return _Game.Contract.MANAGEROLE(&_Game.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Game *GameCaller) SERVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "SERVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Game *GameSession) SERVERROLE() ([32]byte, error) {
	return _Game.Contract.SERVERROLE(&_Game.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Game *GameCallerSession) SERVERROLE() ([32]byte, error) {
	return _Game.Contract.SERVERROLE(&_Game.CallOpts)
}

// Auto is a free data retrieval call binding the contract method 0x7f66a762.
//
// Solidity: function auto_() view returns(address)
func (_Game *GameCaller) Auto(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "auto_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auto is a free data retrieval call binding the contract method 0x7f66a762.
//
// Solidity: function auto_() view returns(address)
func (_Game *GameSession) Auto() (common.Address, error) {
	return _Game.Contract.Auto(&_Game.CallOpts)
}

// Auto is a free data retrieval call binding the contract method 0x7f66a762.
//
// Solidity: function auto_() view returns(address)
func (_Game *GameCallerSession) Auto() (common.Address, error) {
	return _Game.Contract.Auto(&_Game.CallOpts)
}

// BetSingleSwitch is a free data retrieval call binding the contract method 0xfeda4c14.
//
// Solidity: function betSingleSwitch(uint256 ) view returns(bool)
func (_Game *GameCaller) BetSingleSwitch(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "betSingleSwitch", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BetSingleSwitch is a free data retrieval call binding the contract method 0xfeda4c14.
//
// Solidity: function betSingleSwitch(uint256 ) view returns(bool)
func (_Game *GameSession) BetSingleSwitch(arg0 *big.Int) (bool, error) {
	return _Game.Contract.BetSingleSwitch(&_Game.CallOpts, arg0)
}

// BetSingleSwitch is a free data retrieval call binding the contract method 0xfeda4c14.
//
// Solidity: function betSingleSwitch(uint256 ) view returns(bool)
func (_Game *GameCallerSession) BetSingleSwitch(arg0 *big.Int) (bool, error) {
	return _Game.Contract.BetSingleSwitch(&_Game.CallOpts, arg0)
}

// BetSwitch is a free data retrieval call binding the contract method 0x83428687.
//
// Solidity: function betSwitch() view returns(bool)
func (_Game *GameCaller) BetSwitch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "betSwitch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BetSwitch is a free data retrieval call binding the contract method 0x83428687.
//
// Solidity: function betSwitch() view returns(bool)
func (_Game *GameSession) BetSwitch() (bool, error) {
	return _Game.Contract.BetSwitch(&_Game.CallOpts)
}

// BetSwitch is a free data retrieval call binding the contract method 0x83428687.
//
// Solidity: function betSwitch() view returns(bool)
func (_Game *GameCallerSession) BetSwitch() (bool, error) {
	return _Game.Contract.BetSwitch(&_Game.CallOpts)
}

// DrawData is a free data retrieval call binding the contract method 0xc172d6a4.
//
// Solidity: function drawData(uint256 ) view returns(uint256)
func (_Game *GameCaller) DrawData(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "drawData", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DrawData is a free data retrieval call binding the contract method 0xc172d6a4.
//
// Solidity: function drawData(uint256 ) view returns(uint256)
func (_Game *GameSession) DrawData(arg0 *big.Int) (*big.Int, error) {
	return _Game.Contract.DrawData(&_Game.CallOpts, arg0)
}

// DrawData is a free data retrieval call binding the contract method 0xc172d6a4.
//
// Solidity: function drawData(uint256 ) view returns(uint256)
func (_Game *GameCallerSession) DrawData(arg0 *big.Int) (*big.Int, error) {
	return _Game.Contract.DrawData(&_Game.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_Game *GameCaller) Fee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_Game *GameSession) Fee() (common.Address, error) {
	return _Game.Contract.Fee(&_Game.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_Game *GameCallerSession) Fee() (common.Address, error) {
	return _Game.Contract.Fee(&_Game.CallOpts)
}

// GameCategory is a free data retrieval call binding the contract method 0x816928f0.
//
// Solidity: function gameCategory() view returns(address)
func (_Game *GameCaller) GameCategory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "gameCategory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameCategory is a free data retrieval call binding the contract method 0x816928f0.
//
// Solidity: function gameCategory() view returns(address)
func (_Game *GameSession) GameCategory() (common.Address, error) {
	return _Game.Contract.GameCategory(&_Game.CallOpts)
}

// GameCategory is a free data retrieval call binding the contract method 0x816928f0.
//
// Solidity: function gameCategory() view returns(address)
func (_Game *GameCallerSession) GameCategory() (common.Address, error) {
	return _Game.Contract.GameCategory(&_Game.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Game *GameCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Game *GameSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Game.Contract.GetRoleAdmin(&_Game.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Game *GameCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Game.Contract.GetRoleAdmin(&_Game.CallOpts, role)
}

// HasBind is a free data retrieval call binding the contract method 0x5dbdebf6.
//
// Solidity: function hasBind(address ) view returns(bool)
func (_Game *GameCaller) HasBind(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "hasBind", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBind is a free data retrieval call binding the contract method 0x5dbdebf6.
//
// Solidity: function hasBind(address ) view returns(bool)
func (_Game *GameSession) HasBind(arg0 common.Address) (bool, error) {
	return _Game.Contract.HasBind(&_Game.CallOpts, arg0)
}

// HasBind is a free data retrieval call binding the contract method 0x5dbdebf6.
//
// Solidity: function hasBind(address ) view returns(bool)
func (_Game *GameCallerSession) HasBind(arg0 common.Address) (bool, error) {
	return _Game.Contract.HasBind(&_Game.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Game *GameCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Game *GameSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Game.Contract.HasRole(&_Game.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Game *GameCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Game.Contract.HasRole(&_Game.CallOpts, role, account)
}

// JackPot is a free data retrieval call binding the contract method 0x4a78cdba.
//
// Solidity: function jackPot() view returns(address)
func (_Game *GameCaller) JackPot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "jackPot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JackPot is a free data retrieval call binding the contract method 0x4a78cdba.
//
// Solidity: function jackPot() view returns(address)
func (_Game *GameSession) JackPot() (common.Address, error) {
	return _Game.Contract.JackPot(&_Game.CallOpts)
}

// JackPot is a free data retrieval call binding the contract method 0x4a78cdba.
//
// Solidity: function jackPot() view returns(address)
func (_Game *GameCallerSession) JackPot() (common.Address, error) {
	return _Game.Contract.JackPot(&_Game.CallOpts)
}

// Order is a free data retrieval call binding the contract method 0xbf15071d.
//
// Solidity: function order() view returns(address)
func (_Game *GameCaller) Order(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "order")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Order is a free data retrieval call binding the contract method 0xbf15071d.
//
// Solidity: function order() view returns(address)
func (_Game *GameSession) Order() (common.Address, error) {
	return _Game.Contract.Order(&_Game.CallOpts)
}

// Order is a free data retrieval call binding the contract method 0xbf15071d.
//
// Solidity: function order() view returns(address)
func (_Game *GameCallerSession) Order() (common.Address, error) {
	return _Game.Contract.Order(&_Game.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Game *GameCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Game *GameSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Game.Contract.SupportsInterface(&_Game.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Game *GameCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Game.Contract.SupportsInterface(&_Game.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Game *GameCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Game *GameSession) Token() (common.Address, error) {
	return _Game.Contract.Token(&_Game.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Game *GameCallerSession) Token() (common.Address, error) {
	return _Game.Contract.Token(&_Game.CallOpts)
}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(uint256)
func (_Game *GameCaller) UserOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "userOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(uint256)
func (_Game *GameSession) UserOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Game.Contract.UserOrders(&_Game.CallOpts, arg0, arg1)
}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(uint256)
func (_Game *GameCallerSession) UserOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Game.Contract.UserOrders(&_Game.CallOpts, arg0, arg1)
}

// UserParent is a free data retrieval call binding the contract method 0x0dfcaac6.
//
// Solidity: function userParent(address ) view returns(address)
func (_Game *GameCaller) UserParent(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "userParent", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserParent is a free data retrieval call binding the contract method 0x0dfcaac6.
//
// Solidity: function userParent(address ) view returns(address)
func (_Game *GameSession) UserParent(arg0 common.Address) (common.Address, error) {
	return _Game.Contract.UserParent(&_Game.CallOpts, arg0)
}

// UserParent is a free data retrieval call binding the contract method 0x0dfcaac6.
//
// Solidity: function userParent(address ) view returns(address)
func (_Game *GameCallerSession) UserParent(arg0 common.Address) (common.Address, error) {
	return _Game.Contract.UserParent(&_Game.CallOpts, arg0)
}

// UserSons is a free data retrieval call binding the contract method 0xaa836bd7.
//
// Solidity: function userSons(address , uint256 ) view returns(address)
func (_Game *GameCaller) UserSons(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "userSons", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserSons is a free data retrieval call binding the contract method 0xaa836bd7.
//
// Solidity: function userSons(address , uint256 ) view returns(address)
func (_Game *GameSession) UserSons(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Game.Contract.UserSons(&_Game.CallOpts, arg0, arg1)
}

// UserSons is a free data retrieval call binding the contract method 0xaa836bd7.
//
// Solidity: function userSons(address , uint256 ) view returns(address)
func (_Game *GameCallerSession) UserSons(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Game.Contract.UserSons(&_Game.CallOpts, arg0, arg1)
}

// AdminSetBetSwitch is a paid mutator transaction binding the contract method 0x5887c2b0.
//
// Solidity: function admin_set_bet_switch(uint256 gameId, bool betSwitch_) returns()
func (_Game *GameTransactor) AdminSetBetSwitch(opts *bind.TransactOpts, gameId *big.Int, betSwitch_ bool) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "admin_set_bet_switch", gameId, betSwitch_)
}

// AdminSetBetSwitch is a paid mutator transaction binding the contract method 0x5887c2b0.
//
// Solidity: function admin_set_bet_switch(uint256 gameId, bool betSwitch_) returns()
func (_Game *GameSession) AdminSetBetSwitch(gameId *big.Int, betSwitch_ bool) (*types.Transaction, error) {
	return _Game.Contract.AdminSetBetSwitch(&_Game.TransactOpts, gameId, betSwitch_)
}

// AdminSetBetSwitch is a paid mutator transaction binding the contract method 0x5887c2b0.
//
// Solidity: function admin_set_bet_switch(uint256 gameId, bool betSwitch_) returns()
func (_Game *GameTransactorSession) AdminSetBetSwitch(gameId *big.Int, betSwitch_ bool) (*types.Transaction, error) {
	return _Game.Contract.AdminSetBetSwitch(&_Game.TransactOpts, gameId, betSwitch_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x0fd29359.
//
// Solidity: function admin_set_env(address feeAddress_, address tokenAddress_, address orderAddress_, address autoBetAddress_, address jackPotAddress_, address gameCategoryAddress_, bool betSwitch_) returns()
func (_Game *GameTransactor) AdminSetEnv(opts *bind.TransactOpts, feeAddress_ common.Address, tokenAddress_ common.Address, orderAddress_ common.Address, autoBetAddress_ common.Address, jackPotAddress_ common.Address, gameCategoryAddress_ common.Address, betSwitch_ bool) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "admin_set_env", feeAddress_, tokenAddress_, orderAddress_, autoBetAddress_, jackPotAddress_, gameCategoryAddress_, betSwitch_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x0fd29359.
//
// Solidity: function admin_set_env(address feeAddress_, address tokenAddress_, address orderAddress_, address autoBetAddress_, address jackPotAddress_, address gameCategoryAddress_, bool betSwitch_) returns()
func (_Game *GameSession) AdminSetEnv(feeAddress_ common.Address, tokenAddress_ common.Address, orderAddress_ common.Address, autoBetAddress_ common.Address, jackPotAddress_ common.Address, gameCategoryAddress_ common.Address, betSwitch_ bool) (*types.Transaction, error) {
	return _Game.Contract.AdminSetEnv(&_Game.TransactOpts, feeAddress_, tokenAddress_, orderAddress_, autoBetAddress_, jackPotAddress_, gameCategoryAddress_, betSwitch_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x0fd29359.
//
// Solidity: function admin_set_env(address feeAddress_, address tokenAddress_, address orderAddress_, address autoBetAddress_, address jackPotAddress_, address gameCategoryAddress_, bool betSwitch_) returns()
func (_Game *GameTransactorSession) AdminSetEnv(feeAddress_ common.Address, tokenAddress_ common.Address, orderAddress_ common.Address, autoBetAddress_ common.Address, jackPotAddress_ common.Address, gameCategoryAddress_ common.Address, betSwitch_ bool) (*types.Transaction, error) {
	return _Game.Contract.AdminSetEnv(&_Game.TransactOpts, feeAddress_, tokenAddress_, orderAddress_, autoBetAddress_, jackPotAddress_, gameCategoryAddress_, betSwitch_)
}

// AutoBet is a paid mutator transaction binding the contract method 0x1f310457.
//
// Solidity: function auto_bet(uint256 game_id, uint256 token_id, uint256 initial_amount, uint256[] auto_set, bytes32[] hashArr, bytes data) payable returns()
func (_Game *GameTransactor) AutoBet(opts *bind.TransactOpts, game_id *big.Int, token_id *big.Int, initial_amount *big.Int, auto_set []*big.Int, hashArr [][32]byte, data []byte) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "auto_bet", game_id, token_id, initial_amount, auto_set, hashArr, data)
}

// AutoBet is a paid mutator transaction binding the contract method 0x1f310457.
//
// Solidity: function auto_bet(uint256 game_id, uint256 token_id, uint256 initial_amount, uint256[] auto_set, bytes32[] hashArr, bytes data) payable returns()
func (_Game *GameSession) AutoBet(game_id *big.Int, token_id *big.Int, initial_amount *big.Int, auto_set []*big.Int, hashArr [][32]byte, data []byte) (*types.Transaction, error) {
	return _Game.Contract.AutoBet(&_Game.TransactOpts, game_id, token_id, initial_amount, auto_set, hashArr, data)
}

// AutoBet is a paid mutator transaction binding the contract method 0x1f310457.
//
// Solidity: function auto_bet(uint256 game_id, uint256 token_id, uint256 initial_amount, uint256[] auto_set, bytes32[] hashArr, bytes data) payable returns()
func (_Game *GameTransactorSession) AutoBet(game_id *big.Int, token_id *big.Int, initial_amount *big.Int, auto_set []*big.Int, hashArr [][32]byte, data []byte) (*types.Transaction, error) {
	return _Game.Contract.AutoBet(&_Game.TransactOpts, game_id, token_id, initial_amount, auto_set, hashArr, data)
}

// Bet is a paid mutator transaction binding the contract method 0xc8d1bc80.
//
// Solidity: function bet(uint256 game_id, uint256 token_id, uint256 amount, bytes32 hash, bytes data) payable returns()
func (_Game *GameTransactor) Bet(opts *bind.TransactOpts, game_id *big.Int, token_id *big.Int, amount *big.Int, hash [32]byte, data []byte) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "bet", game_id, token_id, amount, hash, data)
}

// Bet is a paid mutator transaction binding the contract method 0xc8d1bc80.
//
// Solidity: function bet(uint256 game_id, uint256 token_id, uint256 amount, bytes32 hash, bytes data) payable returns()
func (_Game *GameSession) Bet(game_id *big.Int, token_id *big.Int, amount *big.Int, hash [32]byte, data []byte) (*types.Transaction, error) {
	return _Game.Contract.Bet(&_Game.TransactOpts, game_id, token_id, amount, hash, data)
}

// Bet is a paid mutator transaction binding the contract method 0xc8d1bc80.
//
// Solidity: function bet(uint256 game_id, uint256 token_id, uint256 amount, bytes32 hash, bytes data) payable returns()
func (_Game *GameTransactorSession) Bet(game_id *big.Int, token_id *big.Int, amount *big.Int, hash [32]byte, data []byte) (*types.Transaction, error) {
	return _Game.Contract.Bet(&_Game.TransactOpts, game_id, token_id, amount, hash, data)
}

// BindParent is a paid mutator transaction binding the contract method 0x1e44cc8f.
//
// Solidity: function bind_parent(address user_address) returns()
func (_Game *GameTransactor) BindParent(opts *bind.TransactOpts, user_address common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "bind_parent", user_address)
}

// BindParent is a paid mutator transaction binding the contract method 0x1e44cc8f.
//
// Solidity: function bind_parent(address user_address) returns()
func (_Game *GameSession) BindParent(user_address common.Address) (*types.Transaction, error) {
	return _Game.Contract.BindParent(&_Game.TransactOpts, user_address)
}

// BindParent is a paid mutator transaction binding the contract method 0x1e44cc8f.
//
// Solidity: function bind_parent(address user_address) returns()
func (_Game *GameTransactorSession) BindParent(user_address common.Address) (*types.Transaction, error) {
	return _Game.Contract.BindParent(&_Game.TransactOpts, user_address)
}

// ChangeUserBalance is a paid mutator transaction binding the contract method 0x4e8cf658.
//
// Solidity: function changeUserBalance(address user, uint256 tokenId, uint256 amount) returns()
func (_Game *GameTransactor) ChangeUserBalance(opts *bind.TransactOpts, user common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "changeUserBalance", user, tokenId, amount)
}

// ChangeUserBalance is a paid mutator transaction binding the contract method 0x4e8cf658.
//
// Solidity: function changeUserBalance(address user, uint256 tokenId, uint256 amount) returns()
func (_Game *GameSession) ChangeUserBalance(user common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Game.Contract.ChangeUserBalance(&_Game.TransactOpts, user, tokenId, amount)
}

// ChangeUserBalance is a paid mutator transaction binding the contract method 0x4e8cf658.
//
// Solidity: function changeUserBalance(address user, uint256 tokenId, uint256 amount) returns()
func (_Game *GameTransactorSession) ChangeUserBalance(user common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Game.Contract.ChangeUserBalance(&_Game.TransactOpts, user, tokenId, amount)
}

// Draw is a paid mutator transaction binding the contract method 0x25b7c593.
//
// Solidity: function draw(uint256 game_id, uint256 order_id, string seed, uint256 rate, bool hash_expired) returns()
func (_Game *GameTransactor) Draw(opts *bind.TransactOpts, game_id *big.Int, order_id *big.Int, seed string, rate *big.Int, hash_expired bool) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "draw", game_id, order_id, seed, rate, hash_expired)
}

// Draw is a paid mutator transaction binding the contract method 0x25b7c593.
//
// Solidity: function draw(uint256 game_id, uint256 order_id, string seed, uint256 rate, bool hash_expired) returns()
func (_Game *GameSession) Draw(game_id *big.Int, order_id *big.Int, seed string, rate *big.Int, hash_expired bool) (*types.Transaction, error) {
	return _Game.Contract.Draw(&_Game.TransactOpts, game_id, order_id, seed, rate, hash_expired)
}

// Draw is a paid mutator transaction binding the contract method 0x25b7c593.
//
// Solidity: function draw(uint256 game_id, uint256 order_id, string seed, uint256 rate, bool hash_expired) returns()
func (_Game *GameTransactorSession) Draw(game_id *big.Int, order_id *big.Int, seed string, rate *big.Int, hash_expired bool) (*types.Transaction, error) {
	return _Game.Contract.Draw(&_Game.TransactOpts, game_id, order_id, seed, rate, hash_expired)
}

// DrawAutoBet is a paid mutator transaction binding the contract method 0x813cdc6a.
//
// Solidity: function draw_auto_bet(uint256 game_id, uint256 auto_id, string[] seed, uint256[] rate, bool[] hash_expired) returns()
func (_Game *GameTransactor) DrawAutoBet(opts *bind.TransactOpts, game_id *big.Int, auto_id *big.Int, seed []string, rate []*big.Int, hash_expired []bool) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "draw_auto_bet", game_id, auto_id, seed, rate, hash_expired)
}

// DrawAutoBet is a paid mutator transaction binding the contract method 0x813cdc6a.
//
// Solidity: function draw_auto_bet(uint256 game_id, uint256 auto_id, string[] seed, uint256[] rate, bool[] hash_expired) returns()
func (_Game *GameSession) DrawAutoBet(game_id *big.Int, auto_id *big.Int, seed []string, rate []*big.Int, hash_expired []bool) (*types.Transaction, error) {
	return _Game.Contract.DrawAutoBet(&_Game.TransactOpts, game_id, auto_id, seed, rate, hash_expired)
}

// DrawAutoBet is a paid mutator transaction binding the contract method 0x813cdc6a.
//
// Solidity: function draw_auto_bet(uint256 game_id, uint256 auto_id, string[] seed, uint256[] rate, bool[] hash_expired) returns()
func (_Game *GameTransactorSession) DrawAutoBet(game_id *big.Int, auto_id *big.Int, seed []string, rate []*big.Int, hash_expired []bool) (*types.Transaction, error) {
	return _Game.Contract.DrawAutoBet(&_Game.TransactOpts, game_id, auto_id, seed, rate, hash_expired)
}

// DrawBatch is a paid mutator transaction binding the contract method 0xaaeddeaa.
//
// Solidity: function draw_batch(uint256[] game_id, uint256[] order_id, string[] seed, uint256[] rate, bool[] hash_expired) returns()
func (_Game *GameTransactor) DrawBatch(opts *bind.TransactOpts, game_id []*big.Int, order_id []*big.Int, seed []string, rate []*big.Int, hash_expired []bool) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "draw_batch", game_id, order_id, seed, rate, hash_expired)
}

// DrawBatch is a paid mutator transaction binding the contract method 0xaaeddeaa.
//
// Solidity: function draw_batch(uint256[] game_id, uint256[] order_id, string[] seed, uint256[] rate, bool[] hash_expired) returns()
func (_Game *GameSession) DrawBatch(game_id []*big.Int, order_id []*big.Int, seed []string, rate []*big.Int, hash_expired []bool) (*types.Transaction, error) {
	return _Game.Contract.DrawBatch(&_Game.TransactOpts, game_id, order_id, seed, rate, hash_expired)
}

// DrawBatch is a paid mutator transaction binding the contract method 0xaaeddeaa.
//
// Solidity: function draw_batch(uint256[] game_id, uint256[] order_id, string[] seed, uint256[] rate, bool[] hash_expired) returns()
func (_Game *GameTransactorSession) DrawBatch(game_id []*big.Int, order_id []*big.Int, seed []string, rate []*big.Int, hash_expired []bool) (*types.Transaction, error) {
	return _Game.Contract.DrawBatch(&_Game.TransactOpts, game_id, order_id, seed, rate, hash_expired)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Game *GameTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Game *GameSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Game.Contract.GrantRole(&_Game.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Game *GameTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Game.Contract.GrantRole(&_Game.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Game *GameTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Game *GameSession) Initialize() (*types.Transaction, error) {
	return _Game.Contract.Initialize(&_Game.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Game *GameTransactorSession) Initialize() (*types.Transaction, error) {
	return _Game.Contract.Initialize(&_Game.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Game *GameTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Game *GameSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Game.Contract.RenounceRole(&_Game.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Game *GameTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Game.Contract.RenounceRole(&_Game.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Game *GameTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Game *GameSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Game.Contract.RevokeRole(&_Game.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Game *GameTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Game.Contract.RevokeRole(&_Game.TransactOpts, role, account)
}

// Withdrawal is a paid mutator transaction binding the contract method 0xc72a8696.
//
// Solidity: function withdrawal(uint256 token_id, uint256 amount) returns()
func (_Game *GameTransactor) Withdrawal(opts *bind.TransactOpts, token_id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "withdrawal", token_id, amount)
}

// Withdrawal is a paid mutator transaction binding the contract method 0xc72a8696.
//
// Solidity: function withdrawal(uint256 token_id, uint256 amount) returns()
func (_Game *GameSession) Withdrawal(token_id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Game.Contract.Withdrawal(&_Game.TransactOpts, token_id, amount)
}

// Withdrawal is a paid mutator transaction binding the contract method 0xc72a8696.
//
// Solidity: function withdrawal(uint256 token_id, uint256 amount) returns()
func (_Game *GameTransactorSession) Withdrawal(token_id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Game.Contract.Withdrawal(&_Game.TransactOpts, token_id, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Game *GameTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Game *GameSession) Receive() (*types.Transaction, error) {
	return _Game.Contract.Receive(&_Game.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Game *GameTransactorSession) Receive() (*types.Transaction, error) {
	return _Game.Contract.Receive(&_Game.TransactOpts)
}

// GameBindParentEventIterator is returned from FilterBindParentEvent and is used to iterate over the raw logs and unpacked data for BindParentEvent events raised by the Game contract.
type GameBindParentEventIterator struct {
	Event *GameBindParentEvent // Event containing the contract specifics and raw log

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
func (it *GameBindParentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameBindParentEvent)
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
		it.Event = new(GameBindParentEvent)
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
func (it *GameBindParentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameBindParentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameBindParentEvent represents a BindParentEvent event raised by the Game contract.
type GameBindParentEvent struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBindParentEvent is a free log retrieval operation binding the contract event 0x4d77ea39d3b5c57d6486cafdc5a83a3b17975aa0bc191b02ea143428ddd73b1a.
//
// Solidity: event BindParentEvent(address indexed user)
func (_Game *GameFilterer) FilterBindParentEvent(opts *bind.FilterOpts, user []common.Address) (*GameBindParentEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "BindParentEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &GameBindParentEventIterator{contract: _Game.contract, event: "BindParentEvent", logs: logs, sub: sub}, nil
}

// WatchBindParentEvent is a free log subscription operation binding the contract event 0x4d77ea39d3b5c57d6486cafdc5a83a3b17975aa0bc191b02ea143428ddd73b1a.
//
// Solidity: event BindParentEvent(address indexed user)
func (_Game *GameFilterer) WatchBindParentEvent(opts *bind.WatchOpts, sink chan<- *GameBindParentEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "BindParentEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameBindParentEvent)
				if err := _Game.contract.UnpackLog(event, "BindParentEvent", log); err != nil {
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

// ParseBindParentEvent is a log parse operation binding the contract event 0x4d77ea39d3b5c57d6486cafdc5a83a3b17975aa0bc191b02ea143428ddd73b1a.
//
// Solidity: event BindParentEvent(address indexed user)
func (_Game *GameFilterer) ParseBindParentEvent(log types.Log) (*GameBindParentEvent, error) {
	event := new(GameBindParentEvent)
	if err := _Game.contract.UnpackLog(event, "BindParentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Game contract.
type GameInitializedIterator struct {
	Event *GameInitialized // Event containing the contract specifics and raw log

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
func (it *GameInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameInitialized)
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
		it.Event = new(GameInitialized)
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
func (it *GameInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameInitialized represents a Initialized event raised by the Game contract.
type GameInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Game *GameFilterer) FilterInitialized(opts *bind.FilterOpts) (*GameInitializedIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GameInitializedIterator{contract: _Game.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Game *GameFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GameInitialized) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameInitialized)
				if err := _Game.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Game *GameFilterer) ParseInitialized(log types.Log) (*GameInitialized, error) {
	event := new(GameInitialized)
	if err := _Game.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Game contract.
type GameRoleAdminChangedIterator struct {
	Event *GameRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GameRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameRoleAdminChanged)
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
		it.Event = new(GameRoleAdminChanged)
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
func (it *GameRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameRoleAdminChanged represents a RoleAdminChanged event raised by the Game contract.
type GameRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Game *GameFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GameRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Game.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GameRoleAdminChangedIterator{contract: _Game.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Game *GameFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GameRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Game.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameRoleAdminChanged)
				if err := _Game.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Game *GameFilterer) ParseRoleAdminChanged(log types.Log) (*GameRoleAdminChanged, error) {
	event := new(GameRoleAdminChanged)
	if err := _Game.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Game contract.
type GameRoleGrantedIterator struct {
	Event *GameRoleGranted // Event containing the contract specifics and raw log

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
func (it *GameRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameRoleGranted)
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
		it.Event = new(GameRoleGranted)
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
func (it *GameRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameRoleGranted represents a RoleGranted event raised by the Game contract.
type GameRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Game *GameFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameRoleGrantedIterator, error) {

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

	logs, sub, err := _Game.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameRoleGrantedIterator{contract: _Game.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Game *GameFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GameRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Game.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameRoleGranted)
				if err := _Game.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Game *GameFilterer) ParseRoleGranted(log types.Log) (*GameRoleGranted, error) {
	event := new(GameRoleGranted)
	if err := _Game.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Game contract.
type GameRoleRevokedIterator struct {
	Event *GameRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GameRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameRoleRevoked)
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
		it.Event = new(GameRoleRevoked)
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
func (it *GameRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameRoleRevoked represents a RoleRevoked event raised by the Game contract.
type GameRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Game *GameFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameRoleRevokedIterator, error) {

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

	logs, sub, err := _Game.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameRoleRevokedIterator{contract: _Game.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Game *GameFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GameRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Game.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameRoleRevoked)
				if err := _Game.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Game *GameFilterer) ParseRoleRevoked(log types.Log) (*GameRoleRevoked, error) {
	event := new(GameRoleRevoked)
	if err := _Game.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameWithdrawalEventIterator is returned from FilterWithdrawalEvent and is used to iterate over the raw logs and unpacked data for WithdrawalEvent events raised by the Game contract.
type GameWithdrawalEventIterator struct {
	Event *GameWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *GameWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameWithdrawalEvent)
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
		it.Event = new(GameWithdrawalEvent)
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
func (it *GameWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameWithdrawalEvent represents a WithdrawalEvent event raised by the Game contract.
type GameWithdrawalEvent struct {
	User    common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0xbda29a9e6b4ed78b3ef7f48ca275518ddc1404cc7fb70813569407d029e59902.
//
// Solidity: event WithdrawalEvent(address indexed user, uint256 indexed tokenId, uint256 indexed amount)
func (_Game *GameFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts, user []common.Address, tokenId []*big.Int, amount []*big.Int) (*GameWithdrawalEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "WithdrawalEvent", userRule, tokenIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &GameWithdrawalEventIterator{contract: _Game.contract, event: "WithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0xbda29a9e6b4ed78b3ef7f48ca275518ddc1404cc7fb70813569407d029e59902.
//
// Solidity: event WithdrawalEvent(address indexed user, uint256 indexed tokenId, uint256 indexed amount)
func (_Game *GameFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *GameWithdrawalEvent, user []common.Address, tokenId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "WithdrawalEvent", userRule, tokenIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameWithdrawalEvent)
				if err := _Game.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0xbda29a9e6b4ed78b3ef7f48ca275518ddc1404cc7fb70813569407d029e59902.
//
// Solidity: event WithdrawalEvent(address indexed user, uint256 indexed tokenId, uint256 indexed amount)
func (_Game *GameFilterer) ParseWithdrawalEvent(log types.Log) (*GameWithdrawalEvent, error) {
	event := new(GameWithdrawalEvent)
	if err := _Game.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
