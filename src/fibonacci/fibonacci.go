package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
