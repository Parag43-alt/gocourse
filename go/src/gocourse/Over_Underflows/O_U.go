package main

import "fmt"

func main() {
	// Unsigned integer (uint8: 0 to 255)
	var smallUint uint8 = 255
	fmt.Println("Original smallUint:", smallUint)

	// ✅ Overflow: 255 + 1 => wraps to 0
	smallUint = smallUint + 1
	fmt.Println("After Overflow (uint8):", smallUint)

	// Signed integer (int8: -128 to 127)
	var smallInt int8 = -128
	fmt.Println("\nOriginal smallInt:", smallInt)

	// ✅ Underflow: -128 - 1 => wraps to 127
	smallInt = smallInt - 1
	fmt.Println("After Underflow (int8):", smallInt)

	// Another signed overflow
	var bigInt int8 = 127
	fmt.Println("\nOriginal bigInt:", bigInt)

	// ✅ Overflow: 127 + 1 => wraps to -128
	bigInt = bigInt + 1
	fmt.Println("After Overflow (int8):", bigInt)
}
