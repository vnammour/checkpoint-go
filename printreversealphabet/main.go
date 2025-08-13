/* Thu Jan 16 08:29:45 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for ch := 'z'; ch >= 'a'; ch-- {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
