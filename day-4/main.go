package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {
		str := fileScanner.Text()
		split := strings.Split(str, ",")
		firstPairNums := strings.Split(split[0], "-")
		secondPairNums := strings.Split(split[1], "-")
		
		num1, num2 := convertHalfsToInts(firstPairNums)
		num3, num4 := convertHalfsToInts(secondPairNums)

		fmt.Printf("%d : %d and ", num1, num2)
		fmt.Printf("%d : %d ", num3, num4)

		if num1 < num3 && num2 > num4 {
			total++
			fmt.Printf("<- true %d\n", total)
		} else if num1 == num3 && num2 > num4 {
			total++
			fmt.Printf("<- true %d\n", total)
		} else if num1 < num3 && num2 == num4 {
			total++
			fmt.Printf("<- true %d\n", total)
		} else if num3 < num1 && num4 > num2 {
			total++
			fmt.Printf("<- true %d\n", total)
		} else if num3 == num1 && num4 > num2 {
			total++
			fmt.Printf("<- true %d\n", total)
		} else if num3 < num1 && num4 == num2 {
			total++
			fmt.Printf("<- true %d\n", total)
		} else if num1 == num3 && num2 == num4 {
			total++
			fmt.Printf("<- true %d\n", total)
		} else {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("Total: %d", total)
}

func convertHalfsToInts(str []string) (int, int) {
	firstNum, err1 := strconv.Atoi(str[0])
	secondNum, err2 := strconv.Atoi(str[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Error")
	}
	
	return firstNum, secondNum
}