package main

import (
	"fmt"
	"math"
	"os"

	"golang.org/x/term"
)

type TermSize struct {
	Width  int
	Height int
}

type CelestialBody struct {
	mass    float32
	gravity float32
	radius  int
}

type Cell struct {
	Char  rune
	Color string
}

type ScreenBuffer struct {
	Width   int
	Height  int
	Current [][]Cell
	Next    [][]Cell
}

func GetTermSize() (width int, height int) {
	var err error

	width, height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Printf("+ Error getting terminal size: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("+ Terminal Width (col): %d\n", width)
	fmt.Printf("+ Terminal Height (row): %d\n", height)

	return
}

func CalculateFormattingPositions() {
	width, height := GetTermSize()
	roundedWidth := math.Round(float64(width) / 3.0)
	clockWidth := int(roundedWidth)
	galaxyWidth := width - clockWidth
	fmt.Printf("++ clock width: %d \n++ galaxy width: %d\n++ height %d\n", clockWidth, galaxyWidth, height)
}

func NewScreenBuffer(width, height int) *ScreenBuffer {
	current := make([][]Cell, height)
	next := make([][]Cell, height)
	for i := 0; i < height; i++ {
		current[i] = make([]Cell, width)
		next[i] = make([]Cell, width)
		for j := 0; j < width; j++ {
			current[i][j] = Cell{Char: ' '}
			next[i][j] = Cell{Char: ' '}
		}
	}
	return &ScreenBuffer{Width: width, Height: height, Current: current, Next: next}
}

func (sb *ScreenBuffer) DrawCell(x, y int, char rune) {
	if x >= 0 && x < sb.Width && y >= 0 && y < sb.Height {
		sb.Next[y][x] = Cell{Char: char}
	}
}
