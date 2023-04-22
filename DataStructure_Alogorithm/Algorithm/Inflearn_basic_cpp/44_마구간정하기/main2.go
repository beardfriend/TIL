package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	N, M int
	arr  []int
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStd := os.Stdin
	defer func() { os.Stdin = oldStd }()
	os.Stdin = file

	fmt.Scanf("%d %d", &N, &M)

	arr = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	sort.Ints(arr)
	var lpt, rpt, res, mid int

	lpt = 0
	rpt = arr[N-1] - arr[0]

	for lpt <= rpt {
		mid = (lpt + rpt) / 2

		if count(mid) >= M {
			res = mid
			lpt = mid + 1
		} else {
			rpt = mid - 1
		}
	}

	fmt.Println(res)
}

func count(max int) int {
	cnt := 1
	tmp := arr[0]
	for i := 1; i < N; i++ {
		if arr[i]-tmp >= max {
			tmp = arr[i]
			cnt++
		}
	}

	return cnt
}
