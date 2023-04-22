package main

import "fmt"

func main() {
	var input string
	var c, h int
	fmt.Scanf("%s", &input)

	var hPosition int
	if string(input[1]) == "H" {
		c = 1
		hPosition = 1
	} else {
		i := 1
		for string(input[i]) != "H" {
			c = c*10 + int(input[i]-48)
			i++
		}
		hPosition = i
	}

	if len(input) == hPosition+1 {
		h = 1
	} else {
		j := 1
		for hPosition+j < len(input) {
			h = h*10 + int(input[hPosition+j]-48)
			j++
		}
	}

	fmt.Println(c*12 + h*1)
}
