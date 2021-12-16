package main

import (
	"github.com/krispykalsi/hackattic/pkg/ha"
	"github.com/krispykalsi/hackattic/pkg/slvr"
	"os"
)

var challenge ha.ChallengeName

var haClient *ha.Client
var solver slvr.Solver

func init() {
	challenge = ha.HelpMeUnpack
	at := os.Getenv("HACKATTIC_ACCESS_TOKEN")

	haClient = ha.NewClient(at)
	solver = slvr.ByName(challenge)
}
