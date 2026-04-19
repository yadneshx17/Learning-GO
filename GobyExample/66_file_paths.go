package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// `Join` takes any number of arguments and constructs a hierarchial path from that.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// normalize paths by removing superfluous separators and directory changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// split a path to the directory and the file.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// split extension out of such names with `Ext`
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// files name  with extension removed
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}
