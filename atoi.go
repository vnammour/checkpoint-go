/* Mon Aug 11 03:10:54 PM IDT 2025 */
/* By: Victor Nammour */
package main

import (
	"fmt"
	"errors"
)

func main() {
	n,err := atoi("   -834008")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("n = %d\n", n)
	}
}

func atoi(s string) (int,error) {
	sign := 1
	i := 0
	for ; i < len(s); i++ {
		if isspace(s[i]) {
			continue
		}
		if s[i] == '-' {
			sign = -1
		}
		if s[i] == '-' || s[i] == '+' {
			continue
		}
		if isdigit(s[i]) {
			break
		}
	}

	fmt.Printf("i = %d\n",i)

	n := 0
	for ; i < len(s); i++ {
		if !isdigit(s[i]) {
			return 0, errors.New("wrong number format")
		}
		n = 10 * n + int(s[i] - '0')
	}
	fmt.Println(n * sign)
	n *= sign

	return n, nil
}

func isspace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\v'
}
func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}
