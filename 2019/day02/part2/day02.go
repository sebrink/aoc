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

	// Just brute force it :pog:
	thisRun := make([]int, len(array))
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			copy(thisRun, array)
			thisRun[1] = n
			thisRun[2] = v
			var i = 0
			for thisRun[i] != 99 {
				// Cover opcode 1
				if thisRun[i] == 1 {
					thisRun[thisRun[i+3]] = thisRun[thisRun[i+1]] + thisRun[thisRun[i+2]]
				}
				if thisRun[i] == 2 {
					thisRun[thisRun[i+3]] = thisRun[thisRun[i+1]] * thisRun[thisRun[i+2]]
				}
				i += 4
			}
			if thisRun[0] == 19690720 {
				fmt.Println(100*n + v)
				break
			}
		}
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
