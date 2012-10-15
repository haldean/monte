package monte

import (
  "flag"
  "image"
  "math"
  "math/rand"
)

var debug = flag.Bool("debug", false, "print debugging information")

func (s *Scene) ColorAtIntersection(geom *Primitive, in *Ray, d uint8) Colorf {
  norm := (*geom).NormalAt(in.Origin)
  c := NewColorf(0, 0, 0, 0)
  n := 0.0

  for i := 0; i < 20; i++ {
    out := NewRay(in.Origin, RandomVectorInHemisphere(norm))
    dc := s.RayCast(out, d + 1)
    c = c.Add(dc)
    n += 1.
  }
  return c.Scale(1 / n).Mix((*geom).BRDF(in))
}

// d is the ray casting depth (number of reflections we've done)
func (s *Scene) RayCast(ray *Ray, d uint8) Colorf {
  if d > 2 {
    return s.Ambient
  }

  minDist := math.Inf(1)
  var minGeom *Primitive = nil
  var intersect *Vector = nil

  for _, geom := range s.Geom {
    i, d := geom.Intersect(ray)
    if i != nil && d < minDist {
      minDist = d
      minGeom = &geom
      intersect = i
    }
  }

  if minGeom != nil && minDist > 0 {
    return s.ColorAtIntersection(minGeom, NewRay(intersect, ray.Direction), d)
  }
  return s.Sky
}

func (s *Scene) Evaluate(u, v float64) Colorf {
  // If we're not antialiasing, this is a super simple function.
  if s.Oversample <= 1 {
    ray := NewRay(s.Look.Origin, s.DirectionAt(u, v))
    return s.RayCast(ray, 0)
  }

  c := NewColorf(0, 0, 0, 0)
  n := 0.0
  for i := 0; i < s.Oversample; i++ {
    du := (rand.Float64() - 0.5) / 200
    dv := (rand.Float64() - 0.5) / 200

    ray := NewRay(s.Look.Origin, s.DirectionAt(u + du, v + dv))
    c = c.Add(s.RayCast(ray, 0))
    n += 1.
  }
  return c.Scale(1 / n)
}

func (s *Scene) SetColor(i, j int, u, v float64, img *image.NRGBA) {
  c := s.Evaluate(u, v).NRGBA()
  img.Set(i, j, c)
}

func (s *Scene) Render(img *image.NRGBA) {
  rect := img.Bounds()
  w, h := float64(rect.Dx()), float64(rect.Dy())
  i_max, j_max := rect.Dx(), rect.Dy()
  for i := 0; i < i_max; i++ {
    u := (float64(i) - w / 2.0) / (w / 2.0)
    for j := 0; j < j_max; j++ {
      // (h / w) term is needed to correct for nonunity aspect ratios
      v := (float64(j) - h / 2.0) / (h / 2.0) * (h / w)

      if *debug {
        //fmt.Printf("%d/%d\n", i * j_max + j, i_max * j_max)
      }

      s.SetColor(i, j, u, v, img)
    }
  }
}
