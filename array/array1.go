package main

import "fmt"

func array() {
	numbers1 := []int{1, 2, 3, 4, 5}
	for index, value := range numbers1 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

func arrayString() {
	colors := map[string]string{"R": "Red", "G": "Green", "B": "Blue"}
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func charLoop() {
	color := "happy"
	for index, char := range color {
		fmt.Printf("index: %d, char: %c\n", index, char)
	}
}
