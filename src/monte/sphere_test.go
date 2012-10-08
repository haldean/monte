package monte

import (
  "testing"
)

func TestSphereIntersect(t *testing.T) {
  s := Sphere{Center: Vect(2, 3, 1), Radius: 3}

  r1 := NewRay(Vect(2, 3, 7), Vect(0, 0, -1))
  if isect, _ := s.Intersect(r1); *isect != *Vect(2, 3, 4) {
    t.Errorf("Bad intersection: expected 2, 3, 4 but got %v", isect)
  }

  r2 := NewRay(Vect(2, 6, 7), Vect(0, 0, -1))
  if isect, _ := s.Intersect(r2); *isect != *Vect(2, 6, 1) {
    t.Errorf("Bad intersection: expected 2, 6, 1 but got %v", isect)
  }
}
