package main

import "fmt"

const {
	appName = "Personal Expense Tracker"
	const MAX_EXPENSES = 10
}

func main() {

	var categories [MAX_EXPENSES]string
	var amounts [MAX_EXPENSES]float64
	var count int

	var choice int


 	ApplicationLoop:
		for{

			fmt.Println("==================================")
			fmt.Println(appName)
			fmt.Println("==================================")

			fmt.Println("1. Add Expense")
			fmt.Println("2. List Expenses")
			fmt.Println("3. Show Total")
			fmt.Println("0. Exit")
			fmt.Println()

			fmt.Println("Enter your choice : ")
			fmt.Scanln(&choice)
switch choice
{
		case 1:
		case 2:
		case 3:
		case 4:
}

		}



	fmt.Println("Today's Expense")
	fmt.Println()

	var choice string = "N"

	for{
		fmt.Print("Enter Category    : ")
		fmt.Scanln(&categories[count])

		fmt.Print("Enter Amount      : ")
		fmt.Scanln(&amounts[count])
		count++

		fmt.Print("Add another expense? (Y/N): ")
		fmt.Scanln(&choice)

		if choice == "N" || choice == "n"{
			break
		}

	}


	for i := 0; i < count; i++{

			fmt.Println("Category :", categories[i])
			fmt.Println("Amount   :", amounts[i])

	}




	fmt.Println("Code run successfully!")
}
