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
	var N, count int
	fmt.Fscanf(file, "%d\n", &N)

	var key [101]int
	for i := 1; i < N; i++ {
		fmt.Fscanf(file, "%d", &key[i])
	}
	max := key[N]

	for i := N; i >= 1; i-- {
		if key[i] > max {
			max = key[i]
			count++
		}
	}

	fmt.Println(count)
}
