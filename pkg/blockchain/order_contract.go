package blockchain

import (
	"betcorgi/pkg/binding/corgi"
	"betcorgi/pkg/event"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"strings"
	"time"
)

type Order struct {
	Cli      *ethclient.Client
	Contract *corgi.Order
}

func NewOrder() (*Order, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewOrder(common.HexToAddress(ChainConfig.OrderContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Order{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Order) AdminSetEnv() {
	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(
			txOpts,
			common.HexToAddress(ChainConfig.FeeContractAddress),
			common.HexToAddress(ChainConfig.GameContractAddress),
			common.HexToAddress(ChainConfig.AutoBetContractAddress),
			common.HexToAddress(ChainConfig.GameCategoryContractAddress),
		)
		if err == nil {
			break
		}
		log.Println(err, 3333)
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

func (c Order) GetLog(startHeight, endHeight *big.Int) {
	query := event.BuildQuery(common.HexToAddress(ChainConfig.OrderContractAddress), event.BetEvent.EventSignature, startHeight, endHeight)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	logs, err := c.Cli.FilterLogs(ctx, query)
	if err != nil {
		log.Println(err)
		return
	}
	abi, err := corgi.OrderMetaData.GetAbi()
	if err != nil {
		log.Println(err)
		return
	}
	for _, eLog := range logs {
		orderId := eLog.Topics[1].Big().Int64()
		user := strings.ToLower(common.HexToAddress(eLog.Topics[2].String()).String())
		amount := eLog.Topics[3].Big().String()
		eventData, err := abi.Unpack(event.BetEvent.EventName, eLog.Data)
		if err != nil {
			log.Println(err)
			return
		}
		gameId := eventData[0].(*big.Int).Int64()
		periodNumber := eventData[1].(*big.Int).Int64()
		guess := eventData[3].([]*big.Int)
		tokenId := eventData[4].(*big.Int).Int64()
		hashBytes32 := eventData[2].([32]byte)

		hash := hexutils.BytesToHex(hashBytes32[:32])
		hashHexStr := fmt.Sprintf("0x%s", hash)
		fmt.Println(fmt.Sprintf("startHeight:%d endHeight:%d height:%d gameId:%d periodNumber:%d orderId:%d hash:%s user:%s guess:%v tokenId:%d amount:%s ",
			startHeight, endHeight,
			eLog.BlockNumber,
			gameId,
			periodNumber,
			orderId,
			hashHexStr,
			user,
			guess,
			tokenId,
			amount,
		))
	}
}

func (c Order) OrderInit() {
	c.AdminSetEnv()
	c.AddAccess()
}

func (c Order) AddAccess() {
	AdminRole := "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
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
}

func (c Order) GetNewOrderId() *big.Int {
	var orderId *big.Int
	var err error

	for {
		orderId, err = c.Contract.GetNewOrderId(nil)
		if err == nil {
			log.Println("orderId:", orderId.String())
			break
		}
		log.Println("GetNewOrderId error:", err)
		time.Sleep(3 * time.Second)
	}
	return orderId
}
