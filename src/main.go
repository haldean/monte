package main

import (
  "flag"
  "image"
  "image/png"
  "log"
  "monte"
  "os"
  "runtime/pprof"
)

var width = flag.Int("width", 300, "width of output image")
var height = flag.Int("height", 300, "height of output image")
var output = flag.String("output", "output.png", "output PNG file")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
  flag.Parse()
  if *cpuprofile != "" {
    f, err := os.Create(*cpuprofile)
    if err != nil {
      log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
  }

  println("Monte is go.")

  geom := []monte.Primitive{monte.Sphere{Center: monte.Vect(0, 0, 4), Radius: 1.0}}

  scene := monte.Scene{
    Geom: geom,
    Look: monte.NewRay(monte.Vect(0, 0, 0), monte.Vect(0, 0, 1)),
    U1: monte.Vect(1, 0, 0),
    U2: monte.Vect(0, 1, 0),
    FDist: 1,
    Sky: monte.NewColor(0, 200, 255, 255),
    Oversample: 8,
  }

  img := image.NewNRGBA(image.Rect(0, 0, *width, *height))
  scene.Render(img)

  out, err := os.Create(*output)
  if err != nil {
    log.Fatalf("Could not open output file: %v\n", err)
    return
  }
  png.Encode(out, img)
  out.Close()

  pprof.StopCPUProfile()
}
