package engine

import (
	"math"

	"github.com/ath0m/DistributedRaytracer/engine/utils"
)

// Vec3 defines a vector in 3D space
type Vec3 struct {
	X, Y, Z float64
}

// Scale scales the vector by the value (return a new vector)
func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{X: v.X * t, Y: v.Y * t, Z: v.Z * t}
}

// Mult multiplies the vector by the other one (return a new vector)
func (v Vec3) Mult(v2 Vec3) Vec3 {
	return Vec3{X: v.X * v2.X, Y: v.Y * v2.Y, Z: v.Z * v2.Z}
}

// Sub substracts the 2 vectors (return a new vector)
func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

// Add adds the 2 vectors (return a new vector)
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{X: v.X + v2.X, Y: v.Y + v2.Y, Z: v.Z + v2.Z}
}

// Length returns the size of the vector
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

// LengthSq returns the squared size of the vector
func (v Vec3) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Unit returns a new vector with same direction and length 1
func (v Vec3) Unit() Vec3 {
	return v.Scale(1.0 / v.Length())
}

// Negate returns a new vector with X/Y/Z negated
func (v Vec3) Negate() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (v Vec3) NearZero() bool {
	eps := 1e-8
	return math.Abs(v.X) < eps && math.Abs(v.Y) < eps && math.Abs(v.Z) < eps
}

// Dot returns the dot product (a scalar) of 2 vectors
func Dot(v1 Vec3, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Cross returns the cross product of 2 vectors (another vector)
func Cross(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{v1.Y*v2.Z - v1.Z*v2.Y, -(v1.X*v2.Z - v1.Z*v2.X), v1.X*v2.Y - v1.Y*v2.X}
}

// Reflect simply reflects the vector based on the normal n
func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Sub(n.Scale(2.0 * Dot(v, n)))
}

// Refract returns a refracted vector (or not if there is no refraction possible)
func (v Vec3) Refract(n Vec3, niOverNt float64) (bool, *Vec3) {
	uv := v.Unit()
	un := n.Unit()

	dt := Dot(uv, un)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted := uv.Sub(un.Scale(dt)).Scale(niOverNt).Sub(un.Scale(math.Sqrt(discriminant)))
		return true, &refracted
	}

	return false, nil
}

func Random(rnd utils.Rnd) Vec3 {
	return Vec3{rnd.Float64(), rnd.Float64(), rnd.Float64()}
}

func RandomInt(rnd utils.Rnd, interval *utils.Interval) Vec3 {
	diff := interval.Max - interval.Min
	x := interval.Min + rnd.Float64()*diff
	y := interval.Min + rnd.Float64()*diff
	z := interval.Min + rnd.Float64()*diff
	return Vec3{x, y, z}
}

func RandomInUnitSphere(rnd utils.Rnd) Vec3 {
	interval := utils.Interval{-1, 1}
	for {
		if v := RandomInt(rnd, &interval); v.LengthSq() < 1 {
			return v
		}
	}
}

func RandomUnitSphere(rnd utils.Rnd) Vec3 {
	return RandomInUnitSphere(rnd).Unit()
}

func RandomInUnitDisk(rnd utils.Rnd) Vec3 {
	for {
		p := Vec3{2.0*rnd.Float64() - 1.0, 2.0*rnd.Float64() - 1.0, 0}
		if Dot(p, p) < 1.0 {
			return p
		}
	}
}
