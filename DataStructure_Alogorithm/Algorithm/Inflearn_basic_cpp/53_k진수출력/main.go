package main

import (
	"fmt"
	"os"
)

var (
	stack = make([]int, 200)
	top   = -1
)

func push(x int) {
	top++
	stack[top] = x
}

func pop() int {
	val := stack[top]
	top--
	return val
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = file

	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	char := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	for N > 0 {
		push(N % K)
		N = N / K
	}

	for top != -1 {
		fmt.Printf("%s", char[pop()])
	}
}
