package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// stdout

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// dump a string or just bytes into a file
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat.txt")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// granular write
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// obvious to defer a Close immediately after opening a file.
	defer f.Close()

	// `Write` byte slices.
	d2 := []byte{115, 111, 109, 101, 10} // some
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// write normal strings
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
