package lib

import (
    "fmt"
    "io/ioutil"
    "syscall"
    "strings"
)

var cgroups = map[string]*Cgroup {
    "cpu": nil,
    "cpuacct": nil,
    "blkio": nil,
}

var devices = map[string]*Device{}


func Start(){
    fmt.Println("Hello from Util")

    scanDevFiles()
    findCgroupMountPoints()
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

func scanDevFiles(){
    files, _ := ioutil.ReadDir("/dev")

    for _, file := range files {
        Sys := file.Sys()

        if  Sys != nil  {
            stat := Sys.(*syscall.Stat_t)
            Rdev := stat.Rdev

            // get first 8 bits only
            minor := 0xFF & Rdev

            // subtracting minor just in case then shiftit it 8 bytes, according to spec it can be 12 bits at most but no check is made
            major := (Rdev - minor) >> 8

            dev := &Device{name: file.Name(), minor: int(minor), major: int(major)}
            devices[file.Name()] = dev
        }
    }

    fmt.Printf("%s\n", devices)
}