/* Sun Feb  9 03:30:36 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
	"ds"
	"fmt"
)

type Pair struct {
	Index1 int
	Index2 int
}

func findpairs(arr []int, sum int) []Pair {
	pairs := make([]Pair, 0, 10)
	fmap := make(map[int]*(ds.List[Pair]))
	for i, v := range arr {
		comp := sum - v
		if l, exists := fmap[comp]; exists {
			for n := l.Head; n != nil; n = n.Next {
				pairs = append(pairs, Pair{n.Data.Index1, i})
			}
		}
        if fmap[v] == nil {
            fmap[v] = &ds.List[Pair]{} 
        }
		fmap[v].Append(Pair{Index1: i})
	}
	return pairs
}

func main() {
	arr := []int{1, 3, 9, 4, 6, 6, 11, -1, -5, 1, 4}
	sum := 10
	pairs := findpairs(arr, sum)
	fmt.Println(pairs)
}
