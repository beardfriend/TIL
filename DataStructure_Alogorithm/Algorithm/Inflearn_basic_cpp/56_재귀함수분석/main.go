package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	recursive(N)
}

func recursive(N int) {
	if N == 0 {
		return
	}
	recursive(N - 1)
	fmt.Println(N)
}
