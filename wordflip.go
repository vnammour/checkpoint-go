/* Wed Aug 13 06:05:49 PM IDT 2025 */
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
	s := "hello      world,     again     "
	fmt.Print(wordflip(s))
	fmt.Print(wordflip("First second last"))
	fmt.Print(wordflip(""))
	fmt.Print(wordflip("     "))
	fmt.Print(wordflip(" hello  all  of  you! "))
}
func wordflip(s string) string {
	if len(s) == 0 {
		return "Invalid Output\n"
	}
	split := strings.Split(s," ")
	flipped := make([]byte,0,len(split)*2)
	for i:= len(split)-1; i >= 0; i-- {
		s = strings.TrimSpace(split[i])
		if len(s) != 0 {
			flipped = append(flipped, s[:]...)
			flipped = append(flipped, ' ')
		}
	}
	if len(flipped) != 0 {
		flipped = flipped[:len(flipped)-1] // remove last space
	}
	return string(append(flipped, '\n'))
}
