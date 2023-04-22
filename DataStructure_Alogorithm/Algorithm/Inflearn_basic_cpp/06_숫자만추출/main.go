package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = file

	var a string
	fmt.Scanf("%s", &a)
	var result string
	for _, v := range a {
		if v > 48 && v <= 57 {
			result += string(v)
		}
	}

	res, _ := strconv.Atoi(result)
	fmt.Println(res)

	i := 1
	count := 0
	for i <= res {
		if res%i == 0 {
			count++
		}
		i++
	}

	fmt.Println(count)
}
