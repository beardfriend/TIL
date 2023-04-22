package main

import (
	"fmt"
	"os"
)

var (
	N, M, sum int
	flag      bool
	count     int
	arr       []int
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	os.Stdin = file
	fmt.Scanf("%d %d", &N, &M)
	arr = make([]int, N)

	DFS(0, sum)

	if flag {
		fmt.Println(count)
	} else {
		fmt.Println(-1)
	}
}

func DFS(num, sum int) {
	if num == N {
		if sum == M {
			count++
			flag = true
		}
	} else {
		DFS(num+1, sum+arr[num+1])
		DFS(num+1, sum-arr[num+1])
	}
}
