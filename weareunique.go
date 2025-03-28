/* Fri Mar 28 12:16:43 PM IDT 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 25
allowed: builtin, strings.Contains
Write a function that takes two strings's and returns the number of characters that are not included in both, without repeating characters.
If there is no unique characters return 0.
If both strings are empty return -1.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(unique("foo", "boo"))
	fmt.Println(unique("", ""))
	fmt.Println(unique("abc", "def"))
}

func unique(s1, s2 string) int {
    if len(s1) + len(s2) == 0 {
        return -1
    }
	m := make(map[rune]bool)
	l := len(s1)
	for _, r := range s1 {
		if _, ok := m[r]; !ok {
			m[r] = true
		}
	}
	for _, r := range s2 {
		if ok, _ := m[r]; ok {
			l--
		} else {
			l++
		}
	}
	return l
}
