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
	var cacheN, N int
	fmt.Scanf("%d %d", &cacheN, &N)

	list := make([]int, N)
	c := make([]int, cacheN)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &list[i])
	}

	for _, v := range list {
		isHit := false
		var hitIndex int
		for i := 0; i < cacheN; i++ {
			if v == c[i] {
				isHit = true
				hitIndex = i
				break
			}
		}

		if isHit {
			j := hitIndex - 1
			for j >= 0 {
				c[j+1] = c[j]
				j--
			}
			c[j+1] = v

		} else {
			for i := cacheN - 1; i > 0; i-- {
				c[i] = c[i-1]
			}
			c[0] = v
		}
	}
	fmt.Println(c)
}
