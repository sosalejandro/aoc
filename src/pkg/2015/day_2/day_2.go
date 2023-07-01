package day_2

import (
	"fmt"
	"github.com/sosalejandro/aoc/src/pkg/helpers"
	"sync"
)

func Run() {
	input := helpers.ReadInput("./src/pkg/2015/day_2/input.txt")
	boxes := parseInput(input)
	fmt.Println("Day 2, Part 1:", Part1(boxes))
	fmt.Println("Day 2, Part 2:", Part2(boxes))
}

// Part1 returns the solution for part 1 of day 2 of the advent of code.
func Part1(boxes []*PresentBox) int {
	return getTotalSurfaceArea(boxes, 50)
}

// Part2 returns the solution for part 2 of day 2 of the advent of code.
func Part2(boxes []*PresentBox) int {
	return getTotalRibbonLength(boxes, 50)
}

func parseInput(input string) []*PresentBox {
	var boxes []*PresentBox
	for _, line := range helpers.GetLines(input) {
		boxes = append(boxes, ParsePresentBox(line))
	}
	return boxes
}

// ParsePresentBox parses a string of dimensions into a PresentBox
func ParsePresentBox(dimensions string) *PresentBox {
	var length, width, height int
	_, err := fmt.Sscanf(dimensions, "%dx%dx%d", &length, &width, &height)
	if err != nil {
		fmt.Println(err)
	}
	return CreatePresentBox(length, width, height)
}

// PresentBox is a box that contains a present, it has length, width, and height
type PresentBox struct {
	Length int
	Width  int
	Height int
}

// GetSurfaceArea returns the surface area of a PresentBox
func (p *PresentBox) GetSurfaceArea() int {
	return 2*p.Length*p.Width + 2*p.Width*p.Height + 2*p.Height*p.Length
}

// GetSmallestSideArea returns the area of the smallest side of a PresentBox
func (p *PresentBox) GetSmallestSideArea() int {
	return helpers.MinInt(p.Length*p.Width, p.Width*p.Height, p.Height*p.Length)
}

// GetWrappingPaperArea returns the total wrapping paper area for a PresentBox
func (p *PresentBox) GetWrappingPaperArea() int {
	return p.GetSurfaceArea() + p.GetSmallestSideArea()
}

func (p *PresentBox) GetSmallestPerimeter() int {
	return helpers.MinInt(2*p.Length+2*p.Width, 2*p.Width+2*p.Height, 2*p.Height+2*p.Length)
}

// GetVolume returns the cubit feet of volume of a PresentBox
func (p *PresentBox) GetVolume() int {
	return p.Length * p.Width * p.Height
}

// CreatePresentBox creates a PresentBox from length, width, and height
func CreatePresentBox(length int, width int, height int) *PresentBox {
	return &PresentBox{length, width, height}
}

// RibbonCalculator calculates the amount of ribbon needed to wrap a PresentBox
type RibbonCalculator struct {
}

// GetRibbonLength returns the length of ribbon needed to wrap a PresentBox
func (r *RibbonCalculator) GetRibbonLength(box *PresentBox) int {
	return box.GetVolume() + box.GetSmallestPerimeter()
}

// getTotalSurfaceArea returns the total surface area of a slice of PresentBoxes
// batchSize is the number of boxes to process concurrently
// if batchSize is 0, all boxes will be processed concurrently
func getTotalSurfaceArea(boxes []*PresentBox, batchSize int) int {
	var total int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < len(boxes); i += batchSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + batchSize
			if end > len(boxes) {
				end = len(boxes)
			}
			batchTotal := 0
			for j := start; j < end; j++ {
				box := boxes[j]
				wrappingPaperArea := box.GetWrappingPaperArea()
				batchTotal += wrappingPaperArea
			}
			mu.Lock()
			total += batchTotal
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	return total
}

// getTotalRibbonLength returns the total length of ribbon needed to wrap a slice of PresentBoxes
// batchSize is the number of boxes to process concurrently
// if batchSize is 0, all boxes will be processed concurrently
func getTotalRibbonLength(boxes []*PresentBox, batchSize int) int {
	var total int
	var wg sync.WaitGroup
	var mu sync.Mutex
	ribbonCalculator := &RibbonCalculator{}

	for i := 0; i < len(boxes); i += batchSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + batchSize
			if end > len(boxes) {
				end = len(boxes)
			}
			batchTotal := 0
			for j := start; j < end; j++ {
				box := boxes[j]
				ribbonLength := ribbonCalculator.GetRibbonLength(box)
				batchTotal += ribbonLength
			}
			mu.Lock()
			total += batchTotal
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	return total
}

// unmeasured code

//func getTotalRibbonLength(boxes []*PresentBox, batchSize int) int {
//	ribbonCalculator := &RibbonCalculator{}
//	return getTotalFromBatch(boxes, batchSize, func(box *PresentBox) int {
//		ribbonLength := ribbonCalculator.GetRibbonLength(box)
//		return ribbonLength
//	})
//}
//
//func getTotalFromBatch(boxes []*PresentBox, batchSize int, callback func(*PresentBox) int) int {
//	var total int
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//
//	for i := 0; i < len(boxes); i += batchSize {
//		wg.Add(1)
//		go func(start int) {
//			defer wg.Done()
//			end := start + batchSize
//			if end > len(boxes) {
//				end = len(boxes)
//			}
//			batchTotal := 0
//			for j := start; j < end; j++ {
//				box := boxes[j]
//				batchTotal += callback(box)
//			}
//			mu.Lock()
//			total += batchTotal
//			mu.Unlock()
//		}(i)
//	}
//
//	wg.Wait()
//
//	return total
//}
