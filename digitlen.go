/* Wed Mar 19 11:08:08 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: none
Write a function DigitLen() that takes two integers as arguments and returns the times the first int can be divided by the second until it reaches zero.
The second int must be between 2 and 36. If not, the function returns -1.
If the first int is negative, reverse the sign and count the digits.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(digitlen(100, 10))
	fmt.Println(digitlen(100, 2))
	fmt.Println(digitlen(-100, 16))
	fmt.Println(digitlen(100, -1))
}

func digitlen(n, d int) int {
	if !(d >= 2 && d <= 36) { // if d <2 || d >36
		return -1
	}
	if n < 0 {
		n *= -1
	}
	count := 0
	for n != 0 {
		n /= d
		count++
	}
	return count
}
