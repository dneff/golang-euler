package main

import (
	"fmt"

	"github.com/dneff/golang-euler/euler"
)

func main() {

	fmt.Println("What is the 10 001st prime number?")

	// let's estimate the value, knowing that 7.8% of numbers under 1MM are prime...
	maxValue := 10001 / 0.078

	primes := euler.FindPrimes(int64(maxValue))

	// indexes start at 0, so subtract one from array index to get result...
	fmt.Println("Answer:", primes[10000])
}
