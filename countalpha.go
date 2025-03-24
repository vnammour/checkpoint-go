/* Thu Mar 20 09:03:14 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 4
level: 9
allowed: none
Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(countalpha("Hello world"))
	fmt.Println(countalpha("H e l l o"))
	fmt.Println(countalpha("H1e2l3l4o"))
}

func countalpha(s string) int {
	count := len(s)
	for i := 0; i < len(s); i++ {
		if !isalpha(s[i]) {
			count--
		}
	}
    return count
}

func isalpha(b byte) bool {
	return b >= 'A' && b <= 'Z' ||
		b >= 'a' && b <= 'z'
}
