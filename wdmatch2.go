/* Thu Mar 13 12:55:31 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: os.Args, github.com/01-edu/z01.PrintRune, builtin
Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string. If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing. If the number of arguments is different from 2, the program displays nothing.
*/

package main

import (
    // "fmt"
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 3 { return }
    wdmatch(os.Args[1], os.Args[2])
}

func wdmatch(s1,s2 string) {
    m := make(map[rune][]int,len(s2))
    for i,r := range s2 {
        if _,ok := m[r]; !ok {
            m[r] = append(make([]int,0,3),i)
        } else { m[r] = append(m[r],i) }
    }

    last_index := -1
    prnt := true // controls whether to continue checking or to shortcut
    for _,r := range s1 {
        if !prnt { break; }
        if v,exists := m[r]; exists {
            if len(v) > 0 {
                set := false    // I need another boolean since "prnt" above affects the whole string s1
                for i,k := range m[r] {
                    if last_index < k {
                        last_index = k
                        set = true
                        m[r] = m[r][i+1:]
                        break;
                    }
                }
                if !set { prnt = false; }
            } else { prnt = false; break; }
        }
    }
    if prnt {
        for _,r := range s1 {
            z01.PrintRune(r)
        }
        z01.PrintRune('\n')
    }
}
