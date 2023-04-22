package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	sort.Ints(A)

	lpt := 0
	rpt := len(A) - 1
	var mid int
	for lpt <= rpt {
		mid = (lpt + rpt) / 2
		if A[mid] == M {
			break
		}
		if A[mid] > M {
			rpt = mid - 1
		} else {
			lpt = mid + 1
		}
	}

	fmt.Println(mid)
}
