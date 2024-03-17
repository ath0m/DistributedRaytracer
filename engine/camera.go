package engine

import (
	"math"

	"github.com/ath0m/DistributedRaytracer/engine/utils"
)

type Camera interface {
	ray(rnd utils.Rnd, u, v float64) *Ray
}

type camera struct {
	Origin          Point3  `json:"origin"`
	LowerLeftCorner Point3  `json:"lowerLeftCorner"`
	Horizontal      Vec3    `json:"horizontal"`
	Vertical        Vec3    `json:"vertical"`
	U               Vec3    `json:"u"`
	V               Vec3    `json:"v"`
	LensRadius      float64 `json:"lensRadius"`
}

// NewCamera computes the parameters necessary for the camera...
//
//	vfov is expressed in degrees (not radians)
func NewCamera(lookFrom Point3, lookAt Point3, vup Vec3, vfov float64, aspect float64, aperture float64, focusDist float64) Camera {
	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	origin := lookFrom
	w := lookFrom.Sub(lookAt).Unit()
	u := Cross(vup, w).Unit()
	v := Cross(w, u)

	lowerLeftCorner := origin.Translate(u.Scale(-(halfWidth * focusDist))).Translate(v.Scale(-(halfHeight * focusDist))).Translate(w.Scale(-focusDist))
	horizontal := u.Scale(2 * halfWidth * focusDist)
	vertical := v.Scale(2 * halfHeight * focusDist)

	return camera{origin, lowerLeftCorner, horizontal, vertical, u, v, aperture / 2.0}
}

// ray implements the main api of the Camera interface according to the book
func (c camera) ray(rnd utils.Rnd, u, v float64) *Ray {
	d := c.LowerLeftCorner.Translate(c.Horizontal.Scale(u)).Translate(c.Vertical.Scale(v)).Sub(c.Origin)
	origin := c.Origin

	if c.LensRadius > 0 {
		rd := RandomInUnitDisk(rnd).Scale(c.LensRadius)
		offset := c.U.Scale(rd.X).Add(c.V.Scale(rd.Y))
		origin = origin.Translate(offset)
		d = d.Sub(offset)
	}
	return &Ray{origin, d, rnd}
}
