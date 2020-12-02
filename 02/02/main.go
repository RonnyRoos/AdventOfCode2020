package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("02/02/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	validPasswordsFound := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		policy := strings.Split(parts[0], " ")
		stringToFind := policy[1]
		limits := strings.Split(policy[0], "-")
		lowerLimit, _ := strconv.Atoi(limits[0])
		upperLimit, _ := strconv.Atoi(limits[1])

		password := parts[1]

		char1 := string(password[lowerLimit-1])
		char2 := string(password[upperLimit-1])
		if (char1 == stringToFind || char2 == stringToFind) && !(char1 == stringToFind && char2 == stringToFind) {
			validPasswordsFound++
		}
	}

	fmt.Println("Valid passwords: ", validPasswordsFound)
}
