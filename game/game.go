package game

import (
	"math/rand"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/nithinputhenveettil/brick-game-car-racing/constants"
)

type Obstacle struct {
	counter int
	pading  int
}

type Game struct {
	store           [][]bool
	start           int
	w               int
	l               int
	GameOver        bool
	border          int
	level           int
	score           int
	currentObstacle *Obstacle
	c               *car
	Animation       *animation
}

func (g *Game) Reset() {
	g.GameOver = false
	g.level = 1
	g.score = 0
	g.w = int(c.Width / c.BlockSize)
	g.l = int(c.Length / c.BlockSize)
	g.store = make([][]bool, g.l)
	for i := range g.store {
		g.store[i] = make([]bool, g.w)
		g.border = (g.border + 1) % 4
		g.store[i][0], g.store[i][g.w-1] = g.border > 0, g.border > 0

	}
	g.border = 3
	g.currentObstacle = &Obstacle{
		counter: 0,
		pading:  3,
	}
	g.c = &car{l: g.l}
	g.Animation = &animation{l: g.l, w: g.w}
	g.Animation.reset()
	g.c.reset()
}

func (g *Game) NextTick() {
	g.start += g.level
	if g.start > c.BlockSize {
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
		} else if c == 4 {
			gg[2+p], gg[4+p] = true, true
		} else if c == 1 || c == 3 {
			gg[3+p] = true
		} else if c == 2 {
			gg[2+p], gg[3+p], gg[4+p] = true, true, true
		}
		g.store = append([][]bool{gg}, g.store[0:g.l-1]...)
		g.start = 0
	}
}

func (g *Game) LitsenKeyboardEvents() {
	if rl.IsKeyPressed(c.LeftArrowKey) {
		if g.c.isRight && !g.GameOver && !g.Animation.IsActive {
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
	if rl.IsKeyPressed(c.RightArrowKey) {
		if !g.c.isRight && !g.GameOver && !g.Animation.IsActive {
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
	if rl.IsKeyDown(c.EnterKey) {
		if g.GameOver {
			g.Reset()
			g.c.reset()
		}
	}
}

func (g *Game) Draw() {
	rl.DrawText("Brick Game Car Racing!", 130, 30, 20, rl.Black)
	rl.DrawRectangle(0, 0, int32(c.Width)+int32(c.Width), int32(c.BlockSize)/2, rl.Black)
	rl.DrawRectangle(int32(c.Width+c.Width-c.BlockSize/2), 0, int32(c.BlockSize)/2, int32(c.HeadPading), rl.Black)
	rl.DrawRectangle(0, 0, int32(c.BlockSize)/2, int32(c.HeadPading), rl.Black)
	rl.DrawRectangle(0, int32(c.HeadPading), int32(c.Width)+int32(c.Width), int32(c.BlockSize), rl.Black)
	for i := 0; i < g.l; i++ {
		for j := 0; j < g.w; j++ {
			x := int32(j * c.BlockSize)
			y := int32(i*c.BlockSize) + int32(c.HeadPading)
			col := c.LightGray
			if g.store[i][j] || (i != 0 && g.store[i-1][j]) {
				col = rl.Black
			}
			rl.DrawLine(x, y+int32(g.start), x+int32(c.BlockSize), y+int32(g.start), col)
			col = c.LightGray
			if g.store[i][j] || (j != 0 && g.store[i][j-1]) {
				col = rl.Black
			}
			rl.DrawLine(x, y+int32(g.start), x, y+int32(c.BlockSize)+int32(g.start), col)
			col = c.LightGray
			if g.store[i][j] {
				col = rl.Black
			}
			rl.DrawRectangle(x+int32(c.BlockPading), y+int32(c.BlockPading)+int32(g.start), int32(c.BlockSize-(c.BlockPading*2)), int32(c.BlockSize-(c.BlockPading*2)), col)
		}
	}
	rl.DrawLine(int32(c.Width), int32(c.HeadPading), int32(c.Width), int32(c.Length+c.HeadPading), rl.Black)
	rl.DrawLine(int32(c.Width), int32(c.Length/2)+int32(c.HeadPading), int32(c.Width+c.Width), int32(c.Length/2)+int32(c.HeadPading), rl.Black)
	rl.DrawText("Score", int32(c.Width+60), int32(c.HeadPading)+60, 40, rl.Black)
	rl.DrawText(strconv.Itoa(g.score), int32(c.Width+90), int32(c.HeadPading)+150, 80, rl.Black)
	rl.DrawText("Level", int32(c.Width+60), int32(c.Length/2)+int32(c.HeadPading)+60, 40, rl.Black)
	rl.DrawText(strconv.Itoa(g.level), int32(c.Width+90), int32(c.Length/2)+int32(c.HeadPading)+150, 80, rl.Black)

	g.c.draw()
}

func (g *Game) CheckGameOver() {
	for _, e := range g.c.store {
		if g.store[e[0]][e[1]] {
			g.GameOver = true
			return
		}
	}
}
