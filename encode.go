package qb

import (
	"encoding/binary"
	"io"
)

func Encode(w io.Writer, in *QbModel) (err error) {
	//Validate matrices
	for i := range in.Matrices {
		if err = in.Matrices[i].validate(); err != nil {
			return
		}
	}

	//Encode header
	if err = in.encodeHeader(w); err != nil {
		return
	}

	//Encode matrix count
	if err = binary.Write(w, qbEndian, uint32(len(in.Matrices))); err != nil {
		return
	}

	//Encode matrices
	for i := range in.Matrices {
		if err = in.Matrices[i].encode(w, &in.Header); err != nil {
			return
		}
	}

	//All good
	return
}

func encodeRle(m *Matrix, w io.Writer, header *Header) (err error) {
	sliceSize := uint32(m.Size.X * m.Size.Y)
	for z := uint32(0); z < uint32(m.Size.Z); z++ {
		index := z * sliceSize
		slice := m.Content[index : index+sliceSize]

		for i := uint32(0); i < sliceSize; {

			//Find repeats
			numRepeats := uint32(1)
			for i2 := i + 1; i2 < sliceSize && slice[i] == slice[i2]; i2++ {
				numRepeats++
			}

			//Encode result
			out := header.ColorFormat.uint32(header.VisibilityMaskEncoding, slice[i])
			if numRepeats != 1 || out == flagCode || out == flagNextSlice {
				binary.Write(w, qbEndian, [3]uint32{
					flagCode,
					numRepeats,
					out,
				})
			} else {
				binary.Write(w, qbEndian, out)
			}

			i += numRepeats
		}

		//Slice done, encode end
		binary.Write(w, qbEndian, flagNextSlice)
	}

	//Looks good to me
	return
}
