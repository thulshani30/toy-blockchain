package chain

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

func TestGenesisBlock(t *testing.T) {

	bc := NewBlockchain()

	if len(bc.Blocks) != 1 {
		t.Fatalf("expected 1 genesis block, got %d", len(bc.Blocks))
	}

	genesis := bc.Blocks[0]

	if genesis.Index != 0 {
		t.Fatalf("expected genesis index 0, got %d", genesis.Index)
	}

	if genesis.PreviousHash != GenesisPreviousHash {
		t.Fatalf("incorrect genesis previous hash")
	}

	if genesis.Hash == "" {
		t.Fatal("genesis hash should not be empty")
	}
}

func TestAddTransaction(t *testing.T) {

	bc := NewBlockchain()

	tx := transaction.Transaction{
		Sender:    transaction.CoinbaseAccount,
		Recipient: "Alice",
		Amount:    100,
	}

	err := bc.AddTransaction(tx)

	if err != nil {
		t.Fatalf("failed to add transaction: %v", err)
	}

	if len(bc.PendingTransactions) != 1 {
		t.Fatalf(
			"expected 1 pending transaction, got %d",
			len(bc.PendingTransactions),
		)
	}
}

func TestRejectInsufficientBalance(t *testing.T) {

	bc := NewBlockchain()

	tx := transaction.Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    50,
	}

	err := bc.AddTransaction(tx)

	if err == nil {
		t.Fatal("expected insufficient balance error")
	}
}
