package main

import (
	"fmt"
)

/*
If we list all the natural numbers below 10 that are multiples
of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

func main() {
	fmt.Println("Find the sum of all the multiples of 3 or 5 below 1000.")
	maxValue := int32(1000)
	var result int32

	for x := int32(0); x < maxValue; x++ {
		if x%3 == 0 || x%5 == 0 {
			result += x
		}
	}
	fmt.Println("Answer: ", result)
}
