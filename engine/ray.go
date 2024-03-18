package engine

import (
	"github.com/ath0m/DistributedRaytracer/engine/geometry"
	"github.com/ath0m/DistributedRaytracer/engine/utils"
)

type Ray struct {
	Origin    geometry.Point3
	Direction geometry.Vec3
	rnd       utils.Rnd
}

// PointAt returns a new point along the ray (0 will return the origin)
func (r *Ray) PointAt(t float64) geometry.Point3 {
	return r.Origin.Translate(r.Direction.Scale(t))
}
