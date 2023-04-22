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

	for i := 1; i < N; i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 {
			if arr[j] > tmp {
				arr[j+1] = arr[j]
			} else {
				break
			}
			j--
		}

		arr[j+1] = tmp
	}

	fmt.Println(arr)
}
