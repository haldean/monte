package monte

type Primitive interface {
  // Find the intersection between the geometric primitive and the given ray.
  // Returns the point of intersection and the distance travelled along the
  // direction vector. Returns (nil, 0) if no intersection.
	Intersect(ray *Ray) (*Vector, float64)

  // Find the normal vector to the geometric primitive at a point on its
  // surface. This function should only be called with points on its surface; if
  // not its behavior is undefined.
  NormalAt(loc *Vector) *Vector

  // The BRDF of a primitive is a function that relates an incoming ray to a
  // color. The BRDF function takes a location on the primitive and a view
  // direction and returns the value of the BRDF at that location.
  BRDF(ray *Ray) Colorf
}

func LambertBrdf(r *Ray, norm *Vector, base Colorf, amb Colorf) Colorf {
  return base.Add(amb).Scale(norm.Dot(r.Direction.ScalarMul(-1)))
}
