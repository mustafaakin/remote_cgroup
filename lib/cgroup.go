package lib

import (
	"fmt"
	"io/ioutil"
	"strings"

	"path/filepath"
	"sync"
)

type Cgroup struct {
	name       string
	mountPoint string
	groups     []string
	sync.Mutex
}

var cgroups = map[string]*Cgroup{
	"cpu":     nil,
	"cpuacct": nil,
	"blkio":   nil,
}

func (c Cgroup) String() string {
	return c.mountPoint
}

func (c *Cgroup) visit(path string) {
	fileInfos, _ := ioutil.ReadDir(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			name := filepath.Join(path, fileInfo.Name())

			// Filter name to remove mount point prefix at the beginning
			filtered := strings.Replace(name, c.mountPoint, "", 1)

			c.groups = append(c.groups, filtered)
			c.visit(name)
		}
	}
}

func (c Cgroup) updateAvailableGroups() {
	// Mutex added since we use the same Cgroup object over time it can create nasty effects if called concurrently
	c.Lock()
	c.groups = make([]string, 0)
	c.visit(c.mountPoint)
	c.Unlock()

	fmt.Printf("%s - %#v \n", c.name, c.groups)
}

func findCgroupMountPoints() {
	// Create Map based on cgroup type and put them to respective places
	mountInfo, _ := ParseMountTable()
	for _, mount := range mountInfo {
		if mount.Fstype == "cgroup" {
			vfsOpts := strings.Split(mount.VfsOpts, ",")
			cgroupName := vfsOpts[1]

			// Check if it is in allowed set of cgroups
			_, ok := cgroups[cgroupName]
			if ok {
				cgroup := &Cgroup{name: cgroupName, mountPoint: mount.Mountpoint}
				cgroups[cgroupName] = cgroup
				cgroup.updateAvailableGroups()
			}
		}
	}

	// fmt.Printf("%s\n", cgroups)
}
