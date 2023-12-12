package aoc

type Solver interface {
	Solve() error
}

var Solutions = make(map[int]map[int]Solver)

func RegisterSolver(year int, solvers map[int]Solver) {
	Solutions[year] = solvers
}
