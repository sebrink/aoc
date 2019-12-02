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
	var fabric[1000][1000]int
	var counter int
	var claims []claim
	var line string

	// Process data for struct
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var c claim
		line = scanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
		claims = append(claims, c)

		// Process data for part one
		for i := c.x; i < c.x+c.w; i++ {
			for j := c.y; j < c.y+c.h; j++ {
				if fabric[i][j] == 0 {
					fabric[i][j] = 1
				} else if fabric[i][j] == 1{
					counter++
					fabric[i][j]++
				}
			}
		}
	}

	// Process data for part two
	for index, claim := range claims {
		var flag bool
		for i := claim.x; i < claim.x+claim.w; i++{
			if flag {
				break
			}
			for j := claim.y; j < claim.y+claim.h; j++{
				if fabric[i][j] != 1{
					flag = true
					break
				}
			}

		}
		if !flag {
			fmt.Printf("Part 1: %v\n", counter)
			fmt.Printf("Part 2: %v\n", index+1)
			break
		}
	}
}