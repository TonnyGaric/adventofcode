package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	// Keep track of how many times two of any letter and
	// three of any letter occurs
	var twoOfAnyLetter = 0
	var threeOfAnyLetter = 0

	// Iterate over each line from inputFile "input.txt"
	for scanner.Scan() {
		// Line from inputFile
		var id = scanner.Text()

		// If a letter appears two or three times in this id,
		// we must only count it once. So we keep track if
		// this id already contains two or three
		// appearances of any letter.
		var containsTwoOfAnyLetter = false
		var containsThreeOfAnyLetter = false

		// Iterate over each character in this id
		for i := 0; i < len(id); i++ {
			// We expect id only consists of ASCII characters,
			// single-byte corresponding to the first 128
			// Unicode characters. Otherwise, we should
			// have used runes, because of UTF-8 chars:
			//
			// string([]rune(id)[i]))
			//
			// See also: https://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
			var letter = string(id[i])

			// Count how many times this letter appears in this id
			var count = strings.Count(id, letter)

			fmt.Println("Letter", letter, "appears", count, "times in ID", id)

			// If containsThreeOfAnyLetter is false
			if !containsThreeOfAnyLetter && count == 3 {
				// If this letter appears three times in this id,
				// we increment threeOfAnyLetter and set
				// containsThreeOfAnyLetter to true
				threeOfAnyLetter++
				containsThreeOfAnyLetter = true
			} else if !containsTwoOfAnyLetter && count == 2 {
				// If this letter appears two times in this id,
				// we increment twoOfAnyLetter and set
				// containsTwoOfAnyLetter to true
				twoOfAnyLetter++
				containsTwoOfAnyLetter = true
			}

			if containsTwoOfAnyLetter && containsThreeOfAnyLetter {
				// If containsTwoOfAnyLetter AND containsThreeOfAnyLetter is true,
				// we can exit this loop, because we can count two and three
				// appearanches of any letter in this id, only once.
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Multiplying twoOfAnyLetter and threeOfAnyLetter
	// produces a checksum
	var checksum = twoOfAnyLetter * threeOfAnyLetter

	// Print the final answer
	fmt.Println("Of these box IDs,",
		twoOfAnyLetter,
		"of them contain a letter which appears exactly twice, and",
		threeOfAnyLetter,
		"of them contain a letter which appears exactly three times. Multiplying these together produces a checksum of",
		checksum)
}
