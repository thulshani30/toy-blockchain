package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thulshani30/toy-blockchain/internal/blockchain/chain"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/ledger"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/storage"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/transaction"
	"github.com/thulshani30/toy-blockchain/internal/blockchain/validation"
)

func StartCLI(bc *chain.Blockchain, difficulty int, dataPath string) {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println()
		fmt.Println("=========================")
		fmt.Println("   Toy Blockchain CLI")
		fmt.Println("=========================")
		fmt.Println("1. View Blockchain")
		fmt.Println("2. Add Transaction")
		fmt.Println("3. Mine Transactions")
		fmt.Println("4. Validate Blockchain")
		fmt.Println("5. Faucet (Create Coinbase Transaction)")
		fmt.Println("6. Show Balance")
		fmt.Println("7. Exit")
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

			_, err := bc.MinePendingTransactions(difficulty)

			if err != nil {
				fmt.Println("Mining error:", err)
			} else {
				fmt.Println("Block mined successfully")
			}

		case "4":

			err := validation.ValidateChain(bc, difficulty)

			if err != nil {
				fmt.Println("Blockchain invalid:", err)
			} else {
				fmt.Println("Blockchain is valid")
			}

		case "5":
			faucet(reader, bc, dataPath)

		case "6":
			showBalances(bc)

		case "7":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}

func viewBlockchain(bc *chain.Blockchain) {

	for _, block := range bc.Blocks {

		fmt.Println("----------------------")
		fmt.Println("Index:", block.Index)
		fmt.Println("Previous:", block.PreviousHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("Transactions:")

		for _, tx := range block.Transactions {
			fmt.Printf("  %s -> %s : %.2f\n",
				tx.Sender,
				tx.Recipient,
				tx.Amount,
			)
		}
	}
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

	tx := transaction.Transaction{
		Sender:    strings.TrimSpace(sender),
		Recipient: strings.TrimSpace(recipient),
		Amount:    amount,
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

	l, err := ledger.BuildLedger(bc)

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
