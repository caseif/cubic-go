package util

import (
    "github.com/go-gl/mathgl/mgl32"
    "strconv"
)

func Vec2ToStr(vec *mgl32.Vec2) string {
    return "(" + strconv.FormatFloat(float64(vec.X()), 'f', -1, 32) + ", " +
    strconv.FormatFloat(float64(vec.Y()), 'f', -1, 32) + ")"
}

func Vec3ToStr(vec *mgl32.Vec3) string {
    return "(" + strconv.FormatFloat(float64(vec.X()), 'f', -1, 32) + ", " +
    strconv.FormatFloat(float64(vec.Y()), 'f', -1, 32) + ", " +
    strconv.FormatFloat(float64(vec.Z()), 'f', -1, 32) + ")"
}
