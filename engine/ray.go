package engine

// Ray represents a ray defined by its origin and direction
type Ray struct {
	Origin    Point3
	Direction Vec3
	rnd       Rnd
}

// PointAt returns a new point along the ray (0 will return the origin)
func (r *Ray) PointAt(t float64) Point3 {
	return r.Origin.Translate(r.Direction.Scale(t))
}
