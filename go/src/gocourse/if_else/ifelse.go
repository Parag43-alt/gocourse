package main

import "fmt"

func ifelse() {
	var age int
	fmt.Println("Enter your age :")
	fmt.Scan(&age) // yaha Scan use kiya hai

	if age >= 18 {
		fmt.Println("You are an adult, you can ride.")
	} else if age <= 16 {
		fmt.Println("You are too young.")
	} else if age >= 14 && age <= 18 {
		fmt.Println("You are in teenage.")
	} else {
		fmt.Println("You cannot ride.")
		fmt.Printf("You have to wait %d years.\n", 18-age)
	}
}

func elseif() {
	var marks int
	fmt.Println("Please enter your marks :")
	fmt.Scan(&marks) // yaha bhi Scan use kiya hai (Scanln ki jagah)

	if marks >= 90 {
		fmt.Println("You are in the Merit List.")
	} else if marks >= 65 {
		fmt.Println("You passed with 1st division.")
	} else if marks >= 50 {
		fmt.Println("You passed with 2nd division.")
	} else if marks >= 33 {
		fmt.Println("You passed with 3rd division.")
	} else {
		fmt.Println("You are fail.")
	}
}
func main() {

	ifelse()
	elseif()
}
