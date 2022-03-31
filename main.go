package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// size of square grid
	size = 30
	// amount of bell rings
	bellRings = 50
	// number of simulations
	simulations = 100
)

// moves contain a list of possible coordinates adjustments
var moves = [][]int{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	average := runSimulations(simulations)
	fmt.Printf("Expected average: %.6f\n", average)
}

// runSimulations executes *simulationsCount* number of flea circus simulations.
// Every simulation is running concurrently inside its own goroutine.
// The result of each goroutine is a number of free squares left after the simulation is done.
// The function returns the expected number of unoccupied squares.
func runSimulations(simulationsCount int) float32 {
	freeSquares := make(chan int)
	var wg sync.WaitGroup
	wg.Add(simulationsCount)
	for i := 0; i < simulationsCount; i++ {
		go func() {
			simulate(freeSquares)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(freeSquares)
	}()

	var result int
	for n := range freeSquares {
		result += n
	}
	return float32(result) / float32(simulationsCount)
}

// simulate performs a simulation by doing jumps for each flea.
func simulate(out chan<- int) {
	occupiedSquares := make(map[string]struct{})
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			i, j := jump(i, j)
			occupiedSquares[fmt.Sprintf("%d-%d", i, j)] = struct{}{}
		}
	}
	out <- size*size - len(occupiedSquares)
}

// jump performs *bellRings* number of jumps starting with [i, j] position.
// The function returns resulting square coordinates.
func jump(i, j int) (int, int) {
	for n := 0; n < bellRings; n++ {
		adjacentSquares := findAdjacentSquares(i, j)
		square := adjacentSquares[rand.Intn(len(adjacentSquares))]
		i, j = square[0], square[1]
	}
	return i, j
}

// findAdjacentSquares lookups possible adjacent squares
func findAdjacentSquares(x, y int) [][]int {
	adjacentSquares := make([][]int, 0, 4)

	var i, j int
	for _, move := range moves {
		i = x + move[0]
		j = y + move[1]
		if onGrid(i, j) {
			adjacentSquares = append(adjacentSquares, []int{i, j})
		}
	}
	return adjacentSquares
}

// onGrid returns true if the coordinates of a square (i, j) are located on a grid.
// If coordinates point outside of the grid - it returns false.
func onGrid(i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= size || j >= size {
		return false
	}
	return true
}
