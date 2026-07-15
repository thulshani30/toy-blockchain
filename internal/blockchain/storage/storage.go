package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
)

// SaveBlockchain safely writes blockchain data using atomic replacement.
func SaveBlockchain(path string, bc *chain.Blockchain) error {

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		return err
	}

	tempPath := path + ".tmp"

	file, err := os.Create(tempPath)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return os.Rename(tempPath, path)
}

// LoadBlockchain loads blockchain data from disk.
func LoadBlockchain(path string) (*chain.Blockchain, error) {

	data, err := os.ReadFile(path)

	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			return chain.NewBlockchain(), nil
		}

		return nil, err
	}

	var bc chain.Blockchain

	if err := json.Unmarshal(data, &bc); err != nil {

		recovered, recoverErr := RecoverBlockchain(path)

		if recoverErr != nil {
			return nil, err
		}

		return recovered, nil
	}

	if bc.CurrentDifficulty == 0 {
		bc.CurrentDifficulty = 3
	}

	return &bc, nil
}

func RecoverBlockchain(path string) (*chain.Blockchain, error) {

	backupPath := path + ".tmp"

	data, err := os.ReadFile(backupPath)

	if err != nil {
		return nil, errors.New("no recovery file found")
	}

	var bc chain.Blockchain

	if err := json.Unmarshal(data, &bc); err != nil {
		return nil, err
	}

	return &bc, nil
}
