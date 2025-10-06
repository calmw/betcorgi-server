package blockchain

import (
	"betcorgi/pkg/binding/corgi"
	"context"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"io"
	"log"
	"math/big"
	random "math/rand"
	"strings"
	"time"
)

type Game struct {
	Cli      *ethclient.Client
	Contract *corgi.Game
}

func NewGame() (*Game, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewGame(common.HexToAddress(ChainConfig.GameContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Game{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Game) GameInit() {
	c.AddAccess()
	c.AdminSetEnv(
		ChainConfig.FeeContractAddress,
		ChainConfig.TokenContractAddress,
		ChainConfig.OrderContractAddress,
		ChainConfig.AutoBetContractAddress,
		ChainConfig.JackPotContractAddress,
		ChainConfig.GameCategoryContractAddress,
		true,
	)
}

func (c Game) AddAccess() {
	AdminRole := "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
	AdminRoleBytes := hexutils.HexToBytes(AdminRole)

	var res *types.Transaction

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
		log.Println(err)
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
		res, err = c.Contract.GrantRole(txOpts, [32]byte(AdminRoleBytes), common.HexToAddress(ChainConfig.AutoBetContractAddress))
		if err == nil {
			break
		}
		log.Println(err)
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

func (c Game) AdminSetEnv(feeAddress, tokenAddress, orderAddress, autoBetAddress, jackPotAddress, gameCategoryAddress string, betSwitch bool) {
	var res *types.Transaction
	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(
			txOpts,
			common.HexToAddress(feeAddress),
			common.HexToAddress(tokenAddress),
			common.HexToAddress(orderAddress),
			common.HexToAddress(autoBetAddress),
			common.HexToAddress(jackPotAddress),
			common.HexToAddress(gameCategoryAddress),
			betSwitch,
		)
		if err == nil {
			break
		}
		log.Println(err)
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

func (c Game) Draw(gameId, orderId int64, seed string, rate int64, hashExpired bool) {

	var res *types.Transaction

	for {
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.Draw(
			txOpts,
			big.NewInt(gameId),
			big.NewInt(orderId),
			seed,
			big.NewInt(rate),
			hashExpired,
		)
		if err == nil {
			break
		}
		log.Println("draw error:", err)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			log.Println("开奖：", res.Hash())
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func (c Game) AutoDraw(gameId int64, seed []string, rate []*big.Int, hashExpired []bool) {

	var res *types.Transaction

	for {
		autobet, _ := NewAutoBet()
		autoId := autobet.GetNewAutoId().Int64() - 1
		txOpts, err := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.DrawAutoBet(
			txOpts,
			big.NewInt(gameId),
			big.NewInt(autoId),
			seed,
			rate,
			hashExpired,
		)
		if err == nil {
			break
		}
		log.Println("draw error:", err)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			log.Println("开奖：", res.Hash())
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func (c Game) Bet(gameId, tokenId, amount int64, hash, data string) {
	var err error
	var hashBytes []byte
	var res *types.Transaction
	var txOpts *bind.TransactOpts

	hashBytes = hexutils.HexToBytes(strings.TrimPrefix(hash, "0x"))

	for {
		if tokenId > 0 {
			txOpts, err = GetAuth(c.Cli)
		} else {
			txOpts, err = GetAuthWithValue(c.Cli, big.NewInt(amount))
		}
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(fmt.Sprintf("gameId:%d", big.NewInt(gameId)))
		fmt.Println(fmt.Sprintf("tokenId:%d", big.NewInt(tokenId)))
		fmt.Println(fmt.Sprintf("amount:%d", big.NewInt(amount)))
		fmt.Println(fmt.Sprintf("hash:0x%s", hexutils.BytesToHex(hashBytes)))
		fmt.Println(fmt.Sprintf("data:%s", data))
		res, err = c.Contract.Bet(
			txOpts,
			big.NewInt(gameId),
			big.NewInt(tokenId),
			big.NewInt(amount),
			[32]byte(hashBytes),
			hexutils.HexToBytes(strings.TrimPrefix(data, "0x")),
		)
		if err == nil {
			break
		}
		log.Println("bet error:", err)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			log.Println("下注：", res.Hash())
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func (c Game) AutoBet(gameId, tokenId, initialAmount int64, autoSet []*big.Int, hashArr [][32]byte) {
	var err error
	var encodeData []byte
	var res *types.Transaction
	var txOpts *bind.TransactOpts
	var totalAmount, addAmount *big.Int

	totalAmount = big.NewInt(initialAmount)
	addAmount = autoSet[2]
	if autoSet[2].Cmp(autoSet[3]) == -1 {
		addAmount = autoSet[3]
	}

	for i := 1; i < len(hashArr); i++ {
		add := totalAmount.Add(totalAmount, addAmount)
		totalAmount = totalAmount.Add(totalAmount, add)
	}

	encodeData = hexutils.HexToBytes("7f4e64900901b9af4f077fe4cca678aec779f34650201df47c7300d55d2bfa53")

	for {
		if tokenId > 0 {
			txOpts, err = GetAuth(c.Cli)
		} else {
			txOpts, err = GetAuthWithValue(c.Cli, totalAmount)
		}
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(fmt.Sprintf("gameId:%d", big.NewInt(gameId)))
		fmt.Println(fmt.Sprintf("tokenId:%d", big.NewInt(tokenId)))
		fmt.Println(fmt.Sprintf("initialAmount:%d", big.NewInt(initialAmount)))
		fmt.Println(fmt.Sprintf("amount:%d", totalAmount))
		fmt.Println(fmt.Sprintf("autoSet%v", autoSet))
		fmt.Println(fmt.Sprintf("hashArr 0 0x%v", hexutils.BytesToHex(hashArr[0][:])))
		fmt.Println(fmt.Sprintf("data:0x%s", hexutils.BytesToHex(encodeData)))
		res, err = c.Contract.AutoBet(
			txOpts,
			big.NewInt(gameId),
			big.NewInt(tokenId),
			big.NewInt(initialAmount),
			autoSet,
			hashArr,
			encodeData,
		)
		if err == nil {
			break
		}
		log.Println("bet error:", err)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			log.Println("下注：", res.Hash())
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func generateRandomHexString() [32]byte {
	var randomBytes [32]byte
	_, err := io.ReadFull(rand.Reader, randomBytes[:])
	if err != nil {
		log.Println("Error generating random bytes:", err)
		return [32]byte{}
	}
	return randomBytes
}

func (c Game) BetTest(gameId, tokenId, amount int64, dataHex string) {
	order, _ := NewOrder()
	orderId := order.GetNewOrderId().Int64()
	for i := 0; i < 5; i++ {
		bytes := generateRandomHexString()
		hexStr := hexutils.BytesToHex(bytes[:])
		var guess []*big.Int
		guess = append(guess, big.NewInt(5))
		guess = append(guess, big.NewInt(4))
		c.Bet(
			gameId,
			tokenId,
			amount,
			hexStr,
			dataHex,
		)
		hashExpired := i%2 == 0
		log.Println("hashExpired:", hashExpired)
		c.Draw(
			1,
			orderId,
			"a49807205ce4d355092ef5a8a",
			2,
			hashExpired,
		)
		time.Sleep(10 * time.Second)
		orderId++
	}
}

func (c Game) BetSingle(gameId, tokenId, amount int64, hashStr, dataHex string) {
	order, _ := NewOrder()
	orderId := order.GetNewOrderId().Int64()
	c.Bet(
		gameId,
		tokenId,
		amount,
		hashStr,
		dataHex,
	)
	r := random.New(random.NewSource(time.Now().UnixNano()))

	// 生成一个 0 到 99 的随机整数
	n := r.Intn(100)
	fmt.Println("随机数:", n)
	hashExpired := n%2 == 0
	log.Println("hashExpired:", hashExpired)
	c.Draw(
		1,
		orderId,
		hashStr,
		2,
		hashExpired,
	)
	time.Sleep(10 * time.Second)
}

func (c Game) AutoBetTest() {

	for i := 0; i < 5; i++ {
		c.AutoBet(
			1,
			1,
			20*1e6,
			[]*big.Int{big.NewInt(2 * 1e10), big.NewInt(40 * 1e10), big.NewInt(5), big.NewInt(3)},
			[][32]byte{generateRandomHexString(), generateRandomHexString(), generateRandomHexString()},
		)
		time.Sleep(5 * time.Second)

		c.AutoDraw(
			1,
			[]string{"a", "b", "c"},
			[]*big.Int{big.NewInt(2), big.NewInt(2), big.NewInt(2)},
			[]bool{false, false, false},
		)

		time.Sleep(10 * time.Second)
	}
}

func (c Game) AdminSetBetSwitch() {

	var res *types.Transaction
	//loopStart:
	for i := 0; i < 22; i++ {
		for {
			txOpts, err := GetAuth(c.Cli)
			if err != nil {
				log.Println(err)
				return
			}
			res, err = c.Contract.AdminSetBetSwitch(
				txOpts,
				big.NewInt(int64(i)),
				true)
			if err == nil {
				break
			}
			log.Println("admin set bet switch error:", err)
			time.Sleep(3 * time.Second)
		}
		log.Println(fmt.Sprintf("成功"))
		for {
			receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
			if err == nil && receipt.Status == 1 {
				log.Println(fmt.Sprintf("确认,游戏ID %d, hash：%s", i, res.Hash()))
				break
			}
			time.Sleep(time.Second * 2)
		}

		log.Println(fmt.Sprintf("确认成功"))
	}
}
