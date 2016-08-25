package util

import (
    "github.com/go-gl/mathgl/mgl32"
)

type BoundingBox struct {
    Pos mgl32.Vec3
    Size mgl32.Vec3
}

func CreateBoundingBox(pos, size *mgl32.Vec3) *BoundingBox {
    return &BoundingBox{Pos: *pos, Size: *size}
}

func (self *BoundingBox) SetPos(pos mgl32.Vec3) {
    self.Pos = pos
}

func (self *BoundingBox) Collides(other *BoundingBox) bool {
    return collides1D(self.Pos.X(), other.Pos.X(), self.Size.X(), other.Size.X()) &&
            collides1D(self.Pos.Y(), other.Pos.Y(), self.Size.Y(), other.Size.Y()) &&
            collides1D(self.Pos.Z(), other.Pos.Z(), self.Size.Z(), other.Size.Z())
}

func collides1D(x1, x2, d1, d2 float32) bool {
    d1 /= 2
    d2 /= 2
    x1 -= d1
    x2 -= d2
    return x1 < x2 + d2 && x1 + d1 > x2
}
