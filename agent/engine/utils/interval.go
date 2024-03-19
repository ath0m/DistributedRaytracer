package utils

import "math"

type Interval struct {
	Min, Max float64
}

var (
	Empty    = Interval{math.MaxFloat64, -math.MaxFloat64}
	Universe = Interval{-math.MaxFloat64, math.MaxFloat64}
)

func (interval Interval) Contains(t float64) bool {
	return interval.Min <= t && t <= interval.Max
}

func (interval Interval) Surrounds(t float64) bool {
	return interval.Min < t && t < interval.Max
}

func (interval Interval) Clamp(t float64) float64 {
	switch {
	case t < interval.Min:
		return interval.Min
	case t > interval.Max:
		return interval.Max
	default:
		return t
	}
}
