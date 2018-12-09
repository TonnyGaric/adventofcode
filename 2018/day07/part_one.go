package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	// Keep track of the order in which our
	// instructions should be completed.
	// This will be our final answer.
	var stepsOrder string

	// This are all staps we can parse from inputfile
	steps := make(map[string]map[string]struct{})

	// Iterate over each line from inputFile "input.txt"
	for scanner.Scan() {
		// Line from inputFile
		var instruction = scanner.Text()

		// An instruction looks like the following:
		//
		// Step G must be finished before step Z can begin.
		//      ^                              ^
		//      5                              36
		//      prerequisite				   step

		// Parse the prerequisite and step
		// from this instruction.
		prerequisite := string(instruction[5])
		stepName := string(instruction[36])

		if _, prs := steps[stepName]; prs {
			// This step is present in map steps,
			// retrieve it and add this
			// prerequisite to it.
			steps[stepName][prerequisite] = struct{}{}
		} else {
			// This step is not present in map steps,
			// so add it.
			prerequisites := make(map[string]struct{})
			prerequisites[prerequisite] = struct{}{}
			steps[stepName] = prerequisites
		}

		// This prerequisite is also a step, so
		// we should add it to steps, if it is
		// not already present.
		if _, prs := steps[prerequisite]; !prs {
			prerequisites := make(map[string]struct{})
			steps[prerequisite] = prerequisites
		}
	}

	// Iterate over steps and apply some logic
	// to determine what the next step is.
	for len(steps) != 0 {
		// Keep track of availableSteps
		var availableSteps []string

		// Find next available step, by iterating over
		// steps and finding the step of which the
		// value is an empty map. This would mean
		// that all prerequisites are done.
		for stepName, prerequisites := range steps {
			if prerequisites == nil || len(prerequisites) == 0 {
				// This step has no prerequisites,
				// so we add it to availableSteps.
				availableSteps = append(availableSteps, stepName)
			}
		}

		// Sort slice availableSteps
		// on alphabetical order.
		sort.Strings(availableSteps)

		if len(availableSteps) != 0 {
			// This is the first alphabetically step
			// of all steps in availableSteps.
			nextStep := availableSteps[0]

			// Delete nextStep from map steps
			delete(steps, nextStep)

			// Delete nextStep as prerequisite from
			// all steps in steps.
			for _, prerequisites := range steps {
				delete(prerequisites, nextStep)
			}

			// Add nextStep to stepsOrder
			stepsOrder = stepsOrder + nextStep
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the final answer
	fmt.Println(stepsOrder, "is the order in which the steps in our instructions should be completed.")
}
