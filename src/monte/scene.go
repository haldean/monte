package monte

type Scene struct {
  Geom []Primitive
  Look *Ray
  // The image vector pair; these orthogonal vectors form the UV-space which we
  // scan to fill in the picture
  U1, U2 *Vector
  FDist float64
  Sky Colorf
  Oversample int
  Ambient Colorf
}

func (s *Scene) DirectionAt(u, v float64) *Vector {
  u1, u2 := s.U1.ScalarMul(u), s.U2.ScalarMul(v)
  return s.Look.Direction.ScalarMul(s.FDist).Add(u1).Add(u2).NormalizeInPlace()
}
