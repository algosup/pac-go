package main

import "github.com/algosup/game"

var bitmaps []game.Bitmap

const pixelWidth = 500
const pixelHeight = 480
const pixelSize = 20

var playground [][]rune

func makePlayground(template string) {
	var row []rune
	playground = nil
	for _, r := range template {
		if r == '\n' {
			playground = append(playground, row)
			row = nil
		} else {
			row = append(row, r)
		}
	}
}

const border = 0
const dot = 1
const ghost = 2
const pacRight = 3
const pacLeft = 5
const pacDown = 7
const pacUp = 9

func init() {

	var names = []string{
		"border.png",
		"dot.png",
		"ghost.png",
		"pac-right1.png",
		"pac-right2.png",
		"pac-left1.png",
		"pac-left2.png",
		"pac-down1.png",
		"pac-down2.png",
		"pac-up1.png",
		"pac-up2.png",
	}

	for _, s := range names {
		b, e := game.LoadBitmap(s)
		if e != nil {
			panic(e)
		}
		bitmaps = append(bitmaps, b)
	}
}
