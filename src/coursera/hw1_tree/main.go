package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	// "strings"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	return listDir(out, path, "", printFiles)
}

func listDir(out io.Writer, path string, prefix string, printFiles bool) error {
	var node = "├───"
	var lastNode = "└───"
	var pass = "│	"
	var lastPass = "	"

	items, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read dir %v", path)
	}
	//------- filter out files ---
	if !printFiles {
		var filteredItems []fs.DirEntry
		for _, item := range items {
			if item.IsDir() {
				filteredItems = append(filteredItems, item)
			}
		}
		items = filteredItems
	}
	//----------------------------
	cnt := len(items)
	for idx, item := range items {
		isLast := idx == cnt-1
		// https://stackoverflow.com/q/45941821
		fullPath := filepath.Join(path, item.Name())
		var newPrefix string
		if item.IsDir() {
			if isLast {
				fmt.Fprintf(out, "%v%v%v\n", prefix, lastNode, item.Name())
				newPrefix = prefix + lastPass
			} else {
				fmt.Fprintf(out, "%v%v%v\n", prefix, node, item.Name())
				newPrefix = prefix + pass
			}
			listDir(out, fullPath, newPrefix, printFiles)
		} else { // file
			f, err := os.Open(fullPath)
			if err != nil {
				return fmt.Errorf("failed to open file %v", fullPath)
			}
			finfo, _ := f.Stat()
			if isLast {
				newPrefix = prefix + lastNode
			} else {
				newPrefix = prefix + node
			}
			if sz := finfo.Size(); sz > 0 {
				fmt.Fprintf(out, "%v%v (%vb)\n", newPrefix, item.Name(), finfo.Size())
			} else {
				fmt.Fprintf(out, "%v%v (empty)\n", newPrefix, item.Name())
			}
		}
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
