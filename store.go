package main

import (
	"fmt"
	"os"
)

const options string = "(W)ithdraw, (D)eposit, (V)iew, (E)xit"

var balance float64 = 100.0

func main() {
	fmt.Println("Welcome to the Bank")
	fmt.Println("What would type of transaction would you like to do?")
	fmt.Println(options, ":")
	var input string
	fmt.Scanln(&input)
	if input == "W" {
		fmt.Println("You are making a WITHDRAWAL.")
		withdraw()
		fmt.Println("Your new balance is: $" + fmt.Sprint(balance))
	} else if input == "D" {
		fmt.Println("You are making a DEPOSIT.")
		deposit()
		fmt.Println("Your new balance is: $" + fmt.Sprint(balance))
	} else if input == "V" {
		fmt.Println("You are VIEWING your balance!")
		fmt.Println("Your current balance is: $" + fmt.Sprint(balance))
	} else if input == "E" {
		fmt.Println("You are EXITING!")
		os.Exit(1)
	} else {
		fmt.Println("[ERROR] Invalid Option : \nYou must pick from: ", options)
	}
}

func withdraw() {
	var amount float64
	fmt.Println("How much would you like to withdraw? : ")
	fmt.Scanln(&amount)
	if amount > balance{
		fmt.Println("You are trying to withdraw more than your current balance. Exiting.")
		os.Exit(1)
	}
	fmt.Println("Withdrawing $" + fmt.Sprint(amount))
	balance = balance - amount
}

func deposit() {
	var amount float64
	fmt.Println("How much would you like to deposit? : ")
	fmt.Scanln(&amount)
	if amount <= 0{
		fmt.Println("You are trying to deposit an invalid amount!")
		os.Exit(1)
	}
	fmt.Println("Depositing $" + fmt.Sprint(amount))
	balance = balance + amount
}