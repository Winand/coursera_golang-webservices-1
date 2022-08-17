package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	// https://gobyexample.com/reading-files
	content, err := os.ReadFile(pidFile)
	if err != nil {
		return err
	}
	str := strings.TrimSpace(string(content))
	// Max PID is 2^22 https://stackoverflow.com/q/6294133/maximum-pid-in-linux
	// Max uint32 is 4`294`967`295 https://stackoverflow.com/a/6878625/1119602
	pid, err := strconv.ParseUint(str, 0, 32)
	if err != nil {
		return errors.Wrap(err, "Invalid pid")
	}
	fmt.Printf("killing server with pid=%d\n", pid)
	if err := os.Remove(pidFile); err != nil {
		log.Printf("Warning: Failed to remove file %s\n", pidFile)
	}
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		log.Fatalf("error: %s", err) // Printf + os.Exit(1)
	}
}
