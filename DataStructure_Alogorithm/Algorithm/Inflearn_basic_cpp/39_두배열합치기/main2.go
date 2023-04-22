package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStd := os.Stdin
	defer func() { os.Stdin = oldStd }()
	os.Stdin = file
	var N, M int
	fmt.Scanf("%d", &N)

	arr1 := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr1[i])
	}

	fmt.Scanf("%d", &M)

	arr2 := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &arr2[i])
	}

	result := make([]int, M+N)

	var p1, p2, p3 int
	for p1 < N && p2 < M {
		if arr1[p1] > arr2[p2] {
			result[p3] = arr2[p2]
			p2++
			p3++
		} else {
			result[p3] = arr1[p1]
			p1++
			p3++
		}
	}

	for p1 < N {
		result[p3] = arr1[p1]
		p1++
		p3++
	}

	for p2 < M {
		result[p3] = arr2[p2]
		p2++
		p3++
	}

	fmt.Println(result)
}
