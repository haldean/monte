package main

import (
  "flag"
  "fmt"
  "image"
  "image/color"
  "image/png"
  "os"
)

var Width = flag.Int("width", 300, "width of output image")
var Height = flag.Int("height", 300, "height of output image")
var Output = flag.String("output", "output.png", "output PNG file")

func NewColor(r, g, b, a uint8) color.Color {
  return color.NRGBA{R: r, G: g, B: b, A: a}
}

func main() {
  flag.Parse()

  println("Monte is go.")
  img := image.NewNRGBA(image.Rect(0, 0, *Width, *Height))
  for i := 0; i < *Width; i++ {
    for j := 0; j < *Height; j++ {
      img.Set(i, j, NewColor(255, 0, 0, 255))
    }
  }

  w, err := os.Create(*Output)
  if err != nil {
    fmt.Printf("Could not open output file: %v\n", err)
    return
  }
  png.Encode(w, img)
  w.Close()
}
