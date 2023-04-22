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
	var N int
	var arr [10]int
	fmt.Scanf("%d", &N)

	for _, v := range strconv.Itoa(N) {
		val, _ := strconv.Atoi(string(v))
		arr[val]++
	}
	var max int
	var val int
	for i, v := range arr {
		if v > max {
			max = v
			val = i
		} else if v == max {
			if v > max {
				max = v
				val = i
			}
		}
	}
	fmt.Println(val)
}
