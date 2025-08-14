/* Mon Mar 17 04:31:05 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 14
level: 32
allowed: builtin, strings.(Split,TrimSpace)
Write a function WordFlip() that takes a string as input and returns it in reverse order.
The output should be followed by a newline \n.
If the string is empty, return Invalid Output.
Ignore multiple spaces between words and trim any leading or trailing spaces in the string.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(wordflip("First second last"))
	fmt.Print(wordflip(""))
	fmt.Print(wordflip("       "))
	fmt.Print(wordflip(" hello   all   of   you!   "))
}

/*
	If I used a []string instead of []byte for result, I'll have to loop at end

on the []string to return a string.
*/
func wordflip(s string) string {
	if len(s) == 0 {
		return "Invalid Output.\n"
	}
	result := make([]byte, 0, len(s))
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) != 0 {
			result = append(result, []byte(strings.TrimSpace(words[i]))...)
			result = append(result, ' ')
		}
	}
	if len(result) != 0 {
		result = result[:len(result)-1] // remove last space
	}
	result = append(result, '\n')
	return string(result)
}
