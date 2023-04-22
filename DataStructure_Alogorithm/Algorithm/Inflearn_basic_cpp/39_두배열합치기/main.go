package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = f

	var n, m int
	var a [100]int

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("%d", &m)
	var Mreq, index int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &Mreq)
		index = -1
		for j := 0; j < n+m; j++ {
			if Mreq >= a[i] {
				a[i] = a[i+1]
			}
		}

		if index == -1 {
			a[m+i] = Mreq
		}

	}

	fmt.Println(a)
}
