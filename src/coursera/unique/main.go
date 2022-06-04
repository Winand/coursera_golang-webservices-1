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
	var prev string
	for in.Scan() { //построчный ввод
		txt := in.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			panic("File not sorted: txt==" + txt + ", prev==" + prev)
		}

		prev = txt

		fmt.Println(">>", txt)
	}
}
