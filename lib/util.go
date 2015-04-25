package lib

import (
	"fmt"
)

var devices = map[string]*Device{}

func Start() {
	fmt.Println("Hello from Util")

	scanDevFiles()
	findCgroupMountPoints()

}
