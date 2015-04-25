package lib

import (
    "strings"
    "fmt"
)
// enum Cgroups?

type Cgroup struct {
    name string
    mountPoint string
}

var cgroups = map[string]*Cgroup {
    "cpu": nil,
    "cpuacct": nil,
    "blkio": nil,
}

func(c Cgroup) String() string {
    return c.mountPoint
}

func(c Cgroup) getAvailableGroups() {
    // ioutil.ReadDir(mountpoint)
    // path.join
    return
}

func findCgroupMountPoints(){
    // manually write them

    // Create Map based on cgroup type and put them to respective places
    mountInfo, _ := ParseMountTable()
    for _, mount := range mountInfo {
        if mount.Fstype == "cgroup" {
            vfsOpts := strings.Split(mount.VfsOpts, ",")
            cgroup := vfsOpts[1]

            // Check if it is in allowed set of cgroups
            _, ok := cgroups[cgroup]
            if ok {
                cgroups[cgroup] = &Cgroup{name: cgroup, mountPoint: mount.Mountpoint}
            }
        }
    }

    fmt.Printf("%s\n", cgroups)
}