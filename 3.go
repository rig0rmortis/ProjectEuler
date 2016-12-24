/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600,851,475,143 ?
*/

package main

import "fmt"

func main() {
	num := int64(600851475143)
	fmt.Println(computeMaxPrimeFactor(num))
}

func computeMaxPrimeFactor(n int64) int64 {
	primeFactors := []int64{}
	for _, factor := range getFactors(n) {
		fmt.Printf("Working on factor: %d", factor)
		if isPrime(factor) {
			primeFactors = append(primeFactors, factor)
		}
	}
	return maxNum(primeFactors)
}

func isPrime(n int64) bool {
	fmt.Printf("Checking if %d is prime...", n)
	for i := int64(2); i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getFactors(n int64) []int64 {
	fmt.Printf("Getting factors for %d", n)
	factors := []int64{1, n}
	for i := int64(2); i <= n/2; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func maxNum(nums []int64) int64 {
	fmt.Printf("Getting the max number")
	max := int64(1)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
