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
	var op string
	var result float64
	v1 := readFloat("Value 1")
	v2 := readFloat("Value 2")
	fmt.Print("Select an operation (+ - * /): ")
	_, _ = fmt.Scanln(&op)
	switch op {
	case "+":
		result = v1 + v2
	case "-":
		result = v1 - v2
	case "*":
		result = v1 * v2
	case "/":
		result = v1 / v2
	default:
		panic(fmt.Sprintf("Unknown operation %v\n", op))
	}
	fmt.Printf("The result is %v\n", math.Round(result*100)/100)
}

func readFloat(prompt string) float64 {
	fmt.Print(prompt + ": ")
	rd := bufio.NewReader(os.Stdin)
	str, _ := rd.ReadString('\n')
	str = strings.TrimSpace(str)
	ret, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(fmt.Sprintf("Wrong input %v!", str))
	}
	return ret
}
