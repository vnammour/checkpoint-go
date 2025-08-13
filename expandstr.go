/* Fri Mar 14 13:08:02 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 8
level: 9
allowed: github.com/01-edu/z01.PrintRune, os.*, builtin
Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.
The string will be followed by a newline ('\n').
A word, in this exercise, is a sequence of visible characters.
If the number of arguments is not 1, or if there are no word, the program displays nothing.
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 { os.Exit(0) }
    fmt.Print(expandstr(os.Args[1]))
}

/* Using ascii, this is what I found. */
func visible(b byte) bool {
    return b >= 33 && b <= 126
}

func expandstr(s string) string {
    out := make([]byte,0,len(s))
    for i:=0; i < len(s); {
        if visible(s[i]) {
            out = append(out,s[i])
            i++
        } else {
            if len(out) == 0 {  // processing invisible bytes at beginning of string
                i++
            } else {
                out = append(out,' ', ' ', ' ')
                for i++;i<len(s) && !visible(s[i]); i++ {}
            }
        }
    }
    // removing invisible bytes at end of slice
    if len(out) > 0 {
        i := len(out) - 1
        for ; i>=0 && !visible(out[i]); i-- {}
        out = append(out[:i+1],'\n')
    }
    return string(out)
}
