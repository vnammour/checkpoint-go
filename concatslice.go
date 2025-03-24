/* Wed Mar 12 05:14:31 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: builtin
Write a function ConcatSlice() that takes two slices of integers as arguments and
returns the concatenation of the two slices.
*/
package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	fmt.Println(ConcatSlice(x, y))
	fmt.Println("x:", x)
	x = []int{}
	y = []int{4, 5, 6, 7, 8, 9}
	fmt.Println(ConcatSlice(x, y))
	fmt.Println("x:", x)
	x = []int{1, 2, 3}
	y = []int{}
	fmt.Println(ConcatSlice(x, y))
	fmt.Println("x:", x)
}

/* Returns the concatenation of x & y into a NEW slice */
func ConcatSlice(x, y []int) []int {
	// t := make([]int,0,len(x)+len(y))
	// return append(append(t,x...),y...)
	return append(append(make([]int, 0, len(x)+len(y)), x...), y...)
}
