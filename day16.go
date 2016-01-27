package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var SUE = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	f, _ := os.Open("day16.input")
	data, _ := ioutil.ReadAll(f)
	sues := strings.Split(string(data), "\n")

	for i := range sues {
		sue := sues[i]
		parts := strings.Split(sue, ",")
		matched := true
		for j := range parts {
			raw := strings.Split(parts[j], " ")
			var key string
			var value int
			if len(raw) == 4 {
				key = raw[2][:len(raw[2])-1] // Strip semicolon
				value, _ = strconv.Atoi(raw[3])
			} else {
				key = raw[1][:len(raw[1])-1] // Strip semicolon
				value, _ = strconv.Atoi(raw[2])
			}

			if key == "cats" || key == "trees" {
				matched = matched && (value > SUE[key])
			} else if key == "pomeranians" || key == "goldfish" {
				matched = matched && (value < SUE[key])
			} else {
				matched = matched && (SUE[key] == value)
			}
		}
		if matched {
			fmt.Println("Sue:", i+1)
		}
	}

}
