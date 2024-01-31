package qb

import (
	"encoding/binary"
	"fmt"
	"image/color"
	"io"
)

type ColorFormat uint32

const (
	ColorFormat_RGBA ColorFormat = iota
	ColorFormat_BGRA
)

func (c ColorFormat) validate() error {
	if c != ColorFormat_RGBA && c != ColorFormat_BGRA {
		return fmt.Errorf("go-qb: unknown color format '%d'", c)
	}
	return nil
}

func (c ColorFormat) decodeColor(r io.Reader, v VisibilityMaskEncoding, out *color.RGBA) (err error) {
	err = binary.Read(r, qbEndian, out)

	//Swap R and B components
	if c == ColorFormat_BGRA {
		b := out.B
		out.B = out.R
		out.R = b
	}

	//Flatten visibility
	if v == VisibilityMaskEncoding_binary && out.A > 0 {
		out.A = 255
	}

	return
}
