package game

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nithinputhenveettil/brick-game-car-racing/utils"
)

type animation struct {
	IsActive bool
	index    int
	isUp     bool
	l        int
	w        int
}

func (a *animation) reset() {
	a.isUp = true
	a.IsActive = true
	a.index = a.l - 1
}

func (a *animation) Animate() {
	if a.isUp {
		if a.index == 0 {
			a.isUp = false
		} else {
			a.index -= 1
		}
		for i := a.l - 1; i >= a.index; i-- {
			for j := 0; j < a.w; j++ {
				utils.DrawSingleBlock(i, j, rl.Black)
			}
		}
	} else {
		if a.index == a.l-1 {
			a.isUp = true
			a.IsActive = false
		} else {
			a.index += 1
		}
		for i := a.index; i < a.l; i++ {
			for j := 0; j < a.w; j++ {
				utils.DrawSingleBlock(i, j, rl.Black)
			}
		}
	}
	time.Sleep(25 * time.Millisecond)
}
