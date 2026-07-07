package ledger

import (
	"errors"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// ValidateTransaction checks whether a transaction is valid
// based on amount rules and sender balance.
func (l *Ledger) ValidateTransaction(tx transaction.Transaction) error {

	if tx.Sender == "" {
		return errors.New("sender cannot be empty")
	}

	if tx.Recipient == "" {
		return errors.New("recipient cannot be empty")
	}

	if tx.Amount <= 0 {
		return errors.New("transaction amount must be positive")
	}

	// Coinbase transactions create new coins
	if tx.Sender == "COINBASE" {
		return nil
	}

	if l.GetBalance(tx.Sender) < tx.Amount {
		return errors.New("insufficient balance")
	}

	return nil
}
