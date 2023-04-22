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
	os.Stdin = file

	var N, K int
	fmt.Scanf("%d\n", &N)
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d\n", &arr[i])
	}

	fmt.Scanf("%d\n", &K)
	var now int
	j := 0
	for i := 0; i < K; i++ {
		for {
			if j == N {
				j = 0
			}
			if arr[j] == 0 {
				j++
				continue
			} else {
				now = j
				arr[j]--
				fmt.Println(arr)
				j++
				break
			}

		}
	}

	fmt.Println(arr, now+1)
}

//
