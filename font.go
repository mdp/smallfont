package smallfont

import (
	"errors"
	"image"
	"image/color"
)

// FontSet holds the ascii table of fonts
type FontSet struct {
	Table  []byte
	Width  int
	Height int
}

// Character return the f
func (f *FontSet) Character(a byte) FontCharacter {
	c := int(a) & 255
	i := c * f.Height
	return FontCharacter{
		Map:    f.Table[i : i+f.Height],
		Width:  f.Width,
		Height: f.Height,
		idx:    0,
	}
}

// FontCharacter holds the representation for the single character
type FontCharacter struct {
	Map    []byte
	Width  int
	Height int
	idx    int
	Y      int
	X      int
	Pixel  bool
}

// NextBit grabs the next bit
func (c *FontCharacter) NextBit() bool {
	if c.idx >= int(c.Width*c.Height) {
		return false
	}
	c.X = c.idx % c.Height
	c.Y = c.idx / c.Width
	b := c.Map[c.Y]
	i := 1 << uint8(c.Width-c.X-1)
	c.Pixel = int(b)&i > 0
	c.idx = c.idx + 1
	return true
}

// Context for rasterizing text to an image
type Context struct {
	Dst    *image.RGBA
	StartX int
	StartY int
	Font   FontSet
	Color  color.Color
}

// Draw a string on an RGBA image
func (context *Context) Draw(str []byte, offsetX, offsetY int) error {
	for i, c := range str {
		fc := context.Font.Character(c)
		offsetX := offsetX + context.StartX + i*fc.Width
		offsetY := offsetY + context.StartY
		for fc.NextBit() {
			x, y := offsetX+fc.X, offsetY+fc.Y
			if !(image.Point{x, y}.In(context.Dst.Rect)) {
				return errors.New("Text drawing outside of image bounds")
			}
			if fc.Pixel {
				context.Dst.Set(x, y, context.Color)
			}
		}
	}
	return nil
}

// Draw a string on an RGBA image
func Draw(img *image.RGBA, str []byte, offsetY, offsetX int, col color.Color) error {
	for i, c := range str {
		fc := Font8x8.Character(c)
		x := offsetX + i*fc.Width
		if !(image.Point{x, offsetY}.In(img.Rect)) {
			return errors.New("Text drawing outside of image bounds")
		}
		for fc.NextBit() {
			if fc.Pixel {
				img.Set(x+fc.X, int(fc.Y), col)
			}
		}
	}
	return nil
}
