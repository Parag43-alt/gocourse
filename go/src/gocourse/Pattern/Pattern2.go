package main

import "fmt"

func Pattern2() {
	// Pattern 2: Inverted Right-Angled Triangle

	rows := 5

	// outer loop
	for i := rows; i >= 1; i-- {
		// inner loop
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
