package main

import (
	"github.com/algosup/game"
)

var bitmaps []game.Bitmap

const pacman = 0
const border = 1
const dot = 2
const redGhost = 3

type direction int

const (
	left direction = iota
	right
	up
	down
)

func init() {
	files := []string{
		"pac-right1",
		"border",
		"dot",
		"ghost",
	}

	for _, v := range files {
		bitmap, e := game.LoadBitmap(v + ".png")
		if e != nil {
			panic(e)
		}
		bitmaps = append(bitmaps, bitmap)
	}

	makePlayground()
}
