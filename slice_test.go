/* Wed Mar 19 11:50:45 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "fmt"
)

func main() {
    // x := []int{1,2,3}
    y := []int{4,5,6}
    x := make([]int,3,6)
    copy(x,[]int{1,2,3})
    fmt.Println(f(x,y))
    fmt.Println(x)
    x = x[:cap(x)]
    fmt.Println(x)
}

func f(x,y []int) []int {
    return append(x,y...)
}
