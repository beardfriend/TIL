package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	os.Stdin = file
	var N int
	fmt.Scanf("%d", &N)
	recursive(N)
}

func recursive(N int) {
	if N == 0 {
		return
	}
	recursive(N / 2)
	fmt.Print(N % 2)
}
