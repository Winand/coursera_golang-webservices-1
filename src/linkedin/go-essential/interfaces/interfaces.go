package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	output io.Writer
}

func (c Capper) Write(p []byte) (n int, err error) {
	return c.output.Write(
		[]byte(
			strings.ToUpper(string(p)),
		),
	)
}

func main() {
	//   ⬐ Why?
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}