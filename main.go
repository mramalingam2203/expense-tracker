package main

import "fmt"

const (
	APP_NAME     = "Personal Expense Tracker"
	MAX_EXPENSES = 10
)

func main() {

	var categories [MAX_EXPENSES]string
	var amounts [MAX_EXPENSES]float64
	var count int
	var choice int

ApplicationLoop:
	for {

		fmt.Println("==================================")
		fmt.Println(APP_NAME)
		fmt.Println("==================================")

		fmt.Println("1. Add Expense")
		fmt.Println("2. List Expenses")
		fmt.Println("3. Show Total")
		fmt.Println("0. Exit")
		fmt.Println()

		fmt.Println("Enter your choice : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if count == MAX_EXPENSES {
				fmt.Println()
				fmt.Println("Max. no. of expenses reached")
				fmt.Println()
			}else{

			fmt.Println()
			fmt.Print("Enter Expense Category: ")
			fmt.Scanln(&categories[count])

			fmt.Print("Enter Expense Amount: ")
			fmt.Scanln(&amounts[count])

			if amounts[count] <= 0 {
				fmt.Println("Amount must be > 0")
				break
			}else{
			count++
			fmt.Println()
			fmt.Println("Expense added successfully")
		}
		}

		case 2:
			fmt.Println()

			if count == 0 {
				fmt.Println("No expenses in record")
			} else {
				fmt.Println("---------Expenses-----------")
				for i := 0; i < count; i++ {
					fmt.Println()
					fmt.Println("Expenses", i+1)
					fmt.Println("Category", categories[i])
					fmt.Printf("Amount 💲%.2f\n", amounts[i])
				}

				fmt.Println()
				fmt.Println("===========================")
			}

		case 3:
			fmt.Println()

			if count == 0 {
				fmt.Println("No expenses in record")
			} else {
				var total float64

				for i := 0; i < count; i++ {
					total += amounts[i]
				}

				fmt.Println("===========================")
				fmt.Println("Total Expenses")
				fmt.Println("===========================")
				fmt.Printf("💲%.2f\n", total)
				fmt.Println()
			}

		case 0:
			fmt.Println()
			fmt.Println("Thank you for using ", APP_NAME)
			break ApplicationLoop

		default:
			fmt.Println()
			fmt.Println("Invalid choice. Please try again")
		}

	}

}
