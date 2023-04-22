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

	var N, K int

	fmt.Scanf("%d", &N)
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	fmt.Scanf("%d", &K)

	count := 0
	i := 1
	for count < K {
		if i > N {
			i = 1
			continue
		}

		if A[i] == 0 {
			i++
			continue
		}
		A[i] = A[i] - 1
		count++
		i++
	}

	for {
		if i > N {
			i = 1
			continue
		}

		if A[i] == 0 {
			i++
			continue
		} else {
			break
		}
	}
	fmt.Println(i)
}
