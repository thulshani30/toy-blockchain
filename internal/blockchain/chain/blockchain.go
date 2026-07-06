package chain

//sync.RWMutex is used to ensure that the blockchain can be accessed by multiple goroutines safely. It allows multiple readers or one writer at a time
//Blocks is a slice of Block structs that represents the entire blockchain.

import (
	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/hashing"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
	"sync"
	"time"
)

// Blockchain represents the complete blockchain.
type Blockchain struct {
	mu     sync.RWMutex
	Blocks []block.Block `json:"blocks"`
}

const GenesisPreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"

func NewGenesisBlock() block.Block {
	genesis := block.Block{
		Index:        0,
		Timestamp:    time.Unix(0, 0),
		Transactions: []transaction.Transaction{},
		PreviousHash: GenesisPreviousHash,
		Nonce:        0,
	}

	hash, _ := hashing.CalculateBlockHash(&genesis)
	genesis.Hash = hash

	return genesis
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []block.Block{
			NewGenesisBlock(),
		},
	}
}

func (bc *Blockchain) getLastBlock() block.Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(transactions []transaction.Transaction) block.Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	lastBlock := bc.getLastBlock()

	newBlock := block.Block{
		Index:        lastBlock.Index + 1,
		Timestamp:    time.Now(),
		Transactions: transactions,
		PreviousHash: lastBlock.Hash,
		Nonce:        0,
	}

	hash, _ := hashing.CalculateBlockHash(&newBlock)
	newBlock.Hash = hash

	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}
