/* Thu Mar 20 08:32:48 AM IST 2025 */
/* By: Victor Nammour */
/* Write a function that returns the first rune of a string. */
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(firstrune("Hello"))
	z01.PrintRune(firstrune("Salut!"))
	z01.PrintRune(firstrune("Ola!"))
	z01.PrintRune('\n')
	z01.PrintRune(firstrune(""))
}

func firstrune(s string) rune {
	var r rune
	for i, c := range s {
		r = c
		if i == 0 {
			break
		}
	}
	return r
}
