package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = file

	var N, i, num uint64
	max := uint64(0)
	min := uint64(255)

	fmt.Scanf("%d", &N)

	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &num)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	fmt.Println(max - min)
}
