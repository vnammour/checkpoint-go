/* Thu Jan 16 11:47:05 AM IST 2025 */
/* By: Victor Nammour */
package piscine

// n!/k!(n-k)!, n = 10, k = 3, I have 10 choose 3, which is 120

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if i == '7' && j == '8' && k == '9' {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
				} else {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
