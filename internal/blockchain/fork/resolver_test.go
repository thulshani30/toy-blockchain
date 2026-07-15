package fork

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

func TestResolveFork(t *testing.T) {

	current := chain.NewBlockchain()
	competing := chain.NewBlockchain()

	// Add one block to current chain
	tx1, _ := transaction.NewCoinbase("Alice", 100)

	current.AddTransaction(tx1)

	_, err := current.MinePendingTransactions(1, 10)

	if err != nil {
		t.Fatal(err)
	}

	// Add two blocks to competing chain
	tx2, _ := transaction.NewCoinbase("Bob", 100)

	competing.AddTransaction(tx2)

	_, err = competing.MinePendingTransactions(1, 10)

	if err != nil {
		t.Fatal(err)
	}

	tx3, _ := transaction.NewCoinbase("Charlie", 50)

	competing.AddTransaction(tx3)

	_, err = competing.MinePendingTransactions(1, 10)

	if err != nil {
		t.Fatal(err)
	}

	result, err := ResolveFork(
		current,
		competing,
		1,
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(result.Blocks) != len(competing.Blocks) {
		t.Fatal("longer competing chain was not selected")
	}
}
