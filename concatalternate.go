/* Mon Mar 24 10:38:47 AM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: builtin
Write a function ConcatAlternate() that receives two slices of an int as arguments and returns a new slice with the result of the alternated values of each slice.
The input slices can be of different lengths.
The new slice should start with an element of the largest slice.
If the slices are of equal length, the new slice should return the elements of the first slice first and then the elements of the second slice.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(concatalternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(concatalternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(concatalternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(concatalternate([]int{1, 2, 3}, []int{}))
}

func concatalternate(x, y []int) []int {
	z := make([]int, 0, len(x)+len(y))
	fst, snd := x, y // snd will be such that len(snd) <= len(fst)
	if len(x) < len(y) {
		fst, snd = y, x
	}
	i := 0
	for ; i < len(snd); i++ {
		z = append(z, fst[i], snd[i])
	}
	z = append(z, fst[i:]...)
	return z
}

func concatalternate2(x, y []int) []int {
	xlen, ylen := len(x), len(y)
	z := make([]int, 0, xlen+ylen)
	i := 0
	if xlen < ylen {
		for ; i < xlen; i++ {
			z = append(z, y[i], x[i])
		}
		z = append(z, y[i:]...)
	} else if xlen > ylen {
		for ; i < ylen; i++ {
			z = append(z, x[i], y[i])
		}
		z = append(z, x[i:]...)
	} else {
		for ; i < xlen; i++ {
			z = append(z, x[i], y[i])
		}
	}
	return z
}

/* Note the difference in the use of make from the other two functions above. */
func concatalternate3(x, y []int) []int {
	xlen, ylen := len(x), len(y)
	z := make([]int, xlen+ylen)
	i, j := 0, 0
	if xlen < ylen {
		for ; i < xlen; i, j = i+1, j+2 {
			z[j], z[j+1] = y[i], x[i]
		}
		for ; i < ylen; i, j = i+1, j+1 {
			z[j] = y[i]
		}
	} else if xlen > ylen {
		for ; i < ylen; i, j = i+1, j+2 {
			z[j], z[j+1] = y[i], x[i]
		}
		for ; i < xlen; i, j = i+1, j+1 {
			z[j] = x[i]
		}
	} else {
		for ; i < xlen; i, j = i+1, j+2 {
			z[j], z[j+1] = x[i], y[i]
		}
	}
	return z
}
