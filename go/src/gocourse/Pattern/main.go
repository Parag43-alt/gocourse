package main

import "fmt"

// Remove the incorrect import. If Pattern1 is defined in another file within the same package, no import is needed.
func main() {
	Pattern1()

	fmt.Println()

	Pattern2()
}
