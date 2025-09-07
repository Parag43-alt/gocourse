package main

import (
	"fmt"
)

func main() {

	// Loop from 1 to 5 : Simple for loop
	for i := 1; i <= 5; i++ {

		fmt.Println("Number:", i)
	}

	// Iterate over a collection

	numbers := []int{10, 20, 30, 40, 50, 60}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//continue statement
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(j) // Print only odd numbers
	}
	fmt.Println("This is Continue Statement Example ")

	// break statement
	for k := 1; k <= 10; k++ {
		if k == 3 {
			break // Exit the loop when k is 3
		}
		fmt.Println(k) // Print numbers from 1 to 2
	}

	fmt.Println("This is Break Statement Example ")

	// timepass

	for m := 1; m <= 10; m++ {
		if m%2 != 0 { // yah pe humne not equal to ka use kiya hai
			// yani odd numbers ko skip ki ya hai
			continue
		}
		if m == 6 { // break se mtlb ki jaise yaha pe 6 pe aayega loop ruk jayega aage nahi badhega mtln 8 aur 10 print nahi hoga
			break
		}
		fmt.Println(m)
	}

	fmt.Println("This is my example ")
}
