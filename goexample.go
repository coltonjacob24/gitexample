package main

import (
	"fmt"
)

func main() {
	fmt.Println("100 Fibonacci numbers")
	for i := 0; i <= 100; i++ {
		fmt.Printf("[%d]%4d\n", i, nthFib(i))
	}
}

func nthFib(n int) int {
	if n <= 1 {
		return n
	}
	return nthFib(n-1) + nthFib(n-2)
}
