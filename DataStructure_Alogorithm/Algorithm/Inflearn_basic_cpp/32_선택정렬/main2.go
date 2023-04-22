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

	var N int
	fmt.Scanf("%d", &N)
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	for i := 0; i < N-1; i++ {
		idx := i
		for j := i + 1; j < N; j++ {
			if arr[i] > arr[idx] {
				idx = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[idx]
		arr[idx] = tmp
	}
	fmt.Println(arr)
}
