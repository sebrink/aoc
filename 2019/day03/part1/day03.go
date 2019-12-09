package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	// Open file
	input, _ := ioutil.ReadFile("sample.txt")

	// Make the input into a list
	directionList := make([]string, 0)
	for _, v := range strings.Split(string(input), ",") {
		directionList = append(directionList, v)
	}

}
