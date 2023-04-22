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

	// 같은 점수 동일 석차
	var N int
	fmt.Scanf("%d", &N)
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	result := make([]int, N)
	for i := 0; i < N; i++ {
		result[i]++
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			if arr[i] < arr[j] {
				result[i]++
			}
		}
	}
	fmt.Println(result)
}
