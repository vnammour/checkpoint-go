/* Wed Mar 19 11:31:32 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 4
level: 9
allowed: builtin
Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.
If it's an empty string return G followed by a newline \n.
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Print(printif("abcdefz"))
}

func printif(s string) string {
    if len(s) >= 3 {
        return "G\n"
    } else {
        return "Invalid Input\n"
    }
}
