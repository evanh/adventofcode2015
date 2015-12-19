package main

import (
	"fmt"
	"io/ioutil"
	"os"
	// "regexp"
	"strings"
)

func main() {
	// var hex, _ = regexp.Compile("\\\\x[[:alnum:]][[:alnum:]]")
	f, _ := os.Open("day8.input")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(string(data), "\n")
	codelen := 0
	enclen := 0
	for _, l := range lines {
		codelen += len(l)
		enc := "\""
		for _, c := range l {
			switch c {
			case '"':
				enc += "\\\""
			case '\\':
				enc += "\\\\"
			default:
				enc += string(c)
			}
		}
		enc += "\""
		fmt.Println(enc)
		enclen += len(enc)
	}
	fmt.Println(enclen - codelen)
}

// func main() {
// 	var hex, _ = regexp.Compile("\\\\x[[:alnum:]][[:alnum:]]")
// 	f, _ := os.Open("day8.input")
// 	data, _ := ioutil.ReadAll(f)
// 	lines := strings.Split(string(data), "\n")
// 	codelen := 0
// 	memlen := 0
// 	for _, l := range lines {
// 		codelen += len(l)
// 		inmem := strings.Replace(l, "\\\\", "a", -1)
// 		inmem = strings.Replace(inmem, "\\\"", "a", -1)
// 		inmem = hex.ReplaceAllString(inmem, "a")
// 		memlen += len(inmem) - 2
// 	}
// 	fmt.Println(codelen - memlen)
// }
