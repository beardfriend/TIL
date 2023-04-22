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

	// 연속된 날짜 온도의 합이 최대인 값을 구하기
	// 입력
	// 날짜수, 합을 구하는 연속적 날짜
	// 온도

	// 입력
	var N, K int
	fmt.Scanf("%d %d", &N, &K)
	temp := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &temp[i])
	}
	// 합을 구한다
	max := -218766660
	for i := 0; i < N; i++ {
		var sum int
		if N-i > K {
			for j := i; j < K+i; j++ {
				sum += temp[j]
			}
		}
		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)
	// 0 ~ K
	// 1 ~ K+1
	// 2 ~ K+2
	// 최대값 측정
}
