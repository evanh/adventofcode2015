package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Inc struct {
	Action   string
	Register string
	Value    int
}

type Computer struct {
	cmp       int
	registers map[string]int
}

func main() {
	f, _ := os.Open("day23.input")
	data, _ := ioutil.ReadAll(f)
	raw_incs := strings.Split(string(data), "\n")

	instructions := []Inc{}
	for _, r := range raw_incs {
		parts := strings.Split(r, " ")
		if len(parts) == 3 {
			var value int
			if parts[2][0] == '+' {
				value, _ = strconv.Atoi(parts[2][1:])
			} else {
				value, _ = strconv.Atoi(parts[2])
			}
			inc := Inc{
				Action:   parts[0],
				Register: parts[1][0:1],
				Value:    value,
			}
			instructions = append(instructions, inc)
		} else if parts[0] == "jmp" {
			var value int
			if parts[1][0] == '+' {
				value, _ = strconv.Atoi(parts[1][1:])
			} else {
				value, _ = strconv.Atoi(parts[1])
			}
			inc := Inc{
				Action: parts[0],
				Value:  value,
			}
			instructions = append(instructions, inc)
		} else {
			inc := Inc{
				Action:   parts[0],
				Register: parts[1],
			}
			instructions = append(instructions, inc)
		}
	}

	fmt.Println(instructions)

	computer := Computer{
		cmp:       0,
		registers: map[string]int{"a": 1, "b": 0},
	}

	for {
		if computer.cmp >= len(instructions) {
			break
		}
		inc := instructions[computer.cmp]

		// Check valid register
		if inc.Register != "" {
			if _, ok := computer.registers[inc.Register]; !ok {
				break
			}
		}

		switch inc.Action {
		case "inc":
			computer.registers[inc.Register]++
			computer.cmp++
		case "tpl":
			computer.registers[inc.Register] = computer.registers[inc.Register] * 3
			computer.cmp++
		case "hlf":
			computer.registers[inc.Register] = computer.registers[inc.Register] / 2
			computer.cmp++
		case "jmp":
			computer.cmp += inc.Value
		case "jio":
			if computer.registers[inc.Register] == 1 {
				computer.cmp += inc.Value
			} else {
				computer.cmp++
			}
		case "jie":
			if computer.registers[inc.Register]%2 == 0 {
				computer.cmp += inc.Value
			} else {
				computer.cmp++
			}
		}
	}
	fmt.Println(computer)
}
