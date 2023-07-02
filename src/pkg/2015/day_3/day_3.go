package day_3

import (
	"fmt"
	"github.com/sosalejandro/aoc/src/pkg/helpers"
)

type CardinalSign string

const (
	North CardinalSign = "^"
	South CardinalSign = "v"
	East  CardinalSign = ">"
	West  CardinalSign = "<"
)

// Run runs the day 3 challenge parts 1 and 2.
func Run() {
	input := helpers.ReadInput("./src/pkg/2015/day_3/input.txt")
	fmt.Println("Day 3, Part 1:", Part1(input))
	fmt.Println("Day 3, Part 2:", Part2(input))
}

// Part1 returns the solution for part 1 of day 3 of the advent of code.
func Part1(input string) int {
	grid := CreateGrid()
	grid.ParseGrid(input)
	return len(grid.Houses)
}

// Part2 returns the solution for part 2 of day 3 of the advent of code.
func Part2(input string) int {
	grid := CreateGrid()
	grid.ParseGrid2(input)
	return len(grid.Houses)
}

// Location represents a location in the grid.
type Location struct {
	x int
	y int
}

// String returns the string representation of the location.
func (l *Location) String() string {
	return fmt.Sprintf("%d,%d", l.x, l.y)
}

// House represents a house in the grid.
type House struct {
	location Location
	Presents int
}

// CreateHouse creates a new house.
func CreateHouse(location Location) *House {
	house := House{location: location, Presents: 0}
	return &house
}

// DeliverPresents delivers presents to the house.
func (h *House) DeliverPresents() {
	h.Presents++
}

// GetLocation returns the location of the house.
func (h *House) GetLocation() string {
	return h.location.String()
}

// Grid represents the grid of houses.
type Grid struct {
	Houses map[string]*House
}

// CreateGrid creates a new grid.
func CreateGrid() *Grid {
	grid := Grid{Houses: make(map[string]*House)}
	return &grid
}

// GetHouse returns the house at the given location.
func (g *Grid) GetHouse(location Location) *House {
	return g.Houses[location.String()]
}

type Deliver interface {
	DeliverPresents(*House)
	GetLocation() *Location
}

// PackageDeliver delivers presents to the houses in the grid.
type PackageDeliver struct {
	location *Location
}

// DeliverPresents delivers presents to the houses in the grid.
func (s *PackageDeliver) DeliverPresents(h *House) {
	h.DeliverPresents()
}

// GetLocation returns the location of the package deliverer.
func (s *PackageDeliver) GetLocation() *Location {
	return s.location
}

// ParseGrid parses the grid from the input.
func (g *Grid) ParseGrid(input string) {
	currentHouse := CreateHouse(Location{0, 0})
	santa := &PackageDeliver{location: &Location{0, 0}}
	g.Houses[currentHouse.GetLocation()] = currentHouse

	for _, sign := range input {
		g.SetLocation(santa.GetLocation(), CardinalSign(sign))
		currentHouse = g.GetHouse(*santa.GetLocation())
		if currentHouse == nil {
			currentHouse = CreateHouse(*santa.GetLocation())
			g.Houses[currentHouse.GetLocation()] = currentHouse
		}
		santa.DeliverPresents(currentHouse)
	}
}

// ParseGrid2 parses the grid from the input.
func (g *Grid) ParseGrid2(input string) {
	currentHouse := CreateHouse(Location{0, 0})
	g.Houses[currentHouse.GetLocation()] = currentHouse

	santa := &PackageDeliver{location: &Location{0, 0}}
	santa.DeliverPresents(currentHouse)

	robotSanta := &PackageDeliver{location: &Location{0, 0}}
	robotSanta.DeliverPresents(currentHouse)

	deliverPackageAtLocation := func(delivery Deliver, sign CardinalSign) {
		g.SetLocation(delivery.GetLocation(), sign)
		currentHouse = g.GetHouse(*delivery.GetLocation())
		if currentHouse == nil {
			currentHouse = CreateHouse(*delivery.GetLocation())
			g.Houses[delivery.GetLocation().String()] = currentHouse
		}
		delivery.DeliverPresents(currentHouse)
	}

	for i, sign := range input {
		if ind := i + 1; ind%2 == 0 {
			deliverPackageAtLocation(robotSanta, CardinalSign(sign))
			continue
		}

		deliverPackageAtLocation(santa, CardinalSign(sign))
		continue
	}
}

// SetLocation sets the location of the house.
func (g *Grid) SetLocation(location *Location, sign CardinalSign) {
	switch sign {
	case North:
		location.y++
		break
	case South:
		location.y--
		break
	case East:
		location.x++
		break
	case West:
		location.x--
		break
	}
}
