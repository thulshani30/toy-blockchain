package mining

import (
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
)

// MiningResult contains statistics about a mining operation.
type MiningResult struct {
	Block    block.Block
	Attempts uint64
	Duration time.Duration
}
