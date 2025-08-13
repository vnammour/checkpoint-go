/* Thu Jan 23 05:50:25 PM IST 2025 */
/* By: Victor Nammour */
package piscine

import (
	"github.com/01-edu/z01"
)

func Printrevcomb() {
  count := 0
  for i := '9'; i >'2'; i-- {
    for j := i-1; j >='1'; j-- {
      for k := j-1; k >='0'; k-- {
        z01.PrintRune(i)
        z01.PrintRune(j)
        z01.PrintRune(k)
        z01.PrintRune(',')
        z01.PrintRune(' ')
        count++
        if count == 20 {
          //return
        }
      }
    }
  }
  z01.PrintRune('2')
  z01.PrintRune('1')
  z01.PrintRune('0')
  z01.PrintRune('\n')
}
