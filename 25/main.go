package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string

	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	parts(input)
}

func parts(input []string) {

	snafu := []int{}

	for _, iv := range input {
		for i := 0; i < len(iv); i++ {
			if i+1 > len(snafu) {
				snafu = append(snafu, 0)
			}
			c := iv[len(iv)-i-1]
			if c == '=' {
				snafu[i] -= 2
			} else if c == '-' {
				snafu[i] -= 1
			} else {
				x, _ := strconv.Atoi(string(c))
				snafu[i] += x
			}
		}
	}

	for i := range snafu {
		v := snafu[i]
		carry := 0
		for v >= 3 {
			carry += 1
			v -= 5
		}
		for v <= -3 {
			carry -= 1
			v += 5
		}
		snafu[i] = v
		if carry != 0 {
			snafu[i+1] += carry
		}
	}

	for i := len(snafu) - 1; i >= 0; i-- {
		v := snafu[i]
		if v == -2 {
			fmt.Print("=")
		} else if v == -1 {
			fmt.Print("-")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}
