package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	file, err := os.Open("./inputs/input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var total int = 0
	arr := make([]int, 0)

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			arr = append(arr, total)
			total = 0
		} else {
			str := fileScanner.Text()
			i, err := strconv.Atoi(str)
			total += i

			if err != nil {
				panic(err)
			}
		}
	}

	var largest int = 0
	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}

	fmt.Printf("%d", largest)
	file.Close()
}