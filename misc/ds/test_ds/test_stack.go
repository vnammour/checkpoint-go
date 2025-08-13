/* Sun Feb 16 02:33:24 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
	"bytes"
	"ds"
	"fmt"
)

func main() {
	stack := &ds.Stack[byte]{}
	stack.Init(16)
	s := string("((\"\"))")
	buf := bytes.NewBufferString(s).Bytes()
	for _, v := range buf {
		switch v {
		case '(', '[', '{':
			stack.Push(v)
		case ')', ']', '}':
			t, ok := stack.Top()
			if !ok {
				break
			}
			if t == '(' && v == ')' || t == '[' && v == ']' || t == '{' && v == '}' {
				stack.Pop()
			}
		}
	}
	if stack.Len() != 0 {
		fmt.Printf("Unbalanced\n")
	} else {
		fmt.Printf("Balanced\n")
	}
	b, err := stack.Pop()
	fmt.Printf("b = %d\n", b)
	if err == nil {
		fmt.Println("nil  .... ")
	} else {
		fmt.Println(err)
	}
}
