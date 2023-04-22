package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var tmp int
			if a[j] > a[i] {
				tmp = a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
	}

	var cnt int
	for i := 1; i < n; i++ {
		if a[i-1] != a[i] {
			cnt++
		}

		if cnt == 2 {
			fmt.Printf("%d", a[i])
			break
		}
	}
}
