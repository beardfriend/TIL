package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	old := os.Stdin
	defer func() { os.Stdin = old }()

	os.Stdin = file
	var N int
	fmt.Scanf("%d", &N)

	var max int
	var maxOrigianl int
	for i := 1; i <= N; i++ {
		var num int
		fmt.Scanf("%d", &num)

		str := strconv.Itoa(num)

		var acc int
		for _, v := range str {
			a, _ := strconv.Atoi(string(v))
			acc += a
		}

		if acc > max {
			max = acc
			maxOrigianl = num
		} else if acc == max {
			if num > maxOrigianl {
				maxOrigianl = num
			}
		}
	}
	fmt.Println(maxOrigianl)
}
