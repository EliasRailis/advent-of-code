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

	arr := make([]string, 0)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		stringLength := len(str)
		strFirstHalf := str[0:(stringLength / 2)]
		strSecondHalf := str[stringLength / 2:stringLength]

		fmt.Printf("%s : %s\n", strFirstHalf, strSecondHalf)
		currentArr := make([]string, 0)
		for i := 0; i < len(strFirstHalf) - 1; i++ {
			currentChar := string(rune(strFirstHalf[i]))

			for x := 0; x < len(strSecondHalf); x++ {
				secCurrentChar := string(rune(strSecondHalf[x]))
				if currentChar == secCurrentChar {
					currentArr = append(currentArr, currentChar)
				} 
			}
		}

		if len(currentArr) != 0 {
			list := removeDuplicateStr(currentArr)
			for _, letter := range list {
				fmt.Print(letter, " ")
				arr = append(arr, letter)
			}
		} else {
			fmt.Print("-- None --")
		}

		fmt.Println()
	}

	var total int = 0
	for _, letter := range arr {
		index := getValue(letter)
		total += index
	}

	fmt.Println(total)
}

func removeDuplicateStr(str []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, item := range str {
		if _, value := keys[item]; !value {
			keys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func getValue(letter string) int {
	alphabetInLowerCase := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	index := strings.Index(alphabetInLowerCase, letter)
	return index + 1;
}
