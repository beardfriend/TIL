package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	os.Stdin = file

	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = 1
	}

	outCnt := 0
	count := 1
	i := 1
	for outCnt < N-1 {
		if i > N {
			i = 1
			continue
		}
		if A[i] == 0 {
			i++
			continue
		}

		if count == 3 {
			fmt.Println(i)
			A[i] = 0
			count = 1
			outCnt++
			continue
		}
		i++
		count++
	}

	fmt.Println(A)
}
