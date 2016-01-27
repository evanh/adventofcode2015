package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var SPLITTER = " => "
var MOLECULE = "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF"

type Outcome struct {
	sync.RWMutex
	Results map[string]bool
}

var OUTCOMES = Outcome{Results: map[string]bool{}}

func Replacement(find, replace string, wg *sync.WaitGroup) {
	defer wg.Done()

	var new_mol string

	index := 0
	location := strings.Index(MOLECULE[index:], find)
	for location != -1 {
		replaced := strings.Replace(MOLECULE[index:], find, replace, 1)

		// copy(MOLECULE, new_mol)
		new_mol = MOLECULE[:index] + replaced
		OUTCOMES.Lock()
		OUTCOMES.Results[new_mol] = true
		OUTCOMES.Unlock()

		location = strings.Index(MOLECULE[index+len(find):], find)
		index += location + len(find)
	}
}

func main() {
	f, _ := os.Open("day19.input")
	data, _ := ioutil.ReadAll(f)
	rule_data := strings.Split(string(data), "\n")

	var wg *sync.WaitGroup = &sync.WaitGroup{}
	for i := range rule_data {
		parts := strings.Split(rule_data[i], SPLITTER)
		find := parts[0]
		replace := parts[1]
		wg.Add(1)
		go Replacement(find, replace, wg)
	}
	wg.Wait()
	fmt.Println(len(OUTCOMES.Results))
}
