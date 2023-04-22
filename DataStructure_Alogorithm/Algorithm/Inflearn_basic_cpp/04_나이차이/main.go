package main

import "fmt"

func main() {
	var n, a int
	max := -2470
	min := 2470

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > max {
			max = a
		}
		if a < min {
			min = a
		}
	}
	fmt.Println(max - min)
}
