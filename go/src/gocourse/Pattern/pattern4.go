// the logic behind this pattern is:
// 1. The first loop is for the number of rows.

// 2. The second loop is for printing spaces before the stars in each row.
// the working of this loop is that it starts from the current row number and goes up to n (the total number of rows).
// the working of i loop is it print rows from 1 to n
// the working of j loop is that it prints (n-i) spaces before the stars in each row.
// let us say n = 5
// when i = 1, j loop will run from 1 to 5, so it will print 4 spaces
// when i = 2, j loop will run from 2 to 5, so it will print 3 spaces
// when i = 3, j loop will run from 3 to 5, so it will print 2 spaces

// now the number of spaces before the stars in each row is decreasing by 1 in each row. That why we are using j = i; j < n; j++ in the second loop.

// 3. The third loop is for printing stars in each row.
// the working of this loop is that it starts from 1 and goes up to (2*i - 1).
// the working of k loop is it prints (2*i - 1) stars in each row.
// let us say n = 5
// when i = 1, k loop will run from 1 to (2*1 - 1) = 1, so it will print 1 star
// when i = 2, k loop will run from 1 to (2*2 - 1) = 3, so it will print 3 stars

// so the number of stars in each row is increasing by 2 in each row. That why we are using k = 1; k <= (2*i - 1); k++ in the third loop.

// 4. Finally, we print a new line after each row is printed.
package main

import (
	"fmt"
)

func pattern4() {

	var n int

	fmt.Println("How many rows you want in pyramid :")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {

		for j := i; j < n; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= (2*i - 1); k++ {

			fmt.Print("*")

		}
		fmt.Println()
	}

}
