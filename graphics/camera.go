package graphics

import (
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
    "github.com/go-gl/gl/v3.3-core/gl"
    "math"
)

const Speed float32 = 0.02

var CAMERA = Camera{}

type Camera struct {
    Translation mgl32.Vec3
    Rotation mgl32.Vec3
    Velocity mgl32.Vec3
    dirtyTranslation, dirtyRotX, dirtyRotY, dirtyRotZ bool
}

func (self *Camera) TranslateBy(translation mgl32.Vec3) {
    self.Translation = self.Translation.Add(translation)
    self.dirtyTranslation = true
}

func (self *Camera) RotateBy(pitch, yaw float32) {
    newPitch := util.Clamp(self.Rotation.X() + pitch, -util.HalfPi, util.HalfPi)
    newYaw := float32(math.Mod(float64(self.Rotation.Y() + yaw), util.TwoPi))
    if newYaw < 0 {
        newYaw += util.TwoPi
    }
    self.Rotation = mgl32.Vec3{newPitch, newYaw, 0}
    self.dirtyRotX = pitch != 0
    self.dirtyRotY = yaw != 0
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
    self.TranslateBy(self.Velocity)
    if self.dirtyTranslation {
        gl.UniformMatrix4fv(TrMatLoc, 1, false, &self.GetTranslationMatrix()[0])
        self.dirtyTranslation = false
    }
    if self.dirtyRotX {
        gl.UniformMatrix4fv(RotXMatLoc, 1, false, &self.GetXRotationMatrix()[0])
        self.dirtyRotX = false
    }
    if self.dirtyRotY {
        gl.UniformMatrix4fv(RotYMatLoc, 1, false, &self.GetYRotationMatrix()[0])
        self.dirtyRotY = false
    }
    if self.dirtyRotZ {
        gl.UniformMatrix4fv(RotZMatLoc, 1, false, &self.GetZRotationMatrix()[0])
        self.dirtyRotZ = false
    }
}
