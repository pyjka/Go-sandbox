package main

import "fmt"

func fib(a uint) uint {
	if a <= 1 {
		return 1
	}
	return fib(a-1) + fib(a-2)
}

func main() {
	fmt.Println(fib(42))
}
