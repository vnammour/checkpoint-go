/* Fri Mar 21 07:28:29 PM IST 2025 */
/* By: Victor Nammour */
/*
difficulty: 10
level: 9
allowed: strconv.Atoi, os.Args, fmt.*, builtin
Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').
Factors must be displayed in ascending order and separated by *.
If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(0)
	}
	fmt.Print(fprime(n))
}

func temp(n int) []int {
	if n <= 0 {
		return nil
	}
	primes := make([]int, n/2+1)
	for i := 2; i < len(primes); i++ {
		for j := i; i*j < len(primes); j++ {
			primes[i*j] = 1 // 1 here means that it's not a prime
		}
	}
	factors := make([]int, 0, len(primes))
	for i := 2; i < len(primes); i++ {
		if primes[i] == 0 && n % i == 0 {
			factors = append(factors, i)
            for j := i*i; j < n/2 && n % j == 0; j = j * i {
			    factors = append(factors, i)
            }
		}
	}
	return factors
}

func fprime(n int) string {
	if n <= 0 {
		return ""
	}
	primes := make([]int, n/2+1)
	for i := 2; i < len(primes); i++ {
		for j := i; i*j < len(primes); j++ {
			primes[i*j] = 1 // 1 here means that it's not a prime
		}
	}
	factors := make([]int, 0, len(primes))
	for i := 2; i < len(primes); i++ {
		if primes[i] == 0 && n % i == 0 {
			factors = append(factors, i)
            for j := i*i; j < n/2 && n % j == 0; j = j * i {
			    factors = append(factors, i)
            }
		}
	}
    out := make([]byte,0,len(factors)*2)
    for i := 0; i < len(factors); i++ {
        out = append(out,itoa(factors[i])...)
        out = append(out,'*')
    }
    if len(out) != 0 {
        out = out[:len(out)-1]
    }
    return string(out)
}

func itoa(n int) []byte {
	if n == 0 {
		return []byte{'0'}
	}
	sign := n
	if sign < 0 {
		n *= -1
	}
	bits := make([]byte, 0, 10) // again, need better approximation for capacity
	for n != 0 {
		bits = append(bits, byte(n%10)+'0')
		n /= 10
	}
	if sign < 0 {
		bits = append(bits, '-')
	}
	return reverse(bits)
}

func reverse(bits []byte) []byte {
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
	return bits
}
