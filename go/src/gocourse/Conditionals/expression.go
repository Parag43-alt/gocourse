package main

import "fmt"

func main() {

	// there are mainly 6 types od conditions are there
	// >  : greater than
	// <  : less than
	// >= : greater than equals to
	// <= : less than equals to
	// == : comparison : means equals to
	// != : not equals to

	val_1 := 20
	val_2 := 30

	// >
	a := val_1 > val_2

	fmt.Println(a)

	// <

	b := val_1 < val_2

	fmt.Println(b)

	// >=
	c := val_1 >= val_2

	fmt.Println(c)

	// <=
	d := val_1 <= val_2

	fmt.Println(d)

	// ==

	e := val_1 == val_2
	fmt.Println(e)

	// !=

	f := val_1 != val_2
	fmt.Println(f)

	// another way
	if int32(val_1)+10 == int32(val_2) {
		fmt.Println("Equal")
	}

}
