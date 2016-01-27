package main

import (
	"fmt"
	"math"
)

func main() {
	var presents = 0.0
	var step = 1.0
	var house = 1.0
	for {
		presents = 0.0
		for i := 1.0; i < math.Sqrt(house); i += 1 {
			if int(house)%int(i) == 0 {
				visits := house / i
				if visits <= 50 {
					presents += 11.0 * i
				}
				if i <= 50 {
					presents += 11.0 * visits
				}
			}
		}

		if presents >= 33100000 {
			break
		}

		if step == 1 || step == -1 {
			house += step
			continue
		}

		if step > 0 {
			if presents < 33100000 {
				house += step
			} else {
				fmt.Println(house, presents, step)
				// We passed the mark, need to go back
				step = step / 2
				if int(step)%2 != 0 {
					step += 1
				}
				step = -1 * step
			}
		} else {
			fmt.Println(house, presents, step)
			if presents > 33100000 {
				house += step
			} else {
				// We passed the mark, need to go back
				step = step / 2
				if int(step)%2 != 0 {
					step += 1
				}
				step = -1 * step
			}
		}
	}
	fmt.Println(house)
}
