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
}

/* Assuming ASCII */
func isupper(r byte) bool { return r >= 'A' && r <= 'Z' }
func islower(r byte) bool { return r >= 'a' && r <= 'z' }
func isletter(r byte) bool { return isupper(r) || islower(r) }
func tolower(r byte) byte { return r - 'A' + 'a' }

func cameltosnakecase(s string) string {
    if len(s) == 0 { return s }
    out:= make([]byte,0,len(s)+3) // initial capacity of len(s) (+3 to account for underscores)
    var prev int = 0
    for i:=0; i < len(s); i++ {
        if !isletter(s[i]) { return s }
        if isupper(s[i]) {
            if i - prev == 1 {     // two consecutive upper case bytes
                return s
            }
            if i == 0 {
                out = append(out,tolower(s[i]))
            } else {
                out = append(out,'_',tolower(s[i]))
            }
            prev = i
        } else { out = append(out,s[i]) }
    }
    return string(out)
}
