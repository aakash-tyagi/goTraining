package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	_ = client

	address := common.HexToAddress("0xb53cC19aD713e00cB71542d064215784c908D387")

	fmt.Println(address)
	// blockNumber := big.NewInt(5532993)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(pendingBalance)

	fbalance.SetString(pendingBalance.String())
	ethValue = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
}
