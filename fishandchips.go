/* Fri Mar 14 12:29:45 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: none
Write a function called FishAndChips() that takes an int and returns a string.
If the number is divisible by 2, print fish.
If the number is divisible by 3, print chips.
If the number is divisible by 2 and 3, print fish and chips.
If the number is negative return error: number is negative.
If the number is non divisible by 2 or 3 return error: non divisible.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f(4))
	fmt.Println(f(9))
	fmt.Println(f(6))
}

func f(n int) string {
	switch {
	case n < 0:
		return "number is negative"
	case n%2 == 0 && n%3 == 0:
		return "fish and chips"
	case n%2 == 0:
		return "fish"
	case n%3 == 0:
		return "chips"
	default:
		return "non divisible"
	}
}
