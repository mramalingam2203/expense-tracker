package main

import "fmt"

const (
	APP_NAME     = "Personal Expense Tracker"
	MAX_EXPENSES = 10
)

var categories [MAX_EXPENSES]string
var amounts [MAX_EXPENSES]float64
var count int
var choice int


func main() {


ApplicationLoop:
	for {

		showMenu()

		fmt.Println("Enter your choice : ")
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			addExpense()
		case 2:
			listExpense()
		case 3:
			showTotal()
		case 4:
			saveWelcomeFile()
		case 0:
			break ApplicationLoop

		default:
			fmt.Println("Invalid Choice")
		}

}


}

func showMenu(){

			fmt.Println("==================================")
			fmt.Println(APP_NAME)
			fmt.Println("==================================")

			fmt.Println("1. Add Expense")
			fmt.Println("2. List Expenses")
			fmt.Println("3. Show Total")
			fmt.Println("4. Save File")
			fmt.Println("0. Exit")
			fmt.Println()
}

func addExpense(){
	if count == MAX_EXPENSES {
		fmt.Println()
		fmt.Println("Max. no. of expenses reached")
		return
	} else {

		fmt.Println()
		fmt.Print("Enter Expense Category: ")
		fmt.Scanln(&categories[count])

		fmt.Print("Enter Expense Amount: ")
		fmt.Scanln(&amounts[count])

		if amounts[count] <= 0 {
			fmt.Println("Amount must be > 0")
			return
		} else {
			count++
			fmt.Println()
			fmt.Println("Expense added successfully")
		}
	}
}


func listExpense(){

	if count == 0 {
		fmt.Println("No expenses in record")
		return
	}

	for i := 0; i < count; i++ {
			fmt.Println()
			fmt.Println("Expense", i+1)
			fmt.Println("Category", categories[i])
			fmt.Printf("Amount 💲%.2f\n", amounts[i])
		}

}


func showTotal(){
	if count == 0 {
		fmt.Println("No expenses in record")
		return
	}
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


func saveWelcomeFile(){

	file, err := os.Create("expenses.txt")

	if err != nil{
		fmt.Println("Unable to create file")
		return
	}

	fmt.Fprintln(file, "Personal Expense Tracker")
	fmt.Fprintln(file, "Welcome!")
	fmt.Fprintln(file, "This is my first file")

	file.Close()

	fmt.Println()


}
