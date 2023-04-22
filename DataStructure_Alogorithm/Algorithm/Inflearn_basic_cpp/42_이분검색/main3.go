package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	os.Stdin = file

	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	sort.Ints(A)

	var lpt, rpt, mid int
	rpt = N - 1
	for lpt <= rpt {
		mid = (rpt + lpt) / 2

		if A[mid] == M {
			break
		}
		if A[mid] > M {
			rpt = mid - 1
		} else {
			lpt = mid + 1
		}
	}
	fmt.Println(mid + 1)
}
