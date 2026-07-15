package block

import (
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// Block represents a single block in the blockchain.
type Block struct {
	// Index is the block height in the chain.
	Index uint64 `json:"index"`

	// Timestamp records when the block was created.
	Timestamp time.Time `json:"timestamp"`

	// Transactions included in this block.
	Transactions []transaction.Transaction `json:"transactions"`

	// Hash of the previous block.
	PreviousHash string `json:"previous_hash"`

	// Nonce discovered during Proof-of-Work mining.
	Nonce uint64 `json:"nonce"`

	// Merkle root of all transactions in this block.
	MerkleRoot string `json:"merkle_root"`

	// SHA-256 hash of this block.
	Hash string `json:"hash"`
}
