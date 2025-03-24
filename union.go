/* Wed Mar 12 05:50:53 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: os.Args, github.com/01-edu/z01.PrintRune, builtin
Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.
The display will be in the same order that the characters appear on the command line and will be followed by a newline ('\n').
If the number of arguments is different from 2, then the program displays a newline ('\n').
*/
/* Sample run:
$ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
zpadintoqefwjy$
$
$ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
$
$ go run . rien "cette phrase ne cache rien" | cat -e
rienct phas$
$
$ go run . | cat -e
$
$ go run . rien | cat -e
$
*/
package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
	} else {
		union(os.Args[1], os.Args[2])
	}
}

func union(s1, s2 string) {
	m := make(map[rune]bool, max(len(s1), len(s2)))
	for _, r := range s1 {
		if seen := m[r]; !seen {
			m[r] = true
			z01.PrintRune(r)
		}
	}

	for _, r := range s2 {
		if seen := m[r]; !seen {
			m[r] = true
			z01.PrintRune(r)
		}
	}

	z01.PrintRune('\n')

}
