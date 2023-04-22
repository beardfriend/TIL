package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()

	os.Stdin = file
	var N, count int
	fmt.Scanf("%d", &N)

	// 총 자리수의 개수

	for i := 1; i <= N; i++ {
		str := strconv.Itoa(i)
		count += len(str)
	}
	fmt.Println(count)
}
