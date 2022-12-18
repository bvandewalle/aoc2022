package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
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
	//parts(input, true)
}

func parts(input []string, part2 bool) {
	moves := []bool{}

	shapesInt := [][][]int{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{1, 2}, {0, 1}, {1, 1}, {2, 1}, {1, 0}},
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		{{0, 3}, {0, 2}, {0, 1}, {0, 0}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	}

	shapes := [][]image.Point{}
	for _, i := range shapesInt {
		l := []image.Point{}
		for _, j := range i {
			l = append(l, image.Point{j[0], j[1]})
		}
		shapes = append(shapes, l)
	}

	for _, iv := range input {
		for _, v := range iv {
			if v == '<' {
				moves = append(moves, true)
			} else {
				moves = append(moves, false)
			}
		}
	}

	grid := map[image.Point]bool{}
	for i := 0; i < 7; i++ {
		grid[image.Point{i, 0}] = true
	}

	i := 0
	currentHeight := 0

	total := 1000000000000
	repeat := len(moves) * len(shapes)
	fmt.Println(len(moves), len(shapes))

	rest := total % repeat
	fmt.Println(repeat, rest)

	for rock := 0; rock < 18875; rock++ {
		if rock%10 == 0 {
			//fmt.Println("tick")
			//fmt.Println(rock, rock%5, i, i%len(moves))
		}
		if (rock%5 == 0) && (i%len(moves) == 21) {
			fmt.Println(rock, currentHeight)
		}
		shape := []image.Point{}
		for _, v := range shapes[rock%5] {
			shape = append(shape, v.Add(image.Point{2, 4 + currentHeight}))
		}

		for {
			movePushed(grid, shape, moves[i%len(moves)])
			i++
			if !moveDown(grid, shape) {
				break
			}
		}
		for _, v := range shape {
			grid[v] = true
			if v.Y > currentHeight {
				currentHeight = v.Y
			}
		}
	}

	fmt.Println(currentHeight)
}

func movePushed(grid map[image.Point]bool, shape []image.Point, moveDir bool) bool {
	m := image.Point{1, 0}
	if moveDir {
		m = image.Point{-1, 0}
	}
	ret := move(grid, shape, m)
	return ret

}

func moveDown(grid map[image.Point]bool, shape []image.Point) bool {
	m := image.Point{0, -1}
	ret := move(grid, shape, m)
	return ret
}

func move(grid map[image.Point]bool, shape []image.Point, m image.Point) bool {
	for _, v := range shape {
		if _, exists := grid[v.Add(m)]; exists {
			return false
		}
		if v.Add(m).X < 0 || v.Add(m).X > 6 {
			return false
		}
	}
	for i := range shape {
		shape[i] = shape[i].Add(m)
	}
	return true
}
