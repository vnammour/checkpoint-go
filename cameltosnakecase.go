/* Fri Mar 14 04:18:02 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 6
level: 9
allowed: builtin
Write a function that converts a string from camelCase to snake_case.
If the string is empty, return an empty string.
If the string is not camelCase, return the string unchanged.
If the string is camelCase, return the snake_case version of the string.
For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:
lowerCamelCase
UpperCamelCase
Rules for writing in camelCase:
The word does not end on a capitalized letter (CamelCasE).
No two capitalized letters shall follow directly each other (CamelCAse).
Numbers or punctuation are not allowed in the word anywhere (camelCase1).
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(cameltosnakecase("CamelCase"))
	fmt.Println(cameltosnakecase("HelloWorld"))
	fmt.Println(cameltosnakecase("helloWorld"))
	fmt.Println(cameltosnakecase("camelCase"))
	fmt.Println(cameltosnakecase("CAMELtoSnackCASE"))
	fmt.Println(cameltosnakecase("camelToSnakeCase"))
	fmt.Println(cameltosnakecase("heyDay2"))
    fmt.Println(cameltosnakecase("A"))
}

/* Assuming ASCII */
func isupper(b byte) bool  { return b >= 'A' && b <= 'Z' }
func islower(b byte) bool  { return b >= 'a' && b <= 'z' }
func isletter(b byte) bool { return isupper(b) || islower(b) }
func tolower(b byte) byte  { return b - 'A' + 'a' }

func cameltosnakecase(s string) string {
    if len(s) == 0 {
        return s
    }
    if isupper(s[len(s)-1]) {
        return s
    }
    out := make([]byte,0,len(s)+2)
    for i:=0; i < len(s); i++ {
        if !isletter(s[i]) {
            return s
        }
        if isupper(s[i]) {
            if i == 0 {
                out = append(out,tolower(s[i]))
            } else {
                if isupper(s[i-1]) { // two consecutive upper case letters
                    return s
                }
                out = append(out,'_',tolower(s[i]))
            }
        } else {
            out = append(out,s[i])
        }
    }
    return string(out)
}
