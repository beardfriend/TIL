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
	var N int
	fmt.Scanf("%d", &N)

	for i := 1; i <= N; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		fmt.Printf("%d ", count)
	}
}
