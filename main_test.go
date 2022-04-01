package main

import (
	"fmt"
	"testing"
)

func BenchmarkRunSimulations(b *testing.B) {
	workers := 1
	b.Run("singe worker", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSimulations(simulations, workers)
		}
	})
	workers = 2
	b.Run(fmt.Sprintf("%d-workers", workers), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSimulations(simulations, workers)
		}
	})
	workers = 4
	b.Run(fmt.Sprintf("%d-workers", workers), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSimulations(simulations, workers)
		}
	})
}
