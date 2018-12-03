package main

// Imports
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Struct to hold each claim
type claim struct {
	id int
	x,y int
	w,h int
}

func main() { 

	//Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Init variables
	//var fabric[1000][1000]int // Not used yet
	var claims []claim
	var line string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		var c claim
		line = scanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
		claims = append(claims, c)
		fmt.Printf("%+v\n", c)
	}
}