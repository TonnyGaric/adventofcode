package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type marble struct {
	next     *marble
	previous *marble
	value    int
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
	// is each line of text. Nevertheless, note that
	// we expect this inputFile to only have a
	// single line.
	scanner := bufio.NewScanner(inputFile)

	// Advance the Scanner to the next token
	scanner.Scan()

	// Number of players, as parsed from this inputFile.
	var numberOfPlayers int

	// Value of last marble, as parsed from this inputFile.
	var lastMarble int

	// Sscanf scans the argument string, storing
	// successive space-separated values into
	// successive arguments as determined by
	// the format.
	// See also: https://golang.org/pkg/fmt/#Sscanf
	fmt.Sscanf(scanner.Text(),
		"%d players; last marble is worth %d points",
		&numberOfPlayers,
		&lastMarble)

	// We need to determine what the new winning
	// Elf's score would be if the number of the
	// last marble were 100 times larger. So we
	// will multiply the last marble with 100.
	lastMarble = lastMarble * 100

	// Create a slice of ints of length
	// of number of players, to keep
	// track of the score of each
	// player.
	players := make([]int, numberOfPlayers)

	currentMarble := &marble{}
	currentMarble.next = currentMarble
	currentMarble.previous = currentMarble

	// Actual value of the current marble
	actualValue := 1

	// Keep track of the number of the
	// current player.
	var currentPlayer int

	for actualValue != lastMarble {
		currentPlayer = (actualValue - 1) % numberOfPlayers

		// If the marble that is about to be placed
		// has a number which is a multiple of 23,
		// we need to apply different logic.
		if actualValue%23 == 0 {
			// First, the current player keeps the
			// marble they would have placed,
			// adding it to their score.
			players[currentPlayer] += actualValue

			// In addition the marble 7 marbles
			// counter-clockwise from the
			// current marble is removed
			// from the circle.
			removedMarble := currentMarble.
				previous. // 1
				previous. // 2
				previous. // 3
				previous. // 4
				previous. // 5
				previous. // 6
				previous  // 7

			removedMarble.previous.next = removedMarble.next
			removedMarble.next.previous = removedMarble.previous

			// And also added to the current
			// player's score.
			players[currentPlayer] += removedMarble.value

			// The marble located immediately
			// clockwise of the marble that
			// was removed becomes the new
			// current marble.
			currentMarble = removedMarble.next
		} else {
			// Create a new marble to place
			marble := &marble{value: actualValue}

			// Set the next and previous marble,
			// based on the current marble.
			marble.next = currentMarble.next.next
			marble.previous = currentMarble.next

			currentMarble.next.next.previous = marble
			currentMarble.next.next = marble

			// The marble that was just placed
			// then becomes the current
			// marble.
			currentMarble = marble
		}

		// Increment the actual value
		actualValue++
	}

	// Keep track of the highest score, thus
	// the winning Elf's score. This will be
	// our final answer.
	var winningElfsScore int

	// Iterate over each value in map players,
	// to determine the winning Elf's score.
	for _, score := range players {
		if score > winningElfsScore {
			winningElfsScore = score
		}
	}

	// Print the final answer
	fmt.Printf("The winning Elf's score is %d.\n", winningElfsScore)
}
