package slvr

type Solver interface {
	Solve(data []byte) []byte
}
