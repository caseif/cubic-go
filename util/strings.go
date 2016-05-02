package util

import (
    "github.com/go-gl/mathgl/mgl32"
    "strconv"
)

func Vec2ToStr(vec *mgl32.Vec2) string {
    return "(" + strconv.Itoa(int(vec.X())) + ", " + strconv.Itoa(int(vec.Y())) + ")"
}

func Vec3ToStr(vec *mgl32.Vec3) string {
    return "(" + strconv.Itoa(int(vec.X())) + ", " + strconv.Itoa(int(vec.Y())) + ", " + strconv.Itoa(int(vec.Z())) +
    ")"
}
