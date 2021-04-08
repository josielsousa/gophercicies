package main

import "fmt"

func recursiveNumbers(n int) int {
	n++
	fmt.Println(n)

	if n < 100 {
		return recursiveNumbers(n)
	}

	return n
}

func main() {
	fmt.Println("# Print 1...100 without using any for/while/loops")
	recursiveNumbers(0)
}
