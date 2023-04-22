package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	open, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer open.Close()
	os.Stdin = open

	var N int

	fmt.Scanf("%d", &N)

	var ages [100]int

	// 값 읽기
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &ages[i])
	}

	var max, tmp int
	for i, v := range ages {
		for j := i + 1; j < N; j++ {
			tmp = Abs(v - ages[j])
			if tmp > max {
				max = tmp
			}
		}
	}

	fmt.Println(max)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
