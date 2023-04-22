package main

import "fmt"

func D(N int) {
	if N > 7 {
		return
	}
	D(N * 2)
	D(N*2 + 1)
	fmt.Printf("%d ", N)
}

func main() {
	D(1)
}
