package main

import (
	"fmt"
	"os"
)

const options string = "(W)ithdraw, (D)eposit, (V)iew, (E)xit"

var balance int64 = 0

func main() {
	fmt.Println("Welcome to the Bank")
	fmt.Println("What would type of transaction would you like to do?")
	fmt.Println(options, ":")
	fmt.Println()
	var input string
	fmt.Scanln(&input)
	if input == "W" {
		fmt.Println("You are making a WITHDRAWAL.")
		withdraw()
	} else if input == "D" {
		fmt.Println("You are making a DEPOSIT.")
	} else if input == "V" {
		fmt.Println("You are VIEWING your balance!")
	} else if input == "E" {
		fmt.Println("You are EXITING!")
		os.Exit(1)
	} else {
		fmt.Println("[ERROR] Invalid Option : \nYou must pick from: ", options)
	}
}

func withdraw() {

}
