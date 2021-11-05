package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/1971f281a1f945d8b10c1d8e94c175a9")
	if err != nil {
		fmt.Println(err)
	}
	getNewAccount()
	getBalance(client)
	//checkAccount(client)
}

func getNewAccount() {
	privateKey, err := crypto.GenerateKey() //生成私钥
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey) //转换成字节
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	key := privateKey.Public() //获取公钥
	publicKeyECDSA, ok := key.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
}

func checkAccount(client *ethclient.Client) {
	code, err := client.CodeAt(context.Background(), common.HexToAddress("0xf8838A920c6bB4C57b9d0ffE3193C26Ae3f5ea88"), nil)
	if err != nil {
		fmt.Println(err)
	}

	if len(code) > 0 {
		fmt.Println("是智能合约")
	} else {
		fmt.Println("普通账户")
	}
}

func getBalance(client *ethclient.Client) {
	address := common.HexToAddress("0x05DA6811DAe17E1C1777A60E493bCfE1d3Feb68F")
	balanceAt, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		fmt.Println(err)
	}

	//待处理余额
	at, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(balanceAt)
	fmt.Println(at)
	checkAccount(client)
}
