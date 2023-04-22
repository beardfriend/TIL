package main

import "fmt"

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	for i := A; i <= B; i++ {
		if i == A {
			fmt.Printf("%d", i)
			continue
		}

		fmt.Printf(" + %d", i)
	}

	fmt.Printf(" = %d", A+B)
}
