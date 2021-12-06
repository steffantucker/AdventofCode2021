package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

type line struct {
	x1, x2 int
	y1, y2 int
}

func getLines(in []string) []line {
	lines := []line{}
	for _, row := range in {
		split1 := strings.Split(row, " -> ")
		asplit := strings.Split(split1[0], ",")
		bsplit := strings.Split(split1[1], ",")
		x1, _ := strconv.Atoi(asplit[0])
		y1, _ := strconv.Atoi(asplit[1])
		x2, _ := strconv.Atoi(bsplit[0])
		y2, _ := strconv.Atoi(bsplit[1])
		lines = append(lines, line{
			x1: x1,
			x2: x2,
			y1: y1,
			y2: y2,
		})
	}
	return lines
}

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	fmt.Printf("test 1: %v\ninput 1: %v\n", puzzle1(testinput), puzzle1(inputinput))
	fmt.Printf("test 2: %v\ninput 2: %v\n", puzzle2(testinput), puzzle2(inputinput))
}

func puzzle1(input []string) (out int) {
	lines := getLines(input)
	grid := make(map[string]int)
	for _, l := range lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			xstart, xend := l.x1, l.x2
			ystart, yend := l.y1, l.y2
			if xstart > xend {
				xstart, xend = xend, xstart
			}
			if ystart > yend {
				ystart, yend = yend, ystart
			}
			for x := xstart; x <= xend; x++ {
				for y := ystart; y <= yend; y++ {
					hash := fmt.Sprintf("%v,%v", x, y)
					grid[hash]++
				}
			}
		}
	}
	count := 0
	for _, point := range grid {
		//fmt.Printf("%v: %v\n", xy, point)
		if point >= 2 {
			count++
		}
	}
	return count
}

func puzzle2(input []string) int {

	lines := getLines(input)
	grid := make(map[string]int)
	for _, l := range lines {
		xstart, xend := l.x1, l.x2
		ystart, yend := l.y1, l.y2
		if xstart > xend {
			xstart, xend = xend, xstart
		}
		if ystart > yend {
			ystart, yend = yend, ystart
		}
		slope := math.Abs((float64(l.y2) - float64(l.y1)) / (float64(l.x2) - float64(l.x1)))
		fmt.Println(slope)
		if l.x1 == l.x2 || l.y1 == l.y2 {
			for x := xstart; x <= xend; x++ {
				for y := ystart; y <= yend; y++ {
					hash := fmt.Sprintf("%v,%v", x, y)
					grid[hash]++
				}
			}
		} else if slope == 1 {
			drawSlope(l.x1, l.x2, l.y1, l.y2, grid)
		}
	}
	count := 0
	for _, point := range grid {
		//fmt.Printf("%v: %v\n", xy, point)
		if point >= 2 {
			count++
		}
	}
	return count
}

func draw(graph [10][10]int) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			fmt.Printf("%v ", graph[y][x])
		}
		fmt.Println()
	}
}

func drawSlope(x1, x2, y1, y2 int, grid map[string]int) {
	bleh := func(x, y int) {
		hash := fmt.Sprintf("%v,%v", x, y)
		grid[hash]++
	}
	if x1 < x2 && y1 < y2 {
		for x, y := x1, y1; x <= x2 && y <= y2; x, y = x+1, y+1 {
			bleh(x, y)
		}
	}
	if x1 > x2 && y1 < y2 {
		for x, y := x1, y1; x >= x2 && y <= y2; x, y = x-1, y+1 {
			bleh(x, y)
		}
	}
	if x1 < x2 && y1 > y2 {
		for x, y := x1, y1; x <= x2 && y >= y2; x, y = x+1, y-1 {
			bleh(x, y)
		}
	}
	if x1 > x2 && y1 > y2 {
		for x, y := x1, y1; x >= x2 && y >= y2; x, y = x-1, y-1 {
			bleh(x, y)
		}
	}
}
