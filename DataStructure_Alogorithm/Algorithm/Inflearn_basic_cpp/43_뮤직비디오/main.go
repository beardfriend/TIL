package main

import (
	"fmt"
	"os"
)

func f() {
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

	var max, res, lpt, rpt, mid int
	for _, v := range arr {
		max += v
	}
	lpt = 1
	rpt = max
	for lpt <= rpt {
		mid = (lpt + rpt) / 2
		if count(arr, mid) <= M {
			res = mid
			rpt = mid - 1
		} else {
			lpt = mid + 1
		}
	}

	fmt.Println(res)
}

func count(arr []int, max int) int {
	var val, cnt int
	cnt = 1

	for _, v := range arr {
		if val+v > max {
			cnt++
			val = v
		} else {
			val += v
		}
	}

	return cnt
}
