package qb

import (
	"fmt"
)

type Matrix struct {
	Name     string
	Size     Vec3
	Position Vec3
	Content  []color.RGBA
}

func (m *Matrix) validate() error {
	if len([]byte(m.Name)) >= 256 {
		return fmt.Errorf("go-qb: matrix name must be less than 256 bytes")
	}
	if len(m.Content) != int(m.Size.X*m.Size.Y*m.Size.Z) {
		return fmt.Errorf("go-qb: matrix contents do not match its size")
	}
	return nil
}
