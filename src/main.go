package main

import (
  "flag"
  "fmt"
  "image"
  "image/png"
  "monte"
  "os"
)

var Width = flag.Int("width", 300, "width of output image")
var Height = flag.Int("height", 300, "height of output image")
var Output = flag.String("output", "output.png", "output PNG file")

func main() {
  flag.Parse()

  println("Monte is go.")

  geom := []monte.Primitive{monte.Sphere{Center: monte.Vect(0, 0, 4), Radius: 1.0}}

  scene := monte.Scene{
    Geom: geom,
    Look: monte.NewRay(monte.Vect(0, 0, 0), monte.Vect(0, 0, 1)),
    U1: monte.Vect(1, 0, 0),
    U2: monte.Vect(0, 1, 0),
    FDist: 1}

  img := image.NewNRGBA(image.Rect(0, 0, *Width, *Height))
  w, h := float64(*Width), float64(*Height)
  for i := 0; i < *Width; i++ {
    u := (float64(i) - w / 2.0) / (w / 2.0)
    for j := 0; j < *Height; j++ {
      v := (float64(j) - h / 2.0) / (h / 2.0)
      go scene.SetColor(i, j, u, v, img)
    }
  }

  out, err := os.Create(*Output)
  if err != nil {
    fmt.Printf("Could not open output file: %v\n", err)
    return
  }
  png.Encode(out, img)
  out.Close()
}
