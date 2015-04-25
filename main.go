package main

import (
    "runtime"
    "fmt"
    "github.com/mustafaakin/remote_cgroup/lib"
    "net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Home")
}

func main() {
    fmt.Println("Hello World From main module")
    runtime.GOMAXPROCS(runtime.NumCPU())
    lib.Start()

    // TODO: Make it more beautiful
    /*
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    http.Handle("/", r)
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatalf("Could not listen on port 8080!")

    }
    */
}
