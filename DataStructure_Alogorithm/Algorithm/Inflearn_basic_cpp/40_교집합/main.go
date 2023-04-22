package main

import (
	"fmt"
	"os"
	"sort"
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

	var result []int

	var p1, p2 int

	sort.Ints(arr1)
	sort.Ints(arr2)

	for p1 < N && p2 < M {
		if arr1[p1] == arr2[p2] {
			result = append(result, arr1[p1])
			p1++
			p2++
		} else if arr1[p1] > arr2[p2] {
			p2++
		} else {
			p1++
		}
	}
	fmt.Println(result)
}
