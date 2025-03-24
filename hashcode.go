/* Fri Mar 14 12:44:15 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: builtin
Write a function called hashcode() that takes a string as an argument and returns a new hashed string.
The hash equation is computed as follows:
(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.
If the resulting character is unprintable add 33 to it.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(hashcode("A"))
	fmt.Println(hashcode("AB"))
	fmt.Println(hashcode("BAC"))
	fmt.Println(hashcode("Hello World"))
}

const (
	NUM    = 127
	OFFSET = 33
)

func hashcode(s string) string {
	hstr := make([]byte, len(s))
	var temp int
	for i := 0; i < len(s); i++ {
		temp = (int(s[i]) + len(s)) % NUM
		if temp < OFFSET { // if unprintable
			temp += OFFSET
		}
		hstr[i] = byte(temp)
	}
	return string(hstr)
}
