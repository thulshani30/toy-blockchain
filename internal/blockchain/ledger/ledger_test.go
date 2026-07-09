package ledger

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

func TestCoinbaseTransaction(t *testing.T) {

	l := NewLedger()

	tx := transaction.Transaction{
		Sender:    transaction.CoinbaseAccount,
		Recipient: "Alice",
		Amount:    100,
	}

	err := l.ApplyTransaction(tx)

	if err != nil {
		t.Fatalf("coinbase transaction failed: %v", err)
	}

	balance := l.GetBalance("Alice")

	if balance != 100 {
		t.Fatalf("expected balance 100, got %f", balance)
	}
}

func TestTransferTransaction(t *testing.T) {

	l := NewLedger()

	coinbase := transaction.Transaction{
		Sender:    transaction.CoinbaseAccount,
		Recipient: "Alice",
		Amount:    100,
	}

	err := l.ApplyTransaction(coinbase)

	if err != nil {
		t.Fatalf("coinbase failed: %v", err)
	}

	tx := transaction.Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    30,
	}

	err = l.ApplyTransaction(tx)

	if err != nil {
		t.Fatalf("transfer failed: %v", err)
	}

	if l.GetBalance("Alice") != 70 {
		t.Fatalf("expected Alice balance 70, got %f",
			l.GetBalance("Alice"))
	}

	if l.GetBalance("Bob") != 30 {
		t.Fatalf("expected Bob balance 30, got %f",
			l.GetBalance("Bob"))
	}
}

func TestInsufficientBalance(t *testing.T) {

	l := NewLedger()

	tx := transaction.Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    50,
	}

	err := l.ApplyTransaction(tx)

	if err == nil {
		t.Fatal("expected insufficient balance error")
	}
}
