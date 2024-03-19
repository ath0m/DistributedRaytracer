package engine

import (
	"encoding/json"
	"os"

	"github.com/ath0m/DistributedRaytracer/agent/engine/camera"
)

type World struct {
	Camera  camera.Camera `json:"camera"`
	Objects HittableList  `json:"objects"`
}

func (w *World) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Camera  json.RawMessage `json:"camera"`
		Objects []Sphere        `json:"objects"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	w.Camera = camera.UnmarshalCamera(aux.Camera)

	// Convert []Sphere to HittableList
	w.Objects = HittableList{}
	for _, obj := range aux.Objects {
		w.Objects = append(w.Objects, obj)
	}

	return nil
}

func LoadWorld(file string) (*World, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var world World
	err = json.Unmarshal(content, &world)
	if err != nil {
		return nil, err
	}

	return &world, nil
}
