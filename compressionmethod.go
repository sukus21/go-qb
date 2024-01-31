package qb

import (
	"fmt"
	"io"
)

type CompressionMethod uint32

const (
	CompressionMethod_none CompressionMethod = iota
	CompressionMethod_rle
)

func (c CompressionMethod) validate() error {
	if c != CompressionMethod_none && c != CompressionMethod_rle {
		return fmt.Errorf("go-qb: unknown compression method '%d'", c)
	}
	return nil
}

func (c CompressionMethod) decodeMatrix(m *Matrix, r io.Reader, header *Header) (err error) {
	//No compression
	if c == CompressionMethod_none {
		for i := range m.Content {
			if err = header.ColorFormat.decodeColor(r, header.VisibilityMaskEncoding, &m.Content[i]); err != nil {
				return
			}
		}
	}

	//RLE compression
	if c == CompressionMethod_rle {
		return decodeRle(m, r, header)
	}

	//unreachable
	return
}

func (c CompressionMethod) encodeMatrix(m *Matrix, w io.Writer, header *Header) (err error) {
	//No compression
	if c == CompressionMethod_none {
		for i := range m.Content {
			if err = header.ColorFormat.encodeColor(w, header.VisibilityMaskEncoding, m.Content[i]); err != nil {
				return
			}
		}
	}

	//RLE compression
	if c == CompressionMethod_rle {
		return encodeRle(m, w, header)
	}

	//unreachable
	return
}
