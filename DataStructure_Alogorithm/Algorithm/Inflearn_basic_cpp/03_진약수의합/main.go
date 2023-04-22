package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	// 진약수를 구한다 == 진약수 약수에서 자기 자신을 제외한 녀석
	sum := 1
	for i := 1; i < N; i++ {
		if i == N {
			continue
		}
		if N%i == 0 {

			if i == 1 {
				fmt.Print("1")
				continue
			}

			fmt.Printf(" + %d", i)
			sum += i
		}
	}

	fmt.Printf(" = %d", sum)
}
