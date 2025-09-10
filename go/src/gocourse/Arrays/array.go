package main

import "fmt"

func main() {

	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println(numbers)
	fmt.Println(numbers[1])
	fmt.Println(len(numbers)) //length of array

	// Infer length of array
	// yaha pe humne size nhi diya hai array ka
	var fruits = []string{"apple ", "mango ", "watermelons"}
	fmt.Println(fruits[2])

	//short hand declaration of array

	cars := [10]string{" bmw ", " ferrari ", " mayback ", " audi", " lamborghini", " rolls royce", " porsche", " bugatti", " mercedes", " honda"}
	fmt.Println(cars)
	fmt.Println(cars[0:5])

	// shorthand declaration of array with infer length

	veggies := [...]string{"potato", "tomato", "cabbage", "onion"}
	fmt.Println(veggies)

	//changing item in an array
	cars[0] = " tata "
	fmt.Println(cars[0])

	// multi dimensional array

	var multiDimensionalArray = [2][4]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
	}
	fmt.Println(multiDimensionalArray)

	// default assignemnt  of array

	var number = [5]int{1, 2}
	fmt.Println(number)

}
