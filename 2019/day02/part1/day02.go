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
