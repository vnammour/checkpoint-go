/* Mon Feb 17 03:50:42 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "fmt"
)

func main() {
  delims := ":;"
  input := ":;this:is;a:test::;a;"
  tokens := strtok(input,delims)
  fmt.Println(tokens)
  for _,v := range tokens {
    fmt.Printf(">%s<\n", v)
  }
}

/* Finding tokens in a string:
Loop over the input string. For each character in input,
check it against all characters in delimiters.
If there is a match, then check if both indices i & j are at same
position: If so, then we have no token; else, append slice of length
j - i starting at i, and set i to point to next character after j.
*/
func strtok(input,delimiters string) []string {
  tokens := []string{}
  i,j:= 0,0
  for ; j < len(input); j++ {
    for k := 0; k < len(delimiters); k++ {
      if delimiters[k] == input[j] {
        if i == j {
          i++
        } else {
          tokens = append(tokens,input[i:j])
          i = j + 1
        }
        break;
      }
    }
  }
  // This in order to get the last token in input string that didn't come across delims (if any)
  if i != j {
    tokens = append(tokens,input[i:j])
  }
  return tokens
}
