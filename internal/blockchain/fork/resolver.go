package fork

import (
	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/validation"
)

// ResolveFork selects the longest valid blockchain.
func ResolveFork(
	current *chain.Blockchain,
	competing *chain.Blockchain,
	difficulty int,
) (*chain.Blockchain, error) {

	if err := validation.ValidateChain(current, difficulty); err != nil {
		return nil, err
	}

	if err := validation.ValidateChain(competing, difficulty); err != nil {
		return nil, err
	}

	if len(competing.Blocks) > len(current.Blocks) {
		return competing, nil
	}

	return current, nil
}
