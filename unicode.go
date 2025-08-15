/* Fri Aug 15 11:22:20 AM IDT 2025 */
/* By: Victor Nammour */
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "\u4e16\u754c"
	fmt.Println(s)
	s = "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(s)
	s = "\U00004e16\U0000754c"
	fmt.Println(s)

	// Unicode escapes may also be used in rune literals. These three literals are equivalent
	// '\u4e16' is same as '\U00004e16'

	// the string contains 13 bytes, but interpreted as UTF-8, it encodes only nine code points or runes:
	// (the hex bytes here are 3 bytes per rune for this asian charset)
	s = "hello, \xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(s)
	// a range loop decodes a utf-8-encoded string.
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("-----------------------")
	// to process those characters, we need a utf-8 decoder. The unicode/utf8 package provides one
	// that we can use like this:
	for i := 0; i < len(s); {
		// each call to decoderuneinstring returns r, the rune itself, and size, the number of
		// bytes occupied by the utf-8 encoding of r.
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// fortunately, Go's range loop, when applied to a string, performs utf-8 decoding implicitly.
	// The index jumps by more than 1 for each non-ASCII rune.
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("-----------------------")
	// we can count the number of runes in a string, like this:
	n := 0
	for range s {
		n++
	}
	// or we can just call utf8.RuneCountInString(s)
}

func hasprefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func hassuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// can be improved by using a hashing technique to search more efficiently
func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if hasprefix(s[i:],substr) {
			return true
		}
	}
	return false
}
