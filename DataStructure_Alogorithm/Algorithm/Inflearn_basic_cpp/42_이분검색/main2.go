package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStd := os.Stdin
	defer func() { os.Stdin = oldStd }()
	os.Stdin = file
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	sort.Ints(arr)

	var lpt, rpt, mid int
	rpt = N - 1
	for lpt <= rpt {
		mid = (rpt + lpt) / 2

		if arr[mid] == M {
			break
		}
		if arr[mid] > M {
			rpt = mid - 1
		} else {
			lpt = mid + 1
		}
	}

	fmt.Println(mid + 1)
}
