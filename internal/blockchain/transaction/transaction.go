package transaction

const CoinbaseAccount = "COINBASE"

// Transaction represents a transfer of funds between two accounts.
// It is the basic unit of data stored inside a block.
type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}
