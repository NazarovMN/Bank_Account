package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------- ")
		//panic("AYOOO YOU HOMELESS GUY AND HAVE NO MONEY,GO FIND A JOB DUDE")
	}

	var transactionHistory []string

	fmt.Println("Welcome to Go Bank!")
	for {

		fmt.Println("What do you want to do ?")

		fmt.Println("1. Check the balance ")
		fmt.Println("2. Deposit Money ")
		fmt.Println("3. Withdraw Money ")
		fmt.Println("4. Your Last 5 Transactions")
		fmt.Println("5. Exit ")

		fmt.Print("Your Choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			outputBalance(accountBalance)
		case 2:
			depositAmount(accountBalance, transactionHistory)
		case 3:
			withdrawalAmount(accountBalance, transactionHistory)
		case 4:
			if len(transactionHistory) == 0 {
				fmt.Println("No transactions available.")
			} else {
				start := len(transactionHistory) - 5
				if start < 0 {
					start = 0
				}
				lastFive := transactionHistory[start:]
				fmt.Println("Your Last 5 Transactions:")
				for _, transaction := range lastFive {
					fmt.Println(transaction)
				}
			}

		case 5:
			fmt.Println("Bye Bye!")
			fmt.Println("Thanks for choosing our bank! ")
			return
		default:
			fmt.Println("Invalid Choice. Please choice a valid otion.")

		}

	}
}

func outputBalance(int) {
	accountBalance, _ := getBalanceFromFile()
	fmt.Println("Your Balance is : ", accountBalance)
}

func depositAmount(accountBalance int, transactionHistory []string) {
	updatedBalance, _ := getBalanceFromFile()
	var depositAmount int
	fmt.Print("Your deposit: ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid Amount ,Must be greater than 0.")
	}

	updatedBalance += depositAmount //accountBalance = accountBalance + deposit
	writeBalanceToFile(updatedBalance)
	fmt.Println("Balance Updated: ", accountBalance+(updatedBalance-accountBalance))
	transaction := fmt.Sprintf("Deposit: %v ", depositAmount)
	transactionHistory = append(transactionHistory, transaction)
	fmt.Printf("Last 5 Transactions: %v \n", transactionHistory)
}

func withdrawalAmount(accountBalance int, transactionHistory []string) {
	updatedBalance, _ := getBalanceFromFile()
	var withdrawalAmount int
	fmt.Print("Your withdrawal: ")
	fmt.Scan(&withdrawalAmount)

	if withdrawalAmount <= 0 {
		fmt.Println("Invalid Amount ,Must be greater than 0.")
	}

	if withdrawalAmount > accountBalance {
		fmt.Println("Invalid Amount. Your Balance is :", accountBalance)
	}

	updatedBalance -= withdrawalAmount
	writeBalanceToFile(updatedBalance)
	fmt.Println("Balance Updated: ", updatedBalance)

	transaction := fmt.Sprintf("Withdrawal: %v", withdrawalAmount)
	transactionHistory = append(transactionHistory, transaction)
	fmt.Printf("Last 5 Transaction : %v \n", transactionHistory)
}

func writeBalanceToFile(balance int) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (int, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("You have no money on a Bank Account. Do you want to Deposit?")
	}

	balanceText := string(data)
	convertedBalance, err := strconv.ParseInt(balanceText, 10, 64)

	if err != nil {
		return 0, errors.New("Converiton Problem.")
	}

	balance := int(convertedBalance)
	return balance, nil
}
