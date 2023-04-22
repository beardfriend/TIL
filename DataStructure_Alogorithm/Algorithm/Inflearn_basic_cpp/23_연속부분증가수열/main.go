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

	// 연속적으로 중가하는 수열의 최대 길이
	// 값이 같을 땐 증가하는 것으로 간주

	var N int
	fmt.Scanf("%d", &N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	cnt := 1
	max := 1
	for i := 1; i < N; i++ {
		if arr[i-1] <= arr[i] {
			cnt++
		} else {
			cnt = 1
		}

		if cnt > max {
			max = cnt
		}
	}

	fmt.Println(max)
}
