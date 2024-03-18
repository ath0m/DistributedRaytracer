package geometry

// Point3 defines a point in 3D space
type Point3 struct {
	X, Y, Z float64
}

// Translate translates the point to a new location (return a new point)
func (p Point3) Translate(v Vec3) Point3 {
	return Point3{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

// Sub subtracts a point to another Point which gives a vector
func (p Point3) Sub(p2 Point3) Vec3 {
	return Vec3{p.X - p2.X, p.Y - p2.Y, p.Z - p2.Z}
}

// Vec3 converts a point to a vector (centered at origin)
func (p Point3) Vec3() Vec3 {
	return Vec3(p)
}
