package help_me_unpack

import (
	"encoding/base64"
	"encoding/binary"
	"github.com/krispykalsi/hackattic/pkg/slvr/utils"
	"log"
	"math"
)

type helpMeUnpack struct{}

func New() *helpMeUnpack {
	return &helpMeUnpack{}
}

func (h helpMeUnpack) Solve(data []byte) []byte {
	p := &problem{}
	utils.FromJson(data, p)

	byteArr := decodeBase64(p.Bytes)
	s := solution{
		Int:             extractInt(byteArr),
		UInt:            extractUInt(byteArr),
		Short:           extractShort(byteArr),
		Float:           extractFloat(byteArr),
		Double:          extractDouble(byteArr),
		BigEndianDouble: extractBigEndianDouble(byteArr),
	}
	return utils.ToJson(s)
}

func extractInt(byteArr []byte) int32 {
	numBytes := byteArr[0:4]
	bits := binary.LittleEndian.Uint32(numBytes)
	return int32(bits)
}

func extractUInt(byteArr []byte) uint {
	numBytes := byteArr[4:8]
	num := binary.LittleEndian.Uint32(numBytes)
	return uint(num)
}

func extractShort(byteArr []byte) int16 {
	numBytes := byteArr[8:10]
	bits := binary.LittleEndian.Uint16(numBytes)
	return int16(bits)
}

func extractFloat(byteArr []byte) float64 {
	numBytes := byteArr[12:16]
	bits := binary.LittleEndian.Uint32(numBytes)
	return float64(math.Float32frombits(bits))
}

func extractDouble(byteArr []byte) float64 {
	numBytes := byteArr[16:24]
	bits := binary.LittleEndian.Uint64(numBytes)
	return math.Float64frombits(bits)
}

func extractBigEndianDouble(byteArr []byte) float64 {
	numBytes := byteArr[24:32]
	bits := binary.BigEndian.Uint64(numBytes)
	return math.Float64frombits(bits)
}

func decodeBase64(code string) []byte {
	byteArr, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		log.Fatalf("Couldn't decode the base64 string: %v", err)
	}
	return byteArr
}
