# Task: Flea circus

A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random (usually 4 possibilities, except for fleas on the edge of the grid or at the corners).

What is the expected number of unoccupied squares after 50 rings of the bell? Give your answer rounded to six decimal places.

# Solution

The script runs `100` simulations of flea circus concurrently inside its own goroutine. Each simulation returns the number of unoccupied squares after 50 bell rings.

Technically, fleas and rings do not exist in the implementation. The simulation takes every square coordinate, performs "jumps" into possible adjacent squares, and records the result. The result of each jump is stored into a map to make sure it's unique. In the end, the length of the map represents the number of occupied squares. After subtracting this number from a total number of squares we acquire the number of free squares.

# Execution

```bash
go run main.go
```

###### Result:

```bash
Expected average: 331.079987
```

# Benchmarking

```bash
go test -bench=. -benchtime=5s -benchmem
```

###### Result:

```bash
goos: darwin
goarch: amd64
pkg: flea-circus
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkRunSimulations-12             6         857128091 ns/op        285255842 B/op  17611066 allocs/op
PASS
ok      flea-circus     6.037s
```
