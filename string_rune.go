package main

import (
	"fmt"
	"unicode/utf8"
	// "unicode/utf8"
)

// String is just a read-only slice of bytes.
// Rune ( int32 ) a huge number that represent the Global ID ( unicode  ) of a character.
// these numbers are called runes. 'A' (single quotes)
// 'A' = 65
// '⌘' = 8984

// character's equivalent in Go rune, an integer represents a unicode code point
// https://go.dev/blog/strings
// Go strings literals are UTF-8 encoded text
// https://news.ycombinator.com/item?id=20325612#:~:text=a%20fe...-,I%20can%20only%20speak%20from%20experience%2C%20but%20Thai%20language%20uses,chance%20of%20exceeding%20510%20bytes.
func main() {
	const s = "สวัสดี"

	const d = "안녕하세요"
	var a = "ab"
	const b = "ab"

	// NOTE: Count how many letters in to the string.
	fmt.Println(len([]rune("a⌘"))) // 2
	fmt.Println("⌘")

	// `len()` counts the boxes(bytes)
	fmt.Println("Len: ", len(s))
	fmt.Println("Len of a:", len(a))
	fmt.Println("Len of b:", len(b))
	fmt.Println("Len of d:", len(d))

	for i := 0; i < len(s); i++ {
		// prints out the raw hex codes
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Run-time of `RuneCountInString` depends on the size of the string, because iot has to decode each UTF-8 rune sequenetially.
	fmt.Println("Rune Count: ", utf8.RuneCountInString(s))

	// When is `range` used, the bytes gets glued back to rune.
	// it handles string specially and decodes each `rune` along side the offset.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts -- %c -- at %d \n", runeValue, runeValue, idx)
	}

	// explicit working of `range` like.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)

	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("Found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
