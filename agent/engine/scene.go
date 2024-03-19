package engine

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/ath0m/DistributedRaytracer/agent/engine/camera"
	clr "github.com/ath0m/DistributedRaytracer/agent/engine/color"
	"github.com/ath0m/DistributedRaytracer/agent/engine/geometry"
	"github.com/ath0m/DistributedRaytracer/agent/engine/utils"
)

// Pixels represents the array of pixels (in packed RGB value) to Render and/or save
type Pixels []uint32

// Scene represents the scene to Render.
type Scene struct {
	width, height int
	raysPerPixel  int
	camera        camera.Camera
	world         Hittable
}

func NewScene(width, height, raysPerPixel int, camera camera.Camera, world Hittable) *Scene {
	return &Scene{
		width:        width,
		height:       height,
		raysPerPixel: raysPerPixel,
		camera:       camera,
		world:        world,
	}
}

// pixel is an internal type which represents the pixel to be processed
//
//	x,y are the coordinates
//	k is the index in the Pixels array
//	color is the color that has been computed by casting raysPerPixel through x/y coordinates (not normalized to avoid accumulating rounding errors)
type pixel struct {
	x, y, k      int
	color        clr.Color
	raysPerPixel int
}

// split is a util function which split an array into an array of array with count elements each (the last one may hold less...)
func split(buf []*pixel, count int) [][]*pixel {
	var chunk []*pixel
	chunks := make([][]*pixel, 0, len(buf)/count+1)
	for len(buf) >= count {
		chunk, buf = buf[:count], buf[count:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf)
	}
	return chunks
}

// render works on a single pixels, casting raysPerPixel through it and accumulating the color
//
//	returns the normalized and gamma corrected value
func (scene *Scene) render(rnd utils.Rnd, pixel *pixel, raysPerPixel int) uint32 {
	c := pixel.color

	for s := 0; s < raysPerPixel; s++ {
		u := (float64(pixel.x) + rnd.Float64()) / float64(scene.width)
		v := (float64(pixel.y) + rnd.Float64()) / float64(scene.height)
		r := scene.camera.Ray(rnd, u, v)
		c = c.Add(color(r, scene.world, 0))
	}

	pixel.color = c
	pixel.raysPerPixel += raysPerPixel

	// normalize the color (average of all the rays cast so far)
	c = c.Scale(1.0 / float64(pixel.raysPerPixel))

	// gamma correction
	c = clr.Color{R: math.Sqrt(c.R), G: math.Sqrt(c.G), B: math.Sqrt(c.B)}

	return c.PixelValue()
}

// Render is the main method of a scene. It is non-blocking and returns right away with the array of pixels
// that will be computed asynchronously and a channel to indicate when the processing is complete. Note that
// no synchronization is required on the array of pixels since it is an array of 32 bits values.
// The image (width x height) will be split in lines each one processed in a separate goroutine (parallelCount
// of them).
func (scene *Scene) Render(parallelCount int) (Pixels, chan struct{}) {
	pixels := make([]uint32, scene.width*scene.height)
	completed := make(chan struct{})

	go func() {
		allPixelsToProcess := make([]*pixel, scene.width*scene.height)

		// initializes the pixels to generate (start with black color)
		k := 0
		for j := scene.height - 1; j >= 0; j-- {
			for i := 0; i < scene.width; i++ {
				allPixelsToProcess[k] = &pixel{x: i, y: j, k: k}
				k++
			}
		}

		// split in lines
		lines := split(allPixelsToProcess, scene.width)

		totalStart := time.Now()

		// creates a channel which will be used to dispatch the line to process to each go routine
		pixelsToProcess := make(chan []*pixel)

		// asynchronously dispatch the lines to process
		go func() {
			for _, p := range lines {
				pixelsToProcess <- p
			}
			// done... signal the end
			close(pixelsToProcess)
		}()

		// create a wait group to wait until all goroutine completes
		wg := sync.WaitGroup{}

		// create parallelCount goroutines
		for c := 0; c < parallelCount; c++ {
			wg.Add(1)
			go func() {
				// due to high contention on global rand, each goroutine uses its own random number generator
				// thus avoiding massive slowdown
				rnd := rand.New(rand.NewSource(rand.Int63()))

				// process a bunch of pixels (in this case a line)
				for ps := range pixelsToProcess {

					// redisplay the line without gamma correction => make it darker to be more visible
					for _, p := range ps {
						if p.raysPerPixel > 0 {
							col := p.color.Scale(1.0 / float64(p.raysPerPixel))
							pixels[p.k] = col.PixelValue()
						}
					}

					// render every pixel in the line
					for _, p := range ps {
						pixels[p.k] = scene.render(rnd, p, scene.raysPerPixel)
					}
				}
				wg.Done()
			}()
		}

		// wait for the pass to be completed
		wg.Wait()

		totalTime := time.Since(totalStart)
		fmt.Printf("Processed %v rays per pixel in %v.", scene.raysPerPixel, totalTime)

		// signal completion
		completed <- struct{}{}
	}()

	return pixels, completed
}

// color computes the color of the ray by checking which hitable gets hit and scattering
// more rays (recursive) depending on material
func color(r *geometry.Ray, world Hittable, depth int) clr.Color {

	if hit, hr := world.Hit(r, &utils.Interval{Min: 0.001, Max: math.MaxFloat64}); hit {
		if depth >= 50 {
			return clr.Black
		}

		if wasScattered, attenuation, scattered := hr.Material.scatter(r, hr); wasScattered {
			return attenuation.Mult(color(scattered, world, depth+1))
		} else {
			return clr.Black
		}
	}

	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)

	return clr.White.Scale(1.0 - t).Add(clr.Color{R: 0.5, G: 0.7, B: 1.0}.Scale(t))
}
