package lib

type Cgroup struct {
    name string
    mountPoint string
}

func(c Cgroup) String() string {
    return c.mountPoint
}