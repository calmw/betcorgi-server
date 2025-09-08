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

// GameCategoryMetaData contains all meta data concerning the GameCategory contract.
var GameCategoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gameName\",\"type\":\"string\"}],\"name\":\"GameCategoryEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"game_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"game_name\",\"type\":\"string\"}],\"name\":\"admin_set_game_category\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameCategories\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gameName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"gameExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"gameName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GameCategoryABI is the input ABI used to generate the binding from.
// Deprecated: Use GameCategoryMetaData.ABI instead.
var GameCategoryABI = GameCategoryMetaData.ABI

// GameCategory is an auto generated Go binding around an Ethereum contract.
type GameCategory struct {
	GameCategoryCaller     // Read-only binding to the contract
	GameCategoryTransactor // Write-only binding to the contract
	GameCategoryFilterer   // Log filterer for contract events
}

// GameCategoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameCategoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameCategoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameCategoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameCategoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameCategoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameCategorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameCategorySession struct {
	Contract     *GameCategory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameCategoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameCategoryCallerSession struct {
	Contract *GameCategoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GameCategoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameCategoryTransactorSession struct {
	Contract     *GameCategoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GameCategoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameCategoryRaw struct {
	Contract *GameCategory // Generic contract binding to access the raw methods on
}

// GameCategoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameCategoryCallerRaw struct {
	Contract *GameCategoryCaller // Generic read-only contract binding to access the raw methods on
}

// GameCategoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameCategoryTransactorRaw struct {
	Contract *GameCategoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameCategory creates a new instance of GameCategory, bound to a specific deployed contract.
func NewGameCategory(address common.Address, backend bind.ContractBackend) (*GameCategory, error) {
	contract, err := bindGameCategory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameCategory{GameCategoryCaller: GameCategoryCaller{contract: contract}, GameCategoryTransactor: GameCategoryTransactor{contract: contract}, GameCategoryFilterer: GameCategoryFilterer{contract: contract}}, nil
}

// NewGameCategoryCaller creates a new read-only instance of GameCategory, bound to a specific deployed contract.
func NewGameCategoryCaller(address common.Address, caller bind.ContractCaller) (*GameCategoryCaller, error) {
	contract, err := bindGameCategory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameCategoryCaller{contract: contract}, nil
}

// NewGameCategoryTransactor creates a new write-only instance of GameCategory, bound to a specific deployed contract.
func NewGameCategoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GameCategoryTransactor, error) {
	contract, err := bindGameCategory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameCategoryTransactor{contract: contract}, nil
}

// NewGameCategoryFilterer creates a new log filterer instance of GameCategory, bound to a specific deployed contract.
func NewGameCategoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GameCategoryFilterer, error) {
	contract, err := bindGameCategory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameCategoryFilterer{contract: contract}, nil
}

// bindGameCategory binds a generic wrapper to an already deployed contract.
func bindGameCategory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GameCategoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameCategory *GameCategoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameCategory.Contract.GameCategoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameCategory *GameCategoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameCategory.Contract.GameCategoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameCategory *GameCategoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameCategory.Contract.GameCategoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameCategory *GameCategoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameCategory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameCategory *GameCategoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameCategory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameCategory *GameCategoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameCategory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameCategory *GameCategoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameCategory *GameCategorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameCategory.Contract.DEFAULTADMINROLE(&_GameCategory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameCategory *GameCategoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameCategory.Contract.DEFAULTADMINROLE(&_GameCategory.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_GameCategory *GameCategoryCaller) SERVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "SERVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_GameCategory *GameCategorySession) SERVERROLE() ([32]byte, error) {
	return _GameCategory.Contract.SERVERROLE(&_GameCategory.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_GameCategory *GameCategoryCallerSession) SERVERROLE() ([32]byte, error) {
	return _GameCategory.Contract.SERVERROLE(&_GameCategory.CallOpts)
}

// GameCategories is a free data retrieval call binding the contract method 0x57b03944.
//
// Solidity: function gameCategories(uint256 ) view returns(uint256 gameId, string gameName)
func (_GameCategory *GameCategoryCaller) GameCategories(opts *bind.CallOpts, arg0 *big.Int) (struct {
	GameId   *big.Int
	GameName string
}, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "gameCategories", arg0)

	outstruct := new(struct {
		GameId   *big.Int
		GameName string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GameName = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GameCategories is a free data retrieval call binding the contract method 0x57b03944.
//
// Solidity: function gameCategories(uint256 ) view returns(uint256 gameId, string gameName)
func (_GameCategory *GameCategorySession) GameCategories(arg0 *big.Int) (struct {
	GameId   *big.Int
	GameName string
}, error) {
	return _GameCategory.Contract.GameCategories(&_GameCategory.CallOpts, arg0)
}

// GameCategories is a free data retrieval call binding the contract method 0x57b03944.
//
// Solidity: function gameCategories(uint256 ) view returns(uint256 gameId, string gameName)
func (_GameCategory *GameCategoryCallerSession) GameCategories(arg0 *big.Int) (struct {
	GameId   *big.Int
	GameName string
}, error) {
	return _GameCategory.Contract.GameCategories(&_GameCategory.CallOpts, arg0)
}

// GameExist is a free data retrieval call binding the contract method 0x59fe9dc9.
//
// Solidity: function gameExist(uint256 gameId) view returns(bool)
func (_GameCategory *GameCategoryCaller) GameExist(opts *bind.CallOpts, gameId *big.Int) (bool, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "gameExist", gameId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GameExist is a free data retrieval call binding the contract method 0x59fe9dc9.
//
// Solidity: function gameExist(uint256 gameId) view returns(bool)
func (_GameCategory *GameCategorySession) GameExist(gameId *big.Int) (bool, error) {
	return _GameCategory.Contract.GameExist(&_GameCategory.CallOpts, gameId)
}

// GameExist is a free data retrieval call binding the contract method 0x59fe9dc9.
//
// Solidity: function gameExist(uint256 gameId) view returns(bool)
func (_GameCategory *GameCategoryCallerSession) GameExist(gameId *big.Int) (bool, error) {
	return _GameCategory.Contract.GameExist(&_GameCategory.CallOpts, gameId)
}

// GameName is a free data retrieval call binding the contract method 0xd55f715c.
//
// Solidity: function gameName(uint256 gameId) view returns(string)
func (_GameCategory *GameCategoryCaller) GameName(opts *bind.CallOpts, gameId *big.Int) (string, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "gameName", gameId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GameName is a free data retrieval call binding the contract method 0xd55f715c.
//
// Solidity: function gameName(uint256 gameId) view returns(string)
func (_GameCategory *GameCategorySession) GameName(gameId *big.Int) (string, error) {
	return _GameCategory.Contract.GameName(&_GameCategory.CallOpts, gameId)
}

// GameName is a free data retrieval call binding the contract method 0xd55f715c.
//
// Solidity: function gameName(uint256 gameId) view returns(string)
func (_GameCategory *GameCategoryCallerSession) GameName(gameId *big.Int) (string, error) {
	return _GameCategory.Contract.GameName(&_GameCategory.CallOpts, gameId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameCategory *GameCategoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameCategory *GameCategorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameCategory.Contract.GetRoleAdmin(&_GameCategory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameCategory *GameCategoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameCategory.Contract.GetRoleAdmin(&_GameCategory.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameCategory *GameCategoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameCategory *GameCategorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameCategory.Contract.HasRole(&_GameCategory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameCategory *GameCategoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameCategory.Contract.HasRole(&_GameCategory.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameCategory *GameCategoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GameCategory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameCategory *GameCategorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameCategory.Contract.SupportsInterface(&_GameCategory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameCategory *GameCategoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameCategory.Contract.SupportsInterface(&_GameCategory.CallOpts, interfaceId)
}

// AdminSetGameCategory is a paid mutator transaction binding the contract method 0x89d04d9e.
//
// Solidity: function admin_set_game_category(uint256 game_id, string game_name) returns()
func (_GameCategory *GameCategoryTransactor) AdminSetGameCategory(opts *bind.TransactOpts, game_id *big.Int, game_name string) (*types.Transaction, error) {
	return _GameCategory.contract.Transact(opts, "admin_set_game_category", game_id, game_name)
}

// AdminSetGameCategory is a paid mutator transaction binding the contract method 0x89d04d9e.
//
// Solidity: function admin_set_game_category(uint256 game_id, string game_name) returns()
func (_GameCategory *GameCategorySession) AdminSetGameCategory(game_id *big.Int, game_name string) (*types.Transaction, error) {
	return _GameCategory.Contract.AdminSetGameCategory(&_GameCategory.TransactOpts, game_id, game_name)
}

// AdminSetGameCategory is a paid mutator transaction binding the contract method 0x89d04d9e.
//
// Solidity: function admin_set_game_category(uint256 game_id, string game_name) returns()
func (_GameCategory *GameCategoryTransactorSession) AdminSetGameCategory(game_id *big.Int, game_name string) (*types.Transaction, error) {
	return _GameCategory.Contract.AdminSetGameCategory(&_GameCategory.TransactOpts, game_id, game_name)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameCategory *GameCategoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameCategory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameCategory *GameCategorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameCategory.Contract.GrantRole(&_GameCategory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameCategory *GameCategoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameCategory.Contract.GrantRole(&_GameCategory.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GameCategory *GameCategoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameCategory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GameCategory *GameCategorySession) Initialize() (*types.Transaction, error) {
	return _GameCategory.Contract.Initialize(&_GameCategory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GameCategory *GameCategoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _GameCategory.Contract.Initialize(&_GameCategory.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GameCategory *GameCategoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GameCategory.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GameCategory *GameCategorySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GameCategory.Contract.RenounceRole(&_GameCategory.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GameCategory *GameCategoryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GameCategory.Contract.RenounceRole(&_GameCategory.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameCategory *GameCategoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameCategory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameCategory *GameCategorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameCategory.Contract.RevokeRole(&_GameCategory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameCategory *GameCategoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameCategory.Contract.RevokeRole(&_GameCategory.TransactOpts, role, account)
}

// GameCategoryGameCategoryEventIterator is returned from FilterGameCategoryEvent and is used to iterate over the raw logs and unpacked data for GameCategoryEvent events raised by the GameCategory contract.
type GameCategoryGameCategoryEventIterator struct {
	Event *GameCategoryGameCategoryEvent // Event containing the contract specifics and raw log

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
func (it *GameCategoryGameCategoryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameCategoryGameCategoryEvent)
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
		it.Event = new(GameCategoryGameCategoryEvent)
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
func (it *GameCategoryGameCategoryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameCategoryGameCategoryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameCategoryGameCategoryEvent represents a GameCategoryEvent event raised by the GameCategory contract.
type GameCategoryGameCategoryEvent struct {
	GameId   *big.Int
	GameName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGameCategoryEvent is a free log retrieval operation binding the contract event 0xbb2f91e5b10d4fb5d90eef16483b6ae5d49e1a1daa8b783dcf7b74ae8d6f8b07.
//
// Solidity: event GameCategoryEvent(uint256 gameId, string gameName)
func (_GameCategory *GameCategoryFilterer) FilterGameCategoryEvent(opts *bind.FilterOpts) (*GameCategoryGameCategoryEventIterator, error) {

	logs, sub, err := _GameCategory.contract.FilterLogs(opts, "GameCategoryEvent")
	if err != nil {
		return nil, err
	}
	return &GameCategoryGameCategoryEventIterator{contract: _GameCategory.contract, event: "GameCategoryEvent", logs: logs, sub: sub}, nil
}

// WatchGameCategoryEvent is a free log subscription operation binding the contract event 0xbb2f91e5b10d4fb5d90eef16483b6ae5d49e1a1daa8b783dcf7b74ae8d6f8b07.
//
// Solidity: event GameCategoryEvent(uint256 gameId, string gameName)
func (_GameCategory *GameCategoryFilterer) WatchGameCategoryEvent(opts *bind.WatchOpts, sink chan<- *GameCategoryGameCategoryEvent) (event.Subscription, error) {

	logs, sub, err := _GameCategory.contract.WatchLogs(opts, "GameCategoryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameCategoryGameCategoryEvent)
				if err := _GameCategory.contract.UnpackLog(event, "GameCategoryEvent", log); err != nil {
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

// ParseGameCategoryEvent is a log parse operation binding the contract event 0xbb2f91e5b10d4fb5d90eef16483b6ae5d49e1a1daa8b783dcf7b74ae8d6f8b07.
//
// Solidity: event GameCategoryEvent(uint256 gameId, string gameName)
func (_GameCategory *GameCategoryFilterer) ParseGameCategoryEvent(log types.Log) (*GameCategoryGameCategoryEvent, error) {
	event := new(GameCategoryGameCategoryEvent)
	if err := _GameCategory.contract.UnpackLog(event, "GameCategoryEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameCategoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GameCategory contract.
type GameCategoryInitializedIterator struct {
	Event *GameCategoryInitialized // Event containing the contract specifics and raw log

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
func (it *GameCategoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameCategoryInitialized)
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
		it.Event = new(GameCategoryInitialized)
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
func (it *GameCategoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameCategoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameCategoryInitialized represents a Initialized event raised by the GameCategory contract.
type GameCategoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GameCategory *GameCategoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*GameCategoryInitializedIterator, error) {

	logs, sub, err := _GameCategory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GameCategoryInitializedIterator{contract: _GameCategory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GameCategory *GameCategoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GameCategoryInitialized) (event.Subscription, error) {

	logs, sub, err := _GameCategory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameCategoryInitialized)
				if err := _GameCategory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GameCategory *GameCategoryFilterer) ParseInitialized(log types.Log) (*GameCategoryInitialized, error) {
	event := new(GameCategoryInitialized)
	if err := _GameCategory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameCategoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GameCategory contract.
type GameCategoryRoleAdminChangedIterator struct {
	Event *GameCategoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GameCategoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameCategoryRoleAdminChanged)
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
		it.Event = new(GameCategoryRoleAdminChanged)
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
func (it *GameCategoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameCategoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameCategoryRoleAdminChanged represents a RoleAdminChanged event raised by the GameCategory contract.
type GameCategoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameCategory *GameCategoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GameCategoryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GameCategory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GameCategoryRoleAdminChangedIterator{contract: _GameCategory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameCategory *GameCategoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GameCategoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GameCategory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameCategoryRoleAdminChanged)
				if err := _GameCategory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GameCategory *GameCategoryFilterer) ParseRoleAdminChanged(log types.Log) (*GameCategoryRoleAdminChanged, error) {
	event := new(GameCategoryRoleAdminChanged)
	if err := _GameCategory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameCategoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GameCategory contract.
type GameCategoryRoleGrantedIterator struct {
	Event *GameCategoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *GameCategoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameCategoryRoleGranted)
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
		it.Event = new(GameCategoryRoleGranted)
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
func (it *GameCategoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameCategoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameCategoryRoleGranted represents a RoleGranted event raised by the GameCategory contract.
type GameCategoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameCategory *GameCategoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameCategoryRoleGrantedIterator, error) {

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

	logs, sub, err := _GameCategory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameCategoryRoleGrantedIterator{contract: _GameCategory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameCategory *GameCategoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GameCategoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GameCategory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameCategoryRoleGranted)
				if err := _GameCategory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GameCategory *GameCategoryFilterer) ParseRoleGranted(log types.Log) (*GameCategoryRoleGranted, error) {
	event := new(GameCategoryRoleGranted)
	if err := _GameCategory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameCategoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GameCategory contract.
type GameCategoryRoleRevokedIterator struct {
	Event *GameCategoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GameCategoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameCategoryRoleRevoked)
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
		it.Event = new(GameCategoryRoleRevoked)
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
func (it *GameCategoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameCategoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameCategoryRoleRevoked represents a RoleRevoked event raised by the GameCategory contract.
type GameCategoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameCategory *GameCategoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameCategoryRoleRevokedIterator, error) {

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

	logs, sub, err := _GameCategory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameCategoryRoleRevokedIterator{contract: _GameCategory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameCategory *GameCategoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GameCategoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GameCategory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameCategoryRoleRevoked)
				if err := _GameCategory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GameCategory *GameCategoryFilterer) ParseRoleRevoked(log types.Log) (*GameCategoryRoleRevoked, error) {
	event := new(GameCategoryRoleRevoked)
	if err := _GameCategory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
