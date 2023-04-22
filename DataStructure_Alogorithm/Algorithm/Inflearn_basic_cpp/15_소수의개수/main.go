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
	var N, count int
	fmt.Scanf("%d", &N)
	for i := 2; i <= N; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
	fmt.Println(count)
}
