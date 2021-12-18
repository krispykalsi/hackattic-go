package slvr

import (
	"github.com/krispykalsi/hackattic/pkg/ha"
	"github.com/krispykalsi/hackattic/pkg/slvr/help_me_unpack"
	"github.com/krispykalsi/hackattic/pkg/slvr/mini_miner"
	"github.com/krispykalsi/hackattic/pkg/slvr/reading_qr"
)

func ByName(challenge ha.ChallengeName) Solver {
	switch challenge {
	case ha.ReadingQr:
		return reading_qr.New()
	case ha.HelpMeUnpack:
		return help_me_unpack.New()
	case ha.MiniMiner:
		return mini_miner.New()
	}
	return nil
}
