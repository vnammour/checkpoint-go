/* Tue Jan 21 09:09:02 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "fmt"
)

func main() {
  slice := []int{5,1,0,6,2,4}
  insertsort(slice)
  fmt.Println(slice)
}

func insertsort(items []int) {
  var i, j int
  for i = 1; i < len(items); i++ {
    j = i
    for j > 0 && items[j] < items[j-1] {
      items[j],items[j-1] = items[j-1],items[j]
      j--
    }
  }
}

func binarysearch(items []int, num int) int {
  var low, high , mid int = 0, len(items), 0
  for low < high {
    mid = (low + high)/2
    switch {
    case items[mid] > num:
      high = mid
      mid = (low+high)/2
    case items[mid] < num:
      low = mid + 1
      mid = (low + high)/2
    default:
      return mid
    }
  }
  return -1
}
