/* Fri Feb 21 11:45:41 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 32
allowed: builtin, strconv.Itoa
Write a function that takes a string and returns a new string that replaces every character with the number
of duplicates and the character itself, deleting the extra duplications.
NOTE: The letters are from the Latin alphabet list only. Any other character, symbols, shall not be tested.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	zipped := zip("YouuungFellllas")
	fmt.Println(zipped)
	fmt.Println(unzip(zipped))
	fmt.Println(zip("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(zip("Helloo Therre!"))
}

/*
Note: This only works for ASCII.
Input: a string.
Output: a count of each character in its ordinal position in the string.
*/
func zip(s string) string {
	count := 1
	out := make([]byte, 0, len(s)+1)
	for i, j := 0, 0; i < len(s); i++ {
		for j = i; j < len(s)-1 && s[j] == s[j+1]; j++ {
			count++
		}
		if count > 1 {
			out = append(out, []byte(strconv.Itoa(count))...)
		} else {
			out = append(out, '1')
		}
		out = append(out, s[i])
		i = j
		count = 1
	}
	return string(out)
}

/* Assumption: digits aren't literal, just counts of the proceeding letters */
func unzip(s string) string {
	out := make([]byte, 0, len(s))
	for i, j := 0, 0; i < len(s); i++ {
		for j = i; j < len(s) && isdigit(s[j]); j++ {
		}
		if i != j { // saw a number
			num, err := strconv.Atoi(s[i:j])
			if err != nil {
				return ""
			}
			if j != len(s) {
				for k := 0; k < num; k++ {
					out = append(out, s[j])
				}
			}
			i = j
		} else {
			out = append(out, s[i])
		}
	}
	return string(out)
}

func isalpha(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isdigit(b byte) bool {
	return b >= '0' && b <= '9'
}
