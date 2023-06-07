package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)	

	var total int = 0;
	for fileScanner.Scan() {
		str := fileScanner.Text()
		charArr := strings.Split(str, "")

		// Wins
		if charArr[2] == "X" && charArr[0] == "C" {
			total += (1 + 6) 
		} else if charArr[2] == "Y" && charArr[0] == "A" {
			total += (2 + 6) 
		} else if charArr[2] == "Z" && charArr[0] == "B" {
			total += (3 + 6) 
		}

		// Draw
		if charArr[2] == "X" && charArr[0] == "A" {
			total += (1 + 3) 
		} else if charArr[2] == "Y" && charArr[0] == "B" {
			total += (2 + 3) 
		} else if charArr[2] == "Z" && charArr[0] == "C" {
			total += (3 + 3)
		}

		// Loss
		if charArr[2] == "X" && charArr[0] == "B" {
			total += 1 
		} else if charArr[2] == "Y" && charArr[0] == "C" {
			total += 2 
		} else if charArr[2] == "Z" && charArr[0] == "A" {
			total += 3
		}
	}

	fmt.Printf("Total Score: %d", total)
}