package main;

import (
	"fmt";
	"runtime";
)

// custom type
type myOwnType int
var u myOwnType

func main() {

	// integer
	var y int = 23
	fmt.Printf("%T", y)
	
	// string
	var name string = "James Bond"
	fmt.Printf("%T", name)
	
	// boolean
	var isBool bool = true
	fmt.Printf("%T\n" , isBool)
	
	// float 32
	var x float32 = 99.99
	fmt.Printf("%T", x)
	
	// float 64
	var k float64 = 99.99
	fmt.Printf("%T", k)

	// using myOwnType custom type
	u = 5
	y = int(u) // convert my custom type to native int

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}