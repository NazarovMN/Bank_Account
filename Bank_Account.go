package main

import "fmt"

func main() {

	var accountBalance int = 1000
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
			fmt.Println("Your Balance is : ", accountBalance)
		case 2:
			var depositAmount int
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount ,Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount //accountBalance = accountBalance + deposit
			fmt.Println("Balance Updated: ", accountBalance)
			transaction := fmt.Sprintf("Deposit: %v ", depositAmount)
			transactionHistory = append(transactionHistory, transaction)
			fmt.Printf("Last 5 Transactions: %v \n", transactionHistory)
		case 3:
			var withdrawalAmount int
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

			transaction := fmt.Sprintf("Withdrawal: %v", withdrawalAmount)
			transactionHistory = append(transactionHistory, transaction)
			fmt.Printf("Last 5 Transaction : %v \n", transactionHistory)
		case 4:
			fmt.Println("Your Last Transactions:", transactionHistory)

			numTransaction := len(transactionHistory)
			start := numTransaction - 5
			if start < 0 {
				start = 0
			}
			for i := start; i < numTransaction; i++ {
				fmt.Println(transactionHistory[i])
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
