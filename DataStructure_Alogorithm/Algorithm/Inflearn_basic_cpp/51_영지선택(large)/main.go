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

	arr := make([][]int, H+2)
	dp := make([][]int, H+2)
	for i := 0; i < H+2; i++ {
		arr2 := make([]int, W+2)
		arr[i] = arr2
		dp[i] = arr2

	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			fmt.Scanf("%d", &arr[i][j])
			dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1] + arr[i][j]
		}
	}
	var maxH, maxW int
	fmt.Scanf("%d %d", &maxH, &maxW)

	// program

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
		}
	}

	var max int
	for i := maxH; i <= H; i++ {
		for j := maxW; j <= W; j++ {
			sum := dp[i][j] - dp[i-maxH][j] - dp[i][j-maxW] + dp[i-maxH][j-maxW]
			if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}
