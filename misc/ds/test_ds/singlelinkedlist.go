/* Wed Feb  5 10:45:34 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "fmt"
  "ds"
)

func main() {
  l := &ds.List[int]{}
  arr := []int{1,2,3,4,5,6}
  for _,v := range arr {
    l.Append(v)
  }
  l.Print("l")
  l.Reverse()
  l.Print("l")
  fmt.Println("# of nodes in l = ", l.Count())
  l2 := &ds.List[int]{}
  for _,v := range arr {
    l2.Prepend(v)
  }
  l2.Delete(10)
  l2.Delete(5)
  l2.Print("l2")
}
