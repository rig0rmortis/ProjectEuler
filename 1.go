/*
Project Euler
Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6, 9. The sum of these multiples is 23.

Find the sum of all multiples of 3 or 5 below 1000.
*/

package main

func main() {

}

func calculateSum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func isMultipleOf(n int, nums []int) bool {
	for _, num := range nums {
		switch n % num {
		case 0:
			return true
		default:
			return false
		}
	}
}
