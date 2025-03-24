/* Thu Mar 20 03:38:23 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 9
Write a function that simulates the behavior of the Itoa function in Go. Itoa transforms a number represented as anint in a number represented as a string.
For this exercise the handling of the signs + or - does have to be taken into account.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(itoa(-123))
	fmt.Println(itoa(12345))
	fmt.Println(itoa(0))
	fmt.Println(itoa(-1234))
	fmt.Println(itoa(987654321))
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := n
	if sign < 0 {
		n *= -1
	}
	out := make([]byte, 0, 10)
	for n != 0 {
		out = append(out, byte(n%10)+'0') // get next digit
		n /= 10                           // delete it
	}
	if sign < 0 {
		out = append(out, '-')
	}
	return string(reverse(out))
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
