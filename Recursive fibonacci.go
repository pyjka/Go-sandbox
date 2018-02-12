package main

import "fmt"

func fib(a uint) uint {
	if a <= 2 {
		return 1
	}
	if a == 0 {
		return 0
	}

	return fib(a-1) + fib(a-2)
}

func main() {
	fmt.Println(fib(42))
}
