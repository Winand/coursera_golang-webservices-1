package main

import (
	"bufio" //буферизованный ввод/вывод
	"fmt"
	"io"
	"os" //стандартный ввод
)

//Логика вынесена в отдельную функцию для упрощения написания тестов
func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() { //построчный ввод
		txt := in.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			return fmt.Errorf("file not sorted: txt==%v, prev==%v", txt, prev)
		}

		prev = txt

		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	// echo "hello" | go run main.go
	// cat text.txt | go run main.go
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
