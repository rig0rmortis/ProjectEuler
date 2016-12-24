/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600,851,475,143 ?
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(computeMaxPrimeFactor(600851475143))
}

func computeMaxPrimeFactor(n int) int {
	primeFactors := []int{}
	for _, factor := range getFactors(n) {
		if isPrime(factor) {
			primeFactors = append(primeFactors, factor)
		}
	}
	return sort.IntSlice(primeFactors)[len(primeFactors)-1]
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getFactors(n int) []int {
	factors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
