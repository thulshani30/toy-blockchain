package mining

import (
	"context"
	"errors"
	"runtime"
	"strings"
	"sync"
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

	workers := runtime.NumCPU()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultCh := make(chan MiningResult, 1)

	var wg sync.WaitGroup

	for worker := 0; worker < workers; worker++ {

		wg.Add(1)

		go func(startNonce uint64) {
			defer wg.Done()

			var attempts uint64

			candidate := b

			for nonce := startNonce; ; nonce += uint64(workers) {

				select {
				case <-ctx.Done():
					return
				default:
				}

				candidate.Nonce = nonce
				attempts++

				hash := hashing.CalculateBlockHash(&candidate)
				candidate.Hash = hash

				if IsValidHash(hash, difficulty) {

					select {
					case resultCh <- MiningResult{
						Block:    candidate,
						Attempts: attempts,
						Duration: time.Since(start),
					}:
						cancel()
					default:
					}

					return
				}
			}
		}(uint64(worker))
	}

	result := <-resultCh

	wg.Wait()

	return result, nil
}
