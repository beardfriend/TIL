package main

import "fmt"

func main() {
	var N int
	var A [100]int
	fmt.Scanf("%d", &N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	for i := 1; i < N; i++ {
		tmp := A[i]
		index := i - 1
		for index >= 0 && A[index] > tmp {
			A[index+1] = A[index]
			index--
		}
		A[index+1] = tmp
	}
}
