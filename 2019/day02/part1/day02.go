package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Incode Program
func intcodeProgram(list string) string {
	//TODO

	// Double check on the return and input type, not sure if those are correct

	// Loop through the list in order to find each of the opcodes (they are 4 apart, positions 0,4,8,12,etc.)

	// Need a case for all of the following

	// 1 is take next two positions add those and store at position of the 3rd

	// 2 is the same as 1 but multiply instead of add

	// 99 is stop (and return the total list)

	return "0"
}

// Processes input file
func fileProccessor(puzzle string) int {

	// Open file
	file, err := os.Open(puzzle)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// I need to learn how to initialize a slice/array rather than doing this
	var input = ""
	for scanner.Scan() {
		input = (scanner.Text())
	}
	var list = strings.Split((input), ",")
	fmt.Println(list)

	// Pass it off to the function

	return 0
}

func main() {
	fmt.Println(fileProccessor("./input.txt"))
}
