/* Thu Jan 23 05:50:19 PM IST 2025 */
/* By: Victor Nammour */
package piscine

import (
  "math"
)

func IsPrime(n int) bool {
  if n <= 1 {
    return false
  } else if n == 2 {
    return true
  } else if n % 2 == 0 {
    return false
  }
  var sqrt = int(math.Sqrt(float64(n)))
  for i := 3; i <= sqrt; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func SieveOfEratosthenes(n int) []int {
  // Find all primes up to n
  if n < 2 {
    return []int{}
  }
  primes := make([]bool, n + 1)
  for i := 2; i < n + 1; i++ {
    primes[i] = true
  }
  // Now as a sieve, remove non-primes, going up to the sqrt of n
  for p := 2; p * p <= n; p++ {
    if primes[p] {
      for i := p * 2; i <=n; i += p {
        primes[i] = false
      }
    }
  }
  // Now return all prime numbers <= n
  var primeNumbers []int
  for i := 2; i <= n; i++ {
    if primes[i] == true {
      primeNumbers = append(primeNumbers, i)
    }
  }
  return primeNumbers
}
