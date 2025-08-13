/* Wed Mar 12 07:37:19 AM IST 2025 */
/* By: Victor Nammour */
/* Go Programming Language, 4.3 Maps */
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    dedup()
}

func dedup() {
    seen := make(map[string]bool)   // a set of strings
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}

/*
Sometimes we need a map or set whose keys are slices, but because a map's keys must be comparable,
this cannot be expressed directly. However, it can be done in two steps. First we define a helper function
k that maps each key to a string, with the property that k(x) == k(y) if and only if we consider x and y
equivalent. Then we create a map whose keys are strings, applying the helper function to each key before
we access the map.

The example below uses a map to record the number of times Add has been called with a given list of strings.
It uses fmt.Sprintf to convert a slice of strings into a single string that is a suitable map key, quoting
each slice element with %q to record string boundaries faithfully:
*/

var m = make(map[string]int)
func k(list []string) string { return fmt.Sprintf("%q", list) }
func Add(list []string) { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
