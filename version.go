package qb

import "fmt"

type Version struct {
	Major, Minor, Release, Build uint8
}

var versionCurrent = Version{
	Major:   1,
	Minor:   1,
	Release: 0,
	Build:   0,
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Release, v.Build)
}

func (v Version) validate() error {
	if v != versionCurrent {
		return fmt.Errorf("go-qb: unknown version '%s'", v)
	}
	return nil
}
