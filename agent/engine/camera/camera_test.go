package camera

import (
	"reflect"
	"testing"

	"github.com/ath0m/DistributedRaytracer/agent/engine/geometry"
)

type mockRnd struct{}

func (rnd mockRnd) Float64() float64 {
	return 0.5
}

func TestCameraRay(t *testing.T) {
	c := camera{
		Origin:          geometry.Point3{X: 0.0, Y: 0.0, Z: 0.0},
		LowerLeftCorner: geometry.Point3{X: -1.0, Y: -1.0, Z: -1.0},
		Horizontal:      geometry.Vec3{X: 2.0, Y: 0.0, Z: 0.0},
		Vertical:        geometry.Vec3{X: 0.0, Y: 2.0, Z: 0.0},
		LensRadius:      0.0,
		U:               geometry.Vec3{X: 1.0, Y: 0.0, Z: 0.0},
		V:               geometry.Vec3{X: 0.0, Y: 1.0, Z: 0.0},
	}

	rnd := mockRnd{}
	u := 0.5
	v := 0.5

	ray := c.Ray(rnd, u, v)

	expectedOrigin := geometry.Point3{X: 0.0, Y: 0.0, Z: 0.0}
	expectedDirection := geometry.Vec3{X: 0.0, Y: 0.0, Z: -1.0}

	if ray.Origin != expectedOrigin {
		t.Errorf("Expected origin %v, but got %v", expectedOrigin, ray.Origin)
	}

	if ray.Direction != expectedDirection {
		t.Errorf("Expected direction %v, but got %v", expectedDirection, ray.Direction)
	}
}
func TestUnmarshalCamera(t *testing.T) {
	data := []byte(`{"Origin":{"X":0.0,"Y":0.0,"Z":0.0},"LowerLeftCorner":{"X":-1.0,"Y":-1.0,"Z":-1.0},"Horizontal":{"X":2.0,"Y":0.0,"Z":0.0},"Vertical":{"X":0.0,"Y":2.0,"Z":0.0},"LensRadius":0.0,"U":{"X":1.0,"Y":0.0,"Z":0.0},"V":{"X":0.0,"Y":1.0,"Z":0.0}}`)
	expectedCamera := camera{
		Origin:          geometry.Point3{X: 0.0, Y: 0.0, Z: 0.0},
		LowerLeftCorner: geometry.Point3{X: -1.0, Y: -1.0, Z: -1.0},
		Horizontal:      geometry.Vec3{X: 2.0, Y: 0.0, Z: 0.0},
		Vertical:        geometry.Vec3{X: 0.0, Y: 2.0, Z: 0.0},
		LensRadius:      0.0,
		U:               geometry.Vec3{X: 1.0, Y: 0.0, Z: 0.0},
		V:               geometry.Vec3{X: 0.0, Y: 1.0, Z: 0.0},
	}

	result := UnmarshalCamera(data)

	if !reflect.DeepEqual(result, expectedCamera) {
		t.Errorf("Expected camera %v, but got %v", expectedCamera, result)
	}
}
