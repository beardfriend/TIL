package main

import "fmt"

func main() {
	var A [10]int
	var B [10]int
	var aCount int
	var bCount int
	var lastWinner int
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &A[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &B[i])
	}

	for i := 0; i < 10; i++ {
		if A[i] > B[i] {
			aCount += 3
			lastWinner = 1
		}

		if A[i] < B[i] {
			bCount += 3
			lastWinner = 2
		}

		if A[i] == B[i] {
			aCount++
			bCount++
		}
	}

	fmt.Println(aCount, bCount)
	if aCount > bCount {
		fmt.Println("A")
	}

	if aCount < bCount {
		fmt.Println("B")
	}

	if aCount == bCount {
		if lastWinner == 1 {
			fmt.Println("A")
		} else if lastWinner == 2 {
			fmt.Println("B")
		} else if lastWinner == 0 {
			fmt.Println("D")
		}
	}
}
