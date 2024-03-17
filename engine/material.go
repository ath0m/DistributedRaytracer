package engine

import (
	"encoding/json"
	"fmt"
	"math"

	clr "github.com/ath0m/DistributedRaytracer/engine/color"
)

// Material defines how a material scatter light
type Material interface {
	scatter(r *Ray, rec *HitRecord) (wasScattered bool, attenuation *clr.Color, scattered *Ray)
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

func (mat Lambertian) scatter(r *Ray, rec *HitRecord) (bool, *clr.Color, *Ray) {
	dir := rec.normal.Add(RandomUnitSphere(r.rnd))
	if dir.NearZero() {
		dir = rec.normal
	}
	scattered := &Ray{rec.p, dir, r.rnd}
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

func (mat Metal) scatter(r *Ray, rec *HitRecord) (bool, *clr.Color, *Ray) {
	reflected := r.Direction.Unit().Reflect(rec.normal)
	reflected = reflected.Add(RandomUnitSphere(r.rnd).Scale(math.Min(mat.fuzz, 1.0)))
	scattered := &Ray{rec.p, reflected, r.rnd}
	attenuation := &mat.albedo

	if Dot(scattered.Direction, rec.normal) > 0 {
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

func (die Dielectric) scatter(r *Ray, rec *HitRecord) (bool, *clr.Color, *Ray) {
	var (
		outwardNormal Vec3
		niOverNt      float64
		cosine        float64
	)

	dotRayNormal := Dot(r.Direction, rec.normal)
	if dotRayNormal > 0 {
		outwardNormal = rec.normal.Negate()
		niOverNt = die.refIdx
		cosine = dotRayNormal / r.Direction.Length()
		cosine = math.Sqrt(1.0 - die.refIdx*die.refIdx*(1.0-cosine*cosine))
	} else {
		outwardNormal = rec.normal
		niOverNt = 1.0 / die.refIdx
		cosine = -dotRayNormal / r.Direction.Length()
	}

	wasRefracted, refracted := r.Direction.Refract(outwardNormal, niOverNt)

	var direction Vec3

	// refract only with some probability
	if wasRefracted && r.rnd.Float64() >= schlick(cosine, die.refIdx) {
		direction = *refracted
	} else {
		direction = r.Direction.Unit().Reflect(rec.normal)
	}

	return true, &clr.White, &Ray{rec.p, direction, r.rnd}
}
