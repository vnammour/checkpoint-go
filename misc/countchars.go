/* Wed Mar 12 09:11:16 AM IST 2025 */
/* By: Victor Nammour */
/*
Go Programming Language, 4.3 Maps
Here’s another example of maps in action, a program that counts the occurrences of each distinct Unicode code point
in its input. Since there are a large number of possible characters, only a small fraction of which would appear in
any particular document, a map is a natural way to keep track of just the ones that have been seen and their
corresponding counts.
*/
package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

func main() {
    fmt.Println(countchars())
}

func countchars() int {
    var m = make(map[rune]int) 
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        // line := input.Text()
        for _,r := range input.Text() {
            m[r]++
        }
    }
    keys := make([]rune,0,len(m))
    for k := range m {
        keys = append(keys,k)
    }
    sort.Slice(keys, func(i,j int) bool {
        return keys[i] < keys[j]
    })
    for _,k := range keys {
        fmt.Printf(">%c<\t%d\n", k, m[k])
    }
    return len(m)
}
