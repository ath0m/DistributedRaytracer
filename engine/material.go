package engine

import "math"

// Material defines how a material scatter light
type Material interface {
	scatter(r *Ray, rec *HitRecord) (wasScattered bool, attenuation *Color, scattered *Ray)
}

type Lambertian struct {
	albedo Color
}

func (mat Lambertian) scatter(r *Ray, rec *HitRecord) (bool, *Color, *Ray) {
	dir := rec.normal.Add(RandomUnitSphere(r.rnd))
	if dir.NearZero() {
		dir = rec.normal
	}
	scattered := &Ray{rec.p, dir, r.rnd}
	attenuation := &mat.albedo
	return true, attenuation, scattered
}

type Metal struct {
	albedo Color
	fuzz   float64
}

func (mat Metal) scatter(r *Ray, rec *HitRecord) (bool, *Color, *Ray) {
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

func schlick(cosine float64, iRefIdx float64) float64 {
	r0 := (1.0 - iRefIdx) / (1.0 + iRefIdx)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1.0-cosine, 5)
}

func (die Dielectric) scatter(r *Ray, rec *HitRecord) (bool, *Color, *Ray) {
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

	return true, &White, &Ray{rec.p, direction, r.rnd}
}
