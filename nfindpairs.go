/* Tue Mar 18 16:12:55 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 14
level: 32
allowed: builtin, fmt.*, os.Args, strconv.Atoi, strings.*
Write a program that finds all pairs of elements in an integer array that sum up to a given target value. The program should output a list of pairs, each representing the indices of the elements that form the pair.
In this exercise you must take in consideration the following:
Ensure it's possible to have positive or negative integers in the array.
Ensure each element is used only once in a pair, although the element can be repeated in different pairs.
Allow for multiple pairs to sum up to the target value.
The output messages should follow the one given in the examples bellow.
Return the message "No pairs found." when no pair is present.
Return the message "Invalid target sum." if the target is invalid.
Return the message "Invalid number: " if the number in the array is invalid.
For any input format that deviates from the specified format "[1, 2, 3, 4, 5]" "6", the program will return an "Invalid input." error message.
Let's consider the input arr = [1, 2, 3, 4, 5] and the target sum targetSum = 6. When we run the program, the findPairs() function will search for pairs in the array that sum up to targetSum.
*/
package main

import (
    "os"
	"fmt"
    "strings"
    "strconv"
    "math/rand"
    "time"
)

func parseinput(input, target string) ([]int,int,string) {
    sum, err := strconv.Atoi(target)
    if err != nil {
        return nil,0,"Invalid target sum."
    }
    input = strings.TrimSpace(input)
    if len(input) <2 || input[0] != '[' || input[len(input)-1] != ']' {
        return nil,0,"Invalid input."
    }
    input = input[1:len(input)-1] // I made sure that len is at least 2
    temp := strings.Split(input,",")
    nums := make([]int,0,len(temp))
    for _,v := range temp {
        n,err := strconv.Atoi(strings.TrimSpace(v))
        if err != nil {
            return nil,0,"Invalid input " + v
        }
        nums = append(nums,n)
    }
    return nums,sum,""
}

func parseinput1(input, target string) ([]int,int,string) {
    sum, err := strconv.Atoi(target)
    if err != nil {
        return nil,0,"Invalid target sum."
    }
    f := func(r rune) bool {
        return r == '[' || r == ',' || r == ']'
    }
    temp := strings.FieldsFunc(input,f)
    nums := make([]int,0,len(temp))
    for i:=0; i<len(temp); i++ {
        num,err := strconv.Atoi(temp[i])
        if err != nil {
            return nil,0,"Invalid number."
        }
        nums = append(nums,num)
    }
    return nums,sum,""
}

func findpairs(arr []int, sum int) [][]int {
	pairs := make([][]int, 0, 10)
	for i, v := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == sum-v {
				pairs = append(pairs, []int{i, j})
			}
		}
	}
	return pairs
}

func main() {
    if len(os.Args) != 3 {
        os.Exit(1)
    }
    nums, sum, errstr := parseinput(os.Args[1], os.Args[2])
    if len(errstr) != 0 {
        fmt.Println(errstr)
        return
    }
    fmt.Println(findpairs(nums,sum))
}

func generate() ([]int,int) {
    size, mod := 1000000,101
    nums := make([]int,size)
    rand.NewSource(int64(time.Since(time.Now())))
    for i:=0; i < size; i++ {
        nums[i] = rand.Int()%mod
    }
    sum := rand.Int()%mod
    return nums,sum
}
