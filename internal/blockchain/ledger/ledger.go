package ledger

import (
	"errors"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// Ledger maintains account balances derived from blockchain transactions.
type Ledger struct {
	balances map[string]float64
}

// NewLedger creates an empty ledger.
func NewLedger() *Ledger {
	return &Ledger{
		balances: make(map[string]float64),
	}
}

// GetBalance returns the current balance of an account.
func (l *Ledger) GetBalance(account string) float64 {
	return l.balances[account]
}

// ApplyTransaction updates the ledger using a validated transaction.
func (l *Ledger) ApplyTransaction(tx transaction.Transaction) error {

	if tx.Amount <= 0 {
		return errors.New("transaction amount must be positive")
	}

	// Coinbase transactions create new coins.
	if tx.Sender != transaction.CoinbaseAccount {

		if l.GetBalance(tx.Sender) < tx.Amount {
			return errors.New("insufficient balance")
		}

		l.balances[tx.Sender] -= tx.Amount
	}

	l.balances[tx.Recipient] += tx.Amount

	return nil
}
