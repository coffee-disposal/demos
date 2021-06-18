package main

import (
	"fmt"
)

func main() {
	n := 10
	fib_n := fib(n)
	formatted_string := fmt.Sprintf("The %dth number of the fibonacci sequence is %d.", n, fib_n)
	fmt.Println(formatted_string)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
