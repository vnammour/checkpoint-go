/* Fri Feb  7 12:35:45 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 12
level: 32
allowed: builtin, strings.String
Write a function FifthAndSkip() that takes a string and returns another string. The function separates every five characters of the string with a space and removes the sixth one.
If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you reach a length of 5.
If the string is less than 5 characters return Invalid Input followed by a newline \n.
If the string is empty return a newline \n.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(fifthandskip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(fifthandskip("This is a short sentence"))
	fmt.Print(fifthandskip("1234"))
}
func isspace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\v' || b == '\t' || b == '\r'
}
/* allowed: strings.String, which is a reciver of type strings.Builder */
func fifthandskip(s string) string {
	const N = 5
    if len(s) == 0 {
        return "\n"
    }
    if len(s) < N {
        return "Invalid Input\n"
    }
    j := 0
	b := &strings.Builder{}
	b.Grow(len(s))
    for i:=0; i < len(s); i++ {
		if isspace(s[i]) {
			continue
		}
		j++ // this has to be before the mod check, otherwise it will only write up to first n bytes
		if j > 1 && j%N == 1 {
			j = 0
			b.WriteRune(' ')
			continue
		}
		b.WriteByte(s[i])
	}
    b.WriteByte('\n')
	return b.String()
}

/* Since I am not using append on out, I have j as an index into it. */
func fifthandskip2(str string) string {
	if len(str) == 0 {
		return "\n"
	}
  	if len(str) <5 {
		return "Invalid Input\n"
	} 
	s := []byte(str)
	out := make([]byte,len(s)+1) // +1 for ending newline
	n, i, j := 0,0,0
    for i < len(s) {
		if isspace(s[i]) {
            i++
			continue
		}
		if (n+1) % 6 == 0 {
			out[j] = ' '
            i,j = i+1,j+1
			n = 0
			continue
		}
		n++
		out[j] = s[i]
        i,j = i+1,j+1
	}
	out[j] = '\n'
    return string(out[:j+1])
}
