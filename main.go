package main

import "fmt"

const appName = "Personal Expense Tracker"
const MAX_EXPENSES = 10

func main() {

	var categories [MAX_EXPENSES]string
	var amounts [MAX_EXPENSES]float64
	var count int = 0

	categories[0] = "Food"
	amounts[0] = 100.00

	categories[1] = "Fuel"
	amounts[1] = 300.00

	categories[2] = "Movie"
	amounts[2] = 200.00




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
	fmt.Println("Category :", categories[count])
	fmt.Println("Amount   :", amounts[count])
	fmt.Println()
	count++

	fmt.Println("Category :", categories[count])
	fmt.Println("Amount   :", amounts[count])
	fmt.Println()
	count++

	fmt.Println("Category :", categories[count])
	fmt.Println("Amount   :", amounts[count])
	fmt.Println()
	count++


	fmt.Println("Project Started Successfully!")
}
