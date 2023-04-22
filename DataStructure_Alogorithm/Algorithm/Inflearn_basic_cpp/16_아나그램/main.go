package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var first, second string
	fmt.Fscan(file, &first, &second)
	firstHash := make(map[string]int, 0)
	secondHash := make(map[string]int, 0)

	for i := 0; i < len(first); i++ {
		firstHash[string(first[i])] = firstHash[string(first[i])] + 1
		secondHash[string(second[i])] = secondHash[string(second[i])] + 1
	}

	for i := 0; i < len(first); i++ {
		if firstHash[string(first[i])] != secondHash[string(first[i])] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
