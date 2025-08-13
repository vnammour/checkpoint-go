/* Mon Jan 20 11:03:30 PM IST 2025 */
/* By: Victor Nammour */
package piscine

import (
)

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool,len(a))
	for i,v := range a {
		result[i] = f(v)
	}
	return result
}
