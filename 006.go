package main

// brute force approach. O(n) for 100 values ain't so bad...

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Find the difference between the sum of the squares of the")
	fmt.Println("first one hundred natural numbers and the square of the sum.")

	var sumSquares, squareSum int64

	maxVal := int64(100)

	for x := int64(1); x <= maxVal; x++ {
		sumSquares += int64(math.Pow(float64(x), 2))
		squareSum += x
	}
	squareSum *= squareSum

	fmt.Println("Answer:", squareSum-sumSquares)

}
