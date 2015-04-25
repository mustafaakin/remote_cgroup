package lib

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

type Device struct {
	name  string
	minor int
	major int
}

func (d Device) String() string {
	return fmt.Sprintf("(%s,%d-%d)", d.name, d.major, d.minor)
}


var devicesByName = map[string]*Device{}
var devicesByMajor = map[int]map[int]*Device{}

func scanDevFiles() {
	files, _ := ioutil.ReadDir("/dev")

	for _, file := range files {
		Sys := file.Sys()

		if Sys != nil {
			stat := Sys.(*syscall.Stat_t)
			Rdev := stat.Rdev

			// get first 8 bits only
			minor := 0xFF & Rdev

			// subtracting minor just in case then shiftit it 8 bytes, according to spec it can be 12 bits at most but no check is made
			major := (Rdev - minor) >> 8

			dev := &Device{name: file.Name(), minor: int(minor), major: int(major)}

            // For fast lookups in the future
            devicesByName[dev.name] = dev
            if devicesByMajor[dev.major] == nil {
                // The reason is that
                devicesByMajor[dev.major] = make(map[int]*Device)
            }

            devicesByMajor[dev.major][dev.minor] = dev
		}
	}

	fmt.Printf("%#v\n", devicesByName)
    fmt.Printf("%#v\n", devicesByMajor)


    fmt.Printf("%#v\n", getDevByMinMajor(8,1)) // should be sda1
}

func getDevByMinMajor(major, minor int) *Device{
    return devicesByMajor[major][minor]
}

