package geometry

import (
	"github.com/ath0m/DistributedRaytracer/agent/engine/utils"
)

type Ray struct {
	Origin    Point3
	Direction Vec3
	Rnd       utils.Rnd
}

// PointAt returns a new point along the ray (0 will return the origin)
func (r *Ray) PointAt(t float64) Point3 {
	return r.Origin.Translate(r.Direction.Scale(t))
}
