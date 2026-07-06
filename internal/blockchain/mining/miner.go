package mining

import (
	"fmt"
	"strings"
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/hashing"
)

func isValidHash(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func MineBlock(b block.Block, difficulty int) (block.Block, error) {
	start := time.Now()

	nonce := uint64(0)

	for {
		b.Nonce = nonce

		hash, err := hashing.CalculateBlockHash(&b)
		if err != nil {
			return block.Block{}, err
		}

		b.Hash = hash

		if isValidHash(hash, difficulty) {
			fmt.Println("Block mined!")
			fmt.Println("Nonce:", nonce)
			fmt.Println("Time taken:", time.Since(start))

			return b, nil
		}

		nonce++
	}
}