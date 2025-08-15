/* Thu Mar 20 12:17:30 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 9
*/
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	printrevcomb()
}

func printrevcomb() {
	two, one, zero := '2', '1', '0'
	for i := '9'; i >= two; i-- {
		for j := i - 1; j >= one; j-- {
			for k := j - 1; k >= zero; k-- {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i != two || j != one || k != zero {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
