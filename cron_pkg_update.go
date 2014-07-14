// This program is a wrapper around brew update that can be run as a cronjob.

package main

import (
    "fmt"
    "os/exec"
)

func main() {
    update := exec.Command("brew", "update")
    out, err := update.Output()
    if err != nil {
        fmt.Print(err)
    } 
    fmt.Printf("%s", out)
}
