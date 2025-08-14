/* Thu Feb 13 09:17:26 PM IST 2025 */
/* By: Victor Nammour */
/*
Write a function rot14 that returns the string within the parameter transformed into a rot14 string. Each letter will be replaced by the letter 14 spots ahead in the alphabetical order.
'z' becomes 'n' and 'Z' becomes 'N'. The case of the letter stays the same.

Assume the set of lower ascii is a-d, with decimal values listed below them:
a b c d
1 2 3 4

Then if we subtract 'a' through 'd' from 'd', i.e. 'd' - 'a' and take the 
result modulo size, which is 4, then we get the following indices:
chars:          a b c d
dec values:     1 2 3 4
('d' - c)%4:    3 2 1 0

So to rotate around this set, if I need to add a number to a set element, I
can do the following for a given character c:
(c - 'a' + n)%4 + 'a'
For the complete character set:
(c - 'a' + n) % 26 + 'a'
In this case, n is equal to 14.
*/

package main

import (
  "fmt"
)

func main() {
  fmt.Println(rot14("hello World!"))
  fmt.Println(rot14("Hello! How are You?"))
}

func rot14(s string) string {
  str := []rune(s)
  for  i,v := range str {
    switch {
    case v >= 'A' && v <= 'Z':
      str[i] = (v - 'A' + 14) % 26 + 'A'
    case v >= 'a' && v <= 'z':
      str[i] = (v - 'a' + 14) % 26 + 'a'
    }
  }
  return string(str)
}
