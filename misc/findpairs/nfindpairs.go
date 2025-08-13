/* Sun Feb  9 03:12:55 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "fmt"
)
type Pair struct {
  Index1 int
  Index2 int
}
func findpairs(arr []int,sum int) []Pair {
  pairs := make([]Pair,0,10)
  for i,v := range arr {
    for j:=i+1; j < len(arr); j++ {
      if arr[j] == sum - v {
        pairs = append(pairs,Pair{i,j})
      }
    }
  }
  return pairs
}

func main() {
  arr:= []int{1,3,9,4,6,6,11,-1,-5,1}
  sum:= 10
  pairs := findpairs(arr,sum)
  fmt.Println(pairs)
}
