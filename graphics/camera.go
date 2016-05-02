package graphics

import (
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
)

var CAMERA = Camera{}

type Camera struct {
    Translation mgl32.Vec3
    Rotation mgl32.Vec3
}

func (self Camera) GetTranslationMatrix() *mgl32.Mat4 {
    return util.GetTranslationMatrix(self.Translation)
}

func (self Camera) GetXRotationMatrix() *mgl32.Mat4 {
    return util.GetXRotationMatrix(self.Rotation.X())
}

func (self Camera) GetYRotationMatrix() *mgl32.Mat4 {
    return util.GetYRotationMatrix(self.Rotation.Y())
}

func (self Camera) GetZRotationMatrix() *mgl32.Mat4 {
    return util.GetZRotationMatrix(self.Rotation.Z())
}
