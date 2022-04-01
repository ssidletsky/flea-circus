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
	// number of bell rings
	bellRings = 50
	// number of simulations
	simulations = 100
	// number of workers
	workers = 1

	// indexes of moves
	moveLeft  = 0
	moveUp    = 1
	moveRight = 2
	moveDown  = 3
)

// moves contain a list of possible coordinates adjustments
var moves = [][]int{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	average := runSimulations(simulations, workers)
	fmt.Printf("Expected average: %.6f\n", average)
}

// runSimulations executes *simulationsCount* number of flea circus simulations.
// Simulations may run concurrently if >1 *workers* provided.
// The result of each simulation is a number of free squares left after the simulation is done.
// The function returns the expected number of unoccupied squares.
func runSimulations(simulationsCount int, workers int) float32 {
	jobsc := jobs(simulationsCount)

	freeSquares := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for range jobsc {
				freeSquares <- simulate()
			}
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

// jobs returns a channel which will be available to read for as many simulation needs to be done.
func jobs(count int) <-chan struct{} {
	out := make(chan struct{}, count)
	for i := 0; i < count; i++ {
		out <- struct{}{}
	}
	close(out)
	return out
}

// simulate performs a simulation by doing jumps for each flea and returns a number of free squares left after simulation
func simulate() int {
	occupiedSquares := make(map[string]struct{})
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			i, j := jump(i, j)
			occupiedSquares[fmt.Sprintf("%d-%d", i, j)] = struct{}{}
		}
	}
	return size*size - len(occupiedSquares)
}

// jump performs *bellRings* number of jumps starting with [i, j] position.
// The function returns resulting square coordinates.
func jump(i, j int) (int, int) {
	for n := 0; n < bellRings; n++ {
		i, j = nextSquare(i, j)
	}
	return i, j
}

// nextSquare calculates which moves are eligible for provided square and returns coordinates of randomly picked move.
func nextSquare(x, y int) (int, int) {
	validMoves := make([]int, 0, 4)
	if x > 0 {
		validMoves = append(validMoves, moveLeft)
	}
	if y > 0 {
		validMoves = append(validMoves, moveUp)
	}
	if x < size-1 {
		validMoves = append(validMoves, moveRight)
	}
	if y < size-1 {
		validMoves = append(validMoves, moveDown)
	}

	moveIndex := validMoves[rand.Intn(len(validMoves))]
	x += moves[moveIndex][0]
	y += moves[moveIndex][1]
	return x, y
}
