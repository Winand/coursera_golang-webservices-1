/*
https://pkg.go.dev/fmt
*/
package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {
	x := 20
	f := 123.45
	// %d	base 10
	fmt.Printf("%d\n", x) // десятичное целое
	// %x	base 16, with lower-case letters for a-f
	fmt.Printf("%x\n", x) // 16-ричное число (hex)
	// %t	the word true or false
	fmt.Printf("%t\n", x > 10) // булево "true"/"false" (bool)
	// %f	decimal point but no exponent, e.g. 123.456
	fmt.Printf("%f\n", f) // число с плавающей запятой (float)
	// %e	scientific notation, e.g. -1.234456e+78
	fmt.Printf("%e\n", f) // экспоненциальное представление
	// https://pkg.go.dev/fmt#hdr-Explicit_argument_indexes
	fmt.Printf("%[2]d %[1]d\n", 52, 40) // обращение к аргументам по индексу
	// '#'	alternate format: add leading 0b for binary (%#b), 0 for octal (%#o),
	// 0x or 0X for hex (%#x or %#X); suppress 0x for %p (%#p);
	// for %q, print a raw (backquoted) string if strconv.CanBackquote
	// returns true;
	// always print a decimal point for %e, %E, %f, %F, %g and %G;
	// do not remove trailing zeros for %g and %G;
	// write e.g. U+0078 'x' if the character is printable for %U (%#U).
	fmt.Printf("%d %#[1]o %#[1]x\n", 52) // печать одного числа в 10-, 8- и 16ричном представлении

	c := circle{
		radius: 20,
		border: 5,
	}
	// %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	// %T	a Go-syntax representation of the type of the value
	fmt.Printf("%T\n", c)

	//fmt.Sprintf - вывод в строковую переменную
	s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
	fmt.Println(s)
}
