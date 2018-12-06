package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// parseTime parses this line to Time, with the format:
// "2006-01-02 15:04".
//
// Example how this line can look like:
//
// [1518-04-09 00:46] wakes up
//  ^              ^
//  1              16
//
// We can expect the first 18 charachters of each line
// to be of the same pattern.
func parseTime(line string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04", line[1:17])
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

	// We will append all lines from inputFile to this
	// slice, so we can later sort this slice to
	// chronological order.
	var lines = []string{}

	// Iterate over each line from inputFile "input.txt",
	// and append it to the slice lines.
	for scanner.Scan() {
		// Line from inputFile
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort slice to have all lines in chronological order,
	// so we can later analyize it.
	sort.Strings(lines)

	// Make a map where:
	// - key: ID of guard
	// - value: minutes this guard is sleeping
	minutesSleptPerGuard := make(map[int]int)

	// Make a map where:
	// - key: ID of guard
	// - value: map where:
	//     - key: minute of hour this guard slept
	//     - value: times this guard slept on
	//              this minute of hour
	minutesSleptPerMinuteOfHour := make(map[int]map[int]int)

	// Keep track of the ID of the last processed guard.
	// So we can map the time of woke up and fell asleep
	// to this guard.
	var idOfLastGuard int

	// Keep track of the ID of the guard that slept the
	// most minutes in total, compared to all other
	// guards.
	var idOfGuardThatSleepsMostMinutes int

	// Keep track of the total minutes slept by the guard
	// that slept the most minutes in total, compared to
	// all other guards.
	var mostMinutesSleeping int

	// Iterate over each line in the chronologically
	// ordered slice lines and analyse each line.
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// We can expect three types of lines.
		// Each line, has a unique suffix.
		//
		// [1518-03-12 00:04] Guard #1987 begins shift
		//                               ^^^^^^^^^^^^^
		//								 unique suffix
		//
		// [1518-03-12 00:21] falls asleep
		//                  ^^^^^^^^^^^^^^
		//					unique suffix
		//
		// [1518-03-12 00:55] wakes up
		//					^^^^^^^^^^
		//				    unique suffix

		if strings.HasSuffix(line, " begins shift") {
			// Parse ID of guard
			guardID, err := strconv.Atoi(line[strings.Index(line, "#")+1 : strings.Index(line, "b")-1])

			if err != nil {
				log.Fatal(err)
			}

			// Keep track of last ID of guard,
			// so we can later match sleeping
			// minutes to this guard.
			idOfLastGuard = guardID

			// Check if guard is present in the maps:
			// - minutesSleptPerGuard;
			// - minutesSleptPerMinuteOfHour.
			if _, prs := minutesSleptPerGuard[guardID]; !prs {
				// Guard is not present in map minutesSleptPerGuard,
				// so add him.
				minutesSleptPerGuard[guardID] = 0

				// If guard is not present in map minutesSleptPerGuard,
				// it means that guard is also not present in map
				// minutesSleptPerMinuteOfHour, so add him.
				minutesSleptPerMinuteOfHour[guardID] = make(map[int]int)
			}
		} else if strings.HasSuffix(line, "] falls asleep") {
			// Parse time that guard fell asleep
			fellAsleepTime, err := parseTime(line)

			if err != nil {
				log.Fatal(err)
			}

			// Next line must be time that guard woke up
			wokeUpLine := lines[i+1]

			// Parse time that guard woke up
			wokeUpTime, err := parseTime(wokeUpLine)

			if err != nil {
				log.Fatal(err)
			}

			// Calculate total minutes that guard was sleeping
			minutesSleeping := int(wokeUpTime.Sub(fellAsleepTime).Minutes())

			// Keep track of how many minutes guard slept in total,
			// by updating the value in map minutesSleptPerGuard.
			minutesSleptPerGuard[idOfLastGuard] = minutesSleptPerGuard[idOfLastGuard] + minutesSleeping

			for i := 0; i < minutesSleeping; i++ {
				// fellAsleepTime.Minute() is the first
				// minute that guard fell asleep.
				minuteOfHour := fellAsleepTime.Minute() + i

				// Keep track how many minutes guard slept
				// on this minuteOfHour, by updating the
				// value in map minutesSleptPerMinuteOfHour.
				minutesSleptPerMinuteOfHour[idOfLastGuard][minuteOfHour] = minutesSleptPerMinuteOfHour[idOfLastGuard][minuteOfHour] + 1
			}

			// Keep track of guard that sleeps the most minutes
			if minutesSleptPerGuard[idOfLastGuard] > mostMinutesSleeping {
				idOfGuardThatSleepsMostMinutes = idOfLastGuard
				mostMinutesSleeping = minutesSleptPerGuard[idOfLastGuard]
			}
		}

		// We basically ignore the line that has
		// the suffix "] falls asleep", because
		// we already handle that line in the
		// `else if` statement.
	}

	// Minute of the hour that the guard that
	// slept most minutes slept on.
	var mostSleptMinuteOfHour int

	// Minutes that the guard that slept most
	// minutes slept on mostSleptMinute.
	var mostSleptMinutes int

	// Determine on what minute of hour the guard
	// that slept most minutes slept most on.
	for sleptMinute, sleptMinutes := range minutesSleptPerMinuteOfHour[idOfGuardThatSleepsMostMinutes] {
		if sleptMinutes > mostSleptMinutes {
			mostSleptMinutes = sleptMinutes
			mostSleptMinuteOfHour = sleptMinute
		}
	}

	// This will be our final answer
	answer := idOfGuardThatSleepsMostMinutes * mostSleptMinuteOfHour

	// Print the final answer
	fmt.Println(answer, "is the ID of the guard (", idOfGuardThatSleepsMostMinutes, ") we chose multiplied by the minute (", mostSleptMinuteOfHour, ") we chose.")
}
