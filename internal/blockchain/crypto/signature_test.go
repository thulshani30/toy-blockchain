package crypto_test

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/crypto"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

func TestSignAndVerifyTransaction(t *testing.T) {

	wallet, err := crypto.NewWallet()

	if err != nil {
		t.Fatalf("wallet creation failed: %v", err)
	}

	tx, err := transaction.New(
		"Alice",
		"Bob",
		10,
	)

	if err != nil {
		t.Fatalf("transaction creation failed: %v", err)
	}

	err = crypto.SignTransaction(
		&tx,
		wallet.PrivateKey,
	)

	if err != nil {
		t.Fatalf("signing failed: %v", err)
	}

	if !crypto.VerifyTransaction(tx) {
		t.Error("signature verification failed")
	}
}
