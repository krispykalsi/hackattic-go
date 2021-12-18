package mini_miner

import (
	"encoding/hex"
	"strconv"
)

type BitString string

func (b BitString) AsByteSlice() []byte {
	var out []byte
	var str string

	for i := len(b); i > 0; i -= 8 {
		if i-8 < 0 {
			str = string(b[0:i])
		} else {
			str = string(b[i-8 : i])
		}
		v, err := strconv.ParseUint(str, 2, 8)
		if err != nil {
			panic(err)
		}
		out = append([]byte{byte(v)}, out...)
	}
	return out
}

func (b BitString) AsHexSlice() []string {
	var out []string
	byteSlice := b.AsByteSlice()
	for _, b := range byteSlice {
		out = append(out, "0x"+hex.EncodeToString([]byte{b}))
	}
	return out
}
