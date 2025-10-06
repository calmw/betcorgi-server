package blockchain

import (
	"betcorgi/pkg/binding/corgi"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

type FeePercentage struct {
	GameId     int64 `json:"game_id"`
	Percentage int64 `json:"percentage"`
}

type Fee struct {
	Cli      *ethclient.Client
	Contract *corgi.Fee
}

func NewFee() (*Fee, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewFee(common.HexToAddress(ChainConfig.FeeContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Fee{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Fee) FeeInit() {
	c.AdminSetFee()
}

func (c Fee) AdminSetFee() {
	var feePercentages []FeePercentage
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     1,
		Percentage: 200,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     2,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     3,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     4,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     5,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     6,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     7,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     8,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     9,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     10,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     11,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     12,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     13,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     14,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     15,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     16,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     17,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     18,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     19,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     20,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     21,
		Percentage: 100,
	})
	feePercentages = append(feePercentages, FeePercentage{
		GameId:     22,
		Percentage: 100,
	})
	for _, feePercentage := range feePercentages {
		c.SetFee(feePercentage)
	}
}

func (c Fee) SetFee(feePercentage FeePercentage) {
	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetFee(txOpts, big.NewInt(feePercentage.GameId), big.NewInt(feePercentage.Percentage))
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(3 * time.Second)
	}

	log.Println(fmt.Sprintf("设置成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}
	log.Println("等待确认成功")
}
