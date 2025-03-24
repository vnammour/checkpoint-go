/* Thu Mar 20 08:25:07 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 4
level: 9
allowed: builtin
write a function that takes a string and a character as arguments and returns the number of times the character appears in the string.
if the character is not in the string return 0
if the string is empty return 0
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(countchar("Hello World", 'l'))
	fmt.Println(countchar("5  balloons", 5))
	fmt.Println(countchar("   ", ' '))
	fmt.Println(countchar("The 7 deadly sins", '7'))
}

func countchar(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
