/* Tue Aug 12 04:55:16 PM IDT 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
group: 5
level: 9
Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').
If there are no argument, the program displays nothing.
allowedFunctions: [ "github.com/01-edu/z01.PrintRune", "os.*", "--allow-builtin" ]
*/
package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	for _,s := range(os.Args[1:]) {
		s = reversestrcap(s)
		/*if len(s) == 0 {
			return
		}*/
		for _,c := range(s) {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		// fmt.Println(reversestrcap(s))
	}
}

/*
Assumption: Ascii characters only
Get all non-letters, then while no space is encountered, consume the chars.
Once a space is encountered (or end of string), check last non-space char: if letter, capitalize it.
*/

func reversestrcap(s string) string {
	out := make([]byte,len(s))
	for i := 0; i < len(s); {
		// consume non-letter
		if !isLetter(s[i]) {
			out[i] = s[i]
			i++
		} else { // letter
			j := i
			for ; j < len(s) && !isSpace(s[j]); j++ {
				if isUpper(s[j]) {
					out[j] = toLower(s[j])
				} else {
					out[j] = s[j]
				}
			}
			c := out[j-1]
			if isLetter(c) {
				out[j-1] = toUpper(c)
			}
			// consume space char if encountered
			if j < len(s) {
				out[j] = s[j]
				j++
			}
			i = j
		}
	}
	return string(out)
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return isUpper(c) || isLower(c)
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func toUpper(c byte) byte {
	return c - 'a' + 'A'
}

func toLower(c byte) byte {
	return c - 'A' + 'a'
}
