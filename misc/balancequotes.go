/* Sun Feb 16 11:31:11 AM IST 2025 */
/* By: Victor Nammour */
package main

import (
	"ds"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("isBalancedQuotes(%s)", os.Args[i])
		fmt.Printf("--->%v\n", isBalancedQuotes(os.Args[i]))
	}
	tests := []string{
		`Hello "world"`,           // true
		`Hello "world`,            // false
		`"It's a test"`,           // true (single inside double)
		`"It\'s a test`,           // false (unclosed double quote)
		`"Hello 'nested' test"`,   // true (nested single inside double)
		`"Hello 'nested test"`,    // false (unclosed single inside double)
		`'Single "inside" works'`, // true (nested double inside single)
		`'Single "inside works'`,  // false (unclosed double inside single)
	}

	for _, test := range tests {
		fmt.Printf(">%s< -> %v\n", test, isBalancedQuotes(test))
	}
}

func isOpeningBracket(b byte) bool {
	return b == '(' || b == '[' || b == '}'
}

func isBalancedQuotes(s string) bool {
	stack := &ds.Stack[byte]{}
	stack.Init(uint(len(s)))
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) { // skip escaped quotes
			i++
			continue
		}
		top, ok := stack.Top()
		switch s[i] {
		case '"':
			if ok && top == '"' {
				stack.Pop()
			} else {
				stack.Push(s[i])
			}
		case '\'':
			if ok && top == '\'' {
				stack.Pop()
			} else {
				stack.Push(s[i])
			}
		}
	}
	//stack.Print()
	return stack.Len() == 0
}
