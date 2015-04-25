package lib

import (
	"fmt"
)

func Start() {
	fmt.Println("Hello from Util")

	scanDevFiles()
	findCgroupMountPoints()
}