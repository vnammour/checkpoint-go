/* Tue Aug 12 02:38:53 AM IDT 2025 */
/* By: Victor Nammour */
package main

import (
	"fmt"
	"math"
)

func main() {
	primes := sieve(50)
	fmt.Println(primes)
	prev, next := surroundingPrimes(50)
	fmt.Printf("For 50: prev prime = %d, next prime = %d\n", prev, next)
}

/*
generate all primes up to n
Sieve of Erathosene
Create a list from 2 to n.
Start from the first number p = 2.
Eliminate all multiples of p greater than p itself.
Move to the next remaining number and repeat until p² > n.
The remaining numbers are all prime.
Time complexity
O(n log log n) — extremely fast for generating many primes.
Memory: O(n) — so if n is very large, it can eat a lot of RAM.
*/
func sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	// create a boolean slice to track primality
	isPrime := make([]bool, n + 1)
	for i := range(isPrime) {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for p := 2; p * p < len(isPrime); p++ {
		if isPrime[p] {
			// mark multiples of p as not prime
			for multiple := p * p; multiple < len(isPrime); multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	// collect all primes into a slice
	primes := []int{}
	for i, prime := range isPrime {
		if prime {
			primes = append(primes, i)
		}
	}

	return primes
}

/*
For small numbers (e.g., under 10⁹)
The trial division method with optimizations is usually fastest and simplest:
Check divisibility by 2 and 3 first.
Then check all numbers of the form 6k ± 1 up to √n.
Time complexity: O(√n).
Checks small cases first.
Only tests numbers of the form 6k ± 1 after 3 (because all primes > 3 are of that form).
Time complexity: O(√n).
*/
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n % 2 == 0 || n % 3 == 0 {
		return false
	}
	// check numbers of the form 6k +- 1 up to sqrt(n)
	limit := int(math.Sqrt(float64(n)))
	for i := 5; i <= limit; i += 6 {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
	}

	return true
}

func surroundingPrimes(n int) (prevPrime int, nextPrime int) {
	// Find previous prime
	for i := n -1; i >=2; i-- {
		if isPrime(i) {
			prevPrime = i
			break
		}
	}

	// Find next prime
	for i := n + 1; ; i++ {
		if isPrime(i) {
			nextPrime = i
			break
		}
	}

	return prevPrime, nextPrime
}
