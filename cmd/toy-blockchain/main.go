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

	flag.Parse()

	bc, err := storage.LoadBlockchain(cfg.DataPath)
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
		cfg.DataPath,
	)
	fmt.Println("----------------")
	fmt.Println("Blocks:", len(bc.Blocks))
	fmt.Println("Difficulty:", *difficulty)
	fmt.Println("Block size:", cfg.BlockSize)
	fmt.Println("Data path:", cfg.DataPath)

	if err := storage.SaveBlockchain(cfg.DataPath, bc); err != nil {
		logger.Error.Fatal(err)
	}

	logger.Info.Println("Blockchain state saved")
}
