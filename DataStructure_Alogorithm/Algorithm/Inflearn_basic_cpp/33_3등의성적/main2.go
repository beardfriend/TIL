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
			if arr[idx] < arr[j] {
				idx = j
			}
		}

		tmp := arr[idx]
		arr[idx] = arr[i]
		arr[i] = tmp
	}

	rank := 1
	tmp := arr[0]
	for i := 1; i < N; i++ {
		if arr[i] == tmp {
			continue
		}
		tmp = arr[i]
		rank++
		if rank == 3 {
			fmt.Println(arr[i])
			break
		}
	}
}
