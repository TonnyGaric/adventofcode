package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	ID                 int
	inchesFromLeftEdge int
	inchesFromTopEdge  int
	inchesWide         int
	inchesTall         int
}

func main() {
	// Create the initial fabric.
	//
	// Each key in this map, represents a row in the
	// piece of fabric. The value of each key, is a
	// slice of ints. Each element in this slice,
	// represents a column in the piece of
	// fabric.
	fabric := make(map[int][]int)

	// According to the README, the fabric is a very
	// large squure, with at least 1000 inches on
	// each side.
	//
	// For each row (key in map), create a slice
	// (value of key) with 1000 ints.
	//
	// The key will be the number of the row. We
	// want the row to start at 1.
	for i := 1; i < 1001; i++ {
		// Set the value of this key to a new slice
		// of 100 ints. Each element in this slice
		// will be initially zero-valued. For
		// ints this means 0.
		fabric[i+1] = make([]int, 1000)
	}

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

	// Save all claims to this map, so we can later
	// search for the claim of which each square
	// inch is not overlapped by any other
	// claim.
	//
	// Note that we do not use a slice, because
	// a slice must have a length and we do not
	// now how many claims we can expect.
	claims := make(map[int]claim)

	// Iterate over each line from inputFile "input.txt"
	for scanner.Scan() {
		// Line from inputFile
		line := scanner.Text()

		// A claim consists of a pattern as follows:
		//
		// #13 @ 176,605: 24x11
		//  ^^   ^^^ ^^^  ^^ ^^
		//  1     2   3   4  5
		//
		// Where:
		// 1 is the claim ID
		// 2 is the inches from the left edge
		// 3 is the inches from the top edge
		// 4 is the inches wide
		// 5 is the inches tall

		// Now extract each item from the line

		claimID, err := strconv.Atoi(line[strings.Index(line, "#")+1 : strings.Index(line, " ")])

		if err != nil {
			log.Fatal(err)
		}

		inchesFromLeftEdge, err := strconv.Atoi(line[strings.Index(line, "@")+2 : strings.Index(line, ",")])

		if err != nil {
			log.Fatal(err)
		}

		inchesFromTopEdge, err := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, ":")])

		if err != nil {
			log.Fatal(err)
		}

		inchesWide, err := strconv.Atoi(line[strings.Index(line, ":")+2 : strings.Index(line, "x")])

		if err != nil {
			log.Fatal(err)
		}

		inchesTall, err := strconv.Atoi(line[strings.Index(line, "x")+1:])

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Claim ID:",
			claimID,
			"inches from left edge:",
			inchesFromLeftEdge,
			"inches from top edge:",
			inchesFromTopEdge,
			"inches wide:",
			inchesWide,
			"inches tall:",
			inchesTall)

		// Add this claim to the fabric
		for row := inchesFromTopEdge; row < inchesFromTopEdge+inchesTall; row++ {
			for column := inchesFromLeftEdge; column < inchesFromLeftEdge+inchesWide; column++ {
				rowSlice := fabric[row]

				// Increment this square of inch,
				// so we now how many claims are
				// within this square inch.
				rowSlice[column] = rowSlice[column] + 1
			}
		}

		// Add this claim to the map claims
		claims[claimID] = claim{ID: claimID, inchesFromLeftEdge: inchesFromLeftEdge, inchesFromTopEdge: inchesFromTopEdge, inchesWide: inchesWide, inchesTall: inchesTall}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// This will be our final answer: the ID of
	// the only claim that does not overlap.
	answer := 0

	// Check for each claim, if each square inch of
	// this claim is or is not overlapped by any
	// other claim.
	for _, claim := range claims {
		timesOverlapped := 0

		for row := claim.inchesFromTopEdge; row < claim.inchesFromTopEdge+claim.inchesTall; row++ {
			for column := claim.inchesFromLeftEdge; column < claim.inchesFromLeftEdge+claim.inchesWide; column++ {
				rowSlice := fabric[row]

				// Increment this square of inch,
				// so we now how many claims are
				// within this square inch.
				//
				// rowSlice[column] is the amount
				// of times that this column is
				// used by a claim. A square
				// inch is overlapped, if
				// this is 2 or higher.
				if rowSlice[column] >= 2 {
					timesOverlapped++
				}
			}
		}

		if timesOverlapped == 0 {
			answer = claim.ID
			// We have found our answer and
			// can safely exit this loop.
			break
		}
	}

	// Print the final answer
	fmt.Println(answer, "is the ID of the only claim that does not overlap.")
}
