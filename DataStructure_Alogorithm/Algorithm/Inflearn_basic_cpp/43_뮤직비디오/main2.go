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
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	var lpt, mid, rpt, res int
	rpt = N - 1

	for lpt <= rpt {
		mid = (lpt + rpt) / 2
		var tmp int

		for i := mid; i < N; i++ {
			tmp += arr[i]
		}

		var sum int
		var cnt int
		cnt = 1
		for i := 0; i < N; i++ {
			if sum+arr[i] > tmp {
				sum = arr[i]
				cnt++
			} else {
				sum += arr[i]
			}
		}

		if M >= cnt {
			res = tmp
			lpt = mid + 1
		} else {
			rpt = mid - 1
		}
	}

	fmt.Println(res)
}
