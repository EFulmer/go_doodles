// This program is a wrapper around brew update that can be run as a cronjob.

package main

import (
        "fmt"
        "os"
        "os/exec"
)

func main() {
        update := exec.Command("brew", "update")
        _, err := update.Output()
        
        // TODO replace with cleanly writing to stdout using the Cmd struct's
        // fields
        if err != nil {
                fmt.Println(os.Stderr, err)
        } 
}
