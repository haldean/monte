package monte

import (
  "image"
  "image/color"
  "math"
  "math/rand"
)

func NewColor(r, g, b, a uint8) color.NRGBA {
  return color.NRGBA{R: r, G: g, B: b, A: a}
}

type Scene struct {
  Geom []Primitive
  Look *Ray
  // The image vector pair; these orthogonal vectors form the UV-space which we
  // scan to fill in the picture
  U1, U2 *Vector
  FDist float64
  Sky color.NRGBA
  Oversample int
}

func (s *Scene) DirectionAt(u, v float64) *Vector {
  u1, u2 := s.U1.ScalarMul(u), s.U2.ScalarMul(v)
  return s.Look.Direction.ScalarMul(s.FDist).Add(u1).Add(u2).NormalizeInPlace()
}

func (s *Scene) ColorAt(u, v float64) color.Color {
  var r, g, b, n int32 = 0, 0, 0, 0
  sr, sg, sb := s.Sky.R, s.Sky.G, s.Sky.B

  for i := 0; i < s.Oversample; i++ {
    du := (rand.Float64() - 0.5) / 200
    dv := (rand.Float64() - 0.5) / 200
    ray := NewRay(s.Look.Origin, s.DirectionAt(u+du, v+dv))

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
      r += 255
      g += 255
      b += 255
    } else {
      r += int32(sr)
      g += int32(sg)
      b += int32(sb)
    }
    n++
  }
  return NewColor(uint8(r / n), uint8(g / n), uint8(b / n), 255)
}

func (s *Scene) SetColor(i, j int, u, v float64, img *image.NRGBA) {
  img.Set(i, j, s.ColorAt(u, v))
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
      go s.SetColor(i, j, u, v, img)
    }
  }
}
