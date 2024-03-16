package engine

import "math"

type Interval struct {
	min, max float64
}

var (
	Empty    = Interval{math.MaxFloat64, -math.MaxFloat64}
	Universe = Interval{-math.MaxFloat64, math.MaxFloat64}
)

func (interval Interval) contains(t float64) bool {
	return interval.min <= t && t <= interval.max
}

func (interval Interval) surrounds(t float64) bool {
	return interval.min < t && t < interval.max
}

func (interval Interval) clamp(t float64) float64 {
	switch {
	case t < interval.min:
		return interval.min
	case t > interval.max:
		return interval.max
	default:
		return t
	}
}
