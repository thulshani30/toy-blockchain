package validation

import (
	"errors"
	"fmt"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/hashing"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/mining"
)

// ValidateChain checks whether the blockchain is valid and has not been tampered with.
func ValidateChain(bc *chain.Blockchain, difficulty int) error {

	if len(bc.Blocks) == 0 {
		return errors.New("blockchain is empty")
	}

	for i, currentBlock := range bc.Blocks {

		// Check block hash integrity
		calculatedHash := hashing.CalculateBlockHash(&currentBlock)

		if calculatedHash != currentBlock.Hash {
			return fmt.Errorf(
				"block %d hash mismatch: stored=%s calculated=%s",
				i,
				currentBlock.Hash,
				calculatedHash,
			)
		}

		// Check proof-of-work validity
		if !mining.IsValidHash(currentBlock.Hash, difficulty) && i != 0 {
			return fmt.Errorf(
				"block %d does not satisfy proof-of-work difficulty",
				i,
			)
		}

		// Skip previous hash check for genesis block
		if i == 0 {
			continue
		}

		previousBlock := bc.Blocks[i-1]

		// Check chain linkage
		if currentBlock.PreviousHash != previousBlock.Hash {
			return fmt.Errorf(
				"block %d previous hash mismatch",
				i,
			)
		}

		// Check block height consistency
		if currentBlock.Index != previousBlock.Index+1 {
			return fmt.Errorf(
				"block %d invalid index",
				i,
			)
		}

		// Timestamp should not go backwards
		if currentBlock.Timestamp.Before(previousBlock.Timestamp) {
			return fmt.Errorf(
				"block %d timestamp is earlier than previous block",
				i,
			)
		}
	}

	return nil
}
