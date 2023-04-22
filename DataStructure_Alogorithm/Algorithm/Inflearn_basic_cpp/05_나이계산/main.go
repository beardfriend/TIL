package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = file

	var identityNumString, sex string
	var age int
	now := 2019
	fmt.Scanf("%s", &identityNumString)
	splited := strings.Split(identityNumString, "-")

	if splited[0][:1] == "0" {
		hasd, _ := strconv.Atoi("20" + splited[0][:2])
		age = now - hasd + 1
	} else {
		hasd, _ := strconv.Atoi("19" + splited[0][:2])
		age = now - hasd + 1
	}

	if splited[1][:1] == "1" || splited[1][:1] == "3" {
		sex = "M"
	} else {
		sex = "W"
	}

	fmt.Println(age, sex)
}
