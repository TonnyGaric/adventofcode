package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// contains validates if e already exists in s.
//
// In Go, there is no "contains" method for slices.
// See also: https://stackoverflow.com/questions/10485743/contains-method-for-a-slice#answer-10485970
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	// Store all reached frequencies in this slice
	var reachedFrequencies []int

	// Keep track if we have not found the first
	// frequency our device reaches twice
	var notFound = true

	// Initial frequency is zero
	var frequency = 0

	// Keep iterating as long as we have not found the
	// first frequency our device reaches twice.
	for notFound {
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

		// Iterate over each line from inputFile "input.txt"
		for scanner.Scan() {
			// Save current frequency, so we can print it later
			var currentFrequency = frequency

			// Line from inputFile
			var line = scanner.Text()

			// Use Atoi to convert line (string) to int.
			//
			// Note that Atoi takes the characters "+"
			// and "-" into account when converting
			// to int.
			//
			// See also: https://golang.org/pkg/strconv/#hdr-Numeric_Conversions
			i, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}

			frequency = frequency + i

			// Print the changes that occur
			fmt.Println("Current frequency", currentFrequency, "change of", line+"; resulting frequency", frequency)

			// Check if this frequency already was reached once
			if contains(reachedFrequencies, frequency) {
				fmt.Println("The first frequency our device reaches twice is", frequency)
				notFound = false
				break
			}

			// Add this frequency to reachedFrequencies
			reachedFrequencies = append(reachedFrequencies, frequency)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
