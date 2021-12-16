//go:build submit

package main

func main() {
	data := haClient.FetchProblemData(challenge)
	solution := solver.Solve(data)
	haClient.Submit(challenge, solution)
}
