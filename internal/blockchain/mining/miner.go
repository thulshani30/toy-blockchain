package mining

import (
	"errors"
	"strings"
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/hashing"
)

// IsValidHash checks whether a hash satisfies the required difficulty.
func IsValidHash(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

// MineBlock performs proof-of-work mining.
func MineBlock(b block.Block, difficulty int) (MiningResult, error) {

	if difficulty < 0 {
		return MiningResult{}, errors.New("difficulty cannot be negative")
	}

	start := time.Now()

	var attempts uint64

	for nonce := uint64(0); ; nonce++ {

		b.Nonce = nonce
		attempts++

		hash := hashing.CalculateBlockHash(&b)
		b.Hash = hash

		if IsValidHash(hash, difficulty) {
			return MiningResult{
				Block:    b,
				Attempts: attempts,
				Duration: time.Since(start),
			}, nil
		}
	}
}
