/* Mon Mar 17 11:51:41 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "fmt"
    "math/rand"
    "time"
)

const SIZE = 10001
const MOD = 101
func main() {
    rand.NewSource(int64(time.Since(time.Now())))
    fmt.Print("\"[")
    for i:=0; i < SIZE; i++ {
        if i == SIZE-1 {
            fmt.Printf("%d]\"",rand.Int()%MOD)
        } else {
            fmt.Printf("%d, ",rand.Int()%MOD)
        }
    }
    fmt.Printf("\t\"%d\"\n",rand.Int()%MOD) }
