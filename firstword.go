/* Fri Mar 21 01:59:09 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: builtin
Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').
A word is a sequence of characters delimited by spaces or by the start/end of the argument.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstword("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(firstword(" lorem,ipsum "))
	fmt.Println(firstword(" "))
}

/* I consider \' as a letter to allow for words like "don't" */
func isletter(b byte) bool {
    return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b == '\''
}

func firstword(s string) string {
    i := 0
    for ; i < len(s) && !isletter(s[i]); i++ {
    }
    s = s[i:]
    for i = 0; i < len(s) && isletter(s[i]); i++ {
    }
    return s[:i]
}
