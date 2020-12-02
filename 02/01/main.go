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
	file, err := os.Open("02/01/input")
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
		parts := strings.Split(line, ":")
		policy := strings.Split(parts[0], " ")
		stringToFind := policy[1]
		limits := strings.Split(policy[0], "-")
		lowerLimit, _ := strconv.Atoi(limits[0])
		upperLimit, _ := strconv.Atoi(limits[1])

		amountFound := 0
		for _, c := range parts[1] {
			if string(c) == stringToFind {
				amountFound++
			}
		}

		if amountFound >= lowerLimit && amountFound <= upperLimit {
			validPasswordsFound++
		}
	}

	fmt.Println("Valid passwords: ", validPasswordsFound)
}
