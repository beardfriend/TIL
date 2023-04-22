package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = file

	arr := make([][]int, 9)

	for i := 0; i < 9; i++ {
		arr[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}
	var result int
	for i, v := range arr {
		var sum, before int
		var avg float64

		sort.Ints(arr[i])
		for _, val := range v {
			sum += val
		}
		avg = math.Abs(float64(sum) / float64(9))
		for _, val := range v {

			cha := int(math.Abs(avg - float64(val)))
			if cha < before {
				result = val
			}
			before = cha
		}

		fmt.Println(avg, result)
	}
}
