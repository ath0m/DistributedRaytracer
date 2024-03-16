// Simple ray tracer based on the Ray Tracing book series by Peter Shirley
package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ath0m/DistributedRaytracer/engine"
)

func main() {
	options := engine.Options{}

	flag.IntVar(&options.Width, "w", 800, "width in pixel")
	flag.IntVar(&options.Height, "h", 400, "height in pixel")
	flag.IntVar(&options.CPU, "cpu", runtime.NumCPU(), "number of CPU to use (default to number of CPU available)")
	flag.Int64Var(&options.Seed, "seed", 2017, "seed for random number generator")
	flag.IntVar(&options.RaysPerPixel, "r", 100, "number of rays per pixel")
	flag.StringVar(&options.Output, "o", "output.png", "path to file for saving (do not save if not defined)")

	flag.Parse()

	camera, world := engine.BuildWorldOneWeekend(options.Width, options.Height)

	scene := engine.NewScene(options.Width, options.Height, options.RaysPerPixel, camera, world)
	pixels, completed := scene.Render(options.CPU)

	<-completed
	fmt.Println("Render complete.")
	err, saved := engine.SaveImage(pixels, options)
	switch {
	case err != nil:
		fmt.Printf("Error while saving the image [%v]\n", err)
	case saved:
		fmt.Printf("Image saved to %v\n", options.Output)
	}
}
