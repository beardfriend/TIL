package main

import "fmt"

func main() {
	var N, tmp int
	fmt.Scanf("%d", &N)
	var A [101]int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	for i := 0; i < N-1; i++ {
		for j := 0; j < N-i-1; j++ {
			if A[j] > A[j+1] {
				tmp = A[j]
				A[j] = A[j+1]
				A[j+1] = tmp
			}
		}
	}
}
