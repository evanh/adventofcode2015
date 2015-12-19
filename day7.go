package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Connection struct {
	Action string
	Value  uint16
}

func IntOrValue(wire string, tree map[string]*Connection) uint16 {
	var v1 uint16
	if v, err := strconv.Atoi(wire); err == nil {
		v1 = uint16(v)
	} else {
		v1 = GetValue(tree[wire], tree)
	}
	return v1
}

func GetValue(c *Connection, tree map[string]*Connection) uint16 {
	if c.Action == "" {
		return c.Value
	}

	parts := strings.Split(c.Action, " ")
	if len(parts) == 1 {
		wire := parts[0]
		c.Value = GetValue(tree[wire], tree)
	} else if len(parts) == 2 {
		wire := parts[1] // Wire to NOT to get value
		c.Value = ^GetValue(tree[wire], tree)
	} else {
		switch parts[1] {
		case "AND":
			wire1 := IntOrValue(parts[0], tree)
			wire2 := IntOrValue(parts[2], tree)
			c.Value = wire1 & wire2
		case "OR":
			wire1 := IntOrValue(parts[0], tree)
			wire2 := IntOrValue(parts[2], tree)
			c.Value = wire1 | wire2
		case "LSHIFT":
			wire1 := IntOrValue(parts[0], tree)
			value, _ := strconv.Atoi(parts[2])
			c.Value = wire1 << uint(value)
		case "RSHIFT":
			wire1 := IntOrValue(parts[0], tree)
			value, _ := strconv.Atoi(parts[2])
			c.Value = wire1 >> uint(value)
		}
	}
	c.Action = ""
	return c.Value
}

func main() {
	var tree = map[string]*Connection{}

	f, _ := os.Open("day7.input")
	data, _ := ioutil.ReadAll(f)
	rules := strings.Split(string(data), "\n")
	for _, r := range rules {
		parts := strings.Split(r, " -> ")
		c := &Connection{}
		if value, err := strconv.Atoi(parts[0]); err == nil {
			c.Value = uint16(value)
		} else {
			c.Action = parts[0]
		}
		if parts[1] == "b" {
			c.Value = 3176
			c.Action = ""
		}
		tree[parts[1]] = c
	}

	fmt.Println(GetValue(tree["a"], tree))
}
