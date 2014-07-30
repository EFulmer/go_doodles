// third version
// runs a timer and writes results to a log file

// TODO do with reg file not log file, write header each time
package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

func AppendFile(name string) (*os.File, error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return os.Create(name)
		}
	}
	return os.OpenFile(name, os.O_APPEND | os.O_RDWR, 0660)
}

func main() {
	cmd := exec.Command("brew", "update")
	var cmdOut bytes.Buffer
	var cmdErr bytes.Buffer

	home := os.Getenv("HOME")
	logPath := path.Join(home, "brew_cron_log.txt")
	logFile, err := AppendFile(logPath)

	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "cron", log.LstdFlags)
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr
	start := time.Now()
	cmd.Run()

	elapsed := time.Since(start)
	logger.Printf("Time needed to run: %v\n", elapsed)
	logger.Printf("Stdout from brew update: %q\n", cmdOut.String())
	logger.Printf("Stderr from brew update: %q\n", cmdErr.String())
}
