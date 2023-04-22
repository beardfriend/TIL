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
	result := make([]int, N+1)
	for i := 2; i <= N; i++ {
		tmp := i
		j := 2
		for tmp > 1 {
			if tmp%j == 0 {
				tmp = tmp / j
				result[j]++
			} else {
				j++
			}
		}
	}
	fmt.Printf("%d! = ", N)
	for _, v := range result {
		if v == 0 {
			continue
		}
		fmt.Printf("%d ", v)
	}
}
