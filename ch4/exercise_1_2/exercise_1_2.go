package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 100)
	for i := 0; i < 100; i++ {
		nums[i] = rand.Intn(100)
	}

	for i, v := range nums {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Printf("Six! index: %d, number: %d\n", i, v)
		case v%2 == 0:
			fmt.Printf("Two! index: %d, number: %d\n", i, v)
		case v%3 == 0:
			fmt.Printf("Three! index: %d, number: %d\n", i, v)
		default:
			fmt.Println("Nevermind.")
		}
	}
}
