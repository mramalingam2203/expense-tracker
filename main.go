package main

import "fmt"

const appName = "Personal Expense Tracker"
const MAX_EXPENSES = 10

func main() {

	var categories [MAX_EXPENSES]string
	var amounts [MAX_EXPENSES]float64
	var count int

	categories[0] = "Food"
	amount[0] = 100.00

	fmt.Println("==================================")
	fmt.Println(appName)
	fmt.Println("==================================")

	fmt.Println()

	fmt.Println("Today's Expense")
	fmt.Println()
/*
	fmt.Print("Enter Expense Category :")
	fmt.Scanln(&category)

	fmt.Print("Enter Expense Amount   :")
	fmt.Scanln(&amount)
	fmt.Println()
*/
	fmt.Println("Category :", category[0])
	fmt.Println("Amount   :", amount[0])

	fmt.Println("Project Started Successfully!")
}
