package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	part1And2(input)
}

func part1And2(input []string) {
	count1 := 0
	count2 := 0

	for _, iv := range input {
		p := strings.Split(iv, ",")
		p0 := strings.Split(p[0], "-")
		p1 := strings.Split(p[1], "-")
		p00v, _ := strconv.Atoi(p0[0])
		p01v, _ := strconv.Atoi(p0[1])
		p10v, _ := strconv.Atoi(p1[0])
		p11v, _ := strconv.Atoi(p1[1])
		pairs := [][]int{{p00v, p01v}, {p10v, p11v}}

		if (pairs[0][0] <= pairs[1][0] && pairs[0][1] >= pairs[1][1]) || (pairs[1][0] <= pairs[0][0] && pairs[1][1] >= pairs[0][1]) {
			count1++
		}
		if (pairs[1][0] <= pairs[0][1] && pairs[1][0] >= pairs[0][0]) || (pairs[0][0] <= pairs[1][1] && pairs[0][0] >= pairs[1][0]) {
			count2++
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}
