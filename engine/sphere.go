package engine

import "math"

type Sphere struct {
	center   Point3
	radius   float64
	material Material
}

// hit implements the hit interface for a Sphere
func (s Sphere) hit(r *Ray, interval *Interval) (bool, *HitRecord) {
	oc := r.Origin.Sub(s.center)
	a := r.Direction.LengthSq()
	b := Dot(oc, r.Direction)
	c := oc.LengthSq() - s.radius*s.radius
	discriminant := b*b - a*c

	if discriminant < 0 {
		return false, nil
	}
	discriminantSquareRoot := math.Sqrt(discriminant)

	root := (-b - discriminantSquareRoot) / a
	if !interval.surrounds(root) {
		root = (-b + discriminantSquareRoot) / a
		if !interval.surrounds(root) {
			return false, nil
		}
	}

	hitPoint := r.PointAt(root)
	hr := HitRecord{
		t:        root,
		p:        hitPoint,
		normal:   hitPoint.Sub(s.center).Scale(1 / s.radius),
		material: s.material,
	}
	return true, &hr
}
