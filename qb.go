package qb

import "encoding/binary"

const (
	flagNextSlice uint32 = 6
	flagCode      uint32 = 2
)

var qbEndian = binary.LittleEndian

type QbModel struct {
	Header
	Matrices []Matrix
}

type Vec3 struct {
	X, Y, Z int32
}
