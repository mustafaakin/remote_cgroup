package lib

import "fmt"

type Device struct {
    name string
    minor int
    major int
}

func (d Device) String() string{
    return fmt.Sprintf("(%s,%d-%d)", d.name, d.major, d.minor)
}