package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
)

// CalculateBlockHash computes the SHA-256 hash of a block using
// deterministic serialization of block fields excluding the Hash field.
func CalculateBlockHash(b *block.Block) string {
	data := serializeBlock(b)

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}

// serializeBlock converts block fields into a deterministic byte sequence.
//
// Hashing order:
// 1. Index
// 2. Timestamp (Unix seconds)
// 3. Transactions (sender, recipient, amount)
// 4. Previous hash
// 5. Nonce
//
// The Hash field is intentionally excluded because it is the value being calculated.
func serializeBlock(b *block.Block) []byte {
	var builder strings.Builder

	// Block index
	fmt.Fprintf(&builder, "%d|", b.Index)

	// Timestamp
	fmt.Fprintf(&builder, "%d|", b.Timestamp.Unix())

	// Transactions
	for _, tx := range b.Transactions {
		fmt.Fprintf(
			&builder,
			"%s|%s|%.8f|",
			tx.Sender,
			tx.Recipient,
			tx.Amount,
		)
	}

	// Previous block hash
	builder.WriteString(b.PreviousHash)
	builder.WriteString("|")

	// Proof-of-work nonce
	fmt.Fprintf(&builder, "%d", b.Nonce)

	return []byte(builder.String())
}
