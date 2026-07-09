package validation_test

import (
	"testing"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/validation"
)

func TestTamperDetection(t *testing.T) {

	bc := chain.NewBlockchain()

	err := validation.ValidateChain(bc, 3)

	if err != nil {
		t.Fatalf("expected valid chain: %v", err)
	}

	// Tamper with genesis block
	bc.Blocks[0].PreviousHash = "tampered"

	err = validation.ValidateChain(bc, 3)

	if err == nil {
		t.Error("expected validation failure after tampering")
	}
}
func TestBrokenPreviousHashDetection(t *testing.T) {

	bc := chain.NewBlockchain()

	bc.Blocks[0].PreviousHash = "invalid_previous_hash"

	err := validation.ValidateChain(bc, 3)

	if err == nil {
		t.Error("expected validation failure for modified previous hash")
	}
}
