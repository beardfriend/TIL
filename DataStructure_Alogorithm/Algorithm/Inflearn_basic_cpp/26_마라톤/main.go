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

	// 현재 순서 - 먼저 있는 숫자가 자기보다 작을 경우를 모두 뺀 값
	var N int
	fmt.Scanf("%d", &N)

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	result := make([]int, N)
	for i := 0; i < N; i++ {
		cnt := 1
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				cnt++
			}
		}
		result[i] = cnt
	}
	fmt.Println(result)
}
