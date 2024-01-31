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
