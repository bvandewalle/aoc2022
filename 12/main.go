package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string

	for scanner.Scan() {
		v := scanner.
			Text()
		input = append(input, v)
	}

	parts(input, false)
	parts(input, true)
}

type point struct {
	x int
	y int
}

func parts(input []string, part2 bool) {
	in := [][]int{}
	m := map[point]int{}
	destination := point{}
	possibleStarts := []point{}

	for i, iv := range input {
		l := []int{}
		for j, c := range iv {
			if c == 'S' {
				c = 'a'
				if !part2 {
					possibleStarts = append(possibleStarts, point{i, j})
				}
			}
			if c == 'E' {
				destination = point{i, j}
				c = 'z'
			}
			if c == 'a' {
				if part2 {
					possibleStarts = append(possibleStarts, point{i, j})
				}
			}
			l = append(l, -(int('a') - int(c)))
		}
		in = append(in, l)
	}

	answer := 10000
	for _, s := range possibleStarts {
		recur(in, m, s, s, 0)
		if m[destination] < answer {
			answer = m[destination]
		}
	}

	fmt.Println(answer)
}

func recur(in [][]int, m map[point]int, previous point, current point, score int) {
	if current.x < 0 || current.x >= len(in) || current.y < 0 || current.y >= len(in[0]) {
		return
	}

	if s, exists := m[current]; exists {
		if s <= score {
			return
		}
	}

	if in[previous.x][previous.y]+1 < in[current.x][current.y] {
		return
	}

	m[current] = score

	recur(in, m, current, point{current.x + 1, current.y}, score+1)
	recur(in, m, current, point{current.x - 1, current.y}, score+1)

	recur(in, m, current, point{current.x, current.y - 1}, score+1)
	recur(in, m, current, point{current.x, current.y + 1}, score+1)
}
