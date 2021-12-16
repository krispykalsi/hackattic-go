package slvr

import (
	"github.com/krispykalsi/hackattic/pkg/ha"
)

func ByName(challenge ha.ChallengeName) Solver {
	switch challenge {
	case ha.ReadingQr:
		return readingQr{}
	}
	return nil
}
