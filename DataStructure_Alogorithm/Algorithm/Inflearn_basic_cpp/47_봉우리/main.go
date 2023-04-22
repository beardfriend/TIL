package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	os.Stdin = file

	var N int

	fmt.Scanf("%d", &N)
	var A [][]int
	A = make([][]int, N+2)

	for i := 0; i <= N+1; i++ {
		A[i] = make([]int, N+2)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			fmt.Scanf("%d", &A[i][j])
		}
	}

	cnt := 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			flag := 0
			if A[i][j] <= A[i+1][j] {
				flag = 1
			}

			if A[i][j] <= A[i-1][j] {
				flag = 1
			}

			if A[i][j] <= A[i][j+1] {
				flag = 1
			}

			if A[i][j] <= A[i][j-1] {
				flag = 1
			}

			if flag == 0 {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
