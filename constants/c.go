package constants

import "image/color"

const (
	Width       int = 250
	Length      int = 600
	BlockSize   int = 25
	BlockPading int = 5
	HeadPading  int = 70
)

const (
	LeftArrowKey  int32 = 263
	RightArrowKey int32 = 262
	EnterKey      int32 = 257
)

var (
	BgColor   = color.RGBA{109, 120, 92, 1}
	LightGray = color.RGBA{0, 0, 20, 40}
)
