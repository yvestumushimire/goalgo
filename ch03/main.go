package main

import (
	"fmt"
	"log"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(fibonacci(144))

}

func fibonacci(n int) int {
	startNumbers := []int{0, 1, 2}
	switch {
	case n < 0:
		log.Fatal("Invalid in put")
	case slices.Contains(startNumbers, n):
		fmt.Println(n)
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
