package chain

import (
	"errors"
	"sync"
	"time"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/hashing"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/merkle"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/mining"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

const GenesisPreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"

// Blockchain represents the blockchain and its pending transaction pool.
type Blockchain struct {
	mu                  sync.RWMutex
	Blocks              []block.Block             `json:"blocks"`
	PendingTransactions []transaction.Transaction `json:"pending_transactions"`
	CurrentDifficulty   int                       `json:"current_difficulty"`
}

// NewGenesisBlock creates the deterministic genesis block.
func NewGenesisBlock() block.Block {
	transactions := []transaction.Transaction{}

	genesis := block.Block{
		Index:        0,
		Timestamp:    time.Unix(0, 0),
		Transactions: transactions,
		PreviousHash: GenesisPreviousHash,
		Nonce:        0,
		MerkleRoot:   merkle.CalculateMerkleRoot(transactions),
	}

	genesis.Hash = hashing.CalculateBlockHash(&genesis)

	return genesis
}

// NewBlockchain creates a blockchain containing only the genesis block.
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []block.Block{
			NewGenesisBlock(),
		},
		PendingTransactions: []transaction.Transaction{},
		CurrentDifficulty:   3,
	}
}

// GetLastBlock returns the latest block in the chain.
func (bc *Blockchain) GetLastBlock() (block.Block, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if len(bc.Blocks) == 0 {
		return block.Block{}, errors.New("blockchain is empty")
	}

	return bc.Blocks[len(bc.Blocks)-1], nil
}

// MinePendingTransactions mines all pending transactions into a new block.
func (bc *Blockchain) MinePendingTransactions(difficulty int, blockSize int) (block.Block, error) {

	bc.mu.Lock()
	defer bc.mu.Unlock()

	if len(bc.PendingTransactions) == 0 {
		return block.Block{}, errors.New("no pending transactions to mine")
	}

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	currentDifficulty := bc.CurrentDifficulty

	if currentDifficulty == 0 {
		currentDifficulty = difficulty
	}

	transactions := bc.PendingTransactions

	if len(transactions) > blockSize {
		transactions = transactions[:blockSize]
	}

	candidate := block.Block{
		Index:        lastBlock.Index + 1,
		Timestamp:    time.Now(),
		Transactions: transactions,
		PreviousHash: lastBlock.Hash,
		MerkleRoot:   merkle.CalculateMerkleRoot(transactions),
	}

	result, err := mining.MineBlock(candidate, currentDifficulty)

	minedBlock := result.Block

	if err != nil {
		return block.Block{}, err
	}

	bc.Blocks = append(bc.Blocks, minedBlock)
	bc.PendingTransactions = bc.PendingTransactions[len(transactions):]

	if lastBlock.Index > 0 {
		bc.CurrentDifficulty = mining.AdjustDifficulty(
			currentDifficulty,
			lastBlock.Timestamp,
			minedBlock.Timestamp,
		)
	}

	return minedBlock, nil
}

// Faucet creates a coinbase transaction and adds it to the pending pool.
func (bc *Blockchain) Faucet(recipient string, amount float64) error {

	tx, err := transaction.NewCoinbase(recipient, amount)
	if err != nil {
		return err
	}

	return bc.AddTransaction(tx)
}

// ReplaceChain replaces the current chain with a longer valid chain.
func (bc *Blockchain) ReplaceChain(newChain *Blockchain) {

	bc.mu.Lock()
	defer bc.mu.Unlock()

	if len(newChain.Blocks) > len(bc.Blocks) {
		bc.Blocks = newChain.Blocks
		bc.PendingTransactions = newChain.PendingTransactions
		bc.CurrentDifficulty = newChain.CurrentDifficulty
	}
}
