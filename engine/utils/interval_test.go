package utils

import "testing"

func TestIntervalContains(t *testing.T) {
	cases := []struct {
		interval Interval
		t        float64
		expected bool
	}{
		{Interval{0.0, 1.0}, 0.5, true},
		{Interval{0.0, 1.0}, 1.5, false},
		{Interval{-1.0, 1.0}, 0.0, true},
		{Interval{-1.0, 1.0}, -1.5, false},
	}

	for _, tc := range cases {
		result := tc.interval.Contains(tc.t)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestIntervalSurrounds(t *testing.T) {
	cases := []struct {
		interval Interval
		t        float64
		expected bool
	}{
		{Interval{0.0, 1.0}, 0.5, true},
		{Interval{0.0, 1.0}, 1.5, false},
		{Interval{-1.0, 1.0}, 0.0, true},
		{Interval{-1.0, 1.0}, -1.5, false},
	}

	for _, tc := range cases {
		result := tc.interval.Surrounds(tc.t)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestIntervalClamp(t *testing.T) {
	cases := []struct {
		interval Interval
		t        float64
		expected float64
	}{
		{Interval{0.0, 1.0}, -0.5, 0.0},
		{Interval{0.0, 1.0}, 0.5, 0.5},
		{Interval{0.0, 1.0}, 1.5, 1.0},
	}

	for _, tc := range cases {
		result := tc.interval.Clamp(tc.t)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
