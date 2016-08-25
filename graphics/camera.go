package graphics

import (
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
    "github.com/go-gl/gl/v3.3-core/gl"
    "math"
    "github.com/caseif/cubic-go/world"
    "fmt"
)

const Speed float32 = 0.02

var size = mgl32.Vec3{1, 2, 1}

var CAMERA = CreateCamera()

type Camera struct {
    Translation mgl32.Vec3
    Rotation mgl32.Vec3
    Velocity mgl32.Vec3
    dirtyTranslation, dirtyRotX, dirtyRotY, dirtyRotZ bool
    boundingBox util.BoundingBox
}

func CreateCamera() *Camera {
    camera := Camera{}
    camera.boundingBox = *util.CreateBoundingBox(&mgl32.Vec3{}, &size)
    return &camera
}

func (self *Camera) TranslateBy(translation mgl32.Vec3) {
    self.Translation = self.Translation.Sub(translation)
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
    self.boundingBox.SetPos(self.Translation.Sub(self.Velocity).Mul(unitLength))
    if doBoundChecks(&self.boundingBox) {
        self.boundingBox.SetPos(self.Translation.Mul(unitLength))
        self.Velocity = mgl32.Vec3{}
        return
    }
    self.TranslateBy(self.Velocity)
    self.boundingBox.SetPos(self.Translation.Mul(unitLength))
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

func doBoundChecks(box *util.BoundingBox) bool {
    minX := int32(math.Floor(float64(box.Pos.X() - box.Size.X() / 2)))
    maxX := int32(math.Floor(float64(box.Pos.X() + box.Size.X() / 2)))
    minY := int32(math.Floor(float64(box.Pos.Y() - box.Size.Y() / 2)))
    maxY := int32(math.Floor(float64(box.Pos.Y() + box.Size.Y() / 2)))
    minZ := int32(math.Floor(float64(box.Pos.Z() - box.Size.Z() / 2)))
    maxZ := int32(math.Floor(float64(box.Pos.Z() + box.Size.Z() / 2)))
    for x := minX; x <= maxX; x++ {
        for y := minY; y <= maxY; y++ {
            for z := minZ; z <= maxZ; z++ {
                block := world.GlobalWorldServer.GetWorld("world").GetBlock(x, y, z)
                if x == 0 && y == 0 && z == 0 {
                    fmt.Println("checking block")
                }
                if block != nil {
                    if block.GetBoundingBox().Collides(box) {
                        return true
                    }
                }
            }
        }
    }
    return false
}
