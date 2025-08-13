/* Wed Aug 13 08:43:56 AM IDT 2025 */
/* By: Victor Nammour */
package main

import (
	// "fmt"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 3 {
		return
	}
	hiddenp(os.Args[1], os.Args[2])
}

/*
If s1 is hidden in s2, the program should display 1 followed by a newline.
If s1 is not hidden in s2, the program should display 0 followed by a newline.
If s1 is an empty string, it is considered hidden in any string.
If the number of arguments is different from 2, the program should display nothing.
*/

// s1 = "ab"
// s2 = "xyabcde"
// Is s1 contained in s2 ?
func hiddenp(s1, s2 string) {
	if s1 == s2 {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}
	if len(s1) == 0 || len(s1) > len(s2) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	// The two blocks above are redundant, but a shortcut
	// i to index s1
	// j to index s2
	ok := true
	i := 0
	for j := 0; ok && i < len(s1) && j < len(s2); i,j = i+1,j+1 {
		for ; j < len(s2) && s1[i] != s2[j]; j++ {}
		if j >= len(s2) {
			ok = false
		}
	}

	ok = ok && i == len(s1)

	if !ok {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		// fmt.Println("Not found")
	} else {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		// fmt.Println("found")
	}
}
