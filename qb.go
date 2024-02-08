package qb

import (
	"encoding/binary"
	"image/color"
)

const (
	flagNextSlice uint32 = 6
	flagCode      uint32 = 2
)

var qbEndian = binary.LittleEndian

type QbModel struct {
	Header
	Matrices []Matrix
}

func (q *QbModel) ColorModel() color.Model {
	return color.RGBAModel
}

func (q *QbModel) Get(x, y, z int) color.Color {
	//Find matrix at these coordinates
	for i := range q.Matrices {
		child := &q.Matrices[i]

		if child.pointWithin(int32(x), int32(y), int32(z)) {
			return child.Get(x, y, z)
		}
	}

	//Nothing found
	return color.Transparent
}

func (q *QbModel) Set(x, y, z int, c color.Color) {
	//Update children
	for i := range q.Matrices {
		q.Matrices[i].Set(x, y, z, c)
	}
}

type Vec3 struct {
	X, Y, Z int32
}

func (v Vec3) Point() Point {
	return Point{
		X: int(v.X),
		Y: int(v.Y),
		Z: int(v.Z),
	}
}
