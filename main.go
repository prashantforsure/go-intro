package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Add(5, 6))

}

func Add(x, y int) int {
	return x + y
}
