package server

import (
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"runtime"

	"github.com/ath0m/DistributedRaytracer/agent/engine"
)

var defaultWorld engine.World

type RenderOptions struct {
	Width        int          `json:"width"`        // width in pixel
	Height       int          `json:"height"`       // height in pixel
	RaysPerPixel int          `json:"raysperpixel"` // number of rays per pixel
	Seed         int64        `json:"seed"`         // seed for random number generator
	World        engine.World `json:"world"`        // Optional world definition
}

func handleRender(w http.ResponseWriter, req *http.Request) {
	requestOptions := RenderOptions{
		Width:        800,
		Height:       400,
		RaysPerPixel: 10,
		Seed:         2024,
		World:        defaultWorld,
	}

	err := json.NewDecoder(req.Body).Decode(&requestOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	scene := engine.NewScene(requestOptions.Width, requestOptions.Height, requestOptions.RaysPerPixel, requestOptions.World.Camera, requestOptions.World.Objects)
	pixels, completed := scene.Render(runtime.NumCPU())

	<-completed
	fmt.Println("Render complete.")

	img := engine.CreateImage(pixels, requestOptions.Width, requestOptions.Height)

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func loadDefaultWorld() error {
	loaded, err := engine.LoadWorld("assets/world.json")
	if err != nil {
		return err
	} else {
		defaultWorld = *loaded
		return nil
	}
}

func Start() {
	err := loadDefaultWorld()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("POST /render", handleRender)

	fmt.Println("Server is starting on port 8090.")
	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server closed.")
}
