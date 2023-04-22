package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = file

	var str string

	fmt.Scanf("%s", &str)

	cnt := 0
	for _, v := range str {
		char := string(v)
		if char == "(" {
			cnt++
		} else {
			cnt--
		}

		if cnt < 0 {
			break
		}
	}

	if cnt == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
