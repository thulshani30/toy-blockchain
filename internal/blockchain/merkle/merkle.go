package merkle

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// CalculateMerkleRoot calculates the Merkle root of transactions.
func CalculateMerkleRoot(transactions []transaction.Transaction) string {

	if len(transactions) == 0 {
		return hashData("")
	}

	var hashes []string

	for _, tx := range transactions {

		data := fmt.Sprintf(
			"%s%s%.8f",
			tx.Sender,
			tx.Recipient,
			tx.Amount,
		)

		hashes = append(hashes, hashData(data))
	}

	for len(hashes) > 1 {

		var nextLevel []string

		for i := 0; i < len(hashes); i += 2 {

			left := hashes[i]
			right := left

			if i+1 < len(hashes) {
				right = hashes[i+1]
			}

			nextLevel = append(
				nextLevel,
				hashData(left+right),
			)
		}

		hashes = nextLevel
	}

	return hashes[0]
}

// hashData calculates SHA-256 hash.
func hashData(data string) string {

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}
