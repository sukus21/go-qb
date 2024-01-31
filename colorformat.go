package qb

import (
	"fmt"
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
