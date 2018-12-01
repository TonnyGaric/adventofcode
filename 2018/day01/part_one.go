package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file "input.txt" for reading
	inputFile, err := os.Open("input.txt")

	// Opening file can have an error. If there is
	// an error, it will be of type *PathError.
	if err != nil {
		log.Fatal(err)
	}

	// Closes the file when we are done
	defer inputFile.Close()

	// Declare scanner to read from inputFile. Note that
	// the split function defaults to ScanLines—which
	// is each line of text.
	scanner := bufio.NewScanner(inputFile)

	// Initial frequency is zero
	var frequency = 0

	// Iterate over each line from inputFile "input.txt"
	for scanner.Scan() {
		// Save current frequency, so we can print it later
		var currentFrequency = frequency

		// Line from inputFile
		var line = scanner.Text()

		// The first character of each line should be a
		// '+' or '-'. The other characters should be a
		// number. So we convert all charachters,
		// except the first one, to an int.
		i, err := strconv.Atoi(line[1:len(line)])

		if err != nil {
			log.Fatal(err)
		}

		// Determine what the first character is, so we
		// know if we must add or substract from
		// frequency.
		if line[0] == '+' {
			frequency = frequency + i
		} else if line[0] == '-' {
			frequency = frequency - i
		}

		// Print the changes that occur
		fmt.Println("Current frequency", currentFrequency, "change of", line+"; resulting frequency", frequency)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the final answer
	fmt.Println("Answer:", frequency)
}
