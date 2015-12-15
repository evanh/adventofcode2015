package main

import (
	// "bytes"
	"crypto/md5"
	"fmt"
)

var KEY = []byte("ckczppom")
var VALID = []byte{0, 0}

func incr(a []byte) []byte {
	var carryover uint8
	i := len(a) - 1
	a[i] += 1
	if a[i] > 57 {
		a[i] = 48
		carryover = 1
	}

	if carryover > 0 {
		i = len(a) - 2
		for carryover > 0 && i >= 0 {
			a[i] += carryover
			if a[i] > 57 {
				a[i] = 48
			} else {
				carryover = 0
			}
			i--
		}
	}

	if carryover > 0 {
		a = append([]byte{48 + carryover}, a...)
	}

	return a
}

func main() {
	start := []byte{48}
	for {
		start = incr(start)
		hash := md5.Sum(append(KEY, start...))
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			fmt.Printf("%x %s\n", hash, start)
			break
		}
	}
}
