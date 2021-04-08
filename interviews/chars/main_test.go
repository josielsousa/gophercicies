package main

import "testing"

func BenchmarkPrintNumberWithouUsingAnyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printNumberWithouUsingAnyNumber()
	}
}
