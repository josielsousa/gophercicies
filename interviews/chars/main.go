package main

import "fmt"

// Print from 1 to 100 without using any numbers in your code.
func printNumberWithouUsingAnyNumber() {
	aCharCode := 'a'
	one := int(aCharCode / aCharCode)

	tenChars := ".........."
	oneHundred := len(tenChars) * len(tenChars)

	for i := one; i <= oneHundred; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("# Print from 1 to 100 without using any numbers in your code.")
	printNumberWithouUsingAnyNumber()
}
