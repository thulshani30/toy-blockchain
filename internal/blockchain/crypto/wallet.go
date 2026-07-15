package crypto

import "crypto/ecdsa"

// Wallet stores a user's key pair.
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  string
}

// NewWallet creates a new wallet.
func NewWallet() (*Wallet, error) {

	privateKey, publicKey, err := GenerateKeyPair()

	if err != nil {
		return nil, err
	}

	encodedPublicKey, err := EncodePublicKey(publicKey)

	if err != nil {
		return nil, err
	}

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  encodedPublicKey,
	}, nil
}
