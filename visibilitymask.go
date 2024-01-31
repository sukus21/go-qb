package qb

import "fmt"

type VisibilityMaskEncoding uint32

const (
	VisibilityMaskEncoding_binary VisibilityMaskEncoding = iota
	VisibilityMaskEncoding_alpha
)

func (v VisibilityMaskEncoding) validate() error {
	if v != VisibilityMaskEncoding_binary && v != VisibilityMaskEncoding_alpha {
		return fmt.Errorf("go-qb: unknown visibility-mask encoding '%d'", v)
	}
	return nil
}
