package main

import (
	"os"
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	input := []string{"hello", "世界"}
	for i := 0; i < b.N; i++ {
		echo(os.Stdout, input)
	}
}
