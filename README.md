![smallfont](https://cloud.githubusercontent.com/assets/2868/22627960/8b74b71c-eb80-11e6-947e-776c3477a12d.png)

# SmallFont

[![Build Status](https://travis-ci.org/mdp/smallfont.svg?branch=master)](https://travis-ci.org/mdp/smallfont)

Rasterize small fonts (8x8 currently supported) for LED projects

## Example usage

```golang
package main

import "github.com/mdp/smallfont"

func main() {
  message := []byte("smallfont")
  img := image.NewRGBA(image.Rect(0, 0, 128, 32)) // 52Pi OLED size
  ctx := smallfont.Context{
    Font:  smallfont.Font8x8,
    Dst:   img,
    Color: color.Black,
  }
  err := ctx.Draw(message, 0, 0)
  if err != nil {
    fmt.Println(err)
  }
  f, _ := os.OpenFile("out.png", os.O_CREATE|os.O_RDWR, 0644)
  defer f.Close()
  png.Encode(f, img)
}
```

## License

MIT - Do with this as you please
