package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func on(x int) int {
	return x + 1
}

func off(x int) int {
	if x == 0 {
		return x
	}
	return x - 1
}

func toggle(x int) int {
	return x + 2
}

type Grid [][]int

func (g Grid) action(x1, y1, x2, y2 int, f func(int) int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			g[x][y] = f(g[x][y])
		}
	}
}

func (g Grid) on(x1, y1, x2, y2 int) {
	g.action(x1, y1, x2, y2, on)
}

func (g Grid) off(x1, y1, x2, y2 int) {
	g.action(x1, y1, x2, y2, off)
}

func (g Grid) toggle(x1, y1, x2, y2 int) {
	g.action(x1, y1, x2, y2, toggle)
}

func (g Grid) count() int {
	count := 0
	for x := 0; x < len(g); x++ {
		for y := 0; y < len(g[x]); y++ {
			count += g[x][y]
		}
	}
	return count
}

func main() {
	grid := make(Grid, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	f, _ := os.Open("day6.input")
	data, _ := ioutil.ReadAll(f)
	moves := strings.Split(string(data), "\n")

	var p1, p2, ac string
	for _, m := range moves {
		parts := strings.Split(m, " ")
		if len(parts) == 5 {
			p1 = parts[2]
			p2 = parts[4]
			ac = parts[1]
		} else {
			p1 = parts[1]
			p2 = parts[3]
			ac = parts[0]
		}

		parts = strings.Split(p1, ",")
		x1, _ := strconv.Atoi(parts[0])
		y1, _ := strconv.Atoi(parts[1])

		parts = strings.Split(p2, ",")
		x2, _ := strconv.Atoi(parts[0])
		y2, _ := strconv.Atoi(parts[1])

		switch ac {
		case "on":
			grid.on(x1, y1, x2, y2)
		case "off":
			grid.off(x1, y1, x2, y2)
		case "toggle":
			grid.toggle(x1, y1, x2, y2)
		}
	}

	fmt.Println(grid.count())
}
