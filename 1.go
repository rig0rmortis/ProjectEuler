/*
Project Euler
Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6, 9. The sum of these multiples is 23.

Find the sum of all multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {
	fmt.Println(sumOfMultiples())
}

func sumOfMultiples() int {
	nums := []int{3, 5}
	sum := 0
	for i := 0; i < 1000; i++ {
		if isMultipleOf(i, nums) {
			sum += i
		}
	}
	return sum
}

func isMultipleOf(n int, nums []int) bool {
	for _, num := range nums {
		if n%num == 0 {
			return true
		}
	}
	return false
}
