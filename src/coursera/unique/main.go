package main

import (
	"bufio" //буферизованный ввод/вывод
	"fmt"
	"os" //стандартный ввод
)

func main() {
	// echo "hello" | go run main.go
	// cat text.txt | go run main.go
	in := bufio.NewScanner(os.Stdin)
	alreadySeen := make(map[string]bool)
	for in.Scan() { //построчный ввод
		txt := in.Text()
		if _, found := alreadySeen[txt]; found {
			continue
		}

		alreadySeen[txt] = true

		fmt.Println(">>", txt)
	}
}
