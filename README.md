# Task: Flea circus

A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random (usually 4 possibilities, except for fleas on the edge of the grid or at the corners).

What is the expected number of unoccupied squares after 50 rings of the bell? Give your answer rounded to six decimal places.

# Solution

The script runs `100` simulations of flea circus concurrently inside its own goroutine. Each simulation returns the number of unoccupied squares after 50 bell rings.

Technically, fleas and rings do not exist in the implementation. The simulation takes every square coordinate, performs "jumps" into possible adjacent squares, and records the result. The result of each jump is stored into a map to make sure it's unique. In the end, the length of the map represents the number of occupied squares. After subtracting this number from a total number of squares we acquire the number of free squares.

###### Edit:

I have changed the concurency model of a progrem to a worker pool. Previously each simulation was running in its own goroutine. To check what is the best number of workers for the task I have updated the benchmars. The results shows, for the task without any I/O there is no advantage to run it concurently on a single processor system.

This conclution is also proven by cpu profiling. Once there is more then one goroutine, runtime starts sppent more time/memory to handle additional overhad to manage them.

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
BenchmarkRunSimulations
BenchmarkRunSimulations/singe_worker
BenchmarkRunSimulations/singe_worker-12                       10         107090039 ns/op         4702520 B/op      92004 allocs/op
BenchmarkRunSimulations/2-workers
BenchmarkRunSimulations/2-workers-12                           8         138092166 ns/op         4701061 B/op      91994 allocs/op
BenchmarkRunSimulations/4-workers
BenchmarkRunSimulations/4-workers-12                           5         228339260 ns/op         4705680 B/op      92004 allocs/op
PASS
ok      flea-circus     6.599s
```
