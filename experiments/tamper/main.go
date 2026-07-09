package main

import (
	"fmt"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/validation"
)

func main() {

	bc := chain.NewBlockchain()

	tx, _ := transaction.NewCoinbase(
		"Alice",
		100,
	)

	bc.AddTransaction(tx)

	_, err := bc.MinePendingTransactions(2, 10)

	if err != nil {
		panic(err)
	}

	fmt.Println("Before tampering:")

	err = validation.ValidateChain(bc, 2)

	fmt.Println(err == nil)

	// Tamper with transaction
	bc.Blocks[1].Transactions[0].Amount = 500

	fmt.Println("\nAfter tampering:")

	err = validation.ValidateChain(bc, 2)

	if err != nil {
		fmt.Println("Detected:", err)
	}
}
