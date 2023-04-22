package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = file

	// N개 정수
	// 인접 한 수의 차이가 1 ~ N-1 모두 가지면 == jolly jumper
	// ex 1 4 2 3
	// 4개 3 2 1
	// 졸리 점퍼임을 판단하는 프로그램

	var N int
	fmt.Scanf("%d", &N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	arr2 := make([]int, N)
	for i := 1; i < N; i++ {
		cha := math.Abs(float64(arr[i-1] - arr[i]))
		arr2[int(cha)] = 1
	}

	for i := 1; i <= N-1; i++ {
		if arr2[i] == 0 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
