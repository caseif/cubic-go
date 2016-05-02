package util

import (
    "github.com/go-gl/mathgl/mgl32"
    "math"
)

func GetTranslationMatrix(translation mgl32.Vec3) *mgl32.Mat4 {
    mat := mgl32.Ident4()
    mat.Set(3, 0, translation.X())
    mat.Set(3, 1, translation.Y())
    mat.Set(3, 2, translation.Z())
    return &mat
}

func GetXRotationMatrix(rotationRads float32) *mgl32.Mat4 {
    var f64 float64 = float64(rotationRads)
    mat := mgl32.Ident4()
    cos := float32(math.Cos(f64))
    sin := float32(math.Sin(f64))
    mat.Set(1, 1, cos)
    mat.Set(1, 2, sin)
    mat.Set(2, 1, -sin)
    mat.Set(2, 2, cos)
    return &mat
}

func GetYRotationMatrix(rotationRads float32) *mgl32.Mat4 {
    var f64 float64 = float64(rotationRads)
    mat := mgl32.Ident4()
    cos := float32(math.Cos(f64))
    sin := float32(math.Sin(f64))
    mat.Set(0, 0, cos)
    mat.Set(2, 0, -sin)
    mat.Set(0, 2, sin)
    mat.Set(2, 2, cos)
    return &mat
}

func GetZRotationMatrix(rotationRads float32) *mgl32.Mat4 {
    var f64 float64 = float64(rotationRads)
    mat := mgl32.Ident4()
    cos := float32(math.Cos(f64))
    sin := float32(math.Sin(f64))
    mat.Set(0, 0, cos)
    mat.Set(0, 1, -sin)
    mat.Set(1, 0, sin)
    mat.Set(1, 1, cos)
    return &mat
}
