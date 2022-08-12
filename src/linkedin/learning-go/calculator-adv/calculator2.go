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
	var result float64
	v1 := readFloat("Value 1")
	v2 := readFloat("Value 2")
	fmt.Print("Select an operation (+ - * /): ")
	switch op := readChar(); op {
	case "+":
		result = add(v1, v2)
	case "-":
		result = subtract(v1, v2)
	case "*":
		result = multiply(v1, v2)
	case "/":
		result = divide(v1, v2)
	default:
		panic(fmt.Sprintf("Unknown operation %v\n", op))
	}
	fmt.Printf("The result is %v\n", math.Round(result*100)/100)
}

func add(v1, v2 float64) float64 {
	// Adds two values
	return v1 + v2
}

func subtract(v1, v2 float64) float64 {
	// Subtracts one value from another
	return v1 - v2
}

func multiply(v1, v2 float64) float64 {
	// Multiplies two values
	return v1 * v2
}

func divide(v1, v2 float64) float64 {
	// Divides one value by another
	return v1 / v2
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

func readChar() string {
	// Read one character string
	var inp string
	_, _ = fmt.Scanln(&inp)
	inp = strings.TrimSpace(inp)
	if len(inp) != 1 {
		panic(fmt.Sprintf("Expected one char, got '%v'!", inp))
	}
	return inp
}
