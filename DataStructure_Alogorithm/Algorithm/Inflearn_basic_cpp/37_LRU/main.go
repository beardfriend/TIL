package main

import "fmt"

func main() {
	var size, n int
	fmt.Scanf("%d %d", &size, &n)
	var now int
	Cache := make([]int, size)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &now)
		hitPosition := -1
		for j := 0; j < size; j++ {
			if Cache[j] == now {
				hitPosition = j
				break
			}
		}

		// miss
		if hitPosition == -1 {
			for j := size - 1; j > 0; j-- {
				Cache[j] = Cache[j-1]
			}
			Cache[0] = now
		} else {
			for j := hitPosition; j > 0; j-- {
				Cache[j] = Cache[j-1]
			}
			Cache[0] = now
		}
	}

	fmt.Println(Cache)
}
