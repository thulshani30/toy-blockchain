# Toy Blockchain

A minimal blockchain and ledger simulator implemented in Go.

This project demonstrates:
- Block creation
- SHA-256 hashing
- Proof-of-Work mining
- Transaction handling
- Ledger management
- Blockchain validation

## Requirements

- Go 1.22+

## Project Structure

```text
toy-blockchain/
│
├── cmd/
│   └── toy-blockchain/
│       └── main.go
│
├── internal/
│   ├── blockchain/
│   │   ├── block/
│   │   ├── chain/
│   │   ├── hashing/
│   │   ├── ledger/
│   │   ├── mining/
│   │   └── transaction/
│   │
│   ├── config/
│   │
│   └── logger/
│
└── data/
```


## Run

Clone the repository:

```bash
git clone https://github.com/thulshani30/toy-blockchain.git
```

Run CLI application:

```bash
go run ./cmd/toy-blockchain
```