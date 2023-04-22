package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var N, answer int
		fmt.Scanf("%d %d", &N, &answer)

		if (1+N)*(N/2) == answer {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
