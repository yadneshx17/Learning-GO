package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Defer is used to ensure that a function call is performed in a programs's execution, usually for purpose of cleaning.

func main() {
	// fmt.Println(os.TempDir()) /tmp
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createfile(path)
	defer closefile(f) // executes at the end of the enclosing function(main), after writeFile has finish
	writefile(f)
}

func createfile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writefile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closefile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
}
