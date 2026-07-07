package config

// Config stores application configuration values.
type Config struct {
	Difficulty int    `json:"difficulty"`
	BlockSize  int    `json:"block_size"`
	DataPath   string `json:"data_path"`
}

// DefaultConfig returns the default blockchain configuration.
func DefaultConfig() Config {
	return Config{
		Difficulty: 3,
		BlockSize:  10,
		DataPath:   "data/blockchain.json",
	}
}
