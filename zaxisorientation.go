package qb

import "fmt"

type ZaxisOrientation uint32

const (
	ZaxisOrientation_lefthanded ZaxisOrientation = iota
	ZaxisOrientation_righthanded
)

func (z ZaxisOrientation) validate() error {
	if z != ZaxisOrientation_lefthanded && z != ZaxisOrientation_righthanded {
		return fmt.Errorf("go-qb: unknown z-axis orientation '%d'", z)
	}
	return nil
}
