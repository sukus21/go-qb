package qb

import (
	"fmt"
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
