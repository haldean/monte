package monte

import "math"

// Vector

type Vector struct {
  X, Y, Z float64
}

func Vect(x, y, z float64) *Vector {
  return &Vector{X: x, Y: y, Z: z}
}

func (v *Vector) Add(w *Vector) *Vector {
  return &Vector{X: v.X + w.X, Y: v.Y + w.Y, Z: v.Z + w.Z}
}

func (v *Vector) Copy() *Vector {
  return &Vector{X: v.X, Y: v.Y, Z: v.Z}
}

func (v *Vector) Dot(w *Vector) float64 {
  return v.X * w.X + v.Y * w.Y + v.Z * w.Z
}

func (v *Vector) Norm() float64 {
  return math.Sqrt(v.NormSqr())
}

func (v *Vector) NormSqr() float64 {
  return v.Dot(v)
}

func (v *Vector) Normalize() *Vector {
  return v.Copy().NormalizeInPlace()
}

// Returns the changed vector for easy chaining in expressions
func (v *Vector) NormalizeInPlace() *Vector {
  norm := v.Norm()
  v.X /= norm
  v.Y /= norm
  v.Z /= norm
  return v
}

func (v *Vector) ScalarMul(s float64) *Vector {
  return &Vector{X: v.X * s, Y: v.Y * s, Z: v.Z * s}
}

// Ray

type Ray struct {
  Origin, Direction *Vector
}

func NewRay(o, d *Vector) *Ray {
  return &Ray{Origin: o, Direction: d.Normalize()}
}
