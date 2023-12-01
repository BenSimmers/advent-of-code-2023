package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() []string {
	// Read input from file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func firstNum(line string, replacements [][]string) int {
	for i := range line { // Loop through each character in the line
		for _, item := range replacements { // Loop through each replacement
			if strings.HasPrefix(line[i:], item[1]) { // If the line starts with the replacement
				num, _ := strconv.Atoi(item[0]) // Convert the replacement to an int
				return num
			}
		}
	}
	return 0
}

func calibrate(line string, replacements [][]string) int {
	reversedReplacements := make([][]string, len(replacements)) // Create a new slice with the same length as replacements
	for i, v := range replacements {                            // Loop through each replacement
		reversedReplacements[i] = []string{v[0], reverseString(v[1])} // Reverse the replacement and add it to the new slice
	}
	return (firstNum(line, replacements) * 10) + firstNum(reverseString(line), reversedReplacements) // Return the sum of the first number in the line and the first number in the reversed line
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // Loop through each character in the string
		runes[i], runes[j] = runes[j], runes[i] // Swap the characters
	}
	return string(runes) 
}

func part1(data []string) int {
	replacements := [][]string{}     // Create a new slice
	for i, v := range "0123456789" { // Loop through each number
		replacements = append(replacements, []string{strconv.Itoa(i), string(v)}) // Add the number and the string to the slice
	}
	sum := 0
	for _, line := range data { // Loop through each line
		sum += calibrate(line, replacements) // Add the calibration to the sum
	}
	return sum // Return the sum
}

func part2(data []string) int { // Same as part 1 but with the words
	replacements := [][]string{}
	for i, v := range "0123456789" { // Loop through each number and add it to the slice
		replacements = append(replacements, []string{strconv.Itoa(i), string(v)})
	}
	for i, word := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} { // Loop through each word and add it to the slice
		replacements = append(replacements, []string{strconv.Itoa(i), word}) // Add the number and the word to the slice
	}
	sum := 0
	for _, line := range data { // Loop through each line
		sum += calibrate(line, replacements) // Add the calibration to the sum
	}
	return sum
}

func main() {
	data := getInput()
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}
