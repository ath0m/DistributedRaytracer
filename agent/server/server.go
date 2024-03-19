package server

import (
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"runtime"

	"github.com/ath0m/DistributedRaytracer/agent/engine"
)

type RenderOptions struct {
	Width        int   `json:"width"`        // width in pixel
	Height       int   `json:"height"`       // height in pixel
	RaysPerPixel int   `json:"raysperpixel"` // number of rays per pixel
	Seed         int64 `json:"seed"`         // seed for random number generator
}

func handleRender(w http.ResponseWriter, req *http.Request) {
	requestOptions := RenderOptions{
		Width:        800,
		Height:       400,
		RaysPerPixel: 10,
		Seed:         2024,
	}

	err := json.NewDecoder(req.Body).Decode(&requestOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	world, err := engine.LoadWorld("assets/world.json")
	if err != nil {
		msg := fmt.Sprintf("Error while loading world [%v]", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	scene := engine.NewScene(requestOptions.Width, requestOptions.Height, requestOptions.RaysPerPixel, world.Camera, world.Objects)
	pixels, completed := scene.Render(runtime.NumCPU())

	<-completed
	fmt.Println("Render complete.")

	img := engine.CreateImage(pixels, requestOptions.Width, requestOptions.Height)

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func Start() {
	http.HandleFunc("POST /render", handleRender)
	http.ListenAndServe(":8090", nil)
}
