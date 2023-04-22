package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, a int
	var primeArr []int
	fmt.Scanf("%d", &N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a)
		rv := reverse2(a)
		if isPrime(rv) {
			primeArr = append(primeArr, rv)
		}
	}

	fmt.Println(primeArr)
}

func reverse2(x int) int {
	var res, tmp int
	for x > 0 {
		tmp = x % 10

		res = res*10 + tmp
		x = x / 10
		fmt.Println(tmp, res, x)
	}
	return res
}

func reverse(x int) int {
	var tmp []string

	for _, v := range strconv.Itoa(x) {
		str := fmt.Sprintf("%c", v)
		if str == "0" {
			continue
		}
		tmp = append(tmp, str)
	}
	var intString string
	for i := len(tmp) - 1; i >= 0; i-- {
		intString += tmp[i]
	}
	value, _ := strconv.Atoi(intString)
	return value
}

// 소수임을 증명하기

// 소수란 자신과 1로 나눠지는 것 외의
// 다른것으로 나눌 수 없는 수

// 2부터  n - 1까지 나눠서
// 나머지 0이 하나도 없으면
// true

func isPrime(val int) bool {
	for i := 2; i < val-1; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}
