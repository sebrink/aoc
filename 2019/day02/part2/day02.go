package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Incode Program
// I literally made this code like 50 times faster, look at those old commits
func intcodeProgram(array []int) int {
	var i = 0
	for array[i] != 99 {
		// Cover opcode 1
		if array[i] == 1 {
			array[array[i+3]] = array[array[i+1]] + array[array[i+2]]
		}
		if array[i] == 2 {
			array[array[i+3]] = array[array[i+1]] * array[array[i+2]]
		}
		i += 4
	}
	return array[0]
}
func main() {
	// Open file
	input, _ := ioutil.ReadFile("input.txt")
	ints := make([]int, 0)
	for _, v := range strings.Split(string(input), ",") {
		i, _ := strconv.Atoi(v)
		ints = append(ints, i)
	}
	fmt.Println(intcodeProgram(ints))
}
