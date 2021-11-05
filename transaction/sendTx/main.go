package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/1971f281a1f945d8b10c1d8e94c175a9")
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err := crypto.HexToECDSA("d12d8480f44007ec0b88cf7a1c6ccf7632657463847f200333ec0a2e1caa8d62")
	if err != nil {
		fmt.Println(err)
	}

	//通过私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//获取账号nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	address := common.HexToAddress("0x05DA6811DAe17E1C1777A60E493bCfE1d3Feb68F")
	//获取gasprice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	tx := types.LegacyTx{
		Nonce:    nonce,
		To:       &address,
		Value:    big.NewInt(1000), //wei
		GasPrice: gasPrice,
		Gas:      30000,
		Data:     nil,
	}
	//newTx := types.NewTx(&tx)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	//signTx, err := types.SignTx(newTx, types.NewEIP155Signer(chainID), privateKey)

	newTx, err := types.SignNewTx(privateKey, types.NewEIP155Signer(chainID), &tx)
	client.SendTransaction(context.Background(), newTx)
	fmt.Println(newTx.Hash().Hex())
}
