# SmallFont

Rasterize small fonts (8x8 currently supported) for LED projects

## Example usage

```golang
message := []byte("smallfont")
img := image.NewRGBA(image.Rect(0, 0, 128, 32)) // 52Pi OLED size
ctx := Context{
  Font:  Font8x8,
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
```

## License

MIT - Do with this as you please
