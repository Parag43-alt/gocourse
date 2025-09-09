package main

import "fmt"

func Pattern1() {

	// Pattern 1: Right-Angled Triangle

	// Outer loop: Iterates over each row of the triangle.
	// 'i' represents the current row number, starting from 1 up to 5.
	for i := 1; i <= 5; i++ {

		// Inner loop: Responsible for printing '*' characters in each row.
		// For each row 'i', it prints 'i' number of '*' characters.
		// 'j' starts from 1 and goes up to the current row number 'i'.
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
			// Prints a '*' followed by a space for better readability.
		}

		println()
		// After printing all '*' in the current row, move to the next line.
	}

}

// Outer loop explanation:
// Controls the number of rows in the triangle (from 1 to 5).
// For each iteration, it represents a new row in the pattern.

// Inner loop explanation:
// Prints '*' for each column in the current row.
// The number of '*' printed equals the current row number (i).
// This creates the right-angled triangle effect, where each row has one more '*' than the previous.
