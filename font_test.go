package smallfont

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestFont(t *testing.T) {
	character := Font8x8.Character(byte('b'))
	expected := [64]int{
		1, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 0, 0,
		1, 1, 0, 0, 0, 1, 1, 0,
		1, 1, 0, 0, 0, 1, 1, 0,
		1, 1, 1, 1, 1, 1, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	for character.NextBit() {
		e := expected[character.Y*character.Height+character.X]
		if (e == 1) != character.Pixel {
			t.Fail()
		}
	}
}

func TestImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 8*8, 16))
	ctx := Context{
		Font:  Font8x8,
		Dst:   img,
		Color: color.Black,
	}
	err := ctx.Draw([]byte("smallfont"), 0, 0)
	if err == nil {
		t.Error("Should error when writing outside of bounds")
	}

	img = image.NewRGBA(image.Rect(0, 0, 128, 32))
	ctx = Context{
		Font:  Font8x8,
		Dst:   img,
		Color: color.Black,
	}
	err = ctx.Draw([]byte("smallfontissmall"), 0, 0)
	err = ctx.Draw([]byte("smallfont"), 0, 8)
	err = ctx.Draw([]byte("smallfont"), 0, 16)
	err = ctx.Draw([]byte("smallfont"), 0, 24)
	if err != nil {
		t.Error("Should not error when writing to a correctly sized image")
	}

	f, _ := os.OpenFile("out.png", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	png.Encode(f, img)
}
