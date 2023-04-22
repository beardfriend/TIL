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

	for i := 1; i <= N; i++ {
		var val int
		fmt.Scanf("%d", &val)
		s := reverse(val)
		if isPrime(s) {
			fmt.Printf("%d ", s)
		}
	}
}

func reverse(v int) int {
	var result int
	strInt := strconv.Itoa(v)
	tmp := make([]string, 0)
	for _, v := range strInt {
		tmp = append(tmp, string(v))
	}
	var tmpResult string
	for i := len(tmp) - 1; i >= 0; i-- {
		tmpResult += tmp[i]
	}
	result, _ = strconv.Atoi(tmpResult)
	return result
}

func isPrime(v int) bool {
	for i := 2; i < v; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}
