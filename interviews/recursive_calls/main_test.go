package main

import "testing"

func BenchmarkRecursiveNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursiveNumbers(0)
	}
}
