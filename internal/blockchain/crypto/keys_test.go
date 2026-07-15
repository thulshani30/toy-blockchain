package crypto_test

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/crypto"
)

func TestGenerateKeyPair(t *testing.T) {

	privateKey, publicKey, err := crypto.GenerateKeyPair()

	if err != nil {
		t.Fatalf("failed generating keys: %v", err)
	}

	if privateKey == nil {
		t.Fatal("private key is nil")
	}

	if publicKey == nil {
		t.Fatal("public key is nil")
	}
}
