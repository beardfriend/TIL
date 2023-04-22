package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	os.Stdin = file

	var N, sum int
	fmt.Scanf("%d", &N)

	front := make([]int, N)
	side := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &front[i])
	}

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &side[i])
	}
	var result [][]int

	for i := 0; i < N; i++ {
		f := make([]int, N)
		result = append(result, f)
	}

	for i, v := range front {
		for j, v2 := range side {
			if v > v2 {
				result[i][j] = v2
			} else {
				result[i][j] = v
			}
		}
	}

	for _, v := range result {
		for _, v2 := range v {
			sum += v2
		}
	}
	fmt.Println(sum)
}
