/* Wed Mar 19 10:54:15 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 9
allowed: builtin, strings.String
Write a function ThirdTimeIsACharm() that takes a string as an argument and returns another string with every third character.
Return the output followed by a newline \n.
If the string is empty, return a newline \n.
If there is no third character, return a newline \n.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(thirdtime("123456789"))
	fmt.Print(thirdtime(""))
	fmt.Print(thirdtime("a b c d e f"))
	fmt.Print(thirdtime("12"))
}

func thirdtime(s string) string {
	if len(s) == 0 || len(s) < 3 {
		return "\n"
	}
	buf := strings.Builder{}
	for i := 2; i < len(s); {
		buf.WriteByte(s[i])
		i += 3
	}
	buf.WriteByte('\n')
	return buf.String()
}
