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

	var N int

	fmt.Fscanf(file, "%d", &N)

	// 1: 가위 2: 바위 3: 보
	// 1  win(3), lose(2)
	// 2  win(1) lose(3)
	// 3  win(2) lose(1)
	//
	var A, B [101]int
	for i := 0; i < N; i++ {
		fmt.Fscanf(file, "%d", &A[i])
	}

	for i := 0; i < N; i++ {
		fmt.Fscanf(file, "%d", &B[i])
	}

	for i := 0; i < N; i++ {
		if A[i] == B[i] {
			fmt.Println("D")
			continue
		}

		if A[i] == 1 && B[i] == 3 || A[i] == 2 && B[i] == 1 || A[i] == 3 && B[i] == 2 {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}

	}
}
