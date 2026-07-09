package hashing

import (
	"testing"
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

func TestDeterministicHash(t *testing.T) {

	b := block.Block{
		Index:     1,
		Timestamp: time.Unix(1000, 0),
		Transactions: []transaction.Transaction{
			{
				Sender:    "Alice",
				Recipient: "Bob",
				Amount:    50,
			},
		},
		PreviousHash: "previous_hash",
		Nonce:        123,
	}

	hash1 := CalculateBlockHash(&b)
	hash2 := CalculateBlockHash(&b)

	if hash1 != hash2 {
		t.Fatal("hash should be deterministic")
	}
}

func TestHashChangesWhenBlockChanges(t *testing.T) {

	b := block.Block{
		Index:        1,
		Timestamp:    time.Unix(1000, 0),
		Transactions: []transaction.Transaction{},
		PreviousHash: "previous_hash",
		Nonce:        1,
	}

	hash1 := CalculateBlockHash(&b)

	b.Nonce = 2

	hash2 := CalculateBlockHash(&b)

	if hash1 == hash2 {
		t.Fatal("hash should change when block data changes")
	}
}
