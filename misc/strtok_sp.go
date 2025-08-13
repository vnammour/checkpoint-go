/* Sun Mar 16 03:11:08 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "fmt"
)

func main() {
    s := "        this is    a short sentence    "
    s1 := "this is    a short sentence"
    fmt.Println(strtok(s))
    fmt.Println(strtok(s1))
    fmt.Println(tostring(strtok(s)))
}

func isspace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\v' || b == '\n' || b == '\r'
}

/* Tokenize using space as a delimiter */
func strtok(s string) []string {
    i := 0
    in := false
    words := make([]string,0,len(s)/2)
    for j:=0; j < len(s); j++ {
        if isspace(s[j]) { continue } 
        in = true
        i = j
        for ; j < len(s) && !isspace(s[j]); j++ {}
        in = false
        if !in {
            words = append(words,s[i:j])
        }
    }
    return words
}

func tostring(s []string) string {
    out := make([]byte,0,len(s)*2)
    for i:= 0; i < len(s); i++ {
        out = append(out,[]byte(s[i])...)
        if i != len(s) - 1 {
            out = append(out,' ')
        }
    }
    return string(out)
}
