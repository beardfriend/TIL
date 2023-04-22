package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = file

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	var result string
	for _, v := range line {
		if string(v) != " " {
			result += string(v)
		}
	}

	fmt.Println(strings.ToLower(result))
}
