package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/1971f281a1f945d8b10c1d8e94c175a9")
	if err != nil {
		fmt.Println(err)
	}

	getTrxMessageByBlockNum(client)
	getTrxMessageByBlockHash(client)
	getSigleTrx(client)

}

func getTrxMessageByBlockNum(client *ethclient.Client) {
	number, err := client.BlockByNumber(context.Background(), big.NewInt(28153996))
	if err != nil {
		fmt.Println(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for _, transaction := range number.Transactions() {
		//fmt.Println(transaction.Hash().Hex())
		//fmt.Println(transaction.Data())
		//fmt.Println(transaction.Type())
		//fmt.Println(transaction.To().Hex())
		//fmt.Println(transaction.Value())
		//fmt.Println(transaction.Gas())
		//fmt.Println(transaction.GasPrice())
		//fmt.Println(transaction)
		//···
		//更新eip-1559后旧签名types.NewEip155Signer解析不了新版的交易信息
		msg, err := transaction.AsMessage(types.NewLondonSigner(chainID), nil)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("from:", msg.From().Hex())
		}
	}
}

func getTrxMessageByBlockHash(client *ethclient.Client) {
	blockHash := common.HexToHash("0xdd99ec010a306ee50f896286d1ed40d50baa32525a4a21f2ab11acec54069e7b")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		fmt.Println(err)
	}
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tx.Hash().Hex())
	}
}

func getSigleTrx(client *ethclient.Client) {
	txHash := common.HexToHash("0x0224de1ecf8e8c9cd7ee8e3ca7862ddb750294e859ba4b32836317e73ddcbeca")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)
}
