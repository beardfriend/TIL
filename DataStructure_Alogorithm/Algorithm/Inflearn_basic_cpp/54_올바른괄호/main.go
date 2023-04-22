package main

import (
	"fmt"
	"os"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = file

	s := make(stack, 0)
	var now string
	fmt.Scanf("%s", &now)

	for _, v := range now {
		str := fmt.Sprintf("%c", v)
		if str == "(" {
			s = s.Push(1)
		} else {
			if len(s) == 0 {
				fmt.Println("NO")
				return
			}
			s, _ = s.Pop()
		}
	}

	if len(s) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
