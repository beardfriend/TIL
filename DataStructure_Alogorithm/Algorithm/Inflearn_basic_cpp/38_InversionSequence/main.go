package main

import "fmt"

func main() {
	var N int

	fmt.Scanf("%d", &N)
	os := make([]int, N)
	is := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &is[i])
	}

	var pos int
	for i := N; i >= 1; i-- {
		pos = i
		for j := 1; j <= is[i]; j++ {
			os[pos] = os[pos+1]
			pos++
		}
		os[pos] = i
	}
}
