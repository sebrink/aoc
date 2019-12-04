package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Incode Program
// If you've ever thought you made spaghetti code, you'll realize nothing beats this function
func intcodeProgram(array []string) string {

	// Getting my actual values
	var x, err1 = strconv.Atoi(array[1])
	if err1 != nil {
		log.Fatal(err1)
	}
	var y, err2 = strconv.Atoi(array[2])
	if err2 != nil {
		log.Fatal(err2)
	}
	var z, err3 = strconv.Atoi(array[3])
	if err3 != nil {
		log.Fatal(err3)
	}
	var addOne, errOne = strconv.Atoi(array[x])
	if errOne != nil {
		log.Fatal(errOne)
	}
	var addTwo, errTwo = strconv.Atoi(array[y])
	if errTwo != nil {
		log.Fatal(errTwo)
	}
	array[z] = strconv.Itoa(addOne + addTwo)

	for i := 0; i < len(array); i += 4 {
		// Cover opcode 99

		if array[i] == "99" {
			break
		}
		// Cover opcode 1
		if array[i] == "1" {

			// Getting my actual values
			var x, err1 = strconv.Atoi(array[i+1])
			if err1 != nil {
				log.Fatal(err1)
			}
			var y, err2 = strconv.Atoi(array[i+2])
			if err2 != nil {
				log.Fatal(err2)
			}
			var z, err3 = strconv.Atoi(array[i+3])
			if err3 != nil {
				log.Fatal(err3)
			}
			var addOne, errOne = strconv.Atoi(array[x])
			if errOne != nil {
				log.Fatal(errOne)
			}
			var addTwo, errTwo = strconv.Atoi(array[y])
			if errTwo != nil {
				log.Fatal(errTwo)
			}
			array[z] = strconv.Itoa(addOne + addTwo)
		}

		// Cover opcode 2
		if array[i] == "2" {

			// Getting my actual values
			var x, err1 = strconv.Atoi(array[i+1])
			if err1 != nil {
				log.Fatal(err1)
			}
			var y, err2 = strconv.Atoi(array[i+2])
			if err2 != nil {
				log.Fatal(err2)
			}
			var z, err3 = strconv.Atoi(array[i+3])
			if err3 != nil {
				log.Fatal(err3)
			}
			var mulOne, errOne = strconv.Atoi(array[x])
			if errOne != nil {
				log.Fatal(errOne)
			}
			var mulTwo, errTwo = strconv.Atoi(array[y])
			if errTwo != nil {
				log.Fatal(errTwo)
			}
			array[z] = strconv.Itoa(mulOne * mulTwo)
		}
	}

	return array[0]
}

// Processes input file
func fileProccessor(puzzle string) string {

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
	fmt.Println(intcodeProgram(list[:]))
	return ""
}

func main() {
	fmt.Println(fileProccessor("./input.txt"))
}
