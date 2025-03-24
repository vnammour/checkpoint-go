/* Thu Mar 20 11:59:27 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: builtin
Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.
'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(repeatalpha("abc"))
	fmt.Println(repeatalpha("Choumi."))
	fmt.Println(repeatalpha(""))
	fmt.Println(repeatalpha("abacadaba 01!"))
}

func repeatalpha(s string) string {
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		ref := reference(s[i])
		if ref != 0 {
			count := int(s[i] - ref)
			for count >= 0 {
				out = append(out, s[i])
				count--
			}
		} else {
			out = append(out, s[i])
		}
	}
	return string(out)
}

func reference(b byte) byte {
	if islower(b) {
		return 'a'
	} else if isupper(b) {
		return 'A'
	}
	return 0
}
func islower(b byte) bool {
	return b >= 'a' && b <= 'z'
}
func isupper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
