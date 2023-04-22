package main

import (
	"fmt"
	"os"

	stc "github.com/golang-collections/collections/stack"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = f

	var N, M int
	fmt.Scanf("%d", &N)

	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = i
	}
	stack := stc.New()

	result := make([]string, 0)
	j := 1
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d", &M)
		stack.Push(M)
		result = append(result, "P")
		for {
			if stack.Len() == 0 {
				break
			}
			if a[j] == stack.Peek() {
				stack.Pop()
				j++
				result = append(result, "O")
			} else {
				break
			}
		}
	}

	if stack.Len() != 0 {
		fmt.Println("impossible")
	} else {
		fmt.Println(result)
	}
}
