package blockchain

import (
	erc20 "betcorgi/pkg/binding/token"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"time"
)

type Erc20 struct {
	Cli      *ethclient.Client
	Contract *erc20.Erc20
}

func NewErc20(tokenAddress string) (*Erc20, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	tokenContract, err := erc20.NewErc20(common.HexToAddress(tokenAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Erc20{
		Cli:      cli,
		Contract: tokenContract,
	}, nil
}

func (c Erc20) Approve(to string, amount int64, decimal int64) {
	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		decimalFloat := math.Pow(10, float64(decimal))
		dcl := big.NewInt(int64(decimalFloat))
		amountRes := dcl.Mul(dcl, big.NewInt(amount))
		res, err = c.Contract.Approve(txOpts, common.HexToAddress(to), amountRes)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("approve 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("等待确认成功: %s", res.Hash().String()))

}

func (c Erc20) Transfer(to string, amount int64, decimal int64) {
	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		decimalFloat := math.Pow(10, float64(decimal))
		dcl := big.NewInt(int64(decimalFloat))
		amountRes := dcl.Mul(dcl, big.NewInt(amount))
		res, err = c.Contract.Transfer(txOpts, common.HexToAddress(to), amountRes)
		if err == nil {
			break
		}
		log.Println(err, 2222)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("transfer 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("等待确认成功: %s", res.Hash().String()))

}
