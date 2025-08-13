/* Thu Mar 13 12:55:31 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: os.Args, github.com/01-edu/z01.PrintRune, builtin
Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string. If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing. If the number of arguments is different from 2, the program displays nothing.
*/

package main

import (
    "fmt"
    "os"
    // "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 3 { return }
	s := wdmatch(os.Args[1], os.Args[2])
	if len(s) > 0 {
    	fmt.Println(s)
	}
}

func wdmatch(s1,s2 string) string {
	if s1 == s2 {
		return s1 // success
	}
	if len(s1) > len(s2) {
		return "" // failure
	}

	// i to index s1
	// j to index s2
	ok := true
	i := 0
	for j:= 0; ok && i < len(s1) && j < len(s2); i,j = i+1, j+1 {
		for ; j < len(s2) && s1[i] != s2[j]; j++ {}
		if j >= len(s2) {
			ok = false
		}
	}

	ok = ok && i == len(s1)

	if ok {
		return s1
	} else {
		return ""
	}
}
