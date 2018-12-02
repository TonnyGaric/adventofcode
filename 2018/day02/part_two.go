package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// compareIDs compares each character of this id
// with each character of eacg ID from input.txt.
//
// If an ID is found that differs by only one
// character from this id, we return the
// index of that character.
func compareIDs(id string) (int, error) {
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
		// Line from inputFile
		var idToCompare = scanner.Text()

		// Keep track how many letters mismatch
		var mismatchedLetters = 0

		// Keep track of the index of the last
		// mismatched letter
		var indexOfMismatchedLetter = 0

		// Iterate over each character in this id
		for i := 0; i < len(id); i++ {
			var letter = string(id[i])
			var letterToCompare = string(idToCompare[i])

			if letter != letterToCompare {
				mismatchedLetters++
				indexOfMismatchedLetter = i
			}
		}

		// If there is only onel mismatched letter, we
		// have found two correct IDs. The common
		// letters are found by removing the
		// differing charachter from either
		// ID. We return this index, so we
		// can later delete it.
		if mismatchedLetters == 1 {
			return indexOfMismatchedLetter, nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0, errors.New("Did not found ID that differs by exactly one character")
}

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

	var commonLetters = ""

	// Iterate over each line from inputFile "input.txt"
	for scanner.Scan() {
		// Line from inputFile
		var id = scanner.Text()

		var indexOfMismatchedLetter, err = compareIDs(id)

		// If compareIDs has no error, it means that a ID
		// is found with only one character different
		// from this id.
		if err == nil {
			fmt.Println("Index of mismatched letter:", indexOfMismatchedLetter)

			// If indexOfMismatchedLetter is the first OR last
			// charachter of this id, we must substring it
			// differently to prevent index out of range
			// and other errors.
			if indexOfMismatchedLetter == 0 {
				commonLetters = id[1:len(id)]
			} else if indexOfMismatchedLetter == len(id)-1 {
				commonLetters = id[0 : len(id)-2]
			} else {
				commonLetters = id[0:indexOfMismatchedLetter] + id[indexOfMismatchedLetter+1:len(id)]
			}

			// We have found the commonLetters, so we can
			// exit this loop.
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the final answer
	fmt.Println("The following letters are common between the two correct box IDs:", commonLetters)
}
