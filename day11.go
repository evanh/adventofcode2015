package main

import (
	"bytes"
	"fmt"
)

var lowest byte = 'a'
var highest byte = 'z'

var NOT_ALLOWED = []byte{'i', 'o', 'l'}

func incr(a []byte) []byte {
	var carryover uint8
	i := len(a) - 1
	a[i] += 1
	if a[i] > highest {
		a[i] = lowest
		carryover = 1
	}

	if carryover > 0 {
		i = len(a) - 2
		for carryover > 0 && i >= 0 {
			a[i] += carryover
			if a[i] > highest {
				a[i] = lowest
			} else {
				carryover = 0
			}
			i--
		}
	}

	return a
}

func verify(a []byte) bool {
	for i := range NOT_ALLOWED {
		if bytes.IndexByte(a, NOT_ALLOWED[i]) != -1 {
			return false
		}
	}

	pairs := []int{}
	consecutive := 1
	has_consecutive := false

	for i := range a {
		if i == 0 {
			continue
		}

		if a[i] == a[i-1] {
			pairs = append(pairs, i)
		}

		if a[i] == a[i-1]+1 {
			consecutive++
			if consecutive >= 3 {
				has_consecutive = true
			}
		} else {
			consecutive = 1
		}

		// Check success
		if has_consecutive && len(pairs) > 1 {
			for j := range pairs {
				if j == 0 {
					continue
				}
				if pairs[j]-pairs[j-1] > 1 {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	start := incr([]byte("cqjxxyzz"))
	for !verify(start) {
		start = incr(start)
	}
	fmt.Println(string(start), verify(start))
}
