package main

import (
	"fmt"
	"os"
)

var (
	N  int
	CH [11]int
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	os.Stdin = file

	fmt.Scanf("%d", &N)

	D(1)
}

func D(j int) {
	if j == N+1 {
		for i := 1; i <= N; i++ {
			if CH[i] == 1 {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
	} else {
		CH[j] = 1
		D(j + 1)
		CH[j] = 0
		D(j + 1)
	}
}
