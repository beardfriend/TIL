package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = file

	var N int
	fmt.Scanf("%d", &N)
	front := make([]int, N)
	right := make([]int, N)
	up := make([][]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &front[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &right[i])
	}

	for i := 0; i < N; i++ {
		up2 := make([]int, N)
		up[i] = up2
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			up[j][i] = front[i]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if up[i][j] > right[i] {
				up[i][j] = right[i]
			}
		}
	}
	var result int
	for _, v := range up {
		for _, val := range v {
			result += val
		}
	}

	fmt.Println(result)
}
