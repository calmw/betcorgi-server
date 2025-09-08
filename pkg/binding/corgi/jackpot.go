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

// JackpotMetaData contains all meta data concerning the Jackpot contract.
var JackpotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PoolAddEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PoolBonusEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderAmount\",\"type\":\"uint256\"}],\"name\":\"addBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOrderAmount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usdtAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"switch_\",\"type\":\"bool\"}],\"name\":\"admin_set_env\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_latest_pool_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"}],\"name\":\"get_users_in_pool\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isUserInPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOrderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolPayoutCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolStatus\",\"outputs\":[{\"internalType\":\"enumIJackpot.PoolStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolUserCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolUsers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// JackpotABI is the input ABI used to generate the binding from.
// Deprecated: Use JackpotMetaData.ABI instead.
var JackpotABI = JackpotMetaData.ABI

// Jackpot is an auto generated Go binding around an Ethereum contract.
type Jackpot struct {
	JackpotCaller     // Read-only binding to the contract
	JackpotTransactor // Write-only binding to the contract
	JackpotFilterer   // Log filterer for contract events
}

// JackpotCaller is an auto generated read-only Go binding around an Ethereum contract.
type JackpotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JackpotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JackpotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JackpotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JackpotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JackpotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JackpotSession struct {
	Contract     *Jackpot          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JackpotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JackpotCallerSession struct {
	Contract *JackpotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// JackpotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JackpotTransactorSession struct {
	Contract     *JackpotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// JackpotRaw is an auto generated low-level Go binding around an Ethereum contract.
type JackpotRaw struct {
	Contract *Jackpot // Generic contract binding to access the raw methods on
}

// JackpotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JackpotCallerRaw struct {
	Contract *JackpotCaller // Generic read-only contract binding to access the raw methods on
}

// JackpotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JackpotTransactorRaw struct {
	Contract *JackpotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJackpot creates a new instance of Jackpot, bound to a specific deployed contract.
func NewJackpot(address common.Address, backend bind.ContractBackend) (*Jackpot, error) {
	contract, err := bindJackpot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jackpot{JackpotCaller: JackpotCaller{contract: contract}, JackpotTransactor: JackpotTransactor{contract: contract}, JackpotFilterer: JackpotFilterer{contract: contract}}, nil
}

// NewJackpotCaller creates a new read-only instance of Jackpot, bound to a specific deployed contract.
func NewJackpotCaller(address common.Address, caller bind.ContractCaller) (*JackpotCaller, error) {
	contract, err := bindJackpot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JackpotCaller{contract: contract}, nil
}

// NewJackpotTransactor creates a new write-only instance of Jackpot, bound to a specific deployed contract.
func NewJackpotTransactor(address common.Address, transactor bind.ContractTransactor) (*JackpotTransactor, error) {
	contract, err := bindJackpot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JackpotTransactor{contract: contract}, nil
}

// NewJackpotFilterer creates a new log filterer instance of Jackpot, bound to a specific deployed contract.
func NewJackpotFilterer(address common.Address, filterer bind.ContractFilterer) (*JackpotFilterer, error) {
	contract, err := bindJackpot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JackpotFilterer{contract: contract}, nil
}

// bindJackpot binds a generic wrapper to an already deployed contract.
func bindJackpot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JackpotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jackpot *JackpotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jackpot.Contract.JackpotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jackpot *JackpotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jackpot.Contract.JackpotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jackpot *JackpotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jackpot.Contract.JackpotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jackpot *JackpotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jackpot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jackpot *JackpotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jackpot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jackpot *JackpotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jackpot.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Jackpot *JackpotCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Jackpot *JackpotSession) ADMINROLE() ([32]byte, error) {
	return _Jackpot.Contract.ADMINROLE(&_Jackpot.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Jackpot *JackpotCallerSession) ADMINROLE() ([32]byte, error) {
	return _Jackpot.Contract.ADMINROLE(&_Jackpot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Jackpot *JackpotCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Jackpot *JackpotSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Jackpot.Contract.DEFAULTADMINROLE(&_Jackpot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Jackpot *JackpotCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Jackpot.Contract.DEFAULTADMINROLE(&_Jackpot.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Jackpot *JackpotCaller) SERVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "SERVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Jackpot *JackpotSession) SERVERROLE() ([32]byte, error) {
	return _Jackpot.Contract.SERVERROLE(&_Jackpot.CallOpts)
}

// SERVERROLE is a free data retrieval call binding the contract method 0x18cf2834.
//
// Solidity: function SERVER_ROLE() view returns(bytes32)
func (_Jackpot *JackpotCallerSession) SERVERROLE() ([32]byte, error) {
	return _Jackpot.Contract.SERVERROLE(&_Jackpot.CallOpts)
}

// AddAmount is a free data retrieval call binding the contract method 0xb9c14577.
//
// Solidity: function addAmount() view returns(uint256)
func (_Jackpot *JackpotCaller) AddAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "addAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddAmount is a free data retrieval call binding the contract method 0xb9c14577.
//
// Solidity: function addAmount() view returns(uint256)
func (_Jackpot *JackpotSession) AddAmount() (*big.Int, error) {
	return _Jackpot.Contract.AddAmount(&_Jackpot.CallOpts)
}

// AddAmount is a free data retrieval call binding the contract method 0xb9c14577.
//
// Solidity: function addAmount() view returns(uint256)
func (_Jackpot *JackpotCallerSession) AddAmount() (*big.Int, error) {
	return _Jackpot.Contract.AddAmount(&_Jackpot.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Jackpot *JackpotCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Jackpot *JackpotSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Jackpot.Contract.GetRoleAdmin(&_Jackpot.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Jackpot *JackpotCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Jackpot.Contract.GetRoleAdmin(&_Jackpot.CallOpts, role)
}

// GetLatestPoolId is a free data retrieval call binding the contract method 0x8f1cc4c3.
//
// Solidity: function get_latest_pool_id() view returns(uint256)
func (_Jackpot *JackpotCaller) GetLatestPoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "get_latest_pool_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPoolId is a free data retrieval call binding the contract method 0x8f1cc4c3.
//
// Solidity: function get_latest_pool_id() view returns(uint256)
func (_Jackpot *JackpotSession) GetLatestPoolId() (*big.Int, error) {
	return _Jackpot.Contract.GetLatestPoolId(&_Jackpot.CallOpts)
}

// GetLatestPoolId is a free data retrieval call binding the contract method 0x8f1cc4c3.
//
// Solidity: function get_latest_pool_id() view returns(uint256)
func (_Jackpot *JackpotCallerSession) GetLatestPoolId() (*big.Int, error) {
	return _Jackpot.Contract.GetLatestPoolId(&_Jackpot.CallOpts)
}

// GetUsersInPool is a free data retrieval call binding the contract method 0xc1cb671b.
//
// Solidity: function get_users_in_pool(uint256 poolId_) view returns(address[])
func (_Jackpot *JackpotCaller) GetUsersInPool(opts *bind.CallOpts, poolId_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "get_users_in_pool", poolId_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersInPool is a free data retrieval call binding the contract method 0xc1cb671b.
//
// Solidity: function get_users_in_pool(uint256 poolId_) view returns(address[])
func (_Jackpot *JackpotSession) GetUsersInPool(poolId_ *big.Int) ([]common.Address, error) {
	return _Jackpot.Contract.GetUsersInPool(&_Jackpot.CallOpts, poolId_)
}

// GetUsersInPool is a free data retrieval call binding the contract method 0xc1cb671b.
//
// Solidity: function get_users_in_pool(uint256 poolId_) view returns(address[])
func (_Jackpot *JackpotCallerSession) GetUsersInPool(poolId_ *big.Int) ([]common.Address, error) {
	return _Jackpot.Contract.GetUsersInPool(&_Jackpot.CallOpts, poolId_)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Jackpot *JackpotCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Jackpot *JackpotSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Jackpot.Contract.HasRole(&_Jackpot.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Jackpot *JackpotCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Jackpot.Contract.HasRole(&_Jackpot.CallOpts, role, account)
}

// IsUserInPool is a free data retrieval call binding the contract method 0x87dbb233.
//
// Solidity: function isUserInPool(uint256 , address ) view returns(bool)
func (_Jackpot *JackpotCaller) IsUserInPool(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "isUserInPool", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserInPool is a free data retrieval call binding the contract method 0x87dbb233.
//
// Solidity: function isUserInPool(uint256 , address ) view returns(bool)
func (_Jackpot *JackpotSession) IsUserInPool(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Jackpot.Contract.IsUserInPool(&_Jackpot.CallOpts, arg0, arg1)
}

// IsUserInPool is a free data retrieval call binding the contract method 0x87dbb233.
//
// Solidity: function isUserInPool(uint256 , address ) view returns(bool)
func (_Jackpot *JackpotCallerSession) IsUserInPool(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Jackpot.Contract.IsUserInPool(&_Jackpot.CallOpts, arg0, arg1)
}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_Jackpot *JackpotCaller) MinOrderAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "minOrderAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_Jackpot *JackpotSession) MinOrderAmount() (*big.Int, error) {
	return _Jackpot.Contract.MinOrderAmount(&_Jackpot.CallOpts)
}

// MinOrderAmount is a free data retrieval call binding the contract method 0x46b62c4a.
//
// Solidity: function minOrderAmount() view returns(uint256)
func (_Jackpot *JackpotCallerSession) MinOrderAmount() (*big.Int, error) {
	return _Jackpot.Contract.MinOrderAmount(&_Jackpot.CallOpts)
}

// PoolAmount is a free data retrieval call binding the contract method 0x02159497.
//
// Solidity: function poolAmount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCaller) PoolAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmount is a free data retrieval call binding the contract method 0x02159497.
//
// Solidity: function poolAmount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotSession) PoolAmount(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolAmount(&_Jackpot.CallOpts, arg0)
}

// PoolAmount is a free data retrieval call binding the contract method 0x02159497.
//
// Solidity: function poolAmount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCallerSession) PoolAmount(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolAmount(&_Jackpot.CallOpts, arg0)
}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCaller) PoolBalance(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 ) view returns(uint256)
func (_Jackpot *JackpotSession) PoolBalance(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolBalance(&_Jackpot.CallOpts, arg0)
}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCallerSession) PoolBalance(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolBalance(&_Jackpot.CallOpts, arg0)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_Jackpot *JackpotCaller) PoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_Jackpot *JackpotSession) PoolId() (*big.Int, error) {
	return _Jackpot.Contract.PoolId(&_Jackpot.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_Jackpot *JackpotCallerSession) PoolId() (*big.Int, error) {
	return _Jackpot.Contract.PoolId(&_Jackpot.CallOpts)
}

// PoolPayoutCount is a free data retrieval call binding the contract method 0x8e6f9ccb.
//
// Solidity: function poolPayoutCount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCaller) PoolPayoutCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolPayoutCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolPayoutCount is a free data retrieval call binding the contract method 0x8e6f9ccb.
//
// Solidity: function poolPayoutCount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotSession) PoolPayoutCount(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolPayoutCount(&_Jackpot.CallOpts, arg0)
}

// PoolPayoutCount is a free data retrieval call binding the contract method 0x8e6f9ccb.
//
// Solidity: function poolPayoutCount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCallerSession) PoolPayoutCount(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolPayoutCount(&_Jackpot.CallOpts, arg0)
}

// PoolStatus is a free data retrieval call binding the contract method 0xf665d10b.
//
// Solidity: function poolStatus(uint256 ) view returns(uint8)
func (_Jackpot *JackpotCaller) PoolStatus(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolStatus is a free data retrieval call binding the contract method 0xf665d10b.
//
// Solidity: function poolStatus(uint256 ) view returns(uint8)
func (_Jackpot *JackpotSession) PoolStatus(arg0 *big.Int) (uint8, error) {
	return _Jackpot.Contract.PoolStatus(&_Jackpot.CallOpts, arg0)
}

// PoolStatus is a free data retrieval call binding the contract method 0xf665d10b.
//
// Solidity: function poolStatus(uint256 ) view returns(uint8)
func (_Jackpot *JackpotCallerSession) PoolStatus(arg0 *big.Int) (uint8, error) {
	return _Jackpot.Contract.PoolStatus(&_Jackpot.CallOpts, arg0)
}

// PoolSwitch is a free data retrieval call binding the contract method 0xa10c27cf.
//
// Solidity: function poolSwitch() view returns(bool)
func (_Jackpot *JackpotCaller) PoolSwitch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolSwitch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolSwitch is a free data retrieval call binding the contract method 0xa10c27cf.
//
// Solidity: function poolSwitch() view returns(bool)
func (_Jackpot *JackpotSession) PoolSwitch() (bool, error) {
	return _Jackpot.Contract.PoolSwitch(&_Jackpot.CallOpts)
}

// PoolSwitch is a free data retrieval call binding the contract method 0xa10c27cf.
//
// Solidity: function poolSwitch() view returns(bool)
func (_Jackpot *JackpotCallerSession) PoolSwitch() (bool, error) {
	return _Jackpot.Contract.PoolSwitch(&_Jackpot.CallOpts)
}

// PoolUserCount is a free data retrieval call binding the contract method 0x8b6dc6e7.
//
// Solidity: function poolUserCount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCaller) PoolUserCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolUserCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolUserCount is a free data retrieval call binding the contract method 0x8b6dc6e7.
//
// Solidity: function poolUserCount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotSession) PoolUserCount(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolUserCount(&_Jackpot.CallOpts, arg0)
}

// PoolUserCount is a free data retrieval call binding the contract method 0x8b6dc6e7.
//
// Solidity: function poolUserCount(uint256 ) view returns(uint256)
func (_Jackpot *JackpotCallerSession) PoolUserCount(arg0 *big.Int) (*big.Int, error) {
	return _Jackpot.Contract.PoolUserCount(&_Jackpot.CallOpts, arg0)
}

// PoolUsers is a free data retrieval call binding the contract method 0x0f2b7cc9.
//
// Solidity: function poolUsers(uint256 , uint256 ) view returns(address)
func (_Jackpot *JackpotCaller) PoolUsers(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "poolUsers", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolUsers is a free data retrieval call binding the contract method 0x0f2b7cc9.
//
// Solidity: function poolUsers(uint256 , uint256 ) view returns(address)
func (_Jackpot *JackpotSession) PoolUsers(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Jackpot.Contract.PoolUsers(&_Jackpot.CallOpts, arg0, arg1)
}

// PoolUsers is a free data retrieval call binding the contract method 0x0f2b7cc9.
//
// Solidity: function poolUsers(uint256 , uint256 ) view returns(address)
func (_Jackpot *JackpotCallerSession) PoolUsers(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Jackpot.Contract.PoolUsers(&_Jackpot.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Jackpot *JackpotCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Jackpot *JackpotSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Jackpot.Contract.SupportsInterface(&_Jackpot.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Jackpot *JackpotCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Jackpot.Contract.SupportsInterface(&_Jackpot.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Jackpot *JackpotCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Jackpot *JackpotSession) Token() (common.Address, error) {
	return _Jackpot.Contract.Token(&_Jackpot.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Jackpot *JackpotCallerSession) Token() (common.Address, error) {
	return _Jackpot.Contract.Token(&_Jackpot.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Jackpot *JackpotCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jackpot.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Jackpot *JackpotSession) Usdt() (common.Address, error) {
	return _Jackpot.Contract.Usdt(&_Jackpot.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Jackpot *JackpotCallerSession) Usdt() (common.Address, error) {
	return _Jackpot.Contract.Usdt(&_Jackpot.CallOpts)
}

// AddBonus is a paid mutator transaction binding the contract method 0xfdddfae8.
//
// Solidity: function addBonus(address user, uint256 orderId, uint256 tokenId, uint256 orderAmount) returns()
func (_Jackpot *JackpotTransactor) AddBonus(opts *bind.TransactOpts, user common.Address, orderId *big.Int, tokenId *big.Int, orderAmount *big.Int) (*types.Transaction, error) {
	return _Jackpot.contract.Transact(opts, "addBonus", user, orderId, tokenId, orderAmount)
}

// AddBonus is a paid mutator transaction binding the contract method 0xfdddfae8.
//
// Solidity: function addBonus(address user, uint256 orderId, uint256 tokenId, uint256 orderAmount) returns()
func (_Jackpot *JackpotSession) AddBonus(user common.Address, orderId *big.Int, tokenId *big.Int, orderAmount *big.Int) (*types.Transaction, error) {
	return _Jackpot.Contract.AddBonus(&_Jackpot.TransactOpts, user, orderId, tokenId, orderAmount)
}

// AddBonus is a paid mutator transaction binding the contract method 0xfdddfae8.
//
// Solidity: function addBonus(address user, uint256 orderId, uint256 tokenId, uint256 orderAmount) returns()
func (_Jackpot *JackpotTransactorSession) AddBonus(user common.Address, orderId *big.Int, tokenId *big.Int, orderAmount *big.Int) (*types.Transaction, error) {
	return _Jackpot.Contract.AddBonus(&_Jackpot.TransactOpts, user, orderId, tokenId, orderAmount)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x75c20f9c.
//
// Solidity: function admin_set_env(uint256 addAmount_, uint256 minOrderAmount_, address usdtAddress_, address tokenAddress_, bool switch_) returns()
func (_Jackpot *JackpotTransactor) AdminSetEnv(opts *bind.TransactOpts, addAmount_ *big.Int, minOrderAmount_ *big.Int, usdtAddress_ common.Address, tokenAddress_ common.Address, switch_ bool) (*types.Transaction, error) {
	return _Jackpot.contract.Transact(opts, "admin_set_env", addAmount_, minOrderAmount_, usdtAddress_, tokenAddress_, switch_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x75c20f9c.
//
// Solidity: function admin_set_env(uint256 addAmount_, uint256 minOrderAmount_, address usdtAddress_, address tokenAddress_, bool switch_) returns()
func (_Jackpot *JackpotSession) AdminSetEnv(addAmount_ *big.Int, minOrderAmount_ *big.Int, usdtAddress_ common.Address, tokenAddress_ common.Address, switch_ bool) (*types.Transaction, error) {
	return _Jackpot.Contract.AdminSetEnv(&_Jackpot.TransactOpts, addAmount_, minOrderAmount_, usdtAddress_, tokenAddress_, switch_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x75c20f9c.
//
// Solidity: function admin_set_env(uint256 addAmount_, uint256 minOrderAmount_, address usdtAddress_, address tokenAddress_, bool switch_) returns()
func (_Jackpot *JackpotTransactorSession) AdminSetEnv(addAmount_ *big.Int, minOrderAmount_ *big.Int, usdtAddress_ common.Address, tokenAddress_ common.Address, switch_ bool) (*types.Transaction, error) {
	return _Jackpot.Contract.AdminSetEnv(&_Jackpot.TransactOpts, addAmount_, minOrderAmount_, usdtAddress_, tokenAddress_, switch_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Jackpot *JackpotTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Jackpot.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Jackpot *JackpotSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Jackpot.Contract.GrantRole(&_Jackpot.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Jackpot *JackpotTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Jackpot.Contract.GrantRole(&_Jackpot.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Jackpot *JackpotTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jackpot.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Jackpot *JackpotSession) Initialize() (*types.Transaction, error) {
	return _Jackpot.Contract.Initialize(&_Jackpot.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Jackpot *JackpotTransactorSession) Initialize() (*types.Transaction, error) {
	return _Jackpot.Contract.Initialize(&_Jackpot.TransactOpts)
}

// Payout is a paid mutator transaction binding the contract method 0xadf58c15.
//
// Solidity: function payout(address user_, uint256 poolId_, uint256 amount_) returns()
func (_Jackpot *JackpotTransactor) Payout(opts *bind.TransactOpts, user_ common.Address, poolId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Jackpot.contract.Transact(opts, "payout", user_, poolId_, amount_)
}

// Payout is a paid mutator transaction binding the contract method 0xadf58c15.
//
// Solidity: function payout(address user_, uint256 poolId_, uint256 amount_) returns()
func (_Jackpot *JackpotSession) Payout(user_ common.Address, poolId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Jackpot.Contract.Payout(&_Jackpot.TransactOpts, user_, poolId_, amount_)
}

// Payout is a paid mutator transaction binding the contract method 0xadf58c15.
//
// Solidity: function payout(address user_, uint256 poolId_, uint256 amount_) returns()
func (_Jackpot *JackpotTransactorSession) Payout(user_ common.Address, poolId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Jackpot.Contract.Payout(&_Jackpot.TransactOpts, user_, poolId_, amount_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Jackpot *JackpotTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Jackpot.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Jackpot *JackpotSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Jackpot.Contract.RenounceRole(&_Jackpot.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Jackpot *JackpotTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Jackpot.Contract.RenounceRole(&_Jackpot.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Jackpot *JackpotTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Jackpot.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Jackpot *JackpotSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Jackpot.Contract.RevokeRole(&_Jackpot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Jackpot *JackpotTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Jackpot.Contract.RevokeRole(&_Jackpot.TransactOpts, role, account)
}

// JackpotInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Jackpot contract.
type JackpotInitializedIterator struct {
	Event *JackpotInitialized // Event containing the contract specifics and raw log

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
func (it *JackpotInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JackpotInitialized)
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
		it.Event = new(JackpotInitialized)
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
func (it *JackpotInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JackpotInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JackpotInitialized represents a Initialized event raised by the Jackpot contract.
type JackpotInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Jackpot *JackpotFilterer) FilterInitialized(opts *bind.FilterOpts) (*JackpotInitializedIterator, error) {

	logs, sub, err := _Jackpot.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JackpotInitializedIterator{contract: _Jackpot.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Jackpot *JackpotFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JackpotInitialized) (event.Subscription, error) {

	logs, sub, err := _Jackpot.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JackpotInitialized)
				if err := _Jackpot.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Jackpot *JackpotFilterer) ParseInitialized(log types.Log) (*JackpotInitialized, error) {
	event := new(JackpotInitialized)
	if err := _Jackpot.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JackpotPoolAddEventIterator is returned from FilterPoolAddEvent and is used to iterate over the raw logs and unpacked data for PoolAddEvent events raised by the Jackpot contract.
type JackpotPoolAddEventIterator struct {
	Event *JackpotPoolAddEvent // Event containing the contract specifics and raw log

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
func (it *JackpotPoolAddEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JackpotPoolAddEvent)
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
		it.Event = new(JackpotPoolAddEvent)
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
func (it *JackpotPoolAddEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JackpotPoolAddEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JackpotPoolAddEvent represents a PoolAddEvent event raised by the Jackpot contract.
type JackpotPoolAddEvent struct {
	User    common.Address
	OrderId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolAddEvent is a free log retrieval operation binding the contract event 0x100682ab2adf1e1f39b36033d5952b9c40526a64e71dbf2859f3f0123568a695.
//
// Solidity: event PoolAddEvent(address indexed user, uint256 indexed orderId, uint256 indexed amount)
func (_Jackpot *JackpotFilterer) FilterPoolAddEvent(opts *bind.FilterOpts, user []common.Address, orderId []*big.Int, amount []*big.Int) (*JackpotPoolAddEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Jackpot.contract.FilterLogs(opts, "PoolAddEvent", userRule, orderIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &JackpotPoolAddEventIterator{contract: _Jackpot.contract, event: "PoolAddEvent", logs: logs, sub: sub}, nil
}

// WatchPoolAddEvent is a free log subscription operation binding the contract event 0x100682ab2adf1e1f39b36033d5952b9c40526a64e71dbf2859f3f0123568a695.
//
// Solidity: event PoolAddEvent(address indexed user, uint256 indexed orderId, uint256 indexed amount)
func (_Jackpot *JackpotFilterer) WatchPoolAddEvent(opts *bind.WatchOpts, sink chan<- *JackpotPoolAddEvent, user []common.Address, orderId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Jackpot.contract.WatchLogs(opts, "PoolAddEvent", userRule, orderIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JackpotPoolAddEvent)
				if err := _Jackpot.contract.UnpackLog(event, "PoolAddEvent", log); err != nil {
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

// ParsePoolAddEvent is a log parse operation binding the contract event 0x100682ab2adf1e1f39b36033d5952b9c40526a64e71dbf2859f3f0123568a695.
//
// Solidity: event PoolAddEvent(address indexed user, uint256 indexed orderId, uint256 indexed amount)
func (_Jackpot *JackpotFilterer) ParsePoolAddEvent(log types.Log) (*JackpotPoolAddEvent, error) {
	event := new(JackpotPoolAddEvent)
	if err := _Jackpot.contract.UnpackLog(event, "PoolAddEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JackpotPoolBonusEventIterator is returned from FilterPoolBonusEvent and is used to iterate over the raw logs and unpacked data for PoolBonusEvent events raised by the Jackpot contract.
type JackpotPoolBonusEventIterator struct {
	Event *JackpotPoolBonusEvent // Event containing the contract specifics and raw log

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
func (it *JackpotPoolBonusEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JackpotPoolBonusEvent)
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
		it.Event = new(JackpotPoolBonusEvent)
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
func (it *JackpotPoolBonusEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JackpotPoolBonusEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JackpotPoolBonusEvent represents a PoolBonusEvent event raised by the Jackpot contract.
type JackpotPoolBonusEvent struct {
	PoolId *big.Int
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolBonusEvent is a free log retrieval operation binding the contract event 0x8e45d6eb8d626f3cbd40d42518d42f4428eea224de0e47af4e054b6f7a56bf6d.
//
// Solidity: event PoolBonusEvent(uint256 indexed poolId, address indexed user, uint256 indexed amount)
func (_Jackpot *JackpotFilterer) FilterPoolBonusEvent(opts *bind.FilterOpts, poolId []*big.Int, user []common.Address, amount []*big.Int) (*JackpotPoolBonusEventIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Jackpot.contract.FilterLogs(opts, "PoolBonusEvent", poolIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &JackpotPoolBonusEventIterator{contract: _Jackpot.contract, event: "PoolBonusEvent", logs: logs, sub: sub}, nil
}

// WatchPoolBonusEvent is a free log subscription operation binding the contract event 0x8e45d6eb8d626f3cbd40d42518d42f4428eea224de0e47af4e054b6f7a56bf6d.
//
// Solidity: event PoolBonusEvent(uint256 indexed poolId, address indexed user, uint256 indexed amount)
func (_Jackpot *JackpotFilterer) WatchPoolBonusEvent(opts *bind.WatchOpts, sink chan<- *JackpotPoolBonusEvent, poolId []*big.Int, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Jackpot.contract.WatchLogs(opts, "PoolBonusEvent", poolIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JackpotPoolBonusEvent)
				if err := _Jackpot.contract.UnpackLog(event, "PoolBonusEvent", log); err != nil {
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

// ParsePoolBonusEvent is a log parse operation binding the contract event 0x8e45d6eb8d626f3cbd40d42518d42f4428eea224de0e47af4e054b6f7a56bf6d.
//
// Solidity: event PoolBonusEvent(uint256 indexed poolId, address indexed user, uint256 indexed amount)
func (_Jackpot *JackpotFilterer) ParsePoolBonusEvent(log types.Log) (*JackpotPoolBonusEvent, error) {
	event := new(JackpotPoolBonusEvent)
	if err := _Jackpot.contract.UnpackLog(event, "PoolBonusEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JackpotRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Jackpot contract.
type JackpotRoleAdminChangedIterator struct {
	Event *JackpotRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *JackpotRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JackpotRoleAdminChanged)
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
		it.Event = new(JackpotRoleAdminChanged)
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
func (it *JackpotRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JackpotRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JackpotRoleAdminChanged represents a RoleAdminChanged event raised by the Jackpot contract.
type JackpotRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Jackpot *JackpotFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*JackpotRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Jackpot.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &JackpotRoleAdminChangedIterator{contract: _Jackpot.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Jackpot *JackpotFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *JackpotRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Jackpot.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JackpotRoleAdminChanged)
				if err := _Jackpot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Jackpot *JackpotFilterer) ParseRoleAdminChanged(log types.Log) (*JackpotRoleAdminChanged, error) {
	event := new(JackpotRoleAdminChanged)
	if err := _Jackpot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JackpotRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Jackpot contract.
type JackpotRoleGrantedIterator struct {
	Event *JackpotRoleGranted // Event containing the contract specifics and raw log

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
func (it *JackpotRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JackpotRoleGranted)
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
		it.Event = new(JackpotRoleGranted)
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
func (it *JackpotRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JackpotRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JackpotRoleGranted represents a RoleGranted event raised by the Jackpot contract.
type JackpotRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Jackpot *JackpotFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*JackpotRoleGrantedIterator, error) {

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

	logs, sub, err := _Jackpot.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &JackpotRoleGrantedIterator{contract: _Jackpot.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Jackpot *JackpotFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *JackpotRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Jackpot.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JackpotRoleGranted)
				if err := _Jackpot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Jackpot *JackpotFilterer) ParseRoleGranted(log types.Log) (*JackpotRoleGranted, error) {
	event := new(JackpotRoleGranted)
	if err := _Jackpot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JackpotRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Jackpot contract.
type JackpotRoleRevokedIterator struct {
	Event *JackpotRoleRevoked // Event containing the contract specifics and raw log

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
func (it *JackpotRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JackpotRoleRevoked)
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
		it.Event = new(JackpotRoleRevoked)
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
func (it *JackpotRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JackpotRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JackpotRoleRevoked represents a RoleRevoked event raised by the Jackpot contract.
type JackpotRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Jackpot *JackpotFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*JackpotRoleRevokedIterator, error) {

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

	logs, sub, err := _Jackpot.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &JackpotRoleRevokedIterator{contract: _Jackpot.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Jackpot *JackpotFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *JackpotRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Jackpot.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JackpotRoleRevoked)
				if err := _Jackpot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Jackpot *JackpotFilterer) ParseRoleRevoked(log types.Log) (*JackpotRoleRevoked, error) {
	event := new(JackpotRoleRevoked)
	if err := _Jackpot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
