/* Mon Mar 17 12:31:52 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: builtin
Given an array of non-negative integers representing the number of steps you can take forward from each position, implement the function CanJump() which takes a slice of unsigned integers []uint as input and returns a boolean value. This function should determine if it's possible to reach and stay at the last index of the array starting from the first index, based on the steps you need to advance. Be aware that:
Each value represents the exact number of steps you must take forward from that position.
The function should return true if it's possible to reach and stay at the last index without stepping out of the array, and false otherwise.
If the input has only one element, that is the last position in the array so the function will return true but if the array is empty it returns false.
Let's take the example array input := []uint{2, 3, 1, 1, 4}.
*/
package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(canjump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(canjump(input2))

	input3 := []uint{0}
	fmt.Println(canjump(input3))
}

func canjump(pos []uint) bool {
	for i := 0; i < len(pos); {
		if i == len(pos)-1 {
			return true
		} else if pos[i] == 0 {
			return false
		} else {
			i += int(pos[i])
		}
	}
	return false
}
