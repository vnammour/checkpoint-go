/* Tue Mar  4 12:28:57 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // fmt.Printf("num of args = %d\n", len(os.Args))
    if len(os.Args) == 1 {  // stdin
        err := cat(os.Stdin.Name())
        if err != nil {
            fmt.Fprintf(os.Stderr,"%s\n",err)
        }
    } else {                // file args
        for i:=1; i < len(os.Args); i++ {
            err := cat(os.Args[i])
            if err != nil {
                fmt.Fprintf(os.Stderr,"%s\n",err)
            }
        }
    }
}

/* open file filename for reading and print contents to stdout after replacing the LF */
func cat(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    input := bufio.NewScanner(file)
    for input.Scan() {
        fmt.Fprintf(os.Stdout,"%s\n", input.Text())
    }
    return file.Close()
}
