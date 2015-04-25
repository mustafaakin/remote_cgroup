package lib

import (
    "fmt"
    "io/ioutil"
    "syscall"
    "strings"
)

func Start(){
    fmt.Println("Hello from Util")
    scanDevFiles()
    findCgroupMountPoints()
}

func findCgroupMountPoints(){
    // Create Map based on cgroup type and put them to respective places
    mountInfo, _ := ParseMountTable()
    for _, mount := range mountInfo {
        if mount.Fstype == "cgroup" {
            vfsOpts := strings.Split(mount.VfsOpts, ",")
            cgroup := vfsOpts[1]

            fmt.Printf("%s %s\n", cgroup, mount.Mountpoint)
        }
    }
}

func scanDevFiles(){
    files, _ := ioutil.ReadDir("/dev")

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