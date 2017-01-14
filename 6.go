/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

Solution: 25164150
*/

package main

import "fmt"

func main() {
	diff := computeDifference(100)
	fmt.Printf("Difference: %d\n", diff)
}

func computeDifference(n int) int {
	sum, squares := 0, 0
	for i := 0; i <= n; i++ {
		sum += i
		squares += i * i
	}
	return (sum * sum) - squares
}
