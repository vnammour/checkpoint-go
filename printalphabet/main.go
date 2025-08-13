/* Wed Jan 15 10:39:53 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for ch := 'a'; ch <= 'z'; ch++ {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
