package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

type Recipe []int

var INGREDIENTS = []Ingredient{}

func incr(a Recipe) {
	var carryover int
	i := len(a) - 1
	a[i] += 1
	if a[i] > 100 {
		a[i] = 0
		carryover = 1
	}

	if carryover > 0 {
		i = len(a) - 2
		for carryover > 0 && i >= 0 {
			a[i] += carryover
			if a[i] > 100 {
				a[i] = 0
			} else {
				carryover = 0
			}
			i--
		}
	}
}

func impossible(r Recipe) bool {
	sum := 0
	cal := 0
	for i := range r {
		sum += r[i]
		cal += r[i] * INGREDIENTS[i].Calories
	}
	return sum != 100 || cal != 500
}

func Goodness(r Recipe) int {
	var capacity, durability, flavor, texture int
	for i, num := range r {
		capacity += num * INGREDIENTS[i].Capacity
		durability += num * INGREDIENTS[i].Durability
		flavor += num * INGREDIENTS[i].Flavor
		texture += num * INGREDIENTS[i].Texture
	}

	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0
	}

	return capacity * durability * flavor * texture
}

func main() {
	f, _ := os.Open("day15.input")
	data, _ := ioutil.ReadAll(f)
	ings := strings.Split(string(data), "\n")
	for _, ing := range ings {
		parts := strings.Split(ing, " ")
		c, _ := strconv.Atoi(parts[2][:len(parts[2])-1]) // Strip comma
		d, _ := strconv.Atoi(parts[4][:len(parts[4])-1]) // Strip comma
		f, _ := strconv.Atoi(parts[6][:len(parts[6])-1]) // Strip comma
		t, _ := strconv.Atoi(parts[8][:len(parts[8])-1]) // Strip comma
		l, _ := strconv.Atoi(parts[10])

		i := Ingredient{
			Capacity:   c,
			Durability: d,
			Flavor:     f,
			Texture:    t,
			Calories:   l,
		}
		INGREDIENTS = append(INGREDIENTS, i)
	}

	fmt.Println(INGREDIENTS)

	r := make(Recipe, len(INGREDIENTS))
	iterations := math.Pow(100.0, float64(len(INGREDIENTS)))
	best := 0
	recipe := ""
	var goodness int
	for i := 0; i <= int(iterations); i++ {
		// fmt.Println(r)
		if impossible(r) {
			incr(r)
			continue
		}
		goodness = Goodness(r)
		if goodness > best {
			best = goodness
			recipe = fmt.Sprintf("%s", r)
		}
		incr(r)
	}
	fmt.Println(recipe, best)
}
