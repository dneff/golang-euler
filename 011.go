package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dneff/golang-euler/euler"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func productArray(arr []int64) int64 {
	product := int64(1)
	for _, x := range arr {
		product *= x
	}
	return product
}

func findMax(g [][]int64) int64 {
	result := int64(0)
	// check rows
	for _, r := range g {
		result = euler.More(result, int64(productArray(r)))
	}
	// check columns
	for c := 0; c < len(g[0]); c++ {
		v := []int64{g[0][c], g[1][c], g[2][c], g[3][c]}
		result = euler.More(result, productArray(v))
	}
	// check diagonals
	d1 := []int64{g[0][0], g[1][1], g[2][2], g[3][3]}
	d2 := []int64{g[0][3], g[1][2], g[2][1], g[3][0]}
	result = euler.More(result, productArray(d1))
	result = euler.More(result, productArray(d2))

	return result
}

func main() {
	fmt.Println("What is the greatest product of four adjacent numbers")
	fmt.Println("in the same direction (up, down, left, right, or")
	fmt.Println("diagonally) in the 20×20 grid?")

	f, err := os.Open("./reference/011.txt")
	check(err)
	grid := [][]int64{}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		r := []int64{}
		for _, c := range strings.Fields(s.Text()) {
			x, err := strconv.Atoi(c)
			check(err)
			r = append(r, int64(x))
		}
		grid = append(grid, r)
	}
	maxValue := int64(0)
	for r := 0; r < len(grid)-3; r++ {
		for c := 0; c < len(grid[r])-3; c++ {
			g := [][]int64{
				grid[r][c : c+4],
				grid[r+1][c : c+4],
				grid[r+2][c : c+4],
				grid[r+3][c : c+4]}

			maxValue = euler.More(maxValue, findMax(g))
		}
	}

	fmt.Println("Answer:", maxValue)
}
