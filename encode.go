package qb

import (
	"encoding/binary"
	"io"
)

func decomposeStitch(voxel VoxelGrid) (out []VoxelGrid) {
	if stitch, ok := voxel.(VoxelStitch); ok {
		for _, part := range stitch.Parts() {
			out = append(out, decomposeStitch(part)...)
		}
	} else {
		out = []VoxelGrid{voxel}
	}
	return
}

func Encode(w io.Writer, in VoxelGrid, settings *Header) (err error) {
	//Break model into its pieces
	models := decomposeStitch(in)

	//Encode header
	if err = settings.encodeHeader(w); err != nil {
		return
	}

	//Encode matrix count
	if err = binary.Write(w, qbEndian, uint32(len(models))); err != nil {
		return
	}

	//Encode matrices
	for i := range models {
		if err = in.Matrices[i].encode(w, settings); err != nil {
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
