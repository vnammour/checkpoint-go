/* Thu Jan 16 11:01:35 AM IST 2025 */
/* By: Victor Nammour */
package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
