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

	shortestPolymerLength := len(polymer)

	// We need to remove all occurances in polymer
	// of "a" and "A", for all characters of the
	// alphabet, then apply some logic to the
	// modified polymer.
	//
	// According to Wikipedia, the upper case
	// alphabet characters are ASCII 65 (A)
	// up to and including 90 (Z)
	//
	// See: https://en.wikipedia.org/wiki/ASCII
	for i := 65; i <= 90; i++ {
		// Create a temporary polymer based on polymer,
		// so the original polymer stays preserved and
		// we start each iteration with the orginal
		// polymer.
		tempPolymer := polymer

		// The character we will remove from the tempPolymer
		upperCharToRemove := string(i)

		// Also remove the lower case variant
		// of charToRemove from tempPolymer
		lowerCharToRemove := strings.ToLower(upperCharToRemove)

		// Remove the upperCharToRemove and lowerCharToRemove
		// from tempPolymer
		tempPolymer = strings.Replace(tempPolymer, upperCharToRemove, "", -1)
		tempPolymer = strings.Replace(tempPolymer, lowerCharToRemove, "", -1)

		// Length of tempPolymer when the for loop starts
		oldPolymerLength := len(tempPolymer)

		// Length of tempPolymer after the for loop finishes.
		//
		// We will start with 0, after the first loop,
		// we have the new accurate length.
		newPolymerLength := 0

		// We will keep iterating, as long as the
		// length of tempPolymer keeps changing.
		for oldPolymerLength != newPolymerLength {
			oldPolymerLength = len(tempPolymer)

			// We need to remove all occurances in tempPolymer
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
				tempPolymer = strings.Replace(tempPolymer, upperChar+lowerChar, "", -1)
				tempPolymer = strings.Replace(tempPolymer, lowerChar+upperChar, "", -1)
			}

			newPolymerLength = len(tempPolymer)
		}

		if newPolymerLength < shortestPolymerLength {
			shortestPolymerLength = newPolymerLength
		}

		fmt.Println("Removing all units of "+upperCharToRemove+"/"+lowerCharToRemove+" and fully reacting the result, produces a polymer with a length of", newPolymerLength)
	}

	// Print final answer
	fmt.Println(shortestPolymerLength, "is the length of the shortest polymer we can remove by removing all units of exactly one type and fully reacting the result.")
}
