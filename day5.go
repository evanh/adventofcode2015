package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func verifyString(s string) bool {
	pairs := map[byte]map[byte]int{}

	paired := false
	middle := false
	for i := range s {
		if i < len(s)-1 {
			if second, ok := pairs[s[i]]; ok {
				if index, ok := second[s[i+1]]; ok {
					if i-index > 1 {
						paired = true
					}
				} else {
					second[s[i+1]] = i
				}
			} else {
				pairs[s[i]] = map[byte]int{s[i+1]: i}
			}
		}

		if i > 1 {
			if s[i-2] == s[i] {
				middle = true
			}
		}
	}

	return paired && middle

}

func main() {
	f, _ := os.Open("day5.input")
	data, _ := ioutil.ReadAll(f)
	inputs := strings.Split(string(data), "\n")
	nice := 0
	for _, input := range inputs {
		fmt.Println(input, verifyString(input))
		if verifyString(input) {
			nice++
		}
	}
	fmt.Println("Nice:", nice)
}
