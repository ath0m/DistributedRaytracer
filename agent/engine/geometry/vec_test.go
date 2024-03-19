package geometry

import (
	"math"
	"testing"
)

func equalVec3(v1, v2 Vec3) bool {
	const eps = 1e-9
	return math.Abs(v1.X-v2.X) < eps && math.Abs(v1.Y-v2.Y) < eps && math.Abs(v1.Z-v2.Z) < eps
}

func TestVec3Scale(t *testing.T) {
	cases := []struct {
		v        Vec3
		t        float64
		expected Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, 2.0, Vec3{2.0, 4.0, 6.0}},
		{Vec3{0.0, 0.0, 0.0}, 2.0, Vec3{0.0, 0.0, 0.0}},
		{Vec3{1.0, 2.0, 3.0}, 0.0, Vec3{0.0, 0.0, 0.0}},
	}

	for _, tc := range cases {
		result := tc.v.Scale(tc.t)
		if !equalVec3(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3Mult(t *testing.T) {
	cases := []struct {
		v1, v2, expected Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{1.0, 2.0, 3.0}, Vec3{1.0, 4.0, 9.0}},
		{Vec3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}, Vec3{0.0, 0.0, 0.0}},
	}

	for _, tc := range cases {
		result := tc.v1.Mult(tc.v2)
		if !equalVec3(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3Sub(t *testing.T) {
	cases := []struct {
		v1, v2, expected Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{1.0, 2.0, 3.0}, Vec3{0.0, 0.0, 0.0}},
		{Vec3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}, Vec3{-1.0, -2.0, -3.0}},
	}

	for _, tc := range cases {
		result := tc.v1.Sub(tc.v2)
		if !equalVec3(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3Add(t *testing.T) {
	cases := []struct {
		v1, v2, expected Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{1.0, 2.0, 3.0}, Vec3{2.0, 4.0, 6.0}},
		{Vec3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}, Vec3{1.0, 2.0, 3.0}},
	}

	for _, tc := range cases {
		result := tc.v1.Add(tc.v2)
		if !equalVec3(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3Length(t *testing.T) {
	cases := []struct {
		v        Vec3
		expected float64
	}{
		{Vec3{1.0, 2.0, 2.0}, 3.0},
		{Vec3{0.0, 0.0, 0.0}, 0.0},
	}

	for _, tc := range cases {
		result := tc.v.Length()
		if math.Abs(result-tc.expected) > 1e-9 {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3Unit(t *testing.T) {
	cases := []struct {
		v, expected Vec3
	}{
		{Vec3{1.0, 2.0, 2.0}, Vec3{1.0 / 3.0, 2.0 / 3.0, 2.0 / 3.0}},
		{Vec3{1.0, 0.0, 0.0}, Vec3{1.0, 0.0, 0.0}},
	}

	for _, tc := range cases {
		result := tc.v.Unit()
		if math.Abs(result.Length()-1.0) > 1e-9 {
			t.Errorf("Expected length 1, but got %v", result.Length())
		}
		if math.Abs(result.X-tc.expected.X) > 1e-9 || math.Abs(result.Y-tc.expected.Y) > 1e-9 || math.Abs(result.Z-tc.expected.Z) > 1e-9 {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3Negate(t *testing.T) {
	cases := []struct {
		v, expected Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{-1.0, -2.0, -3.0}},
		{Vec3{-1.0, -2.0, -3.0}, Vec3{1.0, 2.0, 3.0}},
		{Vec3{0.0, 0.0, 0.0}, Vec3{0.0, 0.0, 0.0}},
	}

	for _, tc := range cases {
		result := tc.v.Negate()
		if !equalVec3(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3NearZero(t *testing.T) {
	cases := []struct {
		v        Vec3
		expected bool
	}{
		{Vec3{1e-9, 1e-9, 1e-9}, true},
		{Vec3{1e-7, 1e-7, 1e-7}, false},
		{Vec3{0.0, 0.0, 0.0}, true},
	}

	for _, tc := range cases {
		result := tc.v.NearZero()
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestDot(t *testing.T) {
	cases := []struct {
		v1, v2   Vec3
		expected float64
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{4.0, 5.0, 6.0}, 32.0},
		{Vec3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}, 0.0},
	}

	for _, tc := range cases {
		result := Dot(tc.v1, tc.v2)
		if math.Abs(result-tc.expected) > 1e-9 {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestCross(t *testing.T) {
	cases := []struct {
		v1, v2, expected Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{4.0, 5.0, 6.0}, Vec3{-3.0, 6.0, -3.0}},
		{Vec3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}, Vec3{0.0, 0.0, 0.0}},
	}

	for _, tc := range cases {
		result := Cross(tc.v1, tc.v2)
		if !equalVec3(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
func TestVec3Reflect(t *testing.T) {
	cases := []struct {
		v, n, expected Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{1.0, 0.0, 0.0}, Vec3{-1.0, 2.0, 3.0}},
		{Vec3{1.0, 2.0, 3.0}, Vec3{0.0, 1.0, 0.0}, Vec3{1.0, -2.0, 3.0}},
		{Vec3{1.0, 2.0, 3.0}, Vec3{0.0, 0.0, 1.0}, Vec3{1.0, 2.0, -3.0}},
	}

	for _, tc := range cases {
		result := tc.v.Reflect(tc.n)
		if !equalVec3(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestVec3Refract(t *testing.T) {
	cases := []struct {
		v           Vec3
		n           Vec3
		niOverNt    float64
		expectedHit bool
		expectedVec *Vec3
	}{
		{Vec3{1.0, 2.0, 3.0}, Vec3{4.0, 5.0, 6.0}, 2.0, true, &Vec3{-0.7616575603770357, -0.5511800876026581, -0.3407026148282802}},
		{Vec3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}, 2.0, false, nil},
	}

	for _, tc := range cases {
		hit, result := tc.v.Refract(tc.n, tc.niOverNt)
		if hit != tc.expectedHit {
			t.Errorf("Expected hit %v, but got %v", tc.expectedHit, hit)
		}
		if hit && !equalVec3(*result, *tc.expectedVec) {
			t.Errorf("Expected %v, but got %v", *tc.expectedVec, *result)
		}
	}
}
