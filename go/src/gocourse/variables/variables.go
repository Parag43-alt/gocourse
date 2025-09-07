package main

import (
	"fmt"
)

func main() {
	// 1️⃣ Single variable declaration
	var name string = "Parag"
	var age int = 22
	fmt.Println("Name:", name, "Age:", age)

	// 2️⃣ Multiple variable declaration
	var city, country string = "Mumbai", "India"
	fmt.Println("City:", city, "Country:", country)

	// 3️⃣ Short declaration (:=)
	score := 95.5    // float64
	isPassed := true // bool
	char := 'A'      // rune (character)
	fmt.Println("Score:", score, "Passed:", isPassed, "Grade:", char)

	// 4️⃣ Zero value variables
	var zeroInt int
	var zeroString string
	var zeroBool bool
	fmt.Println("Zero Values -> Int:", zeroInt, "String:", zeroString, "Bool:", zeroBool)

	// 5️⃣ Constants
	const pi = 3.14159
	const appName = "GoCourse"
	fmt.Println("Pi:", pi, "App Name:", appName)

	// 6️⃣ Type conversion
	var intVal int = 10
	var floatVal float64 = float64(intVal) + 2.5
	fmt.Println("Converted Float:", floatVal)

	// 7️⃣ Multiple assignment
	x, y := 5, 10
	x, y = y, x // swap
	fmt.Println("Swapped Values -> x:", x, "y:", y)

}
