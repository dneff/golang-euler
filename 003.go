package main

import (
	"fmt"
	"math"

	"github.com/dneff/golang-euler/euler"
)

/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

func main() {
	fmt.Println("What is the largest prime factor of the number 600851475143?")
	number := int64(600851475143)
	maxFactor := int64(math.Sqrt(float64(number)) + 1)
	factors := []int64{}
	possiblePrimes := euler.FindPrimes(maxFactor)

	for _, x := range possiblePrimes {
		if number%x == 0 {
			factors = append(factors, x)
		}
	}
	fmt.Println("Answer:", factors[len(factors)-1])

}
