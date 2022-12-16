package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
	"github.com/ernestosuarez/itertools"
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

func parts(input []string, part2 bool) {
	gr := map[string][]string{}
	press := map[string]int{}
	mappingInt := map[string]int{}
	count := 0

	for _, iv := range input {
		iv := strings.TrimPrefix(iv, "Valve ")
		a := strings.Split(iv, " has flow rate=")
		b := strings.Split(a[1], "; tunnels lead to valve")
		if len(b) == 1 {
			b = strings.Split(a[1], "; tunnel leads to valve")
		}
		rate, _ := strconv.Atoi(b[0])

		b[1] = strings.TrimPrefix(b[1], "s ")
		b[1] = strings.TrimPrefix(b[1], " ")
		c := strings.Split(b[1], ", ")

		gr[a[0]] = []string{}
		for _, v := range c {
			gr[a[0]] = append(gr[a[0]], v)
		}
		press[a[0]] = rate

		mappingInt[a[0]] = count
		count++
	}

	graph := dijkstra.NewGraph()
	tunnels := []string{}
	for k := range press {
		if k != "AA" {
			if press[k] != 0 {
				tunnels = append(tunnels, k)
			}
		}
		graph.AddVertex(mappingInt[k])
	}

	for k, v := range gr {
		for _, l := range v {
			graph.AddArc(mappingInt[k], mappingInt[l], 1)
		}
	}

	matrix := calcReachabilityMatrix(graph, mappingInt)

	if !part2 {
		maxPressure := solveRecur(matrix, press, mappingInt, 0, 0, 0, "AA", tunnels, 30)
		fmt.Println(maxPressure)
	} else {
		maxPressure := 0
		for i := 3; i < len(tunnels)/2; i++ {
			for v := range itertools.CombinationsStr(tunnels, i) {
				maxPressure2a := solveRecur(matrix, press, mappingInt, 0, 0, 0, "AA", v, 26)
				maxPressure2b := solveRecur(matrix, press, mappingInt, 0, 0, 0, "AA", createOpposite(tunnels, v), 26)
				if maxPressure2a+maxPressure2b > maxPressure {
					maxPressure = maxPressure2a + maxPressure2b
				}
			}
		}
		fmt.Println(maxPressure)
	}
}

func solveRecur(matrix map[string]map[string]int, pressures map[string]int, mappingInt map[string]int, currentTime int, currentPressure int, currentFlow int, currentTunnel string, remaining []string, limit int) int {
	nScore := currentPressure + (limit-currentTime)*currentFlow
	max := nScore

	for _, v := range remaining {
		distanceAndOpen := matrix[currentTunnel][v] + 1
		if currentTime+distanceAndOpen < limit {
			newTime := currentTime + distanceAndOpen
			newPressure := currentPressure + distanceAndOpen*currentFlow
			newFlow := currentFlow + pressures[v]
			possibleScore := solveRecur(matrix, pressures, mappingInt, newTime, newPressure, newFlow, v, removeFromList(remaining, v), limit)
			if possibleScore > max {
				max = possibleScore
			}
		}
	}

	return max
}

func removeFromList(in []string, v string) []string {
	new := []string{}
	for _, i := range in {
		if i != v {
			new = append(new, i)
		}
	}
	return new
}

func calcReachabilityMatrix(gr *dijkstra.Graph, mappingInt map[string]int) map[string]map[string]int {
	matrix := map[string]map[string]int{}
	for k1, v1 := range mappingInt {
		matrix[k1] = map[string]int{}
		for k2, v2 := range mappingInt {
			best, _ := gr.Shortest(v1, v2)
			matrix[k1][k2] = int(best.Distance)
		}
	}

	return matrix
}

func createOpposite(all []string, partial []string) []string {
	new := []string{}

outer:
	for _, v := range all {
		for _, w := range partial {
			if v == w {
				continue outer
			}
		}
		new = append(new, v)
	}

	return new
}
