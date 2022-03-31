package main

import "testing"

func BenchmarkRunSimulations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runSimulations(simulations)
	}
}
