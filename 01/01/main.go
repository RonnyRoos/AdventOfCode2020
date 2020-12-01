package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("01/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbersInFile []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbersInFile = append(numbersInFile, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Iterate array
	var found = false
	for _, num := range numbersInFile {
		for _, innerNum := range numbersInFile {
			if num+innerNum == 2020 {
				fmt.Println(num * innerNum)
				found = true
				break
			}

			if found {
				break
			}
		}
	}

}
