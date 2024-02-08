package qb

import (
	"encoding/binary"
	"fmt"
	"image/color"
	"io"
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

func (m *Matrix) decode(r io.Reader, h *Header) (err error) {
	//Read name
	buf := [255]byte{}
	if _, err = r.Read(buf[:1]); err != nil {
		return
	}
	strBuf := buf[:buf[0]]
	if _, err = r.Read(strBuf); err != nil {
		return
	}
	m.Name = string(strBuf)

	//Read size and position
	if err = binary.Read(r, qbEndian, &m.Size); err != nil {
		return
	}
	if err = binary.Read(r, qbEndian, &m.Position); err != nil {
		return
	}

	//Read contents
	m.Content = make([]color.RGBA, m.Size.X*m.Size.Y*m.Size.Z)
	return h.CompressionMethod.decodeMatrix(m, r, h)
}

func (m *Matrix) encode(w io.Writer, h *Header) (err error) {
	//Write name
	if _, err = w.Write([]byte{byte(len(m.Name))}); err != nil {
		return
	}
	if _, err = w.Write([]byte(m.Name)); err != nil {
		return
	}

	//Write size and position
	if err = binary.Write(w, qbEndian, m.Size); err != nil {
		return
	}
	if err = binary.Write(w, qbEndian, m.Position); err != nil {
		return
	}

	//Write contents
	return h.CompressionMethod.encodeMatrix(m, w, h)
}

func (m *Matrix) pointWithin(x, y, z int32) bool {
	return m.Bounds().Contains(Point{X: int(x), Y: int(y), Z: int(z)})
}

func (m *Matrix) getIndex(x, y, z int) int {
	index := int32(x) - m.Position.X
	index += (int32(y) - m.Position.Y) * m.Size.X
	index += (int32(z) - m.Position.Z) * m.Size.X * m.Size.Y
	return int(index)
}

func (m *Matrix) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *Matrix) Bounds() Cube {
	return Cube{
		Min: m.Position.Point(),
		Max: Point{
			X: int(m.Position.X + m.Size.X),
			Y: int(m.Position.Y + m.Size.Y),
			Z: int(m.Position.Z + m.Size.Z),
		},
	}
}

func (m *Matrix) Get(x, y, z int) color.Color {
	if !m.pointWithin(int32(x), int32(y), int32(z)) {
		return color.Transparent
	}

	return m.Content[m.getIndex(x, y, z)]
}

func (m *Matrix) Set(x, y, z int, c color.Color) {
	if !m.pointWithin(int32(x), int32(y), int32(z)) {
		return
	}

	//Flatten color to RGBA
	r, g, b, a := c.RGBA()
	m.Content[m.getIndex(x, y, z)] = color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}
