package main

import "fmt"

func main() {

	var accountBalance float64 = 1000
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

		if choice == 1 {

			fmt.Println("Your Balance is : ", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount ,Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount //accountBalance = accountBalance + deposit
			fmt.Println("Balance Updated: ", accountBalance)
			transaction := fmt.Sprintf("Deposit: %.1f ", depositAmount)
			transactionHistory = append(transactionHistory, transaction)
			fmt.Printf("Last 5 Transactions: %v \n", transactionHistory)
		} else if choice == 3 {
			var withdrawalAmount float64
			fmt.Print("Your withdrawal: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount ,Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid Amount. Your Balance is :", accountBalance)
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance Updated: ", accountBalance)

			transaction := fmt.Sprintf("Withdrawal: %.1f", withdrawalAmount)
			transactionHistory = append(transactionHistory, transaction)
			fmt.Printf("Last 5 Transaction : %v \n", transactionHistory)
		} else if choice == 4 {

			fmt.Println("Your Last Transactions:", transactionHistory)

			numTransaction := len(transactionHistory)
			start := numTransaction - 5
			if start < 0 {
				start = 0
			}
			for i := start; i < numTransaction; i++ {
				fmt.Println(transactionHistory[i])
			}

		} else if choice == 5 {
			fmt.Println("Bye Bye!")
			break
		} else {
			fmt.Println("Invalid Choice. Please choice a valid otion.")

		}
	}

	fmt.Println("Thanks for choosing our bank! ")
}
