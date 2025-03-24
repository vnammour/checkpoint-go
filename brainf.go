/* Thu Mar 13 02:10:21 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 20
level: 32
allowed: os.*, github.com/01-edu/z01.PrintRune, --allow-builtin
Write a Brainfuck interpreter program.
The source code will be given as the first parameter, and will always be valid with fewer than 4096 operations.
Your Brainfuck interpreter will consist of an array of 2048 bytes, all initialized to 0, with a pointer to the first byte.
Every operator consists of a single character:
>: increment the pointer.
<: decrement the pointer.
+: increment the pointed byte.
-: decrement the pointed byte.
.: print the pointed byte to standard output.
[: If the byte at the current pointer is 0, skip forward to the command after the matching ].
]: If the byte at the current pointer is not 0, jump back to the command after the matching [.
Any other character is treated as a comment and ignored.

Usage:
$ go run . "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>." | cat -e
Hello World!$
$ go run . "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>." | cat -e
Hi$
$ go run . "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++." | cat -e
abc$
$ go run .
$
*/
package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(0)
    }
    brainf(os.Args[1])
}

const SIZE = 2048
func brainf(program string) {
    if len(program) >= 2 * SIZE { return }
    interpreter := [SIZE]byte{}         // to represent the interpreter, an array of 2048 bytes all init'd to zero
    var rn int                              // pointer to first byte in interpreter
    stack := make([]byte,0,SIZE/100)
    for pc:=0; pc < len(program); pc++ {    // pc is the program pointer/counter
        switch program[pc] {
        // case '>': rn = (rn + 1) % SIZE
        case '>': if rn < SIZE { rn++ }
        // case '<': rn = (rn - 1 + SIZE) % SIZE
        case '<': if rn > 0 { rn-- }
        case '+': interpreter[rn]++;
        case '-': interpreter[rn]--;
        case '.': z01.PrintRune(rune(interpreter[rn]))
        case '[':
            if interpreter[rn] == 0 {
                stack = append(stack,'[')
                for pc++ ; pc < len(program); pc++ {
                    switch program[pc] {
                    case '[': stack = append(stack,'[');
                    case ']': stack = stack[:len(stack)-1];
                    }
                    if len(stack) == 0 {
                        break
                    }
                }
                if len(stack) != 0 {
                    os.Exit(1) // error! mismatched brackets
                }
            }
        case ']':
            if interpreter[rn] != 0 {
                stack = append(stack,']')
                for pc--; pc >=0; pc-- {
                    switch program[pc] {
                    case '[': stack = stack[:len(stack)-1];
                    case ']': stack = append(stack,']')
                    }
                    if len(stack) == 0 {
                        break
                    }
                }
                if len(stack) != 0 {
                    os.Exit(1) // error! mismatched brackets
                }
            }
        }
    }
}
