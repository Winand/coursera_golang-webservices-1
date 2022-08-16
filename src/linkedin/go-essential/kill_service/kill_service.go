package main

import (
	"fmt"
	"log"
)

func killServer(pidFile string) error {
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("sever.pid"); err != nil {
		log.Fatalf("error: %s", err)
	}
}
