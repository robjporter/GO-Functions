package round

import (
	"math"
)

const Epsilon = 0.0000001

func Round(x float64) float64 {
	return ToNearestEven(x)
}

func RoundTo(x float64, dp float64) float64 {
	x = x * math.Pow(10, dp)
	return ToNearestEven(x) / math.Pow(10, dp)
}

func ToNearestEven(x float64) float64 {
	return toNearest(x, true)
}

func ToNearestAway(x float64) float64 {
	return toNearest(x, false)
}

func ToZero(x float64) float64 {
	return math.Trunc(x)
}

func AwayFromZero(x float64) float64 {
	if x >= 0 {
		return math.Ceil(x)
	} else {
		return math.Floor(x)
	}
}

func ToPositiveInf(x float64) float64 {
	return math.Ceil(x)
}

func ToNegativeInf(x float64) float64 {
	return math.Floor(x)
}

func toNearest(x float64, tiesToEven bool) float64 {
	if x == 0 || math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}
	if x < 0.0 {
		return -toNearest(-x, tiesToEven)
	}

	intPart, fracPart := math.Modf(x)

	if math.Abs(fracPart-0.5) < Epsilon {
		if tiesToEven {
			if math.Mod(intPart, 2.0) < Epsilon {
				return intPart
			}
		}
		return math.Ceil(intPart + 0.5)
	}
	return math.Floor(x + 0.5)
}
