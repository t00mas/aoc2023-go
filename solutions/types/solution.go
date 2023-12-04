package types

type Solver interface {
	Init()
	Solve(int) int
}
