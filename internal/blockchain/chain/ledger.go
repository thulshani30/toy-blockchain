package chain

import (
	"github.com/thulshani30/toy-blockchain/internal/blockchain/ledger"
)

// BuildLedger creates a ledger from the blockchain.
func (bc *Blockchain) BuildLedger() (*ledger.Ledger, error) {

	l := ledger.NewLedger()

	for _, b := range bc.Blocks {
		for _, tx := range b.Transactions {

			if err := l.ApplyTransaction(tx); err != nil {
				return nil, err
			}
		}
	}

	return l, nil
}
