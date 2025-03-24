/* Thu Feb 13 09:17:26 PM IST 2025 */
/* By: Victor Nammour */
/*
Write a function rot14 that returns the string within the parameter transformed into a rot14 string. Each letter will be replaced by the letter 14 spots ahead in the alphabetical order.
'z' becomes 'n' and 'Z' becomes 'N'. The case of the letter stays the same.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(rot14("hello World!"))
	fmt.Println(rot14("Hello! How are You?"))
}

func rot14(s string) string {
	str := []rune(s)
	for i, v := range str {
		switch {
		case v >= 'A' && v <= 'Z':
			str[i] = (v-'A'+14)%26 + 'A'
		case v >= 'a' && v <= 'z':
			str[i] = (v-'a'+14)%26 + 'a'
		}
	}
	return string(str)
}
