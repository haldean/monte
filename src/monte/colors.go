package monte

import (
  "image/color"
)

// Color type backed by a float which supports HDR
type Colorf struct {
  R, G, B, A float64
}

func NewColorf(r, g, b, a float64) Colorf {
  return Colorf{R: r, G: g, B: b, A: a}
}

func NewColor(r, g, b, a int32) Colorf {
  return Colorf{
    R: float64(r), G: float64(g), B: float64(b), A: float64(a)}
}

func clamp8(x float64) uint8 {
  x *= 255
  switch {
  case x >= 256:
    return 255
  case x < 0:
    return 0
  }
  return uint8(x)
}

func (c Colorf) NRGBA() color.NRGBA {
  return color.NRGBA{
    R: clamp8(c.R), G: clamp8(c.G), B: clamp8(c.B), A: clamp8(c.A)}
}

func (c1 Colorf) Add(c2 Colorf) Colorf {
  return Colorf{
    R: c1.R + c2.R, G: c1.G + c2.G, B: c1.B + c2.B, A: c1.A + c2.A,
  }
}

func (c Colorf) Scale(s float64) Colorf {
  return Colorf{
    R: c.R * s, G: c.G * s, B: c.B * s, A: c.A * s,
  }
}

func (c1 Colorf) Mix(c2 Colorf) Colorf {
  return Colorf{
    R: c1.R * c2.R, G: c1.G * c2.G, B: c1.B * c2.B, A: c1.A * c2.A,
  }
}
