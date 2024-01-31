package qb

type QbModel struct {
	Header
	Matrices []Matrix
}

type Vec3 struct {
	X, Y, Z uint32
}
