package main

import (
    "fmt"
    "io/ioutil"
    "syscall"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Device struct {
    name string
    minor int
    major int
}

func (d Device) String() string{
    return fmt.Sprintf("Name: %s major: %d minor: %d", d.name, d.major, d.minor)
}

func main() {
    fmt.Println("Hello World")

    files, err := ioutil.ReadDir("/dev")
    check(err)

    for _, file := range files {
        Sys := file.Sys()

        if  Sys != nil  {
            stat := Sys.(*syscall.Stat_t)
            Rdev := stat.Rdev

            minor := 0xFF & Rdev // get first 8 bits only
            major := (Rdev - minor) >> 8 // subtracting minor just in case

            dev := Device{name: file.Name(), minor: int(minor), major: int(major)}
            fmt.Printf("%s \n", dev)
        }
    }
}
