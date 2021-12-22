package jotting_jwts

import (
	"github.com/krispykalsi/hackattic/pkg/slvr/jotting_jwts/server"
	"github.com/krispykalsi/hackattic/pkg/slvr/utils"
)

type jottingJwts struct{}

func New() *jottingJwts {
	return &jottingJwts{}
}

func (j jottingJwts) Solve(data []byte) []byte {
	p := &problem{}
	utils.FromJson(data, p)

	jv := server.NewJwtValidator(p.JwtSecret)
	go jv.Run()

	s := solution{AppUrl: "https://47ac-49-36-184-169.ngrok.io"}
	return utils.ToJson(s)
}
