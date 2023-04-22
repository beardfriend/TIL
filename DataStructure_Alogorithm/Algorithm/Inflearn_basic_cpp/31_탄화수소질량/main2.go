package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStd := os.Stdin
	defer func() { os.Stdin = oldStd }()
	os.Stdin = file

	var hydroCarbon string
	fmt.Scanf("%s", &hydroCarbon)

	var total int
	var hydroNum, carbonNum, character string
	for _, v := range hydroCarbon {
		if v >= 48 && v <= 57 {
			if character == "C" {
				carbonNum += fmt.Sprintf("%c", v)
			} else {
				hydroNum += fmt.Sprintf("%c", v)
			}
		} else {
			character = fmt.Sprintf("%c", v)
		}
	}
	if carbonNum == "" {
		total += 12
	} else {
		i, _ := strconv.Atoi(carbonNum)
		total += i * 12
	}

	if hydroNum == "" {
		total += 1
	} else {
		i, _ := strconv.Atoi(hydroNum)
		total += i
	}
	fmt.Println(total)
}
