package utils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/nithinputhenveettil/brick-game-car-racing/constants"
)

func DrawSingleBlock(i int, j int, col color.RGBA) {

	x := int32(j * c.BlockSize)
	y := int32(i*c.BlockSize) + int32(c.HeadPading)
	rl.DrawLine(x, y, x+int32(c.BlockSize), y, col)
	rl.DrawLine(x, y+int32(c.BlockSize), x+int32(c.BlockSize), y+int32(c.BlockSize), col)
	rl.DrawLine(x, y, x, y+int32(c.BlockSize), col)
	rl.DrawLine(x+int32(c.BlockSize), y, x+int32(c.BlockSize), y+int32(c.BlockSize), col)
	rl.DrawRectangle(x+int32(c.BlockPading), y+int32(c.BlockPading), int32(c.BlockSize-(c.BlockPading*2)), int32(c.BlockSize-(c.BlockPading*2)), col)
}
