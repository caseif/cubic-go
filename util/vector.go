package util

import "github.com/go-gl/mathgl/mgl32"

func VecAdd(vec *mgl32.Vec3, x, y, z float32) *mgl32.Vec3 {
    return &mgl32.Vec3{vec.X() + x, vec.Y() + y, vec.Z() + z}
}

func VecMul(vec *mgl32.Vec3,x, y, z float32) *mgl32.Vec3 {
    return &mgl32.Vec3{vec.X() * x, vec.Y() * y, vec.Z() * z}
}
