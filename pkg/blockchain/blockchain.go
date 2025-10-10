package blockchain

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	LayoutTime = "2006-01-02 15:04:05"
	AdminRole  = "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
	ServerRole = "a8a7bc421f721cb936ea99efdad79237e6ee0b871a2a08cf648691f9584cdc77"
	ManageRole = "a8a7bc421f721cb936ea99efdad79237e6ee0b871a2a08cf648691f9584cdc77"
)

type ChainConfigs struct {
	ChainId                     int64
	RPC                         string
	GameContractAddress         string
	GameCategoryContractAddress string
	OrderContractAddress        string
	TokenContractAddress        string
	AutoBetContractAddress      string
	USDTAddress                 string
	USDCAddress                 string
	PrivateKey                  string
}

var ChainConfig ChainConfigs

func Client(c ChainConfigs) (error, *ethclient.Client) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		log.Fatal("dail failed")
	}
	return nil, client
}

func GetAuth(cli *ethclient.Client) (*bind.TransactOpts, error) {
	//privateKeyEcdsa, err := crypto.HexToECDSA("3741deeaad3b700a09f6ddcf9638c7d2393da6b391ba8f5e53332d77864738a6")
	privateKeyEcdsa, err := crypto.HexToECDSA(ChainConfig.PrivateKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	publicKey := crypto.PubkeyToAddress(privateKeyEcdsa.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(ChainConfig.ChainId))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	head, err := cli.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	gasTipCap := new(big.Int)
	maxFeePerGas := new(big.Int)
	if head.BaseFee != nil {
		gasTipCap, err = cli.SuggestGasTipCap(context.Background())
		if err != nil {
			return nil, err
		}
		// Both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) cannot be specified: https://github.com/ethereum/go-ethereum/blob/95bbd46eabc5d95d9fb2108ec232dd62df2f44ab/accounts/abi/bind/base.go#L254
		gasPrice = nil
		maxFeePerGas = new(big.Int).Add(head.BaseFee, gasTipCap)
		maxFeePerGas = new(big.Int).Mul(maxFeePerGas, big.NewInt(7))
		maxFeePerGas = new(big.Int).Div(maxFeePerGas, big.NewInt(5))
		//fmt.Println(fmt.Sprintf("BaseFee:%v,GasTipCap:%v，maxFeePerGas：%v", head.BaseFee.Int64(), gasTipCap.Int64(), maxFeePerGas.Int64()))
	}

	return &bind.TransactOpts{
		From:      auth.From,
		Nonce:     big.NewInt(int64(nonce)),
		Signer:    auth.Signer,
		Value:     big.NewInt(0),
		GasPrice:  gasPrice,
		GasFeeCap: maxFeePerGas,
		GasTipCap: gasTipCap,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false,
	}, nil
}

func GetAuthWithValue(cli *ethclient.Client, value *big.Int) (*bind.TransactOpts, error) {
	privateKeyEcdsa, err := crypto.HexToECDSA(ChainConfig.PrivateKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	publicKey := crypto.PubkeyToAddress(privateKeyEcdsa.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(ChainConfig.ChainId))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	head, err := cli.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	gasTipCap := new(big.Int)
	maxFeePerGas := new(big.Int)
	if head.BaseFee != nil {
		gasTipCap, err = cli.SuggestGasTipCap(context.Background())
		if err != nil {
			return nil, err
		}
		// Both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) cannot be specified: https://github.com/ethereum/go-ethereum/blob/95bbd46eabc5d95d9fb2108ec232dd62df2f44ab/accounts/abi/bind/base.go#L254
		gasPrice = nil
		maxFeePerGas = new(big.Int).Add(head.BaseFee, gasTipCap)
		maxFeePerGas = new(big.Int).Mul(maxFeePerGas, big.NewInt(7))
		maxFeePerGas = new(big.Int).Div(maxFeePerGas, big.NewInt(5))
		fmt.Println(fmt.Sprintf("BaseFee:%v,GasTipCap:%v，maxFeePerGas：%v", head.BaseFee.Int64(), gasTipCap.Int64(), maxFeePerGas.Int64()))
	}

	return &bind.TransactOpts{
		From:      auth.From,
		Nonce:     big.NewInt(int64(nonce)),
		Signer:    auth.Signer,
		Value:     value,
		GasPrice:  gasPrice,
		GasFeeCap: maxFeePerGas,
		GasTipCap: gasTipCap,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false,
	}, nil
}

func AbiEncode(opts []interface{}) ([]byte, error) {
	abiStr, err := generateAbi(opts)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(bytes.NewBufferString(abiStr))
	if err != nil {
		return nil, err
	}

	callData, err := parsedABI.Pack("draw", opts...)
	if err != nil {
		return nil, err
	}

	return callData[4:], nil
}

func generateAbi(data []interface{}) (string, error) {
	var inputs string
	for i, d := range data {
		t, err := parseType(d)
		if err != nil {
			return "", err
		}
		inputs += fmt.Sprintf(`{
                "internalType": "%s",
                "name": "parameter%d",
                "type": "%s"
            },`, t, i, t)
	}
	inputs = strings.TrimRight(inputs, ",")

	return fmt.Sprintf(`
[
    {
        "inputs": [
            %s
        ],
        "name": "draw",
        "stateMutability": "pure",
        "type": "function"
    }
]`, inputs), nil

}

func parseType(data interface{}) (string, error) {
	t := reflect.TypeOf(data)
	switch t.String() {
	case "*big.Int":
		return "uint256", nil
	case "[]*big.Int":
		return "uint256[]", nil
	default:
		return "", errors.New("opts type error")
	}
}
