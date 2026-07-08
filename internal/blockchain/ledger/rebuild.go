package ledger

import "github.com/thulshani30/toy-blockchain/internal/blockchain/chain"

// BuildLedger creates a ledger from the blockchain.
func BuildLedger(bc *chain.Blockchain) (*Ledger, error) {

	l := NewLedger()

	for _, b := range bc.Blocks {

		for _, tx := range b.Transactions {

			if err := l.ApplyTransaction(tx); err != nil {
				return nil, err
			}
		}
	}

	return l, nil
}
