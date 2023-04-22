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
	var N, twoCnt, fiveCnt int
	fmt.Scanf("%d", &N)

	for i := 2; i <= N; i++ {
		tmp := i
		j := 2
		for tmp > 1 {
			if tmp%j == 0 {
				if j == 2 {
					twoCnt++
				}
				if j == 5 {
					fiveCnt++
				}
				tmp = tmp / j
			} else {
				j++
			}
		}
	}

	if twoCnt > fiveCnt {
		fmt.Println(fiveCnt)
	} else {
		fmt.Println(twoCnt)
	}
}
