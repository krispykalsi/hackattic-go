package slvr

type Solver interface {
	Solve([]byte) []byte
}
