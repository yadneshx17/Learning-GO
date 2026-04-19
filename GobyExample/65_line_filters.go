package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A line filter is a common type of program that reads input on stdin, processes it, and then prints some derived result to stdout.

func main() {

	// wrapper unbuffered `os.Stdin` into buffered scanner
	// it advances the scanner to the next token; which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	// Text returns the current token, here the next line, from the input.
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
