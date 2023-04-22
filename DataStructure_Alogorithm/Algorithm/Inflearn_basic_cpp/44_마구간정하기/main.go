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

	var lt, rt, mid, res int
	lt = 0
	rt = arr[N-1] - arr[0]

	for lt <= rt {
		mid = (lt + rt) / 2

		if count(mid, arr) >= M {
			fmt.Println(mid)
			res = mid
			lt = mid + 1
		} else {
			rt = mid - 1
		}
	}
	fmt.Println(res)
}

func count(res int, arr []int) int {
	var cnt, tmp int
	tmp = arr[0]
	cnt = 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-tmp >= res {
			tmp = arr[i]
			cnt++
		}
	}
	return cnt
}
