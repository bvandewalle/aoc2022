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

	parts(input, false)
	parts(input, true)
}

type point struct {
	x int
	y int
}

func parts(input []string, part2 bool) {

	grid := map[point]int{}
	maxy := 0
	minx := 1000
	maxx := 0

	for _, iv := range input {
		start := point{}
		end := point{}
		a := strings.Split(iv, " -> ")
		for ai, av := range a {
			b := strings.Split(av, ",")
			x, _ := strconv.Atoi(b[0])
			y, _ := strconv.Atoi(b[1])
			if y > maxy {
				maxy = y
			}
			if x < minx {
				minx = x
			}
			if x > maxx {
				maxx = x
			}

			if ai == 0 {
				start = point{x, y}
				end = start
				continue
			}
			start = end
			end = point{x, y}

			fillGrid(grid, start, end)
		}
	}

	if part2 {
		fillGrid(grid, point{minx - 1000, maxy + 2}, point{minx + 1000, maxy + 2})
	}

	ans := 0
	for {
		if !dropSand(grid, maxy) {
			break
		}
		ans++
		if _, e := grid[point{500, 0}]; e {
			break
		}
	}
	fmt.Println(ans)
}

func fillGrid(m map[point]int, start point, end point) {
	if start.x > end.x {
		start, end = end, start
	}
	dx := end.x - start.x
	for i := 0; i <= dx; i++ {
		m[point{start.x + i, start.y}] = 1
	}

	if start.y > end.y {
		start, end = end, start
	}
	dy := end.y - start.y
	for i := 0; i <= dy; i++ {
		m[point{start.x, start.y + i}] = 1
	}
}

func dropSand(m map[point]int, maxy int) bool {
	current := point{500, 0}
	for {
		if current.y > maxy+2 {
			return false
		}

		if _, e1 := m[point{current.x, current.y + 1}]; e1 {
			if _, e2 := m[point{current.x - 1, current.y + 1}]; e2 {
				if _, e3 := m[point{current.x + 1, current.y + 1}]; e3 {
					m[current] = 2
					return true
				} else {
					current = point{current.x + 1, current.y + 1}
				}
			} else {
				current = point{current.x - 1, current.y + 1}
			}
		} else {
			current = point{current.x, current.y + 1}
		}
	}

}
