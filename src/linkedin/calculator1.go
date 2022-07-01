package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	val1 := readFloat("Input 1st value")
	val2 := readFloat("Input 2st value")
	fmt.Printf("Sum is %v\n", val1+val2)
}

func readFloat(prompt string) float64 {
	var fval float64
	fmt.Print(prompt + ": ")
	_, err := fmt.Scanln(&fval)
	if err != nil {
		panic("Wrong input!")
	}
	return fval
}

func readFloatBufio(prompt string) float64 {
	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	fval, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	if err != nil {
		panic("Wrong input!")
	}
	return fval
}
