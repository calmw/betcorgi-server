package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"strings"
)

func DecodeDrawInoutData(inputData string) {
	abiStr := `[
    {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "game_id",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "order_id",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "seed",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "rate",
        "type": "uint256"
      },
      {
        "internalType": "bool",
        "name": "hash_expired",
        "type": "bool"
      }
    ],
    "name": "draw",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
  ]`

	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Println(err)
		return
	}
	parameterBytes, err := contractAbi.Pack(
		"draw",
		big.NewInt(1),
		big.NewInt(1),
		"ascfdv",
		big.NewInt(1),
		true,
	)

	fmt.Println(err, fmt.Sprintf("MethodById %x", parameterBytes[:]))
	inputData = strings.TrimPrefix(inputData, "0x")
	inputBytes := hexutils.HexToBytes(inputData)
	// 解析输入数据
	method, err := contractAbi.MethodById(inputBytes)
	if err != nil {
		log.Println(err, 6666)
		return
	}

	// 获取函数参数
	params := make([]interface{}, len(method.Inputs))
	if params, err = method.Inputs.Unpack(inputBytes[4:]); err != nil {
		log.Println(err)
		return
	}

	gameId := params[0].(*big.Int)
	orderId := params[1].(*big.Int)
	seed := params[2].(string)
	rate := params[3].(*big.Int)
	hashExpired := params[4].(bool)

	log.Println(fmt.Sprintf("gameId:%s, orderId:%s, seed:%s, rate:%s, hashExpired:%v", gameId.String(), orderId.String(), seed, rate.String(), hashExpired))
}

func DecodeBetInoutData(inputData string) {
	abiStr := `[
    {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "game_id",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "token_id",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "hash",
        "type": "bytes32"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "name": "bet",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  }
  ]`

	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Println(err)
		return
	}
	//parameterBytes, err := contractAbi.Pack("bet",
	//	big.NewInt(1),
	//	big.NewInt(1),
	//	big.NewInt(1),
	//	[32]byte(hexutils.HexToBytes("855f7ad8f2ff55074f405edcfc04d9ff996afe7ee638e07a70518c50ac2024f7")),
	//	hexutils.HexToBytes("855f7ad8f2ff55074f405edcfc04d9ff996afe7ee638e07a70518c50ac2024f7"),
	//)
	//
	//fmt.Println(err, fmt.Sprintf("%x", parameterBytes[:4]))
	inputData = strings.TrimPrefix(inputData, "0x")
	inputBytes := hexutils.HexToBytes(inputData)
	// 解析输入数据
	method, err := contractAbi.MethodById(inputBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// 获取函数参数
	params := make([]interface{}, len(method.Inputs))
	if params, err = method.Inputs.Unpack(inputBytes[4:]); err != nil {
		log.Println(err)
		return
	}

	gameId := params[0].(*big.Int)
	tokenId := params[1].(*big.Int)
	amount := params[2].(*big.Int)
	hash := params[3].([32]byte)
	data := params[4].([]byte)

	log.Println(fmt.Sprintf("gameId:%s,tokenId:%s, amount:%s, hash:0x%x, data:0x%x", gameId.String(), tokenId.String(), amount.String(), hash, data))
}
