// fourth version
// checks for network connection asynchronously

package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
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

func testConn(status chan error) {
	conn, err := net.Dial("tcp", "google.com:80")
	status <- err
	if err == nil {
		conn.Close()
	}
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

	status := make(chan error)
	go testConn(status)

	cmd.Stdout = &bufout
	cmd.Stderr = &buferr
	start := time.Now()
	cmd.Run()

	elapsed := time.Since(start)	
	timeTaken := fmt.Sprintf("\nTime needed to run: %v\n", elapsed.String())
	stdOut := fmt.Sprintf("Stdout from brew update: %v\n", bufout.String())
	stdErr := fmt.Sprintf("Stderr from brew update: %v\n", buferr.String())

	netStat := <-status
	if netStat != nil {
		errMsg := fmt.Sprintf("Network error encountered: %v\n", netStat.Error())
		log.Print(errMsg)
		logFile.WriteString(errMsg)
	}

	logFile.WriteString(start.String())
	logFile.WriteString(timeTaken)
	logFile.WriteString(stdOut)
	logFile.WriteString(stdErr)
	logFile.WriteString("\n")
}
