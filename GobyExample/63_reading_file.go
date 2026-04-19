package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := filepath.Join(os.TempDir(), "dat.txt")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// more control
	f, err := os.Open(path) // f _> os.File
	check(err)

	// beginning bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// seek location in the file and read from there
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// other methods of seeking are relative to the current cursor position.
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// and relative to the end of the file.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// no buil-in rewind
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// buffered reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}
