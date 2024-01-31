package qb

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
