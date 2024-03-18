package geometry

import (
	"testing"
)

func TestRayPointAt(t *testing.T) {
	cases := []struct {
		r        Ray
		t        float64
		expected Point3
	}{
		{Ray{Origin: Point3{0.0, 0.0, 0.0}, Direction: Vec3{1.0, 0.0, 0.0}}, 2.0, Point3{2.0, 0.0, 0.0}},
		{Ray{Origin: Point3{1.0, 2.0, 3.0}, Direction: Vec3{0.0, 1.0, 0.0}}, 3.0, Point3{1.0, 5.0, 3.0}},
		{Ray{Origin: Point3{-1.0, -2.0, -3.0}, Direction: Vec3{0.0, 0.0, 1.0}}, 4.0, Point3{-1.0, -2.0, 1.0}},
	}

	for _, tc := range cases {
		result := tc.r.PointAt(tc.t)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
