package main

import (
	"flag"
	"fmt"
	aoc_2015 "github.com/sosalejandro/aoc/src/pkg/2015"
)

type DayFunc func()

func CreateDayFuncArray(fns []func()) []DayFunc {
	dayFuncs := make([]DayFunc, len(fns))
	for i, fn := range fns {
		dayFuncs[i] = fn
	}
	return dayFuncs
}

// create a cmd program that will run the code for the day and year specified by the user
func main() {
	// Define command-line flags for the day and year
	day := flag.Int("day", 0, "The day of the Advent of Code challenge to run (1-25)")
	year := flag.Int("year", 0, "The year of the Advent of Code challenge to run (2015-2021)")

	// Parse the command-line flags
	flag.Parse()

	// Check if the day and year are valid
	if *day < 1 || *day > 25 {
		fmt.Println("Error: Invalid day. Must be between 1 and 25.")
		return
	}
	if *year < 2015 || *year > 2021 {
		fmt.Println("Error: Invalid year. Must be between 2015 and 2021.")
		return
	}

	// Run the code for the specified day and year
	fmt.Printf("Running code for Advent of Code %d, Day %d...\n", *year, *day)
	runYear(*year, *day)
}

// runYear runs the code for the specified year
func runYear(year int, day int) {
	switch year {
	case 2015:
		runDay(year, day, CreateDayFuncArray(aoc_2015.Functions))
	case 2016:
		runDay(year, day, []DayFunc{})
	case 2017:
		runDay(year, day, []DayFunc{})
	case 2018:
		runDay(year, day, []DayFunc{})
	case 2019:
		runDay(year, day, []DayFunc{})
	case 2020:
		runDay(year, day, []DayFunc{})
	case 2021:
		runDay(year, day, []DayFunc{})
	case 2022:
		runDay(year, day, []DayFunc{})
	default:
		fmt.Printf("Error: No code found for Advent of Code %d.\n", year)
		return
	}
}

// runDay runs the code for the specified day and year
func runDay(year int, day int, dayFuncs []DayFunc) {
	if day < 1 || day > 25 {
		fmt.Printf("Error: Invalid day. Must be between 1 and 25.\n")
		return
	}

	if len(dayFuncs) < day {
		fmt.Printf("Error: No code found for Advent of Code %d, Day %d.\n", year, day)
		return
	}

	dayFunc := dayFuncs[day-1]
	if dayFunc == nil {
		fmt.Printf("Error: No code found for Advent of Code %d, Day %d.\n", year, day)
		return
	}

	fmt.Printf("Running code for Advent of Code %d, Day %d...\n", year, day)
	dayFunc()
}
