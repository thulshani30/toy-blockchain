package transaction

import "errors"

// CoinbaseAccount is the special account used to create new coins.
const CoinbaseAccount = "COINBASE"

// Transaction represents a transfer of funds between two accounts.
type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`

	PublicKey string `json:"public+key"`
	Signature string `json:"signature"`
}

// New creates a standard transaction.
func New(sender, recipient string, amount float64) (Transaction, error) {

	tx := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	if err := tx.Validate(); err != nil {
		return Transaction{}, err
	}

	return tx, nil
}

// NewCoinbase creates a coinbase (faucet) transaction.
func NewCoinbase(recipient string, amount float64) (Transaction, error) {

	tx := Transaction{
		Sender:    CoinbaseAccount,
		Recipient: recipient,
		Amount:    amount,
	}

	if err := tx.Validate(); err != nil {
		return Transaction{}, err
	}

	return tx, nil
}

// Validate performs basic transaction validation.
func (tx Transaction) Validate() error {

	if tx.Sender == "" {
		return errors.New("sender cannot be empty")
	}

	if tx.Recipient == "" {
		return errors.New("recipient cannot be empty")
	}

	if tx.Amount <= 0 {
		return errors.New("transaction amount must be positive")
	}

	return nil
}
