package main

import "fmt"

func Pattern1() {

	// Pattern 1: Right-Angled Triangle

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print("* ")

		}
		println()

	}

}
