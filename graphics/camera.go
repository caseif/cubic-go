package graphics

import (
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
    "github.com/go-gl/gl/v3.3-core/gl"
)

const Speed float32 = 2

var CAMERA = Camera{}

type Camera struct {
    Translation mgl32.Vec3
    Rotation mgl32.Vec3
    Velocity mgl32.Vec3
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

func (self *Camera) UpdatePosition() {
    self.Translation = self.Translation.Add(self.Velocity)
    gl.UniformMatrix4fv(TrMatLoc, 1, false, &self.GetTranslationMatrix()[0])
}
