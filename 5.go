/*
2520 is the smallest number that can be divided by each of the numbers from
1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?

Answer: 232792560

go run 5.go  4.20s user 0.05s system 100% cpu 4.239 total
*/

package main

import (
	"fmt"
)

func main() {
	num := tinyNum()
	fmt.Printf("Smallest evenly divisible number: %d\n", num)
}

func tinyNum() int {
	num := 100
	for {
		if testDivisibility(num) {
			return num
		} else {
			num = num + 1
		}
	}
}

func testDivisibility(n int) bool {
	for i := 2; i <= 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
