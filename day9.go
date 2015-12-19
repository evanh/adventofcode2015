package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var EDGES = map[string]map[string]int{}
var CITIES = []string{}

func Contains(a []string, v string) bool {
	for i := range a {
		if a[i] == v {
			return true
		}
	}
	return false
}

func Permutations(output chan []string) {
	// This isn't the cleanest, but have separate functions
	// so we know when we're done with all the permutations.
	for i := range CITIES {
		InnerPermutations(output, []string{CITIES[i]})
	}
	// Send kill signal
	output <- []string{}
}

func InnerPermutations(output chan []string, path []string) {
	if len(path) == len(CITIES) {
		if path[0] == "Tristram" && path[1] == "Tambi" && path[2] == "Snowdin" && path[3] == "AlphaCentauri" && path[4] == "Faerun" && path[5] == "Straylight" && path[6] == "Norrath" {
			fmt.Println("SOLUTION??")
		}
		output <- path
		return
	}

	for i := range CITIES {
		if !Contains(path, CITIES[i]) {
			if _, ok := EDGES[path[len(path)-1]][CITIES[i]]; ok {
				InnerPermutations(output, append(path, CITIES[i]))
			}
		}
	}
}

func main() {
	f, _ := os.Open("day9.input")
	data, _ := ioutil.ReadAll(f)
	inputs := strings.Split(string(data), "\n")

	// Initialize graph and cities to loop over
	for _, input := range inputs {
		parts := strings.Split(input, " to ")
		city1 := parts[0]
		parts = strings.Split(parts[1], " = ")
		city2 := parts[0]
		dist, _ := strconv.Atoi(parts[1])

		if v, ok := EDGES[city1]; ok {
			v[city2] = dist
		} else {
			EDGES[city1] = map[string]int{city2: dist}
		}

		if v, ok := EDGES[city2]; ok {
			v[city1] = dist
		} else {
			EDGES[city2] = map[string]int{city1: dist}
		}

		if !Contains(CITIES, city1) {
			CITIES = append(CITIES, city1)
		}
		if !Contains(CITIES, city2) {
			CITIES = append(CITIES, city2)
		}
	}

	perms := make(chan []string, 100)
	go Permutations(perms)

	shortest := int(math.MaxInt32)
	path := []string{}
	for {
		c := <-perms
		if len(c) == 0 {
			break
		}

		cur := c[0]
		dist := 0
		for i := range c[1:] {
			dist += EDGES[cur][c[i+1]]
			cur = c[i+1]
		}

		// fmt.Println(c, dist, shortest)

		if c[0] == "Tristram" && c[1] == "Tambi" && c[2] == "Snowdin" && c[3] == "AlphaCentauri" && c[4] == "Faerun" && c[5] == "Straylight" && c[6] == "Norrath" {
			fmt.Println(dist)
		}

		if dist < shortest {
			shortest = dist
			path = c
			// fmt.Println(c, dist, shortest)
		}
	}
	fmt.Println(path, shortest)
}
