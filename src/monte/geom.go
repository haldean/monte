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
}
