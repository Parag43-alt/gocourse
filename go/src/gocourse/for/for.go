package main

import "fmt"

func main() {

	// Golang has only one looping construct, the for loop.

	// sum of first 10 digits
	var sum int = 0
	var i int

	for i = 1; i <= 10; i++ {
		sum = sum + i
	}

	fmt.Println("Sum of first 10 digits is:", sum)

}
