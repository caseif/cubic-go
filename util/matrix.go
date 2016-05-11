package util

import (
    "github.com/go-gl/mathgl/mgl32"
    "math"
)

func GetTranslationMatrix(translation mgl32.Vec3) *mgl32.Mat4 {
    mat := mgl32.Ident4()
    mat.Set(0, 3, translation.X())
    mat.Set(1, 3, translation.Y())
    mat.Set(2, 3, translation.Z())
    return &mat
}

func GetXRotationMatrix(rotationRads float32) *mgl32.Mat4 {
    var f64 float64 = float64(rotationRads)
    mat := mgl32.Ident4()
    cos := float32(math.Cos(f64))
    sin := float32(math.Sin(f64))
    mat.Set(1, 1, cos)
    mat.Set(2, 1, sin)
    mat.Set(1, 2, -sin)
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
    mat.Set(1, 0, -sin)
    mat.Set(0, 1, sin)
    mat.Set(1, 1, cos)
    return &mat
}

func Perspective(near, far, fovDegs, aspect float32) *mgl32.Mat4 {
    y2 := near * float32(math.Tan(float64(mgl32.DegToRad(fovDegs))))
    y1 := -y2
    x1 := y1 * aspect
    x2 := y2 * aspect
    return frustum(x1, x2, y1, y2, near, far)
}

func frustum(left, right, bottom, top, near, far float32) *mgl32.Mat4 {
    mat := mgl32.Mat4{}

    mat.Set(0, 0, (2 * near) / (right - left))
    mat.Set(1, 1, (2 * near) / (top - bottom))
    mat.Set(0, 2, (right + left) / (right - left))
    mat.Set(1, 2, (top + bottom) / (top - bottom))
    mat.Set(2, 2, (near + far) / (near - far))
    mat.Set(3, 2, (2 * near * far) / (near - far))
    mat.Set(2, 3, -1)
    mat.Set(3, 3, 0)

    return &mat
}
