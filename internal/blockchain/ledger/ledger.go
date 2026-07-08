package ledger

import (
	"errors"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/block"
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

// GetBalances returns a copy of all account balances.
func (l *Ledger) GetBalances() map[string]float64 {
	copyBalances := make(map[string]float64)

	for account, balance := range l.balances {
		copyBalances[account] = balance
	}

	return copyBalances
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

// Reset clears all balances.
func (l *Ledger) Reset() {
	l.balances = make(map[string]float64)
}

// Rebuild reconstructs the ledger by replaying all blockchain transactions.
func (l *Ledger) Rebuild(blocks []block.Block) error {

	l.Reset()

	for _, blk := range blocks {
		for _, tx := range blk.Transactions {

			if err := l.ApplyTransaction(tx); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetAllBalances returns a copy of all account balances.
func (l *Ledger) GetAllBalances() map[string]float64 {

	balances := make(map[string]float64)

	for account, balance := range l.balances {
		balances[account] = balance
	}

	return balances
}
