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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
)

type TokenInfo struct {
	GameId       int64    `json:"game_id"`
	TokenId      int64    `json:"token_id"`
	TokenAddress string   `json:"token_address"`
	MinBetAmount *big.Int `json:"min_bet_amount"`
	MaxBetAmount *big.Int `json:"max_bet_amount"`
}

type Token struct {
	Cli      *ethclient.Client
	Contract *corgi.Token
}

func NewToken() (*Token, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewToken(common.HexToAddress(ChainConfig.TokenContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Token{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Token) TokenInit() {
	prvKey, err := crypto.HexToECDSA(ChainConfig.PrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	address := crypto.PubkeyToAddress(prvKey.PublicKey)
	c.AddAccess(AdminRole, address.String())
	c.AdminSetTokenExchangeRate()
	c.AdminSetToken()
}

func (c Token) AddAccess(role, account string) {
	roleBytes := hexutils.HexToBytes(role)

	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.GrantRole(txOpts, [32]byte(roleBytes), common.HexToAddress(account))
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

func (c Token) AdminSetToken() {
	var tokenInfos []TokenInfo
	// USDT/USDC
	b1u := big.NewInt(1e6)                                        // 1
	b1000u := big.NewInt(1e9)                                     // 1000
	b2500u := new(big.Int).Mul(big.NewInt(2500), big.NewInt(1e6)) // 2500
	b5000u := new(big.Int).Mul(big.NewInt(5000), big.NewInt(1e6)) // 5000
	b8000u := new(big.Int).Mul(big.NewInt(8000), big.NewInt(1e6)) // 8000
	// ETH
	b18 := big.NewInt(1e18)
	bddd1 := big.NewInt(1e15) //  0.001
	//bdd8 := new(big.Int).Mul(big.NewInt(8), big.NewInt(1e16))  // 0.08
	bd6 := new(big.Int).Mul(big.NewInt(6), big.NewInt(1e17))   // 0.6
	bd25 := new(big.Int).Mul(big.NewInt(25), big.NewInt(1e16)) // 0.25
	b1d6 := new(big.Int).Mul(big.NewInt(16), big.NewInt(1e17)) // 1.6
	// ETH
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       1,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       2,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       3,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // coinflip
		GameId:       4,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: b18,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       5,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       6,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       7,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       8,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       9,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       10,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       11,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // CorgiPharaoh
		GameId:       12,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: b1d6,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       13,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       14,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       15,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       16,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       17,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // roulette
		GameId:       19,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd6,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       20,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       21,
		TokenId:      0,
		TokenAddress: "0x0000000000000000000000000000000000000000",
		MinBetAmount: bddd1,
		MaxBetAmount: bd25,
	})
	// USDT -------------------------------------------------------------------------------------------------------------------------------
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       1,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       2,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       3,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // coinflip
		GameId:       4,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b5000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       5,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: new(big.Int).Mul(b1u, big.NewInt(1000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       6,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       7,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       8,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       9,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       10,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       11,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // CorgiPharaoh
		GameId:       12,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b8000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       13,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       14,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       15,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       16,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       17,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // roulette
		GameId:       19,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b2500u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       20,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       21,
		TokenId:      1,
		TokenAddress: ChainConfig.USDTAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	// USDC -------------------------------------------------------------------------------------------------------------------------------
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       1,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       2,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       3,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // coinflip
		GameId:       4,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b5000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       5,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       6,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       7,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       8,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       9,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       10,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       11,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // CorgiPharaoh
		GameId:       12,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b8000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       13,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       14,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       15,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       16,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       17,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{ // roulette
		GameId:       19,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b2500u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       20,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       21,
		TokenId:      2,
		TokenAddress: ChainConfig.USDCAddress,
		MinBetAmount: b1u,
		MaxBetAmount: b1000u,
	})
	for _, tokenInfo := range tokenInfos {
		//if tokenInfo.GameId != 6 {
		//	continue
		//}
		//fmt.Println(tokenInfo)
		c.SetToken(tokenInfo)
	}
}

func (c Token) SetToken(tokenInfo TokenInfo) {
	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(
			big.NewInt(tokenInfo.GameId),
			big.NewInt(tokenInfo.TokenId),
			common.HexToAddress(tokenInfo.TokenAddress),
			tokenInfo.MinBetAmount,
			tokenInfo.MaxBetAmount,
		)
		res, err = c.Contract.AdminSetToken(
			txOpts,
			big.NewInt(tokenInfo.GameId),
			big.NewInt(tokenInfo.TokenId),
			common.HexToAddress(tokenInfo.TokenAddress),
			tokenInfo.MinBetAmount,
			tokenInfo.MaxBetAmount,
		)
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
	log.Println("等待确认成功", tokenInfo.GameId, tokenInfo.TokenId, res.Hash())
}

func (c Token) AdminSetTokenExchangeRate() {
	var res *types.Transaction

	// ETH
	ethPrice := new(big.Int).Mul(big.NewInt(4618), big.NewInt(1e18))
	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetTokenExchangeRate(
			txOpts,
			big.NewInt(0),
			new(big.Int).Mul(ethPrice, big.NewInt(1e4)),
		)
		if err == nil {
			break
		}
		fmt.Println("AdminSetTokenExchangeRate error", err)
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

	// USDT
	usdtPrice := big.NewInt(1e6)
	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetTokenExchangeRate(
			txOpts,
			big.NewInt(1),
			new(big.Int).Mul(usdtPrice, big.NewInt(1e4)),
		)
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

	// USDC
	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetTokenExchangeRate(
			txOpts,
			big.NewInt(2),
			new(big.Int).Mul(usdtPrice, big.NewInt(1e4)),
		)
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
