/* Fri Mar 21 07:13:47 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 4
level: 9
allowed: builtin, strings.String
Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.
If the length of the string is odd, round it down.
If the string is empty, return an empty string.
If the string length is equal to one, return the string.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(retainfirsthalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(retainfirsthalf("A"))
	fmt.Println(retainfirsthalf(""))
	fmt.Println(retainfirsthalf("Hello World"))
}

func retainfirsthalf(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	return s[:len(s)/2]
}
