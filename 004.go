package main

/*
A palindromic number reads the same both
ways. The largest palindrome made from the
product of two 2-digit numbers is
9009 = 91 × 99.

Find the largest palindrome made from the
product of two 3-digit numbers.
*/

import (
	"fmt"

	"github.com/dneff/golang-euler/euler"
)

func main() {
	// start with maximum 3-digit numbers and work down
	maxValue := int64(999)
	result := int64(0)
	fmt.Println("Find the largest palindrome made from the")
	fmt.Println("product of two 3-digit numbers.")

palindromeCheck:

	for decr := int64(1); decr < maxValue; decr++ {
		i := maxValue - decr
		j := maxValue
		for i <= j {
			if euler.IsPalindrome(i * j) {
				result = i * j
				break palindromeCheck
			}
			i++
			j--
		}
	}
	fmt.Println("Answer:", result)

}
