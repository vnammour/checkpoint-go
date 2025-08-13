/* Thu Mar  6 11:50:41 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
    // "ds"
    "fmt"
    "time"
    // "math"
    "math/rand"
)

func main() {
    fmt.Printf("%v\n", time.Since(time.Now()))
    rand.NewSource(int64(time.Since(time.Now())))
    p := []int{5,6,1,4} // #of diffs  = (n-1)(n-1+1)/2 = n(n-1)/2 => O(n^2)
    for i:=0; i < 10; i++ {
        p = append(p,rand.Int()%11)
    }
    fmt.Println(p)
    b,s := bestbuy(p)
    fmt.Printf("buy @ p[%d] = %d, sell @ p[%d] = %d\n", b,p[b],s,p[s])
    b,s = best(p)
    fmt.Printf("buy @ p[%d] = %d, sell @ p[%d] = %d\n", b,p[b],s,p[s])
    p = []int{6,5,4,3,2,1,0}
    b,s = bestbuy(p)
    fmt.Printf("buy @ p[%d] = %d, sell @ p[%d] = %d\n", b,p[b],s,p[s])
}
// naive way: returns index of lowest buying price and highest selling price
func bestbuy(p []int) (int,int) {
    diff,b,s := 0,0,0
    for i := 0; i < len(p)-1; i++ {
        for j := i+1; j < len(p); j++ {
            if p[j] - p[i] > diff {
                diff = p[j] - p[i]
                b,s = i,j
            }
        }
    }
    return b,s
}
/*
0,1,2,3,4,5,6
-1-1-1-1-1-1 = -6;
*/
func best(p []int) (int,int) {
    diff := 0
    s := -1
    for i:=0; i < len(p)-1; i++ {
        if diff > p[i] - p[i+1] {
            diff = p[i] - p[i+1]
            s = i+1
        }
    }
    return 0,s
}
