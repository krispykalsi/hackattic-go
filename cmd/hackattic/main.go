//go:build !submit

package main

import "log"

func main() {
	data := haClient.FetchProblemData(challenge)
	solution := solver.Solve(data)
	log.Println(string(solution))
}
