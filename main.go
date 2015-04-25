package main

import (
    "runtime"
    "fmt"
    "github.com/mustafaakin/remote_cgroup/lib"
)

func main() {
    fmt.Println("Hello World From main module")
    runtime.GOMAXPROCS(runtime.NumCPU())
    lib.Start()
}
