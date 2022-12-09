package main

import (
	"bufio"
	"fmt"
	"image"
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

	parts(input, 2)
	parts(input, 10)
}

func parts(input []string, ropeLength int) {

	rope := []image.Point{}
	for i := 0; i < ropeLength; i++ {
		rope = append(rope, image.Point{X: 0, Y: 0})
	}

	visited := map[image.Point]bool{}
	moves := map[string]image.Point{"U": image.Point{0, 1}, "D": image.Point{0, -1}, "R": image.Point{1, 0}, "L": image.Point{-1, 0}}

	for _, iv := range input {
		a := strings.Split(iv, " ")
		b, _ := strconv.Atoi(a[1])

		for i := 0; i < b; i++ {
			rope[0] = rope[0].Add(moves[a[0]])

			for j := 1; j < len(rope); j++ {
				dx := rope[j-1].X - rope[j].X
				dy := rope[j-1].Y - rope[j].Y
				if (dx > 1) || (dx < -1) || (dy > 1) || (dy < -1) {
					rope[j] = rope[j].Add(image.Point{sign(dx), sign(dy)})
				}
			}

			visited[rope[ropeLength-1]] = true
		}
	}

	fmt.Println(len(visited))
}

func sign(dx int) int {
	if dx < 0 {
		return -1
	}
	if dx > 0 {
		return 1
	}
	return 0
}
