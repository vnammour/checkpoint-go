/* Sun Mar  9 04:06:19 PM IST 2025 */
/* By: Victor Nammour */
/*
The Go Programming Language
Ex 4.5 p. 93
Write an in-place function to eliminate adjacent duplicates in a []string slice
*/
package main

import (
    "fmt"
)

func main() {
    s := []string{"a","a","a","b","b","c","c","c"}
    sc := removeDuplicates(s)
    fmt.Println(sc)
    fmt.Println(s)
    s =[]string{"a","a","a"}
    sc = removeDuplicates(s)
    fmt.Println(sc)
}

func removeDuplicates(s []string) []string {
    if len(s) == 0 {
        return s
    }
    t := s[1:]
    i := 0
    for j := 0; j < len(t); j++ {
        if s[i] != t[j] {
            s[i+1] = t[j]
            i++
        }
    }
    return s[:i+1] // to account for last element in s
}
