package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/1971f281a1f945d8b10c1d8e94c175a9")
	if err != nil {
		fmt.Println(err)
	}
	getHeadByNum(client, nil)
	getHeadByHash(client, "0x699288f67554ea17faca20fcef39687f614bd37c63592c72fdd0a6adb0a6ffba")
	getBlockByNum(client, nil)

}

func getHeadByNum(client *ethclient.Client, num *big.Int) {
	header, err := client.HeaderByNumber(context.Background(), num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v \n", header)
}

func getHeadByHash(client *ethclient.Client, h string) {
	hash, err := client.HeaderByHash(context.Background(), common.HexToHash(h))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v \n", hash)
}

func getBlockByNum(client *ethclient.Client, num *big.Int) {
	number, err := client.BlockByNumber(context.Background(), num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", number)
}
