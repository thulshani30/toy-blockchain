package crypto

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
)

// SignTransaction creates a digital signature for a transaction.
func SignTransaction(tx *transaction.Transaction, privateKey *ecdsa.PrivateKey) error {

	data := fmt.Sprintf(
		"%s%s%f",
		tx.Sender,
		tx.Recipient,
		tx.Amount,
	)

	hash := sha256.Sum256([]byte(data))

	r, s, err := ecdsa.Sign(
		nil,
		privateKey,
		hash[:],
	)

	if err != nil {
		return err
	}

	signature := append(r.Bytes(), s.Bytes()...)

	tx.Signature = hex.EncodeToString(signature)

	publicKey, err := EncodePublicKey(&privateKey.PublicKey)

	if err != nil {
		return err
	}

	tx.PublicKey = publicKey

	return nil
}
