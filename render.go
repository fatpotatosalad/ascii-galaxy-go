package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

/*  _______________________________
 * |  ___________   _____________  |
 * | |           | |             | |
 * | |           | |   Galaxy    | |
 * | |   Clock   | | Simulation  | |
 * | |  Polygon  | |   Polygon   | |
 * | |           | |             | |
 * | |___________| |_____________| |
 * |_______________________________|
 *
 * The primary structure for the clock + galaxy sim
 * Polygon can have starting x, y  ending x,y calcualted based on term-size.
 * Polygon can also have a Border with Thickness and RGB values as a [3]int.
 * The clock polygon will occupy 1/3 of the total terminal width.
 */

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
	Color [3]int
}

type Border struct {
	Thickness int
	Color     [3]int
}

type Polygon struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
	Border Border
}

type ScreenBuffer struct {
	Width    int
	Height   int
	Current  []Cell
	Next     []Cell
	Polygons []Polygon
}

func GetTermSize() (width int, height int) {
	var err error
	width, height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Printf("+ Error getting terminal size: %v\n", err)
		os.Exit(1)
	}
	return
}

func CalculatePositionAfterBorders(p Polygon) Polygon {
	return Polygon{
		StartX: p.StartX + p.Border.Thickness,
		StartY: p.StartY + p.Border.Thickness,
		EndX:   p.EndX + p.Border.Thickness,
		EndY:   p.EndY + p.Border.Thickness,
		Border: p.Border,
	}
}

func CalculateFormattingPositions() []Polygon {
	// Calculate where the clock and the galaxy sim should be
	// relative to the terminal origin.
	width, height := GetTermSize()
	clockWidth := width / 3
	clock := Polygon{
		StartX: 0,
		StartY: 0,
		EndX:   clockWidth,
		EndY:   height,
		Border: Border{
			Thickness: 1,
			Color:     [3]int{255, 255, 255},
		},
	}

	galaxy := Polygon{
		StartX: clockWidth,
		StartY: 0,
		EndX:   width,
		EndY:   height,
		Border: Border{
			Thickness: 1,
			Color:     [3]int{255, 255, 255},
		},
	}

	return []Polygon{clock, galaxy}
}

func NewScreenBuffer(width, height int) *ScreenBuffer {
	size := width * height

	return &ScreenBuffer{
		Width:   width,
		Height:  height,
		Current: make([]Cell, size),
		Next:    make([]Cell, size),
	}
}

func (sb *ScreenBuffer) DrawCell(x, y int, char rune) {
	//to be implemented
}

func (sb *ScreenBuffer) FlushCells() {
	// Clears out the array with empty values.
	for i := 0; i < len(sb.Next); i++ {
		sb.Next[i] = Cell{Char: ' ', Color: [3]int{255, 255, 255}}
	}
}

func DrawBorders(p Polygon) {
	// to be implemented
}
