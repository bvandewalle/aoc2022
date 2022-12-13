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

	parts(input)
}

func parts(input []string) {
	cycle := 1
	x := 1
	score := 0

	draw(cycle, x)
	for _, iv := range input {
		a := strings.Split(iv, " ")
		cycle++
		draw(cycle-1, x)
		if a[0] == "addx" {
			b, _ := strconv.Atoi(a[1])
			if (cycle-20)%40 == 0 {
				score += x * cycle
			}
			x += b
			cycle++
			draw(cycle-1, x)
		}
		if (cycle-20)%40 == 0 {
			score += x * cycle
		}
	}
	fmt.Println(score)
}

func draw(cycle, x int) {
	if cycle%40 == 0 {
		fmt.Println()
	}
	cycle %= 40
	if x-1 == cycle || x == cycle || x+1 == cycle {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}
