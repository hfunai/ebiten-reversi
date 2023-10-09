package main

import "image/color"

const (
	screenWidth  = 800
	screenHeight = 840
	boardWidth   = 800
	boardHeight  = 800
	numCells     = 8 * 8
	numCellsX    = 8
	numCellsY    = 8
	stoneRadius  = 45
)

var (
	boardColor      color.Color = color.NRGBA{0x2f, 0x9d, 0x33, 0xff}
	boardLineColor  color.Color = color.Black
	stoneColorBlack color.Color = color.Black
	stoneColorWhite color.Color = color.White
)

const (
	PlayerBlack = iota
	PlayerWhite
	PlayerNone
)

const (
	GameStateTitle = iota
	GameStatePlay
	GameStateEnd
)

type Shift struct {
	Row int
	Col int
}

type Shifts []*Shift

type Move struct {
	Row int
	Col int
}

type Moves []*Move

type Flippable struct {
	Row int
	Col int
}

type Flippables []*Flippable
