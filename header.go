package qb

import (
	"encoding/binary"
	"io"
)

type Header struct {
	Version
	ColorFormat
	ZaxisOrientation
	CompressionMethod
	VisibilityMaskEncoding
}

func (h *Header) validate() (err error) {
	if err = h.Version.validate(); err != nil {
		return
	}
	if err = h.ColorFormat.validate(); err != nil {
		return
	}
	if err = h.ZaxisOrientation.validate(); err != nil {
		return
	}
	if err = h.CompressionMethod.validate(); err != nil {
		return
	}
	if err = h.VisibilityMaskEncoding.validate(); err != nil {
		return
	}
	return nil
}

func (h *Header) decodeHeader(r io.Reader) error {
	if err := binary.Read(r, qbEndian, h); err != nil {
		return err
	}
	return h.validate()
}

func (h *Header) encodeHeader(w io.Writer) (err error) {
	if err = h.validate(); err != nil {
		return
	}
	return binary.Write(w, qbEndian, *h)
}
