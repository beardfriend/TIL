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
	os.Stdin = file

	var H, W int

	fmt.Scanf("%d %d", &H, &W)

	arr := make([][]int, H)
	for i := 0; i < H; i++ {
		arr2 := make([]int, W)
		arr[i] = arr2
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}

	var maxH, maxW int

	fmt.Scanf("%d %d", &maxH, &maxW)
	var result int
	for i := 0; i < H-maxH; i++ {
		for j := 0; j < W-maxW; j++ {
			sum := 0
			for k := i; k < i+maxH; k++ {
				for l := j; l < j+maxW; l++ {
					sum += arr[k][l]
				}
			}

			if sum > result {
				result = sum
			}

		}
	}

	fmt.Println(result)
}
