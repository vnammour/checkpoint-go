/* Fri Mar 21 01:49:13 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: builtin
Write a function LastWord that takes a stringand returns its last word followed by a\n`.
A word is a section of string delimited by spaces or by the start/end of the string.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastword("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(lastword(" lorem,ipsum "))
	fmt.Println(lastword(" "))
}

/* I consider \' as a letter to allow for words like "don't" */
func isletter(b byte) bool {
    return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b == '\''
}

func lastword(s string) string {
	i := len(s) - 1
	for ; i >= 0 && !isletter(s[i]); i-- {
	}
	s = s[:i+1]
	for i = len(s) - 1; i >= 0 && isletter(s[i]); i-- {
	}
	return s[i+1:]
}
