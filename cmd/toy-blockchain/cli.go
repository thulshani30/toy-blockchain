package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/crypto"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/storage"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/validation"
)

func StartCLI(bc *chain.Blockchain, difficulty int, blockSize int, dataPath string) {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println()
		fmt.Println("=========================")
		fmt.Println("   Toy Blockchain CLI")
		fmt.Println("=========================")
		fmt.Println("1. View Blockchain")
		fmt.Println("2. Add Transaction")
		fmt.Println("3. View Pending Transactions")
		fmt.Println("4. Mine Transactions")
		fmt.Println("5. Validate Blockchain")
		fmt.Println("6. Faucet (Create Coinbase Transaction)")
		fmt.Println("7. Show Balance")
		fmt.Println("8. Exit")
		fmt.Print("Select option: ")

		input, _ := reader.ReadString('\n')
		option := strings.TrimSpace(input)

		switch option {

		case "1":
			viewBlockchain(bc)

		case "2":
			addTransaction(reader, bc)

			if err := storage.SaveBlockchain(dataPath, bc); err != nil {
				fmt.Println("Auto-save failed:", err)
			} else {
				fmt.Println("Blockchain saved.")
			}

		case "3":
			viewPendingTransactions(bc)

		case "4":

			_, err := bc.MinePendingTransactions(difficulty, blockSize)

			if err != nil {
				fmt.Println("Mining error:", err)
			} else {
				fmt.Println("Block mined successfully")
			}

		case "5":

			err := validation.ValidateChain(bc, difficulty)

			if err != nil {
				fmt.Println("Blockchain invalid:", err)
			} else {
				fmt.Println("Blockchain is valid")
			}

		case "6":
			faucet(reader, bc, dataPath)

		case "7":
			showBalances(bc)

		case "8":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}

func viewBlockchain(bc *chain.Blockchain) {

	if len(bc.Blocks) == 0 {
		fmt.Println("Blockchain is empty.")
		return
	}

	fmt.Println()
	fmt.Println("========== TOY BLOCKCHAIN ==========")

	for _, block := range bc.Blocks {

		fmt.Println("========================================")
		fmt.Printf("Block #%d\n", block.Index)
		fmt.Println("========================================")

		fmt.Printf("Timestamp      : %s\n", block.Timestamp.Format("2006-01-02 15:04:05"))
		fmt.Printf("Previous Hash  : %s\n", block.PreviousHash)
		fmt.Printf("Hash           : %s\n", block.Hash)
		fmt.Printf("Nonce          : %d\n", block.Nonce)

		fmt.Println()
		fmt.Println("Transactions")

		if len(block.Transactions) == 0 {
			fmt.Println("  None")
		} else {
			for i, tx := range block.Transactions {
				fmt.Printf(
					"  %d. %s -> %s : %.2f\n",
					i+1,
					tx.Sender,
					tx.Recipient,
					tx.Amount,
				)
			}
		}

		fmt.Println()
	}

	fmt.Println("========== END OF CHAIN ==========")
}

func addTransaction(reader *bufio.Reader, bc *chain.Blockchain) {

	fmt.Print("Sender: ")
	sender, _ := reader.ReadString('\n')

	fmt.Print("Recipient: ")
	recipient, _ := reader.ReadString('\n')

	fmt.Print("Amount: ")
	amountText, _ := reader.ReadString('\n')

	amount, err := strconv.ParseFloat(strings.TrimSpace(amountText), 64)

	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	wallet, err := crypto.NewWallet()

	if err != nil {
		fmt.Println("Wallet creation failed:", err)
		return
	}

	tx := transaction.Transaction{
		Sender:    strings.TrimSpace(sender),
		Recipient: strings.TrimSpace(recipient),
		Amount:    amount,
	}

	err = crypto.SignTransaction(
		&tx,
		wallet.PrivateKey,
	)

	if err != nil {
		fmt.Println("Transaction signing failed:", err)
		return
	}

	err = bc.AddTransaction(tx)

	if err != nil {
		fmt.Println("Transaction failed:", err)
		return
	}

	fmt.Println("Transaction added")
}

func faucet(reader *bufio.Reader, bc *chain.Blockchain, dataPath string) {

	fmt.Print("Recipient: ")
	recipient, _ := reader.ReadString('\n')

	fmt.Print("Amount: ")
	amountText, _ := reader.ReadString('\n')

	amount, err := strconv.ParseFloat(strings.TrimSpace(amountText), 64)

	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	err = bc.Faucet(
		strings.TrimSpace(recipient),
		amount,
	)

	if err != nil {
		fmt.Println("Faucet failed:", err)
		return
	}

	fmt.Println("Coins added to pending transactions")
}

func showBalances(bc *chain.Blockchain) {

	l, err := bc.BuildLedger()

	if err != nil {
		fmt.Println("Failed to calculate balances:", err)
		return
	}

	balances := l.GetAllBalances()

	if len(balances) == 0 {
		fmt.Println("No account balances available.")
		return
	}

	fmt.Println()
	fmt.Println("Account Balances")
	fmt.Println("----------------")

	for account, balance := range balances {
		fmt.Printf("%-15s %.2f\n", account, balance)
	}
}

func viewPendingTransactions(bc *chain.Blockchain) {

	pending := bc.GetPendingTransactions()

	fmt.Println()
	fmt.Println("========== PENDING TRANSACTIONS ==========")

	if len(pending) == 0 {
		fmt.Println("No pending transactions.")
		return
	}

	for i, tx := range pending {
		fmt.Printf(
			"%d. %s -> %s : %.2f\n",
			i+1,
			tx.Sender,
			tx.Recipient,
			tx.Amount,
		)
	}
}
