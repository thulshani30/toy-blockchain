package crypto

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

// EncodePublicKey converts a public key into a string format.
func EncodePublicKey(publicKey *ecdsa.PublicKey) (string, error) {

	bytes, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return "", err
	}

	block := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	})

	return string(block), nil
}
