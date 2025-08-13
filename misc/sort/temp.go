/* Wed Feb 19 08:19:14 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "fmt"
)

func main() {
  fmt.Println("in main")
  arr := []int{90,-10,4,9}//,11,3,4,6,0,-1,3,8,-1,1,7,55}
  fmt.Println("Starting...",arr)
  out := make([]int,len(arr))
  mergesort(out,arr)
  fmt.Println("finally: ",out)
}

func mergesort(out,arr []int) []int {
  fmt.Printf("len(out) = %d, arr = %v\n", len(out), arr)
  if len(arr) == 0 || len(arr) == 1 {
    return arr
  }
  left := arr[0:len(arr)/2]
  right := arr[len(arr)/2:]
  //fmt.Println("out,mergesort(",left,")", "and mergesort(",right,")")
  return merge(out,mergesort(out[:len(left)],left), mergesort(out[len(left):],right)) 
}

func merge(out, left, right []int) []int {
  i,j,k := 0,0,0
  //fmt.Printf("len(out) = %d\n", len(out))
  for ; i < len(left) && j < len(right); k++ {
    if left[i] < right[j] {
      //fmt.Println("out[",i,"] = ", left[i])
      out[k] = left[i]
      i++
    } else {
      //fmt.Println("out[",j,"] = ", right[j])
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
