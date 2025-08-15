/* Tue Mar 18 11:36:17 PM IST 2025 */
/* By: Victor Nammour */
/*
Difficulty: 20
Level: 32
allowed: fmt.Printf, os.Args, builtin
Write a program that receives two strings and replicates the use of brackets in regular expressions. Brackets in regular expressions returns the words that contain the expression inside of it.
The program should handle the "|" operator, which searches for both strings on each side of the operator.
The output of the program should be, the results of the regular expression, numbered and displayed by the order of appearance in the string.
If the number of arguments is different from 2, if the regular expression is not valid, if the last argument is empty or if there are no matches, the program should print nothing.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	grouping(os.Args[1], os.Args[2])
}

func grouping(exp, input string) {
	exp = trimedges(exp)
	if len(exp) == 0 || len(input) == 0 {
		return
	}
	exps := tokenize(exp, ispipe)
	tokens := tokenize(input, isspace)
	// fmt.Println("tokens:",tokens)
	count := 0
	for _, token := range tokens {
		for _, rg := range exps {
			if contains(token, rg) {
				count++
				fmt.Printf("%d: %s\n", count, token)
			}
		}
	}
}

/* Assuming ASCII */
func isupper(r byte) bool  { return r >= 'A' && r <= 'Z' }
func islower(r byte) bool  { return r >= 'a' && r <= 'z' }
func isletter(r byte) bool { return isupper(r) || islower(r) }
func isspace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\v' || b == '\n' || b == '\r'
}
func ispipe(b byte) bool {
	return b == '|'
}

func hasprefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if hasprefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func trimedges(exp string) string {
	// trim leading and trailing spaces if present
	i, j := 0, len(exp)-1
	for ; i < len(exp) && isspace(exp[i]); i++ {
	}
	for ; j >= 0 && isspace(exp[j]); j-- {
	}
	j++
	exp = exp[i:j]
	// trim opening and closing brackets if present
	if exp[0] != '(' && exp[len(exp)-1] != ')' {
		return ""
	} else {
		exp = exp[1 : len(exp)-1]
	}
	return exp
}

func tokenize(exp string, f func(byte) bool) []string {
	exps := make([]string, 0, len(exp))
	i := 0
	in := false
	for j := 0; j < len(exp); j++ {
		if f(exp[j]) {
			continue
		}
		in = true
		i = j
		for ; j < len(exp) && !f(exp[j]); j++ {
		}
		in = false
		if !in {
			s := exp[i:j]
			// fmt.Printf(">%s<\n",s)
			// remove punctuation from start and end of string if present
			if len(s) > 0 && !isletter(s[0]) {
				k := 0
				for ; k < len(s) && !isletter(s[k]); k++ {
				}
				s = s[k:]
			}
			if len(s) > 0 && !isletter(s[len(s)-1]) {
				// fmt.Printf("len(s) = %d, Before >%s<\n", len(s),s)
				k := len(s) - 1
				for ; k >= 0 && !isletter(s[k]); k-- {
				}
				k++
				s = s[:k]
				// s = s[:len(s)-1]
				// fmt.Printf("len(s) = %d, After >%s<\n", len(s),s)
			}
			exps = append(exps, s)
		}
	}
	return exps
}
