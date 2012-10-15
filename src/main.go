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

var width = flag.Int("width", 500, "width of output image")
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

  geom := []monte.Primitive{
    monte.Sphere{Center: monte.Vect(0, 0, 8), Radius: 1.0},
    monte.Sphere{Center: monte.Vect(2, 2, 8), Radius: 1.0},
  }

  scene := monte.Scene{
    Geom: geom,
    Look: monte.NewRay(monte.Vect(1, 1, 0), monte.Vect(-0.1, -0.1, 1)),
    U1: monte.Vect(1, 0, 0),
    U2: monte.Vect(0, 1, 0),
    FDist: 2,
    Sky: monte.NewColorf(0.6, 0.8, 1, 1),
    Oversample: 1,
    Ambient: monte.NewColorf(0.2, 0.2, 0.2, 1),
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
