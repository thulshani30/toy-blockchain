package main

import (
	"flag"
	"fmt"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/storage"
	"github.com/thulshani30/toy-blockchain/internal/config"
	"github.com/thulshani30/toy-blockchain/internal/logger"
)

func main() {

	logger.Init()

	cfg := config.DefaultConfig()

	difficulty := flag.Int(
		"difficulty",
		cfg.Difficulty,
		"mining difficulty",
	)

	blockSize := flag.Int(
		"block-size",
		cfg.BlockSize,
		"maximum transactions per block",
	)

	dataPath := flag.String(
		"data",
		cfg.DataPath,
		"blockchain data file path",
	)

	flag.Parse()

	bc, err := storage.LoadBlockchain(*dataPath)
	if err != nil {
		logger.Error.Fatal(err)
	}

	if err := storage.ValidateLoadedBlockchain(bc, *difficulty); err != nil {
		logger.Error.Fatal(err)
	}

	logger.Info.Println("Blockchain integrity verified")

	logger.Info.Println("Blockchain loaded successfully")

	StartCLI(
		bc,
		*difficulty,
		*blockSize,
		*dataPath,
	)
	fmt.Println("----------------")
	fmt.Println("Blocks:", len(bc.Blocks))
	fmt.Println("Difficulty:", *difficulty)
	fmt.Println("Block size:", *blockSize)
	fmt.Println("Data path:", *dataPath)

	if err := storage.SaveBlockchain(*dataPath, bc); err != nil {
		logger.Error.Fatal(err)
	}

	logger.Info.Println("Blockchain state saved")
}
