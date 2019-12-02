package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Does the fuel calculation
func calcFuel(fuel int) int {
	// Divide by 3, round down (go does this by default with a "/"), subtract 2
	return ((fuel / 3) - 2)
}

// Processes input file
func fileProccessor(puzzle string) int {

	// Open file
	file, err := os.Open(puzzle)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Take in input and process data
	scanner := bufio.NewScanner(file)
	var total = 0
	for scanner.Scan() {
		// Convert string to int
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		// Calculate the fuel and add it to the total
		total += calcFuel(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total
}

func main() {
	fmt.Println(fileProccessor("./input.txt"))
}
