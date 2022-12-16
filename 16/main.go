package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

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
	reachability := map[string][]string{}
	pressure := map[string]int{}

	// Used for a mapping from string to an index for the dijkstra lib
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

		reachability[a[0]] = []string{}
		for _, v := range c {
			reachability[a[0]] = append(reachability[a[0]], v)
		}
		pressure[a[0]] = rate

		mappingInt[a[0]] = count
		count++
	}

	tunnels := []string{}
	// Nifty  trick: Don't consider the destination with a valve of value zero.
	// This reduces the search space by A LOT
	for k, v := range pressure {
		if v != 0 {
			tunnels = append(tunnels, k)
		}
	}

	// PreCalculate Dijkstra as a full matrix. To save time as those results will be used many times
	matrix := calcReachabilityMatrix(reachability, mappingInt)

	maxPressure := 0

	if !part2 {
		// Part1: Simply brute force on the problem space...
		maxPressure = solveRecur(matrix, pressure, 0, 0, 0, "AA", tunnels, 30)
	} else {
		// Part2: Create all the possible combination of the destination and their opposite
		// Then run both brute force in parallels
		// Nifty trick: Only do the combination to half the length of the problem space as the opposite will take care of the mirroring solution
		mu := sync.Mutex{}
		for i := 1; i < len(tunnels)/2; i++ {
			for v := range itertools.CombinationsStr(tunnels, i) {
				maxPressure2a := solveRecur(matrix, pressure, 0, 0, 0, "AA", v, 26)
				maxPressure2b := solveRecur(matrix, pressure, 0, 0, 0, "AA", createOpposite(tunnels, v), 26)
				mu.Lock()
				if maxPressure2a+maxPressure2b > maxPressure {
					maxPressure = maxPressure2a + maxPressure2b
				}
				mu.Unlock()
			}
		}
	}

	fmt.Println(maxPressure)
}

func solveRecur(matrix map[string]map[string]int, pressures map[string]int, currentTime int, currentPressure int, currentFlow int, currentTunnel string, remaining []string, limit int) int {
	// The score if no other valves are being opened before the $limit time
	nScore := currentPressure + (limit-currentTime)*currentFlow
	max := nScore

	for _, v := range remaining {
		distanceAndOpen := matrix[currentTunnel][v] + 1
		if currentTime+distanceAndOpen < limit {
			newTime := currentTime + distanceAndOpen
			newPressure := currentPressure + distanceAndOpen*currentFlow
			newFlow := currentFlow + pressures[v]
			possibleScore := solveRecur(matrix, pressures, newTime, newPressure, newFlow, v, removeFromList(remaining, v), limit)
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

func calcReachabilityMatrix(reachability map[string][]string, mappingInt map[string]int) map[string]map[string]int {
	graph := dijkstra.NewGraph()

	for k := range reachability {
		graph.AddVertex(mappingInt[k])
	}

	for k, v := range reachability {
		for _, l := range v {
			graph.AddArc(mappingInt[k], mappingInt[l], 1)
		}
	}

	matrix := map[string]map[string]int{}
	for k1, v1 := range mappingInt {
		matrix[k1] = map[string]int{}
		for k2, v2 := range mappingInt {
			best, _ := graph.Shortest(v1, v2)
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
