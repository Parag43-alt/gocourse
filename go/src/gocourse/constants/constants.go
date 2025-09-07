package main

import "fmt"

var name string = "Parag"

// we cannot declare constant outside function using :=

func main() {

	fmt.Println("Declarating constant outside function :", name)

	// single constant declaration
	const lang = "Go"

	fmt.Println("language:", lang)

	// multiple constant declaration : called constant grouping using ()
	const (
		serverport = 8080
		serverhost = "localhost"
		inUse      = true
	)

	fmt.Println("Server Port:", serverport)
	fmt.Println("Server Host:", serverhost)
	fmt.Println("In Use:", inUse)
	// we cannot change the value once it assigned using const
}
