// Simple ray tracer based on the Ray Tracing book series by Peter Shirley
package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ath0m/DistributedRaytracer/engine"
	"github.com/ath0m/DistributedRaytracer/utils"
)

func main() {
	options := utils.Options{}

	flag.IntVar(&options.Width, "w", 800, "width in pixel")
	flag.IntVar(&options.Height, "h", 400, "height in pixel")
	flag.IntVar(&options.CPU, "cpu", runtime.NumCPU(), "number of CPU to use (default to number of CPU available)")
	flag.Int64Var(&options.Seed, "seed", 2024, "seed for random number generator")
	flag.IntVar(&options.RaysPerPixel, "r", 100, "number of rays per pixel")
	flag.StringVar(&options.Output, "o", "output.png", "path to file for saving (do not save if not defined)")
	flag.StringVar(&options.Input, "i", "assets/world.json", "path to file for world definition in json format")

	flag.Parse()

	world, err := utils.LoadWorld(options.Input)
	if err != nil {
		panic(fmt.Sprintf("Error while loading world [%v]", err))
	}

	scene := engine.NewScene(options.Width, options.Height, options.RaysPerPixel, world.Camera, world.Objects)
	pixels, completed := scene.Render(options.CPU)

	<-completed
	fmt.Println("Render complete.")
	err, saved := utils.SaveImage(pixels, options)
	switch {
	case err != nil:
		fmt.Printf("Error while saving the image [%v]\n", err)
	case saved:
		fmt.Printf("Image saved to %v\n", options.Output)
	}
}
