package main

import (
	"fmt"
	"time"
)

var fibSlice [51]int

func main() {
	t1 := time.Now()
	for i := 0; i < 50; i++ {
		fmt.Println(fibonacciWithCache(i))
	}
	t2 := time.Now()

	t := t2.Sub(t1)
	fmt.Println(t)
}

func fibonacciWithCache(n int) int {
	if fibSlice[n] > 0 {
		return fibSlice[n]
	}
	if n <= 2 {
		return 1
	}
	res := fibonacciWithCache(n-1) + fibonacciWithCache(n-2)
	fibSlice[n] = res
	return res
}

// no cache
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
