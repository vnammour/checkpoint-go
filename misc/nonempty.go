/* Mon Mar 10 11:55:56 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "fmt"
)

func main() {
    str := []string{"ab","","","cd"}
    fmt.Println(nonempty(str))
    for _,s := range str {
        fmt.Println(s)
    }
}

func nonempty(strings []string) []string {
    i := 0
    for _,s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}
