package main

import (
	"flag"
	"fmt"

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

	logger.Info.Println("Starting Toy Blockchain")

	fmt.Println("----------------")
	fmt.Println("Difficulty:", *difficulty)
	fmt.Println("Block size:", cfg.BlockSize)
	fmt.Println("Data path:", cfg.DataPath)

	logger.Info.Println("Application started successfully")
}
