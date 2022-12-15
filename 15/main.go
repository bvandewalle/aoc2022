package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	//parts(input, false)
	parts(input, true)
}

type point struct {
	x int
	y int
}

type interval struct {
	f int
	t int
}

func parts(input []string, part2 bool) {

	grid := map[point]int{}
	beaconGrid := map[point]int{}
	line := 10
	limit := 4000000

	intervalGrid := map[int][]interval{}

	for _, iv := range input {
		iv := strings.TrimPrefix(iv, "Sensor at x=")
		a := strings.Split(iv, ": closest beacon is at x=")
		b := strings.Split(a[0], ", y=")
		c := strings.Split(a[1], ", y=")
		x1, _ := strconv.Atoi(b[0])
		y1, _ := strconv.Atoi(b[1])
		x2, _ := strconv.Atoi(c[0])
		y2, _ := strconv.Atoi(c[1])

		if !part2 {
			if y2 == line {
				beaconGrid[point{x2, y2}] = 1

			}
			calculateBeacon(grid, x1, y1, x2, y2, line)
		}

		calculateIntervals(intervalGrid, x1, y1, x2, y2)
	}

	if !part2 {
		fmt.Println(len(grid) - len(beaconGrid))
		return
	}

	for k, v := range intervalGrid {
		if k < 0 {
			continue
		}
		if k > limit {
			continue
		}
		sort.Slice(v, func(i, j int) bool {
			return v[i].f < v[j].f
		})

		currentMax := v[0].t
		for _, vv := range v {
			if currentMax >= vv.f-1 {
				if vv.t >= currentMax {
					currentMax = vv.t
				}
			} else {
				fmt.Println(limit*(currentMax+1) + k)
				return
			}
		}
	}
}

func calculateIntervals(grid map[int][]interval, x1, y1, x2, y2 int) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	dist := dx + dy

	for a := -dist; a <= dist; a++ {
		xtravel := dist - abs(a)
		if _, e := grid[y1+a]; !e {
			grid[y1+a] = []interval{}
		}
		grid[y1+a] = append(grid[y1+a], interval{x1 - xtravel, x1 + xtravel})
	}
}

func calculateBeacon(grid map[point]int, x1, y1, x2, y2 int, line int) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	dist := dx + dy

	dist20 := abs(line - y1)
	xtravel := dist - dist20
	if dist-dist20 >= 0 {
		fmt.Println(dist, dist20, xtravel)
		for i := -xtravel; i <= xtravel; i++ {
			grid[point{x1 + i, line}] = 1
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
