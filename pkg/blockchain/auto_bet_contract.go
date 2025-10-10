package blockchain

import (
	"betcorgi/pkg/binding/corgi"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
)

type AutoBet struct {
	Cli      *ethclient.Client
	Contract *corgi.AutoBet
}

func NewAutoBet() (*AutoBet, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewAutoBet(common.HexToAddress(ChainConfig.AutoBetContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &AutoBet{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c AutoBet) AutoBetInit() {
	fmt.Println("AutoBetInit 开始")
	c.AdminSetEnv(false) // 默认关闭了自动下注
	c.GrantRole()
}

func (c AutoBet) GrantRole() {
	AdminRoleBytes := hexutils.HexToBytes(AdminRole)

	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.GrantRole(txOpts, [32]byte(AdminRoleBytes), common.HexToAddress(ChainConfig.GameContractAddress))
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.GrantRole(txOpts, [32]byte(AdminRoleBytes), common.HexToAddress(ChainConfig.OrderContractAddress))
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func (c AutoBet) AdminSetEnv(autoBetSwitch bool) {
	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(txOpts, common.HexToAddress(ChainConfig.GameContractAddress), autoBetSwitch)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func (c AutoBet) GetNewAutoId() *big.Int {
	var orderId *big.Int
	var err error

	for {
		orderId, err = c.Contract.GetNewAutoId(nil)
		if err == nil {
			log.Println("autoId:", orderId.String())
			break
		}
		log.Println("GetNewAutoId error:", err)
		time.Sleep(3 * time.Second)
	}
	return orderId
}
