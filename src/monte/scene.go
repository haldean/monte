package monte

import (
  "image"
  "image/color"
  "math"
)

func NewColor(r, g, b, a uint8) color.Color {
  return color.NRGBA{R: r, G: g, B: b, A: a}
}

type Scene struct {
  Geom []Primitive
  Look *Ray
  // The image vector pair; these orthogonal vectors form the UV-space which we
  // scan to fill in the picture
  U1, U2 *Vector
  FDist float64
}

func (s *Scene) DirectionAt(u, v float64) *Vector {
  u1, u2 := s.U1.ScalarMul(u), s.U2.ScalarMul(v)
  return s.Look.Direction.ScalarMul(s.FDist).Add(u1).Add(u2).NormalizeInPlace()
}

func (s *Scene) ColorAt(u, v float64) color.Color {
  ray := NewRay(s.Look.Origin, s.DirectionAt(u, v))

  minDist := math.Inf(1)
  var minGeom *Primitive = nil

  for _, geom := range s.Geom {
    i, d := geom.Intersect(ray)
    if i != nil && d < minDist {
      minDist = d
      minGeom = &geom
    }
  }

  if minGeom != nil {
    return NewColor(255, 255, 255, 255)
  }
  return NewColor(0, 0, 0, 255)
}

func (s *Scene) SetColor(i, j int, u, v float64, img *image.NRGBA) {
  img.Set(i, j, s.ColorAt(u, v))
}
