/* Wed Mar 12 08:00:43 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: builtin
Write a function called SaveAndMiss() that takes a string and an int as an argument. The function should move through the string in sets determined by the int, saving the first set, omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the end of the string is reached. Return a string containing the saved characters.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(saveandmiss("abcdef", 1))
	fmt.Println(f("abcdef", 1))
	fmt.Println(saveandmiss("123456789", 1))
	fmt.Println(f("123456789", 1))
	fmt.Println(saveandmiss("123456789", 3))
	fmt.Println(f("123456789", 3))
	fmt.Println(saveandmiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(f("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(saveandmiss("", 3))
	fmt.Println(saveandmiss("hello you all ! ", 0))
	fmt.Println(saveandmiss("what is your name?", 0))
	fmt.Println(saveandmiss("go Exercise Save and Miss", -5))

	fmt.Println(f("", 3))
	fmt.Println(f("hello you all ! ", 0))
	fmt.Println(f("what is your name?", 0))
	fmt.Println(f("go Exercise Save and Miss", -5))
}

func saveandmiss(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	t := []rune(s)
	out := make([]rune, 0, len(t)/2+n)
	i := 0
	for i+n < len(t) {
		out = append(out, t[i:i+n]...)
		i += 2 * n
	}
	if i < len(t) {
		out = append(out, t[i:]...)
	}
	return string(out)
}

// naiive way
func f(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	t := []rune(s)
	out := make([]rune, 0, len(t)/2+n)
	for i := 0; i < len(t); i += 2 * n {
		// abcdefghi
		for j := i; j < i+n && j < len(t); j++ {
			out = append(out, t[j])
		}
	}
	return string(out)
}
