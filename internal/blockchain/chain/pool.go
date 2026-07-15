package chain

import (
	"errors"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/crypto"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// AddTransaction validates and adds a transaction to the pending pool.
func (bc *Blockchain) AddTransaction(tx transaction.Transaction) error {

	bc.mu.Lock()
	defer bc.mu.Unlock()

	if err := ValidatePendingTransaction(tx); err != nil {
		return err
	}

	// Verify digital signature
	if tx.Sender != transaction.CoinbaseAccount {

		if !crypto.VerifyTransaction(tx) {
			return errors.New("invalid transaction signature")
		}

	}

	l, err := bc.BuildLedger()
	if err != nil {
		return err
	}

	for _, pendingTx := range bc.PendingTransactions {
		if err := l.ApplyTransaction(pendingTx); err != nil {
			return err
		}
	}

	if err := l.ValidateTransaction(tx); err != nil {
		return err
	}

	bc.PendingTransactions = append(
		bc.PendingTransactions,
		tx,
	)

	return nil
}

// GetPendingTransactions returns a copy of the pending transactions.
func (bc *Blockchain) GetPendingTransactions() []transaction.Transaction {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	pending := make([]transaction.Transaction, len(bc.PendingTransactions))
	copy(pending, bc.PendingTransactions)

	return pending
}

// ClearPendingTransactions removes all pending transactions.
func (bc *Blockchain) ClearPendingTransactions() {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	bc.PendingTransactions = []transaction.Transaction{}
}

// ValidatePendingTransaction performs basic validation before a transaction
// is accepted into the pending pool.
func ValidatePendingTransaction(tx transaction.Transaction) error {

	if tx.Sender == "" {
		return errors.New("sender cannot be empty")
	}

	if tx.Recipient == "" {
		return errors.New("recipient cannot be empty")
	}

	if tx.Amount <= 0 {
		return errors.New("transaction amount must be greater than zero")
	}

	return nil
}
