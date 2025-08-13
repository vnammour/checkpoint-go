/* Wed Feb 19 08:19:14 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "fmt"
)

func main() {
  fmt.Println("in main")
  arr := []int{90,-10,4,9,11,3,4,6,0,-1,3,8,-1,1,7,55}
  fmt.Println("Starting...",arr)
  out := mergesort(arr)
  fmt.Println("finally: ",out)
}

func mergesort(arr []int) []int {
  if len(arr) == 0 || len(arr) == 1 {
    return arr
  }
  left := arr[0:len(arr)/2]
  right := arr[len(arr)/2:]
  return merge(mergesort(left),mergesort(right))
}

func merge(left, right []int) []int {
  out := make([]int,len(left)+len(right))
  i,j,k := 0,0,0
  for ; i < len(left) && j < len(right); k++ {
    if left[i] < right[j] {
      out[k] = left[i]
      i++
    } else {
      out[k] = right[j]
      j++
    }
  }
  for ; i < len(left); i,k = i+1,k+1 {
    out[k] = left[i]
  }
  for ; j < len(right); j,k = j+1,k+1 {
    out[k] = right[j]
  }
  return out
}
