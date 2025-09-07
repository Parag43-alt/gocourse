package main

func main() {

	// Arithmetic Operators: +, -, *, /, %
	var a, b int = 21, 10

	var c int

	// Addition
	c = a + b
	println("Addition:", c)

	// Subtraction
	c = a - b
	println("Subtraction:", c)

	// Multiplication
	c = a * b
	println("Multiplication:", c)

	// Division
	c = a / b
	println("Division:", c)

	// Modulus  is for remainder
	c = a % b
	println("Modulus:", c)

	// Increment Operator
	a++
	println("Increment:", a)

	// Decrement Operator
	a--
	println("Decrement:", a)

	const p float64 = 22 / 7
	println("Value of p:", p) // Note: This will print 3 due to integer division

	const P float64 = 22.0 / 7.0

	println("Value of P:", P) // This will print a more precise value of pi

}
