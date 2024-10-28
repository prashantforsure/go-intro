package main

import "fmt"

func modifyValue(ptr *int) {
	*ptr = *ptr + 10
}

func main() {
	var x int = 10
	var ptr *int = &x // you will get the address of x
	fmt.Println("print the x", x)
	fmt.Println("print the address of x", &x)

	*ptr = 30
	fmt.Println(" x after change", x)
	modifyValue(&x)          // pass the address of x
	fmt.Println("After:", x) // x is modified by the function
}
