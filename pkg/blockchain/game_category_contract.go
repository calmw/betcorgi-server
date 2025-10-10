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
)

type Category struct {
	GameId   int64  `json:"game_id"`
	GameName string `json:"game_name"`
}

type GameCategory struct {
	Cli      *ethclient.Client
	Contract *corgi.GameCategory
}

func NewGameCategory() (*GameCategory, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewGameCategory(common.HexToAddress(ChainConfig.GameCategoryContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &GameCategory{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c GameCategory) GameCategoryInit() {
	c.AdminSetGameCategory()
}

func (c GameCategory) AdminSetGameCategory() {
	var categories []Category
	categories = append(categories, Category{
		GameId:   1,
		GameName: "dice",
	})
	categories = append(categories, Category{
		GameId:   2,
		GameName: "spacedice",
	})
	categories = append(categories, Category{
		GameId:   3,
		GameName: "limbo",
	})
	categories = append(categories, Category{
		GameId:   4,
		GameName: "coinflip",
	})
	categories = append(categories, Category{
		GameId:   5,
		GameName: "keno",
	})
	categories = append(categories, Category{
		GameId:   6,
		GameName: "plinko",
	})
	categories = append(categories, Category{
		GameId:   7,
		GameName: "crash",
	})
	categories = append(categories, Category{
		GameId:   8,
		GameName: "mines",
	})
	categories = append(categories, Category{
		GameId:   9,
		GameName: "ring",
	})
	categories = append(categories, Category{
		GameId:   10,
		GameName: "cryptos",
	})
	categories = append(categories, Category{
		GameId:   11,
		GameName: "tower",
	})
	categories = append(categories, Category{
		GameId:   12,
		GameName: "corgipharaoh",
	})
	categories = append(categories, Category{
		GameId:   13,
		GameName: "hilo",
	})
	categories = append(categories, Category{
		GameId:   14,
		GameName: "stairs",
	})
	categories = append(categories, Category{
		GameId:   15,
		GameName: "corgiwild",
	})
	categories = append(categories, Category{
		GameId:   16,
		GameName: "circle",
	})
	categories = append(categories, Category{
		GameId:   17,
		GameName: "triple",
	})
	categories = append(categories, Category{
		GameId:   19,
		GameName: "roulette",
	})
	categories = append(categories, Category{
		GameId:   20,
		GameName: "blackjack",
	})
	categories = append(categories, Category{
		GameId:   21,
		GameName: "futures",
	})
	for _, category := range categories {
		c.SetGameCategory(category)
	}
}

func (c GameCategory) SetGameCategory(category Category) {
	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetGameCategory(txOpts, big.NewInt(category.GameId), category.GameName)
		if err == nil {
			break
		}
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
