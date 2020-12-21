package main

import (
	"math/rand"

	"github.com/algosup/game"
	"github.com/algosup/game/key"
)

const pixelWidth = 500
const pixelHeight = 480
const pixelSize = 20

const start = `
 xxxxxxxxxxxxxxxxxxxxxxx
 x..........x..........x
 x.xxx.xxxx.x.xxxx.xxx.x
 x.xxx.xxxx.x.xxxx.xxx.x
 x.....................x
 x.xxx.x.xxxxxxx.x.xxx.x
 x.....x....x....x.....x
 xxxxx.xxxx.x.xxxx.xxxxx
 xxxxx.x.........x.xxxxx
 xxxxx.x.xxxxxxx.x.xxxxx
 x.......x     x.......x
 xxxxx.x.xxxxxxx.x.xxxxx
 xxxxx.x.........x.xxxxx
 xxxxx.x.xxxxxxx.x.xxxxx
 x..........x..........x
 x.xxx.xxxx.x.xxxx.xxx.x
 x...x.............x...x
 xxx.x.x.xxxxxxx.x.x.xxx
 x.....x....x....x.....x
 x.xxxxxxxx.x.xxxxxxxx.x
 x.....................x
 xxxxxxxxxxxxxxxxxxxxxxx
`

var playground [][]rune

type sprite struct {
	direction     direction
	nextDirection direction
	isMoving      bool
	x             int
	y             int
	bitmapIndex   int
}

var p = &sprite{
	direction:     right,
	nextDirection: left,
	x:             2 * pixelSize,
	y:             2 * pixelSize,
	bitmapIndex:   0,
}

var ghosts []*sprite = []*sprite{
	{
		direction:   right,
		isMoving:    true,
		x:           10 * pixelSize,
		y:           13 * pixelSize,
		bitmapIndex: redGhost,
	},
	{
		direction:   right,
		isMoving:    true,
		x:           10 * pixelSize,
		y:           13 * pixelSize,
		bitmapIndex: redGhost,
	},
	{
		direction:   right,
		isMoving:    true,
		x:           10 * pixelSize,
		y:           13 * pixelSize,
		bitmapIndex: redGhost,
	},
	{
		direction:   right,
		isMoving:    true,
		x:           10 * pixelSize,
		y:           13 * pixelSize,
		bitmapIndex: redGhost,
	},
	{
		direction:   right,
		isMoving:    true,
		x:           10 * pixelSize,
		y:           13 * pixelSize,
		bitmapIndex: redGhost,
	},
}

func makePlayground() {
	playground = make([][]rune, 0)
	row := make([]rune, 0)

	for _, r := range start {
		if r == '\n' {
			playground = append(playground, row)
			row = make([]rune, 0)
		} else {
			row = append(row, r)
		}
	}
}

func drawPlayground(surface game.Surface) {
	var x = 0
	var y = 0

	for _, row := range playground {
		for _, c := range row {

			switch c {
			case 'x':
				game.DrawBitmap(surface, x, y, bitmaps[border])

			case '.':
				game.DrawBitmap(surface, x, y, bitmaps[dot])
			}

			x += pixelSize
		}
		x = 0
		y += pixelSize
	}
}

func drawSprite(surface game.Surface, s *sprite) {
	game.DrawBitmap(surface, s.x, s.y, bitmaps[s.bitmapIndex])
}

func draw(surface game.Surface) {
	drawPlayground(surface)
	drawSprite(surface, p)
	for _, g := range ghosts {
		drawSprite(surface, g)
	}

	for _, g := range ghosts {
		if p.x == g.x && p.y == g.y {
			game.DrawText(surface, "GAME OVER", 0, 0)
			return
		}
	}

	move(p)
	collect()

	for _, g := range ghosts {
		move(g)

		if !g.isMoving {
			g.nextDirection = direction(rand.Intn(4))
		}
	}

	if key.IsPressed(key.Left) {
		p.nextDirection = left
	}
	if key.IsPressed(key.Right) {
		p.nextDirection = right
	}
	if key.IsPressed(key.Up) {
		p.nextDirection = up
	}
	if key.IsPressed(key.Down) {
		p.nextDirection = down
	}
}

func move(p *sprite) {
	if mustStop(p) {
		p.isMoving = false
	}
	if p.isMoving {
		switch p.direction {
		case left:
			p.x--
		case right:
			p.x++
		case up:
			p.y--
		case down:
			p.y++
		}
	}

	if canChangeDirection(p) {
		p.direction = p.nextDirection
		p.isMoving = true
	}
}

func collect() {
	if p.x%pixelSize != 0 {
		return
	}

	if p.y%pixelSize != 0 {
		return
	}

	var r = p.y / pixelSize
	var c = p.x / pixelSize

	playground[r][c] = ' '
}

func canChangeDirection(p *sprite) bool {
	if p.x%pixelSize != 0 {
		return false
	}

	if p.y%pixelSize != 0 {
		return false
	}

	var r = p.y / pixelSize
	var c = p.x / pixelSize

	switch p.nextDirection {
	case left:
		c--
	case right:
		c++
	case up:
		r--
	case down:
		r++
	}

	if playground[r][c] == 'x' {
		return false
	} else {
		return true
	}
}

func mustStop(p *sprite) bool {
	if p.x%pixelSize != 0 {
		return false
	}

	if p.y%pixelSize != 0 {
		return false
	}

	var r = p.y / pixelSize
	var c = p.x / pixelSize

	switch p.direction {
	case left:
		c--
	case right:
		c++
	case up:
		r--
	case down:
		r++
	}

	if playground[r][c] == 'x' {
		return true
	} else {
		return false
	}
}

func main() {

	game.Run("Pac-Man", pixelWidth, pixelHeight, draw)

}
