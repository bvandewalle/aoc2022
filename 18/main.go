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

type cube struct {
	x int
	y int
	z int
}

func parts(input []string, part2 bool) {
	cubes := map[cube]bool{}

	for _, iv := range input {
		a := strings.Split(iv, ",")
		x, _ := strconv.Atoi(a[0])
		y, _ := strconv.Atoi(a[1])
		z, _ := strconv.Atoi(a[2])
		cubes[cube{x, y, z}] = true
	}

	n := [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

	if part2 {
		toAdd := recurFindIsland(cubes)
		addElems(cubes, toAdd)
	}

	faces := 0
	for k := range cubes {
		for _, v := range n {
			if _, exist := cubes[cube{k.x + v[0], k.y + v[1], k.z + v[2]}]; !exist {
				faces++
			}
		}
	}

	fmt.Println(faces)
}

func recurFindIsland(cubes map[cube]bool) map[cube]bool {
	visited := map[cube]bool{}
	toRemove := map[cube]bool{}

	for i := 0; i <= 21; i++ {
		for j := 0; j <= 21; j++ {
			for k := 0; k <= 21; k++ {
				innerIsland := map[cube]bool{}
				if !recur(cubes, visited, innerIsland, cube{i, j, k}) {
					addElems(toRemove, innerIsland)
				}
			}
		}
	}

	return toRemove
}

// return true if the island is not an inner island (it leaks ot the outside)
// return false of the island is an inner.
func recur(cubes map[cube]bool, visited map[cube]bool, innerIsland map[cube]bool, current cube) bool {
	if _, e := innerIsland[current]; e {
		return false
	}
	if _, e := cubes[current]; e {
		return false
	}

	if _, e := visited[current]; e {
		return true
	}
	if current.x < 0 || current.x > 21 || current.y < 0 || current.y > 21 || current.z < 0 || current.z > 21 {
		return true
	}

	visited[current] = true
	innerIsland[current] = true

	n := [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}
	for _, v := range n {
		ret := recur(cubes, visited, innerIsland, cube{current.x + v[0], current.y + v[1], current.z + v[2]})
		if ret {
			return true
		}
	}

	return false
}

func addElems(toRemove map[cube]bool, island map[cube]bool) {
	for k, v := range island {
		toRemove[k] = v
	}
}
