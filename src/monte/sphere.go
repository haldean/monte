package monte

import (
  "math"
)

type Sphere struct {
	Center *Vector
	Radius float64
}

// finds the two x's such that ax^2 + bx + c = 0
func quadSolve(a, b, c float64) (float64, float64) {
  disc := math.Sqrt(b * b - 4 * a * c)
  return (-b + disc) / (2 * a), (-b - disc) / (2 * a)
}

// Finds the intersection of a sphere with a ray. Returns nil if no intersection
func (s Sphere) Intersect(ray *Ray) (*Vector, float64) {
  a := ray.Direction.NormSqr()
  b := 2 * (ray.Origin.Dot(ray.Direction) - s.Center.Dot(ray.Direction))
  c := ray.Origin.NormSqr() + s.Center.NormSqr() -
      2 * ray.Origin.Dot(s.Center) - math.Pow(s.Radius, 2)

      s1, s2 := quadSolve(a, b, c)
  switch {
  case math.IsNaN(s1) && math.IsNaN(s2):
    return nil, 0
  case math.IsNaN(s1):
    return ray.Origin.Add(ray.Direction.ScalarMul(s2)), s2
  case math.IsNaN(s2):
    return ray.Origin.Add(ray.Direction.ScalarMul(s1)), s1
  default:
    if s1 < s2 {
      return ray.Origin.Add(ray.Direction.ScalarMul(s1)), s1
    } else {
      return ray.Origin.Add(ray.Direction.ScalarMul(s2)), s2
    }
  }

  return nil, 0
}

func (s Sphere) Normal(loc *Vector) *Vector {
  return Vect(0, 0, 0)
}
