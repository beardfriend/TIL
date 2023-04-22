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
	var N, M, count, val, max int
	fmt.Fscanf(file, "%d %d\n", &N, &M)

	for i := 0; i < N; i++ {
		fmt.Fscanf(file, "%d", &val)
		fmt.Println(val)
		if val >= M {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if max == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(max)
	}
}
