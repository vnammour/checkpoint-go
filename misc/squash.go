/* Sun Mar  9 02:25:09 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
)

func main() {
    fmt.Println("main")
    s := []byte("       this is       a       sentence        ")
    s = squash(s)
    fmt.Printf(">%s<\n", string(s))
}

func isspace(b byte) bool {
    return b == '\t' || b == '\n' || b == '\v' || b == '\f' || b == '\r' || b == ' '
}

// "    this is     a     sent    "
func squash(s []byte) []byte {
    isSpaceAdded := false
    i,j := 0,0
    for i < len(s) {
        r,size := utf8.DecodeRune(s[i:])
        fmt.Printf(">%c<\n", r)
        if unicode.IsSpace(r) {
            if !isSpaceAdded {
                isSpaceAdded = true
                s[j] = ' '
                j++
            }
        } else {
            isSpaceAdded = false
            // I can use either (i.e. XOR) the commented below or the copy after it
            // size = utf8.EncodeRune(s[j:],r)
            copy(s[j:],s[i:i+size])
            j += size
            fmt.Println(j)
        }
        i += size
    }
    fmt.Printf("i = %d, j = %d, -- >%s<\n", i, j, string(s))
    return s[:j]
}
