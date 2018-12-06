package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Slurp the entire content of "input.txt"
	// into our memory.
	inputData, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// inputData is of type []byte, we must convert
	// it to string.
	polymer := string(inputData)

	// Length of polymer when the for loop starts
	oldPolymerLength := len(polymer)

	// Length of polymer after the for loop finishes.
	//
	// We will start with 0, after the first loop,
	// we have the new accurate length.
	newPolymerLength := 0

	// We will keep iterating, as long as the
	// length of polymer keeps changing.
	for oldPolymerLength != newPolymerLength {
		oldPolymerLength = len(polymer)

		// We need to remove all occurances in polymer
		// of "aA" and "Aa", for all characters of the
		// alphabet.
		//
		// According to Wikipedia, the upper case
		// alphabet characters are ASCII 65 (A)
		// up to and including 90 (Z)
		//
		// See: https://en.wikipedia.org/wiki/ASCII
		for i := 65; i <= 90; i++ {
			upperChar := string(i)
			lowerChar := strings.ToLower(upperChar)

			// -1 means that we will replace all occurances
			polymer = strings.Replace(polymer, upperChar+lowerChar, "", -1)
			polymer = strings.Replace(polymer, lowerChar+upperChar, "", -1)
		}

		newPolymerLength = len(polymer)
	}

	// This will be our final answer
	answer := len(polymer)

	// Print final answer
	fmt.Println(answer, "units remain after fully reacting the polymer we scanned.")
}
