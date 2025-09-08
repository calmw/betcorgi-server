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
		GameContractAddress:         "0x29455703c87377e9b5ae64617DD6C5a73DB0FdfA",
		GameCategoryContractAddress: "0xd774870002a35F43DCd62a6d132B362292d186B3",
		FeeContractAddress:          "0x9DaDe402227ba64fA97C9668d2DA21Cdb5C94dAE",
		OrderContractAddress:        "0x5D983e8482CB1426407102E07188851335f3418a",
		TokenContractAddress:        "0x881047e2520eA6B94A57d5987651E7f8978e50E5",
		AutoBetContractAddress:      "0x12B3FF826FFF39266690A7C9496535BdEd7f6705",
		JackPotContractAddress:      "0x078dBf9C5BFE1B9ba140218701CCbAaD58326885",
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
		FeeContractAddress:          "0x61b4aDd1103253c58fe1505E832e460d64265AF4",
		OrderContractAddress:        "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		TokenContractAddress:        "0xC1F4898261Fa0420ec457b48554F719Bcde9c0AD",
		AutoBetContractAddress:      "0xb48Fc516822c845fA825971d2ab1008a9C215D1A",
		JackPotContractAddress:      "0xE267B6fC13378474337eDAbFB47C1865322aA689",
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
