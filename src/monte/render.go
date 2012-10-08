package monte

import (
  //"fmt"
  "image"
  "image/color"
  "math"
  "math/rand"
)

func NewColor(r, g, b, a uint8) color.NRGBA {
  return color.NRGBA{R: r, G: g, B: b, A: a}
}

func (s *Scene) ColorAtIntersection(geom *Primitive, in *Ray) color.NRGBA {
  shade := in.Direction.ScalarMul(-1).Dot((*geom).NormalAt(in.Origin))
  value := uint8(255.0 * shade)
  return NewColor(value, value, value, 255)
}

func (s *Scene) RayCast(u, v float64) color.NRGBA {
  ray := NewRay(s.Look.Origin, s.DirectionAt(u, v))

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

  if minGeom != nil {
    return s.ColorAtIntersection(minGeom, NewRay(intersect, ray.Direction))
  }
  return s.Sky
}

func (s *Scene) Evaluate(u, v float64) color.Color {
  var r, g, b, n uint32 = 0, 0, 0, 0
  for i := 0; i < s.Oversample; i++ {
    du := (rand.Float64() - 0.5) / 200
    dv := (rand.Float64() - 0.5) / 200

    c := s.RayCast(u + du, v + dv)
    r += uint32(c.R); g += uint32(c.G); b += uint32(c.B)
    n++
  }
  return NewColor(uint8(r / n), uint8(g / n), uint8(b / n), 255)
}

func (s *Scene) SetColor(i, j int, u, v float64, img *image.NRGBA) {
  img.Set(i, j, s.Evaluate(u, v))
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
      s.SetColor(i, j, u, v, img)
    }
  }
}
