// third version
// runs a timer and writes results to a log file

package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

const LOG_FILE_NAME = ".brew_update_log"

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
	var bufout bytes.Buffer
	var buferr bytes.Buffer

	home := os.Getenv("HOME")
	logPath := path.Join(home, LOG_FILE_NAME)
	logFile, err := AppendFile(logPath)

	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	cmd.Stdout = &bufout
	cmd.Stderr = &buferr
	start := time.Now()
	cmd.Run()

	elapsed := time.Since(start)
	logFile.WriteString(start.String())
	logFile.WriteString("\nTime needed to run: ")
	logFile.WriteString(elapsed.String())
	logFile.WriteString("\nStdout from brew update: ")
	logFile.WriteString(bufout.String())
	logFile.WriteString("\nStderr from brew update: ")
	logFile.WriteString(buferr.String())
	logFile.WriteString("\n\n")
}
