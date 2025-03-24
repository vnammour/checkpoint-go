/* Tue Mar 18 15:07:26 PM IST 2025 */
/* By: Victor Nammour */
/*
allowed: os.Args, fmt.Println, strconv.*, builtin
https://github.com/01-edu/public/tree/master/subjects/romannumbers
Write a program called rn. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.
The program should have a limit of 4000. In case of an invalid number, for example "hello" or 0 the program should print ERROR: cannot convert to roman digit.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	romannumbers(os.Args[1])
}

func romannumbers(s string) {
	s = trim(s, ' ')
	// initial error checking. 4000 is a limit
	if len(s) > 4 || !isdigit(s[0]) || len(s) == 4 && int(s[0]-'0') >= 4 {
		fmt.Println("ERROR: cannot convert to roman digit.")
		return
	}
	rn := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
		100: "C", 200: "CC", 300: "CCC", 400: "CD", 500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM",
		1000: "M"}
	result := make([]string, 0, 10)
	interm := make([]string, 0, 10)
	ranks := [4]int{1, 10, 100, 1000}
	for i := 0; i < len(s); i++ {
		p := len(s) - i // 4(thousands), 3(hundreds), 2(tens), 1(ones)
		if !isdigit(s[i]) {
			fmt.Println("ERROR: cannot convert to roman digit.")
			return
		}
		n := int(s[i] - '0') // get digit at index i
		switch p {
		case 4: // thousands
			for i := 0; i < n; i++ {
				result = append(result, rn[1000])
				interm = append(interm, rn[1000], "+")
			}
		case 3: // hundreds
			result = append(result, rn[n*100])
		case 2: // tens
			result = append(result, rn[n*10])
		case 1: // ones
			result = append(result, rn[n])
		}
		if p != 4 {
			switch {
			case n >= 1 && n <= 3:
				for i := 0; i < n; i++ {
					interm = append(interm, rn[ranks[p-1]], "+")
				}
			case n == 4:
				interm = append(interm, "(", rn[5*ranks[p-1]], "-", rn[ranks[p-1]], ")", "+")
			case n >= 5 && n <= 8:
				interm = append(interm, rn[5*ranks[p-1]])
				for i := 0; i < n; i++ {
					interm = append(interm, "+", rn[ranks[p-1]])
				}
			case n == 9:
				interm = append(interm, "(", rn[ranks[p]], "-", rn[ranks[p-1]], ")", "+")
			}
		}
	}
	b := interm[len(interm)-1]
	if b == "+" || b == "-" {
		interm = interm[:len(interm)-1]
	}
	for i := 0; i < len(interm); i++ {
		fmt.Printf("%s", interm[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(result); i++ {
		fmt.Printf("%s", result[i])
	}
	fmt.Printf("\n")
}

func isdigit(b byte) bool {
	return b >= '0' && b <= '9'
}

/* remove leading and trailing bytes b */
func trim(s string, b byte) string {
	i := 0
	for ; i < len(s) && s[i] == b; i++ {
	}
	s = s[i:]
	j := len(s) - 1
	for ; j > 0 && s[j] == b; j-- {
	}
	j++
	return s[:j]
}
