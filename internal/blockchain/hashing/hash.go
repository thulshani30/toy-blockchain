package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
)

// CalculateBlockHash computes the SHA-256 hash of a block using
// a deterministic serialization of its fields.
func CalculateBlockHash(b *block.Block) (string, error) {
	data, err := serializeBlock(b)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

// serializeBlock converts a block into a deterministic byte sequence.
// The Hash field is intentionally excluded.
func serializeBlock(b *block.Block) ([]byte, error) {
	var data string

	// 1. Index
	data += fmt.Sprintf("%d|", b.Index)

	// 2. Timestamp (Unix for determinism)
	data += fmt.Sprintf("%d|", b.Timestamp.Unix())

	// 3. Transactions
	for _, tx := range b.Transactions {
		data += fmt.Sprintf("%s|%s|%f|",
			tx.Sender,
			tx.Recipient,
			tx.Amount,
		)
	}

	// 4. Previous Hash
	data += b.PreviousHash + "|"

	// 5. Nonce
	data += fmt.Sprintf("%d", b.Nonce)

	return []byte(data), nil
}
