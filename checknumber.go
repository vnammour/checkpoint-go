/* Thu Mar 20 08:17:59 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 4
level: 9
allowed: none
Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(checknumber("Hello"))
	fmt.Println(checknumber("Hello1"))
}

func checknumber(s string) bool {
	i := 0
	for ; i < len(s) && !isnumber(s[i]); i++ {
	}
	return i < len(s)
}

func isnumber(b byte) bool {
	return b >= '0' && b <= '9'
}
