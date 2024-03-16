package engine

import "math/rand"

// buildWorldDielectrics is the end result book
func BuildWorldOneWeekend(width, height int) (Camera, HittableList) {
	world := []Hittable{}

	maxSpheres := 500
	world = append(world, Sphere{center: Point3{Y: -1000.0}, radius: 1000, material: Lambertian{Color{R: 0.5, G: 0.5, B: 0.5}}})

	for a := -11; a < 11 && len(world) < maxSpheres; a++ {
		for b := -11; b < 11 && len(world) < maxSpheres; b++ {
			chooseMaterial := rand.Float64()
			center := Point3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if center.Sub(Point3{4.0, 0.2, 0}).Length() > 0.9 {
				switch {
				case chooseMaterial < 0.8: // diffuse
					world = append(world,
						Sphere{
							center:   center,
							radius:   0.2,
							material: Lambertian{Color{R: rand.Float64() * rand.Float64(), G: rand.Float64() * rand.Float64(), B: rand.Float64() * rand.Float64()}}})
				case chooseMaterial < 0.95: // metal
					world = append(world,
						Sphere{
							center:   center,
							radius:   0.2,
							material: Metal{Color{R: 0.5 * (1 + rand.Float64()), G: 0.5 * (1 + rand.Float64()), B: 0.5 * (1 + rand.Float64())}, 0.5 * rand.Float64()}})
				default:
					world = append(world,
						Sphere{
							center:   center,
							radius:   0.2,
							material: Dielectric{1.5}})

				}
			}
		}
	}

	world = append(world,
		Sphere{
			center:   Point3{0, 1, 0},
			radius:   1.0,
			material: Dielectric{1.5}},
		Sphere{
			center:   Point3{-4, 1, 0},
			radius:   1.0,
			material: Lambertian{Color{0.4, 0.2, 0.1}}},
		Sphere{
			center:   Point3{4, 1, 0},
			radius:   1.0,
			material: Metal{Color{0.7, 0.6, 0.5}, 0}})

	lookFrom := Point3{13, 2, 3}
	lookAt := Point3{}
	aperture := 0.1
	distToFocus := 10.0
	camera := NewCamera(lookFrom, lookAt, Vec3{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus)

	return camera, world
}
