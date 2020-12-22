package main

import (
	"math/rand"
	"time"

	"github.com/algosup/game"
	"github.com/algosup/game/key"
)

type sprite struct {
	x          int
	y          int
	deltaX     int
	deltaY     int
	nextDeltaX int
	nextDeltaY int
	bitmap     int
}

var pac = &sprite{
	bitmap: pacRight,
}
var ghosts = []*sprite{
	{
		bitmap: ghost,
	},
	{
		bitmap: ghost,
	},
	{
		bitmap: ghost,
	},
	{
		bitmap: ghost,
	},
	{
		bitmap: ghost,
	},
	{
		bitmap: ghost,
	},
}

var frame = 0

func place(s *sprite) {
	h := len(playground)
	w := len(playground[1])

	for {
		y := rand.Intn(h-4) + 2
		x := rand.Intn(w-4) + 2
		if playground[y][x] != 'x' {
			s.x = pixelSize * x
			s.y = pixelSize * y
			return
		}
	}
}

func isEmpty() bool {
	for _, row := range playground {
		for _, r := range row {
			if r == '.' {
				return false
			}
		}
	}

	return true
}

func drawPlayground(surface game.Surface) {
	for i, row := range playground {
		for j, r := range row {
			if r == 'x' {
				game.DrawBitmap(surface, j*pixelSize, i*pixelSize, bitmaps[border])
			}
			if r == '.' {
				game.DrawBitmap(surface, j*pixelSize, i*pixelSize, bitmaps[dot])
			}
		}
	}
}

func drawSprite(surface game.Surface, s *sprite) {
	game.DrawBitmap(surface, s.x, s.y, bitmaps[s.bitmap])
}

func drawPac(surface game.Surface) {
	if pac.deltaX == 1 {
		pac.bitmap = pacRight
	}
	if pac.deltaX == -1 {
		pac.bitmap = pacLeft
	}
	if pac.deltaY == 1 {
		pac.bitmap = pacDown
	}
	if pac.deltaY == -1 {
		pac.bitmap = pacUp
	}

	game.DrawBitmap(surface, pac.x, pac.y, bitmaps[pac.bitmap+(frame/10)%2])
}
func draw(surface game.Surface) {
	frame++
	if isEmpty() {
		makePlayground(start2)
	}
	if key.IsPressed(key.Left) {
		pac.nextDeltaX = -1
		pac.nextDeltaY = 0
	}
	if key.IsPressed(key.Right) {
		pac.nextDeltaX = 1
		pac.nextDeltaY = 0
	}
	if key.IsPressed(key.Up) {
		pac.nextDeltaX = 0
		pac.nextDeltaY = -1
	}
	if key.IsPressed(key.Down) {
		pac.nextDeltaX = 0
		pac.nextDeltaY = 1
	}

	drawPlayground(surface)
	drawPac(surface)
	for _, s := range ghosts {
		drawSprite(surface, s)
	}

	for _, s := range ghosts {
		if isTouchingPac(s) {
			game.DrawText(surface, "GAMER OVER", 1, 1)
			return
		}
		move(s)
		stopIfBlocked(s)
		tryToTurn(s)

		if s.deltaX == 0 && s.deltaY == 0 {
			switch rand.Intn(4) {
			case 0:
				s.nextDeltaX = 1
				s.nextDeltaY = 0
			case 1:
				s.nextDeltaX = -1
				s.nextDeltaY = 0
			case 2:
				s.nextDeltaX = 0
				s.nextDeltaY = 1
			case 3:
				s.nextDeltaX = 0
				s.nextDeltaY = -1
			}

		}

		if isAligned(pac) {
			playground[pac.y/pixelSize][pac.x/pixelSize] = ' '
		}
	}

	move(pac)

	stopIfBlocked(pac)
	tryToTurn(pac)
}

func isTouchingPac(s *sprite) bool {
	if pac.x < s.x && s.x-pac.x > 16 {
		return false
	}
	if pac.x > s.x && pac.x-s.x > 16 {
		return false
	}
	if pac.y < s.y && s.y-pac.y > 16 {
		return false
	}
	if pac.y > s.y && pac.y-s.y > 16 {
		return false
	}
	return true
}

func move(s *sprite) {
	s.x = s.x + s.deltaX
	s.y = s.y + s.deltaY
}

func isAligned(s *sprite) bool {
	if s.x%pixelSize != 0 {
		return false
	}
	if s.y%pixelSize != 0 {
		return false
	}
	return true
}

func stopIfBlocked(s *sprite) {
	if isAligned(s) == false {
		return
	}
	var x = s.x / pixelSize
	var y = s.y / pixelSize
	if playground[y+s.deltaY][x+s.deltaX] == 'x' {
		s.deltaX = 0
		s.deltaY = 0
	}
}

func tryToTurn(s *sprite) {
	if isAligned(s) == false {
		return
	}
	var x = s.x / pixelSize
	var y = s.y / pixelSize
	if playground[y+s.nextDeltaY][x+s.nextDeltaX] == 'x' {
		return
	}
	s.deltaX = s.nextDeltaX
	s.deltaY = s.nextDeltaY
}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	makePlayground(start1)
	place(pac)
	for _, s := range ghosts {
		for {
			place(s)
			if !isTouchingPac(s) {
				break
			}
		}
	}

	game.Run("Pac-Go", pixelWidth, pixelHeight, draw)
}
