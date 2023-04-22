package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStd := os.Stdin
	defer func() { os.Stdin = oldStd }()
	os.Stdin = file
	var N int
	fmt.Scanf("%d", &N)

	var rt, cur, total int
	k := 1
	lt := 1

	for lt != 0 {
		lt = N / (k * 10)
		rt = N % k
		cur = (N / k) % 10

		if cur > 3 {
			total += (lt + 1) * k
		} else if cur < 3 {
			total += lt * k
		} else if cur == 3 {
			total += (lt * k) + (rt + 1)
		}

		k = k * 10

	}

	fmt.Println(total)
}
