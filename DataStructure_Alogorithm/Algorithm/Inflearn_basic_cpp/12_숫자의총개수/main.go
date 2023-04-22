package main

import (
	"fmt"
	"math"
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

	fmt.Scanf("%d", &N)
	// 방법을 생각해본다
	NtoString := strconv.Itoa(N)
	length := len(NtoString)

	var total int
	for i := 1; i < length; i++ {
		total += 9 * int(math.Pow(10, float64(i-1))) * i
	}
	total += (N - int(math.Pow(10, float64(length-1))) + 1) * length

	fmt.Println(total)
}
