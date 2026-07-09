# Toy Blockchain

A lightweight blockchain implementation written in Go that demonstrates the core concepts of blockchain technology, including blocks, transactions, cryptographic hashing, proof-of-work mining, ledger management, validation, and persistent storage.

This project was developed as a backend engineering assessment to demonstrate fundamental blockchain concepts through a functional command-line application.

---

# Features

## Blockchain Core

- Custom block structure containing:
  - Block index
  - Timestamp
  - Transactions
  - Previous block hash
  - Nonce
  - Block hash
- Deterministic genesis block
- Deterministic SHA-256 block hashing
- Hash-linked blockchain
- Secure block creation workflow

## Transactions & Ledger

- Transaction model with sender, recipient, and amount
- Pending transaction pool
- Account balance calculation
- Transaction validation
- Coinbase (faucet) transactions for introducing new coins
- Prevention of invalid transactions
- Prevention of overspending

## Mining

- Proof-of-Work (PoW) mining
- Configurable mining difficulty
- Configurable block size
- Mining statistics:
  - Nonce
  - Number of attempts
  - Mining duration

## Validation & Security

- Blockchain integrity validation
- Hash verification
- Previous hash verification
- Proof-of-Work verification
- Tamper detection

## Persistence

- JSON-based blockchain storage
- Automatic saving after blockchain modifications
- Blockchain loading during application startup

## Command-Line Interface

Interactive CLI supporting:

- View blockchain
- Add transactions
- View pending transactions
- Mine pending transactions
- Validate blockchain
- Create faucet transactions
- View account balances

---

# Architecture

The project follows a modular architecture where each package has a single responsibility.

```
                    +----------------------+
                    |   Command Line CLI   |
                    +----------+-----------+
                               |
                               v
                  +--------------------------+
                  |        Blockchain        |
                  +-----------+--------------+
                              |
      +-----------------------+-----------------------+
      |                       |                       |
      v                       v                       v
+-------------+       +----------------+      +---------------+
|   Ledger    |       |    Mining      |      |  Validation   |
+-------------+       +----------------+      +---------------+
      |                       |                       |
      +-----------------------+-----------------------+
                              |
                              v
                     +------------------+
                     |     Storage      |
                     |   JSON Persist   |
                     +------------------+
```

## Package Responsibilities

| Package | Responsibility |
|----------|----------------|
| `block` | Defines the block structure |
| `chain` | Manages blockchain operations |
| `transaction` | Defines transactions and coinbase transactions |
| `ledger` | Maintains balances and validates transactions |
| `hashing` | Calculates deterministic SHA-256 hashes |
| `mining` | Implements Proof-of-Work mining |
| `validation` | Validates blockchain integrity |
| `storage` | Saves and loads blockchain state |
| `config` | Stores configurable parameters |
| `logger` | Handles application logging |

---

# Project Structure

```
toy-blockchain/
│
├── cmd/
│   └── toy-blockchain/
│       └── main.go
│
├── data/
│   └── blockchain.json
│
├── experiments/
│   ├── difficulty/
│   └── tamper/
│
├── internal/
│   ├── blockchain/
│   │   ├── block/
│   │   ├── chain/
│   │   ├── hashing/
│   │   ├── ledger/
│   │   ├── mining/
│   │   ├── storage/
│   │   ├── transaction/
│   │   └── validation/
│   │
│   ├── config/
│   └── logger/
│
├── go.mod
├── go.sum
└── README.md
```

---

# Requirements

- Go 1.22 or newer
- Git

---

# Installation

Clone the repository:

```bash
git clone https://github.com/thulshani30/toy-blockchain.git
```

Move into the project directory:

```bash
cd toy-blockchain
```

Download dependencies:

```bash
go mod tidy
```

---

# Running the Application

Run the interactive CLI:

```bash
go run ./cmd/toy-blockchain
```

The blockchain is automatically loaded from disk. If no blockchain exists, a new blockchain containing only the genesis block is created.

---

# CLI Usage

The application provides the following menu:

```
=========================
   Toy Blockchain CLI
=========================

1. View Blockchain
2. Add Transaction
3. View Pending Transactions
4. Mine Transactions
5. Validate Blockchain
6. Faucet (Create Coinbase Transaction)
7. Show Balance
8. Exit
```

## Create Initial Funds

Select:

```
6
```

Example:

```
Recipient: Alice
Amount: 1000
```

A coinbase transaction will be added to the pending transaction pool.

---

## Add a Transaction

Select:

```
2
```

Example:

```
Sender: Alice
Recipient: Bob
Amount: 100
```

The transaction is validated before entering the pending transaction pool.

---

## Mine Pending Transactions

Select:

```
4
```

The miner searches for a nonce satisfying the configured Proof-of-Work difficulty.

---

## Validate the Blockchain

Select:

```
5
```

The application verifies:

- Block hashes
- Previous hash links
- Proof-of-Work
- Blockchain integrity

---

## View Account Balance

Select:

```
7
```

Example:

```
Account: Alice

Balance: 900
```

---

# Configuration

The blockchain uses configurable parameters.

| Parameter | Description | Default |
|-----------|-------------|---------|
| Difficulty | Number of leading zeros required for mining | 3 |
| Block Size | Maximum transactions per block | 10 |
| Data Path | Blockchain JSON file | `data/blockchain.json` |

---

# Blockchain Design

## Block Structure

Each block contains:

- Index
- Unix timestamp
- Transaction list
- Previous block hash
- Nonce
- Block hash

## Deterministic Hashing

Blocks are hashed using SHA-256 over the following fields:

1. Index
2. Timestamp
3. Transactions
4. Previous hash
5. Nonce

The block hash itself is excluded from the hash calculation, ensuring deterministic hashing.

## Genesis Block

The blockchain begins with a deterministic genesis block.

```
Index: 0

Previous Hash:
0000000000000000000000000000000000000000000000000000000000000000
```

## Proof-of-Work

Mining requires finding a nonce such that the resulting SHA-256 hash begins with a configurable number of leading zero hexadecimal digits.

Example:

Difficulty = 3

```
000d2f41ab5f...
```

---

# Experiments

## Tamper Detection

A block was intentionally modified after mining.

Output:

```
Before tampering:
true

After tampering:
Detected: block hash mismatch
```

The validator correctly detected the tampered block.

---

## Mining Difficulty

Mining was performed using different difficulty levels.

Example results:

| Difficulty | Attempts |
|------------|----------|
| 1 | 14 |
| 2 | 47 |
| 3 | 11031 |
| 4 | 51179 |

As expected, mining becomes more computationally expensive as difficulty increases.

---

# Testing

Run all automated tests:

```bash
go test ./...
```

Implemented tests include:

- Genesis block creation
- Blockchain initialization
- Transaction validation
- Ledger balance calculation
- Insufficient balance rejection
- Deterministic hashing
- Hash modification detection
- Proof-of-Work validation
- Blockchain tamper detection

Example output:

```
ok   github.com/thulshani30/toy-blockchain/internal/blockchain/chain
ok   github.com/thulshani30/toy-blockchain/internal/blockchain/hashing
ok   github.com/thulshani30/toy-blockchain/internal/blockchain/ledger
ok   github.com/thulshani30/toy-blockchain/internal/blockchain/mining
ok   github.com/thulshani30/toy-blockchain/internal/blockchain/validation
```

---

# Future Improvements

Potential future enhancements include:

- Digital signatures
- Wallet generation
- Peer-to-peer networking
- Node synchronization
- Transaction fees
- Merkle trees
- REST API
- Database-backed persistence
- Multi-node consensus

---

# License

This project was developed for educational and assessment purposes.