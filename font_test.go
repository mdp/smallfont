package smallfont

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestFont8x8(t *testing.T) {
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
		e := expected[character.Y*character.Width+character.X]
		if (e == 1) != character.Pixel {
			t.Fail()
		}
	}
}

func TestFont5x8(t *testing.T) {
	character := Font5x8.Character(byte('b'))
	expected := [64]int{
		0, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 0, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 1, 1, 0, 0,
		0, 0, 0, 0, 0,
	}
	for i, p := range character.PixelMap() {
		if (expected[i] == 1) != p {
			t.Fail()
		}
	}

}

func TestFont6x8(t *testing.T) {
	character := Font6x8.Character(byte('b'))
	expected := [64]int{
		0, 1, 0, 0, 0, 0,
		0, 1, 0, 0, 0, 0,
		0, 1, 1, 1, 1, 0,
		0, 1, 0, 0, 0, 1,
		0, 1, 0, 0, 0, 1,
		0, 1, 0, 0, 0, 1,
		0, 1, 1, 1, 1, 0,
		0, 0, 0, 0, 0, 0,
	}
	for i, p := range character.PixelMap() {
		if (expected[i] == 1) != p {
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
	if err != nil {
		t.Error("Should not error when writing to a correctly sized image")
	}
	err = ctx.Draw([]byte("smallfont"), 0, 8)
	err = ctx.Draw([]byte("smallfont"), 0, 16)
	err = ctx.Draw([]byte("smallfont"), 0, 24)

	f, _ := os.OpenFile("out8x8.png", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	png.Encode(f, img)
}

func TestImage5x8(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 5*8, 16))
	ctx := Context{
		Font:  Font5x8,
		Dst:   img,
		Color: color.Black,
	}
	err := ctx.Draw([]byte("smallfont"), 0, 0)
	if err == nil {
		t.Error("Should error when writing outside of bounds")
	}

	img = image.NewRGBA(image.Rect(0, 0, 128, 32))
	ctx = Context{
		Font:  Font5x8,
		Dst:   img,
		Color: color.Black,
	}
	err = ctx.Draw([]byte("smallfontissmallsmallsmal"), 0, 0)
	if err != nil {
		t.Error("Should not error when writing to a correctly sized image")
	}
	err = ctx.Draw([]byte("smallfont"), 0, 8)
	err = ctx.Draw([]byte("smallfont"), 0, 16)
	err = ctx.Draw([]byte("smallfont"), 0, 24)

	f, _ := os.OpenFile("out5x8.png", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	png.Encode(f, img)
}

func TestImage6x8(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 6*8, 16))
	ctx := Context{
		Font:  Font6x8,
		Dst:   img,
		Color: color.Black,
	}
	err := ctx.Draw([]byte("smallfont"), 0, 0)
	if err == nil {
		t.Error("Should error when writing outside of bounds")
	}

	img = image.NewRGBA(image.Rect(0, 0, 128, 32))
	ctx = Context{
		Font:  Font6x8,
		Dst:   img,
		Color: color.Black,
	}
	err = ctx.Draw([]byte("smallfontissmallsmall"), 0, 0) // 21 max
	if err != nil {
		t.Error("Should not error when writing to a correctly sized image")
	}
	err = ctx.Draw([]byte("smallfont"), 0, 8)
	err = ctx.Draw([]byte("smallfont"), 0, 16)
	err = ctx.Draw([]byte("smallfont"), 0, 24)

	f, _ := os.OpenFile("out6x8.png", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	png.Encode(f, img)
}
