package crypto

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// VerifyTransaction verifies the transaction signature.
func VerifyTransaction(tx transaction.Transaction) bool {

	block, _ := pem.Decode([]byte(tx.PublicKey))

	if block == nil {
		return false
	}

	publicKeyData, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return false
	}

	publicKey, ok := publicKeyData.(*ecdsa.PublicKey)

	if !ok {
		return false
	}

	data := fmt.Sprintf(
		"%s%s%f",
		tx.Sender,
		tx.Recipient,
		tx.Amount,
	)

	hash := sha256.Sum256([]byte(data))

	signatureBytes, err := hex.DecodeString(tx.Signature)

	if err != nil {
		return false
	}

	if len(signatureBytes) < 2 {
		return false
	}

	r := new(big.Int).SetBytes(signatureBytes[:len(signatureBytes)/2])
	s := new(big.Int).SetBytes(signatureBytes[len(signatureBytes)/2:])

	return ecdsa.Verify(
		publicKey,
		hash[:],
		r,
		s,
	)
}
