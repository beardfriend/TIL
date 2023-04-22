package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	var p2, p3, p5, min int
	arr := make([]int, 1501)
	arr[1] = 1
	p2, p3, p5 = 1, 1, 1
	for i := 2; i <= N; i++ {
		if arr[p2]*2 < arr[p3]*3 {
			min = arr[p2] * 2
		} else {
			min = arr[p3] * 3
		}
		if arr[p5]*5 < min {
			min = arr[p5] * 5
		}

		if arr[p5]*5 == min {
			p5++
		}

		if arr[p2]*2 == min {
			p2++
		}

		if arr[p3]*3 == min {
			p3++
		}
		arr[i] = min
	}
	fmt.Println(arr[N])
}
