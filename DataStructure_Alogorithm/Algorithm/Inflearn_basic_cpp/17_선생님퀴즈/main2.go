package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()
	var N int
	fmt.Fscanf(file, "%d\n", &N)
	for i := 0; i < N; i++ {
		var a, total int
		fmt.Fscanf(file, "%d %d", &a, &total)
		if total == (a+1)*(a/2) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
