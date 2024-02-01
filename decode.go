package qb

import (
	"encoding/binary"
	"io"
)

func Decode(r io.Reader) (out *QbModel, err error) {
	out = &QbModel{}

	//Read header
	if err = out.decodeHeader(r); err != nil {
		return
	}

	//Read matrix count
	matrixCount := uint32(0)
	if err = binary.Read(r, qbEndian, &matrixCount); err != nil {
		return
	}

	//Read matrices
	out.Matrices = make([]Matrix, matrixCount)
	for i := range out.Matrices {
		if err = out.Matrices[i].decode(r, &out.Header); err != nil {
			return
		}
	}

	//All is good
	return
}

func decodeRle(m *Matrix, r io.Reader, header *Header) (err error) {
	buf := [8]byte{}

	for z := uint32(0); z < uint32(m.Size.Z); z++ {
		index := z * uint32(m.Size.X*m.Size.Y)
		for {
			//Read instruction
			data := uint32(0)
			if err = binary.Read(r, qbEndian, &data); err != nil {
				return
			}

			//This slice is done now
			if data == flagNextSlice {
				break
			}

			//RLE encoded data
			if data == flagCode {
				if _, err = r.Read(buf[:8]); err != nil {
					return
				}
				count := qbEndian.Uint32(buf[0:4])
				data = qbEndian.Uint32(buf[4:8])
				for i := uint32(0); i < count; i++ {
					if err = header.ColorFormat.decodeColor(r, header.VisibilityMaskEncoding, &m.Content[index]); err != nil {
						return err
					}
					index++
				}
				continue
			}

			//Single datapoint
			if err = header.ColorFormat.decodeColor(r, header.VisibilityMaskEncoding, &m.Content[index]); err != nil {
				return err
			}
			index++
		}
	}

	//Looks good to me
	return
}
