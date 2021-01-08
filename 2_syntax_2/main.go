package main

import "fmt"

func main() {
	// DECLARE a variable and ASSIGN a VALUE
	// short declaration operator
	x := 77
	fmt.Println(x)

	// long declaration operator
	var y = 23
	fmt.Println(y)

	z := x + y
	fmt.Println(z)

	var zz = x - y
	fmt.Printf("%T\n" , zz)

	
}