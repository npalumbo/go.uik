package uik

import (
	"code.google.com/p/freetype-go/freetype/truetype"
	"code.google.com/p/draw2d/draw2d"
)

func init() {
	font, err := truetype.Parse(luxisr_ttf())
	if err != nil {
		// TODO: log error
		println("error!")
		println(err.Error())
	}

	draw2d.RegisterFont(draw2d.FontData{"luxi", draw2d.FontFamilyMono, draw2d.FontStyleBold|draw2d.FontStyleItalic}, font)
}
