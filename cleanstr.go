/* Mon Mar 17 01:35:01 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 32
allowed: os.*, github.com/01-edu/z01.PrintRune, builtin
Write a program that takes a string, and displays this string with exactly:
one space between words.
without spaces nor tabs at the beginning nor at the end.
with the result followed by a newline ("\n").
A "word" is defined as a part of a string delimited either by spaces/tabs, or by the start/end of the string.
If the number of arguments is not 1, or if there are no words to display, the program displays a newline("\n").
*/
package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	cleanstr(os.Args[1])
}

func isspace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\v' || b == '\n' || b == '\r'
}

func cleanstr(s string) {
	if len(s) == 0 {
		z01.PrintRune('\n')
		return
	}
	i := 0
	// trim leading and trailing spaces
	for ; i < len(s) && isspace(s[i]); i++ {
	}
	s = s[i:]
	for i = len(s) - 1; i >= 0 && isspace(s[i]); i-- {
	}
	s = s[:i+1] // i+1, to account for last non-space byte

	// process rest of string
	spaced := false
	for i = 0; i < len(s); i++ {
		if isspace(s[i]) {
			if !spaced {
				z01.PrintRune(' ')
				spaced = true
			}
		} else {
			z01.PrintRune(rune(s[i]))
			spaced = false
		}
	}
	z01.PrintRune('\n')
}
