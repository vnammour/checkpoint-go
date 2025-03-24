/* Wed Mar 19 11:34:34 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: github.com/01-edu/z01.PrintRune, os.Args, builtin
Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).
If the number of arguments is different from 3, the program displays nothing.
If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').


*/
package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 4 {
        os.Exit(0)
    }
    searchreplace(os.Args[1],os.Args[2], os.Args[3])
}

/* I chose to use copy with subslices rather than have f
converted into a rune slice and use that */
func searchreplace(str, srch, rep string) {
    s := []rune(str)
    to := []rune(rep)
    for i,_ := range s {
        if str[i:i+1] == srch {
            copy(s[i:i+1],to[:])
        }
        z01.PrintRune(s[i])
    }
   z01.PrintRune('\n')
}
