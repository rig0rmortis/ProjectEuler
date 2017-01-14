/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	palindrome := generateLargestPalindrome()
	if palindrome == 0 {
		fmt.Println("No Palindrome found!")
	} else {
		fmt.Printf("Largest Palindrome: %d\n", palindrome)
	}
}

/*
Notes: first attempt for a naive solution. Start from 100*100, iterate up to
       999*999... actually, we can do better by starting from 999*999 and
       moving backwards. First palindrome should be largest. Pick 2 numbers,
       j, k both equal to 999. Decrease k till you hit 100, then decrease j
       by 1, and iterate through k from 998 down to 100.
*/
func checkPalindrome(n int) bool {
	stringified := strconv.Itoa(n)
	reversed := Reverse(stringified)
	if strings.Compare(stringified, reversed) == 0 {
		return true
	}
	return false
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func generateLargestPalindrome() int {
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			num := i * j
			if checkPalindrome(num) {
				return num
			}
		}
	}
	return 0
}
