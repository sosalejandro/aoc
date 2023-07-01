package day_1

import (
	"fmt"
	"github.com/sosalejandro/aoc/src/pkg/helpers"
)

func Run() {
	input := helpers.ReadInput("./src/pkg/2015/day_1/input.txt")
	fmt.Println("Day 1, Part 1:", Part1(input))
	fmt.Println("Day 1, Part 2:", Part2(input))
}

// Part1 returns the solution for part 1 of day 1 of the advent of code.
func Part1(input string) int {
	f := CreateFloor()
	f.ReadFloorInstructions1(input)
	return int(*f)
}

// Part2 returns the solution for part 2 of day 1 of the advent of code.
func Part2(input string) int {
	f := CreateFloor()
	result := f.ReadFloorInstructions2(input)
	return result
}

type Floor int16

func CreateFloor() *Floor {
	floor := Floor(0)
	return &floor
}

func (f *Floor) up() {
	*f++
}

func (f *Floor) down() {
	*f--
}

func (f *Floor) ReadFloorInstructions1(instructions string) {
	for _, instruction := range instructions {

		switch instruction {
		case '(':
			f.up()
		case ')':
			f.down()
		}
	}
}

func (f *Floor) ReadFloorInstructions2(instructions string) int {
	for i, instruction := range instructions {
		switch instruction {
		case '(':
			f.up()
		case ')':
			f.down()
		}

		if *f == -1 {
			return i + 1
		}
	}

	return 0
}
