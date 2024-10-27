package main

import "fmt"

func main() {
	numbers := [5]int{4, 5, 3, 6, 7}
	fmt.Println(Sum(numbers))

}

func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}