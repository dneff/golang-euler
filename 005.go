package main

/*
2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all
of the numbers from 1 to 20?
*/

import (
	"fmt"

	"github.com/dneff/golang-euler/euler"
)

func main() {
	fmt.Println("What is the smallest positive number that is evenly divisible by all")
	fmt.Println("of the numbers from 1 to 20?")
	// find factors for 1 to 20, and add unique members to array
	factors := []int64{}

	for x := int64(2); x <= 20; x++ {
		xFactors := euler.FindFactors(int64(x))
		for _, f := range xFactors {
			for j, mf := range factors {
				if mf == f {
					factors = append(factors[:j], factors[j+1:]...)
					break
				}
			}
		}
		factors = append(factors, xFactors...)
	}
	// multiply all factors
	result := int64(1)
	for _, x := range factors {
		result *= x
	}
	fmt.Println("Answer:", result)
}
