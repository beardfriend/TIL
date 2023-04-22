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
		for j := 0; j < N-1-i; j++ {
			if arr[j] > arr[j+1] {
				if arr[j] < 0 && arr[j+1] < 0 {
					continue
				}
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Println(arr)
}
