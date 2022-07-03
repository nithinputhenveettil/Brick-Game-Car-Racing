package game

import (
	"github.com/nithinputhenveettil/brick-game-car-racing/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type car struct {
	store   [][]int
	isRight bool
	l       int
}

func (c *car) draw() {
	for _, e := range c.store {
		utils.DrawSingleBlock(e[0], e[1], rl.Black)
	}
}

func (c *car) reset() {
	c.isRight = false
	c.store = [][]int{
		{c.l - 1, 2},
		{c.l - 1, 4},
		{c.l - 2, 3},
		{c.l - 3, 2},
		{c.l - 3, 3},
		{c.l - 3, 4},
		{c.l - 4, 3},
	}
}
