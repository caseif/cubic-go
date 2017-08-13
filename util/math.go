package util

import "math"

const HalfPi = math.Pi / 2
const TwoPi = math.Pi * 2

func Cos(f float32) float32 {
    return float32(math.Cos(float64(f)))
}

func Sin(f float32) float32 {
    return float32(math.Sin(float64(f)))
}

func Clamp(val, min, max float32) float32 {
    return float32(math.Max(float64(min), math.Min(float64(max), float64(val))))
}

func DivFloor(n, d int) float32 {
    return float32(math.Floor(float64(n) / float64(d)))
}
