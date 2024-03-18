package geometry

import (
	"testing"
)

func TestPoint3Translate(t *testing.T) {
	cases := []struct {
		p        Point3
		v        Vec3
		expected Point3
	}{
		{Point3{1.0, 2.0, 3.0}, Vec3{2.0, 3.0, 4.0}, Point3{3.0, 5.0, 7.0}},
		{Point3{0.0, 0.0, 0.0}, Vec3{2.0, 3.0, 4.0}, Point3{2.0, 3.0, 4.0}},
		{Point3{1.0, 2.0, 3.0}, Vec3{0.0, 0.0, 0.0}, Point3{1.0, 2.0, 3.0}},
	}

	for _, tc := range cases {
		result := tc.p.Translate(tc.v)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestPoint3Sub(t *testing.T) {
	cases := []struct {
		p1       Point3
		p2       Point3
		expected Vec3
	}{
		{Point3{1.0, 2.0, 3.0}, Point3{2.0, 3.0, 4.0}, Vec3{-1.0, -1.0, -1.0}},
		{Point3{0.0, 0.0, 0.0}, Point3{2.0, 3.0, 4.0}, Vec3{-2.0, -3.0, -4.0}},
		{Point3{1.0, 2.0, 3.0}, Point3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}},
	}

	for _, tc := range cases {
		result := tc.p1.Sub(tc.p2)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestPoint3Vec3(t *testing.T) {
	cases := []struct {
		p        Point3
		expected Vec3
	}{
		{Point3{1.0, 2.0, 3.0}, Vec3{1.0, 2.0, 3.0}},
		{Point3{0.0, 0.0, 0.0}, Vec3{0.0, 0.0, 0.0}},
		{Point3{-1.0, -2.0, -3.0}, Vec3{-1.0, -2.0, -3.0}},
	}

	for _, tc := range cases {
		result := tc.p.Vec3()
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
