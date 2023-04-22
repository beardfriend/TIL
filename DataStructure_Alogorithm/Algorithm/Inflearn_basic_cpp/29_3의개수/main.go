package main

import "fmt"

func main() {
	var N, tmp, count int

	fmt.Scanf("%d", &N)

	for i := 1; i <= N; i++ {
		tmp = i
		for tmp > 0 {
			if tmp%10 == 3 {
				count++
			}
			tmp = tmp / 10
		}
	}

	fmt.Println(count)
}
