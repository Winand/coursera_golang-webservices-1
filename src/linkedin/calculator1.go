package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	val1 := readFloat("Input 1st value")
	val2 := readFloat("Input 2st value")
	// 88.6 + 76.3 = 164.89999999999998
	result := math.Round((val1+val2)*100) / 100
	fmt.Printf("Sum is %v\n", result)
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
