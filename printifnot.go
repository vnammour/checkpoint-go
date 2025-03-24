/* Fri Mar 21 07:22:54 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 4
level: 9
allowed: builtin
Write a function that takes a string as an argument and returns the letter G if the argument length is less than 3, otherwise returns Invalid Input followed by a newline \n.
If it's an empty string return G followed by a newline \n.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Print(printifnot("abcdefz"))
	fmt.Print(printifnot("abc"))
	fmt.Print(printifnot(""))
	fmt.Print(printifnot("14"))
}

func printifnot(s string) string {
	if len(s) == 0 {
		return "G\n"
	}
	if len(s) < 3 {
		return "G"
	}
	return "Invalid Input\n"
}
