/* Wed Feb 26 07:42:48 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 32
allowed: github.com/01-edu/z01.PrintRune, unicode.IsGraphic, builtin
Write a function that takes (arr [10]byte), and displays the memory as in the example.
After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.
The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.
*/
package main

import (
	"github.com/01-edu/z01"
	"unicode"
)

func main() {
	arr := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	PrintMemory(arr)
}

func PrintMemory(arr [10]byte) {
	hex := hexvalue(arr[:len(arr)])
	for i := 0; i < len(hex); i++ {
		if hex[i] == 0 {
			z01.PrintRune('0')
		} else {
			z01.PrintRune(rune(hex[i]))
		}
		if i != 0 {
			if (i+1)%8 == 0 {
				z01.PrintRune('\n')
			} else if (i+1)%2 == 0 && i != len(hex)-1 {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
	str := string(arr[:len(arr)])
	for _, r := range str {
		if unicode.IsGraphic(r) {
			z01.PrintRune(r)
		} else {
			z01.PrintRune('.')
		}
	}
}

func hexvalue(arr []byte) []byte {
	base := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f'}
	hex := make([]byte, 2*len(arr))
	for i, j := 0, 1; i < len(arr); i++ {
		b := arr[i]
		k := j
		for b != 0 {
			hex[k] = base[b%16]
			b /= 16
			k--
		}
		j += 2
	}
	return hex
}
