package mini_miner

import (
	"crypto/sha256"
	"fmt"
	"github.com/krispykalsi/hackattic/pkg/slvr/utils"
	"strings"
)

type miniMiner struct{}

func New() *miniMiner {
	return &miniMiner{}
}

func (m miniMiner) Solve(data []byte) []byte {
	p := &problem{}
	utils.FromJson(data, p)

	bytes := (p.Difficulty + p.Difficulty%8) / 8
	zeroes := strings.Repeat("0", p.Difficulty)
	nonce := 0
	for {
		p.Block.Nonce = nonce
		b := utils.ToJson(p.Block)
		h := sha256.Sum256(b)
		bits := ""
		for i := 0; i < bytes; i++ {
			bits += fmt.Sprintf("%08b", h[i])
		}
		if bits[:p.Difficulty] == zeroes {
			break
		}
		nonce++
	}

	s := solution{Nonce: nonce}
	return utils.ToJson(s)
}
