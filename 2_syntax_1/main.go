package main

import "fmt"

func main() {
	fmt.Println("Fasten your seat belts, here we GO ...")

	foo()

	fmt.Println("Some other log")

	for i := 0; i < 10000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm over here foo-ing around")
}
