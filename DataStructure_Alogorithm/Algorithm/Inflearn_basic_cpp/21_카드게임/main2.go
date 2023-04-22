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

	// 0 ~ 9
	// 섞는다
	// 숫자가 큰 사람이 승리
	// 승리 : 3
	// 무승부 : 1
	// 패자 : x
	// 승점이 같으면 제일 마지막에 이긴 사람이 승리
	// 모든 게임이 무승부면 드로우

	// 입력
	var A, B [10]int
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &A[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &B[i])
	}

	var lastWiner int // 0: no one / 1: A / 2: B
	var aCnt, bCnt int
	// 비교
	for i := 0; i < 10; i++ {
		if A[i] > B[i] {
			aCnt = aCnt + 3
			lastWiner = 1
		} else if A[i] < B[i] {
			bCnt = bCnt + 3
			lastWiner = 2
		} else if A[i] == B[i] {
			aCnt++
			bCnt++
		}
	}

	fmt.Println(aCnt, bCnt)
	if aCnt > bCnt {
		fmt.Println("A")
	}

	if aCnt < bCnt {
		fmt.Println("B")
	}

	if aCnt == bCnt {
		if lastWiner == 1 {
			fmt.Println("A")
		} else if lastWiner == 2 {
			fmt.Println("B")
		} else if lastWiner == 0 {
			fmt.Println("D")
		}
	}
	// 승부결과 진단
}
