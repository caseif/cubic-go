package util

import "math"

func Cos(f float32) float32 {
    return float32(math.Cos(float64(f)))
}

func Sin(f float32) float32 {
    return float32(math.Sin(float64(f)))
}
