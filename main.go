package main

import (
	"image/color"
	"math/rand"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width       int = 250
	length      int = 600
	blockSize   int = 25
	blockPading int = 5
	headPading  int = 70
)

const (
	leftArrowKey  int32 = 263
	rightArrowKey int32 = 262
	enterKey      int32 = 257
)

var (
	bgColor   = color.RGBA{109, 120, 92, 1}
	lightGray = color.RGBA{0, 0, 20, 40}
)

type obstacle struct {
	counter int
	pading  int
}

type game struct {
	store           [][]bool
	start           int
	w               int
	l               int
	gameOver        bool
	border          int
	level           int
	score           int
	currentObstacle *obstacle
	c               *car
}

type car struct {
	store   [][]int
	isRight bool
}

func (c *car) reset(l int) {
	c.isRight = false
	c.store = [][]int{
		{l - 1, 2},
		{l - 1, 4},
		{l - 2, 3},
		{l - 3, 2},
		{l - 3, 3},
		{l - 3, 4},
		{l - 4, 3},
	}
}

func (g *game) reset() {
	g.gameOver = false
	g.level = 1
	g.score = 0
	g.w = int(width / blockSize)
	g.l = int(length / blockSize)
	g.store = make([][]bool, g.l)
	for i := range g.store {
		g.store[i] = make([]bool, g.w)
		g.border = (g.border + 1) % 4
		g.store[i][0], g.store[i][g.w-1] = g.border > 0, g.border > 0

	}
	g.border = 3
	g.currentObstacle = &obstacle{
		counter: 0,
		pading:  3,
	}
	g.c = &car{}
	g.c.reset(g.l)
}

func (g *game) nextTick() {
	g.start += g.level
	if g.start > blockSize {
		gg := make([]bool, g.w)
		g.border = (g.border + 1) % 4
		gg[0], gg[g.w-1] = g.border > 0, g.border > 0
		g.currentObstacle.counter = (g.currentObstacle.counter + 1) % 10
		c := g.currentObstacle.counter
		p := g.currentObstacle.pading
		if c == 0 {
			g.currentObstacle.pading = 3
			if x := rand.Intn(2); x == 1 {
				g.currentObstacle.pading = 0
			}
		} else if c == 1 {
			gg[2+p], gg[4+p] = true, true
		} else if c == 2 || c == 4 {
			gg[3+p] = true
		} else if c == 3 {
			gg[2+p], gg[3+p], gg[4+p] = true, true, true
		}
		g.store = append([][]bool{gg}, g.store[0:g.l-1]...)
		g.start = 0
	}
}

func (g *game) litsenKeyboardEvents() {
	if rl.IsKeyPressed(leftArrowKey) {
		if g.c.isRight && !g.gameOver {
			g.c.isRight = false
			for _, e := range g.c.store {
				e[1] -= 3
			}
			g.score++
			if ((g.score + 1) % 10) == 0 {
				g.level++
			}
		}
	}
	if rl.IsKeyPressed(rightArrowKey) {
		if !g.c.isRight && !g.gameOver {
			g.c.isRight = true
			for _, e := range g.c.store {
				e[1] += 3
			}
			g.score++
			if ((g.score + 1) % 10) == 0 {
				g.level++
			}
		}
	}
	if rl.IsKeyDown(enterKey) {
		if g.gameOver {
			g.reset()
			g.c.reset(g.l)
		}
	}
}

func (c *car) drawSingleBlock(i int, j int, col color.RGBA) {

	x := int32(j * blockSize)
	y := int32(i*blockSize) + int32(headPading)
	rl.DrawLine(x, y, x+int32(blockSize), y, col)
	rl.DrawLine(x, y+int32(blockSize), x+int32(blockSize), y+int32(blockSize), col)
	rl.DrawLine(x, y, x, y+int32(blockSize), col)
	rl.DrawLine(x+int32(blockSize), y, x+int32(blockSize), y+int32(blockSize), col)
	rl.DrawRectangle(x+int32(blockPading), y+int32(blockPading), int32(blockSize-(blockPading*2)), int32(blockSize-(blockPading*2)), col)
}

func (c *car) draw() {
	for _, e := range c.store {
		c.drawSingleBlock(e[0], e[1], rl.Black)
	}
}

func (g *game) draw() {
	rl.DrawText("Brick Game Car Racing!", 160, 30, 20, rl.Black)
	rl.DrawRectangle(0, 0, int32(width)+int32(width), int32(blockSize)/2, rl.Black)
	rl.DrawRectangle(int32(width+width-blockSize/2), 0, int32(blockSize)/2, int32(headPading), rl.Black)
	rl.DrawRectangle(0, 0, int32(blockSize)/2, int32(headPading), rl.Black)
	rl.DrawRectangle(0, int32(headPading), int32(width)+int32(width), int32(blockSize), rl.Black)
	for i := 0; i < g.l; i++ {
		for j := 0; j < g.w; j++ {
			x := int32(j * blockSize)
			y := int32(i*blockSize) + int32(headPading)
			col := lightGray
			if g.store[i][j] || (i != 0 && g.store[i-1][j]) {
				col = rl.Black
			}
			rl.DrawLine(x, y+int32(g.start), x+int32(blockSize), y+int32(g.start), col)
			col = lightGray
			if g.store[i][j] || (j != 0 && g.store[i][j-1]) {
				col = rl.Black
			}
			rl.DrawLine(x, y+int32(g.start), x, y+int32(blockSize)+int32(g.start), col)
			col = lightGray
			if g.store[i][j] {
				col = rl.Black
			}
			rl.DrawRectangle(x+int32(blockPading), y+int32(blockPading)+int32(g.start), int32(blockSize-(blockPading*2)), int32(blockSize-(blockPading*2)), col)
		}
	}
	rl.DrawLine(int32(width), int32(headPading), int32(width), int32(length), rl.Black)
	rl.DrawLine(int32(width), int32(length/2)+int32(headPading), int32(width+width), int32(length/2)+int32(headPading), rl.Black)
	rl.DrawText("Score", int32(width+60), int32(headPading)+60, 40, rl.Black)
	rl.DrawText(strconv.Itoa(g.score), int32(width+90), int32(headPading)+150, 80, rl.Black)
	rl.DrawText("Level", int32(width+60), int32(length/2)+int32(headPading)+60, 40, rl.Black)
	rl.DrawText(strconv.Itoa(g.level), int32(width+90), int32(length/2)+int32(headPading)+150, 80, rl.Black)

	g.c.draw()
}

func (g *game) checkGameOver() {
	for _, e := range g.c.store {
		if g.store[e[0]][e[1]] {
			g.gameOver = true
			break
		}
	}
}

func main() {

	g := &game{}
	g.reset()

	rl.InitWindow(int32(width)+int32(width), int32(length)+int32(headPading), "Brick Car Racing")
	rl.SetTargetFPS(120)
	for !rl.WindowShouldClose() {
		g.litsenKeyboardEvents()
		rl.BeginDrawing()
		rl.ClearBackground(bgColor)
		if g.gameOver {
			g.draw()
			rl.DrawText("Game Over!!", 10, 150, 85, rl.White)
			rl.DrawText("Press enter key to continue!", int32(width)-60, int32(length-60), 20, rl.White)
		} else {
			g.nextTick()
			g.checkGameOver()
			g.draw()
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
