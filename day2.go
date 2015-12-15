package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getNeeded(w, l, h float64) float64 {
	sides := []float64{w * l, w * h, h * l}
	smallest := sides[0]
	for i := range sides[1:] {
		if sides[i] < smallest {
			smallest = sides[i]
		}
	}
	return 2*sides[0] + 2*sides[1] + 2*sides[2] + smallest
}

func main() {
	f, _ := os.Open("day2.input")
	input, _ := ioutil.ReadAll(f)
	presents := strings.Split(input, "\n")
	var w, l, h float64
	var parts []string
	total_area := 0
	for _, p := range presents {
		parts = strings.Split(p, "x")
		w = strconv.ParseFloat(parts[0])
		l = strconv.ParseFloat(parts[1])
		h = strconv.ParseFloat(parts[2])
		total_area += getNeeded(w, l, h)
	}

	fmt.Println("We need", total_area)

}
