/* Sat Mar 15 04:41:46 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "fmt"
    "unsafe"
    "reflect"
)

func main() {
    s := make([]int,0,2)
    fmt.Printf("main: len = %d, cap = %d, s = %v\n", len(s), cap(s), s)
    f(s)
    sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
    fmt.Println("main:",sh)
    fmt.Println()
    // t := s[:1]
    // fmt.Println(t)
    // sh = (*reflect.SliceHeader)(unsafe.Pointer(&t))
    // fmt.Println("main:",sh)
}

func f(x []int) {
    x = append(x,1)
    fmt.Printf("f(): len = %d, cap = %d, x = %v\n", len(x), cap(x), x)
    sh := (*reflect.SliceHeader)(unsafe.Pointer(&x))
    fmt.Println("f:",sh)
}
