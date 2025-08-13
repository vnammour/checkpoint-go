/* Sun Feb  9 03:30:36 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
	"fmt"
)

type Pair struct {
	Index1 int
	Index2 int
}

/*
Indices 4,10 are not accounted for because the hashtable's value

	at key = 6 and key = 4 are overwritten: at key 6 the index 4 is overwritten by 5
	and at key 4 the index 3 by 10
	So I need to insert nodes of a linked list in the map.
*/
func findpairs(arr []int, sum int) []Pair {
	pairs := make([]Pair, 0, 10)
	fmap := make(map[int][]int, len(arr))
	for i, v := range arr {
		comp := sum - v
		if indices, exists := fmap[comp]; exists {
            for _,j := range indices {
			    // pairs = append(pairs, Pair{indices[0], i})
			    pairs = append(pairs, Pair{j,i})
            }
            // fmap[v] = indices[:0]
		}
		fmap[v] = append(fmap[v],i)
	}
	return pairs
}

func main() {
	arr := []int{1, 3, 9, 4, 6, 6, 11, -1, -5, 1, 4}
	// arr := []int{2,6,4,8,-12,4}
	sum := 10
	pairs := findpairs(arr, sum)
	fmt.Println(pairs)
}
