/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600,851,475,143 ?

Solution:
6857
*/

package main

import "fmt"

func main() {
	num := int64(600851475143)
	fmt.Println(computeMaxPrimeFactor(num))
}

/*
First Attempt

Notes: obviously the most naive solution, pretty bad in both space and time

Performance:
go run 3.go  3028.36s user 1.04s system 100% cpu 50:28.13 total
*/
func computeMaxPrimeFactor(n int64) int64 {
	primeFactors := []int64{}
	for _, factor := range getFactors(n) {
		if isPrime(factor) {
			primeFactors = append(primeFactors, factor)
		}
	}
	return maxNum(primeFactors)
}

func isPrime(n int64) bool {
	for i := int64(2); i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getFactors(n int64) []int64 {
	factors := []int64{1, n}
	for i := int64(2); i <= n/2; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func maxNum(nums []int64) int64 {
	max := int64(1)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

/*
Second Attempt

Notes: not much better than the first one, just a bit more compact and space efficient
*/
func computeMaxPrimeFactor2(n int64) int64 {
	max := int64(0)
	for i := int64(1); i <= n; i++ {
		if n%i == 0 && isPrime(i) && i > max {
			max = i
		}
	}
	return max
}

/*
Third Attempt

Notes: Attempt to use a better factorization algorithm
*/
