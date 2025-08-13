/* Thu Feb 13 11:59:30 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "github.com/01-edu/z01"
)

func main() {
  printstr("Hello World!")
}

func printstr(s string) {
  for _,r := range s {
    z01.PrintRune(r)
  }
}
