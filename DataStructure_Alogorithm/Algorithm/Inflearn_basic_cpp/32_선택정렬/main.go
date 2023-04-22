package main

import "fmt"

func main() {
	var n, tmp, idx int
	var a [101]int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			idx = i
			if a[i] > a[j] {
				idx = j
			}
		}
		tmp = a[i]
		a[i] = a[idx]
		a[idx] = tmp
	}

	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}
