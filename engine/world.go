package engine

import (
	"encoding/json"

	"github.com/ath0m/DistributedRaytracer/engine/camera"
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
