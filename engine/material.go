package engine

import (
	"encoding/json"
	"fmt"
	"math"

	clr "github.com/ath0m/DistributedRaytracer/engine/color"
	"github.com/ath0m/DistributedRaytracer/engine/geometry"
)

// Material defines how a material scatter light
type Material interface {
	scatter(r *geometry.Ray, rec *HitRecord) (wasScattered bool, attenuation *clr.Color, scattered *geometry.Ray)
}

func UnmarshalMaterial(data json.RawMessage) (Material, error) {
	var m struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	switch m.Type {
	case "Lambertian":
		var l struct {
			Albedo clr.Color `json:"albedo"`
		}
		err := json.Unmarshal(data, &l)
		if err != nil {
			return nil, err
		}
		return Lambertian{albedo: l.Albedo}, nil

	case "Metal":
		var mt struct {
			Albedo clr.Color `json:"albedo"`
			Fuzz   float64   `json:"fuzz"`
		}
		err := json.Unmarshal(data, &mt)
		if err != nil {
			return nil, err
		}
		return Metal{albedo: mt.Albedo, fuzz: mt.Fuzz}, nil

	case "Dielectric":
		var d struct {
			RefIdx float64 `json:"refIdx"`
		}
		err := json.Unmarshal(data, &d)
		if err != nil {
			return nil, err
		}
		return Dielectric{refIdx: d.RefIdx}, nil

	default:
		return nil, fmt.Errorf("unknown material type: %s", m.Type)
	}
}

type Lambertian struct {
	albedo clr.Color
}

func (mat Lambertian) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type   string    `json:"type"`
		Albedo clr.Color `json:"albedo"`
	}{
		Type:   "Lambertian",
		Albedo: mat.albedo,
	})
}

func (mat Lambertian) scatter(r *geometry.Ray, rec *HitRecord) (bool, *clr.Color, *geometry.Ray) {
	dir := rec.Normal.Add(geometry.RandomUnitSphere(r.Rnd))
	if dir.NearZero() {
		dir = rec.Normal
	}
	scattered := &geometry.Ray{Origin: rec.P, Direction: dir, Rnd: r.Rnd}
	attenuation := &mat.albedo
	return true, attenuation, scattered
}

type Metal struct {
	albedo clr.Color
	fuzz   float64
}

func (mat Metal) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type   string    `json:"type"`
		Albedo clr.Color `json:"albedo"`
		Fuzz   float64   `json:"fuzz"`
	}{
		Type:   "Metal",
		Albedo: mat.albedo,
		Fuzz:   mat.fuzz,
	})
}

func (mat Metal) scatter(r *geometry.Ray, rec *HitRecord) (bool, *clr.Color, *geometry.Ray) {
	reflected := r.Direction.Unit().Reflect(rec.Normal)
	reflected = reflected.Add(geometry.RandomUnitSphere(r.Rnd).Scale(math.Min(mat.fuzz, 1.0)))
	scattered := &geometry.Ray{Origin: rec.P, Direction: reflected, Rnd: r.Rnd}
	attenuation := &mat.albedo

	if geometry.Dot(scattered.Direction, rec.Normal) > 0 {
		return true, attenuation, scattered
	}

	return false, nil, nil
}

type Dielectric struct {
	refIdx float64
}

func (die Dielectric) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type   string  `json:"type"`
		RefIdx float64 `json:"refIdx"`
	}{
		Type:   "Dielectric",
		RefIdx: die.refIdx,
	})
}

func schlick(cosine float64, iRefIdx float64) float64 {
	r0 := (1.0 - iRefIdx) / (1.0 + iRefIdx)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1.0-cosine, 5)
}

func (die Dielectric) scatter(r *geometry.Ray, rec *HitRecord) (bool, *clr.Color, *geometry.Ray) {
	var (
		outwardNormal geometry.Vec3
		niOverNt      float64
		cosine        float64
	)

	dotRayNormal := geometry.Dot(r.Direction, rec.Normal)
	if dotRayNormal > 0 {
		outwardNormal = rec.Normal.Negate()
		niOverNt = die.refIdx
		cosine = dotRayNormal / r.Direction.Length()
		cosine = math.Sqrt(1.0 - die.refIdx*die.refIdx*(1.0-cosine*cosine))
	} else {
		outwardNormal = rec.Normal
		niOverNt = 1.0 / die.refIdx
		cosine = -dotRayNormal / r.Direction.Length()
	}

	wasRefracted, refracted := r.Direction.Refract(outwardNormal, niOverNt)

	var direction geometry.Vec3

	// refract only with some probability
	if wasRefracted && r.Rnd.Float64() >= schlick(cosine, die.refIdx) {
		direction = *refracted
	} else {
		direction = r.Direction.Unit().Reflect(rec.Normal)
	}

	return true, &clr.White, &geometry.Ray{Origin: rec.P, Direction: direction, Rnd: r.Rnd}
}
