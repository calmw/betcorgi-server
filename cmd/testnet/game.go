package main

import (
	"betcorgi/pkg/blockchain"
	"betcorgi/services"
	"log"
	"math/big"
)

func main() {
	//services.InitBaseTestnetEnv()
	services.InitTestnetEnv()

	//token, err := blockchain.NewToken()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//token.TokenInit()

	gameCategory, err := blockchain.NewGameCategory()
	if err != nil {
		log.Println(err)
		return
	}
	gameCategory.GameCategoryInit()

	order, err := blockchain.NewOrder()
	if err != nil {
		log.Println(err)
		return
	}
	order.OrderInit()
	order.GetLog(big.NewInt(103315800), big.NewInt(103316000))

	autoBet, err := blockchain.NewAutoBet()
	if err != nil {
		log.Println(err)
		return
	}
	autoBet.AutoBetInit()

	game, err := blockchain.NewGame()
	if err != nil {
		log.Println(err)
		return
	}
	//game.AutoBetTest()
	//game.AdminSetBetSwitch()
	//game.GameInit()

	//game.BindParent("0x37225753153de1241Bc8846A1c816453B0Bfa3f1")
	//game.BetTest(16, 1, 1000000)
	//game.BetSingle(4, 1, 1000000, "0x581d44f5a4fd2e33aa631605bfed5bcdf82a14c248c441a2a02d99a363b9c0e4", "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000777b226973496e223a747275652c2270726564696374696f6e52616e6765223a5b323034352c363530335d2c2268617368223a22307835383164343466356134666432653333616136333136303562666564356263646638326131346332343863343431613261303264393961333633623963306534227d000000000000000000")
	erc20, err := blockchain.NewErc20(blockchain.ChainConfig.USDTAddress)
	if err != nil {
		log.Println(err)
		return
	}
	erc20.Approve(blockchain.ChainConfig.GameContractAddress, 1000, 6)
	game.BetSingle(9, 1, big.NewInt(1e9), "0x15702e5c5654563485d99d16fe2ee82cca6bae8518d6f89de15cf833b40026e3", "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000827b227269736b223a226d656469756d222c2273656c6563746564223a5b33332c33342c33352c33362c33372c33382c33392c34305d2c2268617368223a22307831353730326535633536353435363334383564393964313666653265653832636361366261653835313864366638396465313563663833336234303032366533227d000000000000000000000000000000000000000000000000000000000000")

	// rate:  644500
	//gameID:  12
	//orderID:  401
	//randomSeed:  H1,H2,H3,B,B,N2,B,N2,B_60OACmz1ay5xe//A/nZJhsGRivT
	//expire:  false
	//game.Draw(12, 401, "H1,H2,H3,B,B,N2,B,N2,B_60OACmz1ay5xe//A/nZJhsGRivT", 644500, false)

	//erc20, err := blockchain.NewErc20(blockchain.ChainConfig.USDTAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.GameContractAddress, 2, 6)
	//erc20.Transfer(blockchain.ChainConfig.GameContractAddress, 1000000, 6)
	//erc20, err = blockchain.NewErc20(blockchain.ChainConfig.USDCAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.GameContractAddress, 1000000000000000000, 6)
	//erc20.Transfer(blockchain.ChainConfig.GameContractAddress, 1000000, 6)
	//
	//erc20, err = blockchain.NewErc20(blockchain.ChainConfig.USDTAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.JackPotContractAddress, 1000000000000000000, 6)
	//erc20.Transfer(blockchain.ChainConfig.JackPotContractAddress, 1000000, 6)
	//erc20, err = blockchain.NewErc20(blockchain.ChainConfig.USDCAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.JackPotContractAddress, 1000000000000000000, 6)
	//erc20.Transfer(blockchain.ChainConfig.JackPotContractAddress, 1000000, 6)
}
