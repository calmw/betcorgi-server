package services

import (
	"betcorgi/pkg/blockchain"
	"os"
)

func InitTestnetEnv() {
	key := os.Getenv("BET_CORGI")

	blockchain.ChainConfig = blockchain.ChainConfigs{
		ChainId:                     421614,
		RPC:                         "https://arbitrum-sepolia.gateway.tenderly.co",
		GameContractAddress:         "0x0f618aEad952d298444cFcF8683a7f3Df63B3598",
		GameCategoryContractAddress: "0x9fD4951e888C5de9D75320704E868296Eeb0adE7",
		OrderContractAddress:        "0xE12f247e633Fb773c7387a3266F1f9F287E1a069",
		TokenContractAddress:        "0xd6c44AA3A5503479E2C1248e462aD2eF9EB7Ac19",
		AutoBetContractAddress:      "0x32E3BC35061e6cC8a6072509b51cd254cB9E5AC5",
		USDTAddress:                 "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		USDCAddress:                 "0x70EF97cdFdafa043925225eA4EEeE595DAF22542",
		PrivateKey:                  key,
	}
}

func InitBaseTestnetEnv() {
	key := os.Getenv("BET_CORGI")

	blockchain.ChainConfig = blockchain.ChainConfigs{
		ChainId:                     84532,
		RPC:                         "https://base-sepolia-rpc.publicnode.com",
		GameContractAddress:         "0xd6FE88ae621E60410615E12b174C04f0d25978aC",
		GameCategoryContractAddress: "0x4118f5b3EfDf8CeAA5A9662E5548642e6E222a7a",
		OrderContractAddress:        "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		TokenContractAddress:        "0xC1F4898261Fa0420ec457b48554F719Bcde9c0AD",
		AutoBetContractAddress:      "0xb48Fc516822c845fA825971d2ab1008a9C215D1A",
		USDTAddress:                 "0x6fC29A105AD96F859B402F4c9952622A2aAF6d05",
		USDCAddress:                 "0x2186C83A1866FA0Cf1945501A90C1fa224f33b62",
		PrivateKey:                  key,
	}
}

func InitMainnetEnv() {
	key := os.Getenv("BET_CORGI")
	blockchain.ChainConfig = blockchain.ChainConfigs{
		ChainId:     698,
		RPC:         "https://rpc.matchain.io",
		USDTAddress: "0x14f31776De4155398F1a12eE9beb0E832Bd845e9", // MUSDT代币合约地址
		PrivateKey:  key,
	}
}
