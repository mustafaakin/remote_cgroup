package lib

import "fmt"

type Device struct {
    name string
    minor int
    major int
}

func (d Device) String() string{
    return fmt.Sprintf("Name: %s major: %d minor: %d", d.name, d.major, d.minor)
}