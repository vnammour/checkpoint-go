/* Mon Mar 17 08:20:57 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 18
level: 32
allowed: strconv.Atoi, os.*, fmt.*, strings.Split, builtin
Write a program that takes an undefined number of string in arguments. For each argument, if the expression is correctly bracketed, the program prints on the standard output OK followed by a newline ('\n'), otherwise it prints Error followed by a newline.
Symbols considered as brackets are parentheses ( and ), square brackets [ and ] and curly braces { and }. Every other symbols are simply ignored.
An opening bracket must always be closed by the good closing bracket in the correct order. A string which does not contain any bracket is considered as a correctly bracketed string.
If there is no argument, the program must print nothing.
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    for _,s := range args {
        if brackets(s) {
            fmt.Println("OK!")
        } else {
            fmt.Println("Error")
        }
    }
}

func brackets(s string) bool {
    stack := make([]byte,0,len(s))
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case '(': stack = append(stack,'(')
        case ')':
            if stack[len(stack)-1] == '(' {
                stack = stack[:len(stack)-1]
            }
        case '[': stack = append(stack, '[')
        case ']':
            if stack[len(stack)-1] == '[' {
                stack = stack[:len(stack)-1]
            }
        case '{': stack = append(stack, '{')
        case '}':
            if stack[len(stack)-1] == '{' {
                stack = stack[:len(stack)-1]
            }
        }
    }
    return len(stack) == 0
}
